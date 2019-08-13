// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: deadlock.proto

package deadlockpb

import (
	"fmt"
	"io"
	"math"

	proto "github.com/golang/protobuf/proto"

	_ "github.com/gogo/protobuf/gogoproto"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
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

type DeadlockRequestType int32

const (
	DeadlockRequestType_Detect DeadlockRequestType = 0
	// CleanUpWaitFor cleans a single entry the transaction is waiting.
	DeadlockRequestType_CleanUpWaitFor DeadlockRequestType = 1
	// CleanUp cleans all entries the transaction is waiting.
	DeadlockRequestType_CleanUp DeadlockRequestType = 2
)

var DeadlockRequestType_name = map[int32]string{
	0: "Detect",
	1: "CleanUpWaitFor",
	2: "CleanUp",
}
var DeadlockRequestType_value = map[string]int32{
	"Detect":         0,
	"CleanUpWaitFor": 1,
	"CleanUp":        2,
}

func (x DeadlockRequestType) String() string {
	return proto.EnumName(DeadlockRequestType_name, int32(x))
}
func (DeadlockRequestType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_deadlock_3c9646ad0923eabd, []int{0}
}

type WaitForEntriesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WaitForEntriesRequest) Reset()         { *m = WaitForEntriesRequest{} }
func (m *WaitForEntriesRequest) String() string { return proto.CompactTextString(m) }
func (*WaitForEntriesRequest) ProtoMessage()    {}
func (*WaitForEntriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_deadlock_3c9646ad0923eabd, []int{0}
}
func (m *WaitForEntriesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WaitForEntriesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WaitForEntriesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *WaitForEntriesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WaitForEntriesRequest.Merge(dst, src)
}
func (m *WaitForEntriesRequest) XXX_Size() int {
	return m.Size()
}
func (m *WaitForEntriesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WaitForEntriesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WaitForEntriesRequest proto.InternalMessageInfo

type WaitForEntriesResponse struct {
	Entries              []WaitForEntry `protobuf:"bytes,1,rep,name=entries" json:"entries"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *WaitForEntriesResponse) Reset()         { *m = WaitForEntriesResponse{} }
func (m *WaitForEntriesResponse) String() string { return proto.CompactTextString(m) }
func (*WaitForEntriesResponse) ProtoMessage()    {}
func (*WaitForEntriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_deadlock_3c9646ad0923eabd, []int{1}
}
func (m *WaitForEntriesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WaitForEntriesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WaitForEntriesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *WaitForEntriesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WaitForEntriesResponse.Merge(dst, src)
}
func (m *WaitForEntriesResponse) XXX_Size() int {
	return m.Size()
}
func (m *WaitForEntriesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WaitForEntriesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WaitForEntriesResponse proto.InternalMessageInfo

func (m *WaitForEntriesResponse) GetEntries() []WaitForEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type WaitForEntry struct {
	// The transaction id that is waiting.
	Txn uint64 `protobuf:"varint,1,opt,name=txn,proto3" json:"txn,omitempty"`
	// The transaction id that is being waited for.
	WaitForTxn uint64 `protobuf:"varint,2,opt,name=wait_for_txn,json=waitForTxn,proto3" json:"wait_for_txn,omitempty"`
	// The hash value of the key is being waited for.
	KeyHash              uint64   `protobuf:"varint,3,opt,name=key_hash,json=keyHash,proto3" json:"key_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WaitForEntry) Reset()         { *m = WaitForEntry{} }
func (m *WaitForEntry) String() string { return proto.CompactTextString(m) }
func (*WaitForEntry) ProtoMessage()    {}
func (*WaitForEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_deadlock_3c9646ad0923eabd, []int{2}
}
func (m *WaitForEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WaitForEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WaitForEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *WaitForEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WaitForEntry.Merge(dst, src)
}
func (m *WaitForEntry) XXX_Size() int {
	return m.Size()
}
func (m *WaitForEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_WaitForEntry.DiscardUnknown(m)
}

var xxx_messageInfo_WaitForEntry proto.InternalMessageInfo

func (m *WaitForEntry) GetTxn() uint64 {
	if m != nil {
		return m.Txn
	}
	return 0
}

func (m *WaitForEntry) GetWaitForTxn() uint64 {
	if m != nil {
		return m.WaitForTxn
	}
	return 0
}

func (m *WaitForEntry) GetKeyHash() uint64 {
	if m != nil {
		return m.KeyHash
	}
	return 0
}

type DeadlockRequest struct {
	Tp                   DeadlockRequestType `protobuf:"varint,1,opt,name=tp,proto3,enum=deadlockpb.DeadlockRequestType" json:"tp,omitempty"`
	Entry                WaitForEntry        `protobuf:"bytes,2,opt,name=entry" json:"entry"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *DeadlockRequest) Reset()         { *m = DeadlockRequest{} }
func (m *DeadlockRequest) String() string { return proto.CompactTextString(m) }
func (*DeadlockRequest) ProtoMessage()    {}
func (*DeadlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_deadlock_3c9646ad0923eabd, []int{3}
}
func (m *DeadlockRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeadlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeadlockRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *DeadlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeadlockRequest.Merge(dst, src)
}
func (m *DeadlockRequest) XXX_Size() int {
	return m.Size()
}
func (m *DeadlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeadlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeadlockRequest proto.InternalMessageInfo

func (m *DeadlockRequest) GetTp() DeadlockRequestType {
	if m != nil {
		return m.Tp
	}
	return DeadlockRequestType_Detect
}

func (m *DeadlockRequest) GetEntry() WaitForEntry {
	if m != nil {
		return m.Entry
	}
	return WaitForEntry{}
}

type DeadlockResponse struct {
	// The same entry sent by DeadlockRequest, identifies the sender.
	Entry WaitForEntry `protobuf:"bytes,1,opt,name=entry" json:"entry"`
	// The key hash of the lock that is hold by the waiting transaction.
	DeadlockKeyHash      uint64   `protobuf:"varint,2,opt,name=deadlock_key_hash,json=deadlockKeyHash,proto3" json:"deadlock_key_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeadlockResponse) Reset()         { *m = DeadlockResponse{} }
func (m *DeadlockResponse) String() string { return proto.CompactTextString(m) }
func (*DeadlockResponse) ProtoMessage()    {}
func (*DeadlockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_deadlock_3c9646ad0923eabd, []int{4}
}
func (m *DeadlockResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeadlockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeadlockResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *DeadlockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeadlockResponse.Merge(dst, src)
}
func (m *DeadlockResponse) XXX_Size() int {
	return m.Size()
}
func (m *DeadlockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeadlockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeadlockResponse proto.InternalMessageInfo

func (m *DeadlockResponse) GetEntry() WaitForEntry {
	if m != nil {
		return m.Entry
	}
	return WaitForEntry{}
}

func (m *DeadlockResponse) GetDeadlockKeyHash() uint64 {
	if m != nil {
		return m.DeadlockKeyHash
	}
	return 0
}

func init() {
	proto.RegisterType((*WaitForEntriesRequest)(nil), "deadlockpb.WaitForEntriesRequest")
	proto.RegisterType((*WaitForEntriesResponse)(nil), "deadlockpb.WaitForEntriesResponse")
	proto.RegisterType((*WaitForEntry)(nil), "deadlockpb.WaitForEntry")
	proto.RegisterType((*DeadlockRequest)(nil), "deadlockpb.DeadlockRequest")
	proto.RegisterType((*DeadlockResponse)(nil), "deadlockpb.DeadlockResponse")
	proto.RegisterEnum("deadlockpb.DeadlockRequestType", DeadlockRequestType_name, DeadlockRequestType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Deadlock service

type DeadlockClient interface {
	// Get local wait for entries, should be handle by every node.
	// The owner should sent this request to all members to build the complete wait for graph.
	GetWaitForEntries(ctx context.Context, in *WaitForEntriesRequest, opts ...grpc.CallOption) (*WaitForEntriesResponse, error)
	// Detect should only sent to the owner. only be handled by the owner.
	// The DeadlockResponse is sent back only if there is deadlock detected.
	// CleanUpWaitFor and CleanUp doesn't return responses.
	Detect(ctx context.Context, opts ...grpc.CallOption) (Deadlock_DetectClient, error)
}

type deadlockClient struct {
	cc *grpc.ClientConn
}

func NewDeadlockClient(cc *grpc.ClientConn) DeadlockClient {
	return &deadlockClient{cc}
}

func (c *deadlockClient) GetWaitForEntries(ctx context.Context, in *WaitForEntriesRequest, opts ...grpc.CallOption) (*WaitForEntriesResponse, error) {
	out := new(WaitForEntriesResponse)
	err := c.cc.Invoke(ctx, "/deadlockpb.Deadlock/GetWaitForEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deadlockClient) Detect(ctx context.Context, opts ...grpc.CallOption) (Deadlock_DetectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Deadlock_serviceDesc.Streams[0], "/deadlockpb.Deadlock/Detect", opts...)
	if err != nil {
		return nil, err
	}
	x := &deadlockDetectClient{stream}
	return x, nil
}

type Deadlock_DetectClient interface {
	Send(*DeadlockRequest) error
	Recv() (*DeadlockResponse, error)
	grpc.ClientStream
}

type deadlockDetectClient struct {
	grpc.ClientStream
}

func (x *deadlockDetectClient) Send(m *DeadlockRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *deadlockDetectClient) Recv() (*DeadlockResponse, error) {
	m := new(DeadlockResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Deadlock service

type DeadlockServer interface {
	// Get local wait for entries, should be handle by every node.
	// The owner should sent this request to all members to build the complete wait for graph.
	GetWaitForEntries(context.Context, *WaitForEntriesRequest) (*WaitForEntriesResponse, error)
	// Detect should only sent to the owner. only be handled by the owner.
	// The DeadlockResponse is sent back only if there is deadlock detected.
	// CleanUpWaitFor and CleanUp doesn't return responses.
	Detect(Deadlock_DetectServer) error
}

func RegisterDeadlockServer(s *grpc.Server, srv DeadlockServer) {
	s.RegisterService(&_Deadlock_serviceDesc, srv)
}

func _Deadlock_GetWaitForEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaitForEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeadlockServer).GetWaitForEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deadlockpb.Deadlock/GetWaitForEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeadlockServer).GetWaitForEntries(ctx, req.(*WaitForEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deadlock_Detect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DeadlockServer).Detect(&deadlockDetectServer{stream})
}

type Deadlock_DetectServer interface {
	Send(*DeadlockResponse) error
	Recv() (*DeadlockRequest, error)
	grpc.ServerStream
}

type deadlockDetectServer struct {
	grpc.ServerStream
}

func (x *deadlockDetectServer) Send(m *DeadlockResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *deadlockDetectServer) Recv() (*DeadlockRequest, error) {
	m := new(DeadlockRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Deadlock_serviceDesc = grpc.ServiceDesc{
	ServiceName: "deadlockpb.Deadlock",
	HandlerType: (*DeadlockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWaitForEntries",
			Handler:    _Deadlock_GetWaitForEntries_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Detect",
			Handler:       _Deadlock_Detect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "deadlock.proto",
}

func (m *WaitForEntriesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WaitForEntriesRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *WaitForEntriesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WaitForEntriesResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for _, msg := range m.Entries {
			dAtA[i] = 0xa
			i++
			i = encodeVarintDeadlock(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *WaitForEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WaitForEntry) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Txn != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintDeadlock(dAtA, i, uint64(m.Txn))
	}
	if m.WaitForTxn != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintDeadlock(dAtA, i, uint64(m.WaitForTxn))
	}
	if m.KeyHash != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintDeadlock(dAtA, i, uint64(m.KeyHash))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *DeadlockRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeadlockRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Tp != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintDeadlock(dAtA, i, uint64(m.Tp))
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintDeadlock(dAtA, i, uint64(m.Entry.Size()))
	n1, err := m.Entry.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *DeadlockResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeadlockResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintDeadlock(dAtA, i, uint64(m.Entry.Size()))
	n2, err := m.Entry.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.DeadlockKeyHash != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintDeadlock(dAtA, i, uint64(m.DeadlockKeyHash))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintDeadlock(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *WaitForEntriesRequest) Size() (n int) {
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *WaitForEntriesResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovDeadlock(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *WaitForEntry) Size() (n int) {
	var l int
	_ = l
	if m.Txn != 0 {
		n += 1 + sovDeadlock(uint64(m.Txn))
	}
	if m.WaitForTxn != 0 {
		n += 1 + sovDeadlock(uint64(m.WaitForTxn))
	}
	if m.KeyHash != 0 {
		n += 1 + sovDeadlock(uint64(m.KeyHash))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DeadlockRequest) Size() (n int) {
	var l int
	_ = l
	if m.Tp != 0 {
		n += 1 + sovDeadlock(uint64(m.Tp))
	}
	l = m.Entry.Size()
	n += 1 + l + sovDeadlock(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DeadlockResponse) Size() (n int) {
	var l int
	_ = l
	l = m.Entry.Size()
	n += 1 + l + sovDeadlock(uint64(l))
	if m.DeadlockKeyHash != 0 {
		n += 1 + sovDeadlock(uint64(m.DeadlockKeyHash))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDeadlock(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDeadlock(x uint64) (n int) {
	return sovDeadlock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WaitForEntriesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeadlock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WaitForEntriesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WaitForEntriesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipDeadlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeadlock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WaitForEntriesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeadlock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WaitForEntriesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WaitForEntriesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeadlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDeadlock
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, WaitForEntry{})
			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeadlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeadlock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WaitForEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeadlock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WaitForEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WaitForEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txn", wireType)
			}
			m.Txn = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeadlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Txn |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WaitForTxn", wireType)
			}
			m.WaitForTxn = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeadlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WaitForTxn |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyHash", wireType)
			}
			m.KeyHash = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeadlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyHash |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDeadlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeadlock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DeadlockRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeadlock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeadlockRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeadlockRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tp", wireType)
			}
			m.Tp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeadlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Tp |= (DeadlockRequestType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeadlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDeadlock
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Entry.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeadlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeadlock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DeadlockResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeadlock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeadlockResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeadlockResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeadlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDeadlock
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Entry.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeadlockKeyHash", wireType)
			}
			m.DeadlockKeyHash = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeadlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DeadlockKeyHash |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDeadlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDeadlock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDeadlock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeadlock
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
					return 0, ErrIntOverflowDeadlock
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDeadlock
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthDeadlock
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDeadlock
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipDeadlock(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthDeadlock = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeadlock   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("deadlock.proto", fileDescriptor_deadlock_3c9646ad0923eabd) }

var fileDescriptor_deadlock_3c9646ad0923eabd = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x4e, 0xf2, 0x40,
	0x14, 0xed, 0x14, 0x3e, 0x20, 0x17, 0x02, 0xe5, 0x7e, 0xfe, 0x54, 0x34, 0x05, 0xbb, 0x22, 0x2c,
	0xc0, 0xa0, 0x0b, 0xd7, 0x88, 0x7f, 0x71, 0xd7, 0x60, 0xdc, 0x68, 0x9a, 0x02, 0x23, 0x10, 0x48,
	0xa7, 0xb6, 0x63, 0xa0, 0x6f, 0xe2, 0x7b, 0xf8, 0x12, 0x2c, 0x7d, 0x02, 0x63, 0xf0, 0x45, 0x4c,
	0xff, 0x00, 0x09, 0x62, 0xdc, 0xcd, 0xdc, 0x73, 0xe6, 0xdc, 0x73, 0xee, 0x5c, 0xc8, 0x76, 0xa9,
	0xd1, 0x1d, 0xb1, 0xce, 0xb0, 0x6a, 0xd9, 0x8c, 0x33, 0x84, 0xe8, 0x6e, 0xb5, 0x0b, 0x5b, 0x3d,
	0xd6, 0x63, 0x7e, 0xb9, 0xe6, 0x9d, 0x02, 0x86, 0xba, 0x0b, 0xdb, 0x77, 0xc6, 0x80, 0x5f, 0x30,
	0xfb, 0xdc, 0xe4, 0xf6, 0x80, 0x3a, 0x1a, 0x7d, 0x7a, 0xa6, 0x0e, 0x57, 0x35, 0xd8, 0x59, 0x05,
	0x1c, 0x8b, 0x99, 0x0e, 0xc5, 0x53, 0x48, 0xd2, 0xa0, 0x24, 0x93, 0x52, 0xac, 0x9c, 0xae, 0xcb,
	0xd5, 0x45, 0x9b, 0xea, 0xd2, 0x23, 0xb7, 0x11, 0x9f, 0xbe, 0x17, 0x05, 0x2d, 0xa2, 0xab, 0x0f,
	0x90, 0x59, 0x86, 0x51, 0x82, 0x18, 0x9f, 0x98, 0x32, 0x29, 0x91, 0x72, 0x5c, 0xf3, 0x8e, 0x58,
	0x82, 0xcc, 0xd8, 0x18, 0x70, 0xfd, 0x91, 0xd9, 0xba, 0x07, 0x89, 0x3e, 0x04, 0xe3, 0xe0, 0x55,
	0x6b, 0x62, 0xe2, 0x1e, 0xa4, 0x86, 0xd4, 0xd5, 0xfb, 0x86, 0xd3, 0x97, 0x63, 0x3e, 0x9a, 0x1c,
	0x52, 0xf7, 0xca, 0x70, 0xfa, 0xea, 0x04, 0x72, 0xcd, 0xd0, 0x48, 0x98, 0x02, 0x6b, 0x20, 0x72,
	0xcb, 0x6f, 0x90, 0xad, 0x17, 0x97, 0x6d, 0xae, 0x10, 0x5b, 0xae, 0x45, 0x35, 0x91, 0x5b, 0x78,
	0x02, 0xff, 0x3c, 0xb7, 0xae, 0xdf, 0xf9, 0xf7, 0x68, 0x01, 0x59, 0xe5, 0x20, 0x2d, 0x04, 0xc3,
	0x31, 0xcd, 0x95, 0xc8, 0x1f, 0x94, 0xb0, 0x02, 0xf9, 0x88, 0xa7, 0xcf, 0x73, 0x06, 0x53, 0xc8,
	0x45, 0xc0, 0x4d, 0x90, 0xb7, 0xd2, 0x80, 0xff, 0x6b, 0x62, 0x20, 0x40, 0xa2, 0x49, 0x39, 0xed,
	0x70, 0x49, 0x40, 0x84, 0xec, 0xd9, 0x88, 0x1a, 0xe6, 0xad, 0x15, 0xb6, 0x94, 0x08, 0xa6, 0x21,
	0x19, 0xd6, 0x24, 0xb1, 0xfe, 0x4a, 0x20, 0x15, 0x89, 0xe0, 0x3d, 0xe4, 0x2f, 0x29, 0xff, 0xfe,
	0xed, 0x78, 0xf8, 0x83, 0xf1, 0xc5, 0xae, 0x14, 0xd4, 0x4d, 0x94, 0x60, 0x1c, 0xaa, 0x80, 0xd7,
	0x91, 0x2f, 0xdc, 0xdf, 0xf0, 0x13, 0x85, 0x83, 0xf5, 0x60, 0x24, 0x53, 0x26, 0x47, 0xa4, 0x21,
	0x4d, 0x67, 0x0a, 0x79, 0x9b, 0x29, 0xe4, 0x63, 0xa6, 0x90, 0x97, 0x4f, 0x45, 0x68, 0x27, 0xfc,
	0x75, 0x3e, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x41, 0x43, 0x8b, 0xbf, 0x02, 0x03, 0x00, 0x00,
}
