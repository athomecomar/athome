// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mailer.proto

package pbmailer

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ForgotPasswordRequest struct {
	Email                string           `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	TokenizedUsers       []*TokenizedUser `protobuf:"bytes,2,rep,name=tokenized_users,json=tokenizedUsers,proto3" json:"tokenized_users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ForgotPasswordRequest) Reset()         { *m = ForgotPasswordRequest{} }
func (m *ForgotPasswordRequest) String() string { return proto.CompactTextString(m) }
func (*ForgotPasswordRequest) ProtoMessage()    {}
func (*ForgotPasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5eced95006893a4c, []int{0}
}

func (m *ForgotPasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForgotPasswordRequest.Unmarshal(m, b)
}
func (m *ForgotPasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForgotPasswordRequest.Marshal(b, m, deterministic)
}
func (m *ForgotPasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForgotPasswordRequest.Merge(m, src)
}
func (m *ForgotPasswordRequest) XXX_Size() int {
	return xxx_messageInfo_ForgotPasswordRequest.Size(m)
}
func (m *ForgotPasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ForgotPasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ForgotPasswordRequest proto.InternalMessageInfo

func (m *ForgotPasswordRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ForgotPasswordRequest) GetTokenizedUsers() []*TokenizedUser {
	if m != nil {
		return m.TokenizedUsers
	}
	return nil
}

type TokenizedUser struct {
	Role                 string   `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenizedUser) Reset()         { *m = TokenizedUser{} }
func (m *TokenizedUser) String() string { return proto.CompactTextString(m) }
func (*TokenizedUser) ProtoMessage()    {}
func (*TokenizedUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_5eced95006893a4c, []int{1}
}

func (m *TokenizedUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenizedUser.Unmarshal(m, b)
}
func (m *TokenizedUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenizedUser.Marshal(b, m, deterministic)
}
func (m *TokenizedUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenizedUser.Merge(m, src)
}
func (m *TokenizedUser) XXX_Size() int {
	return xxx_messageInfo_TokenizedUser.Size(m)
}
func (m *TokenizedUser) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenizedUser.DiscardUnknown(m)
}

var xxx_messageInfo_TokenizedUser proto.InternalMessageInfo

func (m *TokenizedUser) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *TokenizedUser) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*ForgotPasswordRequest)(nil), "mailer.ForgotPasswordRequest")
	proto.RegisterType((*TokenizedUser)(nil), "mailer.TokenizedUser")
}

func init() {
	proto.RegisterFile("mailer.proto", fileDescriptor_5eced95006893a4c)
}

var fileDescriptor_5eced95006893a4c = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcd, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x4d, 0x6c, 0xa3, 0xae, 0xb5, 0xca, 0x62, 0x35, 0x44, 0x0a, 0x25, 0x07, 0x09, 0x88,
	0x1b, 0xa8, 0x77, 0x85, 0x80, 0x82, 0x07, 0x41, 0x56, 0xbd, 0x78, 0xd1, 0x7c, 0x4c, 0x63, 0x30,
	0xc9, 0xc6, 0xd9, 0x4d, 0x44, 0xff, 0xf4, 0x9e, 0x24, 0xd9, 0x06, 0x2c, 0xf4, 0x36, 0x6f, 0xdf,
	0x7c, 0xbc, 0xfd, 0x91, 0x51, 0x11, 0x66, 0x39, 0x20, 0xab, 0x50, 0x28, 0x41, 0x2d, 0xad, 0x9c,
	0xd3, 0x26, 0xcc, 0xb3, 0x24, 0x54, 0xe0, 0xf7, 0x85, 0x6e, 0x70, 0xce, 0x52, 0x21, 0xd2, 0x1c,
	0xfc, 0x4e, 0x45, 0xf5, 0xc2, 0x87, 0xa2, 0x52, 0x3f, 0xda, 0x74, 0x1b, 0x32, 0xb9, 0x13, 0x98,
	0x0a, 0xf5, 0x18, 0x4a, 0xf9, 0x2d, 0x30, 0xe1, 0xf0, 0x55, 0x83, 0x54, 0x74, 0x4a, 0x86, 0xd0,
	0x6e, 0xb6, 0x8d, 0x99, 0xe1, 0xed, 0x05, 0x3b, 0xcb, 0x60, 0x80, 0xe6, 0xbb, 0xc1, 0xf5, 0x2b,
	0xbd, 0x26, 0x87, 0x4a, 0x7c, 0x42, 0x99, 0xfd, 0x42, 0xf2, 0x56, 0x4b, 0x40, 0x69, 0x9b, 0xb3,
	0x6d, 0x6f, 0x7f, 0x3e, 0x61, 0xab, 0x74, 0xcf, 0xbd, 0xfd, 0x22, 0x01, 0xf9, 0x58, 0xfd, 0x97,
	0xd2, 0x5d, 0x90, 0x83, 0xb5, 0x06, 0x7a, 0x43, 0x06, 0x28, 0x72, 0x58, 0x9d, 0xbb, 0x58, 0x06,
	0x1e, 0x9e, 0xf3, 0x23, 0x09, 0xd8, 0x64, 0x31, 0x5c, 0x56, 0x28, 0x9a, 0x2c, 0x01, 0xe4, 0xbb,
	0xb1, 0x28, 0x65, 0x5d, 0xb4, 0x55, 0x01, 0x18, 0x7f, 0x84, 0xa5, 0xe2, 0xdd, 0x20, 0x3d, 0x26,
	0xc3, 0xee, 0x86, 0x6d, 0xb6, 0x1b, 0xb8, 0x16, 0xf3, 0x27, 0x62, 0x3d, 0x74, 0x79, 0xe8, 0x3d,
	0x19, 0xaf, 0xff, 0x94, 0x4e, 0xfb, 0xa8, 0x1b, 0x09, 0x38, 0x27, 0x4c, 0x83, 0x63, 0x3d, 0x38,
	0x76, 0xdb, 0x82, 0x73, 0xb7, 0x82, 0xd1, 0x2b, 0x61, 0x7e, 0x15, 0xe9, 0xe9, 0xc8, 0xea, 0xfc,
	0xab, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1b, 0xdb, 0x26, 0x6f, 0x97, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MailerClient is the client API for Mailer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MailerClient interface {
	ForgotPassword(ctx context.Context, in *ForgotPasswordRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type mailerClient struct {
	cc grpc.ClientConnInterface
}

func NewMailerClient(cc grpc.ClientConnInterface) MailerClient {
	return &mailerClient{cc}
}

func (c *mailerClient) ForgotPassword(ctx context.Context, in *ForgotPasswordRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/mailer.Mailer/ForgotPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MailerServer is the server API for Mailer service.
type MailerServer interface {
	ForgotPassword(context.Context, *ForgotPasswordRequest) (*empty.Empty, error)
}

// UnimplementedMailerServer can be embedded to have forward compatible implementations.
type UnimplementedMailerServer struct {
}

func (*UnimplementedMailerServer) ForgotPassword(ctx context.Context, req *ForgotPasswordRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForgotPassword not implemented")
}

func RegisterMailerServer(s *grpc.Server, srv MailerServer) {
	s.RegisterService(&_Mailer_serviceDesc, srv)
}

func _Mailer_ForgotPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgotPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailerServer).ForgotPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mailer.Mailer/ForgotPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailerServer).ForgotPassword(ctx, req.(*ForgotPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mailer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mailer.Mailer",
	HandlerType: (*MailerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ForgotPassword",
			Handler:    _Mailer_ForgotPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mailer.proto",
}