// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/plans/policy.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/lavanet/lava/x/spec/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// the enum below determines the pairing algorithm's behaviour with the selected providers feature
type SELECTED_PROVIDERS_MODE int32

const (
	SELECTED_PROVIDERS_MODE_ALLOWED   SELECTED_PROVIDERS_MODE = 0
	SELECTED_PROVIDERS_MODE_MIXED     SELECTED_PROVIDERS_MODE = 1
	SELECTED_PROVIDERS_MODE_EXCLUSIVE SELECTED_PROVIDERS_MODE = 2
	SELECTED_PROVIDERS_MODE_DISABLED  SELECTED_PROVIDERS_MODE = 3
)

var SELECTED_PROVIDERS_MODE_name = map[int32]string{
	0: "ALLOWED",
	1: "MIXED",
	2: "EXCLUSIVE",
	3: "DISABLED",
}

var SELECTED_PROVIDERS_MODE_value = map[string]int32{
	"ALLOWED":   0,
	"MIXED":     1,
	"EXCLUSIVE": 2,
	"DISABLED":  3,
}

func (x SELECTED_PROVIDERS_MODE) String() string {
	return proto.EnumName(SELECTED_PROVIDERS_MODE_name, int32(x))
}

func (SELECTED_PROVIDERS_MODE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c2388e0faa8deb9b, []int{0}
}

// protobuf expected in YAML format: used "moretags" to simplify parsing
type Policy struct {
	ChainPolicies         []ChainPolicy           `protobuf:"bytes,1,rep,name=chain_policies,json=chainPolicies,proto3" json:"chain_policies"`
	GeolocationProfile    uint64                  `protobuf:"varint,2,opt,name=geolocation_profile,json=geolocationProfile,proto3" json:"geolocation_profile"`
	TotalCuLimit          uint64                  `protobuf:"varint,3,opt,name=total_cu_limit,json=totalCuLimit,proto3" json:"total_cu_limit"`
	EpochCuLimit          uint64                  `protobuf:"varint,4,opt,name=epoch_cu_limit,json=epochCuLimit,proto3" json:"epoch_cu_limit"`
	MaxProvidersToPair    uint64                  `protobuf:"varint,5,opt,name=max_providers_to_pair,json=maxProvidersToPair,proto3" json:"max_providers_to_pair"`
	SelectedProvidersMode SELECTED_PROVIDERS_MODE `protobuf:"varint,6,opt,name=selected_providers_mode,json=selectedProvidersMode,proto3,enum=lavanet.lava.plans.SELECTED_PROVIDERS_MODE" json:"selected_providers_mode"`
	SelectedProviders     []string                `protobuf:"bytes,7,rep,name=selected_providers,json=selectedProviders,proto3" json:"selected_providers"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2388e0faa8deb9b, []int{0}
}
func (m *Policy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return m.Size()
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetChainPolicies() []ChainPolicy {
	if m != nil {
		return m.ChainPolicies
	}
	return nil
}

func (m *Policy) GetGeolocationProfile() uint64 {
	if m != nil {
		return m.GeolocationProfile
	}
	return 0
}

func (m *Policy) GetTotalCuLimit() uint64 {
	if m != nil {
		return m.TotalCuLimit
	}
	return 0
}

func (m *Policy) GetEpochCuLimit() uint64 {
	if m != nil {
		return m.EpochCuLimit
	}
	return 0
}

func (m *Policy) GetMaxProvidersToPair() uint64 {
	if m != nil {
		return m.MaxProvidersToPair
	}
	return 0
}

func (m *Policy) GetSelectedProvidersMode() SELECTED_PROVIDERS_MODE {
	if m != nil {
		return m.SelectedProvidersMode
	}
	return SELECTED_PROVIDERS_MODE_ALLOWED
}

func (m *Policy) GetSelectedProviders() []string {
	if m != nil {
		return m.SelectedProviders
	}
	return nil
}

type ChainPolicy struct {
	ChainId      string             `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id"`
	Apis         []string           `protobuf:"bytes,2,rep,name=apis,proto3" json:"apis"`
	Requirements []ChainRequirement `protobuf:"bytes,3,rep,name=requirements,proto3" json:"requirements"`
}

func (m *ChainPolicy) Reset()         { *m = ChainPolicy{} }
func (m *ChainPolicy) String() string { return proto.CompactTextString(m) }
func (*ChainPolicy) ProtoMessage()    {}
func (*ChainPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2388e0faa8deb9b, []int{1}
}
func (m *ChainPolicy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainPolicy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainPolicy.Merge(m, src)
}
func (m *ChainPolicy) XXX_Size() int {
	return m.Size()
}
func (m *ChainPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_ChainPolicy proto.InternalMessageInfo

func (m *ChainPolicy) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *ChainPolicy) GetApis() []string {
	if m != nil {
		return m.Apis
	}
	return nil
}

func (m *ChainPolicy) GetRequirements() []ChainRequirement {
	if m != nil {
		return m.Requirements
	}
	return nil
}

type ChainRequirement struct {
	Collection types.CollectionData `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection" mapstructure:"collection"`
	Extensions []string             `protobuf:"bytes,2,rep,name=extensions,proto3" json:"extensions" mapstructure:"extensions"`
	Mixed      bool                 `protobuf:"varint,3,opt,name=mixed,proto3" json:"mixed" mapstructure:"mixed"`
}

func (m *ChainRequirement) Reset()         { *m = ChainRequirement{} }
func (m *ChainRequirement) String() string { return proto.CompactTextString(m) }
func (*ChainRequirement) ProtoMessage()    {}
func (*ChainRequirement) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2388e0faa8deb9b, []int{2}
}
func (m *ChainRequirement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainRequirement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainRequirement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainRequirement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainRequirement.Merge(m, src)
}
func (m *ChainRequirement) XXX_Size() int {
	return m.Size()
}
func (m *ChainRequirement) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainRequirement.DiscardUnknown(m)
}

var xxx_messageInfo_ChainRequirement proto.InternalMessageInfo

func (m *ChainRequirement) GetCollection() types.CollectionData {
	if m != nil {
		return m.Collection
	}
	return types.CollectionData{}
}

func (m *ChainRequirement) GetExtensions() []string {
	if m != nil {
		return m.Extensions
	}
	return nil
}

func (m *ChainRequirement) GetMixed() bool {
	if m != nil {
		return m.Mixed
	}
	return false
}

func init() {
	proto.RegisterEnum("lavanet.lava.plans.SELECTED_PROVIDERS_MODE", SELECTED_PROVIDERS_MODE_name, SELECTED_PROVIDERS_MODE_value)
	proto.RegisterType((*Policy)(nil), "lavanet.lava.plans.Policy")
	proto.RegisterType((*ChainPolicy)(nil), "lavanet.lava.plans.ChainPolicy")
	proto.RegisterType((*ChainRequirement)(nil), "lavanet.lava.plans.ChainRequirement")
}

func init() { proto.RegisterFile("lavanet/lava/plans/policy.proto", fileDescriptor_c2388e0faa8deb9b) }

var fileDescriptor_c2388e0faa8deb9b = []byte{
	// 708 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xcd, 0x4e, 0xdb, 0x40,
	0x10, 0x8e, 0x93, 0x00, 0xc9, 0x26, 0xa0, 0x74, 0xcb, 0x8f, 0xa1, 0x95, 0x1d, 0xa2, 0xfe, 0x44,
	0x45, 0x8a, 0x05, 0x3d, 0xb4, 0xea, 0x0d, 0xc7, 0x96, 0x1a, 0x35, 0x94, 0x68, 0x43, 0x29, 0xea,
	0x01, 0x6b, 0xe3, 0x6c, 0xc3, 0x4a, 0x76, 0xd6, 0xb5, 0x1d, 0x14, 0x6e, 0x7d, 0x80, 0x1e, 0xfa,
	0x18, 0x7d, 0x80, 0x9e, 0x7b, 0xe6, 0xc8, 0xb1, 0x27, 0xab, 0x0a, 0x37, 0x1f, 0x79, 0x82, 0xca,
	0x9b, 0x5f, 0x93, 0x70, 0xf1, 0xee, 0xcc, 0xf7, 0x33, 0xa3, 0xf5, 0xec, 0x02, 0xd9, 0xc2, 0x97,
	0xb8, 0x4b, 0x7c, 0x25, 0x5a, 0x15, 0xc7, 0xc2, 0x5d, 0x4f, 0x71, 0x98, 0x45, 0xcd, 0xab, 0x8a,
	0xe3, 0x32, 0x9f, 0x41, 0x38, 0x22, 0x54, 0xa2, 0xb5, 0xc2, 0x09, 0x3b, 0xeb, 0x1d, 0xd6, 0x61,
	0x1c, 0x56, 0xa2, 0xdd, 0x90, 0xb9, 0x23, 0x99, 0xcc, 0xb3, 0x99, 0xa7, 0xb4, 0xb0, 0x47, 0x94,
	0xcb, 0xfd, 0x16, 0xf1, 0xf1, 0xbe, 0x62, 0x32, 0xda, 0x1d, 0xe1, 0x2f, 0x62, 0xa5, 0x3c, 0x87,
	0x98, 0x0a, 0x76, 0xa8, 0x61, 0x32, 0xcb, 0x22, 0xa6, 0x4f, 0xd9, 0x88, 0x57, 0xfa, 0x93, 0x06,
	0xcb, 0x0d, 0xde, 0x02, 0x3c, 0x07, 0x6b, 0xe6, 0x05, 0xa6, 0x5d, 0x83, 0xb7, 0x44, 0x89, 0x27,
	0x0a, 0xc5, 0x54, 0x39, 0x77, 0x20, 0x57, 0xe6, 0xbb, 0xaa, 0x54, 0x23, 0xe6, 0x50, 0xa8, 0x6e,
	0x5e, 0x07, 0x72, 0x22, 0x0c, 0xe4, 0x7b, 0x72, 0xb4, 0x6a, 0x4e, 0x48, 0x94, 0x78, 0xf0, 0x3d,
	0x78, 0xdc, 0x21, 0xcc, 0x62, 0x26, 0x8e, 0xea, 0x1b, 0x8e, 0xcb, 0xbe, 0x52, 0x8b, 0x88, 0xc9,
	0xa2, 0x50, 0x4e, 0xab, 0x5b, 0x61, 0x20, 0x2f, 0x82, 0x11, 0x9c, 0x49, 0x36, 0x86, 0x39, 0xf8,
	0x16, 0xac, 0xf9, 0xcc, 0xc7, 0x96, 0x61, 0xf6, 0x0c, 0x8b, 0xda, 0xd4, 0x17, 0x53, 0xdc, 0x04,
	0x46, 0x4d, 0xc4, 0x11, 0x94, 0xe7, 0x71, 0xb5, 0x57, 0x8f, 0xa2, 0x48, 0x49, 0x1c, 0x66, 0x5e,
	0x4c, 0x95, 0xe9, 0xa9, 0x32, 0x8e, 0xa0, 0x3c, 0x8f, 0xc7, 0xca, 0x3a, 0xd8, 0xb0, 0x71, 0x3f,
	0x6a, 0xeb, 0x92, 0xb6, 0x89, 0xeb, 0x19, 0x3e, 0x33, 0x1c, 0x4c, 0x5d, 0x71, 0x89, 0x1b, 0x6c,
	0x87, 0x81, 0xbc, 0x98, 0x80, 0xa0, 0x8d, 0xfb, 0x8d, 0x71, 0xf6, 0x84, 0x35, 0x30, 0x75, 0xe1,
	0x77, 0x01, 0x6c, 0x79, 0x24, 0xfa, 0x15, 0xa4, 0x3d, 0x23, 0xb1, 0x59, 0x9b, 0x88, 0xcb, 0x45,
	0xa1, 0xbc, 0x76, 0xb0, 0xb7, 0xe8, 0xd4, 0x9b, 0x7a, 0x5d, 0xaf, 0x9e, 0xe8, 0x9a, 0xd1, 0x40,
	0xc7, 0xa7, 0x35, 0x4d, 0x47, 0x4d, 0xe3, 0xe8, 0x58, 0xd3, 0xd5, 0x27, 0x61, 0x20, 0x3f, 0xe4,
	0x87, 0x36, 0xc6, 0xc0, 0xa4, 0x89, 0x23, 0xd6, 0x26, 0x50, 0x07, 0x70, 0x5e, 0x21, 0xae, 0x14,
	0x53, 0xe5, 0xac, 0xba, 0x19, 0x06, 0xf2, 0x02, 0x14, 0x3d, 0x9a, 0xb3, 0x2a, 0xfd, 0x16, 0x40,
	0x6e, 0x66, 0x18, 0xe0, 0x4b, 0x90, 0x19, 0x8e, 0x01, 0x6d, 0x8b, 0x42, 0x51, 0x28, 0x67, 0xd5,
	0x7c, 0x18, 0xc8, 0x93, 0x1c, 0x5a, 0xe1, 0xbb, 0x5a, 0x1b, 0x3e, 0x05, 0x69, 0xec, 0x50, 0x4f,
	0x4c, 0xf2, 0x8a, 0x99, 0x30, 0x90, 0x79, 0x8c, 0xf8, 0x17, 0x9e, 0x83, 0xbc, 0x4b, 0xbe, 0xf5,
	0xa8, 0x4b, 0x6c, 0xd2, 0xf5, 0x3d, 0x31, 0xc5, 0x47, 0xf1, 0xd9, 0x83, 0xa3, 0x88, 0xa6, 0x64,
	0x75, 0x7d, 0x34, 0x8f, 0x31, 0x07, 0x14, 0x8b, 0x4a, 0x3f, 0x92, 0xa0, 0x70, 0x5f, 0x08, 0x5d,
	0x00, 0xa6, 0x17, 0x84, 0x77, 0x9f, 0x3b, 0xd8, 0x8d, 0x97, 0x8c, 0x6e, 0x52, 0xa5, 0x3a, 0x21,
	0x69, 0xd8, 0xc7, 0xaa, 0x32, 0xaa, 0x37, 0x23, 0xbe, 0x0b, 0xe4, 0x6d, 0x1b, 0x3b, 0x9e, 0xef,
	0xf6, 0x4c, 0xbf, 0xe7, 0x92, 0x77, 0xa5, 0x29, 0x56, 0x42, 0x33, 0x44, 0xf8, 0x01, 0x00, 0xd2,
	0xf7, 0x49, 0xd7, 0xa3, 0xac, 0x3b, 0x3e, 0x8c, 0xbd, 0xc8, 0x6c, 0x9a, 0x9d, 0x37, 0x9b, 0x62,
	0x25, 0x34, 0x43, 0x84, 0x6f, 0xc0, 0x92, 0x4d, 0xfb, 0xa4, 0xcd, 0xef, 0x43, 0x46, 0xdd, 0x0d,
	0x03, 0x79, 0x98, 0xb8, 0x0b, 0xe4, 0xf5, 0xb8, 0x05, 0x4f, 0x97, 0xd0, 0x10, 0x7e, 0xf5, 0x11,
	0x6c, 0x3d, 0x30, 0x5b, 0x30, 0x07, 0x56, 0x0e, 0xeb, 0xf5, 0xe3, 0xcf, 0xba, 0x56, 0x48, 0xc0,
	0x2c, 0x58, 0x3a, 0xaa, 0x9d, 0xe9, 0x5a, 0x41, 0x80, 0xab, 0x20, 0xab, 0x9f, 0x55, 0xeb, 0x9f,
	0x9a, 0xb5, 0x53, 0xbd, 0x90, 0x84, 0x79, 0x90, 0xd1, 0x6a, 0xcd, 0x43, 0xb5, 0xae, 0x6b, 0x85,
	0x94, 0x5a, 0xfd, 0x35, 0x90, 0x84, 0xeb, 0x81, 0x24, 0xdc, 0x0c, 0x24, 0xe1, 0xdf, 0x40, 0x12,
	0x7e, 0xde, 0x4a, 0x89, 0x9b, 0x5b, 0x29, 0xf1, 0xf7, 0x56, 0x4a, 0x7c, 0x79, 0xde, 0xa1, 0xfe,
	0x45, 0xaf, 0x55, 0x31, 0x99, 0xad, 0xc4, 0xde, 0xa9, 0xfe, 0xe8, 0x51, 0xf4, 0xaf, 0x1c, 0xe2,
	0xb5, 0x96, 0xf9, 0x13, 0xf5, 0xfa, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6e, 0xf4, 0xfa, 0x71,
	0x37, 0x05, 0x00, 0x00,
}

func (this *Policy) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Policy)
	if !ok {
		that2, ok := that.(Policy)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.ChainPolicies) != len(that1.ChainPolicies) {
		return false
	}
	for i := range this.ChainPolicies {
		if !this.ChainPolicies[i].Equal(&that1.ChainPolicies[i]) {
			return false
		}
	}
	if this.GeolocationProfile != that1.GeolocationProfile {
		return false
	}
	if this.TotalCuLimit != that1.TotalCuLimit {
		return false
	}
	if this.EpochCuLimit != that1.EpochCuLimit {
		return false
	}
	if this.MaxProvidersToPair != that1.MaxProvidersToPair {
		return false
	}
	if this.SelectedProvidersMode != that1.SelectedProvidersMode {
		return false
	}
	if len(this.SelectedProviders) != len(that1.SelectedProviders) {
		return false
	}
	for i := range this.SelectedProviders {
		if this.SelectedProviders[i] != that1.SelectedProviders[i] {
			return false
		}
	}
	return true
}
func (this *ChainPolicy) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ChainPolicy)
	if !ok {
		that2, ok := that.(ChainPolicy)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ChainId != that1.ChainId {
		return false
	}
	if len(this.Apis) != len(that1.Apis) {
		return false
	}
	for i := range this.Apis {
		if this.Apis[i] != that1.Apis[i] {
			return false
		}
	}
	if len(this.Requirements) != len(that1.Requirements) {
		return false
	}
	for i := range this.Requirements {
		if !this.Requirements[i].Equal(&that1.Requirements[i]) {
			return false
		}
	}
	return true
}
func (this *ChainRequirement) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ChainRequirement)
	if !ok {
		that2, ok := that.(ChainRequirement)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Collection.Equal(&that1.Collection) {
		return false
	}
	if len(this.Extensions) != len(that1.Extensions) {
		return false
	}
	for i := range this.Extensions {
		if this.Extensions[i] != that1.Extensions[i] {
			return false
		}
	}
	if this.Mixed != that1.Mixed {
		return false
	}
	return true
}
func (m *Policy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Policy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Policy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SelectedProviders) > 0 {
		for iNdEx := len(m.SelectedProviders) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SelectedProviders[iNdEx])
			copy(dAtA[i:], m.SelectedProviders[iNdEx])
			i = encodeVarintPolicy(dAtA, i, uint64(len(m.SelectedProviders[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.SelectedProvidersMode != 0 {
		i = encodeVarintPolicy(dAtA, i, uint64(m.SelectedProvidersMode))
		i--
		dAtA[i] = 0x30
	}
	if m.MaxProvidersToPair != 0 {
		i = encodeVarintPolicy(dAtA, i, uint64(m.MaxProvidersToPair))
		i--
		dAtA[i] = 0x28
	}
	if m.EpochCuLimit != 0 {
		i = encodeVarintPolicy(dAtA, i, uint64(m.EpochCuLimit))
		i--
		dAtA[i] = 0x20
	}
	if m.TotalCuLimit != 0 {
		i = encodeVarintPolicy(dAtA, i, uint64(m.TotalCuLimit))
		i--
		dAtA[i] = 0x18
	}
	if m.GeolocationProfile != 0 {
		i = encodeVarintPolicy(dAtA, i, uint64(m.GeolocationProfile))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ChainPolicies) > 0 {
		for iNdEx := len(m.ChainPolicies) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChainPolicies[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPolicy(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ChainPolicy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainPolicy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainPolicy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Requirements) > 0 {
		for iNdEx := len(m.Requirements) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Requirements[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPolicy(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Apis) > 0 {
		for iNdEx := len(m.Apis) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Apis[iNdEx])
			copy(dAtA[i:], m.Apis[iNdEx])
			i = encodeVarintPolicy(dAtA, i, uint64(len(m.Apis[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintPolicy(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ChainRequirement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainRequirement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainRequirement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Mixed {
		i--
		if m.Mixed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Extensions) > 0 {
		for iNdEx := len(m.Extensions) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Extensions[iNdEx])
			copy(dAtA[i:], m.Extensions[iNdEx])
			i = encodeVarintPolicy(dAtA, i, uint64(len(m.Extensions[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Collection.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPolicy(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintPolicy(dAtA []byte, offset int, v uint64) int {
	offset -= sovPolicy(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Policy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ChainPolicies) > 0 {
		for _, e := range m.ChainPolicies {
			l = e.Size()
			n += 1 + l + sovPolicy(uint64(l))
		}
	}
	if m.GeolocationProfile != 0 {
		n += 1 + sovPolicy(uint64(m.GeolocationProfile))
	}
	if m.TotalCuLimit != 0 {
		n += 1 + sovPolicy(uint64(m.TotalCuLimit))
	}
	if m.EpochCuLimit != 0 {
		n += 1 + sovPolicy(uint64(m.EpochCuLimit))
	}
	if m.MaxProvidersToPair != 0 {
		n += 1 + sovPolicy(uint64(m.MaxProvidersToPair))
	}
	if m.SelectedProvidersMode != 0 {
		n += 1 + sovPolicy(uint64(m.SelectedProvidersMode))
	}
	if len(m.SelectedProviders) > 0 {
		for _, s := range m.SelectedProviders {
			l = len(s)
			n += 1 + l + sovPolicy(uint64(l))
		}
	}
	return n
}

func (m *ChainPolicy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovPolicy(uint64(l))
	}
	if len(m.Apis) > 0 {
		for _, s := range m.Apis {
			l = len(s)
			n += 1 + l + sovPolicy(uint64(l))
		}
	}
	if len(m.Requirements) > 0 {
		for _, e := range m.Requirements {
			l = e.Size()
			n += 1 + l + sovPolicy(uint64(l))
		}
	}
	return n
}

func (m *ChainRequirement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Collection.Size()
	n += 1 + l + sovPolicy(uint64(l))
	if len(m.Extensions) > 0 {
		for _, s := range m.Extensions {
			l = len(s)
			n += 1 + l + sovPolicy(uint64(l))
		}
	}
	if m.Mixed {
		n += 2
	}
	return n
}

func sovPolicy(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPolicy(x uint64) (n int) {
	return sovPolicy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Policy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolicy
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Policy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Policy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainPolicies", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainPolicies = append(m.ChainPolicies, ChainPolicy{})
			if err := m.ChainPolicies[len(m.ChainPolicies)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GeolocationProfile", wireType)
			}
			m.GeolocationProfile = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GeolocationProfile |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalCuLimit", wireType)
			}
			m.TotalCuLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalCuLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochCuLimit", wireType)
			}
			m.EpochCuLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochCuLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxProvidersToPair", wireType)
			}
			m.MaxProvidersToPair = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxProvidersToPair |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SelectedProvidersMode", wireType)
			}
			m.SelectedProvidersMode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SelectedProvidersMode |= SELECTED_PROVIDERS_MODE(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SelectedProviders", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SelectedProviders = append(m.SelectedProviders, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPolicy
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ChainPolicy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolicy
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ChainPolicy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainPolicy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Apis", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Apis = append(m.Apis, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requirements", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Requirements = append(m.Requirements, ChainRequirement{})
			if err := m.Requirements[len(m.Requirements)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPolicy
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ChainRequirement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolicy
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ChainRequirement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainRequirement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collection", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Collection.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Extensions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Extensions = append(m.Extensions, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mixed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Mixed = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPolicy
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPolicy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPolicy
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPolicy
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPolicy
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPolicy
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPolicy        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPolicy          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPolicy = fmt.Errorf("proto: unexpected end of group")
)
