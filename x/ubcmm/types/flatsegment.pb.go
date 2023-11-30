// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: escbackbone/ubcmm/flatsegment.proto

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

// FlatSegment represents a horizontal line with a dynamic shape and bounded
// interval.
//
// The line has zero length during initialization and could be modified during
// shift up and undergird operations.
//
// The line has a bounded interval [0, p1X].
type FlatSegment struct {
	Y   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=y,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"y"`
	P1X github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=p1X,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"p1X"`
}

func (m *FlatSegment) Reset()         { *m = FlatSegment{} }
func (m *FlatSegment) String() string { return proto.CompactTextString(m) }
func (*FlatSegment) ProtoMessage()    {}
func (*FlatSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_d91dc8a479bea90d, []int{0}
}
func (m *FlatSegment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FlatSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FlatSegment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FlatSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlatSegment.Merge(m, src)
}
func (m *FlatSegment) XXX_Size() int {
	return m.Size()
}
func (m *FlatSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_FlatSegment.DiscardUnknown(m)
}

var xxx_messageInfo_FlatSegment proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FlatSegment)(nil), "escbackbone.ubcmm.FlatSegment")
}

func init() {
	proto.RegisterFile("escbackbone/ubcmm/flatsegment.proto", fileDescriptor_d91dc8a479bea90d)
}

var fileDescriptor_d91dc8a479bea90d = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0x2d, 0x4e, 0x4e,
	0x4a, 0x4c, 0xce, 0x4e, 0xca, 0xcf, 0x4b, 0xd5, 0x2f, 0x4d, 0x4a, 0xce, 0xcd, 0xd5, 0x4f, 0xcb,
	0x49, 0x2c, 0x29, 0x4e, 0x4d, 0xcf, 0x4d, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x44, 0x52, 0xa4, 0x07, 0x56, 0x24, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x96, 0xd5, 0x07,
	0xb1, 0x20, 0x0a, 0x95, 0x7a, 0x19, 0xb9, 0xb8, 0xdd, 0x72, 0x12, 0x4b, 0x82, 0x21, 0xda, 0x85,
	0x6c, 0xb8, 0x18, 0x2b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x9d, 0xf4, 0x4e, 0xdc, 0x93, 0x67,
	0xb8, 0x75, 0x4f, 0x5e, 0x2d, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f,
	0x39, 0xbf, 0x38, 0x37, 0xbf, 0x18, 0x4a, 0xe9, 0x16, 0xa7, 0x64, 0xeb, 0x97, 0x54, 0x16, 0xa4,
	0x16, 0xeb, 0xb9, 0xa4, 0x26, 0x07, 0x31, 0x56, 0x0a, 0x39, 0x70, 0x31, 0x17, 0x18, 0x46, 0x48,
	0x30, 0x91, 0xa5, 0x1f, 0xa4, 0xd5, 0xc9, 0xfd, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18,
	0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5,
	0x18, 0xa2, 0x74, 0x91, 0x8d, 0x49, 0x2c, 0x49, 0xcd, 0x4b, 0xac, 0xd0, 0x4f, 0x2d, 0x4e, 0xd6,
	0x85, 0x87, 0x45, 0x05, 0x34, 0x34, 0xc0, 0x26, 0x26, 0xb1, 0x81, 0xfd, 0x67, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x50, 0x49, 0xf9, 0x7f, 0x2f, 0x01, 0x00, 0x00,
}

func (m *FlatSegment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FlatSegment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FlatSegment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.P1X.Size()
		i -= size
		if _, err := m.P1X.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintFlatsegment(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Y.Size()
		i -= size
		if _, err := m.Y.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintFlatsegment(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintFlatsegment(dAtA []byte, offset int, v uint64) int {
	offset -= sovFlatsegment(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FlatSegment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Y.Size()
	n += 1 + l + sovFlatsegment(uint64(l))
	l = m.P1X.Size()
	n += 1 + l + sovFlatsegment(uint64(l))
	return n
}

func sovFlatsegment(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFlatsegment(x uint64) (n int) {
	return sovFlatsegment(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FlatSegment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFlatsegment
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
			return fmt.Errorf("proto: FlatSegment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FlatSegment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Y", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlatsegment
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
				return ErrInvalidLengthFlatsegment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFlatsegment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Y.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field P1X", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFlatsegment
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
				return ErrInvalidLengthFlatsegment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFlatsegment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.P1X.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFlatsegment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFlatsegment
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
func skipFlatsegment(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFlatsegment
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
					return 0, ErrIntOverflowFlatsegment
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
					return 0, ErrIntOverflowFlatsegment
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
				return 0, ErrInvalidLengthFlatsegment
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFlatsegment
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFlatsegment
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFlatsegment        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFlatsegment          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFlatsegment = fmt.Errorf("proto: unexpected end of group")
)
