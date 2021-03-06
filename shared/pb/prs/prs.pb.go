// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/prs.proto

package prs

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ParseRequest struct {
	ReplayUrl            string   `protobuf:"bytes,1,opt,name=replayUrl,proto3" json:"replayUrl,omitempty"`
	PlayerId             uint64   `protobuf:"varint,2,opt,name=playerId,proto3" json:"playerId,omitempty"`
	MatchId              uint64   `protobuf:"varint,3,opt,name=matchId,proto3" json:"matchId,omitempty"`
	ReplaySalt           uint64   `protobuf:"varint,4,opt,name=replaySalt,proto3" json:"replaySalt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParseRequest) Reset()         { *m = ParseRequest{} }
func (m *ParseRequest) String() string { return proto.CompactTextString(m) }
func (*ParseRequest) ProtoMessage()    {}
func (*ParseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6a3b567db2d869d, []int{0}
}

func (m *ParseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParseRequest.Unmarshal(m, b)
}
func (m *ParseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParseRequest.Marshal(b, m, deterministic)
}
func (m *ParseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParseRequest.Merge(m, src)
}
func (m *ParseRequest) XXX_Size() int {
	return xxx_messageInfo_ParseRequest.Size(m)
}
func (m *ParseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ParseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ParseRequest proto.InternalMessageInfo

func (m *ParseRequest) GetReplayUrl() string {
	if m != nil {
		return m.ReplayUrl
	}
	return ""
}

func (m *ParseRequest) GetPlayerId() uint64 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *ParseRequest) GetMatchId() uint64 {
	if m != nil {
		return m.MatchId
	}
	return 0
}

func (m *ParseRequest) GetReplaySalt() uint64 {
	if m != nil {
		return m.ReplaySalt
	}
	return 0
}

type ParseResult struct {
	ReplayData           *ReplayData `protobuf:"bytes,1,opt,name=replayData,proto3" json:"replayData,omitempty"`
	Success              bool        `protobuf:"varint,2,opt,name=Success,proto3" json:"Success,omitempty"`
	ErrorInfo            string      `protobuf:"bytes,3,opt,name=ErrorInfo,proto3" json:"ErrorInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ParseResult) Reset()         { *m = ParseResult{} }
func (m *ParseResult) String() string { return proto.CompactTextString(m) }
func (*ParseResult) ProtoMessage()    {}
func (*ParseResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6a3b567db2d869d, []int{1}
}

func (m *ParseResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParseResult.Unmarshal(m, b)
}
func (m *ParseResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParseResult.Marshal(b, m, deterministic)
}
func (m *ParseResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParseResult.Merge(m, src)
}
func (m *ParseResult) XXX_Size() int {
	return xxx_messageInfo_ParseResult.Size(m)
}
func (m *ParseResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ParseResult.DiscardUnknown(m)
}

var xxx_messageInfo_ParseResult proto.InternalMessageInfo

func (m *ParseResult) GetReplayData() *ReplayData {
	if m != nil {
		return m.ReplayData
	}
	return nil
}

func (m *ParseResult) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ParseResult) GetErrorInfo() string {
	if m != nil {
		return m.ErrorInfo
	}
	return ""
}

type ReplayData struct {
	GameTotalTimeSec     uint64      `protobuf:"varint,1,opt,name=gameTotalTimeSec,proto3" json:"gameTotalTimeSec,omitempty"`
	MovesMap             []*MovesMap `protobuf:"bytes,2,rep,name=movesMap,proto3" json:"movesMap,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ReplayData) Reset()         { *m = ReplayData{} }
func (m *ReplayData) String() string { return proto.CompactTextString(m) }
func (*ReplayData) ProtoMessage()    {}
func (*ReplayData) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6a3b567db2d869d, []int{2}
}

func (m *ReplayData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplayData.Unmarshal(m, b)
}
func (m *ReplayData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplayData.Marshal(b, m, deterministic)
}
func (m *ReplayData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplayData.Merge(m, src)
}
func (m *ReplayData) XXX_Size() int {
	return xxx_messageInfo_ReplayData.Size(m)
}
func (m *ReplayData) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplayData.DiscardUnknown(m)
}

var xxx_messageInfo_ReplayData proto.InternalMessageInfo

func (m *ReplayData) GetGameTotalTimeSec() uint64 {
	if m != nil {
		return m.GameTotalTimeSec
	}
	return 0
}

func (m *ReplayData) GetMovesMap() []*MovesMap {
	if m != nil {
		return m.MovesMap
	}
	return nil
}

type MovesMap struct {
	HeroName             string   `protobuf:"bytes,1,opt,name=heroName,proto3" json:"heroName,omitempty"`
	Moves                []*Move  `protobuf:"bytes,2,rep,name=moves,proto3" json:"moves,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MovesMap) Reset()         { *m = MovesMap{} }
func (m *MovesMap) String() string { return proto.CompactTextString(m) }
func (*MovesMap) ProtoMessage()    {}
func (*MovesMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6a3b567db2d869d, []int{3}
}

func (m *MovesMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MovesMap.Unmarshal(m, b)
}
func (m *MovesMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MovesMap.Marshal(b, m, deterministic)
}
func (m *MovesMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MovesMap.Merge(m, src)
}
func (m *MovesMap) XXX_Size() int {
	return xxx_messageInfo_MovesMap.Size(m)
}
func (m *MovesMap) XXX_DiscardUnknown() {
	xxx_messageInfo_MovesMap.DiscardUnknown(m)
}

var xxx_messageInfo_MovesMap proto.InternalMessageInfo

func (m *MovesMap) GetHeroName() string {
	if m != nil {
		return m.HeroName
	}
	return ""
}

func (m *MovesMap) GetMoves() []*Move {
	if m != nil {
		return m.Moves
	}
	return nil
}

type Move struct {
	Time                 uint64   `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	X                    uint64   `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
	Y                    uint64   `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`
	FullX                uint64   `protobuf:"varint,4,opt,name=fullX,proto3" json:"fullX,omitempty"`
	FullY                uint64   `protobuf:"varint,5,opt,name=fullY,proto3" json:"fullY,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Move) Reset()         { *m = Move{} }
func (m *Move) String() string { return proto.CompactTextString(m) }
func (*Move) ProtoMessage()    {}
func (*Move) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6a3b567db2d869d, []int{4}
}

func (m *Move) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Move.Unmarshal(m, b)
}
func (m *Move) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Move.Marshal(b, m, deterministic)
}
func (m *Move) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Move.Merge(m, src)
}
func (m *Move) XXX_Size() int {
	return xxx_messageInfo_Move.Size(m)
}
func (m *Move) XXX_DiscardUnknown() {
	xxx_messageInfo_Move.DiscardUnknown(m)
}

var xxx_messageInfo_Move proto.InternalMessageInfo

func (m *Move) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Move) GetX() uint64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Move) GetY() uint64 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Move) GetFullX() uint64 {
	if m != nil {
		return m.FullX
	}
	return 0
}

func (m *Move) GetFullY() uint64 {
	if m != nil {
		return m.FullY
	}
	return 0
}

func init() {
	proto.RegisterType((*ParseRequest)(nil), "ParseRequest")
	proto.RegisterType((*ParseResult)(nil), "ParseResult")
	proto.RegisterType((*ReplayData)(nil), "ReplayData")
	proto.RegisterType((*MovesMap)(nil), "MovesMap")
	proto.RegisterType((*Move)(nil), "Move")
}

func init() { proto.RegisterFile("pb/prs.proto", fileDescriptor_b6a3b567db2d869d) }

var fileDescriptor_b6a3b567db2d869d = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xcd, 0x6a, 0xe3, 0x30,
	0x10, 0x80, 0xd7, 0x89, 0xbd, 0x6b, 0x8f, 0x1d, 0x76, 0x11, 0x7b, 0x30, 0x69, 0x29, 0xc1, 0xd0,
	0x12, 0x5a, 0x70, 0x20, 0x85, 0x3e, 0x40, 0x7f, 0x0e, 0x39, 0xa4, 0x14, 0x39, 0x85, 0xa6, 0x97,
	0xa2, 0x38, 0x93, 0x26, 0x60, 0xd7, 0xae, 0x24, 0x87, 0xe4, 0xda, 0x27, 0x2f, 0x96, 0xe5, 0x1f,
	0xe8, 0x4d, 0xdf, 0x37, 0xd6, 0x68, 0x3c, 0x33, 0xe0, 0xe5, 0xab, 0x49, 0xce, 0x45, 0x98, 0xf3,
	0x4c, 0x66, 0xc1, 0x97, 0x01, 0xde, 0x13, 0xe3, 0x02, 0x29, 0x7e, 0x16, 0x28, 0x24, 0x39, 0x05,
	0x87, 0x63, 0x9e, 0xb0, 0xe3, 0x33, 0x4f, 0x7c, 0x63, 0x64, 0x8c, 0x1d, 0xda, 0x0a, 0x32, 0x04,
	0xbb, 0x3c, 0x22, 0x9f, 0xad, 0xfd, 0xde, 0xc8, 0x18, 0x9b, 0xb4, 0x61, 0xe2, 0xc3, 0x9f, 0x94,
	0xc9, 0x78, 0x3b, 0x5b, 0xfb, 0x7d, 0x15, 0xaa, 0x91, 0x9c, 0x01, 0x54, 0x29, 0x22, 0x96, 0x48,
	0xdf, 0x54, 0xc1, 0x8e, 0x09, 0x38, 0xb8, 0xba, 0x06, 0x51, 0x24, 0x92, 0x5c, 0xd5, 0x9f, 0xdf,
	0x33, 0xc9, 0x54, 0x0d, 0xee, 0xd4, 0x0d, 0x69, 0xa3, 0x68, 0x27, 0x5c, 0xbe, 0x1a, 0x15, 0x71,
	0x8c, 0x42, 0xa8, 0x82, 0x6c, 0x5a, 0x63, 0xf9, 0x27, 0x0f, 0x9c, 0x67, 0x7c, 0xf6, 0xb1, 0xc9,
	0x54, 0x45, 0x0e, 0x6d, 0x45, 0xf0, 0x06, 0xd0, 0x66, 0x24, 0x97, 0xf0, 0xef, 0x9d, 0xa5, 0xb8,
	0xc8, 0x24, 0x4b, 0x16, 0xbb, 0x14, 0x23, 0x8c, 0xd5, 0xc3, 0x26, 0xfd, 0xe1, 0xc9, 0x39, 0xd8,
	0x69, 0xb6, 0x47, 0x31, 0x67, 0xb9, 0xdf, 0x1b, 0xf5, 0xc7, 0xee, 0xd4, 0x09, 0xe7, 0x5a, 0xd0,
	0x26, 0x14, 0xdc, 0x81, 0x5d, 0xdb, 0xb2, 0x6d, 0x5b, 0xe4, 0xd9, 0x23, 0x4b, 0x51, 0xf7, 0xb4,
	0x61, 0x72, 0x02, 0x96, 0xba, 0xa3, 0x73, 0x59, 0x2a, 0x17, 0xad, 0x5c, 0xb0, 0x02, 0xb3, 0x44,
	0x42, 0xc0, 0x94, 0x3b, 0x7d, 0xd9, 0xa4, 0xea, 0x4c, 0x3c, 0x30, 0x0e, 0x7a, 0x08, 0xc6, 0xa1,
	0xa4, 0xa3, 0xee, 0xbb, 0x71, 0x24, 0xff, 0xc1, 0xda, 0x14, 0x49, 0xf2, 0xa2, 0x9b, 0x5d, 0x41,
	0x6d, 0x97, 0xbe, 0xd5, 0xda, 0xe5, 0xf4, 0x46, 0x6f, 0x40, 0x84, 0x7c, 0xbf, 0x8b, 0x91, 0x5c,
	0x80, 0xa5, 0x98, 0x0c, 0xc2, 0xee, 0x66, 0x0c, 0xbd, 0xb0, 0x33, 0xa4, 0xe0, 0xd7, 0xed, 0xdf,
	0xd7, 0x81, 0xd8, 0x32, 0x8e, 0xeb, 0x49, 0xb5, 0x51, 0xab, 0xdf, 0x6a, 0xa5, 0xae, 0xbf, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xdf, 0x12, 0xe1, 0x3a, 0x62, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ParseServiceClient is the client API for ParseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ParseServiceClient interface {
	Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParseResult, error)
}

type parseServiceClient struct {
	cc *grpc.ClientConn
}

func NewParseServiceClient(cc *grpc.ClientConn) ParseServiceClient {
	return &parseServiceClient{cc}
}

func (c *parseServiceClient) Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParseResult, error) {
	out := new(ParseResult)
	err := c.cc.Invoke(ctx, "/ParseService/Parse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParseServiceServer is the server API for ParseService service.
type ParseServiceServer interface {
	Parse(context.Context, *ParseRequest) (*ParseResult, error)
}

func RegisterParseServiceServer(s *grpc.Server, srv ParseServiceServer) {
	s.RegisterService(&_ParseService_serviceDesc, srv)
}

func _ParseService_Parse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParseServiceServer).Parse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ParseService/Parse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParseServiceServer).Parse(ctx, req.(*ParseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ParseService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ParseService",
	HandlerType: (*ParseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Parse",
			Handler:    _ParseService_Parse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/prs.proto",
}
