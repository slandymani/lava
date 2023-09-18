package badgegenerator

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"

	"google.golang.org/grpc/metadata"

	btcSecp256k1 "github.com/btcsuite/btcd/btcec"
	"github.com/lavanet/lava/protocol/badgegenerator/grpc"
	"github.com/lavanet/lava/protocol/lavasession"
	"github.com/lavanet/lava/utils"
	"github.com/lavanet/lava/utils/sigs"
	pairingtypes "github.com/lavanet/lava/x/pairing/types"
	spectypes "github.com/lavanet/lava/x/spec/types"
)

const dummyApiInterface = "badgeApiInterface"

type Server struct {
	pairingtypes.UnimplementedBadgeGeneratorServer
	ProjectsConfiguration map[string]map[string]*ProjectConfiguration // geolocation/project_id/project_data
	epoch                 uint64
	virtualEpoch          uint64
	grpcFetcher           *grpc.GRPCFetcher
	ChainId               string
	IpService             *IpService
	metrics               *MetricsService
	stateTracker          *BadgeStateTracker
	specs                 map[string]spectypes.Spec // holding the specs for all chains
	specLock              sync.RWMutex
}

func NewServer(ipService *IpService, grpcUrl, chainId, userData string) (*Server, error) {
	server := &Server{
		ProjectsConfiguration: map[string]map[string]*ProjectConfiguration{},
		ChainId:               chainId,
		IpService:             ipService,
		specs:                 map[string]spectypes.Spec{},
	}

	if userData != "" {
		projectsData := make(map[string]map[string]*ProjectConfiguration)
		err := json.Unmarshal([]byte(userData), &projectsData)
		if err != nil {
			utils.LavaFormatWarning("provided information: ", err, utils.Attribute{Key: "userData", Value: userData})
			return nil, err
		}
		server.ProjectsConfiguration = projectsData
	}
	grpcFetch, err := grpc.NewGRPCFetcher(grpcUrl)
	if err != nil {
		return nil, err
	}
	server.grpcFetcher = grpcFetch
	server.metrics = InitMetrics()
	return server, nil
	//
}

func (s *Server) InitializeStateTracker(tracker *BadgeStateTracker) {
	if s.stateTracker != nil {
		utils.LavaFormatFatal("state tracker already initialized", nil)
	}
	s.stateTracker = tracker
}

func (s *Server) SetSpec(specUpdate spectypes.Spec) {
	s.specLock.Lock()
	defer s.specLock.Unlock()
	s.specs[specUpdate.Index] = specUpdate
}

func (s *Server) UpdateEpoch(epoch uint64) {
	utils.LavaFormatDebug("Got epoch update", utils.Attribute{Key: "epoch", Value: epoch})
	atomic.StoreUint64(&s.epoch, epoch)
	// reset virtual epoch after the end of emergency mode
	atomic.StoreUint64(&s.virtualEpoch, 0)
}

func (s *Server) UpdateVirtualEpoch(epoch uint64, virtualEpoch uint64) {
	utils.LavaFormatDebug("Got virtual epoch update", utils.Attribute{Key: "virtual_epoch", Value: virtualEpoch})
	atomic.StoreUint64(&s.virtualEpoch, virtualEpoch)
}

func (s *Server) GetEpoch() uint64 {
	return atomic.LoadUint64(&s.epoch)
}

func (s *Server) GetVirtualEpoch() uint64 {
	return atomic.LoadUint64(&s.virtualEpoch)
}

func (s *Server) checkSpecExists(specID string) (spectypes.Spec, bool) {
	s.specLock.RLock()
	defer s.specLock.RUnlock()
	spec, found := s.specs[specID]
	return spec, found
}

func (s *Server) getSpec(ctx context.Context, specId string) (spectypes.Spec, error) {
	_, found := s.checkSpecExists(specId)
	if !found {
		err := s.stateTracker.RegisterForSpecUpdates(ctx, s, lavasession.RPCEndpoint{ChainID: specId, ApiInterface: dummyApiInterface})
		if err != nil {
			return spectypes.Spec{}, utils.LavaFormatError("BadgeServer Failed registering for spec updates", err)
		}
	}
	// we should have the spec now after fetching it from the chain. if we don't have it badge server failed getting the spec
	spec, found := s.checkSpecExists(specId)
	if !found {
		return spectypes.Spec{}, utils.LavaFormatError("Failed fetching spec without getting error, shouldn't get here", nil)
	}
	return spec, nil
}

func (s *Server) GenerateBadge(ctx context.Context, req *pairingtypes.GenerateBadgeRequest) (*pairingtypes.GenerateBadgeResponse, error) {
	spec, err := s.getSpec(ctx, req.SpecId)
	if err != nil {
		return nil, utils.LavaFormatError("badge server failed fetching spec", err)
	}
	metadata, _ := metadata.FromIncomingContext(ctx)
	clientAddress := metadata.Get(RefererHeaderKey)
	ipAddress := ""
	if len(clientAddress) > 0 {
		ipAddress = clientAddress[0]
	}
	projectData, err := s.validateRequest(ipAddress, req)
	if err != nil {
		s.metrics.AddRequest(false)
		return nil, err
	}
	badge := pairingtypes.Badge{
		CuAllocation: uint64(projectData.EpochsMaxCu) * (s.GetVirtualEpoch() + 1), // add additional CU due to emergency mode
		Epoch:        s.GetEpoch(),
		Address:      req.BadgeAddress,
		LavaChainId:  s.ChainId,
	}

	result := pairingtypes.GenerateBadgeResponse{
		Badge:              &badge,
		BadgeSignerAddress: projectData.ProjectPublicKey,
		Spec:               &spec,
	}

	err = s.addPairingListToResponse(req, projectData, &result)
	if err != nil {
		s.metrics.AddRequest(false)
		return nil, err
	}

	err = signTheResponse(projectData.ProjectPrivateKey, &result)
	if err != nil {
		s.metrics.AddRequest(false)
		return nil, err
	}
	s.metrics.AddRequest(true)
	return &result, nil
}

func (s *Server) validateRequest(clientAddress string, in *pairingtypes.GenerateBadgeRequest) (*ProjectConfiguration, error) {
	if in == nil {
		err := fmt.Errorf("invalid request, no input data provided")
		utils.LavaFormatError("Validation failed", err)
		return nil, err
	}
	if in.BadgeAddress == "" || in.ProjectId == "" {
		fmt.Println("In: ", in)
		err := fmt.Errorf("bad request, no valid input data provided")
		utils.LavaFormatError("Validation failed", err)
		return nil, err
	}
	geolocation := s.getClientGeolocationOrDefault(clientAddress)
	geolocationData, exist := s.ProjectsConfiguration[geolocation]
	if !exist {
		err := fmt.Errorf("invalid configuration for this geolocation")
		utils.LavaFormatError(
			"invalid configuration",
			err,
			utils.Attribute{
				Key:   "BadgeAddress",
				Value: in.BadgeAddress,
			}, utils.Attribute{
				Key:   "ProjectId",
				Value: in.ProjectId,
			},
			utils.Attribute{
				Key:   "geolocation",
				Value: geolocation,
			},
			utils.Attribute{
				Key:   "ip",
				Value: clientAddress,
			},
		)
		return nil, err
	}
	projectData, exist := geolocationData[in.ProjectId]
	if !exist {
		projectData, exist = geolocationData[DefaultProjectId]
		if !exist {
			err := fmt.Errorf("default project not found")
			utils.LavaFormatError(
				"Validation failed",
				err,
				utils.Attribute{
					Key:   "BadgeAddress",
					Value: in.BadgeAddress,
				}, utils.Attribute{
					Key:   "ProjectId",
					Value: in.ProjectId,
				},
				utils.Attribute{
					Key:   "geolocation",
					Value: geolocation,
				},
			)
			return nil, err
		}
	}
	return projectData, nil
}

func (s *Server) getClientGeolocationOrDefault(clientIpAddress string) string {
	if s.IpService != nil && len(clientIpAddress) > 0 {
		utils.LavaFormatDebug("searching for ip", utils.Attribute{
			Key:   "clientIp",
			Value: clientIpAddress,
		})
		ip, err := s.IpService.SearchForIp(clientIpAddress)
		if err != nil {
			utils.LavaFormatError("error searching for client ip-geolocation", err)
		} else if ip == nil {
			utils.LavaFormatInfo("ip not found")
		} else {
			return fmt.Sprintf("%d", ip.Geolocation)
		}
	} else {
		utils.LavaFormatInfo("Ip service not configured correctly")
	}
	return fmt.Sprintf("%d", s.IpService.DefaultGeolocation)
}

func (s *Server) addPairingListToResponse(request *pairingtypes.GenerateBadgeRequest, configurations *ProjectConfiguration, response *pairingtypes.GenerateBadgeResponse) error {
	if request.SpecId != "" {
		if configurations.PairingList == nil || response.Badge.Epoch != configurations.UpdatedEpoch {
			pairings, err := s.grpcFetcher.FetchPairings(request.SpecId, configurations.ProjectPublicKey)
			if err != nil {
				utils.LavaFormatError("Failed to get pairings", err,
					utils.Attribute{Key: "epoch", Value: s.GetEpoch()},
					utils.Attribute{Key: "BadgeAddress", Value: request.GetBadgeAddress()},
					utils.Attribute{Key: "ProjectId", Value: request.ProjectId})
				return err
			}
			configurations.PairingList = pairings
			configurations.UpdatedEpoch = response.Badge.Epoch
		}
		// return the pairingList we have stored in configurations
		response.GetPairingResponse = configurations.PairingList
	}
	return nil
}

// note this update the signature of the response
func signTheResponse(privateKeyString string, response *pairingtypes.GenerateBadgeResponse) error {
	privateKeyBytes, _ := hex.DecodeString(privateKeyString)
	privateKey, _ := btcSecp256k1.PrivKeyFromBytes(btcSecp256k1.S256(), privateKeyBytes)
	signature, err := sigs.Sign(privateKey, *response.Badge)
	if err != nil {
		return err
	}

	response.Badge.ProjectSig = signature
	return nil
}
