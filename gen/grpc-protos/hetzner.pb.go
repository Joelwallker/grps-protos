// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: hetzner/hetzner.proto

package grpc_protos

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

type GetServersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens []string `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"` // Список токенов Hetzner
	Labels []string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"` // Список лейблов
}

func (x *GetServersRequest) Reset() {
	*x = GetServersRequest{}
	mi := &file_hetzner_hetzner_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetServersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServersRequest) ProtoMessage() {}

func (x *GetServersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hetzner_hetzner_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServersRequest.ProtoReflect.Descriptor instead.
func (*GetServersRequest) Descriptor() ([]byte, []int) {
	return file_hetzner_hetzner_proto_rawDescGZIP(), []int{0}
}

func (x *GetServersRequest) GetTokens() []string {
	if x != nil {
		return x.Tokens
	}
	return nil
}

func (x *GetServersRequest) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type GetServersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerData map[string]*GetServersResponse_ServerList `protobuf:"bytes,1,rep,name=server_data,json=serverData,proto3" json:"server_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Сопоставление лейблов и серверов
}

func (x *GetServersResponse) Reset() {
	*x = GetServersResponse{}
	mi := &file_hetzner_hetzner_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetServersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServersResponse) ProtoMessage() {}

func (x *GetServersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hetzner_hetzner_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServersResponse.ProtoReflect.Descriptor instead.
func (*GetServersResponse) Descriptor() ([]byte, []int) {
	return file_hetzner_hetzner_proto_rawDescGZIP(), []int{1}
}

func (x *GetServersResponse) GetServerData() map[string]*GetServersResponse_ServerList {
	if x != nil {
		return x.ServerData
	}
	return nil
}

type GetServersResponse_ServerList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Servers []*ServerInfo `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty"`
}

func (x *GetServersResponse_ServerList) Reset() {
	*x = GetServersResponse_ServerList{}
	mi := &file_hetzner_hetzner_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetServersResponse_ServerList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServersResponse_ServerList) ProtoMessage() {}

func (x *GetServersResponse_ServerList) ProtoReflect() protoreflect.Message {
	mi := &file_hetzner_hetzner_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServersResponse_ServerList.ProtoReflect.Descriptor instead.
func (*GetServersResponse_ServerList) Descriptor() ([]byte, []int) {
	return file_hetzner_hetzner_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetServersResponse_ServerList) GetServers() []*ServerInfo {
	if x != nil {
		return x.Servers
	}
	return nil
}

var File_hetzner_hetzner_proto protoreflect.FileDescriptor

var file_hetzner_hetzner_proto_rawDesc = []byte{
	0x0a, 0x15, 0x68, 0x65, 0x74, 0x7a, 0x6e, 0x65, 0x72, 0x2f, 0x68, 0x65, 0x74, 0x7a, 0x6e, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x68, 0x65, 0x74, 0x7a, 0x6e, 0x65, 0x72,
	0x1a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22,
	0x85, 0x02, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x68, 0x65,
	0x74, 0x7a, 0x6e, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x1a, 0x3a, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x2c, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73,
	0x1a, 0x65, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x68, 0x65, 0x74, 0x7a, 0x6e, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x57, 0x0a, 0x0e, 0x48, 0x65, 0x74, 0x7a, 0x6e,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x2e, 0x68, 0x65, 0x74, 0x7a, 0x6e, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x65, 0x74, 0x7a, 0x6e, 0x65, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x0e, 0x5a, 0x0c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hetzner_hetzner_proto_rawDescOnce sync.Once
	file_hetzner_hetzner_proto_rawDescData = file_hetzner_hetzner_proto_rawDesc
)

func file_hetzner_hetzner_proto_rawDescGZIP() []byte {
	file_hetzner_hetzner_proto_rawDescOnce.Do(func() {
		file_hetzner_hetzner_proto_rawDescData = protoimpl.X.CompressGZIP(file_hetzner_hetzner_proto_rawDescData)
	})
	return file_hetzner_hetzner_proto_rawDescData
}

var file_hetzner_hetzner_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_hetzner_hetzner_proto_goTypes = []any{
	(*GetServersRequest)(nil),             // 0: hetzner.GetServersRequest
	(*GetServersResponse)(nil),            // 1: hetzner.GetServersResponse
	(*GetServersResponse_ServerList)(nil), // 2: hetzner.GetServersResponse.ServerList
	nil,                                   // 3: hetzner.GetServersResponse.ServerDataEntry
	(*ServerInfo)(nil),                    // 4: common.ServerInfo
}
var file_hetzner_hetzner_proto_depIdxs = []int32{
	3, // 0: hetzner.GetServersResponse.server_data:type_name -> hetzner.GetServersResponse.ServerDataEntry
	4, // 1: hetzner.GetServersResponse.ServerList.servers:type_name -> common.ServerInfo
	2, // 2: hetzner.GetServersResponse.ServerDataEntry.value:type_name -> hetzner.GetServersResponse.ServerList
	0, // 3: hetzner.HetznerService.GetServers:input_type -> hetzner.GetServersRequest
	1, // 4: hetzner.HetznerService.GetServers:output_type -> hetzner.GetServersResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_hetzner_hetzner_proto_init() }
func file_hetzner_hetzner_proto_init() {
	if File_hetzner_hetzner_proto != nil {
		return
	}
	file_common_server_info_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hetzner_hetzner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hetzner_hetzner_proto_goTypes,
		DependencyIndexes: file_hetzner_hetzner_proto_depIdxs,
		MessageInfos:      file_hetzner_hetzner_proto_msgTypes,
	}.Build()
	File_hetzner_hetzner_proto = out.File
	file_hetzner_hetzner_proto_rawDesc = nil
	file_hetzner_hetzner_proto_goTypes = nil
	file_hetzner_hetzner_proto_depIdxs = nil
}