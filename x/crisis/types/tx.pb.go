// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/crisis/v1beta1/tx.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// MsgVerifyInvariant represents a message to verify a particular invariance.
type MsgVerifyInvariant struct {
	Sender              github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=sender,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"sender,omitempty"`
	InvariantModuleName string                                        `protobuf:"bytes,2,opt,name=invariant_module_name,json=invariantModuleName,proto3" json:"invariant_module_name,omitempty" yaml:"invariant_module_name"`
	InvariantRoute      string                                        `protobuf:"bytes,3,opt,name=invariant_route,json=invariantRoute,proto3" json:"invariant_route,omitempty" yaml:"invariant_route"`
}

func (m *MsgVerifyInvariant) Reset()         { *m = MsgVerifyInvariant{} }
func (m *MsgVerifyInvariant) String() string { return proto.CompactTextString(m) }
func (*MsgVerifyInvariant) ProtoMessage()    {}
func (*MsgVerifyInvariant) Descriptor() ([]byte, []int) {
	return fileDescriptor_61276163172fe867, []int{0}
}
func (m *MsgVerifyInvariant) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgVerifyInvariant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgVerifyInvariant.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgVerifyInvariant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgVerifyInvariant.Merge(m, src)
}
func (m *MsgVerifyInvariant) XXX_Size() int {
	return m.Size()
}
func (m *MsgVerifyInvariant) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgVerifyInvariant.DiscardUnknown(m)
}

var xxx_messageInfo_MsgVerifyInvariant proto.InternalMessageInfo

func (m *MsgVerifyInvariant) GetSender() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *MsgVerifyInvariant) GetInvariantModuleName() string {
	if m != nil {
		return m.InvariantModuleName
	}
	return ""
}

func (m *MsgVerifyInvariant) GetInvariantRoute() string {
	if m != nil {
		return m.InvariantRoute
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgVerifyInvariant)(nil), "cosmos.crisis.v1beta1.MsgVerifyInvariant")
}

func init() { proto.RegisterFile("cosmos/crisis/v1beta1/tx.proto", fileDescriptor_61276163172fe867) }

var fileDescriptor_61276163172fe867 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x4f, 0x2e, 0xca, 0x2c, 0xce, 0x2c, 0xd6, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49,
	0x34, 0xd4, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x85, 0xc8, 0xeb, 0x41,
	0xe4, 0xf5, 0xa0, 0xf2, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x15, 0xfa, 0x20, 0x16, 0x44,
	0xb1, 0x52, 0x03, 0x13, 0x97, 0x90, 0x6f, 0x71, 0x7a, 0x58, 0x6a, 0x51, 0x66, 0x5a, 0xa5, 0x67,
	0x5e, 0x59, 0x62, 0x51, 0x66, 0x62, 0x5e, 0x89, 0x90, 0x27, 0x17, 0x5b, 0x71, 0x6a, 0x5e, 0x4a,
	0x6a, 0x91, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x8f, 0x93, 0xe1, 0xaf, 0x7b, 0xf2, 0xba, 0xe9, 0x99,
	0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x30, 0x27, 0x80, 0x29, 0xdd, 0xe2, 0x94,
	0x6c, 0xfd, 0x92, 0xca, 0x82, 0xd4, 0x62, 0x3d, 0xc7, 0xe4, 0x64, 0xc7, 0x94, 0x94, 0xa2, 0xd4,
	0xe2, 0xe2, 0x20, 0xa8, 0x01, 0x42, 0x21, 0x5c, 0xa2, 0x99, 0x30, 0x73, 0xe3, 0x73, 0xf3, 0x53,
	0x4a, 0x73, 0x52, 0xe3, 0xf3, 0x12, 0x73, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x9d, 0x14,
	0x3e, 0xdd, 0x93, 0x97, 0xa9, 0x4c, 0xcc, 0xcd, 0xb1, 0x52, 0xc2, 0xaa, 0x4c, 0x29, 0x48, 0x18,
	0x2e, 0xee, 0x0b, 0x16, 0xf6, 0x4b, 0xcc, 0x4d, 0x15, 0x72, 0xe6, 0xe2, 0x47, 0x28, 0x2f, 0xca,
	0x2f, 0x2d, 0x49, 0x95, 0x60, 0x06, 0x9b, 0x27, 0xf5, 0xe9, 0x9e, 0xbc, 0x18, 0xba, 0x79, 0x60,
	0x05, 0x4a, 0x41, 0x7c, 0x70, 0x91, 0x20, 0x90, 0x80, 0x15, 0xcb, 0x8b, 0x05, 0xf2, 0x8c, 0x4e,
	0xae, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7,
	0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x8d, 0xd7, 0xc7, 0x15,
	0xb0, 0x18, 0x00, 0x7b, 0x3d, 0x89, 0x0d, 0x1c, 0xa0, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xb0, 0xb6, 0x10, 0xbd, 0x9f, 0x01, 0x00, 0x00,
}

func (this *MsgVerifyInvariant) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgVerifyInvariant)
	if !ok {
		that2, ok := that.(MsgVerifyInvariant)
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
	if !bytes.Equal(this.Sender, that1.Sender) {
		return false
	}
	if this.InvariantModuleName != that1.InvariantModuleName {
		return false
	}
	if this.InvariantRoute != that1.InvariantRoute {
		return false
	}
	return true
}
func (m *MsgVerifyInvariant) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgVerifyInvariant) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgVerifyInvariant) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.InvariantRoute) > 0 {
		i -= len(m.InvariantRoute)
		copy(dAtA[i:], m.InvariantRoute)
		i = encodeVarintTx(dAtA, i, uint64(len(m.InvariantRoute)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.InvariantModuleName) > 0 {
		i -= len(m.InvariantModuleName)
		copy(dAtA[i:], m.InvariantModuleName)
		i = encodeVarintTx(dAtA, i, uint64(len(m.InvariantModuleName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
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
func (m *MsgVerifyInvariant) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.InvariantModuleName)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.InvariantRoute)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgVerifyInvariant) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgVerifyInvariant: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgVerifyInvariant: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InvariantModuleName", wireType)
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
			m.InvariantModuleName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InvariantRoute", wireType)
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
			m.InvariantRoute = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
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
