// Copyright (c) 2022-2023 Contributors to the Eclipse Foundation
//
// See the NOTICE file(s) distributed with this work for additional
// information regarding copyright ownership.
//
// This program and the accompanying materials are made available under the
// terms of the Apache License, Version 2.0 which is available at
// https://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: escbackbone/ubcmm/curve.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Curve represents the undergirding bonding curve (ubc), consisting of
// multiple segments.
type Curve struct {
	S0                        *FlatSegment                           `protobuf:"bytes,1,opt,name=s0,proto3" json:"s0,omitempty"`
	S1                        *BezierSegment                         `protobuf:"bytes,2,opt,name=s1,proto3" json:"s1,omitempty"`
	S2                        *BezierSegment                         `protobuf:"bytes,3,opt,name=s2,proto3" json:"s2,omitempty"`
	S3                        *FixedBezierSegment                    `protobuf:"bytes,4,opt,name=s3,proto3" json:"s3,omitempty"`
	S4                        *FixedQuadraticSegment                 `protobuf:"bytes,5,opt,name=s4,proto3" json:"s4,omitempty"`
	RefProfitFactor           github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=refProfitFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"refProfitFactor"`
	RefTokenSupply            github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=refTokenSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"refTokenSupply"`
	RefTokenPrice             github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=refTokenPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"refTokenPrice"`
	BPool                     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,9,opt,name=bPool,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"bPool"`
	BPoolUnder                github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,10,opt,name=bPoolUnder,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"bPoolUnder"`
	FactorFy                  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,11,opt,name=factorFy,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"factorFy"`
	FactorFxy                 github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,12,opt,name=factorFxy,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"factorFxy"`
	TradingPoint              github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,13,opt,name=tradingPoint,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"tradingPoint"`
	CurrentSupply             github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,14,opt,name=currentSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"currentSupply"`
	SlopeP2                   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,15,opt,name=slopeP2,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slopeP2"`
	SlopeP3                   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,16,opt,name=slopeP3,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slopeP3"`
	NumericalErrorAccumulator github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,17,opt,name=numericalErrorAccumulator,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"numericalErrorAccumulator"`
	Segments                  Segments                               `protobuf:"bytes,18,opt,name=segments,proto3,customtype=Segments" json:"segments"`
}

func (m *Curve) Reset()         { *m = Curve{} }
func (m *Curve) String() string { return proto.CompactTextString(m) }
func (*Curve) ProtoMessage()    {}
func (*Curve) Descriptor() ([]byte, []int) {
	return fileDescriptor_399925d0dbf34629, []int{0}
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

func (m *Curve) GetS0() *FlatSegment {
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

func (m *Curve) GetS2() *BezierSegment {
	if m != nil {
		return m.S2
	}
	return nil
}

func (m *Curve) GetS3() *FixedBezierSegment {
	if m != nil {
		return m.S3
	}
	return nil
}

func (m *Curve) GetS4() *FixedQuadraticSegment {
	if m != nil {
		return m.S4
	}
	return nil
}

func init() {
	proto.RegisterType((*Curve)(nil), "escbackbone.ubcmm.Curve")
}

func init() { proto.RegisterFile("escbackbone/ubcmm/curve.proto", fileDescriptor_399925d0dbf34629) }

var fileDescriptor_399925d0dbf34629 = []byte{
	// 551 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xdd, 0x6e, 0x12, 0x41,
	0x14, 0xc7, 0xd9, 0x55, 0x5a, 0x98, 0x7e, 0x4f, 0xbc, 0x18, 0x9b, 0xb8, 0x25, 0x9a, 0x1a, 0x62,
	0x64, 0x97, 0x42, 0x4d, 0xbc, 0x15, 0x2b, 0x1a, 0x63, 0x0c, 0xd2, 0x6a, 0x8c, 0x77, 0xb3, 0xc3,
	0x01, 0x37, 0xec, 0xee, 0xe0, 0xcc, 0xac, 0x01, 0x9f, 0xc2, 0x27, 0xf0, 0x79, 0x7a, 0xd9, 0x4b,
	0xe3, 0x45, 0x63, 0xe0, 0x45, 0x0c, 0xc3, 0xf2, 0xb1, 0x40, 0x13, 0xb3, 0xbd, 0xe2, 0x64, 0xe6,
	0xff, 0xfb, 0x71, 0x66, 0x06, 0x0e, 0x7a, 0x00, 0x92, 0xb9, 0x94, 0x75, 0x5d, 0x1e, 0x82, 0x13,
	0xb9, 0x2c, 0x08, 0x1c, 0x16, 0x89, 0xef, 0x60, 0xf7, 0x04, 0x57, 0x1c, 0x1f, 0x2c, 0x6c, 0xdb,
	0x7a, 0xfb, 0xf0, 0x78, 0x95, 0x70, 0xe1, 0x87, 0x07, 0x42, 0x42, 0x27, 0x80, 0x50, 0x4d, 0xc8,
	0xc3, 0xe2, 0x6a, 0xec, 0x5b, 0x44, 0x5b, 0x82, 0x2a, 0x8f, 0x25, 0x93, 0x8f, 0x56, 0x93, 0x6d,
	0x9f, 0xaa, 0x64, 0xe8, 0xc9, 0x9a, 0x90, 0xd7, 0x87, 0xd6, 0xba, 0xaf, 0xbe, 0xd7, 0xe1, 0x1d,
	0xae, 0x4b, 0x67, 0x5c, 0x4d, 0x56, 0x1f, 0xfe, 0x42, 0x28, 0xfb, 0x72, 0x7c, 0x34, 0x6c, 0x23,
	0x53, 0x96, 0x89, 0x51, 0x30, 0x8a, 0x5b, 0x15, 0xcb, 0x5e, 0x39, 0xa1, 0x5d, 0xf7, 0xa9, 0x3a,
	0x9f, 0x18, 0x9b, 0xa6, 0x2c, 0xe3, 0x32, 0x32, 0xe5, 0x09, 0x31, 0x75, 0xbe, 0xb0, 0x26, 0x5f,
	0xd3, 0x3d, 0xcc, 0x89, 0x13, 0x4d, 0x54, 0xc8, 0x9d, 0xff, 0x26, 0x2a, 0xf8, 0x19, 0x32, 0x65,
	0x95, 0xdc, 0xd5, 0xc4, 0xf1, 0xba, 0x9e, 0xc6, 0x87, 0x5d, 0xc6, 0xaa, 0xf8, 0x39, 0x32, 0xe5,
	0x29, 0xc9, 0x6a, 0xac, 0x78, 0x13, 0xf6, 0x61, 0x7a, 0xef, 0x73, 0xf2, 0x14, 0x7f, 0x46, 0x7b,
	0x02, 0xda, 0x0d, 0xc1, 0xdb, 0x9e, 0xaa, 0x53, 0xa6, 0xb8, 0x20, 0x1b, 0x05, 0xa3, 0x98, 0xaf,
	0xd9, 0x97, 0xd7, 0x47, 0x99, 0x3f, 0xd7, 0x47, 0x8f, 0x3b, 0x9e, 0xfa, 0x1a, 0xb9, 0x36, 0xe3,
	0x81, 0xc3, 0xb8, 0x0c, 0xb8, 0x8c, 0x3f, 0x4a, 0xb2, 0xd5, 0x75, 0xd4, 0xa0, 0x07, 0xd2, 0x3e,
	0x03, 0xd6, 0x5c, 0xd6, 0xe0, 0x4f, 0x68, 0x57, 0x40, 0xfb, 0x82, 0x77, 0x21, 0x3c, 0x8f, 0x7a,
	0x3d, 0x7f, 0x40, 0x36, 0x53, 0x89, 0x97, 0x2c, 0xf8, 0x02, 0xed, 0x4c, 0x57, 0x1a, 0xc2, 0x63,
	0x40, 0x72, 0xa9, 0xb4, 0x49, 0x09, 0x3e, 0x43, 0x59, 0xb7, 0xc1, 0xb9, 0x4f, 0xf2, 0xa9, 0x6c,
	0x13, 0x18, 0xbf, 0x47, 0x48, 0x17, 0x1f, 0xc3, 0x16, 0x08, 0x82, 0x52, 0xa9, 0x16, 0x0c, 0xf8,
	0x2d, 0xca, 0xb5, 0xf5, 0x6d, 0xd6, 0x07, 0x64, 0x2b, 0x95, 0x6d, 0xc6, 0xe3, 0x77, 0x28, 0x1f,
	0xd7, 0xfd, 0x01, 0xd9, 0x4e, 0x25, 0x9b, 0x0b, 0x70, 0x13, 0x6d, 0x2b, 0x41, 0x5b, 0x5e, 0xd8,
	0x69, 0x70, 0x2f, 0x54, 0x64, 0x27, 0x95, 0x30, 0xe1, 0x18, 0xbf, 0x2c, 0x8b, 0x84, 0x80, 0x50,
	0xc5, 0x3f, 0x98, 0xdd, 0x74, 0x2f, 0x9b, 0x90, 0xe0, 0x37, 0x68, 0x53, 0xfa, 0xbc, 0x07, 0x8d,
	0x0a, 0xd9, 0x4b, 0xe5, 0x9b, 0xe2, 0x73, 0x53, 0x95, 0xec, 0xdf, 0xc6, 0x54, 0xc5, 0x3e, 0xba,
	0x1f, 0x46, 0x01, 0x08, 0x8f, 0x51, 0xff, 0x95, 0x10, 0x5c, 0xbc, 0x60, 0x2c, 0x0a, 0x22, 0x9f,
	0x8e, 0xff, 0x7f, 0x07, 0xa9, 0xdc, 0x37, 0x0b, 0xf1, 0x53, 0x94, 0x8b, 0x27, 0xa3, 0x24, 0x58,
	0xcb, 0xf7, 0x63, 0x79, 0x2e, 0x1e, 0x05, 0xb2, 0x39, 0x4b, 0xd4, 0x5e, 0x5f, 0x0e, 0x2d, 0xe3,
	0x6a, 0x68, 0x19, 0x7f, 0x87, 0x96, 0xf1, 0x73, 0x64, 0x65, 0xae, 0x46, 0x56, 0xe6, 0xf7, 0xc8,
	0xca, 0x7c, 0x29, 0x2d, 0xb6, 0x42, 0x15, 0x84, 0xb4, 0xef, 0x80, 0x64, 0xa5, 0xd9, 0x40, 0xee,
	0xc7, 0x23, 0x59, 0x77, 0xe5, 0x6e, 0xe8, 0x81, 0x5b, 0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0xef,
	0x86, 0x9b, 0xb9, 0x5c, 0x06, 0x00, 0x00,
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
		size := m.Segments.Size()
		i -= size
		if _, err := m.Segments.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCurve(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0x92
	{
		size := m.NumericalErrorAccumulator.Size()
		i -= size
		if _, err := m.NumericalErrorAccumulator.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCurve(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0x8a
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
	if m.S4 != nil {
		{
			size, err := m.S4.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCurve(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.S3 != nil {
		{
			size, err := m.S3.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCurve(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
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
		dAtA[i] = 0x1a
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
		dAtA[i] = 0x12
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
	if m.S3 != nil {
		l = m.S3.Size()
		n += 1 + l + sovCurve(uint64(l))
	}
	if m.S4 != nil {
		l = m.S4.Size()
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
	l = m.NumericalErrorAccumulator.Size()
	n += 2 + l + sovCurve(uint64(l))
	l = m.Segments.Size()
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
				m.S0 = &FlatSegment{}
			}
			if err := m.S0.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
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
		case 3:
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
				m.S2 = &BezierSegment{}
			}
			if err := m.S2.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field S3", wireType)
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
			if m.S3 == nil {
				m.S3 = &FixedBezierSegment{}
			}
			if err := m.S3.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field S4", wireType)
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
			if m.S4 == nil {
				m.S4 = &FixedQuadraticSegment{}
			}
			if err := m.S4.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumericalErrorAccumulator", wireType)
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
			if err := m.NumericalErrorAccumulator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 18:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Segments", wireType)
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
			if err := m.Segments.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
