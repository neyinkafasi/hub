// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/provider/v2/msg.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/sentinel-official/hub/v1/types"
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

// MsgRegisterRequest defines the SDK message for registering a provider.
type MsgRegisterRequest struct {
	// Field 1: Sender's address initiating the registration.
	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	// Field 2: Provider name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Field 3: Identity of the provider.
	Identity string `protobuf:"bytes,3,opt,name=identity,proto3" json:"identity,omitempty"`
	// Field 4: Website associated with the provider.
	Website string `protobuf:"bytes,4,opt,name=website,proto3" json:"website,omitempty"`
	// Field 5: Description of the provider.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *MsgRegisterRequest) Reset()         { *m = MsgRegisterRequest{} }
func (m *MsgRegisterRequest) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterRequest) ProtoMessage()    {}
func (*MsgRegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_558ff44089bcc434, []int{0}
}
func (m *MsgRegisterRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterRequest.Merge(m, src)
}
func (m *MsgRegisterRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterRequest proto.InternalMessageInfo

// MsgUpdateRequest defines the SDK message for updating a provider.
type MsgUpdateRequest struct {
	// Field 1: Sender's address initiating the update.
	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	// Field 2: New provider name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Field 3: New identity of the provider.
	Identity string `protobuf:"bytes,3,opt,name=identity,proto3" json:"identity,omitempty"`
	// Field 4: New website associated with the provider.
	Website string `protobuf:"bytes,4,opt,name=website,proto3" json:"website,omitempty"`
	// Field 5: New description of the provider.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Field 6: New status of the provider, using the sentinel.types.v1.Status enum.
	Status types.Status `protobuf:"varint,6,opt,name=status,proto3,enum=sentinel.types.v1.Status" json:"status,omitempty"`
}

func (m *MsgUpdateRequest) Reset()         { *m = MsgUpdateRequest{} }
func (m *MsgUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateRequest) ProtoMessage()    {}
func (*MsgUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_558ff44089bcc434, []int{1}
}
func (m *MsgUpdateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateRequest.Merge(m, src)
}
func (m *MsgUpdateRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateRequest proto.InternalMessageInfo

// MsgRegisterResponse defines the response of message MsgRegisterRequest.
type MsgRegisterResponse struct {
}

func (m *MsgRegisterResponse) Reset()         { *m = MsgRegisterResponse{} }
func (m *MsgRegisterResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterResponse) ProtoMessage()    {}
func (*MsgRegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_558ff44089bcc434, []int{2}
}
func (m *MsgRegisterResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterResponse.Merge(m, src)
}
func (m *MsgRegisterResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterResponse proto.InternalMessageInfo

// MsgUpdateResponse defines the response of message MsgUpdateRequest.
type MsgUpdateResponse struct {
}

func (m *MsgUpdateResponse) Reset()         { *m = MsgUpdateResponse{} }
func (m *MsgUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateResponse) ProtoMessage()    {}
func (*MsgUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_558ff44089bcc434, []int{3}
}
func (m *MsgUpdateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateResponse.Merge(m, src)
}
func (m *MsgUpdateResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRegisterRequest)(nil), "sentinel.provider.v2.MsgRegisterRequest")
	proto.RegisterType((*MsgUpdateRequest)(nil), "sentinel.provider.v2.MsgUpdateRequest")
	proto.RegisterType((*MsgRegisterResponse)(nil), "sentinel.provider.v2.MsgRegisterResponse")
	proto.RegisterType((*MsgUpdateResponse)(nil), "sentinel.provider.v2.MsgUpdateResponse")
}

func init() { proto.RegisterFile("sentinel/provider/v2/msg.proto", fileDescriptor_558ff44089bcc434) }

var fileDescriptor_558ff44089bcc434 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0xbb, 0x0e, 0xd3, 0x30,
	0x14, 0x8d, 0xa1, 0x14, 0xea, 0x4a, 0x08, 0xdc, 0x22, 0x85, 0x0c, 0x56, 0x95, 0x01, 0xca, 0x80,
	0xad, 0x06, 0xbe, 0x80, 0xbd, 0x4b, 0x2a, 0x18, 0x10, 0x4b, 0x1e, 0xb7, 0xae, 0xa5, 0x26, 0x0e,
	0xb1, 0x13, 0xe8, 0x5f, 0x30, 0xf0, 0x11, 0xfc, 0x08, 0x52, 0x27, 0xd4, 0x91, 0x11, 0xd2, 0x1f,
	0x41, 0x71, 0x48, 0x1f, 0x02, 0x04, 0x23, 0xdb, 0xf5, 0x39, 0xc7, 0xd7, 0x37, 0xe7, 0xdc, 0x60,
	0xaa, 0x21, 0x37, 0x32, 0x87, 0x2d, 0x2f, 0x4a, 0x55, 0xcb, 0x14, 0x4a, 0x5e, 0x07, 0x3c, 0xd3,
	0x82, 0x15, 0xa5, 0x32, 0x8a, 0x4c, 0x7b, 0x9e, 0xf5, 0x3c, 0xab, 0x03, 0x6f, 0x2a, 0x94, 0x50,
	0x56, 0xc0, 0xdb, 0xaa, 0xd3, 0x7a, 0xe7, 0x5e, 0x66, 0x57, 0x80, 0xe6, 0xf5, 0x82, 0x6b, 0x13,
	0x99, 0x4a, 0x77, 0xbc, 0xff, 0x11, 0x61, 0xb2, 0xd4, 0x22, 0x04, 0x21, 0xb5, 0x81, 0x32, 0x84,
	0xb7, 0x15, 0x68, 0x43, 0x08, 0x1e, 0xac, 0x4b, 0x95, 0xb9, 0x68, 0x86, 0xe6, 0xa3, 0xd0, 0xd6,
	0x2d, 0x96, 0x47, 0x19, 0xb8, 0x37, 0x3a, 0xac, 0xad, 0x89, 0x87, 0xef, 0xc8, 0xb4, 0x7d, 0xc1,
	0xec, 0xdc, 0x9b, 0x16, 0x3f, 0x9d, 0x89, 0x8b, 0x6f, 0xbf, 0x83, 0x58, 0x4b, 0x03, 0xee, 0xc0,
	0x52, 0xfd, 0x91, 0xcc, 0xf0, 0x38, 0x05, 0x9d, 0x94, 0xb2, 0x30, 0x52, 0xe5, 0xee, 0x2d, 0xcb,
	0x5e, 0x42, 0xfe, 0x67, 0x84, 0xef, 0x2d, 0xb5, 0x78, 0x59, 0xa4, 0x91, 0x81, 0xff, 0x64, 0x28,
	0xb2, 0xc0, 0xc3, 0xce, 0x3b, 0x77, 0x38, 0x43, 0xf3, 0xbb, 0xc1, 0x43, 0x76, 0x0a, 0xc2, 0x9a,
	0xcb, 0xea, 0x05, 0x5b, 0x59, 0x41, 0xf8, 0x53, 0xe8, 0x3f, 0xc0, 0x93, 0x2b, 0x77, 0x75, 0xa1,
	0x72, 0x0d, 0xfe, 0x04, 0xdf, 0xbf, 0xf8, 0xba, 0x0e, 0x0c, 0xbe, 0x20, 0x8c, 0x97, 0x5a, 0xac,
	0xa0, 0xac, 0x65, 0x02, 0x24, 0xc6, 0xe3, 0x8b, 0xab, 0x64, 0xce, 0x7e, 0x97, 0x3a, 0xfb, 0x35,
	0x3b, 0xef, 0xc9, 0x3f, 0x28, 0xbb, 0x27, 0xc9, 0x1b, 0x3c, 0x3a, 0xcd, 0x41, 0x1e, 0xfd, 0xf1,
	0xde, 0x55, 0x0c, 0xde, 0xe3, 0xbf, 0xea, 0xba, 0xee, 0x2f, 0x5e, 0xed, 0xbf, 0x53, 0xe7, 0x53,
	0x43, 0x9d, 0x7d, 0x43, 0xd1, 0xa1, 0xa1, 0xe8, 0x5b, 0x43, 0xd1, 0x87, 0x23, 0x75, 0x0e, 0x47,
	0xea, 0x7c, 0x3d, 0x52, 0xe7, 0xf5, 0x73, 0x21, 0xcd, 0xa6, 0x8a, 0x59, 0xa2, 0x32, 0xde, 0x37,
	0x7d, 0xaa, 0xd6, 0x6b, 0x99, 0xc8, 0x68, 0xcb, 0x37, 0x55, 0xdc, 0xee, 0xeb, 0xfb, 0xf3, 0x6f,
	0x60, 0x6d, 0x8e, 0x87, 0x76, 0x75, 0x9f, 0xfd, 0x08, 0x00, 0x00, 0xff, 0xff, 0xda, 0xf8, 0xc2,
	0x82, 0x28, 0x03, 0x00, 0x00,
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
	// RPC method for registering a provider.
	MsgRegister(ctx context.Context, in *MsgRegisterRequest, opts ...grpc.CallOption) (*MsgRegisterResponse, error)
	// RPC method for updating a provider.
	MsgUpdate(ctx context.Context, in *MsgUpdateRequest, opts ...grpc.CallOption) (*MsgUpdateResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) MsgRegister(ctx context.Context, in *MsgRegisterRequest, opts ...grpc.CallOption) (*MsgRegisterResponse, error) {
	out := new(MsgRegisterResponse)
	err := c.cc.Invoke(ctx, "/sentinel.provider.v2.MsgService/MsgRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) MsgUpdate(ctx context.Context, in *MsgUpdateRequest, opts ...grpc.CallOption) (*MsgUpdateResponse, error) {
	out := new(MsgUpdateResponse)
	err := c.cc.Invoke(ctx, "/sentinel.provider.v2.MsgService/MsgUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	// RPC method for registering a provider.
	MsgRegister(context.Context, *MsgRegisterRequest) (*MsgRegisterResponse, error)
	// RPC method for updating a provider.
	MsgUpdate(context.Context, *MsgUpdateRequest) (*MsgUpdateResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) MsgRegister(ctx context.Context, req *MsgRegisterRequest) (*MsgRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgRegister not implemented")
}
func (*UnimplementedMsgServiceServer) MsgUpdate(ctx context.Context, req *MsgUpdateRequest) (*MsgUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgUpdate not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_MsgRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).MsgRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentinel.provider.v2.MsgService/MsgRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).MsgRegister(ctx, req.(*MsgRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_MsgUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).MsgUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sentinel.provider.v2.MsgService/MsgUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).MsgUpdate(ctx, req.(*MsgUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sentinel.provider.v2.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MsgRegister",
			Handler:    _MsgService_MsgRegister_Handler,
		},
		{
			MethodName: "MsgUpdate",
			Handler:    _MsgService_MsgUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sentinel/provider/v2/msg.proto",
}

func (m *MsgRegisterRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Website) > 0 {
		i -= len(m.Website)
		copy(dAtA[i:], m.Website)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Website)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Identity) > 0 {
		i -= len(m.Identity)
		copy(dAtA[i:], m.Identity)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Identity)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
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

func (m *MsgUpdateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Website) > 0 {
		i -= len(m.Website)
		copy(dAtA[i:], m.Website)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Website)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Identity) > 0 {
		i -= len(m.Identity)
		copy(dAtA[i:], m.Identity)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Identity)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
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

func (m *MsgRegisterResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgRegisterRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = len(m.Identity)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = len(m.Website)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	return n
}

func (m *MsgUpdateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = len(m.Identity)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = len(m.Website)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovMsg(uint64(m.Status))
	}
	return n
}

func (m *MsgRegisterResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateResponse) Size() (n int) {
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
func (m *MsgRegisterRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRegisterRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identity", wireType)
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
			m.Identity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Website", wireType)
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
			m.Website = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
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
			m.Description = string(dAtA[iNdEx:postIndex])
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
func (m *MsgUpdateRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identity", wireType)
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
			m.Identity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Website", wireType)
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
			m.Website = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
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
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= types.Status(b&0x7F) << shift
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
func (m *MsgRegisterResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRegisterResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgUpdateResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
