// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: escbackbone/ubc/segment.proto

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

type Segment struct {
	P0     *github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=p0,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"p0,omitempty"`
	A      *github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=a,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"a,omitempty"`
	B      *github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=b,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"b,omitempty"`
	P1     *github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=p1,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"p1,omitempty"`
	P0X    *github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=p0X,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"p0X,omitempty"`
	P1X    *github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=p1X,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"p1X,omitempty"`
	DelatX *github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=delatX,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"delatX,omitempty"`
}

func (m *Segment) Reset()         { *m = Segment{} }
func (m *Segment) String() string { return proto.CompactTextString(m) }
func (*Segment) ProtoMessage()    {}
func (*Segment) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7293bbfb2f004b1, []int{0}
}
func (m *Segment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Segment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Segment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Segment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Segment.Merge(m, src)
}
func (m *Segment) XXX_Size() int {
	return m.Size()
}
func (m *Segment) XXX_DiscardUnknown() {
	xxx_messageInfo_Segment.DiscardUnknown(m)
}

var xxx_messageInfo_Segment proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Segment)(nil), "escbackbone.ubc.Segment")
}

func init() { proto.RegisterFile("escbackbone/ubc/segment.proto", fileDescriptor_c7293bbfb2f004b1) }

var fileDescriptor_c7293bbfb2f004b1 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x2d, 0x4e, 0x4e,
	0x4a, 0x4c, 0xce, 0x4e, 0xca, 0xcf, 0x4b, 0xd5, 0x2f, 0x4d, 0x4a, 0xd6, 0x2f, 0x4e, 0x4d, 0xcf,
	0x4d, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x47, 0x92, 0xd6, 0x2b, 0x4d,
	0x4a, 0x96, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xcb, 0xe9, 0x83, 0x58, 0x10, 0x65, 0x4a, 0x5b,
	0x99, 0xb9, 0xd8, 0x83, 0x21, 0x1a, 0x85, 0xac, 0xb8, 0x98, 0x0a, 0x0c, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x9d, 0xb4, 0x6e, 0xdd, 0x93, 0x57, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b,
	0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0x86, 0x52, 0xba, 0xc5, 0x29, 0xd9, 0xfa,
	0x25, 0x95, 0x05, 0xa9, 0xc5, 0x7a, 0x2e, 0xa9, 0xc9, 0x41, 0x4c, 0x05, 0x06, 0x42, 0x16, 0x5c,
	0x8c, 0x89, 0x12, 0x4c, 0x24, 0x6b, 0x65, 0x4c, 0x04, 0xe9, 0x4c, 0x92, 0x60, 0x26, 0x5d, 0x67,
	0x12, 0xd8, 0xbd, 0x86, 0x12, 0x2c, 0x64, 0xb8, 0xd7, 0x50, 0xc8, 0x86, 0x8b, 0xb9, 0xc0, 0x20,
	0x42, 0x82, 0x95, 0x64, 0xcd, 0x20, 0x6d, 0x60, 0xdd, 0x86, 0x11, 0x12, 0x6c, 0x64, 0xe8, 0x36,
	0x8c, 0x10, 0x72, 0xe2, 0x62, 0x4b, 0x49, 0xcd, 0x49, 0x2c, 0x89, 0x90, 0x60, 0x27, 0xd9, 0x00,
	0xa8, 0x4e, 0x27, 0xd7, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e,
	0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x46,
	0x36, 0x29, 0xb1, 0x24, 0x35, 0x2f, 0xb1, 0x42, 0x3f, 0xb5, 0x38, 0x59, 0x17, 0x9e, 0x56, 0x2a,
	0xc0, 0xa9, 0x05, 0x6c, 0x64, 0x12, 0x1b, 0x38, 0x15, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xbe, 0x44, 0xaa, 0x8c, 0x4d, 0x02, 0x00, 0x00,
}

func (m *Segment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Segment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Segment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DelatX != nil {
		{
			size := m.DelatX.Size()
			i -= size
			if _, err := m.DelatX.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintSegment(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.P1X != nil {
		{
			size := m.P1X.Size()
			i -= size
			if _, err := m.P1X.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintSegment(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.P0X != nil {
		{
			size := m.P0X.Size()
			i -= size
			if _, err := m.P0X.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintSegment(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.P1 != nil {
		{
			size := m.P1.Size()
			i -= size
			if _, err := m.P1.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintSegment(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.B != nil {
		{
			size := m.B.Size()
			i -= size
			if _, err := m.B.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintSegment(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.A != nil {
		{
			size := m.A.Size()
			i -= size
			if _, err := m.A.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintSegment(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.P0 != nil {
		{
			size := m.P0.Size()
			i -= size
			if _, err := m.P0.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintSegment(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSegment(dAtA []byte, offset int, v uint64) int {
	offset -= sovSegment(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Segment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.P0 != nil {
		l = m.P0.Size()
		n += 1 + l + sovSegment(uint64(l))
	}
	if m.A != nil {
		l = m.A.Size()
		n += 1 + l + sovSegment(uint64(l))
	}
	if m.B != nil {
		l = m.B.Size()
		n += 1 + l + sovSegment(uint64(l))
	}
	if m.P1 != nil {
		l = m.P1.Size()
		n += 1 + l + sovSegment(uint64(l))
	}
	if m.P0X != nil {
		l = m.P0X.Size()
		n += 1 + l + sovSegment(uint64(l))
	}
	if m.P1X != nil {
		l = m.P1X.Size()
		n += 1 + l + sovSegment(uint64(l))
	}
	if m.DelatX != nil {
		l = m.DelatX.Size()
		n += 1 + l + sovSegment(uint64(l))
	}
	return n
}

func sovSegment(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSegment(x uint64) (n int) {
	return sovSegment(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Segment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSegment
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
			return fmt.Errorf("proto: Segment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Segment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field P0", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSegment
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
				return ErrInvalidLengthSegment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSegment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.P0 = &v
			if err := m.P0.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSegment
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
				return ErrInvalidLengthSegment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSegment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.A = &v
			if err := m.A.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSegment
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
				return ErrInvalidLengthSegment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSegment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.B = &v
			if err := m.B.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field P1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSegment
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
				return ErrInvalidLengthSegment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSegment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.P1 = &v
			if err := m.P1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field P0X", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSegment
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
				return ErrInvalidLengthSegment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSegment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.P0X = &v
			if err := m.P0X.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field P1X", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSegment
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
				return ErrInvalidLengthSegment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSegment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.P1X = &v
			if err := m.P1X.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelatX", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSegment
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
				return ErrInvalidLengthSegment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSegment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.DelatX = &v
			if err := m.DelatX.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSegment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSegment
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
func skipSegment(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSegment
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
					return 0, ErrIntOverflowSegment
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
					return 0, ErrIntOverflowSegment
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
				return 0, ErrInvalidLengthSegment
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSegment
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSegment
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSegment        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSegment          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSegment = fmt.Errorf("proto: unexpected end of group")
)
