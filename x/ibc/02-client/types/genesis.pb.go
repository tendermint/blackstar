// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/client/genesis.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
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

// GenesisClientState defines an identified ClientState as protobuf Any format.
type GenesisClientState struct {
	ClientID    string     `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty" yaml:"client_id"`
	ClientState *types.Any `protobuf:"bytes,2,opt,name=client_state,json=clientState,proto3" json:"client_state,omitempty" yaml:"client_state"`
}

func (m *GenesisClientState) Reset()         { *m = GenesisClientState{} }
func (m *GenesisClientState) String() string { return proto.CompactTextString(m) }
func (*GenesisClientState) ProtoMessage()    {}
func (*GenesisClientState) Descriptor() ([]byte, []int) {
	return fileDescriptor_2eb5d7ff040be5c2, []int{0}
}
func (m *GenesisClientState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisClientState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisClientState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisClientState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisClientState.Merge(m, src)
}
func (m *GenesisClientState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisClientState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisClientState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisClientState proto.InternalMessageInfo

func (m *GenesisClientState) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *GenesisClientState) GetClientState() *types.Any {
	if m != nil {
		return m.ClientState
	}
	return nil
}

// ClientConsensusStates defines all the stored consensus states for a given client.
type ClientConsensusStates struct {
	ClientID        string       `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ConsensusStates []*types.Any `protobuf:"bytes,2,rep,name=consensus_states,json=consensusStates,proto3" json:"consensus_states,omitempty" yaml:"consensus_states"`
}

func (m *ClientConsensusStates) Reset()         { *m = ClientConsensusStates{} }
func (m *ClientConsensusStates) String() string { return proto.CompactTextString(m) }
func (*ClientConsensusStates) ProtoMessage()    {}
func (*ClientConsensusStates) Descriptor() ([]byte, []int) {
	return fileDescriptor_2eb5d7ff040be5c2, []int{1}
}
func (m *ClientConsensusStates) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientConsensusStates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientConsensusStates.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientConsensusStates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientConsensusStates.Merge(m, src)
}
func (m *ClientConsensusStates) XXX_Size() int {
	return m.Size()
}
func (m *ClientConsensusStates) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientConsensusStates.DiscardUnknown(m)
}

var xxx_messageInfo_ClientConsensusStates proto.InternalMessageInfo

func (m *ClientConsensusStates) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *ClientConsensusStates) GetConsensusStates() []*types.Any {
	if m != nil {
		return m.ConsensusStates
	}
	return nil
}

// GenesisState defines the ibc client submodule's genesis state.
type GenesisState struct {
	Clients          []GenesisClientState    `protobuf:"bytes,1,rep,name=clients,proto3" json:"clients"`
	ClientsConsensus []ClientConsensusStates `protobuf:"bytes,2,rep,name=clients_consensus,json=clientsConsensus,proto3" json:"clients_consensus" yaml:"clients_consensus"`
	CreateLocalhost  bool                    `protobuf:"varint,3,opt,name=create_localhost,json=createLocalhost,proto3" json:"create_localhost,omitempty" yaml:"create_localhost"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_2eb5d7ff040be5c2, []int{2}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetClients() []GenesisClientState {
	if m != nil {
		return m.Clients
	}
	return nil
}

func (m *GenesisState) GetClientsConsensus() []ClientConsensusStates {
	if m != nil {
		return m.ClientsConsensus
	}
	return nil
}

func (m *GenesisState) GetCreateLocalhost() bool {
	if m != nil {
		return m.CreateLocalhost
	}
	return false
}

func init() {
	proto.RegisterType((*GenesisClientState)(nil), "ibc.client.GenesisClientState")
	proto.RegisterType((*ClientConsensusStates)(nil), "ibc.client.ClientConsensusStates")
	proto.RegisterType((*GenesisState)(nil), "ibc.client.GenesisState")
}

func init() { proto.RegisterFile("ibc/client/genesis.proto", fileDescriptor_2eb5d7ff040be5c2) }

var fileDescriptor_2eb5d7ff040be5c2 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x6e, 0xd3, 0x30,
	0x1c, 0xc7, 0xe3, 0x0e, 0x41, 0xe7, 0x55, 0x22, 0x84, 0xa1, 0x85, 0x21, 0x25, 0xc1, 0xa7, 0x72,
	0x98, 0x8d, 0xca, 0x0d, 0x09, 0x24, 0x32, 0x04, 0x9a, 0xb4, 0x03, 0x0a, 0x37, 0x84, 0x54, 0x25,
	0xae, 0xc9, 0x22, 0xd2, 0xb8, 0x9a, 0x5d, 0x89, 0xbc, 0x01, 0x47, 0x1e, 0x81, 0x0b, 0xef, 0xb2,
	0xe3, 0x8e, 0x9c, 0x22, 0x94, 0xbe, 0x41, 0x9e, 0x00, 0xc5, 0x76, 0xda, 0xb4, 0x85, 0x53, 0x9c,
	0xdf, 0x9f, 0xef, 0xf7, 0xa3, 0xaf, 0x0d, 0xdd, 0x2c, 0xa1, 0x84, 0xe6, 0x19, 0x2b, 0x24, 0x49,
	0x59, 0xc1, 0x44, 0x26, 0xf0, 0xe2, 0x9a, 0x4b, 0xee, 0xc0, 0x2c, 0xa1, 0x58, 0x77, 0x4e, 0x8f,
	0x53, 0x9e, 0x72, 0x55, 0x26, 0xed, 0x49, 0x4f, 0x9c, 0x3e, 0x4e, 0x39, 0x4f, 0x73, 0x46, 0xd4,
	0x5f, 0xb2, 0xfc, 0x42, 0xe2, 0xa2, 0xd4, 0x2d, 0xf4, 0x0b, 0x40, 0xe7, 0xbd, 0x96, 0x3b, 0x57,
	0x12, 0x1f, 0x65, 0x2c, 0x99, 0xf3, 0x0a, 0x1e, 0x6a, 0xc5, 0x69, 0x36, 0x73, 0x41, 0x00, 0xc6,
	0x87, 0x61, 0x50, 0x57, 0xfe, 0x50, 0xcf, 0x5c, 0xbc, 0x6d, 0x2a, 0xdf, 0x2e, 0xe3, 0x79, 0xfe,
	0x12, 0xad, 0xc7, 0x50, 0x34, 0xd4, 0xe7, 0x8b, 0x99, 0xf3, 0x01, 0x8e, 0x4c, 0x5d, 0xb4, 0x72,
	0xee, 0x20, 0x00, 0xe3, 0xa3, 0xc9, 0x31, 0xd6, 0x1c, 0xb8, 0xe3, 0xc0, 0x6f, 0x8a, 0x32, 0x3c,
	0x69, 0x2a, 0xff, 0xe1, 0x96, 0x96, 0xda, 0x41, 0xd1, 0x11, 0xdd, 0x00, 0xa1, 0x9f, 0x00, 0x3e,
	0xd2, 0xe6, 0xe7, 0xbc, 0x10, 0xac, 0x10, 0x4b, 0xa1, 0x1a, 0xc2, 0x79, 0xb6, 0x8f, 0x3a, 0xea,
	0xa3, 0xf6, 0xb0, 0x3e, 0x43, 0x9b, 0x76, 0xdb, 0xda, 0x45, 0xb8, 0x83, 0xe0, 0xe0, 0xbf, 0x68,
	0x4f, 0x9a, 0xca, 0x3f, 0x31, 0x68, 0x3b, 0x7b, 0x28, 0xba, 0x4f, 0xb7, 0x41, 0xd0, 0xf7, 0x01,
	0x1c, 0x99, 0x28, 0x75, 0x88, 0xaf, 0xe1, 0x3d, 0x6d, 0x2d, 0x5c, 0xa0, 0x5c, 0x3c, 0xbc, 0xb9,
	0x2a, 0xbc, 0x9f, 0x7a, 0x78, 0xe7, 0xa6, 0xf2, 0xad, 0xa8, 0x5b, 0x72, 0x16, 0xf0, 0x81, 0x39,
	0x4e, 0xd7, 0x5e, 0x86, 0xf7, 0x69, 0x5f, 0xe9, 0x9f, 0xb9, 0x84, 0x41, 0x2b, 0xd6, 0x54, 0xbe,
	0xdb, 0xcf, 0xb6, 0xa7, 0x84, 0x22, 0xdb, 0xd4, 0xd6, 0x9b, 0xce, 0x3b, 0x68, 0xd3, 0x6b, 0x16,
	0x4b, 0x36, 0xcd, 0x39, 0x8d, 0xf3, 0x2b, 0x2e, 0xa4, 0x7b, 0x10, 0x80, 0xf1, 0x70, 0x2b, 0x8a,
	0x9d, 0x89, 0x36, 0x0a, 0x55, 0xba, 0xec, 0x2a, 0xe1, 0xe5, 0x4d, 0xed, 0x81, 0xdb, 0xda, 0x03,
	0x7f, 0x6a, 0x0f, 0xfc, 0x58, 0x79, 0xd6, 0xed, 0xca, 0xb3, 0x7e, 0xaf, 0x3c, 0xeb, 0xd3, 0x24,
	0xcd, 0xe4, 0xd5, 0x32, 0xc1, 0x94, 0xcf, 0x09, 0xe5, 0x62, 0xce, 0x85, 0xf9, 0x9c, 0x89, 0xd9,
	0x57, 0xf2, 0x8d, 0xb4, 0xaf, 0xfc, 0xf9, 0xe4, 0xcc, 0x3c, 0x74, 0x59, 0x2e, 0x98, 0x48, 0xee,
	0xaa, 0x4b, 0x79, 0xf1, 0x37, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x16, 0xc1, 0x42, 0x03, 0x03, 0x00,
	0x00,
}

func (m *GenesisClientState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisClientState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisClientState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ClientState != nil {
		{
			size, err := m.ClientState.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClientID) > 0 {
		i -= len(m.ClientID)
		copy(dAtA[i:], m.ClientID)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ClientID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClientConsensusStates) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientConsensusStates) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientConsensusStates) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ConsensusStates) > 0 {
		for iNdEx := len(m.ConsensusStates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ConsensusStates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ClientID) > 0 {
		i -= len(m.ClientID)
		copy(dAtA[i:], m.ClientID)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ClientID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreateLocalhost {
		i--
		if m.CreateLocalhost {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.ClientsConsensus) > 0 {
		for iNdEx := len(m.ClientsConsensus) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClientsConsensus[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Clients) > 0 {
		for iNdEx := len(m.Clients) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Clients[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisClientState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientID)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.ClientState != nil {
		l = m.ClientState.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *ClientConsensusStates) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientID)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.ConsensusStates) > 0 {
		for _, e := range m.ConsensusStates {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Clients) > 0 {
		for _, e := range m.Clients {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ClientsConsensus) > 0 {
		for _, e := range m.ClientsConsensus {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.CreateLocalhost {
		n += 2
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisClientState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisClientState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisClientState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClientState == nil {
				m.ClientState = &types.Any{}
			}
			if err := m.ClientState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *ClientConsensusStates) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: ClientConsensusStates: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientConsensusStates: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusStates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsensusStates = append(m.ConsensusStates, &types.Any{})
			if err := m.ConsensusStates[len(m.ConsensusStates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Clients", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Clients = append(m.Clients, GenesisClientState{})
			if err := m.Clients[len(m.Clients)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientsConsensus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientsConsensus = append(m.ClientsConsensus, ClientConsensusStates{})
			if err := m.ClientsConsensus[len(m.ClientsConsensus)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateLocalhost", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
			m.CreateLocalhost = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
