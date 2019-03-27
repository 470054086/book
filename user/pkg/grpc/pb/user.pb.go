// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type UserInfo struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Phone                int64    `protobuf:"varint,2,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	Age                  int64    `protobuf:"varint,4,opt,name=Age,proto3" json:"Age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfo) GetPhone() int64 {
	if m != nil {
		return m.Phone
	}
	return 0
}

func (m *UserInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserInfo) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

type RegisterRequest struct {
	User                 *UserInfo `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetUser() *UserInfo {
	if m != nil {
		return m.User
	}
	return nil
}

type RegisterReply struct {
	User                 *UserInfo `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RegisterReply) Reset()         { *m = RegisterReply{} }
func (m *RegisterReply) String() string { return proto.CompactTextString(m) }
func (*RegisterReply) ProtoMessage()    {}
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *RegisterReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReply.Unmarshal(m, b)
}
func (m *RegisterReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReply.Marshal(b, m, deterministic)
}
func (m *RegisterReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReply.Merge(m, src)
}
func (m *RegisterReply) XXX_Size() int {
	return xxx_messageInfo_RegisterReply.Size(m)
}
func (m *RegisterReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReply.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReply proto.InternalMessageInfo

func (m *RegisterReply) GetUser() *UserInfo {
	if m != nil {
		return m.User
	}
	return nil
}

type UserInfoByIdRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoByIdRequest) Reset()         { *m = UserInfoByIdRequest{} }
func (m *UserInfoByIdRequest) String() string { return proto.CompactTextString(m) }
func (*UserInfoByIdRequest) ProtoMessage()    {}
func (*UserInfoByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *UserInfoByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoByIdRequest.Unmarshal(m, b)
}
func (m *UserInfoByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoByIdRequest.Marshal(b, m, deterministic)
}
func (m *UserInfoByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoByIdRequest.Merge(m, src)
}
func (m *UserInfoByIdRequest) XXX_Size() int {
	return xxx_messageInfo_UserInfoByIdRequest.Size(m)
}
func (m *UserInfoByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoByIdRequest proto.InternalMessageInfo

func (m *UserInfoByIdRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UserInfoByIdReply struct {
	User                 *UserInfo `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UserInfoByIdReply) Reset()         { *m = UserInfoByIdReply{} }
func (m *UserInfoByIdReply) String() string { return proto.CompactTextString(m) }
func (*UserInfoByIdReply) ProtoMessage()    {}
func (*UserInfoByIdReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *UserInfoByIdReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoByIdReply.Unmarshal(m, b)
}
func (m *UserInfoByIdReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoByIdReply.Marshal(b, m, deterministic)
}
func (m *UserInfoByIdReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoByIdReply.Merge(m, src)
}
func (m *UserInfoByIdReply) XXX_Size() int {
	return xxx_messageInfo_UserInfoByIdReply.Size(m)
}
func (m *UserInfoByIdReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoByIdReply.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoByIdReply proto.InternalMessageInfo

func (m *UserInfoByIdReply) GetUser() *UserInfo {
	if m != nil {
		return m.User
	}
	return nil
}

type UserInfoByPhoneRequest struct {
	Phone                int64    `protobuf:"varint,1,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoByPhoneRequest) Reset()         { *m = UserInfoByPhoneRequest{} }
func (m *UserInfoByPhoneRequest) String() string { return proto.CompactTextString(m) }
func (*UserInfoByPhoneRequest) ProtoMessage()    {}
func (*UserInfoByPhoneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *UserInfoByPhoneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoByPhoneRequest.Unmarshal(m, b)
}
func (m *UserInfoByPhoneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoByPhoneRequest.Marshal(b, m, deterministic)
}
func (m *UserInfoByPhoneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoByPhoneRequest.Merge(m, src)
}
func (m *UserInfoByPhoneRequest) XXX_Size() int {
	return xxx_messageInfo_UserInfoByPhoneRequest.Size(m)
}
func (m *UserInfoByPhoneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoByPhoneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoByPhoneRequest proto.InternalMessageInfo

func (m *UserInfoByPhoneRequest) GetPhone() int64 {
	if m != nil {
		return m.Phone
	}
	return 0
}

type UserInfoByPhoneReply struct {
	User                 *UserInfo `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UserInfoByPhoneReply) Reset()         { *m = UserInfoByPhoneReply{} }
func (m *UserInfoByPhoneReply) String() string { return proto.CompactTextString(m) }
func (*UserInfoByPhoneReply) ProtoMessage()    {}
func (*UserInfoByPhoneReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *UserInfoByPhoneReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoByPhoneReply.Unmarshal(m, b)
}
func (m *UserInfoByPhoneReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoByPhoneReply.Marshal(b, m, deterministic)
}
func (m *UserInfoByPhoneReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoByPhoneReply.Merge(m, src)
}
func (m *UserInfoByPhoneReply) XXX_Size() int {
	return xxx_messageInfo_UserInfoByPhoneReply.Size(m)
}
func (m *UserInfoByPhoneReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoByPhoneReply.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoByPhoneReply proto.InternalMessageInfo

func (m *UserInfoByPhoneReply) GetUser() *UserInfo {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*UserInfo)(nil), "pb.UserInfo")
	proto.RegisterType((*RegisterRequest)(nil), "pb.RegisterRequest")
	proto.RegisterType((*RegisterReply)(nil), "pb.RegisterReply")
	proto.RegisterType((*UserInfoByIdRequest)(nil), "pb.UserInfoByIdRequest")
	proto.RegisterType((*UserInfoByIdReply)(nil), "pb.UserInfoByIdReply")
	proto.RegisterType((*UserInfoByPhoneRequest)(nil), "pb.UserInfoByPhoneRequest")
	proto.RegisterType((*UserInfoByPhoneReply)(nil), "pb.UserInfoByPhoneReply")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x25, 0x1f, 0x95, 0xf8, 0xac, 0xad, 0x9d, 0x46, 0x0d, 0x39, 0x85, 0x80, 0xd0, 0x53, 0xc0,
	0x14, 0xc1, 0x93, 0xa0, 0x17, 0xc9, 0xad, 0x04, 0xbc, 0x0a, 0x86, 0xac, 0x35, 0x50, 0x9a, 0x35,
	0x9b, 0x22, 0xf9, 0x7d, 0xfe, 0x31, 0xd9, 0x0d, 0xdb, 0x26, 0xad, 0x87, 0xdc, 0x76, 0xde, 0xbc,
	0x37, 0x6f, 0xe6, 0xb1, 0xc0, 0x4e, 0xb0, 0x2a, 0xe2, 0x55, 0x59, 0x97, 0x64, 0xf2, 0x2c, 0x7c,
	0x87, 0xf3, 0x26, 0x58, 0x95, 0x6c, 0x3f, 0x4b, 0x9a, 0xc0, 0x4c, 0x72, 0xcf, 0x08, 0x8c, 0x85,
	0x95, 0x9a, 0x49, 0x4e, 0x2e, 0x46, 0xab, 0xaf, 0x72, 0xcb, 0x3c, 0x53, 0x41, 0x6d, 0x41, 0x3e,
	0x9c, 0xd5, 0x87, 0x10, 0x3f, 0x65, 0x95, 0x7b, 0x56, 0x60, 0x2c, 0xce, 0xd3, 0x7d, 0x4d, 0x57,
	0xb0, 0x9e, 0xd7, 0xcc, 0xb3, 0x15, 0x5f, 0x3e, 0xc3, 0x25, 0xa6, 0x29, 0x5b, 0x17, 0xa2, 0x66,
	0x55, 0xca, 0xbe, 0x77, 0x4c, 0xd4, 0x14, 0xc0, 0x96, 0x4b, 0x28, 0xa3, 0x8b, 0x78, 0x1c, 0xf1,
	0x2c, 0xd2, 0x2b, 0xa4, 0xaa, 0x13, 0xde, 0xe3, 0xf2, 0x20, 0xe2, 0x9b, 0x66, 0x80, 0xe4, 0x0e,
	0x73, 0x8d, 0xbc, 0x34, 0x49, 0xae, 0xbd, 0x26, 0x30, 0x8b, 0xfd, 0x49, 0x45, 0x1e, 0x3e, 0x60,
	0xd6, 0xa7, 0x0d, 0x9b, 0x1e, 0xe1, 0xe6, 0x20, 0x53, 0x31, 0x68, 0x03, 0x17, 0x23, 0xae, 0x32,
	0x6a, 0x3d, 0xda, 0x22, 0x7c, 0x84, 0x7b, 0xc2, 0x1f, 0xe4, 0x14, 0xff, 0x1a, 0xb0, 0x25, 0x44,
	0x31, 0x1c, 0x9d, 0x01, 0xcd, 0x25, 0xf1, 0x28, 0x46, 0x7f, 0xd6, 0x07, 0xe5, 0xf8, 0x27, 0x8c,
	0xbb, 0xd7, 0xd1, 0x6d, 0xd7, 0xa0, 0x13, 0x8b, 0x7f, 0x7d, 0xda, 0x90, 0xfa, 0x57, 0x4c, 0x8f,
	0xd6, 0x26, 0xbf, 0xcf, 0xec, 0xde, 0xee, 0x7b, 0xff, 0xf6, 0xf8, 0xa6, 0xc9, 0xce, 0xd4, 0x07,
	0x5b, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x25, 0xf7, 0x21, 0x69, 0x6e, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
	UserInfoById(ctx context.Context, in *UserInfoByIdRequest, opts ...grpc.CallOption) (*UserInfoByIdReply, error)
	UserInfoByPhone(ctx context.Context, in *UserInfoByPhoneRequest, opts ...grpc.CallOption) (*UserInfoByPhoneReply, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, "/pb.User/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserInfoById(ctx context.Context, in *UserInfoByIdRequest, opts ...grpc.CallOption) (*UserInfoByIdReply, error) {
	out := new(UserInfoByIdReply)
	err := c.cc.Invoke(ctx, "/pb.User/UserInfoById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserInfoByPhone(ctx context.Context, in *UserInfoByPhoneRequest, opts ...grpc.CallOption) (*UserInfoByPhoneReply, error) {
	out := new(UserInfoByPhoneReply)
	err := c.cc.Invoke(ctx, "/pb.User/UserInfoByPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	UserInfoById(context.Context, *UserInfoByIdRequest) (*UserInfoByIdReply, error)
	UserInfoByPhone(context.Context, *UserInfoByPhoneRequest) (*UserInfoByPhoneReply, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserInfoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserInfoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/UserInfoById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserInfoById(ctx, req.(*UserInfoByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserInfoByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoByPhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserInfoByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/UserInfoByPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserInfoByPhone(ctx, req.(*UserInfoByPhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _User_Register_Handler,
		},
		{
			MethodName: "UserInfoById",
			Handler:    _User_UserInfoById_Handler,
		},
		{
			MethodName: "UserInfoByPhone",
			Handler:    _User_UserInfoByPhone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}