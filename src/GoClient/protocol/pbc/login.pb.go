// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.17.3
// source: src/login.proto

package pbc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CMDDef int32

const (
	CMDDef_invalid_id                  CMDDef = 0
	CMDDef_common_udp_heartbeat        CMDDef = 1
	CMDDef_login_cs_get_battle_servers CMDDef = 1001
	CMDDef_login_sc_get_battle_servers CMDDef = 1002
	// 进入世界
	CMDDef_world_cs_enter_world CMDDef = 2001
	CMDDef_world_sc_enter_world CMDDef = 2002
	// 世界转发
	CMDDef_world_cs_forward_data      CMDDef = 2003
	CMDDef_world_sc_forward_data      CMDDef = 2004
	CMDDef_world_push_forward_data    CMDDef = 2005 //转发推送
	CMDDef_world_push_disconnected    CMDDef = 2006 //转发掉线
	CMDDef_world_sc_other_enter_world CMDDef = 2007
)

// Enum value maps for CMDDef.
var (
	CMDDef_name = map[int32]string{
		0:    "invalid_id",
		1:    "common_udp_heartbeat",
		1001: "login_cs_get_battle_servers",
		1002: "login_sc_get_battle_servers",
		2001: "world_cs_enter_world",
		2002: "world_sc_enter_world",
		2003: "world_cs_forward_data",
		2004: "world_sc_forward_data",
		2005: "world_push_forward_data",
		2006: "world_push_disconnected",
		2007: "world_sc_other_enter_world",
	}
	CMDDef_value = map[string]int32{
		"invalid_id":                  0,
		"common_udp_heartbeat":        1,
		"login_cs_get_battle_servers": 1001,
		"login_sc_get_battle_servers": 1002,
		"world_cs_enter_world":        2001,
		"world_sc_enter_world":        2002,
		"world_cs_forward_data":       2003,
		"world_sc_forward_data":       2004,
		"world_push_forward_data":     2005,
		"world_push_disconnected":     2006,
		"world_sc_other_enter_world":  2007,
	}
)

func (x CMDDef) Enum() *CMDDef {
	p := new(CMDDef)
	*p = x
	return p
}

func (x CMDDef) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CMDDef) Descriptor() protoreflect.EnumDescriptor {
	return file_src_login_proto_enumTypes[0].Descriptor()
}

func (CMDDef) Type() protoreflect.EnumType {
	return &file_src_login_proto_enumTypes[0]
}

func (x CMDDef) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CMDDef.Descriptor instead.
func (CMDDef) EnumDescriptor() ([]byte, []int) {
	return file_src_login_proto_rawDescGZIP(), []int{0}
}

// client-> server
type CsMessageWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd   uint32 `protobuf:"fixed32,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	SeqId uint32 `protobuf:"fixed32,2,opt,name=seq_id,json=seqId,proto3" json:"seq_id,omitempty"`
	Data  []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CsMessageWrapper) Reset() {
	*x = CsMessageWrapper{}
	mi := &file_src_login_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CsMessageWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CsMessageWrapper) ProtoMessage() {}

func (x *CsMessageWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_src_login_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CsMessageWrapper.ProtoReflect.Descriptor instead.
func (*CsMessageWrapper) Descriptor() ([]byte, []int) {
	return file_src_login_proto_rawDescGZIP(), []int{0}
}

func (x *CsMessageWrapper) GetCmd() uint32 {
	if x != nil {
		return x.Cmd
	}
	return 0
}

func (x *CsMessageWrapper) GetSeqId() uint32 {
	if x != nil {
		return x.SeqId
	}
	return 0
}

func (x *CsMessageWrapper) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// server->client
type ScMessageWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd    uint32 `protobuf:"fixed32,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	SeqId  uint32 `protobuf:"fixed32,2,opt,name=seq_id,json=seqId,proto3" json:"seq_id,omitempty"`
	Result int32  `protobuf:"varint,3,opt,name=result,proto3" json:"result,omitempty"`
	Data   []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ScMessageWrapper) Reset() {
	*x = ScMessageWrapper{}
	mi := &file_src_login_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScMessageWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScMessageWrapper) ProtoMessage() {}

func (x *ScMessageWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_src_login_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScMessageWrapper.ProtoReflect.Descriptor instead.
func (*ScMessageWrapper) Descriptor() ([]byte, []int) {
	return file_src_login_proto_rawDescGZIP(), []int{1}
}

func (x *ScMessageWrapper) GetCmd() uint32 {
	if x != nil {
		return x.Cmd
	}
	return 0
}

func (x *ScMessageWrapper) GetSeqId() uint32 {
	if x != nil {
		return x.SeqId
	}
	return 0
}

func (x *ScMessageWrapper) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *ScMessageWrapper) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type CsGetBattleServers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CsGetBattleServers) Reset() {
	*x = CsGetBattleServers{}
	mi := &file_src_login_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CsGetBattleServers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CsGetBattleServers) ProtoMessage() {}

func (x *CsGetBattleServers) ProtoReflect() protoreflect.Message {
	mi := &file_src_login_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CsGetBattleServers.ProtoReflect.Descriptor instead.
func (*CsGetBattleServers) Descriptor() ([]byte, []int) {
	return file_src_login_proto_rawDescGZIP(), []int{2}
}

type ScGetBattleServers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerList []string `protobuf:"bytes,1,rep,name=serverList,proto3" json:"serverList,omitempty"`
}

func (x *ScGetBattleServers) Reset() {
	*x = ScGetBattleServers{}
	mi := &file_src_login_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScGetBattleServers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScGetBattleServers) ProtoMessage() {}

func (x *ScGetBattleServers) ProtoReflect() protoreflect.Message {
	mi := &file_src_login_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScGetBattleServers.ProtoReflect.Descriptor instead.
func (*ScGetBattleServers) Descriptor() ([]byte, []int) {
	return file_src_login_proto_rawDescGZIP(), []int{3}
}

func (x *ScGetBattleServers) GetServerList() []string {
	if x != nil {
		return x.ServerList
	}
	return nil
}

type PlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  uint64  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	PosX float32 `protobuf:"fixed32,2,opt,name=pos_x,json=posX,proto3" json:"pos_x,omitempty"`
	PosY float32 `protobuf:"fixed32,3,opt,name=pos_y,json=posY,proto3" json:"pos_y,omitempty"`
	PosZ float32 `protobuf:"fixed32,4,opt,name=pos_z,json=posZ,proto3" json:"pos_z,omitempty"`
}

func (x *PlayerInfo) Reset() {
	*x = PlayerInfo{}
	mi := &file_src_login_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerInfo) ProtoMessage() {}

func (x *PlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_src_login_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerInfo.ProtoReflect.Descriptor instead.
func (*PlayerInfo) Descriptor() ([]byte, []int) {
	return file_src_login_proto_rawDescGZIP(), []int{4}
}

func (x *PlayerInfo) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *PlayerInfo) GetPosX() float32 {
	if x != nil {
		return x.PosX
	}
	return 0
}

func (x *PlayerInfo) GetPosY() float32 {
	if x != nil {
		return x.PosY
	}
	return 0
}

func (x *PlayerInfo) GetPosZ() float32 {
	if x != nil {
		return x.PosZ
	}
	return 0
}

type CsEnterWorld struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"` //玩家uid
}

func (x *CsEnterWorld) Reset() {
	*x = CsEnterWorld{}
	mi := &file_src_login_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CsEnterWorld) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CsEnterWorld) ProtoMessage() {}

func (x *CsEnterWorld) ProtoReflect() protoreflect.Message {
	mi := &file_src_login_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CsEnterWorld.ProtoReflect.Descriptor instead.
func (*CsEnterWorld) Descriptor() ([]byte, []int) {
	return file_src_login_proto_rawDescGZIP(), []int{5}
}

func (x *CsEnterWorld) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type ScEnterWorld struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid              uint64        `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"` //玩家uid
	PosX             float32       `protobuf:"fixed32,2,opt,name=pos_x,json=posX,proto3" json:"pos_x,omitempty"`
	PosY             float32       `protobuf:"fixed32,3,opt,name=pos_y,json=posY,proto3" json:"pos_y,omitempty"`
	PosZ             float32       `protobuf:"fixed32,4,opt,name=pos_z,json=posZ,proto3" json:"pos_z,omitempty"`
	OnlinePlayerList []*PlayerInfo `protobuf:"bytes,5,rep,name=online_player_list,json=onlinePlayerList,proto3" json:"online_player_list,omitempty"`
}

func (x *ScEnterWorld) Reset() {
	*x = ScEnterWorld{}
	mi := &file_src_login_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScEnterWorld) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScEnterWorld) ProtoMessage() {}

func (x *ScEnterWorld) ProtoReflect() protoreflect.Message {
	mi := &file_src_login_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScEnterWorld.ProtoReflect.Descriptor instead.
func (*ScEnterWorld) Descriptor() ([]byte, []int) {
	return file_src_login_proto_rawDescGZIP(), []int{6}
}

func (x *ScEnterWorld) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ScEnterWorld) GetPosX() float32 {
	if x != nil {
		return x.PosX
	}
	return 0
}

func (x *ScEnterWorld) GetPosY() float32 {
	if x != nil {
		return x.PosY
	}
	return 0
}

func (x *ScEnterWorld) GetPosZ() float32 {
	if x != nil {
		return x.PosZ
	}
	return 0
}

func (x *ScEnterWorld) GetOnlinePlayerList() []*PlayerInfo {
	if x != nil {
		return x.OnlinePlayerList
	}
	return nil
}

type ScOtherEnterWorld struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OtherPlayer *PlayerInfo `protobuf:"bytes,1,opt,name=other_player,json=otherPlayer,proto3" json:"other_player,omitempty"`
}

func (x *ScOtherEnterWorld) Reset() {
	*x = ScOtherEnterWorld{}
	mi := &file_src_login_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScOtherEnterWorld) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScOtherEnterWorld) ProtoMessage() {}

func (x *ScOtherEnterWorld) ProtoReflect() protoreflect.Message {
	mi := &file_src_login_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScOtherEnterWorld.ProtoReflect.Descriptor instead.
func (*ScOtherEnterWorld) Descriptor() ([]byte, []int) {
	return file_src_login_proto_rawDescGZIP(), []int{7}
}

func (x *ScOtherEnterWorld) GetOtherPlayer() *PlayerInfo {
	if x != nil {
		return x.OtherPlayer
	}
	return nil
}

type CsForwardData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte  `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	PosX float32 `protobuf:"fixed32,2,opt,name=pos_x,json=posX,proto3" json:"pos_x,omitempty"`
	PosY float32 `protobuf:"fixed32,3,opt,name=pos_y,json=posY,proto3" json:"pos_y,omitempty"`
	PosZ float32 `protobuf:"fixed32,4,opt,name=pos_z,json=posZ,proto3" json:"pos_z,omitempty"`
}

func (x *CsForwardData) Reset() {
	*x = CsForwardData{}
	mi := &file_src_login_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CsForwardData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CsForwardData) ProtoMessage() {}

func (x *CsForwardData) ProtoReflect() protoreflect.Message {
	mi := &file_src_login_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CsForwardData.ProtoReflect.Descriptor instead.
func (*CsForwardData) Descriptor() ([]byte, []int) {
	return file_src_login_proto_rawDescGZIP(), []int{8}
}

func (x *CsForwardData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CsForwardData) GetPosX() float32 {
	if x != nil {
		return x.PosX
	}
	return 0
}

func (x *CsForwardData) GetPosY() float32 {
	if x != nil {
		return x.PosY
	}
	return 0
}

func (x *CsForwardData) GetPosZ() float32 {
	if x != nil {
		return x.PosZ
	}
	return 0
}

type ScForwardData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ScForwardData) Reset() {
	*x = ScForwardData{}
	mi := &file_src_login_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScForwardData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScForwardData) ProtoMessage() {}

func (x *ScForwardData) ProtoReflect() protoreflect.Message {
	mi := &file_src_login_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScForwardData.ProtoReflect.Descriptor instead.
func (*ScForwardData) Descriptor() ([]byte, []int) {
	return file_src_login_proto_rawDescGZIP(), []int{9}
}

type PushForwardData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ForwardPlayer *PlayerInfo `protobuf:"bytes,1,opt,name=forward_player,json=forwardPlayer,proto3" json:"forward_player,omitempty"`
	Data          []byte      `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"` //所转发的消息数据
}

func (x *PushForwardData) Reset() {
	*x = PushForwardData{}
	mi := &file_src_login_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PushForwardData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushForwardData) ProtoMessage() {}

func (x *PushForwardData) ProtoReflect() protoreflect.Message {
	mi := &file_src_login_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushForwardData.ProtoReflect.Descriptor instead.
func (*PushForwardData) Descriptor() ([]byte, []int) {
	return file_src_login_proto_rawDescGZIP(), []int{10}
}

func (x *PushForwardData) GetForwardPlayer() *PlayerInfo {
	if x != nil {
		return x.ForwardPlayer
	}
	return nil
}

func (x *PushForwardData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// 玩家掉线推送
type PushDisconnected struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"` //玩家uid
}

func (x *PushDisconnected) Reset() {
	*x = PushDisconnected{}
	mi := &file_src_login_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PushDisconnected) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushDisconnected) ProtoMessage() {}

func (x *PushDisconnected) ProtoReflect() protoreflect.Message {
	mi := &file_src_login_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushDisconnected.ProtoReflect.Descriptor instead.
func (*PushDisconnected) Descriptor() ([]byte, []int) {
	return file_src_login_proto_rawDescGZIP(), []int{11}
}

func (x *PushDisconnected) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type CsReqUdpHeartbeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seqid uint32 `protobuf:"varint,1,opt,name=seqid,proto3" json:"seqid,omitempty"`
}

func (x *CsReqUdpHeartbeat) Reset() {
	*x = CsReqUdpHeartbeat{}
	mi := &file_src_login_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CsReqUdpHeartbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CsReqUdpHeartbeat) ProtoMessage() {}

func (x *CsReqUdpHeartbeat) ProtoReflect() protoreflect.Message {
	mi := &file_src_login_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CsReqUdpHeartbeat.ProtoReflect.Descriptor instead.
func (*CsReqUdpHeartbeat) Descriptor() ([]byte, []int) {
	return file_src_login_proto_rawDescGZIP(), []int{12}
}

func (x *CsReqUdpHeartbeat) GetSeqid() uint32 {
	if x != nil {
		return x.Seqid
	}
	return 0
}

type ScRspUdpHeartbeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seqid uint32 `protobuf:"varint,1,opt,name=seqid,proto3" json:"seqid,omitempty"`
}

func (x *ScRspUdpHeartbeat) Reset() {
	*x = ScRspUdpHeartbeat{}
	mi := &file_src_login_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScRspUdpHeartbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScRspUdpHeartbeat) ProtoMessage() {}

func (x *ScRspUdpHeartbeat) ProtoReflect() protoreflect.Message {
	mi := &file_src_login_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScRspUdpHeartbeat.ProtoReflect.Descriptor instead.
func (*ScRspUdpHeartbeat) Descriptor() ([]byte, []int) {
	return file_src_login_proto_rawDescGZIP(), []int{13}
}

func (x *ScRspUdpHeartbeat) GetSeqid() uint32 {
	if x != nil {
		return x.Seqid
	}
	return 0
}

var File_src_login_proto protoreflect.FileDescriptor

var file_src_login_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x72, 0x63, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x22, 0x51, 0x0a, 0x12, 0x63, 0x73, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x07, 0x52, 0x03, 0x63, 0x6d, 0x64,
	0x12, 0x15, 0x0a, 0x06, 0x73, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x07,
	0x52, 0x05, 0x73, 0x65, 0x71, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x69, 0x0a, 0x12, 0x73,
	0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x07, 0x52, 0x03,
	0x63, 0x6d, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x07, 0x52, 0x05, 0x73, 0x65, 0x71, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x17, 0x0a, 0x15, 0x63, 0x73, 0x5f, 0x67, 0x65, 0x74,
	0x5f, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x22,
	0x37, 0x0a, 0x15, 0x73, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x5e, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x70, 0x6f, 0x73,
	0x5f, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x58, 0x12, 0x13,
	0x0a, 0x05, 0x70, 0x6f, 0x73, 0x5f, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x70,
	0x6f, 0x73, 0x59, 0x12, 0x13, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x5f, 0x7a, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x5a, 0x22, 0x22, 0x0a, 0x0e, 0x63, 0x73, 0x5f, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0xa3, 0x01, 0x0a,
	0x0e, 0x73, 0x63, 0x5f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x12, 0x13, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x5f, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x04, 0x70, 0x6f, 0x73, 0x58, 0x12, 0x13, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x5f, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x59, 0x12, 0x13, 0x0a, 0x05, 0x70,
	0x6f, 0x73, 0x5f, 0x7a, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x5a,
	0x12, 0x40, 0x0a, 0x12, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x52, 0x10, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x4d, 0x0a, 0x14, 0x73, 0x63, 0x5f, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x35, 0x0a, 0x0c, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x22, 0x64, 0x0a, 0x0f, 0x63, 0x73, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x13, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x5f,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x58, 0x12, 0x13, 0x0a,
	0x05, 0x70, 0x6f, 0x73, 0x5f, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x70, 0x6f,
	0x73, 0x59, 0x12, 0x13, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x5f, 0x7a, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x5a, 0x22, 0x11, 0x0a, 0x0f, 0x73, 0x63, 0x5f, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0x62, 0x0a, 0x11, 0x70, 0x75,
	0x73, 0x68, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x39, 0x0a, 0x0e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x25,
	0x0a, 0x11, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x14, 0x63, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x5f,
	0x75, 0x64, 0x70, 0x5f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x65, 0x71, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x65,
	0x71, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x14, 0x73, 0x63, 0x5f, 0x72, 0x73, 0x70, 0x5f, 0x75, 0x64,
	0x70, 0x5f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x65, 0x71, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x65, 0x71, 0x69,
	0x64, 0x2a, 0xc1, 0x02, 0x0a, 0x06, 0x43, 0x4d, 0x44, 0x44, 0x65, 0x66, 0x12, 0x0e, 0x0a, 0x0a,
	0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x69, 0x64, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x75, 0x64, 0x70, 0x5f, 0x68, 0x65, 0x61, 0x72, 0x74,
	0x62, 0x65, 0x61, 0x74, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f,
	0x63, 0x73, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x73, 0x10, 0xe9, 0x07, 0x12, 0x20, 0x0a, 0x1b, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x5f, 0x73, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x10, 0xea, 0x07, 0x12, 0x19, 0x0a, 0x14, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x5f, 0x63, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x10, 0xd1, 0x0f, 0x12, 0x19, 0x0a, 0x14, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x73,
	0x63, 0x5f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x10, 0xd2, 0x0f,
	0x12, 0x1a, 0x0a, 0x15, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x63, 0x73, 0x5f, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x10, 0xd3, 0x0f, 0x12, 0x1a, 0x0a, 0x15,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x73, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x10, 0xd4, 0x0f, 0x12, 0x1c, 0x0a, 0x17, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x5f, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x10, 0xd5, 0x0f, 0x12, 0x1c, 0x0a, 0x17, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f,
	0x70, 0x75, 0x73, 0x68, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x10, 0xd6, 0x0f, 0x12, 0x1f, 0x0a, 0x1a, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x73, 0x63,
	0x5f, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x10, 0xd7, 0x0f, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x70, 0x62, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_login_proto_rawDescOnce sync.Once
	file_src_login_proto_rawDescData = file_src_login_proto_rawDesc
)

func file_src_login_proto_rawDescGZIP() []byte {
	file_src_login_proto_rawDescOnce.Do(func() {
		file_src_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_login_proto_rawDescData)
	})
	return file_src_login_proto_rawDescData
}

var file_src_login_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_src_login_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_src_login_proto_goTypes = []any{
	(CMDDef)(0),                // 0: login.CMDDef
	(*CsMessageWrapper)(nil),   // 1: login.cs_message_wrapper
	(*ScMessageWrapper)(nil),   // 2: login.sc_message_wrapper
	(*CsGetBattleServers)(nil), // 3: login.cs_get_battle_servers
	(*ScGetBattleServers)(nil), // 4: login.sc_get_battle_servers
	(*PlayerInfo)(nil),         // 5: login.player_info
	(*CsEnterWorld)(nil),       // 6: login.cs_enter_world
	(*ScEnterWorld)(nil),       // 7: login.sc_enter_world
	(*ScOtherEnterWorld)(nil),  // 8: login.sc_other_enter_world
	(*CsForwardData)(nil),      // 9: login.cs_forward_data
	(*ScForwardData)(nil),      // 10: login.sc_forward_data
	(*PushForwardData)(nil),    // 11: login.push_forward_data
	(*PushDisconnected)(nil),   // 12: login.push_disconnected
	(*CsReqUdpHeartbeat)(nil),  // 13: login.cs_req_udp_heartbeat
	(*ScRspUdpHeartbeat)(nil),  // 14: login.sc_rsp_udp_heartbeat
}
var file_src_login_proto_depIdxs = []int32{
	5, // 0: login.sc_enter_world.online_player_list:type_name -> login.player_info
	5, // 1: login.sc_other_enter_world.other_player:type_name -> login.player_info
	5, // 2: login.push_forward_data.forward_player:type_name -> login.player_info
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_src_login_proto_init() }
func file_src_login_proto_init() {
	if File_src_login_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_src_login_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_login_proto_goTypes,
		DependencyIndexes: file_src_login_proto_depIdxs,
		EnumInfos:         file_src_login_proto_enumTypes,
		MessageInfos:      file_src_login_proto_msgTypes,
	}.Build()
	File_src_login_proto = out.File
	file_src_login_proto_rawDesc = nil
	file_src_login_proto_goTypes = nil
	file_src_login_proto_depIdxs = nil
}