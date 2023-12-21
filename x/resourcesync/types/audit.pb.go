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
// source: escbackbone/resourcesync/audit.proto

package types

import (
	fmt "fmt"
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

type ActionLog int32

const (
	ActionLog_ACTION_NOT_SET  ActionLog = 0
	ActionLog_RESOURCE_CREATE ActionLog = 1
	ActionLog_RESOURCE_UPDATE ActionLog = 2
	ActionLog_RESOURCE_DELETE ActionLog = 3
)

var ActionLog_name = map[int32]string{
	0: "ACTION_NOT_SET",
	1: "RESOURCE_CREATE",
	2: "RESOURCE_UPDATE",
	3: "RESOURCE_DELETE",
}

var ActionLog_value = map[string]int32{
	"ACTION_NOT_SET":  0,
	"RESOURCE_CREATE": 1,
	"RESOURCE_UPDATE": 2,
	"RESOURCE_DELETE": 3,
}

func (x ActionLog) String() string {
	return proto.EnumName(ActionLog_name, int32(x))
}

func (ActionLog) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3df3f7dc30997e93, []int{0}
}

type Audit struct {
	Action    ActionLog  `protobuf:"varint,1,opt,name=action,proto3,enum=escbackbone.resourcesync.ActionLog" json:"action,omitempty"`
	Signature *Signature `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Audit) Reset()         { *m = Audit{} }
func (m *Audit) String() string { return proto.CompactTextString(m) }
func (*Audit) ProtoMessage()    {}
func (*Audit) Descriptor() ([]byte, []int) {
	return fileDescriptor_3df3f7dc30997e93, []int{0}
}
func (m *Audit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Audit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Audit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Audit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Audit.Merge(m, src)
}
func (m *Audit) XXX_Size() int {
	return m.Size()
}
func (m *Audit) XXX_DiscardUnknown() {
	xxx_messageInfo_Audit.DiscardUnknown(m)
}

var xxx_messageInfo_Audit proto.InternalMessageInfo

func (m *Audit) GetAction() ActionLog {
	if m != nil {
		return m.Action
	}
	return ActionLog_ACTION_NOT_SET
}

func (m *Audit) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Signature struct {
	Pubkey string `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
}

func (m *Signature) Reset()         { *m = Signature{} }
func (m *Signature) String() string { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()    {}
func (*Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_3df3f7dc30997e93, []int{1}
}
func (m *Signature) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Signature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Signature.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Signature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signature.Merge(m, src)
}
func (m *Signature) XXX_Size() int {
	return m.Size()
}
func (m *Signature) XXX_DiscardUnknown() {
	xxx_messageInfo_Signature.DiscardUnknown(m)
}

var xxx_messageInfo_Signature proto.InternalMessageInfo

func (m *Signature) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

func init() {
	proto.RegisterEnum("escbackbone.resourcesync.ActionLog", ActionLog_name, ActionLog_value)
	proto.RegisterType((*Audit)(nil), "escbackbone.resourcesync.Audit")
	proto.RegisterType((*Signature)(nil), "escbackbone.resourcesync.Signature")
}

func init() {
	proto.RegisterFile("escbackbone/resourcesync/audit.proto", fileDescriptor_3df3f7dc30997e93)
}

var fileDescriptor_3df3f7dc30997e93 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0x2d, 0x4e, 0x4e,
	0x4a, 0x4c, 0xce, 0x4e, 0xca, 0xcf, 0x4b, 0xd5, 0x2f, 0x4a, 0x2d, 0xce, 0x2f, 0x2d, 0x4a, 0x4e,
	0x2d, 0xae, 0xcc, 0x4b, 0xd6, 0x4f, 0x2c, 0x4d, 0xc9, 0x2c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x92, 0x40, 0x52, 0xa5, 0x87, 0xac, 0x4a, 0xa9, 0x9d, 0x91, 0x8b, 0xd5, 0x11, 0xa4, 0x52,
	0xc8, 0x9a, 0x8b, 0x2d, 0x31, 0xb9, 0x24, 0x33, 0x3f, 0x4f, 0x82, 0x51, 0x81, 0x51, 0x83, 0xcf,
	0x48, 0x59, 0x0f, 0x97, 0x26, 0x3d, 0x47, 0xb0, 0x3a, 0x9f, 0xfc, 0xf4, 0x20, 0xa8, 0x16, 0x21,
	0x47, 0x2e, 0xce, 0xe2, 0xcc, 0xf4, 0xbc, 0xc4, 0x92, 0xd2, 0xa2, 0x54, 0x09, 0x26, 0x05, 0x46,
	0x0d, 0x6e, 0x7c, 0xfa, 0x83, 0x61, 0x4a, 0x83, 0x10, 0xba, 0x94, 0x94, 0xb9, 0x38, 0xe1, 0xe2,
	0x42, 0x62, 0x5c, 0x6c, 0x05, 0xa5, 0x49, 0xd9, 0xa9, 0x95, 0x60, 0xc7, 0x70, 0x06, 0x41, 0x79,
	0x5a, 0x71, 0x5c, 0x9c, 0x70, 0xcb, 0x85, 0x84, 0xb8, 0xf8, 0x1c, 0x9d, 0x43, 0x3c, 0xfd, 0xfd,
	0xe2, 0xfd, 0xfc, 0x43, 0xe2, 0x83, 0x5d, 0x43, 0x04, 0x18, 0x84, 0x84, 0xb9, 0xf8, 0x83, 0x5c,
	0x83, 0xfd, 0x43, 0x83, 0x9c, 0x5d, 0xe3, 0x9d, 0x83, 0x5c, 0x1d, 0x43, 0x5c, 0x05, 0x18, 0x51,
	0x04, 0x43, 0x03, 0x5c, 0x40, 0x82, 0x4c, 0x28, 0x82, 0x2e, 0xae, 0x3e, 0xae, 0x21, 0xae, 0x02,
	0xcc, 0x4e, 0x7e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3,
	0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x92, 0x9e,
	0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x9c, 0x58, 0x92, 0x9a, 0x97, 0x58,
	0xa1, 0x9f, 0x5a, 0x9c, 0xac, 0x0b, 0x0f, 0xfc, 0x0a, 0xd4, 0xe0, 0x2f, 0xa9, 0x2c, 0x48, 0x2d,
	0x4e, 0x62, 0x03, 0x87, 0xbf, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x4c, 0x07, 0x85, 0xa7,
	0x01, 0x00, 0x00,
}

func (m *Audit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Audit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Audit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Signature != nil {
		{
			size, err := m.Signature.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAudit(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Action != 0 {
		i = encodeVarintAudit(dAtA, i, uint64(m.Action))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Signature) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Signature) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Signature) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Pubkey) > 0 {
		i -= len(m.Pubkey)
		copy(dAtA[i:], m.Pubkey)
		i = encodeVarintAudit(dAtA, i, uint64(len(m.Pubkey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAudit(dAtA []byte, offset int, v uint64) int {
	offset -= sovAudit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Audit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Action != 0 {
		n += 1 + sovAudit(uint64(m.Action))
	}
	if m.Signature != nil {
		l = m.Signature.Size()
		n += 1 + l + sovAudit(uint64(l))
	}
	return n
}

func (m *Signature) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Pubkey)
	if l > 0 {
		n += 1 + l + sovAudit(uint64(l))
	}
	return n
}

func sovAudit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAudit(x uint64) (n int) {
	return sovAudit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Audit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAudit
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
			return fmt.Errorf("proto: Audit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Audit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			m.Action = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAudit
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Action |= ActionLog(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAudit
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
				return ErrInvalidLengthAudit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAudit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Signature == nil {
				m.Signature = &Signature{}
			}
			if err := m.Signature.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAudit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAudit
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
func (m *Signature) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAudit
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
			return fmt.Errorf("proto: Signature: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Signature: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAudit
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
				return ErrInvalidLengthAudit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAudit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pubkey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAudit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAudit
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
func skipAudit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAudit
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
					return 0, ErrIntOverflowAudit
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
					return 0, ErrIntOverflowAudit
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
				return 0, ErrInvalidLengthAudit
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAudit
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAudit
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAudit        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAudit          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAudit = fmt.Errorf("proto: unexpected end of group")
)
