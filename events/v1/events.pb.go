// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/v1/events.proto

/*
Package events is a generated protocol buffer package.

It is generated from these files:
	events/v1/events.proto

It has these top-level messages:
	PlayerPlay
*/
package events

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PlayerPlay struct {
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	QueueId   string                     `protobuf:"bytes,2,opt,name=queueId" json:"queueId,omitempty"`
	Provider  string                     `protobuf:"bytes,3,opt,name=provider" json:"provider,omitempty"`
	TrackId   string                     `protobuf:"bytes,4,opt,name=trackId" json:"trackId,omitempty"`
	UserId    string                     `protobuf:"bytes,5,opt,name=userId" json:"userId,omitempty"`
}

func (m *PlayerPlay) Reset()                    { *m = PlayerPlay{} }
func (m *PlayerPlay) String() string            { return proto.CompactTextString(m) }
func (*PlayerPlay) ProtoMessage()               {}
func (*PlayerPlay) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PlayerPlay) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *PlayerPlay) GetQueueId() string {
	if m != nil {
		return m.QueueId
	}
	return ""
}

func (m *PlayerPlay) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *PlayerPlay) GetTrackId() string {
	if m != nil {
		return m.TrackId
	}
	return ""
}

func (m *PlayerPlay) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func init() {
	proto.RegisterType((*PlayerPlay)(nil), "fm.events.v1.PlayerPlay")
}

func init() { proto.RegisterFile("events/v1/events.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x86, 0x65, 0x3e, 0x0a, 0x35, 0x4c, 0x1e, 0x2a, 0x2b, 0x0b, 0x15, 0x53, 0x17, 0x6c, 0xb5,
	0x30, 0x20, 0xb1, 0xb1, 0x65, 0x43, 0x15, 0x13, 0x5b, 0x52, 0x5f, 0x5c, 0x8b, 0xba, 0x17, 0xfc,
	0x25, 0xf1, 0xa7, 0xf8, 0x8d, 0xa8, 0x76, 0x92, 0x2e, 0xd6, 0xbd, 0x7e, 0x9e, 0x7b, 0x25, 0x9b,
	0x2e, 0x20, 0xc1, 0x31, 0x78, 0x99, 0xd6, 0xb2, 0x4c, 0xa2, 0x77, 0x18, 0x90, 0xdd, 0x77, 0x56,
	0x0c, 0x17, 0x69, 0x5d, 0x3d, 0x68, 0x44, 0x7d, 0x00, 0x99, 0x59, 0x1b, 0x3b, 0x19, 0x8c, 0x05,
	0x1f, 0x1a, 0xdb, 0x17, 0xfd, 0xf1, 0x8f, 0x50, 0xfa, 0x71, 0x68, 0x7e, 0xc1, 0x9d, 0x4e, 0xf6,
	0x4a, 0xe7, 0x93, 0xc1, 0xc9, 0x92, 0xac, 0xee, 0x36, 0x95, 0x28, 0x1d, 0x62, 0xec, 0x10, 0x9f,
	0xa3, 0xb1, 0x3d, 0xcb, 0x8c, 0xd3, 0x9b, 0x9f, 0x08, 0x11, 0x6a, 0xc5, 0x2f, 0x96, 0x64, 0x35,
	0xdf, 0x8e, 0x91, 0x55, 0xf4, 0xb6, 0x77, 0x98, 0x8c, 0x02, 0xc7, 0x2f, 0x33, 0x9a, 0xf2, 0x69,
	0x2b, 0xb8, 0x66, 0xf7, 0x5d, 0x2b, 0x7e, 0x55, 0xb6, 0x86, 0xc8, 0x16, 0x74, 0x16, 0x3d, 0xb8,
	0x5a, 0xf1, 0xeb, 0x0c, 0x86, 0xf4, 0xfe, 0xf2, 0xb5, 0xd1, 0x26, 0xec, 0x63, 0x2b, 0x76, 0x68,
	0x65, 0xd8, 0x1b, 0x6f, 0xbc, 0x47, 0x3c, 0x3e, 0x75, 0xb6, 0xbc, 0x52, 0x68, 0x94, 0xd3, 0xdf,
	0xbc, 0x95, 0xa9, 0x9d, 0x65, 0xf4, 0xfc, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xc0, 0x09, 0xf3,
	0x36, 0x01, 0x00, 0x00,
}
