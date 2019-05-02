// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package go_micro_srv_user

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

// 用户信息
type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Company              string   `protobuf:"bytes,3,opt,name=company,proto3" json:"company,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Tel                  string   `protobuf:"bytes,5,opt,name=tel,proto3" json:"tel,omitempty"`
	Password             string   `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetTel() string {
	if m != nil {
		return m.Tel
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
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
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
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

type ToRevokeToken struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToRevokeToken) Reset()         { *m = ToRevokeToken{} }
func (m *ToRevokeToken) String() string { return proto.CompactTextString(m) }
func (*ToRevokeToken) ProtoMessage()    {}
func (*ToRevokeToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *ToRevokeToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToRevokeToken.Unmarshal(m, b)
}
func (m *ToRevokeToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToRevokeToken.Marshal(b, m, deterministic)
}
func (m *ToRevokeToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToRevokeToken.Merge(m, src)
}
func (m *ToRevokeToken) XXX_Size() int {
	return xxx_messageInfo_ToRevokeToken.Size(m)
}
func (m *ToRevokeToken) XXX_DiscardUnknown() {
	xxx_messageInfo_ToRevokeToken.DiscardUnknown(m)
}

var xxx_messageInfo_ToRevokeToken proto.InternalMessageInfo

func (m *ToRevokeToken) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ToRevokeToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Users struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Users) Reset()         { *m = Users{} }
func (m *Users) String() string { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()    {}
func (*Users) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{4}
}

func (m *Users) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Users.Unmarshal(m, b)
}
func (m *Users) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Users.Marshal(b, m, deterministic)
}
func (m *Users) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Users.Merge(m, src)
}
func (m *Users) XXX_Size() int {
	return xxx_messageInfo_Users.Size(m)
}
func (m *Users) XXX_DiscardUnknown() {
	xxx_messageInfo_Users.DiscardUnknown(m)
}

var xxx_messageInfo_Users proto.InternalMessageInfo

func (m *Users) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{5}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "go.micro.srv.user.User")
	proto.RegisterType((*Request)(nil), "go.micro.srv.user.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.user.Response")
	proto.RegisterType((*ToRevokeToken)(nil), "go.micro.srv.user.ToRevokeToken")
	proto.RegisterType((*Users)(nil), "go.micro.srv.user.Users")
	proto.RegisterType((*Token)(nil), "go.micro.srv.user.Token")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x5d, 0x4b, 0xe3, 0x40,
	0x14, 0x6d, 0x9a, 0xa4, 0x1f, 0xb7, 0xec, 0xb2, 0x7b, 0xe9, 0xb2, 0x43, 0x96, 0x85, 0x32, 0x4f,
	0x7d, 0x31, 0x42, 0x45, 0x41, 0x7c, 0xb1, 0x58, 0x28, 0xbe, 0x49, 0x6c, 0x7f, 0x40, 0x6c, 0xae,
	0x35, 0x34, 0xc9, 0xc4, 0x99, 0x69, 0xc5, 0x7f, 0xe0, 0x0f, 0xf4, 0x07, 0xc9, 0x4c, 0xb4, 0x08,
	0x6d, 0xfa, 0xa0, 0x2f, 0x61, 0xce, 0xb9, 0x77, 0xce, 0xb9, 0x27, 0x97, 0x81, 0x3f, 0xa5, 0x14,
	0x5a, 0x1c, 0xaf, 0x15, 0x49, 0xfb, 0x09, 0x2d, 0xc6, 0xdf, 0x4b, 0x11, 0xe6, 0xe9, 0x42, 0x8a,
	0x50, 0xc9, 0x4d, 0x68, 0x0a, 0xfc, 0xc5, 0x01, 0x6f, 0xae, 0x48, 0xe2, 0x4f, 0x68, 0xa6, 0x09,
	0x73, 0x06, 0xce, 0xb0, 0x1b, 0x35, 0xd3, 0x04, 0x11, 0xbc, 0x22, 0xce, 0x89, 0x35, 0x2d, 0x63,
	0xcf, 0xc8, 0xa0, 0xbd, 0x10, 0x79, 0x19, 0x17, 0xcf, 0xcc, 0xb5, 0xf4, 0x07, 0xc4, 0x3e, 0xf8,
	0x94, 0xc7, 0x69, 0xc6, 0x3c, 0xcb, 0x57, 0x00, 0x7f, 0x81, 0xab, 0x29, 0x63, 0xbe, 0xe5, 0xcc,
	0x11, 0x03, 0xe8, 0x94, 0xb1, 0x52, 0x4f, 0x42, 0x26, 0xac, 0x65, 0xe9, 0x2d, 0xe6, 0x5d, 0x68,
	0x47, 0xf4, 0xb8, 0x26, 0xa5, 0x39, 0x40, 0x27, 0x22, 0x55, 0x8a, 0x42, 0x11, 0x3f, 0x85, 0x1f,
	0x33, 0x11, 0xd1, 0x46, 0xac, 0x68, 0x26, 0x56, 0x54, 0xec, 0x4c, 0xda, 0x07, 0x5f, 0x9b, 0xc2,
	0xfb, 0xa8, 0x15, 0xe0, 0x67, 0xe0, 0x9b, 0x5c, 0x0a, 0x8f, 0xc0, 0x37, 0x49, 0x15, 0x73, 0x06,
	0xee, 0xb0, 0x37, 0xfa, 0x1b, 0xee, 0xfc, 0x84, 0xd0, 0x34, 0x46, 0x55, 0x17, 0xff, 0x0f, 0x7e,
	0x65, 0xb3, 0x95, 0x75, 0x3e, 0xc9, 0x8e, 0x5e, 0x5d, 0xe8, 0x99, 0xf6, 0x5b, 0x92, 0x9b, 0x74,
	0x41, 0x78, 0x09, 0xad, 0x2b, 0x49, 0xb1, 0x26, 0xac, 0x13, 0x0e, 0xfe, 0xed, 0x29, 0x6c, 0xd3,
	0x35, 0x8c, 0xc2, 0x84, 0x32, 0xfa, 0x86, 0xc2, 0x39, 0xb8, 0x53, 0xd2, 0xf5, 0xd7, 0xeb, 0x0a,
	0x95, 0xf9, 0x94, 0xf4, 0x38, 0xcb, 0x30, 0xd8, 0xeb, 0x61, 0xd7, 0x11, 0xb0, 0x1a, 0x01, 0xc5,
	0x1b, 0x38, 0x01, 0x98, 0x97, 0x49, 0xac, 0xe9, 0xba, 0xb8, 0x17, 0x5f, 0x8e, 0x70, 0x01, 0xde,
	0x78, 0xad, 0x1f, 0xea, 0xef, 0xef, 0x1b, 0xc1, 0xee, 0x89, 0x37, 0x70, 0x0c, 0xde, 0x4d, 0x5a,
	0x2c, 0x0f, 0x46, 0x38, 0xec, 0x7f, 0xd7, 0xb2, 0x0f, 0xe4, 0xe4, 0x2d, 0x00, 0x00, 0xff, 0xff,
	0xbf, 0xc6, 0x85, 0x0d, 0x39, 0x03, 0x00, 0x00,
}
