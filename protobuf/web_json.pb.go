// Code generated by protoc-gen-go. DO NOT EDIT.
// source: web_json.proto

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	web_json.proto

It has these top-level messages:
	WebJSON
*/
package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// WebJSON : web json with google/protobuf/any.proto
type WebJSON struct {
	Code    int32                `protobuf:"varint,1,opt,name=Code" json:"Code,omitempty"`
	Message string               `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
	Data    *google_protobuf.Any `protobuf:"bytes,3,opt,name=Data" json:"Data,omitempty"`
}

func (m *WebJSON) Reset()                    { *m = WebJSON{} }
func (m *WebJSON) String() string            { return proto.CompactTextString(m) }
func (*WebJSON) ProtoMessage()               {}
func (*WebJSON) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *WebJSON) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *WebJSON) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *WebJSON) GetData() *google_protobuf.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*WebJSON)(nil), "protobuf.WebJSON")
}

func init() { proto.RegisterFile("web_json.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4f, 0x4d, 0x8a,
	0xcf, 0x2a, 0xce, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x49, 0xa5,
	0x69, 0x52, 0x92, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x30, 0x01, 0xfd, 0xc4, 0xbc, 0x4a,
	0x88, 0x22, 0xa5, 0x44, 0x2e, 0xf6, 0xf0, 0xd4, 0x24, 0xaf, 0x60, 0x7f, 0x3f, 0x21, 0x21, 0x2e,
	0x16, 0xe7, 0xfc, 0x94, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x30, 0x5b, 0x48, 0x82,
	0x8b, 0xdd, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x3d, 0x55, 0x82, 0x49, 0x81, 0x51, 0x83, 0x33, 0x08,
	0xc6, 0x15, 0xd2, 0xe0, 0x62, 0x71, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x56, 0x60, 0xd4, 0xe0, 0x36,
	0x12, 0xd1, 0x83, 0x58, 0xa1, 0x07, 0xb3, 0x42, 0xcf, 0x31, 0xaf, 0x32, 0x08, 0xac, 0xc2, 0x89,
	0x2b, 0x0a, 0xee, 0x92, 0x24, 0x36, 0x30, 0xcb, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xe8, 0xf1,
	0x91, 0x9d, 0xac, 0x00, 0x00, 0x00,
}
