// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibcdex/buyOrderBook.proto

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

type BuyOrderBook struct {
	Index        string   `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	OrderIDTrack int32    `protobuf:"varint,3,opt,name=orderIDTrack,proto3" json:"orderIDTrack,omitempty"`
	AmountDenom  string   `protobuf:"bytes,4,opt,name=amountDenom,proto3" json:"amountDenom,omitempty"`
	PriceDenom   string   `protobuf:"bytes,5,opt,name=priceDenom,proto3" json:"priceDenom,omitempty"`
	Orders       []*Order `protobuf:"bytes,6,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (m *BuyOrderBook) Reset()         { *m = BuyOrderBook{} }
func (m *BuyOrderBook) String() string { return proto.CompactTextString(m) }
func (*BuyOrderBook) ProtoMessage()    {}
func (*BuyOrderBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5716e4857f17ae7, []int{0}
}
func (m *BuyOrderBook) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BuyOrderBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BuyOrderBook.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BuyOrderBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuyOrderBook.Merge(m, src)
}
func (m *BuyOrderBook) XXX_Size() int {
	return m.Size()
}
func (m *BuyOrderBook) XXX_DiscardUnknown() {
	xxx_messageInfo_BuyOrderBook.DiscardUnknown(m)
}

var xxx_messageInfo_BuyOrderBook proto.InternalMessageInfo

func (m *BuyOrderBook) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *BuyOrderBook) GetOrderIDTrack() int32 {
	if m != nil {
		return m.OrderIDTrack
	}
	return 0
}

func (m *BuyOrderBook) GetAmountDenom() string {
	if m != nil {
		return m.AmountDenom
	}
	return ""
}

func (m *BuyOrderBook) GetPriceDenom() string {
	if m != nil {
		return m.PriceDenom
	}
	return ""
}

func (m *BuyOrderBook) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

func init() {
	proto.RegisterType((*BuyOrderBook)(nil), "tendermint.interchange.ibcdex.BuyOrderBook")
}

func init() { proto.RegisterFile("ibcdex/buyOrderBook.proto", fileDescriptor_c5716e4857f17ae7) }

var fileDescriptor_c5716e4857f17ae7 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4b, 0xc4, 0x30,
	0x18, 0x86, 0x1b, 0xcf, 0x16, 0xcc, 0xdd, 0x14, 0x1c, 0xaa, 0x60, 0x28, 0x87, 0x43, 0xa7, 0x44,
	0x74, 0x75, 0x2a, 0xb7, 0xe8, 0x22, 0x14, 0x27, 0xb7, 0x36, 0xfd, 0xb8, 0x0b, 0x47, 0x93, 0x92,
	0x4b, 0xa1, 0xf7, 0x2f, 0xfc, 0x59, 0x37, 0x76, 0x74, 0x94, 0xf6, 0x8f, 0x88, 0x69, 0xc1, 0xb8,
	0xdc, 0x98, 0x97, 0xf7, 0x79, 0xc8, 0xf7, 0xe2, 0x1b, 0x59, 0x8a, 0x0a, 0x3a, 0x5e, 0xb6, 0xc7,
	0x37, 0x53, 0x81, 0xc9, 0xb4, 0xde, 0xb3, 0xc6, 0x68, 0xab, 0xc9, 0x9d, 0x05, 0x55, 0x81, 0xa9,
	0xa5, 0xb2, 0x4c, 0x2a, 0x0b, 0x46, 0xec, 0x0a, 0xb5, 0x05, 0x36, 0x11, 0xb7, 0x64, 0x26, 0xf5,
	0x2f, 0x36, 0x21, 0xeb, 0x13, 0xc2, 0xab, 0xcc, 0x33, 0x91, 0x6b, 0x1c, 0x4a, 0x55, 0x41, 0x17,
	0x5f, 0x24, 0x28, 0xbd, 0xca, 0xa7, 0x07, 0x59, 0xe3, 0x95, 0xa3, 0x5e, 0x36, 0xef, 0xa6, 0x10,
	0xfb, 0x78, 0x91, 0xa0, 0x34, 0xcc, 0xff, 0x65, 0x24, 0xc1, 0xcb, 0xa2, 0xd6, 0xad, 0xb2, 0x1b,
	0x50, 0xba, 0x8e, 0x2f, 0x1d, 0xef, 0x47, 0x84, 0x62, 0xdc, 0x18, 0x29, 0x60, 0x2a, 0x84, 0xae,
	0xe0, 0x25, 0xe4, 0x19, 0x47, 0xce, 0x78, 0x88, 0xa3, 0x64, 0x91, 0x2e, 0x1f, 0xef, 0xd9, 0xd9,
	0x83, 0x98, 0xfb, 0x75, 0x3e, 0x33, 0xd9, 0xeb, 0x69, 0xa0, 0xa8, 0x1f, 0x28, 0xfa, 0x1e, 0x28,
	0xfa, 0x1c, 0x69, 0xd0, 0x8f, 0x34, 0xf8, 0x1a, 0x69, 0xf0, 0xf1, 0xb0, 0x95, 0x76, 0xd7, 0x96,
	0x4c, 0xe8, 0x9a, 0xff, 0x19, 0xb9, 0x67, 0xe4, 0x1d, 0x9f, 0xc7, 0xb1, 0xc7, 0x06, 0x0e, 0x65,
	0xe4, 0xd6, 0x79, 0xfa, 0x09, 0x00, 0x00, 0xff, 0xff, 0x58, 0x5f, 0xc0, 0xc7, 0x6d, 0x01, 0x00,
	0x00,
}

func (m *BuyOrderBook) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BuyOrderBook) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BuyOrderBook) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
				i = encodeVarintBuyOrderBook(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.PriceDenom) > 0 {
		i -= len(m.PriceDenom)
		copy(dAtA[i:], m.PriceDenom)
		i = encodeVarintBuyOrderBook(dAtA, i, uint64(len(m.PriceDenom)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.AmountDenom) > 0 {
		i -= len(m.AmountDenom)
		copy(dAtA[i:], m.AmountDenom)
		i = encodeVarintBuyOrderBook(dAtA, i, uint64(len(m.AmountDenom)))
		i--
		dAtA[i] = 0x22
	}
	if m.OrderIDTrack != 0 {
		i = encodeVarintBuyOrderBook(dAtA, i, uint64(m.OrderIDTrack))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintBuyOrderBook(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func encodeVarintBuyOrderBook(dAtA []byte, offset int, v uint64) int {
	offset -= sovBuyOrderBook(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BuyOrderBook) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovBuyOrderBook(uint64(l))
	}
	if m.OrderIDTrack != 0 {
		n += 1 + sovBuyOrderBook(uint64(m.OrderIDTrack))
	}
	l = len(m.AmountDenom)
	if l > 0 {
		n += 1 + l + sovBuyOrderBook(uint64(l))
	}
	l = len(m.PriceDenom)
	if l > 0 {
		n += 1 + l + sovBuyOrderBook(uint64(l))
	}
	if len(m.Orders) > 0 {
		for _, e := range m.Orders {
			l = e.Size()
			n += 1 + l + sovBuyOrderBook(uint64(l))
		}
	}
	return n
}

func sovBuyOrderBook(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBuyOrderBook(x uint64) (n int) {
	return sovBuyOrderBook(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BuyOrderBook) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBuyOrderBook
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
			return fmt.Errorf("proto: BuyOrderBook: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BuyOrderBook: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBuyOrderBook
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
				return ErrInvalidLengthBuyOrderBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBuyOrderBook
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
					return ErrIntOverflowBuyOrderBook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrderIDTrack |= int32(b&0x7F) << shift
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
					return ErrIntOverflowBuyOrderBook
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
				return ErrInvalidLengthBuyOrderBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBuyOrderBook
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
					return ErrIntOverflowBuyOrderBook
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
				return ErrInvalidLengthBuyOrderBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBuyOrderBook
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
					return ErrIntOverflowBuyOrderBook
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
				return ErrInvalidLengthBuyOrderBook
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBuyOrderBook
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
			skippy, err := skipBuyOrderBook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBuyOrderBook
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
func skipBuyOrderBook(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBuyOrderBook
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
					return 0, ErrIntOverflowBuyOrderBook
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
					return 0, ErrIntOverflowBuyOrderBook
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
				return 0, ErrInvalidLengthBuyOrderBook
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBuyOrderBook
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBuyOrderBook
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBuyOrderBook        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBuyOrderBook          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBuyOrderBook = fmt.Errorf("proto: unexpected end of group")
)
