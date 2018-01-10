// Code generated by protoc-gen-go. DO NOT EDIT.
// source: LandlordsProto.proto

/*
Package msg is a generated protocol buffer package.

It is generated from these files:
	LandlordsProto.proto

It has these top-level messages:
	CS_PlayerLogin
	SC_PlayerLogin
	CS_PlayerRegister
	SC_PlayerRegister
	CS_Heartbeat
	SC_Heartbeat
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

type CS_PlayerLogin struct {
	Account string `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
	Pwd     string `protobuf:"bytes,2,opt,name=pwd" json:"pwd,omitempty"`
}

func (m *CS_PlayerLogin) Reset()                    { *m = CS_PlayerLogin{} }
func (m *CS_PlayerLogin) String() string            { return proto.CompactTextString(m) }
func (*CS_PlayerLogin) ProtoMessage()               {}
func (*CS_PlayerLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CS_PlayerLogin) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *CS_PlayerLogin) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type SC_PlayerLogin struct {
	// 0成功 1账号错误 2密码错误
	Result int32 `protobuf:"zigzag32,1,opt,name=result" json:"result,omitempty"`
}

func (m *SC_PlayerLogin) Reset()                    { *m = SC_PlayerLogin{} }
func (m *SC_PlayerLogin) String() string            { return proto.CompactTextString(m) }
func (*SC_PlayerLogin) ProtoMessage()               {}
func (*SC_PlayerLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SC_PlayerLogin) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type CS_PlayerRegister struct {
	Account string `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
	Pwd     string `protobuf:"bytes,2,opt,name=pwd" json:"pwd,omitempty"`
}

func (m *CS_PlayerRegister) Reset()                    { *m = CS_PlayerRegister{} }
func (m *CS_PlayerRegister) String() string            { return proto.CompactTextString(m) }
func (*CS_PlayerRegister) ProtoMessage()               {}
func (*CS_PlayerRegister) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CS_PlayerRegister) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *CS_PlayerRegister) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type SC_PlayerRegister struct {
	// 0成功 1账号错误
	Result int32 `protobuf:"zigzag32,1,opt,name=result" json:"result,omitempty"`
}

func (m *SC_PlayerRegister) Reset()                    { *m = SC_PlayerRegister{} }
func (m *SC_PlayerRegister) String() string            { return proto.CompactTextString(m) }
func (*SC_PlayerRegister) ProtoMessage()               {}
func (*SC_PlayerRegister) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SC_PlayerRegister) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type CS_Heartbeat struct {
}

func (m *CS_Heartbeat) Reset()                    { *m = CS_Heartbeat{} }
func (m *CS_Heartbeat) String() string            { return proto.CompactTextString(m) }
func (*CS_Heartbeat) ProtoMessage()               {}
func (*CS_Heartbeat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type SC_Heartbeat struct {
	ServerTime int64 `protobuf:"zigzag64,1,opt,name=serverTime" json:"serverTime,omitempty"`
}

func (m *SC_Heartbeat) Reset()                    { *m = SC_Heartbeat{} }
func (m *SC_Heartbeat) String() string            { return proto.CompactTextString(m) }
func (*SC_Heartbeat) ProtoMessage()               {}
func (*SC_Heartbeat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SC_Heartbeat) GetServerTime() int64 {
	if m != nil {
		return m.ServerTime
	}
	return 0
}

func init() {
	proto.RegisterType((*CS_PlayerLogin)(nil), "msg.CS_PlayerLogin")
	proto.RegisterType((*SC_PlayerLogin)(nil), "msg.SC_PlayerLogin")
	proto.RegisterType((*CS_PlayerRegister)(nil), "msg.CS_PlayerRegister")
	proto.RegisterType((*SC_PlayerRegister)(nil), "msg.SC_PlayerRegister")
	proto.RegisterType((*CS_Heartbeat)(nil), "msg.CS_Heartbeat")
	proto.RegisterType((*SC_Heartbeat)(nil), "msg.SC_Heartbeat")
}

func init() { proto.RegisterFile("LandlordsProto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xf1, 0x49, 0xcc, 0x4b,
	0xc9, 0xc9, 0x2f, 0x4a, 0x29, 0x0e, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2b, 0x00, 0x91, 0x42, 0xcc,
	0xb9, 0xc5, 0xe9, 0x4a, 0x36, 0x5c, 0x7c, 0xce, 0xc1, 0xf1, 0x01, 0x39, 0x89, 0x95, 0xa9, 0x45,
	0x3e, 0xf9, 0xe9, 0x99, 0x79, 0x42, 0x12, 0x5c, 0xec, 0x89, 0xc9, 0xc9, 0xf9, 0xa5, 0x79, 0x25,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x90, 0x00, 0x17, 0x73, 0x41, 0x79, 0x8a,
	0x04, 0x13, 0x58, 0x14, 0xc4, 0x54, 0xd2, 0xe0, 0xe2, 0x0b, 0x76, 0x46, 0xd1, 0x2d, 0xc6, 0xc5,
	0x56, 0x94, 0x5a, 0x5c, 0x9a, 0x03, 0xd1, 0x2c, 0x18, 0x04, 0xe5, 0x29, 0xd9, 0x73, 0x09, 0xc2,
	0xed, 0x09, 0x4a, 0x4d, 0xcf, 0x2c, 0x2e, 0x49, 0x2d, 0x22, 0xc9, 0x2a, 0x6d, 0x2e, 0x41, 0xb8,
	0x55, 0x70, 0x03, 0x70, 0xd9, 0xc6, 0xc7, 0xc5, 0xe3, 0x1c, 0x1c, 0xef, 0x91, 0x9a, 0x58, 0x54,
	0x92, 0x94, 0x9a, 0x58, 0xa2, 0xa4, 0xc7, 0xc5, 0x13, 0xec, 0x8c, 0xe0, 0x0b, 0xc9, 0x71, 0x71,
	0x15, 0xa7, 0x16, 0x95, 0xa5, 0x16, 0x85, 0x64, 0xe6, 0xa6, 0x82, 0xf5, 0x0a, 0x05, 0x21, 0x89,
	0x24, 0xb1, 0x81, 0x43, 0xc8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x33, 0x3d, 0xa1, 0x92, 0x39,
	0x01, 0x00, 0x00,
}
