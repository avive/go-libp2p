// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: p2p.proto

/*
Package protocols_p2p is a generated protocol buffer package.

It is generated from these files:
	p2p.proto

It has these top-level messages:
	MessageData
	PingRequest
	PingResponse
	EchoRequest
	EchoResponse
*/
package protocols_p2p

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// designed to be shared between all app protocols
type MessageData struct {
	// shared between all requests
	ClientVersion string `protobuf:"bytes,1,opt,name=clientVersion,proto3" json:"clientVersion,omitempty"`
	Timestamp     int64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Id            string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Gossip        bool   `protobuf:"varint,4,opt,name=gossip,proto3" json:"gossip,omitempty"`
	NodeId        string `protobuf:"bytes,5,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	NodePubKey    []byte `protobuf:"bytes,6,opt,name=nodePubKey,proto3" json:"nodePubKey,omitempty"`
	Sign          string `protobuf:"bytes,7,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (m *MessageData) Reset()                    { *m = MessageData{} }
func (m *MessageData) String() string            { return proto.CompactTextString(m) }
func (*MessageData) ProtoMessage()               {}
func (*MessageData) Descriptor() ([]byte, []int) { return fileDescriptorP2P, []int{0} }

func (m *MessageData) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

func (m *MessageData) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *MessageData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MessageData) GetGossip() bool {
	if m != nil {
		return m.Gossip
	}
	return false
}

func (m *MessageData) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *MessageData) GetNodePubKey() []byte {
	if m != nil {
		return m.NodePubKey
	}
	return nil
}

func (m *MessageData) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

// a protocol define a set of reuqest and responses
type PingRequest struct {
	MessageData *MessageData `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	// method specific data
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptorP2P, []int{1} }

func (m *PingRequest) GetMessageData() *MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *PingRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PingResponse struct {
	MessageData *MessageData `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	// response specific data
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptorP2P, []int{2} }

func (m *PingResponse) GetMessageData() *MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *PingResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// a protocol define a set of reuqest and responses
type EchoRequest struct {
	MessageData *MessageData `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	// method specific data
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *EchoRequest) Reset()                    { *m = EchoRequest{} }
func (m *EchoRequest) String() string            { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()               {}
func (*EchoRequest) Descriptor() ([]byte, []int) { return fileDescriptorP2P, []int{3} }

func (m *EchoRequest) GetMessageData() *MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *EchoRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type EchoResponse struct {
	MessageData *MessageData `protobuf:"bytes,1,opt,name=messageData" json:"messageData,omitempty"`
	// response specific data
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *EchoResponse) Reset()                    { *m = EchoResponse{} }
func (m *EchoResponse) String() string            { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()               {}
func (*EchoResponse) Descriptor() ([]byte, []int) { return fileDescriptorP2P, []int{4} }

func (m *EchoResponse) GetMessageData() *MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (m *EchoResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*MessageData)(nil), "protocols.p2p.MessageData")
	proto.RegisterType((*PingRequest)(nil), "protocols.p2p.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "protocols.p2p.PingResponse")
	proto.RegisterType((*EchoRequest)(nil), "protocols.p2p.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "protocols.p2p.EchoResponse")
}

func init() { proto.RegisterFile("p2p.proto", fileDescriptorP2P) }

var fileDescriptorP2P = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x8f, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xe5, 0xb6, 0xa4, 0xe4, 0xdc, 0x32, 0xdc, 0x80, 0x2c, 0x84, 0x50, 0x14, 0x31, 0x64,
	0xca, 0x10, 0x56, 0x46, 0x18, 0x10, 0x42, 0xaa, 0x3c, 0xb0, 0xa7, 0xc9, 0x11, 0x2c, 0x35, 0xb6,
	0xe9, 0xb9, 0x03, 0x0f, 0xc8, 0x7b, 0xa1, 0xba, 0x41, 0x4d, 0x1f, 0xa0, 0x4c, 0xbe, 0xff, 0xd3,
	0x9d, 0x7f, 0x7d, 0x90, 0xfa, 0xca, 0x97, 0x7e, 0xeb, 0x82, 0xc3, 0x65, 0x7c, 0x1a, 0xb7, 0xe1,
	0xd2, 0x57, 0x3e, 0xff, 0x11, 0x20, 0xdf, 0x88, 0xb9, 0xee, 0xe8, 0xa9, 0x0e, 0x35, 0xde, 0xc3,
	0xb2, 0xd9, 0x18, 0xb2, 0xe1, 0x9d, 0xb6, 0x6c, 0x9c, 0x55, 0x22, 0x13, 0x45, 0xaa, 0x4f, 0x21,
	0xde, 0x42, 0x1a, 0x4c, 0x4f, 0x1c, 0xea, 0xde, 0xab, 0x49, 0x26, 0x8a, 0xa9, 0x3e, 0x02, 0xbc,
	0x82, 0x89, 0x69, 0xd5, 0x34, 0x1e, 0x4e, 0x4c, 0x8b, 0xd7, 0x90, 0x74, 0x8e, 0xd9, 0x78, 0x35,
	0xcb, 0x44, 0x71, 0xa9, 0x87, 0xb4, 0xe7, 0xd6, 0xb5, 0xf4, 0xd2, 0xaa, 0x8b, 0xb8, 0x3b, 0x24,
	0xbc, 0x03, 0xd8, 0x4f, 0xab, 0xdd, 0xfa, 0x95, 0xbe, 0x55, 0x92, 0x89, 0x62, 0xa1, 0x47, 0x04,
	0x11, 0x66, 0x6c, 0x3a, 0xab, 0xe6, 0xf1, 0x2a, 0xce, 0x39, 0x81, 0x5c, 0x19, 0xdb, 0x69, 0xfa,
	0xda, 0x11, 0x07, 0x7c, 0x04, 0xd9, 0x1f, 0xad, 0xa2, 0x84, 0xac, 0x6e, 0xca, 0x13, 0xf7, 0x72,
	0xe4, 0xad, 0xc7, 0xeb, 0xa8, 0x60, 0x3e, 0xc4, 0x28, 0x97, 0xea, 0xbf, 0x98, 0x7f, 0xc0, 0xe2,
	0x50, 0xc3, 0xde, 0x59, 0xa6, 0xb3, 0xf5, 0x10, 0xc8, 0xe7, 0xe6, 0xd3, 0xfd, 0x83, 0xce, 0xa1,
	0xe6, 0xbc, 0x3a, 0xeb, 0x24, 0xfe, 0xf0, 0xf0, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x96, 0x51, 0x39,
	0x88, 0x88, 0x02, 0x00, 0x00,
}
