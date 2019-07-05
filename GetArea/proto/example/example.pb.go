// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_GetArea

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	//返回错误码
	Errno string `protobuf:"bytes,1,opt,name=Errno,proto3" json:"Errno,omitempty"`
	//返回错误信息
	Errmsg               string              `protobuf:"bytes,2,opt,name=Errmsg,proto3" json:"Errmsg,omitempty"`
	Data                 []*Response_Address `protobuf:"bytes,3,rep,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *Response) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *Response) GetData() []*Response_Address {
	if m != nil {
		return m.Data
	}
	return nil
}

//返回数据的类型
type Response_Address struct {
	Aid                  int32    `protobuf:"varint,1,opt,name=aid,proto3" json:"aid,omitempty"`
	Aname                string   `protobuf:"bytes,2,opt,name=aname,proto3" json:"aname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response_Address) Reset()         { *m = Response_Address{} }
func (m *Response_Address) String() string { return proto.CompactTextString(m) }
func (*Response_Address) ProtoMessage()    {}
func (*Response_Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{1, 0}
}

func (m *Response_Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_Address.Unmarshal(m, b)
}
func (m *Response_Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_Address.Marshal(b, m, deterministic)
}
func (m *Response_Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_Address.Merge(m, src)
}
func (m *Response_Address) XXX_Size() int {
	return xxx_messageInfo_Response_Address.Size(m)
}
func (m *Response_Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Response_Address proto.InternalMessageInfo

func (m *Response_Address) GetAid() int32 {
	if m != nil {
		return m.Aid
	}
	return 0
}

func (m *Response_Address) GetAname() string {
	if m != nil {
		return m.Aname
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.GetArea.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.GetArea.Response")
	proto.RegisterType((*Response_Address)(nil), "go.micro.srv.GetArea.Response.Address")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4b, 0x87, 0x30,
	0x18, 0xc6, 0x33, 0xd3, 0xe5, 0xdb, 0x25, 0x86, 0x84, 0x18, 0x85, 0x78, 0x08, 0x4f, 0x8b, 0xec,
	0xd6, 0x4d, 0x48, 0x82, 0x8e, 0x83, 0x3e, 0xc0, 0xca, 0x17, 0x11, 0x9a, 0xb3, 0x77, 0x2b, 0xfa,
	0x46, 0x7d, 0xcd, 0x70, 0xce, 0x5b, 0xfd, 0x4f, 0xdb, 0xf3, 0xc0, 0x9e, 0xdf, 0xf3, 0x0c, 0x2e,
	0x17, 0x32, 0xce, 0xdc, 0xe2, 0xb7, 0xd2, 0xcb, 0x3b, 0xee, 0xa7, 0xf0, 0x2e, 0xcf, 0x47, 0x23,
	0xf4, 0xf4, 0x46, 0x46, 0x58, 0xfa, 0x12, 0x4f, 0xe8, 0x3a, 0x42, 0x55, 0x67, 0xc0, 0x24, 0x7e,
	0x7c, 0xa2, 0x75, 0xf5, 0x4f, 0x04, 0xa7, 0x12, 0xed, 0x62, 0x66, 0x8b, 0x3c, 0x87, 0xa4, 0x27,
	0x9a, 0x4d, 0x11, 0x55, 0x51, 0x93, 0xc9, 0x4d, 0xf0, 0x0b, 0x48, 0x7b, 0x22, 0x6d, 0xc7, 0xe2,
	0xd8, 0xdb, 0x41, 0xf1, 0x07, 0x38, 0x79, 0x54, 0x4e, 0x15, 0x71, 0x15, 0x37, 0x67, 0xed, 0x8d,
	0xf8, 0x0b, 0x25, 0xf6, 0x6c, 0xd1, 0x0d, 0x03, 0xa1, 0xb5, 0xd2, 0xbf, 0x29, 0xef, 0x80, 0x05,
	0x83, 0x9f, 0x43, 0xac, 0xa6, 0xc1, 0x23, 0x13, 0xb9, 0x5e, 0xd7, 0x1a, 0x6a, 0x56, 0x1a, 0x03,
	0x6f, 0x13, 0xed, 0x0b, 0xb0, 0x7e, 0xdb, 0xc6, 0x9f, 0x81, 0x85, 0x7c, 0x7e, 0xf5, 0x1f, 0xd6,
	0xcf, 0x2b, 0xaf, 0x0f, 0xb7, 0xaa, 0x8f, 0x5e, 0x53, 0xff, 0x51, 0xf7, 0xbf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x5d, 0x52, 0xc2, 0x85, 0x47, 0x01, 0x00, 0x00,
}