// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: simapp/types.proto

package simapp

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_x_auth_exported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	types "github.com/cosmos/cosmos-sdk/x/auth/types"
	types1 "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

// Account defines the application-level Account type.
type Account struct {
	// sum defines a list of all acceptable concrete Account implementations.
	//
	// Types that are valid to be assigned to Sum:
	//	*Account_BaseAccount
	//	*Account_ContinuousVestingAccount
	//	*Account_DelayedVestingAccount
	//	*Account_PeriodicVestingAccount
	Sum isAccount_Sum `protobuf_oneof:"sum"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcef5072c001b067, []int{0}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Account.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return m.Size()
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

type isAccount_Sum interface {
	isAccount_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Account_BaseAccount struct {
	BaseAccount *types.BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3,oneof" json:"base_account,omitempty"`
}
type Account_ContinuousVestingAccount struct {
	ContinuousVestingAccount *types1.ContinuousVestingAccount `protobuf:"bytes,2,opt,name=continuous_vesting_account,json=continuousVestingAccount,proto3,oneof" json:"continuous_vesting_account,omitempty"`
}
type Account_DelayedVestingAccount struct {
	DelayedVestingAccount *types1.DelayedVestingAccount `protobuf:"bytes,3,opt,name=delayed_vesting_account,json=delayedVestingAccount,proto3,oneof" json:"delayed_vesting_account,omitempty"`
}
type Account_PeriodicVestingAccount struct {
	PeriodicVestingAccount *types1.PeriodicVestingAccount `protobuf:"bytes,4,opt,name=periodic_vesting_account,json=periodicVestingAccount,proto3,oneof" json:"periodic_vesting_account,omitempty"`
}

func (*Account_BaseAccount) isAccount_Sum()              {}
func (*Account_ContinuousVestingAccount) isAccount_Sum() {}
func (*Account_DelayedVestingAccount) isAccount_Sum()    {}
func (*Account_PeriodicVestingAccount) isAccount_Sum()   {}

func (m *Account) GetSum() isAccount_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Account) GetBaseAccount() *types.BaseAccount {
	if x, ok := m.GetSum().(*Account_BaseAccount); ok {
		return x.BaseAccount
	}
	return nil
}

func (m *Account) GetContinuousVestingAccount() *types1.ContinuousVestingAccount {
	if x, ok := m.GetSum().(*Account_ContinuousVestingAccount); ok {
		return x.ContinuousVestingAccount
	}
	return nil
}

func (m *Account) GetDelayedVestingAccount() *types1.DelayedVestingAccount {
	if x, ok := m.GetSum().(*Account_DelayedVestingAccount); ok {
		return x.DelayedVestingAccount
	}
	return nil
}

func (m *Account) GetPeriodicVestingAccount() *types1.PeriodicVestingAccount {
	if x, ok := m.GetSum().(*Account_PeriodicVestingAccount); ok {
		return x.PeriodicVestingAccount
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Account) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Account_BaseAccount)(nil),
		(*Account_ContinuousVestingAccount)(nil),
		(*Account_DelayedVestingAccount)(nil),
		(*Account_PeriodicVestingAccount)(nil),
	}
}

func init() {
	proto.RegisterType((*Account)(nil), "cosmos_sdk.simapp.v1.Account")
}

func init() { proto.RegisterFile("simapp/types.proto", fileDescriptor_dcef5072c001b067) }

var fileDescriptor_dcef5072c001b067 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x4b, 0xfb, 0x30,
	0x18, 0xc7, 0xdb, 0xdf, 0x7e, 0x2a, 0x64, 0x9e, 0x8a, 0x7f, 0x4a, 0x0f, 0x65, 0xea, 0x45, 0x94,
	0xa5, 0xcc, 0xa9, 0xa0, 0xe0, 0xc1, 0x29, 0xe2, 0x51, 0x3c, 0x78, 0xf0, 0x52, 0xd2, 0x24, 0x6c,
	0x61, 0xb6, 0x09, 0x4d, 0x52, 0xb6, 0x77, 0xe1, 0x8b, 0xf1, 0x45, 0x88, 0xa7, 0x1d, 0x3d, 0xca,
	0x76, 0xf5, 0x45, 0xc8, 0x9a, 0xb8, 0x0d, 0x3a, 0x77, 0x7c, 0xfa, 0x7d, 0xbe, 0x9f, 0x4f, 0x21,
	0x0f, 0xf0, 0x24, 0x4b, 0x91, 0x10, 0x91, 0x1a, 0x0a, 0x2a, 0xa1, 0xc8, 0xb9, 0xe2, 0xde, 0x16,
	0xe6, 0x32, 0xe5, 0x32, 0x96, 0xa4, 0x0f, 0x4d, 0x0c, 0x8b, 0x56, 0x70, 0xac, 0x7a, 0x2c, 0x27,
	0xb1, 0x40, 0xb9, 0x1a, 0x46, 0xe5, 0x62, 0x64, 0xf6, 0x9a, 0x8b, 0x83, 0x41, 0x04, 0xfe, 0x20,
	0x42, 0x5a, 0xf5, 0x0c, 0x76, 0x11, 0x1e, 0x34, 0x6c, 0x52, 0x50, 0xa9, 0x58, 0xd6, 0xad, 0x6e,
	0xec, 0x7f, 0xd7, 0xc0, 0xc6, 0x35, 0xc6, 0x5c, 0x67, 0xca, 0xbb, 0x03, 0x9b, 0x09, 0x92, 0x34,
	0x46, 0x66, 0xf6, 0xdd, 0x86, 0x7b, 0x58, 0x3f, 0xd9, 0x83, 0x0b, 0x7f, 0x38, 0x80, 0x53, 0x1e,
	0x2c, 0x5a, 0xb0, 0x83, 0x24, 0xb5, 0xc5, 0x7b, 0xe7, 0xb1, 0x9e, 0xcc, 0x47, 0xaf, 0x00, 0x01,
	0xe6, 0x99, 0x62, 0x99, 0xe6, 0x5a, 0xc6, 0xd6, 0x3d, 0xa3, 0xfe, 0x2b, 0xa9, 0xe7, 0xcb, 0xa8,
	0x66, 0x73, 0x4a, 0xbf, 0x99, 0xf5, 0x9f, 0xcc, 0xc7, 0xb9, 0xca, 0xc7, 0x7f, 0x64, 0x5e, 0x0a,
	0x76, 0x09, 0x7d, 0x41, 0x43, 0x4a, 0x2a, 0xd2, 0x5a, 0x29, 0x6d, 0xaf, 0x96, 0xde, 0x9a, 0x72,
	0xc5, 0xb8, 0x4d, 0x96, 0x05, 0x9e, 0x00, 0xbe, 0xa0, 0x39, 0xe3, 0x84, 0xe1, 0x8a, 0xef, 0x7f,
	0xe9, 0x3b, 0x5d, 0xed, 0x7b, 0xb0, 0xed, 0x8a, 0x70, 0x47, 0x2c, 0x4d, 0x2e, 0x2f, 0x3e, 0xde,
	0x9a, 0x67, 0x47, 0x5d, 0xa6, 0x7a, 0x3a, 0x81, 0x98, 0xa7, 0xf6, 0x0c, 0x7e, 0x4f, 0x43, 0x92,
	0x7e, 0x64, 0x1f, 0x9c, 0x0e, 0x04, 0xcf, 0x15, 0x25, 0xd0, 0x56, 0x3b, 0x6b, 0xa0, 0x26, 0x75,
	0xda, 0xb9, 0x7a, 0x1f, 0x87, 0xee, 0x68, 0x1c, 0xba, 0x5f, 0xe3, 0xd0, 0x7d, 0x9d, 0x84, 0xce,
	0x68, 0x12, 0x3a, 0x9f, 0x93, 0xd0, 0x79, 0x3e, 0x58, 0x89, 0x35, 0x97, 0x99, 0xac, 0x97, 0x47,
	0xd3, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x1b, 0xc5, 0xeb, 0x96, 0xc9, 0x02, 0x00, 0x00,
}

func (this *Account) GetAccount() github_com_cosmos_cosmos_sdk_x_auth_exported.Account {
	if x := this.GetBaseAccount(); x != nil {
		return x
	}
	if x := this.GetContinuousVestingAccount(); x != nil {
		return x
	}
	if x := this.GetDelayedVestingAccount(); x != nil {
		return x
	}
	if x := this.GetPeriodicVestingAccount(); x != nil {
		return x
	}
	return nil
}

func (this *Account) SetAccount(value github_com_cosmos_cosmos_sdk_x_auth_exported.Account) error {
	if value == nil {
		this.Sum = nil
		return nil
	}
	switch vt := value.(type) {
	case *types.BaseAccount:
		this.Sum = &Account_BaseAccount{vt}
		return nil
	case *types1.ContinuousVestingAccount:
		this.Sum = &Account_ContinuousVestingAccount{vt}
		return nil
	case *types1.DelayedVestingAccount:
		this.Sum = &Account_DelayedVestingAccount{vt}
		return nil
	case *types1.PeriodicVestingAccount:
		this.Sum = &Account_PeriodicVestingAccount{vt}
		return nil
	}
	return fmt.Errorf("can't encode value of type %T as message Account", value)
}

func (m *Account) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Account) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Account_BaseAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account_BaseAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.BaseAccount != nil {
		{
			size, err := m.BaseAccount.MarshalToSizedBuffer(dAtA[:i])
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
func (m *Account_ContinuousVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account_ContinuousVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ContinuousVestingAccount != nil {
		{
			size, err := m.ContinuousVestingAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *Account_DelayedVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account_DelayedVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.DelayedVestingAccount != nil {
		{
			size, err := m.DelayedVestingAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *Account_PeriodicVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account_PeriodicVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PeriodicVestingAccount != nil {
		{
			size, err := m.PeriodicVestingAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
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
func (m *Account) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *Account_BaseAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseAccount != nil {
		l = m.BaseAccount.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Account_ContinuousVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ContinuousVestingAccount != nil {
		l = m.ContinuousVestingAccount.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Account_DelayedVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DelayedVestingAccount != nil {
		l = m.DelayedVestingAccount.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *Account_PeriodicVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PeriodicVestingAccount != nil {
		l = m.PeriodicVestingAccount.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Account) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Account: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Account: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAccount", wireType)
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
			v := &types.BaseAccount{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Account_BaseAccount{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContinuousVestingAccount", wireType)
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
			v := &types1.ContinuousVestingAccount{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Account_ContinuousVestingAccount{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelayedVestingAccount", wireType)
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
			v := &types1.DelayedVestingAccount{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Account_DelayedVestingAccount{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeriodicVestingAccount", wireType)
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
			v := &types1.PeriodicVestingAccount{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Account_PeriodicVestingAccount{v}
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
