// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: league/league.proto

package league

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LeagueID int32

const (
	LeagueID_INVALID       LeagueID = 0
	LeagueID_QUIZ_PLEASE   LeagueID = 1
	LeagueID_SQUIZ         LeagueID = 2
	LeagueID_SIXTY_SECONDS LeagueID = 3
	LeagueID_SHAKER        LeagueID = 4
	LeagueID_WOW           LeagueID = 5
	LeagueID_MOZGOBOYNYA   LeagueID = 6
	LeagueID_NASH_QUIZ     LeagueID = 7
	LeagueID_SMUZI         LeagueID = 8
	LeagueID_QUIZ_PEACE    LeagueID = 9
	LeagueID_JACK_QUIZ     LeagueID = 10
	LeagueID_OWL_CUP       LeagueID = 11
)

// Enum value maps for LeagueID.
var (
	LeagueID_name = map[int32]string{
		0:  "INVALID",
		1:  "QUIZ_PLEASE",
		2:  "SQUIZ",
		3:  "SIXTY_SECONDS",
		4:  "SHAKER",
		5:  "WOW",
		6:  "MOZGOBOYNYA",
		7:  "NASH_QUIZ",
		8:  "SMUZI",
		9:  "QUIZ_PEACE",
		10: "JACK_QUIZ",
		11: "OWL_CUP",
	}
	LeagueID_value = map[string]int32{
		"INVALID":       0,
		"QUIZ_PLEASE":   1,
		"SQUIZ":         2,
		"SIXTY_SECONDS": 3,
		"SHAKER":        4,
		"WOW":           5,
		"MOZGOBOYNYA":   6,
		"NASH_QUIZ":     7,
		"SMUZI":         8,
		"QUIZ_PEACE":    9,
		"JACK_QUIZ":     10,
		"OWL_CUP":       11,
	}
)

func (x LeagueID) Enum() *LeagueID {
	p := new(LeagueID)
	*p = x
	return p
}

func (x LeagueID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LeagueID) Descriptor() protoreflect.EnumDescriptor {
	return file_league_league_proto_enumTypes[0].Descriptor()
}

func (LeagueID) Type() protoreflect.EnumType {
	return &file_league_league_proto_enumTypes[0]
}

func (x LeagueID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LeagueID.Descriptor instead.
func (LeagueID) EnumDescriptor() ([]byte, []int) {
	return file_league_league_proto_rawDescGZIP(), []int{0}
}

type League struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ShortName string `protobuf:"bytes,3,opt,name=short_name,json=shortName,proto3" json:"short_name,omitempty"`
	LogoLink  string `protobuf:"bytes,4,opt,name=logo_link,json=logoLink,proto3" json:"logo_link,omitempty"`
	WebSite   string `protobuf:"bytes,5,opt,name=web_site,json=webSite,proto3" json:"web_site,omitempty"`
}

func (x *League) Reset() {
	*x = League{}
	if protoimpl.UnsafeEnabled {
		mi := &file_league_league_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *League) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*League) ProtoMessage() {}

func (x *League) ProtoReflect() protoreflect.Message {
	mi := &file_league_league_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use League.ProtoReflect.Descriptor instead.
func (*League) Descriptor() ([]byte, []int) {
	return file_league_league_proto_rawDescGZIP(), []int{0}
}

func (x *League) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *League) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *League) GetShortName() string {
	if x != nil {
		return x.ShortName
	}
	return ""
}

func (x *League) GetLogoLink() string {
	if x != nil {
		return x.LogoLink
	}
	return ""
}

func (x *League) GetWebSite() string {
	if x != nil {
		return x.WebSite
	}
	return ""
}

type GetLeagueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetLeagueRequest) Reset() {
	*x = GetLeagueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_league_league_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLeagueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLeagueRequest) ProtoMessage() {}

func (x *GetLeagueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_league_league_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLeagueRequest.ProtoReflect.Descriptor instead.
func (*GetLeagueRequest) Descriptor() ([]byte, []int) {
	return file_league_league_proto_rawDescGZIP(), []int{1}
}

func (x *GetLeagueRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_league_league_proto protoreflect.FileDescriptor

var file_league_league_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x2f, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x22, 0x83, 0x01,
	0x0a, 0x06, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x6f, 0x67, 0x6f, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x6f, 0x67, 0x6f, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x19, 0x0a, 0x08, 0x77, 0x65, 0x62, 0x5f,
	0x73, 0x69, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x53,
	0x69, 0x74, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x2a, 0xb2, 0x01, 0x0a, 0x08, 0x4c, 0x65, 0x61, 0x67,
	0x75, 0x65, 0x49, 0x44, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10,
	0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x51, 0x55, 0x49, 0x5a, 0x5f, 0x50, 0x4c, 0x45, 0x41, 0x53, 0x45,
	0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x51, 0x55, 0x49, 0x5a, 0x10, 0x02, 0x12, 0x11, 0x0a,
	0x0d, 0x53, 0x49, 0x58, 0x54, 0x59, 0x5f, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x53, 0x10, 0x03,
	0x12, 0x0a, 0x0a, 0x06, 0x53, 0x48, 0x41, 0x4b, 0x45, 0x52, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03,
	0x57, 0x4f, 0x57, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x4f, 0x5a, 0x47, 0x4f, 0x42, 0x4f,
	0x59, 0x4e, 0x59, 0x41, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x41, 0x53, 0x48, 0x5f, 0x51,
	0x55, 0x49, 0x5a, 0x10, 0x07, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x4d, 0x55, 0x5a, 0x49, 0x10, 0x08,
	0x12, 0x0e, 0x0a, 0x0a, 0x51, 0x55, 0x49, 0x5a, 0x5f, 0x50, 0x45, 0x41, 0x43, 0x45, 0x10, 0x09,
	0x12, 0x0d, 0x0a, 0x09, 0x4a, 0x41, 0x43, 0x4b, 0x5f, 0x51, 0x55, 0x49, 0x5a, 0x10, 0x0a, 0x12,
	0x0b, 0x0a, 0x07, 0x4f, 0x57, 0x4c, 0x5f, 0x43, 0x55, 0x50, 0x10, 0x0b, 0x32, 0x42, 0x0a, 0x07,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4c, 0x65,
	0x61, 0x67, 0x75, 0x65, 0x12, 0x18, 0x2e, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x2e, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x22, 0x00,
	0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e,
	0x69, 0x6b, 0x69, 0x74, 0x61, 0x35, 0x36, 0x33, 0x37, 0x2f, 0x71, 0x75, 0x69, 0x7a, 0x2d, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x3b, 0x6c, 0x65, 0x61,
	0x67, 0x75, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_league_league_proto_rawDescOnce sync.Once
	file_league_league_proto_rawDescData = file_league_league_proto_rawDesc
)

func file_league_league_proto_rawDescGZIP() []byte {
	file_league_league_proto_rawDescOnce.Do(func() {
		file_league_league_proto_rawDescData = protoimpl.X.CompressGZIP(file_league_league_proto_rawDescData)
	})
	return file_league_league_proto_rawDescData
}

var file_league_league_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_league_league_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_league_league_proto_goTypes = []interface{}{
	(LeagueID)(0),            // 0: league.LeagueID
	(*League)(nil),           // 1: league.League
	(*GetLeagueRequest)(nil), // 2: league.GetLeagueRequest
}
var file_league_league_proto_depIdxs = []int32{
	2, // 0: league.Service.GetLeague:input_type -> league.GetLeagueRequest
	1, // 1: league.Service.GetLeague:output_type -> league.League
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_league_league_proto_init() }
func file_league_league_proto_init() {
	if File_league_league_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_league_league_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*League); i {
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
		file_league_league_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLeagueRequest); i {
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
			RawDescriptor: file_league_league_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_league_league_proto_goTypes,
		DependencyIndexes: file_league_league_proto_depIdxs,
		EnumInfos:         file_league_league_proto_enumTypes,
		MessageInfos:      file_league_league_proto_msgTypes,
	}.Build()
	File_league_league_proto = out.File
	file_league_league_proto_rawDesc = nil
	file_league_league_proto_goTypes = nil
	file_league_league_proto_depIdxs = nil
}
