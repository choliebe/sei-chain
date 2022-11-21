// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the dex module's genesis state.
type GenesisState struct {
	Params        Params          `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	ContractState []ContractState `protobuf:"bytes,2,rep,name=contractState,proto3" json:"contractState"`
	// this line is used by starport scaffolding # genesis/proto/state
	LastEpoch uint64 `protobuf:"varint,7,opt,name=lastEpoch,proto3" json:"lastEpoch,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a803aaabd08db59d, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetContractState() []ContractState {
	if m != nil {
		return m.ContractState
	}
	return nil
}

func (m *GenesisState) GetLastEpoch() uint64 {
	if m != nil {
		return m.LastEpoch
	}
	return 0
}

type ContractState struct {
	ContractInfo        ContractInfoV2 `protobuf:"bytes,1,opt,name=contractInfo,proto3" json:"contractInfo"`
	LongBookList        []LongBook     `protobuf:"bytes,2,rep,name=longBookList,proto3" json:"longBookList"`
	ShortBookList       []ShortBook    `protobuf:"bytes,3,rep,name=shortBookList,proto3" json:"shortBookList"`
	TriggeredOrdersList []Order        `protobuf:"bytes,4,rep,name=triggeredOrdersList,proto3" json:"triggeredOrdersList"`
	PairList            []Pair         `protobuf:"bytes,5,rep,name=pairList,proto3" json:"pairList"`
}

func (m *ContractState) Reset()         { *m = ContractState{} }
func (m *ContractState) String() string { return proto.CompactTextString(m) }
func (*ContractState) ProtoMessage()    {}
func (*ContractState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a803aaabd08db59d, []int{1}
}
func (m *ContractState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractState.Merge(m, src)
}
func (m *ContractState) XXX_Size() int {
	return m.Size()
}
func (m *ContractState) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractState.DiscardUnknown(m)
}

var xxx_messageInfo_ContractState proto.InternalMessageInfo

func (m *ContractState) GetContractInfo() ContractInfoV2 {
	if m != nil {
		return m.ContractInfo
	}
	return ContractInfoV2{}
}

func (m *ContractState) GetLongBookList() []LongBook {
	if m != nil {
		return m.LongBookList
	}
	return nil
}

func (m *ContractState) GetShortBookList() []ShortBook {
	if m != nil {
		return m.ShortBookList
	}
	return nil
}

func (m *ContractState) GetTriggeredOrdersList() []Order {
	if m != nil {
		return m.TriggeredOrdersList
	}
	return nil
}

func (m *ContractState) GetPairList() []Pair {
	if m != nil {
		return m.PairList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "seiprotocol.seichain.dex.GenesisState")
	proto.RegisterType((*ContractState)(nil), "seiprotocol.seichain.dex.ContractState")
}

func init() { proto.RegisterFile("dex/genesis.proto", fileDescriptor_a803aaabd08db59d) }

var fileDescriptor_a803aaabd08db59d = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0x87, 0x1b, 0x5b, 0x57, 0x9d, 0x6d, 0xfd, 0x33, 0xbb, 0x48, 0x29, 0x92, 0x2d, 0xf5, 0x60,
	0x2e, 0x9b, 0x40, 0xbd, 0x8b, 0x54, 0x64, 0x11, 0x0a, 0x2b, 0x2d, 0x28, 0x78, 0x29, 0xe9, 0x64,
	0x4c, 0x86, 0xa6, 0x79, 0xc3, 0xcc, 0x2c, 0x54, 0x3f, 0x85, 0x5f, 0xc0, 0x0f, 0xe3, 0x6d, 0x8f,
	0x7b, 0xf4, 0x24, 0xd2, 0x7e, 0x11, 0x99, 0x37, 0x33, 0xb5, 0x01, 0xe3, 0xde, 0x66, 0x7e, 0xf3,
	0xbc, 0x4f, 0xdf, 0xb7, 0x33, 0x21, 0x4f, 0x12, 0xbe, 0x89, 0x52, 0x5e, 0x70, 0x25, 0x54, 0x58,
	0x4a, 0xd0, 0x40, 0xfb, 0x8a, 0x0b, 0x5c, 0x31, 0xc8, 0x43, 0xc5, 0x05, 0xcb, 0x62, 0x51, 0x84,
	0x09, 0xdf, 0x0c, 0x4e, 0x53, 0x48, 0x01, 0x8f, 0x22, 0xb3, 0xaa, 0xf8, 0xc1, 0x63, 0xa3, 0x28,
	0x63, 0x19, 0xaf, 0xad, 0x61, 0x70, 0x62, 0x92, 0x1c, 0x8a, 0x74, 0xb1, 0x04, 0x58, 0xd9, 0xf0,
	0xd4, 0x84, 0x2a, 0x03, 0xa9, 0x0f, 0x53, 0x44, 0xb5, 0x60, 0xab, 0x85, 0x12, 0x5f, 0xb9, 0x0d,
	0x1f, 0x99, 0x10, 0x64, 0xc2, 0xa5, 0x0d, 0xa8, 0x09, 0x18, 0x14, 0x5a, 0xc6, 0x4c, 0xdb, 0xec,
	0x61, 0xf5, 0xb3, 0xc2, 0x31, 0x4f, 0xcd, 0x7e, 0x1d, 0x6b, 0x96, 0x2d, 0x24, 0x57, 0x57, 0xb9,
	0xe5, 0x46, 0x3f, 0x3c, 0xd2, 0xbd, 0xa8, 0x06, 0x9c, 0xeb, 0x58, 0x73, 0xfa, 0x8a, 0x1c, 0x55,
	0xdd, 0xf6, 0xbd, 0xa1, 0x17, 0x1c, 0x8f, 0x87, 0x61, 0xd3, 0xc0, 0xe1, 0x7b, 0xe4, 0x26, 0x9d,
	0xeb, 0x5f, 0x67, 0xad, 0x99, 0xad, 0xa2, 0x73, 0xd2, 0x73, 0xad, 0xa0, 0xb0, 0x7f, 0x67, 0xd8,
	0x0e, 0x8e, 0xc7, 0x2f, 0x9a, 0x35, 0x6f, 0x0e, 0x71, 0x6b, 0xab, 0x3b, 0xe8, 0x33, 0xf2, 0x20,
	0x8f, 0x95, 0x7e, 0x5b, 0x02, 0xcb, 0xfa, 0xf7, 0x86, 0x5e, 0xd0, 0x99, 0xfd, 0x0d, 0x46, 0xdf,
	0xdb, 0xa4, 0x57, 0x93, 0xd0, 0x19, 0xe9, 0x3a, 0xc1, 0xbb, 0xe2, 0x33, 0xd8, 0x51, 0x82, 0xdb,
	0x7b, 0x30, 0xf4, 0x87, 0xb1, 0x6d, 0xa2, 0xe6, 0xa0, 0x53, 0xd2, 0x35, 0x97, 0x36, 0x01, 0x58,
	0x4d, 0x85, 0xd2, 0x76, 0xae, 0x51, 0xb3, 0x73, 0x6a, 0x69, 0x67, 0x3b, 0xac, 0xa6, 0x97, 0xa4,
	0x87, 0xb7, 0xbd, 0xd7, 0xb5, 0x51, 0xf7, 0xbc, 0x59, 0x37, 0x77, 0xb8, 0xfb, 0x8b, 0x6a, 0xf5,
	0xf4, 0x23, 0x39, 0xd1, 0x52, 0xa4, 0x29, 0x97, 0x3c, 0xb9, 0x34, 0x8f, 0x43, 0xa1, 0xb6, 0x83,
	0xda, 0xb3, 0x66, 0x2d, 0xb2, 0x56, 0xf9, 0x2f, 0x03, 0x7d, 0x4d, 0xee, 0x9b, 0x77, 0x84, 0xb6,
	0xbb, 0x68, 0xf3, 0xff, 0xf7, 0x24, 0x84, 0x93, 0xed, 0xab, 0x26, 0x17, 0xd7, 0x5b, 0xdf, 0xbb,
	0xd9, 0xfa, 0xde, 0xef, 0xad, 0xef, 0x7d, 0xdb, 0xf9, 0xad, 0x9b, 0x9d, 0xdf, 0xfa, 0xb9, 0xf3,
	0x5b, 0x9f, 0xce, 0x53, 0xa1, 0xb3, 0xab, 0x65, 0xc8, 0x60, 0x1d, 0x29, 0x2e, 0xce, 0x9d, 0x14,
	0x37, 0x68, 0x8d, 0x36, 0x11, 0x7e, 0x03, 0x5f, 0x4a, 0xae, 0x96, 0x47, 0x78, 0xfe, 0xf2, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xd8, 0xe0, 0xa9, 0xf2, 0x97, 0x03, 0x00, 0x00,
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
	if m.LastEpoch != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LastEpoch))
		i--
		dAtA[i] = 0x38
	}
	if len(m.ContractState) > 0 {
		for iNdEx := len(m.ContractState) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ContractState[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ContractState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PairList) > 0 {
		for iNdEx := len(m.PairList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PairList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.TriggeredOrdersList) > 0 {
		for iNdEx := len(m.TriggeredOrdersList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TriggeredOrdersList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ShortBookList) > 0 {
		for iNdEx := len(m.ShortBookList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ShortBookList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.LongBookList) > 0 {
		for iNdEx := len(m.LongBookList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LongBookList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	{
		size, err := m.ContractInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.ContractState) > 0 {
		for _, e := range m.ContractState {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.LastEpoch != 0 {
		n += 1 + sovGenesis(uint64(m.LastEpoch))
	}
	return n
}

func (m *ContractState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ContractInfo.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.LongBookList) > 0 {
		for _, e := range m.LongBookList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ShortBookList) > 0 {
		for _, e := range m.ShortBookList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TriggeredOrdersList) > 0 {
		for _, e := range m.TriggeredOrdersList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PairList) > 0 {
		for _, e := range m.PairList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
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
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractState", wireType)
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
			m.ContractState = append(m.ContractState, ContractState{})
			if err := m.ContractState[len(m.ContractState)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastEpoch", wireType)
			}
			m.LastEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastEpoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *ContractState) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ContractState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractInfo", wireType)
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
			if err := m.ContractInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LongBookList", wireType)
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
			m.LongBookList = append(m.LongBookList, LongBook{})
			if err := m.LongBookList[len(m.LongBookList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShortBookList", wireType)
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
			m.ShortBookList = append(m.ShortBookList, ShortBook{})
			if err := m.ShortBookList[len(m.ShortBookList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TriggeredOrdersList", wireType)
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
			m.TriggeredOrdersList = append(m.TriggeredOrdersList, Order{})
			if err := m.TriggeredOrdersList[len(m.TriggeredOrdersList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairList", wireType)
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
			m.PairList = append(m.PairList, Pair{})
			if err := m.PairList[len(m.PairList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
