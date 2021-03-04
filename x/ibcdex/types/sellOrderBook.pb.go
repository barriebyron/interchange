// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibcdex/sellOrderBook.proto

package types

import (
	fmt "fmt"
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

type SellOrderBook struct {
	Index        string   `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	OrderIDTrack uint32   `protobuf:"varint,3,opt,name=orderIDTrack,proto3" json:"orderIDTrack,omitempty"`
	AmountDenom  string   `protobuf:"bytes,4,opt,name=amountDenom,proto3" json:"amountDenom,omitempty"`
	PriceDenom   string   `protobuf:"bytes,5,opt,name=priceDenom,proto3" json:"priceDenom,omitempty"`
	Orders       []*Order `protobuf:"bytes,6,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (m *SellOrderBook) Reset()         { *m = SellOrderBook{} }
func (m *SellOrderBook) String() string { return proto.CompactTextString(m) }
func (*SellOrderBook) ProtoMessage()    {}
func (*SellOrderBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_fddb06350ff171e6, []int{0}
}
func (m *SellOrderBook) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SellOrderBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SellOrderBook.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SellOrderBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SellOrderBook.Merge(m, src)
}
func (m *SellOrderBook) XXX_Size() int {
	return m.Size()
}
func (m *SellOrderBook) XXX_DiscardUnknown() {
	xxx_messageInfo_SellOrderBook.DiscardUnknown(m)
}

var xxx_messageInfo_SellOrderBook proto.InternalMessageInfo

func (m *SellOrderBook) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *SellOrderBook) GetOrderIDTrack() uint32 {
	if m != nil {
		return m.OrderIDTrack
	}
	return 0
}

func (m *SellOrderBook) GetAmountDenom() string {
	if m != nil {
		return m.AmountDenom
	}
	return ""
}

func (m *SellOrderBook) GetPriceDenom() string {
	if m != nil {
		return m.PriceDenom
	}
	return ""
}

func (m *SellOrderBook) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

func init() {
	proto.RegisterType((*SellOrderBook)(nil), "tendermint.interchange.ibcdex.SellOrderBook")
}

func init() { proto.RegisterFile("ibcdex/sellOrderBook.proto", fileDescriptor_fddb06350ff171e6) }

var fileDescriptor_fddb06350ff171e6 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xca, 0x4c, 0x4a, 0x4e,
	0x49, 0xad, 0xd0, 0x2f, 0x4e, 0xcd, 0xc9, 0xf1, 0x2f, 0x4a, 0x49, 0x2d, 0x72, 0xca, 0xcf, 0xcf,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x2d, 0x49, 0xcd, 0x4b, 0x49, 0x2d, 0xca, 0xcd,
	0xcc, 0x2b, 0xd1, 0xcb, 0xcc, 0x2b, 0x49, 0x2d, 0x4a, 0xce, 0x48, 0xcc, 0x4b, 0x4f, 0xd5, 0x83,
	0x68, 0x91, 0x12, 0x82, 0x6a, 0xcd, 0x07, 0x69, 0x83, 0x68, 0x51, 0x3a, 0xc9, 0xc8, 0xc5, 0x1b,
	0x8c, 0x6c, 0x94, 0x90, 0x08, 0x17, 0x6b, 0x66, 0x5e, 0x4a, 0x6a, 0x85, 0x04, 0x93, 0x02, 0xa3,
	0x06, 0x67, 0x10, 0x84, 0x23, 0xa4, 0xc4, 0xc5, 0x03, 0xd6, 0xe6, 0xe9, 0x12, 0x52, 0x94, 0x98,
	0x9c, 0x2d, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x1b, 0x84, 0x22, 0x26, 0xa4, 0xc0, 0xc5, 0x9d, 0x98,
	0x9b, 0x5f, 0x9a, 0x57, 0xe2, 0x92, 0x9a, 0x97, 0x9f, 0x2b, 0xc1, 0x02, 0xd6, 0x8f, 0x2c, 0x24,
	0x24, 0xc7, 0xc5, 0x55, 0x50, 0x94, 0x99, 0x9c, 0x0a, 0x51, 0xc0, 0x0a, 0x56, 0x80, 0x24, 0x22,
	0x64, 0xc3, 0xc5, 0x06, 0x36, 0xb1, 0x58, 0x82, 0x4d, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x45, 0x0f,
	0xaf, 0x8f, 0xf4, 0xc0, 0xae, 0x0e, 0x82, 0xea, 0x71, 0xf2, 0x3a, 0xf1, 0x48, 0x8e, 0xf1, 0xc2,
	0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1,
	0xc6, 0x63, 0x39, 0x86, 0x28, 0x83, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c,
	0x7d, 0x84, 0x89, 0xfa, 0x48, 0x26, 0xea, 0x57, 0xe8, 0x43, 0x43, 0xa7, 0xa4, 0xb2, 0x20, 0xb5,
	0x38, 0x89, 0x0d, 0x1c, 0x3c, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0x36, 0x4c, 0x9c,
	0x6f, 0x01, 0x00, 0x00,
}

func (m *SellOrderBook) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SellOrderBook) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SellOrderBook) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Orders) > 0 {
		for iNdEx := len(m.Orders) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Orders[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSellOrderBook(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.PriceDenom) > 0 {
		i -= len(m.PriceDenom)
		copy(dAtA[i:], m.PriceDenom)
		i = encodeVarintSellOrderBook(dAtA, i, uint64(len(m.PriceDenom)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.AmountDenom) > 0 {
		i -= len(m.AmountDenom)
		copy(dAtA[i:], m.AmountDenom)
		i = encodeVarintSellOrderBook(dAtA, i, uint64(len(m.AmountDenom)))
		i--
		dAtA[i] = 0x22
	}
	if m.OrderIDTrack != 0 {
		i = encodeVarintSellOrderBook(dAtA, i, uint64(m.OrderIDTrack))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintSellOrderBook(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func encodeVarintSellOrderBook(dAtA []byte, offset int, v uint64) int {
	offset -= sovSellOrderBook(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SellOrderBook) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovSellOrderBook(uint64(l))
	}
	if m.OrderIDTrack != 0 {
		n += 1 + sovSellOrderBook(uint64(m.OrderIDTrack))
	}
	l = len(m.AmountDenom)
	if l > 0 {
		n += 1 + l + sovSellOrderBook(uint64(l))
	}
	l = len(m.PriceDenom)
	if l > 0 {
		n += 1 + l + sovSellOrderBook(uint64(l))
	}
	if len(m.Orders) > 0 {
		for _, e := range m.Orders {
			l = e.Size()
			n += 1 + l + sovSellOrderBook(uint64(l))
		}
	}
	return n
}

func sovSellOrderBook(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSellOrderBook(x uint64) (n int) {
	return sovSellOrderBook(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SellOrderBook) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSellOrderBook
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
			return fmt.Errorf("proto: SellOrderBook: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SellOrderBook: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSellOrderBook
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
				return ErrInvalidLengthSellOrderBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSellOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderIDTrack", wireType)
			}
			m.OrderIDTrack = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSellOrderBook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrderIDTrack |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSellOrderBook
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
				return ErrInvalidLengthSellOrderBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSellOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AmountDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSellOrderBook
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
				return ErrInvalidLengthSellOrderBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSellOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PriceDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Orders", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSellOrderBook
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
				return ErrInvalidLengthSellOrderBook
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSellOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Orders = append(m.Orders, &Order{})
			if err := m.Orders[len(m.Orders)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSellOrderBook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSellOrderBook
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
func skipSellOrderBook(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSellOrderBook
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
					return 0, ErrIntOverflowSellOrderBook
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
					return 0, ErrIntOverflowSellOrderBook
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
				return 0, ErrInvalidLengthSellOrderBook
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSellOrderBook
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSellOrderBook
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSellOrderBook        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSellOrderBook          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSellOrderBook = fmt.Errorf("proto: unexpected end of group")
)
