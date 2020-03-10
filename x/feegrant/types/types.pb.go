// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/feegrant/types/types.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgRevokeFeeAllowance struct {
	Granter github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=Granter,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"Granter,omitempty" yaml:"granter"`
	Grantee github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=Grantee,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"Grantee,omitempty" yaml:"grantee"`
}

func (m *MsgRevokeFeeAllowance) Reset()         { *m = MsgRevokeFeeAllowance{} }
func (m *MsgRevokeFeeAllowance) String() string { return proto.CompactTextString(m) }
func (*MsgRevokeFeeAllowance) ProtoMessage()    {}
func (*MsgRevokeFeeAllowance) Descriptor() ([]byte, []int) {
	return fileDescriptor_86c534389d2c5768, []int{0}
}
func (m *MsgRevokeFeeAllowance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRevokeFeeAllowance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRevokeFeeAllowance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRevokeFeeAllowance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRevokeFeeAllowance.Merge(m, src)
}
func (m *MsgRevokeFeeAllowance) XXX_Size() int {
	return m.Size()
}
func (m *MsgRevokeFeeAllowance) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRevokeFeeAllowance.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRevokeFeeAllowance proto.InternalMessageInfo

func (m *MsgRevokeFeeAllowance) GetGranter() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Granter
	}
	return nil
}

func (m *MsgRevokeFeeAllowance) GetGrantee() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Grantee
	}
	return nil
}

// BasicFeeAllowance implements FeeAllowance with a one-time grant of tokens
// that optionally expires. The delegatee can use up to SpendLimit to cover fees.
type BasicFeeAllowance struct {
	SpendLimit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=SpendLimit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"SpendLimit"`
	Expiration *ExpiresAt                               `protobuf:"bytes,2,opt,name=Expiration,proto3" json:"Expiration,omitempty"`
}

func (m *BasicFeeAllowance) Reset()         { *m = BasicFeeAllowance{} }
func (m *BasicFeeAllowance) String() string { return proto.CompactTextString(m) }
func (*BasicFeeAllowance) ProtoMessage()    {}
func (*BasicFeeAllowance) Descriptor() ([]byte, []int) {
	return fileDescriptor_86c534389d2c5768, []int{1}
}
func (m *BasicFeeAllowance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BasicFeeAllowance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BasicFeeAllowance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BasicFeeAllowance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasicFeeAllowance.Merge(m, src)
}
func (m *BasicFeeAllowance) XXX_Size() int {
	return m.Size()
}
func (m *BasicFeeAllowance) XXX_DiscardUnknown() {
	xxx_messageInfo_BasicFeeAllowance.DiscardUnknown(m)
}

var xxx_messageInfo_BasicFeeAllowance proto.InternalMessageInfo

func (m *BasicFeeAllowance) GetSpendLimit() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.SpendLimit
	}
	return nil
}

func (m *BasicFeeAllowance) GetExpiration() *ExpiresAt {
	if m != nil {
		return m.Expiration
	}
	return nil
}

type PeriodicFeeAllowance struct {
	Basic            *BasicFeeAllowance                       `protobuf:"bytes,1,opt,name=Basic,proto3" json:"Basic,omitempty"`
	Period           *Duration                                `protobuf:"bytes,2,opt,name=Period,proto3" json:"Period,omitempty"`
	PeriodSpendLimit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=PeriodSpendLimit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"PeriodSpendLimit"`
	PeriodCanSpend   github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=PeriodCanSpend,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"PeriodCanSpend"`
	PeriodReset      *ExpiresAt                               `protobuf:"bytes,5,opt,name=PeriodReset,proto3" json:"PeriodReset,omitempty"`
}

func (m *PeriodicFeeAllowance) Reset()         { *m = PeriodicFeeAllowance{} }
func (m *PeriodicFeeAllowance) String() string { return proto.CompactTextString(m) }
func (*PeriodicFeeAllowance) ProtoMessage()    {}
func (*PeriodicFeeAllowance) Descriptor() ([]byte, []int) {
	return fileDescriptor_86c534389d2c5768, []int{2}
}
func (m *PeriodicFeeAllowance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PeriodicFeeAllowance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PeriodicFeeAllowance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PeriodicFeeAllowance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeriodicFeeAllowance.Merge(m, src)
}
func (m *PeriodicFeeAllowance) XXX_Size() int {
	return m.Size()
}
func (m *PeriodicFeeAllowance) XXX_DiscardUnknown() {
	xxx_messageInfo_PeriodicFeeAllowance.DiscardUnknown(m)
}

var xxx_messageInfo_PeriodicFeeAllowance proto.InternalMessageInfo

func (m *PeriodicFeeAllowance) GetBasic() *BasicFeeAllowance {
	if m != nil {
		return m.Basic
	}
	return nil
}

func (m *PeriodicFeeAllowance) GetPeriod() *Duration {
	if m != nil {
		return m.Period
	}
	return nil
}

func (m *PeriodicFeeAllowance) GetPeriodSpendLimit() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.PeriodSpendLimit
	}
	return nil
}

func (m *PeriodicFeeAllowance) GetPeriodCanSpend() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.PeriodCanSpend
	}
	return nil
}

func (m *PeriodicFeeAllowance) GetPeriodReset() *ExpiresAt {
	if m != nil {
		return m.PeriodReset
	}
	return nil
}

// Duration is a repeating unit of either clock time or number of blocks.
// This is designed to be added to an ExpiresAt struct.
type Duration struct {
	Clock time.Duration `protobuf:"bytes,1,opt,name=Clock,proto3,stdduration" json:"Clock" yaml:"time"`
	Block int64         `protobuf:"varint,2,opt,name=Block,proto3" json:"Block,omitempty"`
}

func (m *Duration) Reset()         { *m = Duration{} }
func (m *Duration) String() string { return proto.CompactTextString(m) }
func (*Duration) ProtoMessage()    {}
func (*Duration) Descriptor() ([]byte, []int) {
	return fileDescriptor_86c534389d2c5768, []int{3}
}
func (m *Duration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Duration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Duration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Duration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Duration.Merge(m, src)
}
func (m *Duration) XXX_Size() int {
	return m.Size()
}
func (m *Duration) XXX_DiscardUnknown() {
	xxx_messageInfo_Duration.DiscardUnknown(m)
}

var xxx_messageInfo_Duration proto.InternalMessageInfo

func (m *Duration) GetClock() time.Duration {
	if m != nil {
		return m.Clock
	}
	return 0
}

func (m *Duration) GetBlock() int64 {
	if m != nil {
		return m.Block
	}
	return 0
}

// ExpiresAt is a point in time where something expires.
// It may be *either* block time or block height
type ExpiresAt struct {
	Time   time.Time `protobuf:"bytes,1,opt,name=Time,proto3,stdtime" json:"Time" yaml:"time"`
	Height int64     `protobuf:"varint,2,opt,name=Height,proto3" json:"Height,omitempty"`
}

func (m *ExpiresAt) Reset()         { *m = ExpiresAt{} }
func (m *ExpiresAt) String() string { return proto.CompactTextString(m) }
func (*ExpiresAt) ProtoMessage()    {}
func (*ExpiresAt) Descriptor() ([]byte, []int) {
	return fileDescriptor_86c534389d2c5768, []int{4}
}
func (m *ExpiresAt) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExpiresAt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExpiresAt.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExpiresAt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpiresAt.Merge(m, src)
}
func (m *ExpiresAt) XXX_Size() int {
	return m.Size()
}
func (m *ExpiresAt) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpiresAt.DiscardUnknown(m)
}

var xxx_messageInfo_ExpiresAt proto.InternalMessageInfo

func (m *ExpiresAt) GetTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

func (m *ExpiresAt) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgRevokeFeeAllowance)(nil), "cosmos_sdk.x.feegrant.v1.MsgRevokeFeeAllowance")
	proto.RegisterType((*BasicFeeAllowance)(nil), "cosmos_sdk.x.feegrant.v1.BasicFeeAllowance")
	proto.RegisterType((*PeriodicFeeAllowance)(nil), "cosmos_sdk.x.feegrant.v1.PeriodicFeeAllowance")
	proto.RegisterType((*Duration)(nil), "cosmos_sdk.x.feegrant.v1.Duration")
	proto.RegisterType((*ExpiresAt)(nil), "cosmos_sdk.x.feegrant.v1.ExpiresAt")
}

func init() { proto.RegisterFile("x/feegrant/types/types.proto", fileDescriptor_86c534389d2c5768) }

var fileDescriptor_86c534389d2c5768 = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x63, 0xd2, 0x14, 0xd8, 0xa0, 0x8a, 0x2e, 0x05, 0xac, 0x08, 0xd9, 0x95, 0x91, 0x50,
	0x24, 0xd4, 0x35, 0x0d, 0xb7, 0xde, 0xe2, 0x50, 0x02, 0x12, 0x48, 0xc8, 0x70, 0x42, 0x42, 0x95,
	0x63, 0x4f, 0x9d, 0x55, 0x6c, 0xaf, 0xf1, 0x6e, 0x42, 0xf2, 0x16, 0x3d, 0xf2, 0x0c, 0x3c, 0x04,
	0x5c, 0x7b, 0xec, 0x05, 0x89, 0x53, 0x8a, 0x92, 0x37, 0xe8, 0x91, 0x13, 0xb2, 0xd7, 0xa6, 0x26,
	0x21, 0xa8, 0xaa, 0x7a, 0xb1, 0x77, 0xc7, 0x33, 0xdf, 0xfc, 0x33, 0xa3, 0x31, 0x7a, 0x30, 0x36,
	0x0f, 0x01, 0xfc, 0xc4, 0x89, 0x84, 0x29, 0x26, 0x31, 0x70, 0xf9, 0x24, 0x71, 0xc2, 0x04, 0xc3,
	0xaa, 0xcb, 0x78, 0xc8, 0xf8, 0x01, 0xf7, 0x06, 0x64, 0x4c, 0x0a, 0x47, 0x32, 0xda, 0x6d, 0x3c,
	0x12, 0x7d, 0x9a, 0x78, 0x07, 0xb1, 0x93, 0x88, 0x89, 0x99, 0x39, 0x9b, 0x3e, 0xf3, 0xd9, 0xf9,
	0x49, 0x12, 0x1a, 0x9b, 0x4b, 0xd0, 0x86, 0xee, 0x33, 0xe6, 0x07, 0x20, 0xa3, 0x7a, 0xc3, 0x43,
	0x53, 0xd0, 0x10, 0xb8, 0x70, 0xc2, 0x58, 0x3a, 0x18, 0xdf, 0x15, 0x74, 0xf7, 0x35, 0xf7, 0x6d,
	0x18, 0xb1, 0x01, 0x3c, 0x07, 0x68, 0x07, 0x01, 0xfb, 0xe4, 0x44, 0x2e, 0xe0, 0x0f, 0xe8, 0x7a,
	0x37, 0x55, 0x00, 0x89, 0xaa, 0x6c, 0x2b, 0xcd, 0x5b, 0x56, 0xe7, 0x6c, 0xaa, 0x6f, 0x4c, 0x9c,
	0x30, 0xd8, 0x33, 0x7c, 0xf9, 0xc1, 0xf8, 0x35, 0xd5, 0x77, 0x7c, 0x2a, 0xfa, 0xc3, 0x1e, 0x71,
	0x59, 0x68, 0xca, 0x0a, 0xf2, 0xd7, 0x0e, 0xf7, 0x06, 0xb9, 0x96, 0xb6, 0xeb, 0xb6, 0x3d, 0x2f,
	0x01, 0xce, 0xed, 0x82, 0x79, 0x8e, 0x07, 0xf5, 0xda, 0xbf, 0xf1, 0x70, 0x79, 0x3c, 0x18, 0xdf,
	0x14, 0xb4, 0x69, 0x39, 0x9c, 0xba, 0x7f, 0xd5, 0x04, 0x08, 0xbd, 0x8d, 0x21, 0xf2, 0x5e, 0xd1,
	0x90, 0x0a, 0x55, 0xd9, 0xae, 0x36, 0xeb, 0xad, 0x3b, 0xa4, 0xd4, 0xf8, 0xd1, 0x2e, 0xe9, 0x30,
	0x1a, 0x59, 0x4f, 0x8e, 0xa7, 0x7a, 0xe5, 0xcb, 0xa9, 0xde, 0xbc, 0x40, 0xfa, 0x34, 0x80, 0xdb,
	0x25, 0x30, 0xee, 0x20, 0xb4, 0x3f, 0x8e, 0x69, 0xe2, 0x08, 0xca, 0xa2, 0xac, 0xbc, 0x7a, 0xeb,
	0x21, 0x59, 0x35, 0x5f, 0x92, 0xf9, 0x02, 0x6f, 0x0b, 0xbb, 0x14, 0x66, 0x7c, 0xad, 0xa2, 0xad,
	0x37, 0x90, 0x50, 0xe6, 0x2d, 0x14, 0xd1, 0x46, 0xb5, 0xac, 0xb2, 0x6c, 0x2c, 0xf5, 0xd6, 0xe3,
	0xd5, 0xe0, 0xa5, 0x06, 0xd8, 0x32, 0x12, 0xef, 0xa1, 0x75, 0x89, 0xce, 0xc5, 0x19, 0xab, 0x19,
	0xcf, 0x86, 0x52, 0x8f, 0x9d, 0x47, 0xe0, 0x8f, 0xe8, 0xb6, 0x3c, 0x95, 0x3a, 0x59, 0xbd, 0xca,
	0x4e, 0x2e, 0xe1, 0x71, 0x88, 0x36, 0xa4, 0xad, 0xe3, 0x44, 0x99, 0x59, 0x5d, 0xbb, 0xca, 0x84,
	0x0b, 0x70, 0xbc, 0x8f, 0xea, 0xd2, 0x62, 0x03, 0x07, 0xa1, 0xd6, 0x2e, 0x3e, 0xbf, 0x72, 0x9c,
	0x31, 0x40, 0x37, 0x8a, 0xe6, 0xe1, 0x97, 0xa8, 0xd6, 0x09, 0x98, 0x3b, 0xc8, 0x67, 0xd6, 0x20,
	0x72, 0x2f, 0x49, 0xb1, 0x97, 0xe4, 0x5d, 0xb1, 0x97, 0xd6, 0xfd, 0x54, 0xff, 0xd9, 0x54, 0xaf,
	0xcb, 0x5d, 0x48, 0x17, 0xd6, 0xf8, 0x7c, 0xaa, 0x2b, 0xb6, 0x24, 0xe0, 0x2d, 0x54, 0xb3, 0x32,
	0x54, 0x3a, 0xba, 0xaa, 0x2d, 0x2f, 0x46, 0x80, 0x6e, 0xfe, 0x91, 0x81, 0xbb, 0x68, 0x2d, 0xe5,
	0x5d, 0x32, 0xd9, 0x51, 0x9a, 0x2c, 0x03, 0xe0, 0x7b, 0x68, 0xfd, 0x05, 0x50, 0xbf, 0x2f, 0xf2,
	0x64, 0xf9, 0xcd, 0xea, 0x1e, 0xcf, 0x34, 0xe5, 0x64, 0xa6, 0x29, 0x3f, 0x67, 0x9a, 0x72, 0x34,
	0xd7, 0x2a, 0x27, 0x73, 0xad, 0xf2, 0x63, 0xae, 0x55, 0xde, 0xff, 0x7f, 0x5f, 0x17, 0x7f, 0x80,
	0xbd, 0xf5, 0x4c, 0xd3, 0xd3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x40, 0xee, 0xe5, 0x1d, 0x1b,
	0x05, 0x00, 0x00,
}

func (m *MsgRevokeFeeAllowance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRevokeFeeAllowance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRevokeFeeAllowance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BasicFeeAllowance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BasicFeeAllowance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BasicFeeAllowance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Expiration != nil {
		{
			size, err := m.Expiration.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.SpendLimit) > 0 {
		for iNdEx := len(m.SpendLimit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SpendLimit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *PeriodicFeeAllowance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PeriodicFeeAllowance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PeriodicFeeAllowance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PeriodReset != nil {
		{
			size, err := m.PeriodReset.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PeriodCanSpend) > 0 {
		for iNdEx := len(m.PeriodCanSpend) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PeriodCanSpend[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.PeriodSpendLimit) > 0 {
		for iNdEx := len(m.PeriodSpendLimit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PeriodSpendLimit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Period != nil {
		{
			size, err := m.Period.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Basic != nil {
		{
			size, err := m.Basic.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Duration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Duration) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Duration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Block != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x10
	}
	n5, err5 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Clock, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Clock):])
	if err5 != nil {
		return 0, err5
	}
	i -= n5
	i = encodeVarintTypes(dAtA, i, uint64(n5))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ExpiresAt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExpiresAt) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExpiresAt) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	n6, err6 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Time):])
	if err6 != nil {
		return 0, err6
	}
	i -= n6
	i = encodeVarintTypes(dAtA, i, uint64(n6))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRevokeFeeAllowance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *BasicFeeAllowance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SpendLimit) > 0 {
		for _, e := range m.SpendLimit {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.Expiration != nil {
		l = m.Expiration.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *PeriodicFeeAllowance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Basic != nil {
		l = m.Basic.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Period != nil {
		l = m.Period.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.PeriodSpendLimit) > 0 {
		for _, e := range m.PeriodSpendLimit {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.PeriodCanSpend) > 0 {
		for _, e := range m.PeriodCanSpend {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.PeriodReset != nil {
		l = m.PeriodReset.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Duration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Clock)
	n += 1 + l + sovTypes(uint64(l))
	if m.Block != 0 {
		n += 1 + sovTypes(uint64(m.Block))
	}
	return n
}

func (m *ExpiresAt) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovTypes(uint64(l))
	if m.Height != 0 {
		n += 1 + sovTypes(uint64(m.Height))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRevokeFeeAllowance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: MsgRevokeFeeAllowance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRevokeFeeAllowance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Granter = append(m.Granter[:0], dAtA[iNdEx:postIndex]...)
			if m.Granter == nil {
				m.Granter = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Grantee = append(m.Grantee[:0], dAtA[iNdEx:postIndex]...)
			if m.Grantee == nil {
				m.Grantee = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *BasicFeeAllowance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: BasicFeeAllowance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BasicFeeAllowance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpendLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpendLimit = append(m.SpendLimit, types.Coin{})
			if err := m.SpendLimit[len(m.SpendLimit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Expiration == nil {
				m.Expiration = &ExpiresAt{}
			}
			if err := m.Expiration.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *PeriodicFeeAllowance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: PeriodicFeeAllowance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PeriodicFeeAllowance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Basic", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Basic == nil {
				m.Basic = &BasicFeeAllowance{}
			}
			if err := m.Basic.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Period == nil {
				m.Period = &Duration{}
			}
			if err := m.Period.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeriodSpendLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeriodSpendLimit = append(m.PeriodSpendLimit, types.Coin{})
			if err := m.PeriodSpendLimit[len(m.PeriodSpendLimit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeriodCanSpend", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeriodCanSpend = append(m.PeriodCanSpend, types.Coin{})
			if err := m.PeriodCanSpend[len(m.PeriodCanSpend)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeriodReset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PeriodReset == nil {
				m.PeriodReset = &ExpiresAt{}
			}
			if err := m.PeriodReset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Duration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Duration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Duration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Clock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Clock, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ExpiresAt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ExpiresAt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExpiresAt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
