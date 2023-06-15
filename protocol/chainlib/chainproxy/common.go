package chainproxy

import (
	"encoding/json"
	"fmt"

	"github.com/lavanet/lava/protocol/parser"
	pairingtypes "github.com/lavanet/lava/x/pairing/types"
	spectypes "github.com/lavanet/lava/x/spec/types"
)

const (
	LavaErrorCode       = 555
	InternalErrorString = "Internal Error"
)

type CustomParsingMessage interface {
	NewParsableRPCInput(input json.RawMessage) (parser.RPCInput, error)
}

type BaseMessage struct {
	Headers                 []pairingtypes.Metadata
	LatestBlockHeaderSetter *spectypes.ParseDirective
}

func (bm *BaseMessage) SetLatestBlockWithHeader(latestBlock uint64, modifyContent bool) (done bool) {
	if bm.LatestBlockHeaderSetter == nil {
		return false
	}
	headerValue := fmt.Sprintf(bm.LatestBlockHeaderSetter.FunctionTemplate, latestBlock)
	for idx, header := range bm.Headers {
		if header.Name == bm.LatestBlockHeaderSetter.ApiName {
			if modifyContent {
				bm.Headers[idx].Value = headerValue
			}
			return true
		}
	}
	if modifyContent {
		bm.Headers = append(bm.Headers, pairingtypes.Metadata{
			Name:  bm.LatestBlockHeaderSetter.ApiName,
			Value: headerValue,
		})
	}
	return true
}

func (bm BaseMessage) GetHeaders() []pairingtypes.Metadata {
	return bm.Headers
}

type DefaultRPCInput struct {
	Result json.RawMessage
	BaseMessage
}

func (dri DefaultRPCInput) GetParams() interface{} {
	return nil
}

func (dri DefaultRPCInput) GetResult() json.RawMessage {
	return dri.Result
}

func (dri DefaultRPCInput) ParseBlock(inp string) (int64, error) {
	return parser.ParseDefaultBlockParameter(inp)
}

func DefaultParsableRPCInput(input json.RawMessage) parser.RPCInput {
	return DefaultRPCInput{Result: input}
}
