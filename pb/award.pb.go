// Code generated by protoc-gen-go. DO NOT EDIT.
// source: award.proto

package pb

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

// The request message containing the user's name.
type DrawRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DrawRequest) Reset()         { *m = DrawRequest{} }
func (m *DrawRequest) String() string { return proto.CompactTextString(m) }
func (*DrawRequest) ProtoMessage()    {}
func (*DrawRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3994c66428a105c7, []int{0}
}

func (m *DrawRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DrawRequest.Unmarshal(m, b)
}
func (m *DrawRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DrawRequest.Marshal(b, m, deterministic)
}
func (m *DrawRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DrawRequest.Merge(m, src)
}
func (m *DrawRequest) XXX_Size() int {
	return xxx_messageInfo_DrawRequest.Size(m)
}
func (m *DrawRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DrawRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DrawRequest proto.InternalMessageInfo

func (m *DrawRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type DrawReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DrawReply) Reset()         { *m = DrawReply{} }
func (m *DrawReply) String() string { return proto.CompactTextString(m) }
func (*DrawReply) ProtoMessage()    {}
func (*DrawReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_3994c66428a105c7, []int{1}
}

func (m *DrawReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DrawReply.Unmarshal(m, b)
}
func (m *DrawReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DrawReply.Marshal(b, m, deterministic)
}
func (m *DrawReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DrawReply.Merge(m, src)
}
func (m *DrawReply) XXX_Size() int {
	return xxx_messageInfo_DrawReply.Size(m)
}
func (m *DrawReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DrawReply.DiscardUnknown(m)
}

var xxx_messageInfo_DrawReply proto.InternalMessageInfo

func (m *DrawReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*DrawRequest)(nil), "pb.DrawRequest")
	proto.RegisterType((*DrawReply)(nil), "pb.DrawReply")
}

func init() { proto.RegisterFile("award.proto", fileDescriptor_3994c66428a105c7) }

var fileDescriptor_3994c66428a105c7 = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2c, 0x4f, 0x2c,
	0x4a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x52, 0xe4, 0xe2, 0x76,
	0x29, 0x4a, 0x2c, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b,
	0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x54, 0xb9, 0x38, 0x21,
	0x4a, 0x0a, 0x72, 0x2a, 0x85, 0x24, 0xb8, 0xd8, 0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x61, 0x6a,
	0x60, 0x5c, 0x23, 0x43, 0x2e, 0x56, 0x47, 0x90, 0xe1, 0x42, 0x1a, 0x5c, 0x2c, 0x20, 0xf5, 0x42,
	0xfc, 0x7a, 0x05, 0x49, 0x7a, 0x48, 0x86, 0x4b, 0xf1, 0x22, 0x04, 0x0a, 0x72, 0x2a, 0x95, 0x18,
	0x92, 0xd8, 0xc0, 0xee, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x5f, 0xac, 0xfa, 0x96,
	0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AwardClient is the client API for Award service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AwardClient interface {
	// Sends a greeting
	Draw(ctx context.Context, in *DrawRequest, opts ...grpc.CallOption) (*DrawReply, error)
}

type awardClient struct {
	cc *grpc.ClientConn
}

func NewAwardClient(cc *grpc.ClientConn) AwardClient {
	return &awardClient{cc}
}

func (c *awardClient) Draw(ctx context.Context, in *DrawRequest, opts ...grpc.CallOption) (*DrawReply, error) {
	out := new(DrawReply)
	err := c.cc.Invoke(ctx, "/pb.Award/Draw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AwardServer is the server API for Award service.
type AwardServer interface {
	// Sends a greeting
	Draw(context.Context, *DrawRequest) (*DrawReply, error)
}

// UnimplementedAwardServer can be embedded to have forward compatible implementations.
type UnimplementedAwardServer struct {
}

func (*UnimplementedAwardServer) Draw(ctx context.Context, req *DrawRequest) (*DrawReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Draw not implemented")
}

func RegisterAwardServer(s *grpc.Server, srv AwardServer) {
	s.RegisterService(&_Award_serviceDesc, srv)
}

func _Award_Draw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AwardServer).Draw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Award/Draw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AwardServer).Draw(ctx, req.(*DrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Award_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Award",
	HandlerType: (*AwardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Draw",
			Handler:    _Award_Draw_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "award.proto",
}