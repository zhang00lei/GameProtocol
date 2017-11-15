// Code generated by protoc-gen-go. DO NOT EDIT.
// source: login.proto

/*
Package platform is a generated protocol buffer package.

It is generated from these files:
	login.proto

It has these top-level messages:
	RequestLogin
	ResponseLogin
	RequestServerList
	ResponseServerList
*/
package platform

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

type RequestLogin struct {
	Account          *string `protobuf:"bytes,1,req,name=account" json:"account,omitempty"`
	Password         *string `protobuf:"bytes,2,req,name=password" json:"password,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RequestLogin) Reset()                    { *m = RequestLogin{} }
func (m *RequestLogin) String() string            { return proto.CompactTextString(m) }
func (*RequestLogin) ProtoMessage()               {}
func (*RequestLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RequestLogin) GetAccount() string {
	if m != nil && m.Account != nil {
		return *m.Account
	}
	return ""
}

func (m *RequestLogin) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

type ResponseLogin struct {
	Error            *int32  `protobuf:"zigzag32,1,req,name=error" json:"error,omitempty"`
	Key              *string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ResponseLogin) Reset()                    { *m = ResponseLogin{} }
func (m *ResponseLogin) String() string            { return proto.CompactTextString(m) }
func (*ResponseLogin) ProtoMessage()               {}
func (*ResponseLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ResponseLogin) GetError() int32 {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return 0
}

func (m *ResponseLogin) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

type RequestServerList struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RequestServerList) Reset()                    { *m = RequestServerList{} }
func (m *RequestServerList) String() string            { return proto.CompactTextString(m) }
func (*RequestServerList) ProtoMessage()               {}
func (*RequestServerList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ResponseServerList struct {
	SrvInfo          []*ResponseServerListServer `protobuf:"bytes,1,rep,name=srv_info,json=srvInfo" json:"srv_info,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *ResponseServerList) Reset()                    { *m = ResponseServerList{} }
func (m *ResponseServerList) String() string            { return proto.CompactTextString(m) }
func (*ResponseServerList) ProtoMessage()               {}
func (*ResponseServerList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ResponseServerList) GetSrvInfo() []*ResponseServerListServer {
	if m != nil {
		return m.SrvInfo
	}
	return nil
}

type ResponseServerListServer struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Ip               *string `protobuf:"bytes,2,opt,name=ip" json:"ip,omitempty"`
	Port             *int32  `protobuf:"zigzag32,3,opt,name=port" json:"port,omitempty"`
	Status           *int32  `protobuf:"zigzag32,4,opt,name=status" json:"status,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ResponseServerListServer) Reset()                    { *m = ResponseServerListServer{} }
func (m *ResponseServerListServer) String() string            { return proto.CompactTextString(m) }
func (*ResponseServerListServer) ProtoMessage()               {}
func (*ResponseServerListServer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

func (m *ResponseServerListServer) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ResponseServerListServer) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *ResponseServerListServer) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func (m *ResponseServerListServer) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*RequestLogin)(nil), "platform.request_login")
	proto.RegisterType((*ResponseLogin)(nil), "platform.response_login")
	proto.RegisterType((*RequestServerList)(nil), "platform.request_server_list")
	proto.RegisterType((*ResponseServerList)(nil), "platform.response_server_list")
	proto.RegisterType((*ResponseServerListServer)(nil), "platform.response_server_list.server")
}

func init() { proto.RegisterFile("login.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x31, 0x4f, 0xf3, 0x30,
	0x10, 0x86, 0xe5, 0xa4, 0x5f, 0x9b, 0x5e, 0xf5, 0x55, 0xd4, 0x14, 0x64, 0x75, 0x8a, 0x22, 0x21,
	0x65, 0xca, 0xc0, 0xc4, 0xc8, 0xc2, 0xc0, 0x9a, 0x89, 0x2d, 0xb2, 0x8a, 0x83, 0x2c, 0x52, 0x9f,
	0xb9, 0x73, 0x82, 0xf8, 0x5b, 0xfc, 0x42, 0x54, 0x27, 0xa9, 0x18, 0xd8, 0x9e, 0xe7, 0xee, 0xfc,
	0xea, 0x35, 0x6c, 0x3a, 0x7c, 0xb3, 0xae, 0xf2, 0x84, 0x01, 0x65, 0xe6, 0x3b, 0x1d, 0x5a, 0xa4,
	0x53, 0xf1, 0x04, 0xff, 0xc9, 0x7c, 0xf4, 0x86, 0x43, 0x13, 0x0f, 0xa4, 0x82, 0x95, 0x3e, 0x1e,
	0xb1, 0x77, 0x41, 0x89, 0x3c, 0x29, 0xd7, 0xf5, 0xac, 0xf2, 0x00, 0x99, 0xd7, 0xcc, 0x9f, 0x48,
	0xaf, 0x2a, 0x89, 0xab, 0x8b, 0x17, 0x0f, 0xb0, 0x25, 0xc3, 0x1e, 0x1d, 0x9b, 0x29, 0x67, 0x0f,
	0xff, 0x0c, 0x11, 0x52, 0x4c, 0xd9, 0xd5, 0xa3, 0xc8, 0x2b, 0x48, 0xdf, 0xcd, 0x97, 0x4a, 0x72,
	0x51, 0xae, 0xeb, 0x33, 0x16, 0x37, 0x70, 0x3d, 0x17, 0x60, 0x43, 0x83, 0xa1, 0xa6, 0xb3, 0x1c,
	0x8a, 0x6f, 0x01, 0xfb, 0x4b, 0xe2, 0xaf, 0x85, 0x7c, 0x84, 0x8c, 0x69, 0x68, 0xac, 0x6b, 0x51,
	0x89, 0x3c, 0x2d, 0x37, 0xf7, 0x77, 0xd5, 0xfc, 0x9b, 0xea, 0xaf, 0x17, 0xd5, 0xc8, 0xf5, 0x8a,
	0x69, 0x78, 0x76, 0x2d, 0x1e, 0x5e, 0x60, 0x39, 0x8e, 0xa4, 0x84, 0x85, 0xd3, 0x27, 0xa3, 0x44,
	0xac, 0x13, 0x59, 0x6e, 0x21, 0xb1, 0x7e, 0x2a, 0x98, 0x58, 0x7f, 0xbe, 0xf1, 0x48, 0x41, 0xa5,
	0xb9, 0x28, 0x77, 0x75, 0x64, 0x79, 0x0b, 0x4b, 0x0e, 0x3a, 0xf4, 0xac, 0x16, 0x71, 0x3a, 0xd9,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x04, 0x61, 0xa2, 0x5e, 0x64, 0x01, 0x00, 0x00,
}
