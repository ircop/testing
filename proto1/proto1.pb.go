// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto1.proto

package proto1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	proto3enum "../proto3enum"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Enum1 struct {
	TestE                proto3enum.TestEnum `protobuf:"varint,1,opt,name=testE,proto3,enum=proto3enum.TestEnum" json:"testE,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Enum1) Reset()         { *m = Enum1{} }
func (m *Enum1) String() string { return proto.CompactTextString(m) }
func (*Enum1) ProtoMessage()    {}
func (*Enum1) Descriptor() ([]byte, []int) {
	return fileDescriptor_077598b0e4a82848, []int{0}
}

func (m *Enum1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Enum1.Unmarshal(m, b)
}
func (m *Enum1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Enum1.Marshal(b, m, deterministic)
}
func (m *Enum1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Enum1.Merge(m, src)
}
func (m *Enum1) XXX_Size() int {
	return xxx_messageInfo_Enum1.Size(m)
}
func (m *Enum1) XXX_DiscardUnknown() {
	xxx_messageInfo_Enum1.DiscardUnknown(m)
}

var xxx_messageInfo_Enum1 proto.InternalMessageInfo

func (m *Enum1) GetTestE() proto3enum.TestEnum {
	if m != nil {
		return m.TestE
	}
	return proto3enum.TestEnum_FOO
}

func init() {
	proto.RegisterType((*Enum1)(nil), "proto1.Enum1")
}

func init() { proto.RegisterFile("proto1.proto", fileDescriptor_077598b0e4a82848) }

var fileDescriptor_077598b0e4a82848 = []byte{
	// 94 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0x37, 0xd4, 0x03, 0x53, 0x42, 0x6c, 0x10, 0x9e, 0x94, 0x10, 0x98, 0x36, 0x4e, 0xcd, 0x2b,
	0xcd, 0xd5, 0x4f, 0x85, 0xc8, 0x29, 0x19, 0x73, 0xb1, 0xba, 0xe6, 0x95, 0xe6, 0x1a, 0x0a, 0x69,
	0x71, 0xb1, 0x96, 0xa4, 0x16, 0x97, 0xb8, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x19, 0x89, 0xe8,
	0x21, 0x14, 0xeb, 0x85, 0x80, 0x24, 0xf2, 0x4a, 0x73, 0x83, 0x20, 0x4a, 0x92, 0x20, 0x06, 0x1a,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x08, 0x6d, 0x71, 0xde, 0x67, 0x00, 0x00, 0x00,
}
