// Code generated by protoc-gen-go.
// source: raft.proto
// DO NOT EDIT!

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	raft.proto

It has these top-level messages:
	LogEntry
	AppendEntriesRequest
	AppendEntriesResponse
	RequestVoteRequest
	RequestVoteResponse
	SnapshotRecoveryRequest
	SnapshotRecoveryResponse
	SnapshotRequest
	SnapshotResponse
*/
package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type LogEntry struct {
	Index       uint64 `protobuf:"varint,1,opt,name=Index" json:"Index,omitempty"`
	Term        uint64 `protobuf:"varint,2,opt,name=Term" json:"Term,omitempty"`
	CommandName string `protobuf:"bytes,3,opt,name=CommandName" json:"CommandName,omitempty"`
	Command     []byte `protobuf:"bytes,4,opt,name=Command,proto3" json:"Command,omitempty"`
}

func (m *LogEntry) Reset()                    { *m = LogEntry{} }
func (m *LogEntry) String() string            { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()               {}
func (*LogEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LogEntry) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *LogEntry) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *LogEntry) GetCommandName() string {
	if m != nil {
		return m.CommandName
	}
	return ""
}

func (m *LogEntry) GetCommand() []byte {
	if m != nil {
		return m.Command
	}
	return nil
}

type AppendEntriesRequest struct {
	Term         uint64      `protobuf:"varint,1,opt,name=Term" json:"Term,omitempty"`
	PrevLogIndex uint64      `protobuf:"varint,2,opt,name=PrevLogIndex" json:"PrevLogIndex,omitempty"`
	PrevLogTerm  uint64      `protobuf:"varint,3,opt,name=PrevLogTerm" json:"PrevLogTerm,omitempty"`
	CommitIndex  uint64      `protobuf:"varint,4,opt,name=CommitIndex" json:"CommitIndex,omitempty"`
	LeaderName   string      `protobuf:"bytes,5,opt,name=LeaderName" json:"LeaderName,omitempty"`
	Entries      []*LogEntry `protobuf:"bytes,6,rep,name=Entries" json:"Entries,omitempty"`
}

func (m *AppendEntriesRequest) Reset()                    { *m = AppendEntriesRequest{} }
func (m *AppendEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*AppendEntriesRequest) ProtoMessage()               {}
func (*AppendEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AppendEntriesRequest) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *AppendEntriesRequest) GetPrevLogIndex() uint64 {
	if m != nil {
		return m.PrevLogIndex
	}
	return 0
}

func (m *AppendEntriesRequest) GetPrevLogTerm() uint64 {
	if m != nil {
		return m.PrevLogTerm
	}
	return 0
}

func (m *AppendEntriesRequest) GetCommitIndex() uint64 {
	if m != nil {
		return m.CommitIndex
	}
	return 0
}

func (m *AppendEntriesRequest) GetLeaderName() string {
	if m != nil {
		return m.LeaderName
	}
	return ""
}

func (m *AppendEntriesRequest) GetEntries() []*LogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type AppendEntriesResponse struct {
	Term        uint64 `protobuf:"varint,1,opt,name=Term" json:"Term,omitempty"`
	Index       uint64 `protobuf:"varint,2,opt,name=Index" json:"Index,omitempty"`
	CommitIndex uint64 `protobuf:"varint,3,opt,name=CommitIndex" json:"CommitIndex,omitempty"`
	Success     bool   `protobuf:"varint,4,opt,name=Success" json:"Success,omitempty"`
}

func (m *AppendEntriesResponse) Reset()                    { *m = AppendEntriesResponse{} }
func (m *AppendEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*AppendEntriesResponse) ProtoMessage()               {}
func (*AppendEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AppendEntriesResponse) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *AppendEntriesResponse) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *AppendEntriesResponse) GetCommitIndex() uint64 {
	if m != nil {
		return m.CommitIndex
	}
	return 0
}

func (m *AppendEntriesResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type RequestVoteRequest struct {
	Term          uint64 `protobuf:"varint,1,opt,name=Term" json:"Term,omitempty"`
	LastLogIndex  uint64 `protobuf:"varint,2,opt,name=LastLogIndex" json:"LastLogIndex,omitempty"`
	LastLogTerm   uint64 `protobuf:"varint,3,opt,name=LastLogTerm" json:"LastLogTerm,omitempty"`
	CandidateName string `protobuf:"bytes,4,opt,name=CandidateName" json:"CandidateName,omitempty"`
}

func (m *RequestVoteRequest) Reset()                    { *m = RequestVoteRequest{} }
func (m *RequestVoteRequest) String() string            { return proto.CompactTextString(m) }
func (*RequestVoteRequest) ProtoMessage()               {}
func (*RequestVoteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RequestVoteRequest) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RequestVoteRequest) GetLastLogIndex() uint64 {
	if m != nil {
		return m.LastLogIndex
	}
	return 0
}

func (m *RequestVoteRequest) GetLastLogTerm() uint64 {
	if m != nil {
		return m.LastLogTerm
	}
	return 0
}

func (m *RequestVoteRequest) GetCandidateName() string {
	if m != nil {
		return m.CandidateName
	}
	return ""
}

type RequestVoteResponse struct {
	Term        uint64 `protobuf:"varint,1,opt,name=Term" json:"Term,omitempty"`
	VoteGranted bool   `protobuf:"varint,2,opt,name=VoteGranted" json:"VoteGranted,omitempty"`
}

func (m *RequestVoteResponse) Reset()                    { *m = RequestVoteResponse{} }
func (m *RequestVoteResponse) String() string            { return proto.CompactTextString(m) }
func (*RequestVoteResponse) ProtoMessage()               {}
func (*RequestVoteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RequestVoteResponse) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RequestVoteResponse) GetVoteGranted() bool {
	if m != nil {
		return m.VoteGranted
	}
	return false
}

type SnapshotRecoveryRequest struct {
	LeaderName string                          `protobuf:"bytes,1,opt,name=LeaderName" json:"LeaderName,omitempty"`
	LastIndex  uint64                          `protobuf:"varint,2,opt,name=LastIndex" json:"LastIndex,omitempty"`
	LastTerm   uint64                          `protobuf:"varint,3,opt,name=LastTerm" json:"LastTerm,omitempty"`
	Peers      []*SnapshotRecoveryRequest_Peer `protobuf:"bytes,4,rep,name=Peers" json:"Peers,omitempty"`
	State      []byte                          `protobuf:"bytes,5,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *SnapshotRecoveryRequest) Reset()                    { *m = SnapshotRecoveryRequest{} }
func (m *SnapshotRecoveryRequest) String() string            { return proto.CompactTextString(m) }
func (*SnapshotRecoveryRequest) ProtoMessage()               {}
func (*SnapshotRecoveryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SnapshotRecoveryRequest) GetLeaderName() string {
	if m != nil {
		return m.LeaderName
	}
	return ""
}

func (m *SnapshotRecoveryRequest) GetLastIndex() uint64 {
	if m != nil {
		return m.LastIndex
	}
	return 0
}

func (m *SnapshotRecoveryRequest) GetLastTerm() uint64 {
	if m != nil {
		return m.LastTerm
	}
	return 0
}

func (m *SnapshotRecoveryRequest) GetPeers() []*SnapshotRecoveryRequest_Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

func (m *SnapshotRecoveryRequest) GetState() []byte {
	if m != nil {
		return m.State
	}
	return nil
}

type SnapshotRecoveryRequest_Peer struct {
	Name             string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	ConnectionString string `protobuf:"bytes,2,opt,name=ConnectionString" json:"ConnectionString,omitempty"`
}

func (m *SnapshotRecoveryRequest_Peer) Reset()                    { *m = SnapshotRecoveryRequest_Peer{} }
func (m *SnapshotRecoveryRequest_Peer) String() string            { return proto.CompactTextString(m) }
func (*SnapshotRecoveryRequest_Peer) ProtoMessage()               {}
func (*SnapshotRecoveryRequest_Peer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *SnapshotRecoveryRequest_Peer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SnapshotRecoveryRequest_Peer) GetConnectionString() string {
	if m != nil {
		return m.ConnectionString
	}
	return ""
}

type SnapshotRecoveryResponse struct {
	Term        uint64 `protobuf:"varint,1,opt,name=Term" json:"Term,omitempty"`
	Success     bool   `protobuf:"varint,2,opt,name=Success" json:"Success,omitempty"`
	CommitIndex uint64 `protobuf:"varint,3,opt,name=CommitIndex" json:"CommitIndex,omitempty"`
}

func (m *SnapshotRecoveryResponse) Reset()                    { *m = SnapshotRecoveryResponse{} }
func (m *SnapshotRecoveryResponse) String() string            { return proto.CompactTextString(m) }
func (*SnapshotRecoveryResponse) ProtoMessage()               {}
func (*SnapshotRecoveryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SnapshotRecoveryResponse) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *SnapshotRecoveryResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *SnapshotRecoveryResponse) GetCommitIndex() uint64 {
	if m != nil {
		return m.CommitIndex
	}
	return 0
}

type SnapshotRequest struct {
	LeaderName string `protobuf:"bytes,1,opt,name=LeaderName" json:"LeaderName,omitempty"`
	LastIndex  uint64 `protobuf:"varint,2,opt,name=LastIndex" json:"LastIndex,omitempty"`
	LastTerm   uint64 `protobuf:"varint,3,opt,name=LastTerm" json:"LastTerm,omitempty"`
}

func (m *SnapshotRequest) Reset()                    { *m = SnapshotRequest{} }
func (m *SnapshotRequest) String() string            { return proto.CompactTextString(m) }
func (*SnapshotRequest) ProtoMessage()               {}
func (*SnapshotRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SnapshotRequest) GetLeaderName() string {
	if m != nil {
		return m.LeaderName
	}
	return ""
}

func (m *SnapshotRequest) GetLastIndex() uint64 {
	if m != nil {
		return m.LastIndex
	}
	return 0
}

func (m *SnapshotRequest) GetLastTerm() uint64 {
	if m != nil {
		return m.LastTerm
	}
	return 0
}

type SnapshotResponse struct {
	Success bool `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
}

func (m *SnapshotResponse) Reset()                    { *m = SnapshotResponse{} }
func (m *SnapshotResponse) String() string            { return proto.CompactTextString(m) }
func (*SnapshotResponse) ProtoMessage()               {}
func (*SnapshotResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *SnapshotResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*LogEntry)(nil), "protobuf.LogEntry")
	proto.RegisterType((*AppendEntriesRequest)(nil), "protobuf.AppendEntriesRequest")
	proto.RegisterType((*AppendEntriesResponse)(nil), "protobuf.AppendEntriesResponse")
	proto.RegisterType((*RequestVoteRequest)(nil), "protobuf.RequestVoteRequest")
	proto.RegisterType((*RequestVoteResponse)(nil), "protobuf.RequestVoteResponse")
	proto.RegisterType((*SnapshotRecoveryRequest)(nil), "protobuf.SnapshotRecoveryRequest")
	proto.RegisterType((*SnapshotRecoveryRequest_Peer)(nil), "protobuf.SnapshotRecoveryRequest.Peer")
	proto.RegisterType((*SnapshotRecoveryResponse)(nil), "protobuf.SnapshotRecoveryResponse")
	proto.RegisterType((*SnapshotRequest)(nil), "protobuf.SnapshotRequest")
	proto.RegisterType((*SnapshotResponse)(nil), "protobuf.SnapshotResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Raft service

type RaftClient interface {
	OnSendVoteRequest(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error)
	OnSendAppendEntriesRequest(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error)
	OnSendSnapshotRequest(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (*SnapshotResponse, error)
	OnSendSnapshotRecoveryRequest(ctx context.Context, in *SnapshotRecoveryRequest, opts ...grpc.CallOption) (*SnapshotRecoveryResponse, error)
}

type raftClient struct {
	cc *grpc.ClientConn
}

func NewRaftClient(cc *grpc.ClientConn) RaftClient {
	return &raftClient{cc}
}

func (c *raftClient) OnSendVoteRequest(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error) {
	out := new(RequestVoteResponse)
	err := grpc.Invoke(ctx, "/protobuf.Raft/OnSendVoteRequest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftClient) OnSendAppendEntriesRequest(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error) {
	out := new(AppendEntriesResponse)
	err := grpc.Invoke(ctx, "/protobuf.Raft/OnSendAppendEntriesRequest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftClient) OnSendSnapshotRequest(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (*SnapshotResponse, error) {
	out := new(SnapshotResponse)
	err := grpc.Invoke(ctx, "/protobuf.Raft/OnSendSnapshotRequest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftClient) OnSendSnapshotRecoveryRequest(ctx context.Context, in *SnapshotRecoveryRequest, opts ...grpc.CallOption) (*SnapshotRecoveryResponse, error) {
	out := new(SnapshotRecoveryResponse)
	err := grpc.Invoke(ctx, "/protobuf.Raft/OnSendSnapshotRecoveryRequest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Raft service

type RaftServer interface {
	OnSendVoteRequest(context.Context, *RequestVoteRequest) (*RequestVoteResponse, error)
	OnSendAppendEntriesRequest(context.Context, *AppendEntriesRequest) (*AppendEntriesResponse, error)
	OnSendSnapshotRequest(context.Context, *SnapshotRequest) (*SnapshotResponse, error)
	OnSendSnapshotRecoveryRequest(context.Context, *SnapshotRecoveryRequest) (*SnapshotRecoveryResponse, error)
}

func RegisterRaftServer(s *grpc.Server, srv RaftServer) {
	s.RegisterService(&_Raft_serviceDesc, srv)
}

func _Raft_OnSendVoteRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).OnSendVoteRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Raft/OnSendVoteRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).OnSendVoteRequest(ctx, req.(*RequestVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Raft_OnSendAppendEntriesRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).OnSendAppendEntriesRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Raft/OnSendAppendEntriesRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).OnSendAppendEntriesRequest(ctx, req.(*AppendEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Raft_OnSendSnapshotRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).OnSendSnapshotRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Raft/OnSendSnapshotRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).OnSendSnapshotRequest(ctx, req.(*SnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Raft_OnSendSnapshotRecoveryRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SnapshotRecoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).OnSendSnapshotRecoveryRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Raft/OnSendSnapshotRecoveryRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).OnSendSnapshotRecoveryRequest(ctx, req.(*SnapshotRecoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Raft_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Raft",
	HandlerType: (*RaftServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnSendVoteRequest",
			Handler:    _Raft_OnSendVoteRequest_Handler,
		},
		{
			MethodName: "OnSendAppendEntriesRequest",
			Handler:    _Raft_OnSendAppendEntriesRequest_Handler,
		},
		{
			MethodName: "OnSendSnapshotRequest",
			Handler:    _Raft_OnSendSnapshotRequest_Handler,
		},
		{
			MethodName: "OnSendSnapshotRecoveryRequest",
			Handler:    _Raft_OnSendSnapshotRecoveryRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "raft.proto",
}

func init() { proto.RegisterFile("raft.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 571 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x89, 0xdb, 0x24, 0x93, 0x20, 0xca, 0xd2, 0x0a, 0x63, 0xb5, 0xc5, 0xac, 0x10, 0x8a,
	0x50, 0x95, 0x43, 0xb9, 0x72, 0x41, 0x11, 0x20, 0x44, 0x04, 0xd1, 0x06, 0x71, 0x45, 0xdb, 0x78,
	0x92, 0x06, 0x94, 0x5d, 0x63, 0x6f, 0x2a, 0x7a, 0xe0, 0x37, 0x70, 0xe1, 0x07, 0xf2, 0x2f, 0xb8,
	0xa2, 0xdd, 0xb5, 0xe3, 0x75, 0x3e, 0x4f, 0x9c, 0xe2, 0x79, 0x33, 0x3b, 0xf3, 0xde, 0xdb, 0xd9,
	0x00, 0xa4, 0x7c, 0xa2, 0x7a, 0x49, 0x2a, 0x95, 0x24, 0x4d, 0xf3, 0x73, 0xb5, 0x98, 0xd0, 0x04,
	0x9a, 0x03, 0x39, 0x7d, 0x2d, 0x54, 0x7a, 0x4b, 0x8e, 0xe1, 0xe0, 0x9d, 0x88, 0xf1, 0x47, 0xe0,
	0x45, 0x5e, 0xd7, 0x67, 0x36, 0x20, 0x04, 0xfc, 0x4f, 0x98, 0xce, 0x83, 0x9a, 0x01, 0xcd, 0x37,
	0x89, 0xa0, 0xdd, 0x97, 0xf3, 0x39, 0x17, 0xf1, 0x07, 0x3e, 0xc7, 0xa0, 0x1e, 0x79, 0xdd, 0x16,
	0x73, 0x21, 0x12, 0x40, 0x23, 0x0f, 0x03, 0x3f, 0xf2, 0xba, 0x1d, 0x56, 0x84, 0xf4, 0x8f, 0x07,
	0xc7, 0xaf, 0x92, 0x04, 0x45, 0xac, 0xa7, 0xce, 0x30, 0x63, 0xf8, 0x7d, 0x81, 0x99, 0x5a, 0x0e,
	0xf2, 0x9c, 0x41, 0x14, 0x3a, 0xc3, 0x14, 0x6f, 0x06, 0x72, 0x6a, 0x99, 0x59, 0x12, 0x15, 0x4c,
	0x93, 0xc9, 0x63, 0x73, 0xbc, 0x6e, 0x4a, 0x5c, 0xa8, 0xa0, 0x3b, 0x53, 0xb6, 0x89, 0x6f, 0x2b,
	0x1c, 0x88, 0x9c, 0x03, 0x0c, 0x90, 0xc7, 0x98, 0x1a, 0x3d, 0x07, 0x46, 0x8f, 0x83, 0x90, 0x0b,
	0x68, 0xe4, 0x6c, 0x83, 0xc3, 0xa8, 0xde, 0x6d, 0x5f, 0x92, 0x5e, 0x61, 0x61, 0xaf, 0xf0, 0x8f,
	0x15, 0x25, 0xf4, 0x27, 0x9c, 0xac, 0x28, 0xcc, 0x12, 0x29, 0x32, 0xdc, 0x28, 0x71, 0xe9, 0x7a,
	0xcd, 0x75, 0x7d, 0x85, 0x72, 0x7d, 0x9d, 0x72, 0x00, 0x8d, 0xd1, 0x62, 0x3c, 0xc6, 0x2c, 0x33,
	0x82, 0x9a, 0xac, 0x08, 0xe9, 0x6f, 0x0f, 0x48, 0x6e, 0xea, 0x67, 0xa9, 0x70, 0x8f, 0xbf, 0x03,
	0x9e, 0xa9, 0x55, 0x7f, 0x5d, 0x4c, 0x53, 0xc9, 0x63, 0xd7, 0x5f, 0x07, 0x22, 0x4f, 0xe1, 0x6e,
	0x9f, 0x8b, 0x78, 0x16, 0x73, 0x85, 0xc6, 0x40, 0xdf, 0x18, 0x58, 0x05, 0xe9, 0x7b, 0x78, 0x50,
	0x61, 0xb5, 0xc3, 0x93, 0x08, 0xda, 0xba, 0xe6, 0x6d, 0xca, 0x85, 0xc2, 0xd8, 0xb0, 0x6a, 0x32,
	0x17, 0xa2, 0xbf, 0x6a, 0xf0, 0x70, 0x24, 0x78, 0x92, 0x5d, 0x4b, 0xc5, 0x70, 0x2c, 0x6f, 0x30,
	0xbd, 0x2d, 0x84, 0x56, 0x2f, 0xd3, 0x5b, 0xbb, 0xcc, 0x53, 0x68, 0x69, 0xf6, 0xae, 0xe2, 0x12,
	0x20, 0x21, 0x34, 0x75, 0xe0, 0x68, 0x5d, 0xc6, 0xe4, 0x25, 0x1c, 0x0c, 0x11, 0x53, 0xed, 0xb8,
	0x5e, 0x82, 0x67, 0xe5, 0x12, 0x6c, 0xe1, 0xd2, 0xd3, 0xe5, 0xcc, 0x1e, 0xd2, 0x37, 0x3d, 0x52,
	0x5c, 0xd9, 0xfd, 0xea, 0x30, 0x1b, 0x84, 0x6f, 0xc0, 0xd7, 0x69, 0xed, 0x83, 0xc3, 0xd7, 0x7c,
	0x93, 0xe7, 0x70, 0xd4, 0x97, 0x42, 0xe0, 0x58, 0xcd, 0xa4, 0x18, 0xa9, 0x74, 0x26, 0xa6, 0x86,
	0x70, 0x8b, 0xad, 0xe1, 0xf4, 0x2b, 0x04, 0xeb, 0x24, 0x76, 0x78, 0xec, 0xec, 0x4f, 0xad, 0xb2,
	0x3f, 0xfb, 0x77, 0x8f, 0x7e, 0x83, 0x7b, 0xe5, 0xac, 0xff, 0x6c, 0x3a, 0xbd, 0x80, 0xa3, 0x72,
	0x58, 0x2e, 0xc8, 0x21, 0xef, 0x55, 0xc8, 0x5f, 0xfe, 0xad, 0x81, 0xcf, 0xf8, 0x44, 0x11, 0x06,
	0xf7, 0x3f, 0x8a, 0x11, 0x8a, 0xd8, 0x7d, 0x03, 0xa7, 0xe5, 0x8d, 0xad, 0xbf, 0x90, 0xf0, 0x6c,
	0x4b, 0xd6, 0x0e, 0xa5, 0x77, 0xc8, 0x17, 0x08, 0x6d, 0xcf, 0x8d, 0x7f, 0x60, 0xe7, 0xe5, 0xf1,
	0x4d, 0xf9, 0xf0, 0xf1, 0xd6, 0xfc, 0x72, 0xc0, 0x10, 0x4e, 0xec, 0x80, 0x55, 0x7b, 0x1f, 0x6d,
	0x5a, 0x35, 0xdb, 0x36, 0xdc, 0x94, 0x5a, 0x76, 0xbc, 0x86, 0xb3, 0xd5, 0x8e, 0xd5, 0xd7, 0xf2,
	0x64, 0xef, 0x12, 0x87, 0x74, 0x57, 0x49, 0x31, 0xe9, 0xea, 0xd0, 0x14, 0xbd, 0xf8, 0x17, 0x00,
	0x00, 0xff, 0xff, 0x14, 0x90, 0x63, 0x15, 0x69, 0x06, 0x00, 0x00,
}
