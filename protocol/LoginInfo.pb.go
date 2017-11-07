// Code generated by protoc-gen-go. DO NOT EDIT.
// source: LoginInfo.proto

/*
Package PlayerLogin is a generated protocol buffer package.

It is generated from these files:
	LoginInfo.proto

It has these top-level messages:
	CS_PlayerLogin
	SC_PlayerLogin
	CS_PlayerRegister
	SC_PlayerRegister
*/
package PlayerLogin

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

type SC_PlayerLogin_LoginResult int32

const (
	SC_PlayerLogin_SUCCESS       SC_PlayerLogin_LoginResult = 0
	SC_PlayerLogin_ACCOUNT_ERROR SC_PlayerLogin_LoginResult = 1
	SC_PlayerLogin_PWD_ERROR     SC_PlayerLogin_LoginResult = 2
)

var SC_PlayerLogin_LoginResult_name = map[int32]string{
	0: "SUCCESS",
	1: "ACCOUNT_ERROR",
	2: "PWD_ERROR",
}
var SC_PlayerLogin_LoginResult_value = map[string]int32{
	"SUCCESS":       0,
	"ACCOUNT_ERROR": 1,
	"PWD_ERROR":     2,
}

func (x SC_PlayerLogin_LoginResult) String() string {
	return proto.EnumName(SC_PlayerLogin_LoginResult_name, int32(x))
}
func (SC_PlayerLogin_LoginResult) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

type SC_PlayerRegister_RegisterResult int32

const (
	SC_PlayerRegister_SUCCESS       SC_PlayerRegister_RegisterResult = 0
	SC_PlayerRegister_ACCOUNT_ERROR SC_PlayerRegister_RegisterResult = 1
)

var SC_PlayerRegister_RegisterResult_name = map[int32]string{
	0: "SUCCESS",
	1: "ACCOUNT_ERROR",
}
var SC_PlayerRegister_RegisterResult_value = map[string]int32{
	"SUCCESS":       0,
	"ACCOUNT_ERROR": 1,
}

func (x SC_PlayerRegister_RegisterResult) String() string {
	return proto.EnumName(SC_PlayerRegister_RegisterResult_name, int32(x))
}
func (SC_PlayerRegister_RegisterResult) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

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
	LoginResult SC_PlayerLogin_LoginResult `protobuf:"varint,1,opt,name=loginResult,enum=PlayerLogin.SC_PlayerLogin_LoginResult" json:"loginResult,omitempty"`
}

func (m *SC_PlayerLogin) Reset()                    { *m = SC_PlayerLogin{} }
func (m *SC_PlayerLogin) String() string            { return proto.CompactTextString(m) }
func (*SC_PlayerLogin) ProtoMessage()               {}
func (*SC_PlayerLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SC_PlayerLogin) GetLoginResult() SC_PlayerLogin_LoginResult {
	if m != nil {
		return m.LoginResult
	}
	return SC_PlayerLogin_SUCCESS
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
	RegisterResult SC_PlayerRegister_RegisterResult `protobuf:"varint,1,opt,name=registerResult,enum=PlayerLogin.SC_PlayerRegister_RegisterResult" json:"registerResult,omitempty"`
}

func (m *SC_PlayerRegister) Reset()                    { *m = SC_PlayerRegister{} }
func (m *SC_PlayerRegister) String() string            { return proto.CompactTextString(m) }
func (*SC_PlayerRegister) ProtoMessage()               {}
func (*SC_PlayerRegister) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SC_PlayerRegister) GetRegisterResult() SC_PlayerRegister_RegisterResult {
	if m != nil {
		return m.RegisterResult
	}
	return SC_PlayerRegister_SUCCESS
}

func init() {
	proto.RegisterType((*CS_PlayerLogin)(nil), "PlayerLogin.CS_PlayerLogin")
	proto.RegisterType((*SC_PlayerLogin)(nil), "PlayerLogin.SC_PlayerLogin")
	proto.RegisterType((*CS_PlayerRegister)(nil), "PlayerLogin.CS_PlayerRegister")
	proto.RegisterType((*SC_PlayerRegister)(nil), "PlayerLogin.SC_PlayerRegister")
	proto.RegisterEnum("PlayerLogin.SC_PlayerLogin_LoginResult", SC_PlayerLogin_LoginResult_name, SC_PlayerLogin_LoginResult_value)
	proto.RegisterEnum("PlayerLogin.SC_PlayerRegister_RegisterResult", SC_PlayerRegister_RegisterResult_name, SC_PlayerRegister_RegisterResult_value)
}

func init() { proto.RegisterFile("LoginInfo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xf7, 0xc9, 0x4f, 0xcf,
	0xcc, 0xf3, 0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x0e, 0xc8, 0x49,
	0xac, 0x4c, 0x2d, 0x02, 0x0b, 0x2b, 0xd9, 0x70, 0xf1, 0x39, 0x07, 0xc7, 0x23, 0x89, 0x08, 0x49,
	0x70, 0xb1, 0x27, 0x26, 0x27, 0xe7, 0x97, 0xe6, 0x95, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0xc1, 0xb8, 0x42, 0x02, 0x5c, 0xcc, 0x05, 0xe5, 0x29, 0x12, 0x4c, 0x60, 0x51, 0x10, 0x53, 0x69,
	0x26, 0x23, 0x17, 0x5f, 0xb0, 0x33, 0x8a, 0x76, 0x4f, 0x2e, 0xee, 0x1c, 0x10, 0x23, 0x28, 0xb5,
	0xb8, 0x34, 0x07, 0x62, 0x04, 0x9f, 0x91, 0xba, 0x1e, 0x92, 0x12, 0x3d, 0x54, 0x1d, 0x7a, 0x3e,
	0x08, 0xe5, 0x41, 0xc8, 0x7a, 0x95, 0x6c, 0xb8, 0xb8, 0x91, 0xe4, 0x84, 0xb8, 0xb9, 0xd8, 0x83,
	0x43, 0x9d, 0x9d, 0x5d, 0x83, 0x83, 0x05, 0x18, 0x84, 0x04, 0xb9, 0x78, 0x1d, 0x9d, 0x9d, 0xfd,
	0x43, 0xfd, 0x42, 0xe2, 0x5d, 0x83, 0x82, 0xfc, 0x83, 0x04, 0x18, 0x85, 0x78, 0xb9, 0x38, 0x03,
	0xc2, 0x5d, 0xa0, 0x5c, 0x26, 0x25, 0x7b, 0x2e, 0x41, 0xb8, 0xcf, 0x82, 0x52, 0xd3, 0x33, 0x8b,
	0x4b, 0x52, 0x8b, 0x48, 0xf2, 0xdc, 0x1c, 0x46, 0x2e, 0x41, 0xb8, 0x53, 0xe1, 0x26, 0x84, 0x72,
	0xf1, 0x15, 0x41, 0xd9, 0x28, 0x5e, 0xd4, 0xc5, 0xee, 0x45, 0x98, 0x3e, 0xbd, 0x20, 0x14, 0x4d,
	0x41, 0x68, 0x86, 0x28, 0x19, 0x70, 0xf1, 0xa1, 0xaa, 0x20, 0xe4, 0xdd, 0x24, 0x36, 0x70, 0x6c,
	0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x37, 0x51, 0x5d, 0xe0, 0x01, 0x00, 0x00,
}
