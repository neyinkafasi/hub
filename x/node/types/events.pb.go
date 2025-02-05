// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/node/v2/events.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/sentinel-official/hub/v1/types"
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

// EventRegister represents an event for registration.
type EventRegister struct {
	// Field 1: Address associated with the registration event.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
}

func (m *EventRegister) Reset()         { *m = EventRegister{} }
func (m *EventRegister) String() string { return proto.CompactTextString(m) }
func (*EventRegister) ProtoMessage()    {}
func (*EventRegister) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02d789f2b133ff1, []int{0}
}
func (m *EventRegister) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRegister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRegister.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRegister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRegister.Merge(m, src)
}
func (m *EventRegister) XXX_Size() int {
	return m.Size()
}
func (m *EventRegister) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRegister.DiscardUnknown(m)
}

var xxx_messageInfo_EventRegister proto.InternalMessageInfo

// EventUpdateDetails represents an event for updating details.
type EventUpdateDetails struct {
	// Field 1: Address associated with the update details event.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
}

func (m *EventUpdateDetails) Reset()         { *m = EventUpdateDetails{} }
func (m *EventUpdateDetails) String() string { return proto.CompactTextString(m) }
func (*EventUpdateDetails) ProtoMessage()    {}
func (*EventUpdateDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02d789f2b133ff1, []int{1}
}
func (m *EventUpdateDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventUpdateDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventUpdateDetails.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventUpdateDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUpdateDetails.Merge(m, src)
}
func (m *EventUpdateDetails) XXX_Size() int {
	return m.Size()
}
func (m *EventUpdateDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUpdateDetails.DiscardUnknown(m)
}

var xxx_messageInfo_EventUpdateDetails proto.InternalMessageInfo

// EventUpdateStatus represents an event for updating status.
type EventUpdateStatus struct {
	// Field 1: Status to be updated in the event.
	Status types.Status `protobuf:"varint,1,opt,name=status,proto3,enum=sentinel.types.v1.Status" json:"status,omitempty" yaml:"status"`
	// Field 2: Address associated with the update status event.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
}

func (m *EventUpdateStatus) Reset()         { *m = EventUpdateStatus{} }
func (m *EventUpdateStatus) String() string { return proto.CompactTextString(m) }
func (*EventUpdateStatus) ProtoMessage()    {}
func (*EventUpdateStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02d789f2b133ff1, []int{2}
}
func (m *EventUpdateStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventUpdateStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventUpdateStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventUpdateStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUpdateStatus.Merge(m, src)
}
func (m *EventUpdateStatus) XXX_Size() int {
	return m.Size()
}
func (m *EventUpdateStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUpdateStatus.DiscardUnknown(m)
}

var xxx_messageInfo_EventUpdateStatus proto.InternalMessageInfo

// EventCreateSubscription represents an event for creating a subscription.
type EventCreateSubscription struct {
	// Field 1: Address associated with the create subscription event.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	// Field 2: Node address associated with the create subscription event.
	NodeAddress string `protobuf:"bytes,2,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty" yaml:"node_address"`
	// Field 3: ID associated with the create subscription event.
	ID uint64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
}

func (m *EventCreateSubscription) Reset()         { *m = EventCreateSubscription{} }
func (m *EventCreateSubscription) String() string { return proto.CompactTextString(m) }
func (*EventCreateSubscription) ProtoMessage()    {}
func (*EventCreateSubscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02d789f2b133ff1, []int{3}
}
func (m *EventCreateSubscription) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCreateSubscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCreateSubscription.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCreateSubscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCreateSubscription.Merge(m, src)
}
func (m *EventCreateSubscription) XXX_Size() int {
	return m.Size()
}
func (m *EventCreateSubscription) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCreateSubscription.DiscardUnknown(m)
}

var xxx_messageInfo_EventCreateSubscription proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventRegister)(nil), "sentinel.node.v2.EventRegister")
	proto.RegisterType((*EventUpdateDetails)(nil), "sentinel.node.v2.EventUpdateDetails")
	proto.RegisterType((*EventUpdateStatus)(nil), "sentinel.node.v2.EventUpdateStatus")
	proto.RegisterType((*EventCreateSubscription)(nil), "sentinel.node.v2.EventCreateSubscription")
}

func init() { proto.RegisterFile("sentinel/node/v2/events.proto", fileDescriptor_d02d789f2b133ff1) }

var fileDescriptor_d02d789f2b133ff1 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xb1, 0x8e, 0xda, 0x30,
	0x18, 0xc7, 0xe3, 0xb4, 0xa2, 0xc2, 0x2d, 0xa8, 0x84, 0x4a, 0x50, 0xa4, 0x3a, 0x28, 0x5d, 0x18,
	0xda, 0xb8, 0xd0, 0x0d, 0xa9, 0x43, 0x53, 0x3a, 0x74, 0x4d, 0xd5, 0xa5, 0x4b, 0x95, 0x10, 0x13,
	0x2c, 0x85, 0x38, 0x8a, 0x9d, 0xa8, 0x3c, 0xc1, 0xad, 0xf7, 0x18, 0x3c, 0x0a, 0x23, 0xe3, 0x4d,
	0xd1, 0x5d, 0x78, 0x83, 0x3c, 0xc1, 0x29, 0x76, 0x40, 0xdc, 0x74, 0x62, 0xb3, 0xf5, 0xff, 0xfe,
	0x3f, 0xff, 0xbf, 0xcf, 0x1f, 0xfc, 0xc0, 0x49, 0x2c, 0x68, 0x4c, 0x22, 0x1c, 0xb3, 0x80, 0xe0,
	0x7c, 0x86, 0x49, 0x4e, 0x62, 0xc1, 0xed, 0x24, 0x65, 0x82, 0x19, 0x6f, 0x4f, 0xb2, 0x5d, 0xcb,
	0x76, 0x3e, 0x1b, 0xbd, 0x0b, 0x59, 0xc8, 0xa4, 0x88, 0xeb, 0x93, 0xaa, 0x1b, 0xa1, 0x33, 0x46,
	0x6c, 0x13, 0xc2, 0x71, 0x3e, 0xc5, 0x5c, 0x78, 0x22, 0x6b, 0x38, 0xd6, 0x37, 0xd8, 0xf9, 0x59,
	0x73, 0x5d, 0x12, 0x52, 0x2e, 0x48, 0x6a, 0x7c, 0x82, 0xaf, 0xbc, 0x20, 0x48, 0x09, 0xe7, 0x43,
	0x30, 0x06, 0x93, 0xb6, 0x63, 0x54, 0x85, 0xd9, 0xdd, 0x7a, 0x9b, 0x68, 0x6e, 0x35, 0x82, 0xe5,
	0x9e, 0x4a, 0x2c, 0x07, 0x1a, 0xd2, 0xfe, 0x27, 0x09, 0x3c, 0x41, 0x16, 0x44, 0x78, 0x34, 0xe2,
	0x57, 0x32, 0x6e, 0x00, 0xec, 0x5d, 0x40, 0x7e, 0xcb, 0x78, 0xc6, 0x02, 0xb6, 0x54, 0x50, 0x89,
	0xe8, 0xce, 0xde, 0xdb, 0xe7, 0x8e, 0x65, 0x27, 0x76, 0x3e, 0xb5, 0x55, 0xa9, 0xd3, 0xab, 0x0a,
	0xb3, 0xa3, 0xe8, 0xca, 0x62, 0xb9, 0x8d, 0xf7, 0x32, 0x89, 0xfe, 0x7c, 0x92, 0x1d, 0x80, 0x03,
	0x99, 0xe4, 0x47, 0x4a, 0xea, 0x24, 0x99, 0xcf, 0x97, 0x29, 0x4d, 0x04, 0x65, 0xf1, 0x75, 0x3d,
	0x19, 0x73, 0xf8, 0xa6, 0xfe, 0x97, 0x7f, 0x4f, 0x1f, 0x1f, 0x54, 0x85, 0xd9, 0x57, 0x96, 0x4b,
	0xd5, 0x72, 0x5f, 0xd7, 0xd7, 0xef, 0x8d, 0xf7, 0x23, 0xd4, 0x69, 0x30, 0x7c, 0x31, 0x06, 0x93,
	0x97, 0x4e, 0xbf, 0x2c, 0x4c, 0xfd, 0xd7, 0xa2, 0x2a, 0xcc, 0xb6, 0xf2, 0xd1, 0xc0, 0x72, 0x75,
	0x1a, 0x38, 0xee, 0xfe, 0x01, 0x69, 0xbb, 0x12, 0x69, 0xfb, 0x12, 0x81, 0x43, 0x89, 0xc0, 0x7d,
	0x89, 0xc0, 0xed, 0x11, 0x69, 0x87, 0x23, 0xd2, 0xee, 0x8e, 0x48, 0xfb, 0xfb, 0x25, 0xa4, 0x62,
	0x9d, 0xf9, 0xf6, 0x92, 0x6d, 0xf0, 0x69, 0x74, 0x9f, 0xd9, 0x6a, 0x45, 0x97, 0xd4, 0x8b, 0xf0,
	0x3a, 0xf3, 0xeb, 0x5d, 0xf8, 0xaf, 0xb6, 0x4b, 0x4e, 0xd4, 0x6f, 0xc9, 0x95, 0xf8, 0xfa, 0x18,
	0x00, 0x00, 0xff, 0xff, 0x34, 0x7e, 0xf8, 0x50, 0x7b, 0x02, 0x00, 0x00,
}

func (m *EventRegister) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRegister) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRegister) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventUpdateDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventUpdateDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventUpdateDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventUpdateStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventUpdateStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventUpdateStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.Status != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventCreateSubscription) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCreateSubscription) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCreateSubscription) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.NodeAddress) > 0 {
		i -= len(m.NodeAddress)
		copy(dAtA[i:], m.NodeAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.NodeAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventRegister) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventUpdateDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventUpdateStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovEvents(uint64(m.Status))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventCreateSubscription) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.NodeAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.ID != 0 {
		n += 1 + sovEvents(uint64(m.ID))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventRegister) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventRegister: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRegister: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventUpdateDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventUpdateDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventUpdateDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventUpdateStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventUpdateStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventUpdateStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventCreateSubscription) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventCreateSubscription: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCreateSubscription: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
