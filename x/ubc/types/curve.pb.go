// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: escbackbone/ubc/curve.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type Curve struct {
	FS0             *Flatsegment                           `protobuf:"bytes,1,opt,name=fS0,proto3" json:"fS0,omitempty"`
	S0              *BezierSegment                         `protobuf:"bytes,2,opt,name=s0,proto3" json:"s0,omitempty"`
	S1              *BezierSegment                         `protobuf:"bytes,3,opt,name=s1,proto3" json:"s1,omitempty"`
	S2              *FixedBezierSegment                    `protobuf:"bytes,4,opt,name=s2,proto3" json:"s2,omitempty"`
	QS3             *Quadraticsegment                      `protobuf:"bytes,5,opt,name=qS3,proto3" json:"qS3,omitempty"`
	RefProfitFactor github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=refProfitFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"refProfitFactor"`
	RefTokenSupply  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=refTokenSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"refTokenSupply"`
	RefTokenPrice   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=refTokenPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"refTokenPrice"`
	BPool           github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,9,opt,name=bPool,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"bPool"`
	BPoolUnder      github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,10,opt,name=bPoolUnder,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"bPoolUnder"`
	FactorFy        github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,11,opt,name=factorFy,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"factorFy"`
	FactorFxy       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,12,opt,name=factorFxy,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"factorFxy"`
	TradingPoint    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,13,opt,name=tradingPoint,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"tradingPoint"`
	CurrentSupply   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,14,opt,name=currentSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"currentSupply"`
	SlopeP2         github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,15,opt,name=slopeP2,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slopeP2"`
	SlopeP3         github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,16,opt,name=slopeP3,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slopeP3"`
}

func (m *Curve) Reset()         { *m = Curve{} }
func (m *Curve) String() string { return proto.CompactTextString(m) }
func (*Curve) ProtoMessage()    {}
func (*Curve) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c5281ef0eb0b3ff, []int{0}
}
func (m *Curve) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Curve) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Curve.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Curve) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Curve.Merge(m, src)
}
func (m *Curve) XXX_Size() int {
	return m.Size()
}
func (m *Curve) XXX_DiscardUnknown() {
	xxx_messageInfo_Curve.DiscardUnknown(m)
}

var xxx_messageInfo_Curve proto.InternalMessageInfo

func (m *Curve) GetFS0() *Flatsegment {
	if m != nil {
		return m.FS0
	}
	return nil
}

func (m *Curve) GetS0() *BezierSegment {
	if m != nil {
		return m.S0
	}
	return nil
}

func (m *Curve) GetS1() *BezierSegment {
	if m != nil {
		return m.S1
	}
	return nil
}

func (m *Curve) GetS2() *FixedBezierSegment {
	if m != nil {
		return m.S2
	}
	return nil
}

func (m *Curve) GetQS3() *Quadraticsegment {
	if m != nil {
		return m.QS3
	}
	return nil
}

func init() {
	proto.RegisterType((*Curve)(nil), "escbackbone.ubc.Curve")
}

func init() { proto.RegisterFile("escbackbone/ubc/curve.proto", fileDescriptor_7c5281ef0eb0b3ff) }

var fileDescriptor_7c5281ef0eb0b3ff = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc7, 0x9b, 0x94, 0x6e, 0xad, 0xf7, 0x51, 0x64, 0x71, 0x61, 0x0d, 0xc8, 0x3e, 0x90, 0xa6,
	0x49, 0xa8, 0xc9, 0xd6, 0xbc, 0x41, 0x19, 0x15, 0x42, 0x08, 0x85, 0x74, 0x20, 0xc4, 0x5d, 0xe2,
	0x9c, 0x84, 0xa8, 0x6d, 0x9c, 0xd9, 0x0e, 0x6a, 0x79, 0x0a, 0x1e, 0x6b, 0x97, 0xbb, 0x44, 0x5c,
	0x4c, 0xa8, 0x7d, 0x02, 0xde, 0x00, 0xc5, 0xed, 0xd6, 0x26, 0xbd, 0x41, 0xe1, 0x2a, 0x47, 0xce,
	0xef, 0xff, 0x4b, 0x8e, 0x6d, 0x1d, 0xf4, 0x14, 0x04, 0xf5, 0x3d, 0x3a, 0xf4, 0x59, 0x02, 0x56,
	0xe6, 0x53, 0x8b, 0x66, 0xfc, 0x1b, 0x98, 0x29, 0x67, 0x92, 0xe1, 0xf6, 0xda, 0x4b, 0x33, 0xf3,
	0xe9, 0xc1, 0x8b, 0x32, 0xed, 0xc3, 0xf7, 0x18, 0xb8, 0x80, 0x68, 0x0c, 0x89, 0x5c, 0xa4, 0x0e,
	0x4e, 0xcb, 0xd0, 0x75, 0xe6, 0x05, 0xdc, 0x93, 0x31, 0x2d, 0x72, 0xc7, 0x65, 0x2e, 0x1c, 0x79,
	0xb2, 0x88, 0x9c, 0x6c, 0x20, 0xf1, 0x04, 0x82, 0x22, 0xf3, 0x24, 0x62, 0x11, 0x53, 0xa5, 0x95,
	0x57, 0x8b, 0xd5, 0x93, 0x3f, 0x4d, 0xd4, 0x78, 0x95, 0xb7, 0x82, 0x4d, 0x54, 0x0f, 0x07, 0xe7,
	0x44, 0x3b, 0xd2, 0xce, 0x76, 0xba, 0xcf, 0xcc, 0x52, 0x4b, 0x66, 0x7f, 0xf5, 0x51, 0x37, 0x07,
	0xb1, 0x89, 0x74, 0x71, 0x4e, 0x74, 0x85, 0x1b, 0x1b, 0x78, 0x4f, 0x35, 0x3c, 0x58, 0x06, 0x74,
	0xb1, 0xe0, 0x2f, 0x48, 0xfd, 0x1f, 0xf9, 0x0b, 0xdc, 0x41, 0xba, 0xe8, 0x92, 0x47, 0x8a, 0x7f,
	0xbe, 0xf9, 0x3b, 0x6b, 0x0d, 0xba, 0xba, 0xe8, 0x62, 0x1b, 0xd5, 0xaf, 0x07, 0x36, 0x69, 0x28,
	0xfe, 0x78, 0x83, 0xff, 0x50, 0xda, 0x5b, 0x37, 0xa7, 0xf1, 0x67, 0xd4, 0xe6, 0x10, 0x3a, 0x9c,
	0x85, 0xb1, 0xec, 0x7b, 0x54, 0x32, 0x4e, 0xb6, 0x8e, 0xb4, 0xb3, 0x56, 0xcf, 0xbc, 0xb9, 0x3b,
	0xac, 0xfd, 0xba, 0x3b, 0x3c, 0x8d, 0x62, 0xf9, 0x35, 0xf3, 0x4d, 0xca, 0xc6, 0x16, 0x65, 0x62,
	0xcc, 0xc4, 0xf2, 0xd1, 0x11, 0xc1, 0xd0, 0x92, 0xd3, 0x14, 0x84, 0x79, 0x09, 0xd4, 0x2d, 0x6b,
	0xf0, 0x27, 0xb4, 0xcf, 0x21, 0xbc, 0x62, 0x43, 0x48, 0x06, 0x59, 0x9a, 0x8e, 0xa6, 0x64, 0xbb,
	0x92, 0xb8, 0x64, 0xc1, 0x57, 0x68, 0xef, 0x7e, 0xc5, 0xe1, 0x31, 0x05, 0xd2, 0xac, 0xa4, 0x2d,
	0x4a, 0xf0, 0x25, 0x6a, 0xf8, 0x0e, 0x63, 0x23, 0xd2, 0xaa, 0x64, 0x5b, 0x84, 0xf1, 0x7b, 0x84,
	0x54, 0xf1, 0x31, 0x09, 0x80, 0x13, 0x54, 0x49, 0xb5, 0x66, 0xc0, 0x6f, 0x51, 0x33, 0x54, 0xbb,
	0xd9, 0x9f, 0x92, 0x9d, 0x4a, 0xb6, 0x87, 0x3c, 0x7e, 0x87, 0x5a, 0xcb, 0x7a, 0x32, 0x25, 0xbb,
	0x95, 0x64, 0x2b, 0x01, 0x76, 0xd1, 0xae, 0xe4, 0x5e, 0x10, 0x27, 0x91, 0xc3, 0xe2, 0x44, 0x92,
	0xbd, 0x4a, 0xc2, 0x82, 0x23, 0x3f, 0x59, 0x9a, 0x71, 0x0e, 0x89, 0x5c, 0x5e, 0x98, 0xfd, 0x6a,
	0x27, 0x5b, 0x90, 0xe0, 0x37, 0x68, 0x5b, 0x8c, 0x58, 0x0a, 0x4e, 0x97, 0xb4, 0x2b, 0xf9, 0xee,
	0xe3, 0x2b, 0x93, 0x4d, 0x1e, 0xff, 0x8f, 0xc9, 0xee, 0xbd, 0xbe, 0x99, 0x19, 0xda, 0xed, 0xcc,
	0xd0, 0x7e, 0xcf, 0x0c, 0xed, 0xc7, 0xdc, 0xa8, 0xdd, 0xce, 0x8d, 0xda, 0xcf, 0xb9, 0x51, 0xfb,
	0xf2, 0x72, 0x5d, 0xe5, 0x49, 0x48, 0xbc, 0x89, 0x05, 0x82, 0x76, 0x1e, 0x66, 0xdb, 0x44, 0x4d,
	0x37, 0xe5, 0xf4, 0xb7, 0xd4, 0x04, 0xb3, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x57, 0xa9, 0xce,
	0x0a, 0x9b, 0x05, 0x00, 0x00,
}

func (m *Curve) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Curve) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Curve) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.SlopeP3.Size()
		i -= size
		if _, err := m.SlopeP3.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCurve(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0x82
	{
		size := m.SlopeP2.Size()
		i -= size
		if _, err := m.SlopeP2.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCurve(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x7a
	{
		size := m.CurrentSupply.Size()
		i -= size
		if _, err := m.CurrentSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCurve(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x72
	{
		size := m.TradingPoint.Size()
		i -= size
		if _, err := m.TradingPoint.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCurve(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	{
		size := m.FactorFxy.Size()
		i -= size
		if _, err := m.FactorFxy.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCurve(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	{
		size := m.FactorFy.Size()
		i -= size
		if _, err := m.FactorFy.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCurve(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size := m.BPoolUnder.Size()
		i -= size
		if _, err := m.BPoolUnder.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCurve(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size := m.BPool.Size()
		i -= size
		if _, err := m.BPool.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCurve(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.RefTokenPrice.Size()
		i -= size
		if _, err := m.RefTokenPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCurve(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.RefTokenSupply.Size()
		i -= size
		if _, err := m.RefTokenSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCurve(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.RefProfitFactor.Size()
		i -= size
		if _, err := m.RefProfitFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCurve(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if m.QS3 != nil {
		{
			size, err := m.QS3.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCurve(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.S2 != nil {
		{
			size, err := m.S2.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCurve(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.S1 != nil {
		{
			size, err := m.S1.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCurve(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.S0 != nil {
		{
			size, err := m.S0.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCurve(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.FS0 != nil {
		{
			size, err := m.FS0.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCurve(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCurve(dAtA []byte, offset int, v uint64) int {
	offset -= sovCurve(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Curve) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FS0 != nil {
		l = m.FS0.Size()
		n += 1 + l + sovCurve(uint64(l))
	}
	if m.S0 != nil {
		l = m.S0.Size()
		n += 1 + l + sovCurve(uint64(l))
	}
	if m.S1 != nil {
		l = m.S1.Size()
		n += 1 + l + sovCurve(uint64(l))
	}
	if m.S2 != nil {
		l = m.S2.Size()
		n += 1 + l + sovCurve(uint64(l))
	}
	if m.QS3 != nil {
		l = m.QS3.Size()
		n += 1 + l + sovCurve(uint64(l))
	}
	l = m.RefProfitFactor.Size()
	n += 1 + l + sovCurve(uint64(l))
	l = m.RefTokenSupply.Size()
	n += 1 + l + sovCurve(uint64(l))
	l = m.RefTokenPrice.Size()
	n += 1 + l + sovCurve(uint64(l))
	l = m.BPool.Size()
	n += 1 + l + sovCurve(uint64(l))
	l = m.BPoolUnder.Size()
	n += 1 + l + sovCurve(uint64(l))
	l = m.FactorFy.Size()
	n += 1 + l + sovCurve(uint64(l))
	l = m.FactorFxy.Size()
	n += 1 + l + sovCurve(uint64(l))
	l = m.TradingPoint.Size()
	n += 1 + l + sovCurve(uint64(l))
	l = m.CurrentSupply.Size()
	n += 1 + l + sovCurve(uint64(l))
	l = m.SlopeP2.Size()
	n += 1 + l + sovCurve(uint64(l))
	l = m.SlopeP3.Size()
	n += 2 + l + sovCurve(uint64(l))
	return n
}

func sovCurve(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCurve(x uint64) (n int) {
	return sovCurve(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Curve) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCurve
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
			return fmt.Errorf("proto: Curve: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Curve: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FS0", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCurve
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
				return ErrInvalidLengthCurve
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCurve
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FS0 == nil {
				m.FS0 = &Flatsegment{}
			}
			if err := m.FS0.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field S0", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCurve
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
				return ErrInvalidLengthCurve
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCurve
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.S0 == nil {
				m.S0 = &BezierSegment{}
			}
			if err := m.S0.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field S1", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCurve
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
				return ErrInvalidLengthCurve
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCurve
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.S1 == nil {
				m.S1 = &BezierSegment{}
			}
			if err := m.S1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field S2", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCurve
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
				return ErrInvalidLengthCurve
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCurve
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.S2 == nil {
				m.S2 = &FixedBezierSegment{}
			}
			if err := m.S2.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QS3", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCurve
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
				return ErrInvalidLengthCurve
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCurve
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.QS3 == nil {
				m.QS3 = &Quadraticsegment{}
			}
			if err := m.QS3.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefProfitFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCurve
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
				return ErrInvalidLengthCurve
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCurve
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RefProfitFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefTokenSupply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCurve
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
				return ErrInvalidLengthCurve
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCurve
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RefTokenSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefTokenPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCurve
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
				return ErrInvalidLengthCurve
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCurve
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RefTokenPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BPool", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCurve
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
				return ErrInvalidLengthCurve
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCurve
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BPool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BPoolUnder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCurve
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
				return ErrInvalidLengthCurve
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCurve
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BPoolUnder.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FactorFy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCurve
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
				return ErrInvalidLengthCurve
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCurve
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FactorFy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FactorFxy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCurve
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
				return ErrInvalidLengthCurve
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCurve
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FactorFxy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradingPoint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCurve
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
				return ErrInvalidLengthCurve
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCurve
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TradingPoint.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentSupply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCurve
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
				return ErrInvalidLengthCurve
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCurve
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlopeP2", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCurve
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
				return ErrInvalidLengthCurve
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCurve
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlopeP2.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlopeP3", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCurve
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
				return ErrInvalidLengthCurve
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCurve
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlopeP3.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCurve(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCurve
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
func skipCurve(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCurve
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
					return 0, ErrIntOverflowCurve
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
					return 0, ErrIntOverflowCurve
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
				return 0, ErrInvalidLengthCurve
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCurve
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCurve
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCurve        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCurve          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCurve = fmt.Errorf("proto: unexpected end of group")
)
