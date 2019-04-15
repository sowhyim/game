// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pbServer/server.proto

package server

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ConnRequest struct {
	Login                string   `protobuf:"bytes,1,opt,name=Login,proto3" json:"Login,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	RmoteAddr            string   `protobuf:"bytes,3,opt,name=RmoteAddr,proto3" json:"RmoteAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnRequest) Reset()         { *m = ConnRequest{} }
func (m *ConnRequest) String() string { return proto.CompactTextString(m) }
func (*ConnRequest) ProtoMessage()    {}
func (*ConnRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_226f711b2eb24214, []int{0}
}

func (m *ConnRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnRequest.Unmarshal(m, b)
}
func (m *ConnRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnRequest.Marshal(b, m, deterministic)
}
func (m *ConnRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnRequest.Merge(m, src)
}
func (m *ConnRequest) XXX_Size() int {
	return xxx_messageInfo_ConnRequest.Size(m)
}
func (m *ConnRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnRequest proto.InternalMessageInfo

func (m *ConnRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *ConnRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ConnRequest) GetRmoteAddr() string {
	if m != nil {
		return m.RmoteAddr
	}
	return ""
}

type ConnReply struct {
	Status               int32    `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnReply) Reset()         { *m = ConnReply{} }
func (m *ConnReply) String() string { return proto.CompactTextString(m) }
func (*ConnReply) ProtoMessage()    {}
func (*ConnReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_226f711b2eb24214, []int{1}
}

func (m *ConnReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnReply.Unmarshal(m, b)
}
func (m *ConnReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnReply.Marshal(b, m, deterministic)
}
func (m *ConnReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnReply.Merge(m, src)
}
func (m *ConnReply) XXX_Size() int {
	return xxx_messageInfo_ConnReply.Size(m)
}
func (m *ConnReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnReply.DiscardUnknown(m)
}

var xxx_messageInfo_ConnReply proto.InternalMessageInfo

func (m *ConnReply) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*ConnRequest)(nil), "ConnRequest")
	proto.RegisterType((*ConnReply)(nil), "ConnReply")
}

func init() { proto.RegisterFile("pbServer/server.proto", fileDescriptor_226f711b2eb24214) }

var fileDescriptor_226f711b2eb24214 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x48, 0x0a, 0x4e,
	0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x2f, 0x06, 0x53, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x4a, 0xa1,
	0x5c, 0xdc, 0xce, 0xf9, 0x79, 0x79, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x22, 0x5c,
	0xac, 0x3e, 0xf9, 0xe9, 0x99, 0x79, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x90,
	0x10, 0x17, 0x8b, 0x5f, 0x62, 0x6e, 0xaa, 0x04, 0x13, 0x58, 0x10, 0xcc, 0x16, 0x92, 0xe1, 0xe2,
	0x0c, 0xca, 0xcd, 0x2f, 0x49, 0x75, 0x4c, 0x49, 0x29, 0x92, 0x60, 0x06, 0x4b, 0x20, 0x04, 0x94,
	0x94, 0xb9, 0x38, 0x21, 0xc6, 0x16, 0xe4, 0x54, 0x0a, 0x89, 0x71, 0xb1, 0x05, 0x97, 0x24, 0x96,
	0x94, 0x16, 0x83, 0x4d, 0x65, 0x0d, 0x82, 0xf2, 0x8c, 0x4c, 0xb8, 0x38, 0x21, 0x4e, 0x0a, 0x2e,
	0x2a, 0x13, 0x52, 0xe7, 0xe2, 0x06, 0x5b, 0x06, 0x11, 0x11, 0xe2, 0xd1, 0x43, 0x72, 0x96, 0x14,
	0x97, 0x1e, 0xdc, 0xb4, 0x24, 0x36, 0xb0, 0xc3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0e,
	0x3a, 0x14, 0x3b, 0xd1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServerSrvClient is the client API for ServerSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerSrvClient interface {
	LoginServer(ctx context.Context, in *ConnRequest, opts ...grpc.CallOption) (*ConnReply, error)
}

type serverSrvClient struct {
	cc *grpc.ClientConn
}

func NewServerSrvClient(cc *grpc.ClientConn) ServerSrvClient {
	return &serverSrvClient{cc}
}

func (c *serverSrvClient) LoginServer(ctx context.Context, in *ConnRequest, opts ...grpc.CallOption) (*ConnReply, error) {
	out := new(ConnReply)
	err := c.cc.Invoke(ctx, "/ServerSrv/LoginServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerSrvServer is the server API for ServerSrv service.
type ServerSrvServer interface {
	LoginServer(context.Context, *ConnRequest) (*ConnReply, error)
}

// UnimplementedServerSrvServer can be embedded to have forward compatible implementations.
type UnimplementedServerSrvServer struct {
}

func (*UnimplementedServerSrvServer) LoginServer(ctx context.Context, req *ConnRequest) (*ConnReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginServer not implemented")
}

func RegisterServerSrvServer(s *grpc.Server, srv ServerSrvServer) {
	s.RegisterService(&_ServerSrv_serviceDesc, srv)
}

func _ServerSrv_LoginServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerSrvServer).LoginServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ServerSrv/LoginServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerSrvServer).LoginServer(ctx, req.(*ConnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServerSrv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ServerSrv",
	HandlerType: (*ServerSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginServer",
			Handler:    _ServerSrv_LoginServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pbServer/server.proto",
}
