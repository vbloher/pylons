// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pylons/fighter.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
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

type Fighter struct {
	Creator    string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	ID         uint64 `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	CookbookID string `protobuf:"bytes,3,opt,name=cookbookID,proto3" json:"cookbookID,omitempty"`
	LHitem     string `protobuf:"bytes,4,opt,name=LHitem,proto3" json:"LHitem,omitempty"`
	RHitem     string `protobuf:"bytes,5,opt,name=RHitem,proto3" json:"RHitem,omitempty"`
	Armoritem  string `protobuf:"bytes,6,opt,name=Armoritem,proto3" json:"Armoritem,omitempty"`
	Status     string `protobuf:"bytes,7,opt,name=Status,proto3" json:"Status,omitempty"`
	Log        string `protobuf:"bytes,8,opt,name=Log,proto3" json:"Log,omitempty"`
}

func (m *Fighter) Reset()         { *m = Fighter{} }
func (m *Fighter) String() string { return proto.CompactTextString(m) }
func (*Fighter) ProtoMessage()    {}
func (*Fighter) Descriptor() ([]byte, []int) {
	return fileDescriptor_93f2c4e6ff2ec5f7, []int{0}
}
func (m *Fighter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Fighter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Fighter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Fighter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fighter.Merge(m, src)
}
func (m *Fighter) XXX_Size() int {
	return m.Size()
}
func (m *Fighter) XXX_DiscardUnknown() {
	xxx_messageInfo_Fighter.DiscardUnknown(m)
}

var xxx_messageInfo_Fighter proto.InternalMessageInfo

func (m *Fighter) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Fighter) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Fighter) GetCookbookID() string {
	if m != nil {
		return m.CookbookID
	}
	return ""
}

func (m *Fighter) GetLHitem() string {
	if m != nil {
		return m.LHitem
	}
	return ""
}

func (m *Fighter) GetRHitem() string {
	if m != nil {
		return m.RHitem
	}
	return ""
}

func (m *Fighter) GetArmoritem() string {
	if m != nil {
		return m.Armoritem
	}
	return ""
}

func (m *Fighter) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Fighter) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

func init() {
	proto.RegisterType((*Fighter)(nil), "Pylonstech.pylons.pylons.Fighter")
}

func init() { proto.RegisterFile("pylons/fighter.proto", fileDescriptor_93f2c4e6ff2ec5f7) }

var fileDescriptor_93f2c4e6ff2ec5f7 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x1c, 0xc5, 0x9b, 0x6e, 0xb6, 0x2e, 0x07, 0x91, 0x30, 0x24, 0x0c, 0x09, 0xc3, 0xd3, 0x0e, 0xda,
	0x30, 0xfc, 0x04, 0x4a, 0x19, 0x16, 0x76, 0x90, 0x7a, 0xf3, 0xd6, 0x86, 0xd8, 0x96, 0xd9, 0xfd,
	0x4b, 0x92, 0x89, 0xfb, 0x16, 0x7e, 0x2c, 0xc1, 0xcb, 0x8e, 0x1e, 0xa5, 0xfd, 0x22, 0xd2, 0xa4,
	0x43, 0x4f, 0x79, 0xef, 0xf7, 0xfe, 0x8f, 0xc0, 0xc3, 0xd3, 0x66, 0xff, 0x0a, 0x5b, 0xcd, 0x5f,
	0xaa, 0xa2, 0x34, 0x52, 0x45, 0x8d, 0x02, 0x03, 0x84, 0x3e, 0x5a, 0x6a, 0xa4, 0x28, 0x23, 0x77,
	0x30, 0x3c, 0x33, 0x26, 0x40, 0xd7, 0xa0, 0x79, 0x9e, 0x69, 0xc9, 0xdf, 0x96, 0xb9, 0x34, 0xd9,
	0x92, 0x0b, 0xa8, 0xb6, 0xae, 0x39, 0x9b, 0x16, 0x50, 0x80, 0x95, 0xbc, 0x57, 0x8e, 0x5e, 0x7d,
	0x21, 0x1c, 0xae, 0xdc, 0x0f, 0x84, 0xe2, 0x50, 0x28, 0x99, 0x19, 0x50, 0x14, 0xcd, 0xd1, 0x62,
	0x92, 0x1e, 0x2d, 0x39, 0xc3, 0x7e, 0x12, 0x53, 0x7f, 0x8e, 0x16, 0xe3, 0xd4, 0x4f, 0x62, 0xc2,
	0x30, 0x16, 0x00, 0x9b, 0x1c, 0x60, 0x93, 0xc4, 0x74, 0x64, 0x8f, 0xff, 0x11, 0x72, 0x81, 0x83,
	0xf5, 0x43, 0x65, 0x64, 0x4d, 0xc7, 0x36, 0x1b, 0x5c, 0xcf, 0x53, 0xc7, 0x4f, 0x1c, 0x77, 0x8e,
	0x5c, 0xe2, 0xc9, 0x9d, 0xaa, 0x41, 0xd9, 0x28, 0xb0, 0xd1, 0x1f, 0xe8, 0x5b, 0x4f, 0x26, 0x33,
	0x3b, 0x4d, 0x43, 0xd7, 0x72, 0x8e, 0x9c, 0xe3, 0xd1, 0x1a, 0x0a, 0x7a, 0x6a, 0x61, 0x2f, 0xef,
	0x57, 0x9f, 0x2d, 0x43, 0x87, 0x96, 0xa1, 0x9f, 0x96, 0xa1, 0x8f, 0x8e, 0x79, 0x87, 0x8e, 0x79,
	0xdf, 0x1d, 0xf3, 0x9e, 0xaf, 0x8b, 0xca, 0x94, 0xbb, 0x3c, 0x12, 0x50, 0x73, 0x37, 0xe1, 0x4d,
	0xbf, 0x21, 0x1f, 0x46, 0x7e, 0x3f, 0x0a, 0xb3, 0x6f, 0xa4, 0xce, 0x03, 0x3b, 0xce, 0xed, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x01, 0xc2, 0xe1, 0x3e, 0x84, 0x01, 0x00, 0x00,
}

func (m *Fighter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Fighter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Fighter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Log) > 0 {
		i -= len(m.Log)
		copy(dAtA[i:], m.Log)
		i = encodeVarintFighter(dAtA, i, uint64(len(m.Log)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintFighter(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Armoritem) > 0 {
		i -= len(m.Armoritem)
		copy(dAtA[i:], m.Armoritem)
		i = encodeVarintFighter(dAtA, i, uint64(len(m.Armoritem)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.RHitem) > 0 {
		i -= len(m.RHitem)
		copy(dAtA[i:], m.RHitem)
		i = encodeVarintFighter(dAtA, i, uint64(len(m.RHitem)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.LHitem) > 0 {
		i -= len(m.LHitem)
		copy(dAtA[i:], m.LHitem)
		i = encodeVarintFighter(dAtA, i, uint64(len(m.LHitem)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CookbookID) > 0 {
		i -= len(m.CookbookID)
		copy(dAtA[i:], m.CookbookID)
		i = encodeVarintFighter(dAtA, i, uint64(len(m.CookbookID)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ID != 0 {
		i = encodeVarintFighter(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintFighter(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFighter(dAtA []byte, offset int, v uint64) int {
	offset -= sovFighter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Fighter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovFighter(uint64(l))
	}
	if m.ID != 0 {
		n += 1 + sovFighter(uint64(m.ID))
	}
	l = len(m.CookbookID)
	if l > 0 {
		n += 1 + l + sovFighter(uint64(l))
	}
	l = len(m.LHitem)
	if l > 0 {
		n += 1 + l + sovFighter(uint64(l))
	}
	l = len(m.RHitem)
	if l > 0 {
		n += 1 + l + sovFighter(uint64(l))
	}
	l = len(m.Armoritem)
	if l > 0 {
		n += 1 + l + sovFighter(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovFighter(uint64(l))
	}
	l = len(m.Log)
	if l > 0 {
		n += 1 + l + sovFighter(uint64(l))
	}
	return n
}

func sovFighter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFighter(x uint64) (n int) {
	return sovFighter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Fighter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFighter
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
			return fmt.Errorf("proto: Fighter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fighter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFighter
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
				return ErrInvalidLengthFighter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFighter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFighter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CookbookID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFighter
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
				return ErrInvalidLengthFighter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFighter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CookbookID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LHitem", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFighter
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
				return ErrInvalidLengthFighter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFighter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LHitem = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RHitem", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFighter
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
				return ErrInvalidLengthFighter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFighter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RHitem = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Armoritem", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFighter
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
				return ErrInvalidLengthFighter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFighter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Armoritem = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFighter
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
				return ErrInvalidLengthFighter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFighter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Log", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFighter
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
				return ErrInvalidLengthFighter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFighter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Log = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFighter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFighter
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
func skipFighter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFighter
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
					return 0, ErrIntOverflowFighter
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
					return 0, ErrIntOverflowFighter
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
				return 0, ErrInvalidLengthFighter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFighter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFighter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFighter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFighter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFighter = fmt.Errorf("proto: unexpected end of group")
)