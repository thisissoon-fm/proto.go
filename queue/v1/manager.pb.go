// Code generated by protoc-gen-go. DO NOT EDIT.
// source: queue/v1/manager.proto

/*
Package queue is a generated protocol buffer package.

It is generated from these files:
	queue/v1/manager.proto

It has these top-level messages:
	PutRequest
	PutResponse
	PopResponse
	NextResponse
	CountResponse
	QueueRequest
	QueueTrack
	DeleteRequest
*/
package queue

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"
import fm_providers_v1 "github.com/thisissoon-fm/proto.go/providers/v1"

import (
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

// Request for [QueueManager.Put][fm.queue.v1.QueueManager.Put]
type PutRequest struct {
	// Stream provider, for example Spotify, Google etc
	Provider fm_providers_v1.Provider `protobuf:"varint,1,opt,name=provider,enum=fm.providers.v1.Provider" json:"provider,omitempty"`
	// Provider track ID to add to the Queue
	TrackId string `protobuf:"bytes,2,opt,name=trackId" json:"trackId,omitempty"`
	// ID of the user
	UserId string `protobuf:"bytes,3,opt,name=userId" json:"userId,omitempty"`
}

func (m *PutRequest) Reset()                    { *m = PutRequest{} }
func (m *PutRequest) String() string            { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()               {}
func (*PutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PutRequest) GetProvider() fm_providers_v1.Provider {
	if m != nil {
		return m.Provider
	}
	return fm_providers_v1.Provider_Spotify
}

func (m *PutRequest) GetTrackId() string {
	if m != nil {
		return m.TrackId
	}
	return ""
}

func (m *PutRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

// Response for [QueueManager.Put][fm.queue.v1.QueueManager.Put]
type PutResponse struct {
	// The queue item created
	QueueItem *QueueTrack `protobuf:"bytes,1,opt,name=queueItem" json:"queueItem,omitempty"`
}

func (m *PutResponse) Reset()                    { *m = PutResponse{} }
func (m *PutResponse) String() string            { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()               {}
func (*PutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PutResponse) GetQueueItem() *QueueTrack {
	if m != nil {
		return m.QueueItem
	}
	return nil
}

// Response for [QueueManager.Pop][fm.queue.v1.QueueManager.Pop]
type PopResponse struct {
	// The queue item popped from the top of the queue
	QueueItem *QueueTrack `protobuf:"bytes,1,opt,name=queueItem" json:"queueItem,omitempty"`
}

func (m *PopResponse) Reset()                    { *m = PopResponse{} }
func (m *PopResponse) String() string            { return proto.CompactTextString(m) }
func (*PopResponse) ProtoMessage()               {}
func (*PopResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PopResponse) GetQueueItem() *QueueTrack {
	if m != nil {
		return m.QueueItem
	}
	return nil
}

// Response for [QueueManager.Next][fm.queue.v1.QueueManager.Next]
type NextResponse struct {
	// The next item in the queue
	QueueItem *QueueTrack `protobuf:"bytes,1,opt,name=queueItem" json:"queueItem,omitempty"`
}

func (m *NextResponse) Reset()                    { *m = NextResponse{} }
func (m *NextResponse) String() string            { return proto.CompactTextString(m) }
func (*NextResponse) ProtoMessage()               {}
func (*NextResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *NextResponse) GetQueueItem() *QueueTrack {
	if m != nil {
		return m.QueueItem
	}
	return nil
}

// Response for [QueueManager.Count][fm.queue.v1.QueueManager.Count]
type CountResponse struct {
	// Total number of items in the queue
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *CountResponse) Reset()                    { *m = CountResponse{} }
func (m *CountResponse) String() string            { return proto.CompactTextString(m) }
func (*CountResponse) ProtoMessage()               {}
func (*CountResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CountResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

// Request for [QueueManager.Queue][fm.queue.v1.QueueManager.Queue]
type QueueRequest struct {
	// Number of items to retrun
	Limit int64 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	// Offset from 0 to return items from
	Offset int64 `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
}

func (m *QueueRequest) Reset()                    { *m = QueueRequest{} }
func (m *QueueRequest) String() string            { return proto.CompactTextString(m) }
func (*QueueRequest) ProtoMessage()               {}
func (*QueueRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *QueueRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *QueueRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

// A QueueTrack is an individual item in the playlist queue
type QueueTrack struct {
	// UUIDv4 Id
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Timestamp the queue item was added
	Created *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=created" json:"created,omitempty"`
	// Provider track ID to add to the Queue
	TrackId string `protobuf:"bytes,3,opt,name=trackId" json:"trackId,omitempty"`
	// ID of the user
	UserId string `protobuf:"bytes,4,opt,name=userId" json:"userId,omitempty"`
}

func (m *QueueTrack) Reset()                    { *m = QueueTrack{} }
func (m *QueueTrack) String() string            { return proto.CompactTextString(m) }
func (*QueueTrack) ProtoMessage()               {}
func (*QueueTrack) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *QueueTrack) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *QueueTrack) GetCreated() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *QueueTrack) GetTrackId() string {
	if m != nil {
		return m.TrackId
	}
	return ""
}

func (m *QueueTrack) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

// Request for [QueueManager.Delete][fm.queue.v1.QueueManager.Delete]
type DeleteRequest struct {
	// Id of the queue item to delete
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*PutRequest)(nil), "fm.queue.v1.PutRequest")
	proto.RegisterType((*PutResponse)(nil), "fm.queue.v1.PutResponse")
	proto.RegisterType((*PopResponse)(nil), "fm.queue.v1.PopResponse")
	proto.RegisterType((*NextResponse)(nil), "fm.queue.v1.NextResponse")
	proto.RegisterType((*CountResponse)(nil), "fm.queue.v1.CountResponse")
	proto.RegisterType((*QueueRequest)(nil), "fm.queue.v1.QueueRequest")
	proto.RegisterType((*QueueTrack)(nil), "fm.queue.v1.QueueTrack")
	proto.RegisterType((*DeleteRequest)(nil), "fm.queue.v1.DeleteRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for QueueManager service

type QueueManagerClient interface {
	// Put adds a new track to the queue
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
	// Next returns the next track in the queue
	Next(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*NextResponse, error)
	// Pop takes the first item in the queue, removing it
	// from the queue and returning it
	Pop(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*PopResponse, error)
	// Count returns the totoal number of items in the queue
	Count(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*CountResponse, error)
	// Queue returns the current queue as a stream of queue items
	Queue(ctx context.Context, in *QueueRequest, opts ...grpc.CallOption) (QueueManager_QueueClient, error)
	// Delete removes a specific item from the queue
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type queueManagerClient struct {
	cc *grpc.ClientConn
}

func NewQueueManagerClient(cc *grpc.ClientConn) QueueManagerClient {
	return &queueManagerClient{cc}
}

func (c *queueManagerClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := grpc.Invoke(ctx, "/fm.queue.v1.QueueManager/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueManagerClient) Next(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*NextResponse, error) {
	out := new(NextResponse)
	err := grpc.Invoke(ctx, "/fm.queue.v1.QueueManager/Next", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueManagerClient) Pop(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*PopResponse, error) {
	out := new(PopResponse)
	err := grpc.Invoke(ctx, "/fm.queue.v1.QueueManager/Pop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueManagerClient) Count(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*CountResponse, error) {
	out := new(CountResponse)
	err := grpc.Invoke(ctx, "/fm.queue.v1.QueueManager/Count", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueManagerClient) Queue(ctx context.Context, in *QueueRequest, opts ...grpc.CallOption) (QueueManager_QueueClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_QueueManager_serviceDesc.Streams[0], c.cc, "/fm.queue.v1.QueueManager/Queue", opts...)
	if err != nil {
		return nil, err
	}
	x := &queueManagerQueueClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type QueueManager_QueueClient interface {
	Recv() (*QueueTrack, error)
	grpc.ClientStream
}

type queueManagerQueueClient struct {
	grpc.ClientStream
}

func (x *queueManagerQueueClient) Recv() (*QueueTrack, error) {
	m := new(QueueTrack)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *queueManagerClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/fm.queue.v1.QueueManager/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for QueueManager service

type QueueManagerServer interface {
	// Put adds a new track to the queue
	Put(context.Context, *PutRequest) (*PutResponse, error)
	// Next returns the next track in the queue
	Next(context.Context, *google_protobuf.Empty) (*NextResponse, error)
	// Pop takes the first item in the queue, removing it
	// from the queue and returning it
	Pop(context.Context, *google_protobuf.Empty) (*PopResponse, error)
	// Count returns the totoal number of items in the queue
	Count(context.Context, *google_protobuf.Empty) (*CountResponse, error)
	// Queue returns the current queue as a stream of queue items
	Queue(*QueueRequest, QueueManager_QueueServer) error
	// Delete removes a specific item from the queue
	Delete(context.Context, *DeleteRequest) (*google_protobuf.Empty, error)
}

func RegisterQueueManagerServer(s *grpc.Server, srv QueueManagerServer) {
	s.RegisterService(&_QueueManager_serviceDesc, srv)
}

func _QueueManager_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueManagerServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fm.queue.v1.QueueManager/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueManagerServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueManager_Next_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueManagerServer).Next(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fm.queue.v1.QueueManager/Next",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueManagerServer).Next(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueManager_Pop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueManagerServer).Pop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fm.queue.v1.QueueManager/Pop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueManagerServer).Pop(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueManager_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueManagerServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fm.queue.v1.QueueManager/Count",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueManagerServer).Count(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueManager_Queue_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueueRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueueManagerServer).Queue(m, &queueManagerQueueServer{stream})
}

type QueueManager_QueueServer interface {
	Send(*QueueTrack) error
	grpc.ServerStream
}

type queueManagerQueueServer struct {
	grpc.ServerStream
}

func (x *queueManagerQueueServer) Send(m *QueueTrack) error {
	return x.ServerStream.SendMsg(m)
}

func _QueueManager_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueManagerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fm.queue.v1.QueueManager/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueManagerServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueueManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fm.queue.v1.QueueManager",
	HandlerType: (*QueueManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _QueueManager_Put_Handler,
		},
		{
			MethodName: "Next",
			Handler:    _QueueManager_Next_Handler,
		},
		{
			MethodName: "Pop",
			Handler:    _QueueManager_Pop_Handler,
		},
		{
			MethodName: "Count",
			Handler:    _QueueManager_Count_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _QueueManager_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Queue",
			Handler:       _QueueManager_Queue_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "queue/v1/manager.proto",
}

func init() { proto.RegisterFile("queue/v1/manager.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0x4e, 0x96, 0xb5, 0xa3, 0xaf, 0xdb, 0x0e, 0xd6, 0x54, 0xd2, 0x80, 0xb4, 0xc9, 0x12, 0xd2,
	0x2e, 0x38, 0x6b, 0x60, 0x07, 0x18, 0x3f, 0x24, 0xd8, 0x0e, 0x3d, 0x80, 0x4a, 0xb4, 0x13, 0xb7,
	0xb4, 0x71, 0xb2, 0x88, 0xba, 0xce, 0x62, 0xbb, 0x82, 0x2b, 0xe2, 0x0f, 0x47, 0xb1, 0x93, 0xa6,
	0x69, 0x95, 0xd3, 0x6e, 0x7e, 0x79, 0xef, 0x7b, 0xcf, 0xef, 0xfb, 0xbe, 0x18, 0x46, 0x8f, 0x8a,
	0x2a, 0xea, 0xaf, 0x27, 0x3e, 0x8b, 0x56, 0x51, 0x4a, 0x0b, 0x92, 0x17, 0x5c, 0x72, 0x34, 0x4c,
	0x18, 0xd1, 0x29, 0xb2, 0x9e, 0x78, 0x2f, 0x52, 0xce, 0xd3, 0x25, 0xf5, 0x75, 0x6a, 0xae, 0x12,
	0x9f, 0xb2, 0x5c, 0xfe, 0x31, 0x95, 0xde, 0xf9, 0x6e, 0x52, 0x66, 0x8c, 0x0a, 0x19, 0xb1, 0xbc,
	0x2a, 0x78, 0x99, 0x17, 0x7c, 0x9d, 0xc5, 0xb4, 0x10, 0xe5, 0x98, 0x4d, 0x60, 0xb2, 0x58, 0x01,
	0xcc, 0x94, 0x0c, 0xe9, 0xa3, 0xa2, 0x42, 0xa2, 0x6b, 0x78, 0x56, 0x17, 0xb8, 0xf6, 0x85, 0x7d,
	0x79, 0x1a, 0x8c, 0x49, 0xc2, 0x48, 0x03, 0x5a, 0x4f, 0xc8, 0xac, 0x0a, 0xc2, 0x4d, 0x29, 0x72,
	0xe1, 0x48, 0x16, 0xd1, 0xe2, 0xd7, 0x34, 0x76, 0x0f, 0x2e, 0xec, 0xcb, 0x41, 0x58, 0x87, 0x68,
	0x04, 0x7d, 0x25, 0x68, 0x31, 0x8d, 0x5d, 0x47, 0x27, 0xaa, 0x08, 0xdf, 0xc2, 0x50, 0x8f, 0x15,
	0x39, 0x5f, 0x09, 0x8a, 0xae, 0x61, 0xa0, 0xb7, 0x9d, 0x4a, 0xca, 0xf4, 0xe0, 0x61, 0xf0, 0x9c,
	0x6c, 0x51, 0x40, 0x7e, 0x94, 0x87, 0xfb, 0xb2, 0x69, 0xd8, 0x54, 0xea, 0x2e, 0x3c, 0x7f, 0x6a,
	0x97, 0x3b, 0x38, 0xfe, 0x4e, 0x7f, 0x3f, 0xf9, 0x32, 0xaf, 0xe0, 0xe4, 0x2b, 0x57, 0xab, 0xa6,
	0xcf, 0x19, 0xf4, 0x16, 0xe5, 0x07, 0xdd, 0xc3, 0x09, 0x4d, 0x80, 0x3f, 0xc0, 0xb1, 0xc6, 0xd7,
	0x94, 0x9f, 0x41, 0x6f, 0x99, 0xb1, 0x6c, 0x53, 0xa5, 0x83, 0x92, 0x37, 0x9e, 0x24, 0x82, 0x4a,
	0x4d, 0xa8, 0x13, 0x56, 0x11, 0xfe, 0x67, 0x03, 0x34, 0xe3, 0xd1, 0x29, 0x1c, 0x64, 0xb1, 0x46,
	0x0e, 0xc2, 0x83, 0x2c, 0x46, 0x6f, 0xe1, 0x68, 0x51, 0xd0, 0x48, 0x52, 0x23, 0xc4, 0x30, 0xf0,
	0x88, 0xb1, 0x07, 0xa9, 0xed, 0x41, 0xee, 0x6b, 0x7b, 0x84, 0x75, 0xe9, 0xb6, 0x7c, 0x4e, 0x97,
	0x7c, 0x87, 0x2d, 0xf9, 0xce, 0xe1, 0xe4, 0x96, 0x2e, 0xa9, 0xdc, 0x6c, 0xb1, 0x73, 0x91, 0xe0,
	0xaf, 0x53, 0xad, 0xf9, 0xcd, 0xd8, 0x1a, 0xbd, 0x07, 0x67, 0xa6, 0x24, 0x6a, 0x13, 0xd9, 0x38,
	0xcf, 0x73, 0xf7, 0x13, 0x86, 0x46, 0x6c, 0xa1, 0x1b, 0x38, 0x2c, 0x05, 0x42, 0xa3, 0xbd, 0x65,
	0xee, 0xca, 0x1f, 0xc1, 0x1b, 0xb7, 0xb0, 0xdb, 0x5a, 0x62, 0x0b, 0xbd, 0x03, 0x67, 0xc6, 0xf3,
	0x4e, 0xec, 0xce, 0xdc, 0xc6, 0x4d, 0xd8, 0x42, 0x1f, 0xa1, 0xa7, 0x15, 0xed, 0x04, 0x7b, 0x2d,
	0x70, 0x4b, 0x7d, 0x6c, 0xa1, 0xcf, 0xd0, 0xd3, 0x14, 0xa0, 0xf1, 0xbe, 0x7b, 0xea, 0xb5, 0xbb,
	0x8c, 0x85, 0xad, 0x2b, 0x1b, 0x7d, 0x82, 0xbe, 0x61, 0x19, 0xb5, 0x07, 0xb5, 0xa8, 0xf7, 0x3a,
	0x2e, 0x87, 0xad, 0x2f, 0xc1, 0xcf, 0xab, 0x34, 0x93, 0x0f, 0x6a, 0x4e, 0x16, 0x9c, 0xf9, 0xf2,
	0x21, 0x13, 0x99, 0x10, 0x9c, 0xaf, 0x5e, 0x27, 0xcc, 0x3c, 0x17, 0x24, 0xe5, 0x7e, 0xfd, 0x00,
	0xdd, 0xe8, 0xc3, 0xbc, 0xaf, 0x13, 0x6f, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x04, 0xbc,
	0xd2, 0x99, 0x04, 0x00, 0x00,
}
