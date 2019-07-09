// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envelope.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type Envelope struct {
	InnerMessage         *any.Any `protobuf:"bytes,1,opt,name=innerMessage,proto3" json:"innerMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Envelope) Reset()         { *m = Envelope{} }
func (m *Envelope) String() string { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()    {}
func (*Envelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee266e8c558e9dc5, []int{0}
}

func (m *Envelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Envelope.Unmarshal(m, b)
}
func (m *Envelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Envelope.Marshal(b, m, deterministic)
}
func (m *Envelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Envelope.Merge(m, src)
}
func (m *Envelope) XXX_Size() int {
	return xxx_messageInfo_Envelope.Size(m)
}
func (m *Envelope) XXX_DiscardUnknown() {
	xxx_messageInfo_Envelope.DiscardUnknown(m)
}

var xxx_messageInfo_Envelope proto.InternalMessageInfo

func (m *Envelope) GetInnerMessage() *any.Any {
	if m != nil {
		return m.InnerMessage
	}
	return nil
}

func init() {
	proto.RegisterType((*Envelope)(nil), "pb.Envelope")
}

func init() { proto.RegisterFile("envelope.proto", fileDescriptor_ee266e8c558e9dc5) }

var fileDescriptor_ee266e8c558e9dc5 = []byte{
	// 113 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xcd, 0x2b, 0x4b,
	0xcd, 0xc9, 0x2f, 0x48, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0x92,
	0x4c, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x24, 0x95, 0xa6, 0xe9, 0x27, 0xe6, 0x55,
	0x42, 0xa4, 0x95, 0x5c, 0xb8, 0x38, 0x5c, 0xa1, 0x1a, 0x84, 0x2c, 0xb8, 0x78, 0x32, 0xf3, 0xf2,
	0x52, 0x8b, 0x7c, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d,
	0x44, 0xf4, 0x20, 0xba, 0xf5, 0x60, 0xba, 0xf5, 0x1c, 0xf3, 0x2a, 0x83, 0x50, 0x54, 0x26, 0xb1,
	0x81, 0xe5, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x68, 0x59, 0x9e, 0xdf, 0x7d, 0x00, 0x00,
	0x00,
}
