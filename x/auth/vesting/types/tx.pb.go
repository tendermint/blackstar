// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/vesting/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type VestingAccountType int32

const (
	VestingAccountType_DELAYED_VESTING_ACCOUNT_TYPE          VestingAccountType = 0
	VestingAccountType_CONTINUOUS_VESTING_ACCOUNT_TYPE       VestingAccountType = 1
	VestingAccountType_PERMANENT_LOCKED_VESTING_ACCOUNT_TYPE VestingAccountType = 2
)

var VestingAccountType_name = map[int32]string{
	0: "DELAYED_VESTING_ACCOUNT_TYPE",
	1: "CONTINUOUS_VESTING_ACCOUNT_TYPE",
	2: "PERMANENT_LOCKED_VESTING_ACCOUNT_TYPE",
}

var VestingAccountType_value = map[string]int32{
	"DELAYED_VESTING_ACCOUNT_TYPE":          0,
	"CONTINUOUS_VESTING_ACCOUNT_TYPE":       1,
	"PERMANENT_LOCKED_VESTING_ACCOUNT_TYPE": 2,
}

func (x VestingAccountType) String() string {
	return proto.EnumName(VestingAccountType_name, int32(x))
}

func (VestingAccountType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5338ca97811f9792, []int{0}
}

// MsgCreateVestingAccount defines a message that enables creating a vesting
// account.
type MsgCreateVestingAccount struct {
	FromAddress        string                                   `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	ToAddress          string                                   `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty" yaml:"to_address"`
	Amount             github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
	EndTime            int64                                    `protobuf:"varint,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty" yaml:"end_time"`
	VestingAccountType VestingAccountType                       `protobuf:"varint,5,opt,name=vesting_account_type,json=vestingAccountType,proto3,enum=cosmos.vesting.v1beta1.VestingAccountType" json:"vesting_account_type,omitempty" yaml:"vesting_account_type"`
}

func (m *MsgCreateVestingAccount) Reset()         { *m = MsgCreateVestingAccount{} }
func (m *MsgCreateVestingAccount) String() string { return proto.CompactTextString(m) }
func (*MsgCreateVestingAccount) ProtoMessage()    {}
func (*MsgCreateVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_5338ca97811f9792, []int{0}
}
func (m *MsgCreateVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateVestingAccount.Merge(m, src)
}
func (m *MsgCreateVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateVestingAccount proto.InternalMessageInfo

func (m *MsgCreateVestingAccount) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgCreateVestingAccount) GetToAddress() string {
	if m != nil {
		return m.ToAddress
	}
	return ""
}

func (m *MsgCreateVestingAccount) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *MsgCreateVestingAccount) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *MsgCreateVestingAccount) GetVestingAccountType() VestingAccountType {
	if m != nil {
		return m.VestingAccountType
	}
	return VestingAccountType_DELAYED_VESTING_ACCOUNT_TYPE
}

// MsgCreateVestingAccountResponse defines the Msg/CreateVestingAccount response type.
type MsgCreateVestingAccountResponse struct {
}

func (m *MsgCreateVestingAccountResponse) Reset()         { *m = MsgCreateVestingAccountResponse{} }
func (m *MsgCreateVestingAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateVestingAccountResponse) ProtoMessage()    {}
func (*MsgCreateVestingAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5338ca97811f9792, []int{1}
}
func (m *MsgCreateVestingAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateVestingAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateVestingAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateVestingAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateVestingAccountResponse.Merge(m, src)
}
func (m *MsgCreateVestingAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateVestingAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateVestingAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateVestingAccountResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("cosmos.vesting.v1beta1.VestingAccountType", VestingAccountType_name, VestingAccountType_value)
	proto.RegisterType((*MsgCreateVestingAccount)(nil), "cosmos.vesting.v1beta1.MsgCreateVestingAccount")
	proto.RegisterType((*MsgCreateVestingAccountResponse)(nil), "cosmos.vesting.v1beta1.MsgCreateVestingAccountResponse")
}

func init() { proto.RegisterFile("cosmos/vesting/v1beta1/tx.proto", fileDescriptor_5338ca97811f9792) }

var fileDescriptor_5338ca97811f9792 = []byte{
	// 521 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xbf, 0x6f, 0xd3, 0x40,
	0x18, 0xcd, 0x35, 0xa5, 0xd0, 0x2b, 0xe2, 0xc7, 0x35, 0xd0, 0x10, 0x90, 0x2f, 0x18, 0x21, 0x85,
	0x4a, 0xd8, 0x24, 0x20, 0x21, 0x65, 0x4b, 0x5c, 0x0b, 0x55, 0x6d, 0x9c, 0xca, 0x75, 0x2a, 0x95,
	0xc5, 0x72, 0xec, 0xc3, 0xb5, 0xc0, 0xbe, 0x28, 0x77, 0x89, 0x9a, 0x01, 0x89, 0x09, 0x31, 0xf2,
	0x27, 0xb0, 0xb0, 0xf0, 0x97, 0x74, 0xec, 0xc8, 0x64, 0x50, 0xb2, 0x30, 0xe7, 0x2f, 0x40, 0xf6,
	0x39, 0x01, 0xd4, 0x04, 0xa9, 0x93, 0x7d, 0xfa, 0xde, 0x7b, 0xdf, 0xbd, 0x77, 0xdf, 0x07, 0xb1,
	0x4b, 0x59, 0x48, 0x99, 0x3a, 0x24, 0x8c, 0x07, 0x91, 0xaf, 0x0e, 0xab, 0x5d, 0xc2, 0x9d, 0xaa,
	0xca, 0x4f, 0x95, 0x5e, 0x9f, 0x72, 0x8a, 0xee, 0x0a, 0x80, 0x92, 0x01, 0x94, 0x0c, 0x50, 0x2a,
	0xf8, 0xd4, 0xa7, 0x29, 0x44, 0x4d, 0xfe, 0x04, 0xba, 0x24, 0x65, 0x72, 0x5d, 0x87, 0x91, 0xb9,
	0x96, 0x4b, 0x83, 0x48, 0xd4, 0xe5, 0xaf, 0x79, 0xb8, 0xd5, 0x62, 0xbe, 0xd6, 0x27, 0x0e, 0x27,
	0x47, 0x42, 0xb2, 0xe1, 0xba, 0x74, 0x10, 0x71, 0x54, 0x87, 0xd7, 0xdf, 0xf4, 0x69, 0x68, 0x3b,
	0x9e, 0xd7, 0x27, 0x8c, 0x15, 0x41, 0x19, 0x54, 0xd6, 0x9b, 0x5b, 0xd3, 0x18, 0x6f, 0x8e, 0x9c,
	0xf0, 0x5d, 0x5d, 0xfe, 0xbb, 0x2a, 0x9b, 0x1b, 0xc9, 0xb1, 0x21, 0x4e, 0xe8, 0x05, 0x84, 0x9c,
	0xce, 0x99, 0x2b, 0x29, 0xf3, 0xce, 0x34, 0xc6, 0xb7, 0x05, 0xf3, 0x4f, 0x4d, 0x36, 0xd7, 0x39,
	0x9d, 0xb1, 0x5c, 0xb8, 0xe6, 0x84, 0x49, 0xef, 0x62, 0xbe, 0x9c, 0xaf, 0x6c, 0xd4, 0xee, 0x29,
	0x99, 0xd9, 0xe4, 0xfa, 0x33, 0xa7, 0x8a, 0x46, 0x83, 0xa8, 0xf9, 0xec, 0x2c, 0xc6, 0xb9, 0x6f,
	0x3f, 0x70, 0xc5, 0x0f, 0xf8, 0xc9, 0xa0, 0xab, 0xb8, 0x34, 0x54, 0x33, 0xaf, 0xe2, 0xf3, 0x94,
	0x79, 0x6f, 0x55, 0x3e, 0xea, 0x11, 0x96, 0x12, 0x98, 0x99, 0x49, 0x23, 0x05, 0x5e, 0x23, 0x91,
	0x67, 0xf3, 0x20, 0x24, 0xc5, 0xd5, 0x32, 0xa8, 0xe4, 0x9b, 0x9b, 0xd3, 0x18, 0xdf, 0x14, 0x17,
	0x9b, 0x55, 0x64, 0xf3, 0x2a, 0x89, 0x3c, 0x2b, 0x08, 0x09, 0x7a, 0x0f, 0x0b, 0x59, 0xd6, 0xb6,
	0x23, 0x92, 0xb1, 0x13, 0xd9, 0xe2, 0x95, 0x32, 0xa8, 0xdc, 0xa8, 0x6d, 0x2b, 0x8b, 0xdf, 0x43,
	0xf9, 0x37, 0x4c, 0x6b, 0xd4, 0x23, 0x4d, 0x3c, 0x8d, 0xf1, 0x7d, 0xd1, 0x67, 0x91, 0xa2, 0x6c,
	0xa2, 0xe1, 0x05, 0x52, 0x7d, 0xf5, 0xd7, 0x17, 0x0c, 0xe4, 0x87, 0x10, 0x2f, 0x79, 0x26, 0x93,
	0xb0, 0x1e, 0x8d, 0x18, 0xd9, 0xfe, 0x08, 0x20, 0xba, 0xd8, 0x14, 0x95, 0xe1, 0x83, 0x1d, 0x7d,
	0xbf, 0x71, 0xac, 0xef, 0xd8, 0x47, 0xfa, 0xa1, 0xb5, 0x6b, 0xbc, 0xb2, 0x1b, 0x9a, 0xd6, 0xee,
	0x18, 0x96, 0x6d, 0x1d, 0x1f, 0xe8, 0xb7, 0x72, 0xe8, 0x11, 0xc4, 0x5a, 0xdb, 0xb0, 0x76, 0x8d,
	0x4e, 0xbb, 0x73, 0xb8, 0x18, 0x04, 0xd0, 0x13, 0xf8, 0xf8, 0x40, 0x37, 0x5b, 0x0d, 0x43, 0x37,
	0x2c, 0x7b, 0xbf, 0xad, 0xed, 0x2d, 0xd3, 0x5b, 0xa9, 0x7d, 0x02, 0x30, 0xdf, 0x62, 0x3e, 0xfa,
	0x00, 0x60, 0x61, 0xe1, 0x60, 0xa9, 0xcb, 0x32, 0x5b, 0x62, 0xb1, 0xf4, 0xf2, 0x92, 0x84, 0x59,
	0x26, 0xcd, 0xbd, 0xb3, 0xb1, 0x04, 0xce, 0xc7, 0x12, 0xf8, 0x39, 0x96, 0xc0, 0xe7, 0x89, 0x94,
	0x3b, 0x9f, 0x48, 0xb9, 0xef, 0x13, 0x29, 0xf7, 0xba, 0xfa, 0xdf, 0xb9, 0x39, 0x55, 0x9d, 0x01,
	0x3f, 0x99, 0x2f, 0x61, 0x3a, 0x46, 0xdd, 0xb5, 0x74, 0x65, 0x9e, 0xff, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0x39, 0x69, 0x73, 0x02, 0xa3, 0x03, 0x00, 0x00,
}

func (this *MsgCreateVestingAccount) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgCreateVestingAccount)
	if !ok {
		that2, ok := that.(MsgCreateVestingAccount)
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
	if this.FromAddress != that1.FromAddress {
		return false
	}
	if this.ToAddress != that1.ToAddress {
		return false
	}
	if len(this.Amount) != len(that1.Amount) {
		return false
	}
	for i := range this.Amount {
		if !this.Amount[i].Equal(&that1.Amount[i]) {
			return false
		}
	}
	if this.EndTime != that1.EndTime {
		return false
	}
	if this.VestingAccountType != that1.VestingAccountType {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// CreateVestingAccount defines a method that enables creating a vesting
	// account.
	CreateVestingAccount(ctx context.Context, in *MsgCreateVestingAccount, opts ...grpc.CallOption) (*MsgCreateVestingAccountResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateVestingAccount(ctx context.Context, in *MsgCreateVestingAccount, opts ...grpc.CallOption) (*MsgCreateVestingAccountResponse, error) {
	out := new(MsgCreateVestingAccountResponse)
	err := c.cc.Invoke(ctx, "/cosmos.vesting.v1beta1.Msg/CreateVestingAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateVestingAccount defines a method that enables creating a vesting
	// account.
	CreateVestingAccount(context.Context, *MsgCreateVestingAccount) (*MsgCreateVestingAccountResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateVestingAccount(ctx context.Context, req *MsgCreateVestingAccount) (*MsgCreateVestingAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVestingAccount not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateVestingAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateVestingAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateVestingAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.vesting.v1beta1.Msg/CreateVestingAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateVestingAccount(ctx, req.(*MsgCreateVestingAccount))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.vesting.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVestingAccount",
			Handler:    _Msg_CreateVestingAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/vesting/v1beta1/tx.proto",
}

func (m *MsgCreateVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VestingAccountType != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.VestingAccountType))
		i--
		dAtA[i] = 0x28
	}
	if m.EndTime != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.EndTime))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ToAddress) > 0 {
		i -= len(m.ToAddress)
		copy(dAtA[i:], m.ToAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ToAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateVestingAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateVestingAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateVestingAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ToAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.EndTime != 0 {
		n += 1 + sovTx(uint64(m.EndTime))
	}
	if m.VestingAccountType != 0 {
		n += 1 + sovTx(uint64(m.VestingAccountType))
	}
	return n
}

func (m *MsgCreateVestingAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			m.EndTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingAccountType", wireType)
			}
			m.VestingAccountType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VestingAccountType |= VestingAccountType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateVestingAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateVestingAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateVestingAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
