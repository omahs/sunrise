// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sunrise/liquiditypool/pool.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	cosmossdk_io_math "cosmossdk.io/math"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type Pool struct {
	Id                   uint64                      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DenomBase            string                      `protobuf:"bytes,2,opt,name=denom_base,json=denomBase,proto3" json:"denom_base,omitempty"`
	DenomQuote           string                      `protobuf:"bytes,3,opt,name=denom_quote,json=denomQuote,proto3" json:"denom_quote,omitempty"`
	FeeRate              cosmossdk_io_math.LegacyDec `protobuf:"bytes,4,opt,name=fee_rate,json=feeRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"fee_rate"`
	TickParams           TickParams                  `protobuf:"bytes,5,opt,name=tick_params,json=tickParams,proto3" json:"tick_params"`
	CurrentTick          int64                       `protobuf:"varint,6,opt,name=current_tick,json=currentTick,proto3" json:"current_tick,omitempty"`
	CurrentTickLiquidity cosmossdk_io_math.LegacyDec `protobuf:"bytes,7,opt,name=current_tick_liquidity,json=currentTickLiquidity,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"current_tick_liquidity"`
	CurrentSqrtPrice     cosmossdk_io_math.LegacyDec `protobuf:"bytes,8,opt,name=current_sqrt_price,json=currentSqrtPrice,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"current_sqrt_price"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5e900f89b7804df, []int{0}
}
func (m *Pool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return m.Size()
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

func (m *Pool) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Pool) GetDenomBase() string {
	if m != nil {
		return m.DenomBase
	}
	return ""
}

func (m *Pool) GetDenomQuote() string {
	if m != nil {
		return m.DenomQuote
	}
	return ""
}

func (m *Pool) GetTickParams() TickParams {
	if m != nil {
		return m.TickParams
	}
	return TickParams{}
}

func (m *Pool) GetCurrentTick() int64 {
	if m != nil {
		return m.CurrentTick
	}
	return 0
}

// PriceRatio^(Tick - BaseOffSet)
type TickParams struct {
	// Basically 1.0001
	PriceRatio cosmossdk_io_math.LegacyDec `protobuf:"bytes,1,opt,name=price_ratio,json=priceRatio,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"price_ratio"`
	// basically 0 and (-1,0]. In the 1:1 stable pair, -0.5 would work
	BaseOffset cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=base_offset,json=baseOffset,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"base_offset"`
}

func (m *TickParams) Reset()         { *m = TickParams{} }
func (m *TickParams) String() string { return proto.CompactTextString(m) }
func (*TickParams) ProtoMessage()    {}
func (*TickParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_b5e900f89b7804df, []int{1}
}
func (m *TickParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TickParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TickParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TickParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TickParams.Merge(m, src)
}
func (m *TickParams) XXX_Size() int {
	return m.Size()
}
func (m *TickParams) XXX_DiscardUnknown() {
	xxx_messageInfo_TickParams.DiscardUnknown(m)
}

var xxx_messageInfo_TickParams proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Pool)(nil), "sunrise.liquiditypool.Pool")
	proto.RegisterType((*TickParams)(nil), "sunrise.liquiditypool.TickParams")
}

func init() { proto.RegisterFile("sunrise/liquiditypool/pool.proto", fileDescriptor_b5e900f89b7804df) }

var fileDescriptor_b5e900f89b7804df = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xeb, 0xb6, 0x6c, 0xab, 0x83, 0x10, 0x58, 0x03, 0x85, 0x21, 0xd2, 0x6c, 0xa7, 0x0a,
	0x89, 0x44, 0x02, 0xc1, 0x07, 0xa8, 0x76, 0x1c, 0x5a, 0x17, 0x90, 0x90, 0xb8, 0x44, 0x6e, 0xf2,
	0xda, 0x59, 0x4d, 0xfa, 0x52, 0xdb, 0x91, 0xe8, 0xb7, 0xe0, 0x63, 0x70, 0xe4, 0xc0, 0x99, 0xf3,
	0x8e, 0x13, 0x27, 0xc4, 0x61, 0x42, 0xed, 0x81, 0x23, 0x5f, 0x01, 0xd9, 0x4e, 0xc7, 0x40, 0x9c,
	0xb6, 0x8b, 0xd5, 0xf7, 0x7f, 0x7f, 0xff, 0xfe, 0xae, 0xf3, 0x4c, 0x43, 0x55, 0xcf, 0xa5, 0x50,
	0x10, 0x17, 0x62, 0x51, 0x8b, 0x5c, 0xe8, 0x65, 0x85, 0x58, 0xc4, 0x66, 0x89, 0x2a, 0x89, 0x1a,
	0xd9, 0xfd, 0xc6, 0x11, 0xfd, 0xe5, 0xd8, 0xbb, 0xc7, 0x4b, 0x31, 0xc7, 0xd8, 0xae, 0xce, 0xb9,
	0xf7, 0x30, 0x43, 0x55, 0xa2, 0x4a, 0x6d, 0x15, 0xbb, 0xa2, 0x69, 0xed, 0x4e, 0x71, 0x8a, 0x4e,
	0x37, 0xbf, 0x9c, 0x7a, 0xf0, 0xab, 0x43, 0xbb, 0x23, 0xc4, 0x82, 0xdd, 0xa1, 0x6d, 0x91, 0xfb,
	0x24, 0x24, 0x83, 0x6e, 0xd2, 0x16, 0x39, 0x7b, 0x4c, 0x69, 0x0e, 0x73, 0x2c, 0xd3, 0x31, 0x57,
	0xe0, 0xb7, 0x43, 0x32, 0xe8, 0x25, 0x3d, 0xab, 0x0c, 0xb9, 0x02, 0xd6, 0xa7, 0x9e, 0x6b, 0x2f,
	0x6a, 0xd4, 0xe0, 0x77, 0x6c, 0xdf, 0xed, 0x38, 0x31, 0x0a, 0x3b, 0xa1, 0x3b, 0x13, 0x80, 0x54,
	0x72, 0x0d, 0x7e, 0xd7, 0x74, 0x87, 0x2f, 0xcf, 0x2e, 0xfa, 0xad, 0xef, 0x17, 0xfd, 0x47, 0xee,
	0x58, 0x2a, 0x9f, 0x45, 0x02, 0xe3, 0x92, 0xeb, 0xd3, 0xe8, 0x08, 0xa6, 0x3c, 0x5b, 0x1e, 0x42,
	0xf6, 0xf5, 0xf3, 0x53, 0xda, 0x9c, 0xfa, 0x10, 0xb2, 0x8f, 0x3f, 0x3f, 0x3d, 0x21, 0xc9, 0xf6,
	0x04, 0x20, 0xe1, 0x1a, 0xd8, 0x2b, 0xea, 0x69, 0x91, 0xcd, 0xd2, 0x8a, 0x4b, 0x5e, 0x2a, 0xff,
	0x56, 0x48, 0x06, 0xde, 0xb3, 0xfd, 0xe8, 0xbf, 0x97, 0x13, 0xbd, 0x11, 0xd9, 0x6c, 0x64, 0x8d,
	0xc3, 0x9e, 0x09, 0x76, 0x2c, 0xaa, 0x2f, 0x65, 0xb6, 0x4f, 0x6f, 0x67, 0xb5, 0x94, 0x30, 0xd7,
	0xa9, 0x51, 0xfd, 0xad, 0x90, 0x0c, 0x3a, 0x89, 0xd7, 0x68, 0x66, 0x3f, 0x2b, 0xe8, 0x83, 0xab,
	0x96, 0xf4, 0x32, 0xc2, 0xdf, 0xbe, 0xd1, 0x5f, 0xda, 0xbd, 0x12, 0x72, 0xb4, 0x61, 0xb2, 0x9c,
	0xb2, 0x4d, 0x9a, 0x5a, 0x48, 0x9d, 0x56, 0x52, 0x64, 0xe0, 0xef, 0xdc, 0x28, 0xe9, 0x6e, 0x43,
	0x7c, 0xbd, 0x90, 0x7a, 0x64, 0x78, 0x07, 0x5f, 0x08, 0xa5, 0x7f, 0x2e, 0x87, 0xbd, 0xa5, 0x9e,
	0xcd, 0x31, 0x5f, 0x4a, 0xa0, 0x1d, 0x80, 0xeb, 0xa7, 0x51, 0x8b, 0x4a, 0x0c, 0xc9, 0x80, 0xcd,
	0xe8, 0xa4, 0x38, 0x99, 0x28, 0xd0, 0x6e, 0x82, 0xae, 0x0f, 0x36, 0xa8, 0x63, 0x4b, 0x1a, 0x1e,
	0x9f, 0xad, 0x02, 0x72, 0xbe, 0x0a, 0xc8, 0x8f, 0x55, 0x40, 0x3e, 0xac, 0x83, 0xd6, 0xf9, 0x3a,
	0x68, 0x7d, 0x5b, 0x07, 0xad, 0x77, 0x2f, 0xa6, 0x42, 0x9f, 0xd6, 0xe3, 0x28, 0xc3, 0x32, 0x6e,
	0xa6, 0xa2, 0xe0, 0x4b, 0x90, 0x9b, 0x22, 0x7e, 0xff, 0xcf, 0x1b, 0xd3, 0xcb, 0x0a, 0xd4, 0x78,
	0xcb, 0x3e, 0x85, 0xe7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x38, 0x40, 0x39, 0x9a, 0x89, 0x03,
	0x00, 0x00,
}

func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.CurrentSqrtPrice.Size()
		i -= size
		if _, err := m.CurrentSqrtPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.CurrentTickLiquidity.Size()
		i -= size
		if _, err := m.CurrentTickLiquidity.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.CurrentTick != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.CurrentTick))
		i--
		dAtA[i] = 0x30
	}
	{
		size, err := m.TickParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.FeeRate.Size()
		i -= size
		if _, err := m.FeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.DenomQuote) > 0 {
		i -= len(m.DenomQuote)
		copy(dAtA[i:], m.DenomQuote)
		i = encodeVarintPool(dAtA, i, uint64(len(m.DenomQuote)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DenomBase) > 0 {
		i -= len(m.DenomBase)
		copy(dAtA[i:], m.DenomBase)
		i = encodeVarintPool(dAtA, i, uint64(len(m.DenomBase)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TickParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TickParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TickParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.BaseOffset.Size()
		i -= size
		if _, err := m.BaseOffset.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.PriceRatio.Size()
		i -= size
		if _, err := m.PriceRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPool(uint64(m.Id))
	}
	l = len(m.DenomBase)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = len(m.DenomQuote)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	l = m.FeeRate.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.TickParams.Size()
	n += 1 + l + sovPool(uint64(l))
	if m.CurrentTick != 0 {
		n += 1 + sovPool(uint64(m.CurrentTick))
	}
	l = m.CurrentTickLiquidity.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.CurrentSqrtPrice.Size()
	n += 1 + l + sovPool(uint64(l))
	return n
}

func (m *TickParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PriceRatio.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.BaseOffset.Size()
	n += 1 + l + sovPool(uint64(l))
	return n
}

func sovPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPool(x uint64) (n int) {
	return sovPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomBase", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomBase = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomQuote", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomQuote = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TickParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentTick", wireType)
			}
			m.CurrentTick = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentTick |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentTickLiquidity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentTickLiquidity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentSqrtPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentSqrtPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
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
func (m *TickParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: TickParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TickParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PriceRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseOffset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BaseOffset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
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
func skipPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
				return 0, ErrInvalidLengthPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPool = fmt.Errorf("proto: unexpected end of group")
)
