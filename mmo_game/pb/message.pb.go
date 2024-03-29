// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

//同步玩家id
type SyncPid struct {
	Pid                  int32    `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncPid) Reset()         { *m = SyncPid{} }
func (m *SyncPid) String() string { return proto.CompactTextString(m) }
func (*SyncPid) ProtoMessage()    {}
func (*SyncPid) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *SyncPid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncPid.Unmarshal(m, b)
}
func (m *SyncPid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncPid.Marshal(b, m, deterministic)
}
func (m *SyncPid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncPid.Merge(m, src)
}
func (m *SyncPid) XXX_Size() int {
	return xxx_messageInfo_SyncPid.Size(m)
}
func (m *SyncPid) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncPid.DiscardUnknown(m)
}

var xxx_messageInfo_SyncPid proto.InternalMessageInfo

func (m *SyncPid) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

//玩家位置信息
type Position struct {
	X                    float32  `protobuf:"fixed32,1,opt,name=X,proto3" json:"X,omitempty"`
	Y                    float32  `protobuf:"fixed32,2,opt,name=Y,proto3" json:"Y,omitempty"`
	Z                    float32  `protobuf:"fixed32,3,opt,name=Z,proto3" json:"Z,omitempty"`
	V                    float32  `protobuf:"fixed32,4,opt,name=V,proto3" json:"V,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}
func (*Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *Position) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Position.Unmarshal(m, b)
}
func (m *Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Position.Marshal(b, m, deterministic)
}
func (m *Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Position.Merge(m, src)
}
func (m *Position) XXX_Size() int {
	return xxx_messageInfo_Position.Size(m)
}
func (m *Position) XXX_DiscardUnknown() {
	xxx_messageInfo_Position.DiscardUnknown(m)
}

var xxx_messageInfo_Position proto.InternalMessageInfo

func (m *Position) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Position) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Position) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

func (m *Position) GetV() float32 {
	if m != nil {
		return m.V
	}
	return 0
}

//广播消息
type BroadCast struct {
	Pid int32 `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"`
	Tp  int32 `protobuf:"varint,2,opt,name=Tp,proto3" json:"Tp,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*BroadCast_Content
	//	*BroadCast_P
	//	*BroadCast_ActionData
	Data                 isBroadCast_Data `protobuf_oneof:"Data"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BroadCast) Reset()         { *m = BroadCast{} }
func (m *BroadCast) String() string { return proto.CompactTextString(m) }
func (*BroadCast) ProtoMessage()    {}
func (*BroadCast) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *BroadCast) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BroadCast.Unmarshal(m, b)
}
func (m *BroadCast) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BroadCast.Marshal(b, m, deterministic)
}
func (m *BroadCast) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BroadCast.Merge(m, src)
}
func (m *BroadCast) XXX_Size() int {
	return xxx_messageInfo_BroadCast.Size(m)
}
func (m *BroadCast) XXX_DiscardUnknown() {
	xxx_messageInfo_BroadCast.DiscardUnknown(m)
}

var xxx_messageInfo_BroadCast proto.InternalMessageInfo

func (m *BroadCast) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *BroadCast) GetTp() int32 {
	if m != nil {
		return m.Tp
	}
	return 0
}

type isBroadCast_Data interface {
	isBroadCast_Data()
}

type BroadCast_Content struct {
	Content string `protobuf:"bytes,3,opt,name=Content,proto3,oneof"`
}

type BroadCast_P struct {
	P *Position `protobuf:"bytes,4,opt,name=P,proto3,oneof"`
}

type BroadCast_ActionData struct {
	ActionData int32 `protobuf:"varint,5,opt,name=ActionData,proto3,oneof"`
}

func (*BroadCast_Content) isBroadCast_Data() {}

func (*BroadCast_P) isBroadCast_Data() {}

func (*BroadCast_ActionData) isBroadCast_Data() {}

func (m *BroadCast) GetData() isBroadCast_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *BroadCast) GetContent() string {
	if x, ok := m.GetData().(*BroadCast_Content); ok {
		return x.Content
	}
	return ""
}

func (m *BroadCast) GetP() *Position {
	if x, ok := m.GetData().(*BroadCast_P); ok {
		return x.P
	}
	return nil
}

func (m *BroadCast) GetActionData() int32 {
	if x, ok := m.GetData().(*BroadCast_ActionData); ok {
		return x.ActionData
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BroadCast) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BroadCast_Content)(nil),
		(*BroadCast_P)(nil),
		(*BroadCast_ActionData)(nil),
	}
}

//世界聊天
type Talk struct {
	Content              string   `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Talk) Reset()         { *m = Talk{} }
func (m *Talk) String() string { return proto.CompactTextString(m) }
func (*Talk) ProtoMessage()    {}
func (*Talk) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{3}
}

func (m *Talk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Talk.Unmarshal(m, b)
}
func (m *Talk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Talk.Marshal(b, m, deterministic)
}
func (m *Talk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Talk.Merge(m, src)
}
func (m *Talk) XXX_Size() int {
	return xxx_messageInfo_Talk.Size(m)
}
func (m *Talk) XXX_DiscardUnknown() {
	xxx_messageInfo_Talk.DiscardUnknown(m)
}

var xxx_messageInfo_Talk proto.InternalMessageInfo

func (m *Talk) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

//同步玩家的显示数据
type SyncPlayers struct {
	Ps                   []*Player `protobuf:"bytes,1,rep,name=ps,proto3" json:"ps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SyncPlayers) Reset()         { *m = SyncPlayers{} }
func (m *SyncPlayers) String() string { return proto.CompactTextString(m) }
func (*SyncPlayers) ProtoMessage()    {}
func (*SyncPlayers) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{4}
}

func (m *SyncPlayers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncPlayers.Unmarshal(m, b)
}
func (m *SyncPlayers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncPlayers.Marshal(b, m, deterministic)
}
func (m *SyncPlayers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncPlayers.Merge(m, src)
}
func (m *SyncPlayers) XXX_Size() int {
	return xxx_messageInfo_SyncPlayers.Size(m)
}
func (m *SyncPlayers) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncPlayers.DiscardUnknown(m)
}

var xxx_messageInfo_SyncPlayers proto.InternalMessageInfo

func (m *SyncPlayers) GetPs() []*Player {
	if m != nil {
		return m.Ps
	}
	return nil
}

//玩家信息
type Player struct {
	Pid                  int32     `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"`
	P                    *Position `protobuf:"bytes,2,opt,name=P,proto3" json:"P,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{5}
}

func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (m *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(m, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *Player) GetP() *Position {
	if m != nil {
		return m.P
	}
	return nil
}

func init() {
	proto.RegisterType((*SyncPid)(nil), "pb.SyncPid")
	proto.RegisterType((*Position)(nil), "pb.Position")
	proto.RegisterType((*BroadCast)(nil), "pb.BroadCast")
	proto.RegisterType((*Talk)(nil), "pb.Talk")
	proto.RegisterType((*SyncPlayers)(nil), "pb.SyncPlayers")
	proto.RegisterType((*Player)(nil), "pb.Player")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x6b, 0xe7, 0xa7, 0xe4, 0xa6, 0x20, 0xe4, 0xc9, 0x0a, 0x0c, 0x91, 0xa7, 0xb2, 0x64,
	0x28, 0x12, 0x3b, 0x29, 0x43, 0x46, 0xcb, 0x44, 0x55, 0xdb, 0xcd, 0x69, 0x22, 0x14, 0x51, 0x62,
	0x2b, 0xf6, 0xd2, 0xc7, 0xe0, 0x35, 0x78, 0x4a, 0x64, 0x47, 0x05, 0xa4, 0x32, 0xd9, 0xdf, 0x3d,
	0xd2, 0x39, 0xc7, 0xbe, 0x70, 0xfd, 0xd1, 0x19, 0x23, 0xdf, 0xba, 0x42, 0x8f, 0xca, 0x2a, 0x82,
	0x75, 0xc3, 0xee, 0x60, 0xfe, 0x7a, 0x1a, 0x0e, 0xbc, 0x6f, 0xc9, 0x2d, 0x04, 0xbc, 0x6f, 0x29,
	0xca, 0xd1, 0x32, 0x12, 0xee, 0xca, 0x4a, 0xb8, 0xe2, 0xca, 0xf4, 0xb6, 0x57, 0x03, 0x59, 0x00,
	0xda, 0x7a, 0x0d, 0x0b, 0xb4, 0x75, 0xb4, 0xa3, 0x78, 0xa2, 0x9d, 0xa3, 0x3d, 0x0d, 0x26, 0xda,
	0x3b, 0xda, 0xd0, 0x70, 0xa2, 0x0d, 0xfb, 0x44, 0x90, 0x94, 0xa3, 0x92, 0xed, 0x5a, 0x1a, 0x7b,
	0x99, 0x41, 0x6e, 0x00, 0xd7, 0xda, 0x5b, 0x45, 0x02, 0xd7, 0x9a, 0x64, 0x30, 0x5f, 0xab, 0xc1,
	0x76, 0x83, 0xf5, 0x8e, 0x49, 0x35, 0x13, 0xe7, 0x01, 0xb9, 0x07, 0xc4, 0xbd, 0x73, 0xba, 0x5a,
	0x14, 0xba, 0x29, 0xce, 0xe5, 0xaa, 0x99, 0x40, 0x9c, 0xe4, 0x00, 0xcf, 0x07, 0x87, 0x2f, 0xd2,
	0x4a, 0x1a, 0x39, 0xc7, 0x6a, 0x26, 0xfe, 0xcc, 0xca, 0x18, 0x42, 0x77, 0xb2, 0x1c, 0xc2, 0x5a,
	0x1e, 0xdf, 0x09, 0xfd, 0xcd, 0x72, 0x8d, 0x92, 0x9f, 0x24, 0xf6, 0x00, 0xa9, 0xff, 0x96, 0xa3,
	0x3c, 0x75, 0xa3, 0x21, 0x19, 0x60, 0x6d, 0x28, 0xca, 0x83, 0x65, 0xba, 0x02, 0x9f, 0xec, 0x05,
	0x81, 0xb5, 0x61, 0x4f, 0x10, 0x4f, 0xf4, 0xcf, 0xe3, 0x32, 0x57, 0x18, 0x5f, 0x16, 0x16, 0x88,
	0x97, 0xd1, 0x17, 0xc6, 0xbc, 0x69, 0x62, 0xbf, 0x8b, 0xc7, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x2f, 0x8d, 0x58, 0x72, 0x9c, 0x01, 0x00, 0x00,
}
