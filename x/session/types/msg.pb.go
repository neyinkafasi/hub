// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/session/v2/msg.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgStartRequest defines the SDK message for starting a session
type MsgStartRequest struct {
	From        string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	ID          uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	NodeAddress string `protobuf:"bytes,3,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty"`
}

func (m *MsgStartRequest) Reset()         { *m = MsgStartRequest{} }
func (m *MsgStartRequest) String() string { return proto.CompactTextString(m) }
func (*MsgStartRequest) ProtoMessage()    {}
func (*MsgStartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6db1be5556425867, []int{0}
}
func (m *MsgStartRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStartRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStartRequest.Merge(m, src)
}
func (m *MsgStartRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgStartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStartRequest proto.InternalMessageInfo

// MsgUpdateDetailsRequest defines the SDK message for updating a session
type MsgUpdateDetailsRequest struct {
	From      string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Proof     Proof  `protobuf:"bytes,2,opt,name=proof,proto3" json:"proof"`
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *MsgUpdateDetailsRequest) Reset()         { *m = MsgUpdateDetailsRequest{} }
func (m *MsgUpdateDetailsRequest) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateDetailsRequest) ProtoMessage()    {}
func (*MsgUpdateDetailsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6db1be5556425867, []int{1}
}
func (m *MsgUpdateDetailsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateDetailsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateDetailsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateDetailsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateDetailsRequest.Merge(m, src)
}
func (m *MsgUpdateDetailsRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateDetailsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateDetailsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateDetailsRequest proto.InternalMessageInfo

// MsgEndRequest defines the SDK message for ending a session
type MsgEndRequest struct {
	From   string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	ID     uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Rating uint64 `protobuf:"varint,3,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (m *MsgEndRequest) Reset()         { *m = MsgEndRequest{} }
func (m *MsgEndRequest) String() string { return proto.CompactTextString(m) }
func (*MsgEndRequest) ProtoMessage()    {}
func (*MsgEndRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6db1be5556425867, []int{2}
}
func (m *MsgEndRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgEndRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgEndRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgEndRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgEndRequest.Merge(m, src)
}
func (m *MsgEndRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgEndRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgEndRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgEndRequest proto.InternalMessageInfo

// MsgStartResponse defines the response of message MsgStartRequest
type MsgStartResponse struct {
}

func (m *MsgStartResponse) Reset()         { *m = MsgStartResponse{} }
func (m *MsgStartResponse) String() string { return proto.CompactTextString(m) }
func (*MsgStartResponse) ProtoMessage()    {}
func (*MsgStartResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6db1be5556425867, []int{3}
}
func (m *MsgStartResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStartResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStartResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStartResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStartResponse.Merge(m, src)
}
func (m *MsgStartResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgStartResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStartResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStartResponse proto.InternalMessageInfo

// MsgUpdateDetailsResponse defines the response of message
// MsgUpdateDetailsRequest
type MsgUpdateDetailsResponse struct {
}

func (m *MsgUpdateDetailsResponse) Reset()         { *m = MsgUpdateDetailsResponse{} }
func (m *MsgUpdateDetailsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateDetailsResponse) ProtoMessage()    {}
func (*MsgUpdateDetailsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6db1be5556425867, []int{4}
}
func (m *MsgUpdateDetailsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateDetailsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateDetailsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateDetailsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateDetailsResponse.Merge(m, src)
}
func (m *MsgUpdateDetailsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateDetailsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateDetailsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateDetailsResponse proto.InternalMessageInfo

// MsgEndResponse defines the response of message MsgEndRequest
type MsgEndResponse struct {
}

func (m *MsgEndResponse) Reset()         { *m = MsgEndResponse{} }
func (m *MsgEndResponse) String() string { return proto.CompactTextString(m) }
func (*MsgEndResponse) ProtoMessage()    {}
func (*MsgEndResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6db1be5556425867, []int{5}
}
func (m *MsgEndResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgEndResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgEndResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgEndResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgEndResponse.Merge(m, src)
}
func (m *MsgEndResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgEndResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgEndResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgEndResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgStartRequest)(nil), "sentinel.session.v2.MsgStartRequest")
	proto.RegisterType((*MsgUpdateDetailsRequest)(nil), "sentinel.session.v2.MsgUpdateDetailsRequest")
	proto.RegisterType((*MsgEndRequest)(nil), "sentinel.session.v2.MsgEndRequest")
	proto.RegisterType((*MsgStartResponse)(nil), "sentinel.session.v2.MsgStartResponse")
	proto.RegisterType((*MsgUpdateDetailsResponse)(nil), "sentinel.session.v2.MsgUpdateDetailsResponse")
	proto.RegisterType((*MsgEndResponse)(nil), "sentinel.session.v2.MsgEndResponse")
}

func init() { proto.RegisterFile("sentinel/session/v2/msg.proto", fileDescriptor_6db1be5556425867) }

var fileDescriptor_6db1be5556425867 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xb6, 0x4d, 0x88, 0xe8, 0xb4, 0x40, 0xb5, 0xa0, 0x12, 0x59, 0xb0, 0x29, 0x06, 0xa4, 0x1e,
	0xa8, 0x8d, 0x82, 0xc4, 0x9d, 0xa8, 0x1c, 0x38, 0x44, 0x02, 0x57, 0x08, 0x89, 0x0b, 0x38, 0xf1,
	0x66, 0xbb, 0x52, 0xb2, 0x6b, 0x76, 0x36, 0x11, 0x9c, 0x79, 0x01, 0x1e, 0x03, 0xde, 0x24, 0xc7,
	0x1e, 0x39, 0x55, 0xe0, 0xbc, 0x08, 0xca, 0xae, 0x4d, 0x04, 0xb8, 0x01, 0xf5, 0xb6, 0x9e, 0xf9,
	0xe6, 0xfb, 0xf1, 0x68, 0xe0, 0x0e, 0x32, 0x69, 0x84, 0x64, 0x93, 0x04, 0x19, 0xa2, 0x50, 0x32,
	0x99, 0xf7, 0x92, 0x29, 0xf2, 0xb8, 0xd0, 0xca, 0x28, 0x72, 0xa3, 0x6e, 0xc7, 0x55, 0x3b, 0x9e,
	0xf7, 0xc2, 0x9b, 0x5c, 0x71, 0x65, 0xfb, 0xc9, 0xea, 0xe5, 0xa0, 0x61, 0xb7, 0x89, 0xa9, 0xd0,
	0x4a, 0x8d, 0x1d, 0x20, 0x7a, 0x07, 0xd7, 0x07, 0xc8, 0x8f, 0x4d, 0xa6, 0x4d, 0xca, 0xde, 0xcf,
	0x18, 0x1a, 0x42, 0xa0, 0x35, 0xd6, 0x6a, 0xda, 0xf1, 0xf7, 0xfd, 0x83, 0xad, 0xd4, 0xbe, 0xc9,
	0x1e, 0x04, 0x22, 0xef, 0x04, 0xfb, 0xfe, 0x41, 0xab, 0xdf, 0x2e, 0xcf, 0xba, 0xc1, 0xf3, 0xa3,
	0x34, 0x10, 0x39, 0xb9, 0x0b, 0x3b, 0x52, 0xe5, 0xec, 0x6d, 0x96, 0xe7, 0x9a, 0x21, 0x76, 0x2e,
	0xd9, 0x99, 0xed, 0x55, 0xed, 0xa9, 0x2b, 0x45, 0x9f, 0x7c, 0xb8, 0x35, 0x40, 0xfe, 0xaa, 0xc8,
	0x33, 0xc3, 0x8e, 0x98, 0xc9, 0xc4, 0x04, 0x37, 0x49, 0x3d, 0x81, 0xcb, 0xd6, 0xa0, 0x55, 0xdb,
	0xee, 0x85, 0x71, 0x43, 0xda, 0xf8, 0xc5, 0x0a, 0xd1, 0x6f, 0x2d, 0xce, 0xba, 0x5e, 0xea, 0xe0,
	0xe4, 0x36, 0x6c, 0xa1, 0xe0, 0x32, 0x33, 0x33, 0xcd, 0xac, 0x8f, 0x9d, 0x74, 0x5d, 0x88, 0x8e,
	0xe1, 0xea, 0x00, 0xf9, 0x33, 0x99, 0x5f, 0x24, 0xe5, 0x1e, 0xb4, 0x75, 0x66, 0x84, 0xe4, 0x96,
	0xb7, 0x95, 0x56, 0x5f, 0x11, 0x81, 0xdd, 0xf5, 0xcf, 0xc3, 0x42, 0x49, 0x64, 0x51, 0x08, 0x9d,
	0xbf, 0xd3, 0x56, 0xbd, 0x5d, 0xb8, 0x56, 0x9b, 0x70, 0x95, 0xde, 0xd7, 0x00, 0x60, 0x45, 0xc1,
	0xf4, 0x5c, 0x8c, 0x18, 0x79, 0x0d, 0x57, 0x6a, 0x42, 0x72, 0xbf, 0x31, 0xf8, 0x1f, 0xcb, 0x0a,
	0x1f, 0xfc, 0x03, 0xe5, 0x74, 0x88, 0xb2, 0x4e, 0x7f, 0x73, 0x45, 0x1e, 0x9e, 0x37, 0xda, 0xb4,
	0xaa, 0xf0, 0xf0, 0x3f, 0xd1, 0x95, 0xe0, 0x4b, 0x68, 0xbb, 0xa8, 0x24, 0x3a, 0x6f, 0x70, 0xbd,
	0x8c, 0xf0, 0xde, 0x46, 0x8c, 0xa3, 0xec, 0xa7, 0x8b, 0x1f, 0xd4, 0xfb, 0x52, 0x52, 0x6f, 0x51,
	0x52, 0xff, 0xb4, 0xa4, 0xfe, 0xf7, 0x92, 0xfa, 0x9f, 0x97, 0xd4, 0x3b, 0x5d, 0x52, 0xef, 0xdb,
	0x92, 0x7a, 0x6f, 0x1e, 0x71, 0x61, 0x4e, 0x66, 0xc3, 0x78, 0xa4, 0xa6, 0x49, 0x4d, 0x78, 0xa8,
	0xc6, 0x63, 0x31, 0x12, 0xd9, 0x24, 0x39, 0x99, 0x0d, 0x93, 0x0f, 0xbf, 0xee, 0xc0, 0x7c, 0x2c,
	0x18, 0x0e, 0xdb, 0xf6, 0x0a, 0x1e, 0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x21, 0x0a, 0x63, 0xab,
	0x72, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
	MsgStart(ctx context.Context, in *MsgStartRequest, opts ...grpc.CallOption) (*MsgStartResponse, error)
	MsgUpdateDetails(ctx context.Context, in *MsgUpdateDetailsRequest, opts ...grpc.CallOption) (*MsgUpdateDetailsResponse, error)
	MsgEnd(ctx context.Context, in *MsgEndRequest, opts ...grpc.CallOption) (*MsgEndResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) MsgStart(ctx context.Context, in *MsgStartRequest, opts ...grpc.CallOption) (*MsgStartResponse, error) {
	out := new(MsgStartResponse)
	err := c.cc.Invoke(ctx, "/sentinel.session.v2.MsgService/MsgStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) MsgUpdateDetails(ctx context.Context, in *MsgUpdateDetailsRequest, opts ...grpc.CallOption) (*MsgUpdateDetailsResponse, error) {
	out := new(MsgUpdateDetailsResponse)
	err := c.cc.Invoke(ctx, "/sentinel.session.v2.MsgService/MsgUpdateDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) MsgEnd(ctx context.Context, in *MsgEndRequest, opts ...grpc.CallOption) (*MsgEndResponse, error) {
	out := new(MsgEndResponse)
	err := c.cc.Invoke(ctx, "/sentinel.session.v2.MsgService/MsgEnd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	MsgStart(context.Context, *MsgStartRequest) (*MsgStartResponse, error)
	MsgUpdateDetails(context.Context, *MsgUpdateDetailsRequest) (*MsgUpdateDetailsResponse, error)
	MsgEnd(context.Context, *MsgEndRequest) (*MsgEndResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) MsgStart(ctx context.Context, req *MsgStartRequest) (*MsgStartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgStart not implemented")
}
func (*UnimplementedMsgServiceServer) MsgUpdateDetails(ctx context.Context, req *MsgUpdateDetailsRequest) (*MsgUpdateDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgUpdateDetails not implemented")
}
func (*UnimplementedMsgServiceServer) MsgEnd(ctx context.Context, req *MsgEndRequest) (*MsgEndResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgEnd not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_MsgStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).MsgStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentinel.session.v2.MsgService/MsgStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).MsgStart(ctx, req.(*MsgStartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_MsgUpdateDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).MsgUpdateDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentinel.session.v2.MsgService/MsgUpdateDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).MsgUpdateDetails(ctx, req.(*MsgUpdateDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_MsgEnd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgEndRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).MsgEnd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentinel.session.v2.MsgService/MsgEnd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).MsgEnd(ctx, req.(*MsgEndRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sentinel.session.v2.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MsgStart",
			Handler:    _MsgService_MsgStart_Handler,
		},
		{
			MethodName: "MsgUpdateDetails",
			Handler:    _MsgService_MsgUpdateDetails_Handler,
		},
		{
			MethodName: "MsgEnd",
			Handler:    _MsgService_MsgEnd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sentinel/session/v2/msg.proto",
}

func (m *MsgStartRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStartRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStartRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NodeAddress) > 0 {
		i -= len(m.NodeAddress)
		copy(dAtA[i:], m.NodeAddress)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.NodeAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ID != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateDetailsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateDetailsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateDetailsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.Proof.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgEndRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgEndRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgEndRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Rating != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.Rating))
		i--
		dAtA[i] = 0x18
	}
	if m.ID != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgStartResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStartResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStartResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateDetailsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateDetailsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateDetailsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgEndResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgEndResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgEndResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgStartRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.ID != 0 {
		n += 1 + sovMsg(uint64(m.ID))
	}
	l = len(m.NodeAddress)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	return n
}

func (m *MsgUpdateDetailsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = m.Proof.Size()
	n += 1 + l + sovMsg(uint64(l))
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	return n
}

func (m *MsgEndRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.ID != 0 {
		n += 1 + sovMsg(uint64(m.ID))
	}
	if m.Rating != 0 {
		n += 1 + sovMsg(uint64(m.Rating))
	}
	return n
}

func (m *MsgStartResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateDetailsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgEndResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsg(x uint64) (n int) {
	return sovMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgStartRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgStartRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStartRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateDetailsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateDetailsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateDetailsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Proof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgEndRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgEndRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgEndRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rating", wireType)
			}
			m.Rating = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rating |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgStartResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgStartResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStartResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateDetailsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateDetailsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateDetailsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgEndResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgEndResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgEndResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsg
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsg = fmt.Errorf("proto: unexpected end of group")
)
