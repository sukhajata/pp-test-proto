// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pp-test.proto

package test

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

type ReceivedMessage_Source int32

const (
	ReceivedMessage_NO_SOURCE ReceivedMessage_Source = 0
	ReceivedMessage_MQTT      ReceivedMessage_Source = 1
	ReceivedMessage_KAFKA     ReceivedMessage_Source = 2
	ReceivedMessage_TIMESCALE ReceivedMessage_Source = 3
)

var ReceivedMessage_Source_name = map[int32]string{
	0: "NO_SOURCE",
	1: "MQTT",
	2: "KAFKA",
	3: "TIMESCALE",
}

var ReceivedMessage_Source_value = map[string]int32{
	"NO_SOURCE": 0,
	"MQTT":      1,
	"KAFKA":     2,
	"TIMESCALE": 3,
}

func (x ReceivedMessage_Source) String() string {
	return proto.EnumName(ReceivedMessage_Source_name, int32(x))
}

func (ReceivedMessage_Source) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a04cf5b57750b424, []int{0, 0}
}

type ReceivedMessage struct {
	Source               ReceivedMessage_Source `protobuf:"varint,1,opt,name=source,proto3,enum=test.ReceivedMessage_Source" json:"source,omitempty"`
	Type                 string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	DeviceEUI            string                 `protobuf:"bytes,3,opt,name=deviceEUI,proto3" json:"deviceEUI,omitempty"`
	Timestamp            string                 `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Phase                int32                  `protobuf:"varint,5,opt,name=phase,proto3" json:"phase,omitempty"`
	Content              string                 `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ReceivedMessage) Reset()         { *m = ReceivedMessage{} }
func (m *ReceivedMessage) String() string { return proto.CompactTextString(m) }
func (*ReceivedMessage) ProtoMessage()    {}
func (*ReceivedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a04cf5b57750b424, []int{0}
}

func (m *ReceivedMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceivedMessage.Unmarshal(m, b)
}
func (m *ReceivedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceivedMessage.Marshal(b, m, deterministic)
}
func (m *ReceivedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceivedMessage.Merge(m, src)
}
func (m *ReceivedMessage) XXX_Size() int {
	return xxx_messageInfo_ReceivedMessage.Size(m)
}
func (m *ReceivedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceivedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ReceivedMessage proto.InternalMessageInfo

func (m *ReceivedMessage) GetSource() ReceivedMessage_Source {
	if m != nil {
		return m.Source
	}
	return ReceivedMessage_NO_SOURCE
}

func (m *ReceivedMessage) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ReceivedMessage) GetDeviceEUI() string {
	if m != nil {
		return m.DeviceEUI
	}
	return ""
}

func (m *ReceivedMessage) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *ReceivedMessage) GetPhase() int32 {
	if m != nil {
		return m.Phase
	}
	return 0
}

func (m *ReceivedMessage) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type Response struct {
	Reply                string   `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_a04cf5b57750b424, []int{1}
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

func (m *Response) GetReply() string {
	if m != nil {
		return m.Reply
	}
	return ""
}

func init() {
	proto.RegisterEnum("test.ReceivedMessage_Source", ReceivedMessage_Source_name, ReceivedMessage_Source_value)
	proto.RegisterType((*ReceivedMessage)(nil), "test.ReceivedMessage")
	proto.RegisterType((*Response)(nil), "test.Response")
}

func init() { proto.RegisterFile("pp-test.proto", fileDescriptor_a04cf5b57750b424) }

var fileDescriptor_a04cf5b57750b424 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x6d, 0xda, 0x24, 0x36, 0x23, 0xad, 0x71, 0x50, 0x58, 0xa4, 0x87, 0x90, 0x53, 0x2f, 0xf6,
	0x50, 0xbd, 0xe9, 0x25, 0x94, 0x08, 0x25, 0xc6, 0xe0, 0x26, 0x3d, 0x4b, 0x4c, 0x07, 0x0d, 0xd8,
	0x64, 0xc9, 0xae, 0x85, 0xde, 0xfd, 0xe1, 0x92, 0x4d, 0xa3, 0x20, 0xde, 0xe6, 0x7d, 0x0c, 0xf3,
	0xf6, 0x2d, 0x4c, 0x84, 0xb8, 0x56, 0x24, 0xd5, 0x42, 0x34, 0xb5, 0xaa, 0xd1, 0x6c, 0x67, 0xff,
	0x6b, 0x08, 0x67, 0x9c, 0x0a, 0x2a, 0xf7, 0xb4, 0x8d, 0x49, 0xca, 0xfc, 0x8d, 0xf0, 0x16, 0x6c,
	0x59, 0x7f, 0x36, 0x05, 0x31, 0xc3, 0x33, 0xe6, 0xd3, 0xe5, 0x6c, 0xa1, 0xd7, 0xfe, 0xd8, 0x16,
	0xa9, 0xf6, 0xf0, 0xa3, 0x17, 0x11, 0x4c, 0x75, 0x10, 0xc4, 0x86, 0x9e, 0x31, 0x77, 0xb8, 0x9e,
	0x71, 0x06, 0xce, 0x96, 0xf6, 0x65, 0x41, 0xe1, 0x66, 0xcd, 0x46, 0x5a, 0xf8, 0x25, 0x5a, 0x55,
	0x95, 0x3b, 0x92, 0x2a, 0xdf, 0x09, 0x66, 0x76, 0xea, 0x0f, 0x81, 0x17, 0x60, 0x89, 0xf7, 0x5c,
	0x12, 0xb3, 0x3c, 0x63, 0x6e, 0xf1, 0x0e, 0x20, 0x83, 0x93, 0xa2, 0xae, 0x14, 0x55, 0x8a, 0xd9,
	0x7a, 0xa3, 0x87, 0xfe, 0x1d, 0xd8, 0x5d, 0x22, 0x9c, 0x80, 0xf3, 0x94, 0xbc, 0xa4, 0xc9, 0x86,
	0xaf, 0x42, 0x77, 0x80, 0x63, 0x30, 0xe3, 0xe7, 0x2c, 0x73, 0x0d, 0x74, 0xc0, 0x8a, 0x82, 0x87,
	0x28, 0x70, 0x87, 0xad, 0x27, 0x5b, 0xc7, 0x61, 0xba, 0x0a, 0x1e, 0x43, 0x77, 0xe4, 0x7b, 0x30,
	0xe6, 0x24, 0x45, 0x5d, 0x49, 0x6a, 0x0f, 0x37, 0x24, 0x3e, 0x0e, 0xfa, 0xf5, 0x0e, 0xef, 0xc0,
	0x32, 0x82, 0xd3, 0x8c, 0xa4, 0x4a, 0xa9, 0x69, 0xe3, 0xe3, 0x3d, 0x9c, 0x27, 0xd5, 0xb1, 0x89,
	0xbe, 0x18, 0xbc, 0xfc, 0xb7, 0xa8, 0xab, 0x69, 0x4f, 0x77, 0x07, 0xfc, 0xc1, 0xab, 0xad, 0xbf,
	0xe0, 0xe6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xae, 0x05, 0x98, 0x93, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestServiceClient interface {
	OnMessageReceived(ctx context.Context, in *ReceivedMessage, opts ...grpc.CallOption) (*Response, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) OnMessageReceived(ctx context.Context, in *ReceivedMessage, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/test.TestService/OnMessageReceived", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceServer is the server API for TestService service.
type TestServiceServer interface {
	OnMessageReceived(context.Context, *ReceivedMessage) (*Response, error)
}

// UnimplementedTestServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct {
}

func (*UnimplementedTestServiceServer) OnMessageReceived(ctx context.Context, req *ReceivedMessage) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnMessageReceived not implemented")
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_OnMessageReceived_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceivedMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).OnMessageReceived(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.TestService/OnMessageReceived",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).OnMessageReceived(ctx, req.(*ReceivedMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnMessageReceived",
			Handler:    _TestService_OnMessageReceived_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pp-test.proto",
}
