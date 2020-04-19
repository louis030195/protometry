// Code generated by protoc-gen-go. DO NOT EDIT.
// source: quaternionn.proto

package protometry

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type QuaternionN struct {
	Value                *VectorN `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuaternionN) Reset()         { *m = QuaternionN{} }
func (m *QuaternionN) String() string { return proto.CompactTextString(m) }
func (*QuaternionN) ProtoMessage()    {}
func (*QuaternionN) Descriptor() ([]byte, []int) {
	return fileDescriptor_09a15d631415a433, []int{0}
}

func (m *QuaternionN) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuaternionN.Unmarshal(m, b)
}
func (m *QuaternionN) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuaternionN.Marshal(b, m, deterministic)
}
func (m *QuaternionN) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuaternionN.Merge(m, src)
}
func (m *QuaternionN) XXX_Size() int {
	return xxx_messageInfo_QuaternionN.Size(m)
}
func (m *QuaternionN) XXX_DiscardUnknown() {
	xxx_messageInfo_QuaternionN.DiscardUnknown(m)
}

var xxx_messageInfo_QuaternionN proto.InternalMessageInfo

func (m *QuaternionN) GetValue() *VectorN {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*QuaternionN)(nil), "protometry.QuaternionN")
}

func init() {
	proto.RegisterFile("quaternionn.proto", fileDescriptor_09a15d631415a433)
}

var fileDescriptor_09a15d631415a433 = []byte{
	// 122 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x2c, 0x4d, 0x2c,
	0x49, 0x2d, 0xca, 0xcb, 0xcc, 0xcf, 0xcb, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x02,
	0x53, 0xb9, 0xa9, 0x25, 0x45, 0x95, 0x52, 0xbc, 0x65, 0xa9, 0xc9, 0x25, 0xf9, 0x45, 0x50, 0x29,
	0x25, 0x0b, 0x2e, 0xee, 0x40, 0xb8, 0x7a, 0x3f, 0x21, 0x4d, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2,
	0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x61, 0x3d, 0x84, 0x4e, 0xbd, 0x30, 0xb0, 0x46,
	0xbf, 0x20, 0x88, 0x0a, 0x27, 0x6d, 0x2e, 0xbe, 0xe4, 0xfc, 0x5c, 0x24, 0x05, 0x4e, 0x5c, 0x01,
	0x70, 0x76, 0x00, 0xe3, 0x2a, 0x26, 0x24, 0x6e, 0x12, 0x1b, 0x58, 0x99, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x71, 0x10, 0x9f, 0x47, 0x9d, 0x00, 0x00, 0x00,
}
