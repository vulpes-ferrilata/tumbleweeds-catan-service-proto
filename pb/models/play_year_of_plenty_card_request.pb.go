// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: protobuf/models/play_year_of_plenty_card_request.proto

package models

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

type PlayYearOfPlentyCardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameID                     string   `protobuf:"bytes,1,opt,name=GameID,proto3" json:"GameID,omitempty"`
	UserID                     string   `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	DevelopmentCardID          string   `protobuf:"bytes,3,opt,name=DevelopmentCardID,proto3" json:"DevelopmentCardID,omitempty"`
	DemandingResourceCardTypes []string `protobuf:"bytes,4,rep,name=DemandingResourceCardTypes,proto3" json:"DemandingResourceCardTypes,omitempty"`
}

func (x *PlayYearOfPlentyCardRequest) Reset() {
	*x = PlayYearOfPlentyCardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_models_play_year_of_plenty_card_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayYearOfPlentyCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayYearOfPlentyCardRequest) ProtoMessage() {}

func (x *PlayYearOfPlentyCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_models_play_year_of_plenty_card_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayYearOfPlentyCardRequest.ProtoReflect.Descriptor instead.
func (*PlayYearOfPlentyCardRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_models_play_year_of_plenty_card_request_proto_rawDescGZIP(), []int{0}
}

func (x *PlayYearOfPlentyCardRequest) GetGameID() string {
	if x != nil {
		return x.GameID
	}
	return ""
}

func (x *PlayYearOfPlentyCardRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *PlayYearOfPlentyCardRequest) GetDevelopmentCardID() string {
	if x != nil {
		return x.DevelopmentCardID
	}
	return ""
}

func (x *PlayYearOfPlentyCardRequest) GetDemandingResourceCardTypes() []string {
	if x != nil {
		return x.DemandingResourceCardTypes
	}
	return nil
}

var File_protobuf_models_play_year_of_plenty_card_request_proto protoreflect.FileDescriptor

var file_protobuf_models_play_year_of_plenty_card_request_proto_rawDesc = []byte{
	0x0a, 0x36, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x70,
	0x6c, 0x65, 0x6e, 0x74, 0x79, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x22, 0xbb, 0x01, 0x0a, 0x1b, 0x50, 0x6c, 0x61, 0x79, 0x59, 0x65, 0x61, 0x72, 0x4f, 0x66, 0x50,
	0x6c, 0x65, 0x6e, 0x74, 0x79, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x2c, 0x0a, 0x11, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x43,
	0x61, 0x72, 0x64, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x44, 0x65, 0x76,
	0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x64, 0x49, 0x44, 0x12, 0x3e,
	0x0a, 0x1a, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x43, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x1a, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x73, 0x42, 0x3b,
	0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x75, 0x6c,
	0x70, 0x65, 0x73, 0x2d, 0x66, 0x65, 0x72, 0x72, 0x69, 0x6c, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x61,
	0x74, 0x61, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_protobuf_models_play_year_of_plenty_card_request_proto_rawDescOnce sync.Once
	file_protobuf_models_play_year_of_plenty_card_request_proto_rawDescData = file_protobuf_models_play_year_of_plenty_card_request_proto_rawDesc
)

func file_protobuf_models_play_year_of_plenty_card_request_proto_rawDescGZIP() []byte {
	file_protobuf_models_play_year_of_plenty_card_request_proto_rawDescOnce.Do(func() {
		file_protobuf_models_play_year_of_plenty_card_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_models_play_year_of_plenty_card_request_proto_rawDescData)
	})
	return file_protobuf_models_play_year_of_plenty_card_request_proto_rawDescData
}

var file_protobuf_models_play_year_of_plenty_card_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protobuf_models_play_year_of_plenty_card_request_proto_goTypes = []interface{}{
	(*PlayYearOfPlentyCardRequest)(nil), // 0: models.PlayYearOfPlentyCardRequest
}
var file_protobuf_models_play_year_of_plenty_card_request_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_models_play_year_of_plenty_card_request_proto_init() }
func file_protobuf_models_play_year_of_plenty_card_request_proto_init() {
	if File_protobuf_models_play_year_of_plenty_card_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_models_play_year_of_plenty_card_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayYearOfPlentyCardRequest); i {
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
			RawDescriptor: file_protobuf_models_play_year_of_plenty_card_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_models_play_year_of_plenty_card_request_proto_goTypes,
		DependencyIndexes: file_protobuf_models_play_year_of_plenty_card_request_proto_depIdxs,
		MessageInfos:      file_protobuf_models_play_year_of_plenty_card_request_proto_msgTypes,
	}.Build()
	File_protobuf_models_play_year_of_plenty_card_request_proto = out.File
	file_protobuf_models_play_year_of_plenty_card_request_proto_rawDesc = nil
	file_protobuf_models_play_year_of_plenty_card_request_proto_goTypes = nil
	file_protobuf_models_play_year_of_plenty_card_request_proto_depIdxs = nil
}
