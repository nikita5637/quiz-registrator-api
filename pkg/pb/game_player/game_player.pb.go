// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: game_player/game_player.proto

package gameplayer

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Degree int32

const (
	Degree_DEGREE_INVALID  Degree = 0
	Degree_DEGREE_LIKELY   Degree = 1
	Degree_DEGREE_UNLIKELY Degree = 2
)

// Enum value maps for Degree.
var (
	Degree_name = map[int32]string{
		0: "DEGREE_INVALID",
		1: "DEGREE_LIKELY",
		2: "DEGREE_UNLIKELY",
	}
	Degree_value = map[string]int32{
		"DEGREE_INVALID":  0,
		"DEGREE_LIKELY":   1,
		"DEGREE_UNLIKELY": 2,
	}
)

func (x Degree) Enum() *Degree {
	p := new(Degree)
	*p = x
	return p
}

func (x Degree) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Degree) Descriptor() protoreflect.EnumDescriptor {
	return file_game_player_game_player_proto_enumTypes[0].Descriptor()
}

func (Degree) Type() protoreflect.EnumType {
	return &file_game_player_game_player_proto_enumTypes[0]
}

func (x Degree) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Degree.Descriptor instead.
func (Degree) EnumDescriptor() ([]byte, []int) {
	return file_game_player_game_player_proto_rawDescGZIP(), []int{0}
}

type GamePlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	GameId       int32                  `protobuf:"varint,2,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	UserId       *wrapperspb.Int32Value `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RegisteredBy int32                  `protobuf:"varint,4,opt,name=registered_by,json=registeredBy,proto3" json:"registered_by,omitempty"`
	Degree       Degree                 `protobuf:"varint,5,opt,name=degree,proto3,enum=game_player.Degree" json:"degree,omitempty"`
}

func (x *GamePlayer) Reset() {
	*x = GamePlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_game_player_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GamePlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GamePlayer) ProtoMessage() {}

func (x *GamePlayer) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_game_player_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GamePlayer.ProtoReflect.Descriptor instead.
func (*GamePlayer) Descriptor() ([]byte, []int) {
	return file_game_player_game_player_proto_rawDescGZIP(), []int{0}
}

func (x *GamePlayer) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GamePlayer) GetGameId() int32 {
	if x != nil {
		return x.GameId
	}
	return 0
}

func (x *GamePlayer) GetUserId() *wrapperspb.Int32Value {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *GamePlayer) GetRegisteredBy() int32 {
	if x != nil {
		return x.RegisteredBy
	}
	return 0
}

func (x *GamePlayer) GetDegree() Degree {
	if x != nil {
		return x.Degree
	}
	return Degree_DEGREE_INVALID
}

type CreateGamePlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GamePlayer *GamePlayer `protobuf:"bytes,1,opt,name=game_player,json=gamePlayer,proto3" json:"game_player,omitempty"`
}

func (x *CreateGamePlayerRequest) Reset() {
	*x = CreateGamePlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_game_player_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGamePlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGamePlayerRequest) ProtoMessage() {}

func (x *CreateGamePlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_game_player_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGamePlayerRequest.ProtoReflect.Descriptor instead.
func (*CreateGamePlayerRequest) Descriptor() ([]byte, []int) {
	return file_game_player_game_player_proto_rawDescGZIP(), []int{1}
}

func (x *CreateGamePlayerRequest) GetGamePlayer() *GamePlayer {
	if x != nil {
		return x.GamePlayer
	}
	return nil
}

type DeleteGamePlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteGamePlayerRequest) Reset() {
	*x = DeleteGamePlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_game_player_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGamePlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGamePlayerRequest) ProtoMessage() {}

func (x *DeleteGamePlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_game_player_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGamePlayerRequest.ProtoReflect.Descriptor instead.
func (*DeleteGamePlayerRequest) Descriptor() ([]byte, []int) {
	return file_game_player_game_player_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteGamePlayerRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetGamePlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetGamePlayerRequest) Reset() {
	*x = GetGamePlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_game_player_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGamePlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGamePlayerRequest) ProtoMessage() {}

func (x *GetGamePlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_game_player_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGamePlayerRequest.ProtoReflect.Descriptor instead.
func (*GetGamePlayerRequest) Descriptor() ([]byte, []int) {
	return file_game_player_game_player_proto_rawDescGZIP(), []int{3}
}

func (x *GetGamePlayerRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetGamePlayersByGameIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId int32 `protobuf:"varint,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
}

func (x *GetGamePlayersByGameIDRequest) Reset() {
	*x = GetGamePlayersByGameIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_game_player_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGamePlayersByGameIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGamePlayersByGameIDRequest) ProtoMessage() {}

func (x *GetGamePlayersByGameIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_game_player_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGamePlayersByGameIDRequest.ProtoReflect.Descriptor instead.
func (*GetGamePlayersByGameIDRequest) Descriptor() ([]byte, []int) {
	return file_game_player_game_player_proto_rawDescGZIP(), []int{4}
}

func (x *GetGamePlayersByGameIDRequest) GetGameId() int32 {
	if x != nil {
		return x.GameId
	}
	return 0
}

type GetGamePlayersByGameIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GamePlayers []*GamePlayer `protobuf:"bytes,1,rep,name=game_players,json=gamePlayers,proto3" json:"game_players,omitempty"`
}

func (x *GetGamePlayersByGameIDResponse) Reset() {
	*x = GetGamePlayersByGameIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_game_player_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGamePlayersByGameIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGamePlayersByGameIDResponse) ProtoMessage() {}

func (x *GetGamePlayersByGameIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_game_player_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGamePlayersByGameIDResponse.ProtoReflect.Descriptor instead.
func (*GetGamePlayersByGameIDResponse) Descriptor() ([]byte, []int) {
	return file_game_player_game_player_proto_rawDescGZIP(), []int{5}
}

func (x *GetGamePlayersByGameIDResponse) GetGamePlayers() []*GamePlayer {
	if x != nil {
		return x.GamePlayers
	}
	return nil
}

type PatchGamePlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GamePlayer *GamePlayer            `protobuf:"bytes,1,opt,name=game_player,json=gamePlayer,proto3" json:"game_player,omitempty"`
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *PatchGamePlayerRequest) Reset() {
	*x = PatchGamePlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_game_player_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchGamePlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchGamePlayerRequest) ProtoMessage() {}

func (x *PatchGamePlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_game_player_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchGamePlayerRequest.ProtoReflect.Descriptor instead.
func (*PatchGamePlayerRequest) Descriptor() ([]byte, []int) {
	return file_game_player_game_player_proto_rawDescGZIP(), []int{6}
}

func (x *PatchGamePlayerRequest) GetGamePlayer() *GamePlayer {
	if x != nil {
		return x.GamePlayer
	}
	return nil
}

func (x *PatchGamePlayerRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type RegisterPlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GamePlayer *GamePlayer `protobuf:"bytes,1,opt,name=game_player,json=gamePlayer,proto3" json:"game_player,omitempty"`
}

func (x *RegisterPlayerRequest) Reset() {
	*x = RegisterPlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_game_player_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterPlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterPlayerRequest) ProtoMessage() {}

func (x *RegisterPlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_game_player_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterPlayerRequest.ProtoReflect.Descriptor instead.
func (*RegisterPlayerRequest) Descriptor() ([]byte, []int) {
	return file_game_player_game_player_proto_rawDescGZIP(), []int{7}
}

func (x *RegisterPlayerRequest) GetGamePlayer() *GamePlayer {
	if x != nil {
		return x.GamePlayer
	}
	return nil
}

type UnregisterPlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GamePlayer *GamePlayer `protobuf:"bytes,1,opt,name=game_player,json=gamePlayer,proto3" json:"game_player,omitempty"`
}

func (x *UnregisterPlayerRequest) Reset() {
	*x = UnregisterPlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_game_player_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnregisterPlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnregisterPlayerRequest) ProtoMessage() {}

func (x *UnregisterPlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_game_player_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnregisterPlayerRequest.ProtoReflect.Descriptor instead.
func (*UnregisterPlayerRequest) Descriptor() ([]byte, []int) {
	return file_game_player_game_player_proto_rawDescGZIP(), []int{8}
}

func (x *UnregisterPlayerRequest) GetGamePlayer() *GamePlayer {
	if x != nil {
		return x.GamePlayer
	}
	return nil
}

var File_game_player_game_player_proto protoreflect.FileDescriptor

var file_game_player_game_player_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x67, 0x61,
	0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x0a,
	0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61,
	0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x61, 0x6d,
	0x65, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x42, 0x79, 0x12, 0x2b,
	0x0a, 0x06, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x67,
	0x72, 0x65, 0x65, 0x52, 0x06, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x22, 0x53, 0x0a, 0x17, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0b, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x22, 0x29, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x26, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x42, 0x79, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x22, 0x5c, 0x0a,
	0x1e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x42,
	0x79, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3a, 0x0a, 0x0c, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x0b,
	0x67, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x22, 0x8f, 0x01, 0x0a, 0x16,
	0x50, 0x61, 0x74, 0x63, 0x68, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0b, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73,
	0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x51, 0x0a,
	0x15, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0b, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x22, 0x53, 0x0a, 0x17, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0b, 0x67,
	0x61, 0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x47,
	0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x2a, 0x44, 0x0a, 0x06, 0x44, 0x65, 0x67, 0x72, 0x65, 0x65, 0x12,
	0x12, 0x0a, 0x0e, 0x44, 0x45, 0x47, 0x52, 0x45, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x45, 0x47, 0x52, 0x45, 0x45, 0x5f, 0x4c, 0x49,
	0x4b, 0x45, 0x4c, 0x59, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x45, 0x47, 0x52, 0x45, 0x45,
	0x5f, 0x55, 0x4e, 0x4c, 0x49, 0x4b, 0x45, 0x4c, 0x59, 0x10, 0x02, 0x32, 0xc9, 0x03, 0x0a, 0x07,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e,
	0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x10,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x12, 0x24, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x4d, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x12, 0x21, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x00, 0x12,
	0x73, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x42, 0x79, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x12, 0x2a, 0x2e, 0x67, 0x61, 0x6d, 0x65,
	0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x42, 0x79, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x42, 0x79, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0f, 0x50, 0x61, 0x74, 0x63, 0x68, 0x47, 0x61, 0x6d,
	0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x47, 0x61, 0x6d, 0x65, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x00, 0x32, 0xb8, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e,
	0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x12, 0x22, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x52,
	0x0a, 0x10, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x12, 0x24, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x2e, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6e, 0x69, 0x6b, 0x69, 0x74, 0x61, 0x35, 0x36, 0x33, 0x37, 0x2f, 0x71, 0x75, 0x69, 0x7a,
	0x2d, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2d, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x3b, 0x67, 0x61, 0x6d, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_player_game_player_proto_rawDescOnce sync.Once
	file_game_player_game_player_proto_rawDescData = file_game_player_game_player_proto_rawDesc
)

func file_game_player_game_player_proto_rawDescGZIP() []byte {
	file_game_player_game_player_proto_rawDescOnce.Do(func() {
		file_game_player_game_player_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_player_game_player_proto_rawDescData)
	})
	return file_game_player_game_player_proto_rawDescData
}

var file_game_player_game_player_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_game_player_game_player_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_game_player_game_player_proto_goTypes = []interface{}{
	(Degree)(0),                            // 0: game_player.Degree
	(*GamePlayer)(nil),                     // 1: game_player.GamePlayer
	(*CreateGamePlayerRequest)(nil),        // 2: game_player.CreateGamePlayerRequest
	(*DeleteGamePlayerRequest)(nil),        // 3: game_player.DeleteGamePlayerRequest
	(*GetGamePlayerRequest)(nil),           // 4: game_player.GetGamePlayerRequest
	(*GetGamePlayersByGameIDRequest)(nil),  // 5: game_player.GetGamePlayersByGameIDRequest
	(*GetGamePlayersByGameIDResponse)(nil), // 6: game_player.GetGamePlayersByGameIDResponse
	(*PatchGamePlayerRequest)(nil),         // 7: game_player.PatchGamePlayerRequest
	(*RegisterPlayerRequest)(nil),          // 8: game_player.RegisterPlayerRequest
	(*UnregisterPlayerRequest)(nil),        // 9: game_player.UnregisterPlayerRequest
	(*wrapperspb.Int32Value)(nil),          // 10: google.protobuf.Int32Value
	(*fieldmaskpb.FieldMask)(nil),          // 11: google.protobuf.FieldMask
	(*emptypb.Empty)(nil),                  // 12: google.protobuf.Empty
}
var file_game_player_game_player_proto_depIdxs = []int32{
	10, // 0: game_player.GamePlayer.user_id:type_name -> google.protobuf.Int32Value
	0,  // 1: game_player.GamePlayer.degree:type_name -> game_player.Degree
	1,  // 2: game_player.CreateGamePlayerRequest.game_player:type_name -> game_player.GamePlayer
	1,  // 3: game_player.GetGamePlayersByGameIDResponse.game_players:type_name -> game_player.GamePlayer
	1,  // 4: game_player.PatchGamePlayerRequest.game_player:type_name -> game_player.GamePlayer
	11, // 5: game_player.PatchGamePlayerRequest.update_mask:type_name -> google.protobuf.FieldMask
	1,  // 6: game_player.RegisterPlayerRequest.game_player:type_name -> game_player.GamePlayer
	1,  // 7: game_player.UnregisterPlayerRequest.game_player:type_name -> game_player.GamePlayer
	2,  // 8: game_player.Service.CreateGamePlayer:input_type -> game_player.CreateGamePlayerRequest
	3,  // 9: game_player.Service.DeleteGamePlayer:input_type -> game_player.DeleteGamePlayerRequest
	4,  // 10: game_player.Service.GetGamePlayer:input_type -> game_player.GetGamePlayerRequest
	5,  // 11: game_player.Service.GetGamePlayersByGameID:input_type -> game_player.GetGamePlayersByGameIDRequest
	7,  // 12: game_player.Service.PatchGamePlayer:input_type -> game_player.PatchGamePlayerRequest
	8,  // 13: game_player.RegistratorService.RegisterPlayer:input_type -> game_player.RegisterPlayerRequest
	9,  // 14: game_player.RegistratorService.UnregisterPlayer:input_type -> game_player.UnregisterPlayerRequest
	1,  // 15: game_player.Service.CreateGamePlayer:output_type -> game_player.GamePlayer
	12, // 16: game_player.Service.DeleteGamePlayer:output_type -> google.protobuf.Empty
	1,  // 17: game_player.Service.GetGamePlayer:output_type -> game_player.GamePlayer
	6,  // 18: game_player.Service.GetGamePlayersByGameID:output_type -> game_player.GetGamePlayersByGameIDResponse
	1,  // 19: game_player.Service.PatchGamePlayer:output_type -> game_player.GamePlayer
	12, // 20: game_player.RegistratorService.RegisterPlayer:output_type -> google.protobuf.Empty
	12, // 21: game_player.RegistratorService.UnregisterPlayer:output_type -> google.protobuf.Empty
	15, // [15:22] is the sub-list for method output_type
	8,  // [8:15] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_game_player_game_player_proto_init() }
func file_game_player_game_player_proto_init() {
	if File_game_player_game_player_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_player_game_player_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GamePlayer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_player_game_player_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGamePlayerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_player_game_player_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGamePlayerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_player_game_player_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGamePlayerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_player_game_player_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGamePlayersByGameIDRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_player_game_player_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGamePlayersByGameIDResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_player_game_player_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchGamePlayerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_player_game_player_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterPlayerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_player_game_player_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnregisterPlayerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_game_player_game_player_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_game_player_game_player_proto_goTypes,
		DependencyIndexes: file_game_player_game_player_proto_depIdxs,
		EnumInfos:         file_game_player_game_player_proto_enumTypes,
		MessageInfos:      file_game_player_game_player_proto_msgTypes,
	}.Build()
	File_game_player_game_player_proto = out.File
	file_game_player_game_player_proto_rawDesc = nil
	file_game_player_game_player_proto_goTypes = nil
	file_game_player_game_player_proto_depIdxs = nil
}
