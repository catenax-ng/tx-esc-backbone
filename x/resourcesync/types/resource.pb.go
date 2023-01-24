// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: escbackbone/resourcesync/resource.proto

package types

import (
	fmt "fmt"
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

type ReceiptLog int32

const (
	ReceiptLog_RECEIPT_NOT_SET ReceiptLog = 0
	ReceiptLog_RECEIPT_APPLIED ReceiptLog = 1
)

var ReceiptLog_name = map[int32]string{
	0: "RECEIPT_NOT_SET",
	1: "RECEIPT_APPLIED",
}

var ReceiptLog_value = map[string]int32{
	"RECEIPT_NOT_SET": 0,
	"RECEIPT_APPLIED": 1,
}

func (x ReceiptLog) String() string {
	return proto.EnumName(ReceiptLog_name, int32(x))
}

func (ReceiptLog) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3149da7e1850b6f5, []int{0}
}

type Resource struct {
	// issue of resource change
	Originator string `protobuf:"bytes,1,opt,name=originator,proto3" json:"originator,omitempty"`
	// id of the resource by originator -  unique per originator
	OrigResId string `protobuf:"bytes,2,opt,name=origResId,proto3" json:"origResId,omitempty"`
	// pointer to the system holding the information of the resource
	TargetSystem string `protobuf:"bytes,3,opt,name=targetSystem,proto3" json:"targetSystem,omitempty"`
	// Id of the resource to access it at the target system
	ResourceKey string `protobuf:"bytes,4,opt,name=resourceKey,proto3" json:"resourceKey,omitempty"`
	// Hash of the resource
	DataHash []byte `protobuf:"bytes,5,opt,name=dataHash,proto3" json:"dataHash,omitempty"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_3149da7e1850b6f5, []int{0}
}
func (m *Resource) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return m.Size()
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("escbackbone.resourcesync.ReceiptLog", ReceiptLog_name, ReceiptLog_value)
	proto.RegisterType((*Resource)(nil), "escbackbone.resourcesync.Resource")
}

func init() {
	proto.RegisterFile("escbackbone/resourcesync/resource.proto", fileDescriptor_3149da7e1850b6f5)
}

var fileDescriptor_3149da7e1850b6f5 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4f, 0x2d, 0x4e, 0x4e,
	0x4a, 0x4c, 0xce, 0x4e, 0xca, 0xcf, 0x4b, 0xd5, 0x2f, 0x4a, 0x2d, 0xce, 0x2f, 0x2d, 0x4a, 0x4e,
	0x2d, 0xae, 0xcc, 0x4b, 0x86, 0x73, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x24, 0x90, 0x14,
	0xea, 0x21, 0x2b, 0x94, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x2b, 0xd2, 0x07, 0xb1, 0x20, 0xea,
	0xa5, 0x54, 0x70, 0x1a, 0x9c, 0x58, 0x9a, 0x92, 0x59, 0x02, 0x51, 0xa5, 0xb4, 0x81, 0x91, 0x8b,
	0x23, 0x08, 0x2a, 0x29, 0x24, 0xc7, 0xc5, 0x95, 0x5f, 0x94, 0x99, 0x9e, 0x99, 0x97, 0x58, 0x92,
	0x5f, 0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x24, 0x22, 0x24, 0xc3, 0xc5, 0x09, 0xe2,
	0x05, 0xa5, 0x16, 0x7b, 0xa6, 0x48, 0x30, 0x81, 0xa5, 0x11, 0x02, 0x42, 0x4a, 0x5c, 0x3c, 0x25,
	0x89, 0x45, 0xe9, 0xa9, 0x25, 0xc1, 0x95, 0xc5, 0x25, 0xa9, 0xb9, 0x12, 0xcc, 0x60, 0x05, 0x28,
	0x62, 0x42, 0x0a, 0x5c, 0xdc, 0x30, 0xa7, 0x78, 0xa7, 0x56, 0x4a, 0xb0, 0x80, 0x95, 0x20, 0x0b,
	0x09, 0x49, 0x71, 0x71, 0xa4, 0x24, 0x96, 0x24, 0x7a, 0x24, 0x16, 0x67, 0x48, 0xb0, 0x2a, 0x30,
	0x6a, 0xf0, 0x04, 0xc1, 0xf9, 0x56, 0x2c, 0x1d, 0x0b, 0xe4, 0x19, 0xb4, 0xcc, 0xb8, 0xb8, 0x82,
	0x52, 0x93, 0x53, 0x33, 0x0b, 0x4a, 0x7c, 0xf2, 0xd3, 0x85, 0x84, 0xb9, 0xf8, 0x83, 0x5c, 0x9d,
	0x5d, 0x3d, 0x03, 0x42, 0xe2, 0xfd, 0xfc, 0x43, 0xe2, 0x83, 0x5d, 0x43, 0x04, 0x18, 0x90, 0x05,
	0x1d, 0x03, 0x02, 0x7c, 0x3c, 0x5d, 0x5d, 0x04, 0x18, 0x9d, 0xfc, 0x4e, 0x3c, 0x92, 0x63, 0xbc,
	0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63,
	0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x24, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f,
	0x57, 0x3f, 0x39, 0xb1, 0x24, 0x35, 0x2f, 0xb1, 0x42, 0x3f, 0xb5, 0x38, 0x59, 0x17, 0x1e, 0x7c,
	0x15, 0xa8, 0x01, 0x58, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x0e, 0x41, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x42, 0x0e, 0xfd, 0x73, 0xc2, 0x01, 0x00, 0x00,
}

func (m *Resource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Resource) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Resource) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DataHash) > 0 {
		i -= len(m.DataHash)
		copy(dAtA[i:], m.DataHash)
		i = encodeVarintResource(dAtA, i, uint64(len(m.DataHash)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ResourceKey) > 0 {
		i -= len(m.ResourceKey)
		copy(dAtA[i:], m.ResourceKey)
		i = encodeVarintResource(dAtA, i, uint64(len(m.ResourceKey)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.TargetSystem) > 0 {
		i -= len(m.TargetSystem)
		copy(dAtA[i:], m.TargetSystem)
		i = encodeVarintResource(dAtA, i, uint64(len(m.TargetSystem)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OrigResId) > 0 {
		i -= len(m.OrigResId)
		copy(dAtA[i:], m.OrigResId)
		i = encodeVarintResource(dAtA, i, uint64(len(m.OrigResId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Originator) > 0 {
		i -= len(m.Originator)
		copy(dAtA[i:], m.Originator)
		i = encodeVarintResource(dAtA, i, uint64(len(m.Originator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintResource(dAtA []byte, offset int, v uint64) int {
	offset -= sovResource(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Resource) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Originator)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	l = len(m.OrigResId)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	l = len(m.TargetSystem)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	l = len(m.ResourceKey)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	l = len(m.DataHash)
	if l > 0 {
		n += 1 + l + sovResource(uint64(l))
	}
	return n
}

func sovResource(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozResource(x uint64) (n int) {
	return sovResource(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Resource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResource
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
			return fmt.Errorf("proto: Resource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Resource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Originator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Originator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrigResId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrigResId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetSystem", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TargetSystem = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
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
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataHash = append(m.DataHash[:0], dAtA[iNdEx:postIndex]...)
			if m.DataHash == nil {
				m.DataHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthResource
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
func skipResource(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResource
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
					return 0, ErrIntOverflowResource
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
					return 0, ErrIntOverflowResource
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
				return 0, ErrInvalidLengthResource
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupResource
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthResource
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthResource        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResource          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupResource = fmt.Errorf("proto: unexpected end of group")
)
