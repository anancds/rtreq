// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

/*
Package msg is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	BasicMessage
*/
package msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BasicMessage struct {
	Sender  string `protobuf:"bytes,1,opt,name=sender" json:"sender,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *BasicMessage) Reset()                    { *m = BasicMessage{} }
func (m *BasicMessage) String() string            { return proto.CompactTextString(m) }
func (*BasicMessage) ProtoMessage()               {}
func (*BasicMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BasicMessage) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *BasicMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*BasicMessage)(nil), "msg.BasicMessage")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 91 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2d, 0x4e, 0x57, 0x72,
	0xe0, 0xe2, 0x71, 0x4a, 0x2c, 0xce, 0x4c, 0xf6, 0x85, 0x48, 0x09, 0x89, 0x71, 0xb1, 0x15, 0xa7,
	0xe6, 0xa5, 0xa4, 0x16, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x79, 0x42, 0x12, 0x5c,
	0xec, 0x50, 0xdd, 0x12, 0x4c, 0x60, 0x09, 0x18, 0x37, 0x89, 0x0d, 0x6c, 0x9a, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0xf8, 0xac, 0x16, 0xde, 0x5e, 0x00, 0x00, 0x00,
}
