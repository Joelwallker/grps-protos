// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: reg_ru/reg_ru.proto

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

type GetResourcesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens []string `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"` // Список токенов Reg.ru
	Labels []string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"` // Список лейблов
}

func (x *GetResourcesRequest) Reset() {
	*x = GetResourcesRequest{}
	mi := &file_reg_ru_reg_ru_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResourcesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourcesRequest) ProtoMessage() {}

func (x *GetResourcesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reg_ru_reg_ru_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourcesRequest.ProtoReflect.Descriptor instead.
func (*GetResourcesRequest) Descriptor() ([]byte, []int) {
	return file_reg_ru_reg_ru_proto_rawDescGZIP(), []int{0}
}

func (x *GetResourcesRequest) GetTokens() []string {
	if x != nil {
		return x.Tokens
	}
	return nil
}

func (x *GetResourcesRequest) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type GetResourcesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceData map[string]*GetResourcesResponse_ResourceList `protobuf:"bytes,1,rep,name=resource_data,json=resourceData,proto3" json:"resource_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Сопоставление лейблов и ресурсов
}

func (x *GetResourcesResponse) Reset() {
	*x = GetResourcesResponse{}
	mi := &file_reg_ru_reg_ru_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResourcesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourcesResponse) ProtoMessage() {}

func (x *GetResourcesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_reg_ru_reg_ru_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourcesResponse.ProtoReflect.Descriptor instead.
func (*GetResourcesResponse) Descriptor() ([]byte, []int) {
	return file_reg_ru_reg_ru_proto_rawDescGZIP(), []int{1}
}

func (x *GetResourcesResponse) GetResourceData() map[string]*GetResourcesResponse_ResourceList {
	if x != nil {
		return x.ResourceData
	}
	return nil
}

type GetResourcesResponse_ResourceList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resources []*ServerInfo `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *GetResourcesResponse_ResourceList) Reset() {
	*x = GetResourcesResponse_ResourceList{}
	mi := &file_reg_ru_reg_ru_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResourcesResponse_ResourceList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourcesResponse_ResourceList) ProtoMessage() {}

func (x *GetResourcesResponse_ResourceList) ProtoReflect() protoreflect.Message {
	mi := &file_reg_ru_reg_ru_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourcesResponse_ResourceList.ProtoReflect.Descriptor instead.
func (*GetResourcesResponse_ResourceList) Descriptor() ([]byte, []int) {
	return file_reg_ru_reg_ru_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetResourcesResponse_ResourceList) GetResources() []*ServerInfo {
	if x != nil {
		return x.Resources
	}
	return nil
}

var File_reg_ru_reg_ru_proto protoreflect.FileDescriptor

var file_reg_ru_reg_ru_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x65, 0x67, 0x5f, 0x72, 0x75, 0x2f, 0x72, 0x65, 0x67, 0x5f, 0x72, 0x75, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x65, 0x67, 0x5f, 0x72, 0x75, 0x1a, 0x18, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0x99,
	0x02, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x72, 0x65, 0x67, 0x5f, 0x72, 0x75, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x40, 0x0a, 0x0c,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x09,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x6a,
	0x0a, 0x11, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x72, 0x65, 0x67, 0x5f, 0x72, 0x75, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x59, 0x0a, 0x0c, 0x52, 0x65,
	0x67, 0x52, 0x75, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x67,
	0x5f, 0x72, 0x75, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x65, 0x67, 0x5f, 0x72, 0x75,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reg_ru_reg_ru_proto_rawDescOnce sync.Once
	file_reg_ru_reg_ru_proto_rawDescData = file_reg_ru_reg_ru_proto_rawDesc
)

func file_reg_ru_reg_ru_proto_rawDescGZIP() []byte {
	file_reg_ru_reg_ru_proto_rawDescOnce.Do(func() {
		file_reg_ru_reg_ru_proto_rawDescData = protoimpl.X.CompressGZIP(file_reg_ru_reg_ru_proto_rawDescData)
	})
	return file_reg_ru_reg_ru_proto_rawDescData
}

var file_reg_ru_reg_ru_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_reg_ru_reg_ru_proto_goTypes = []any{
	(*GetResourcesRequest)(nil),               // 0: reg_ru.GetResourcesRequest
	(*GetResourcesResponse)(nil),              // 1: reg_ru.GetResourcesResponse
	(*GetResourcesResponse_ResourceList)(nil), // 2: reg_ru.GetResourcesResponse.ResourceList
	nil,                // 3: reg_ru.GetResourcesResponse.ResourceDataEntry
	(*ServerInfo)(nil), // 4: common.ServerInfo
}
var file_reg_ru_reg_ru_proto_depIdxs = []int32{
	3, // 0: reg_ru.GetResourcesResponse.resource_data:type_name -> reg_ru.GetResourcesResponse.ResourceDataEntry
	4, // 1: reg_ru.GetResourcesResponse.ResourceList.resources:type_name -> common.ServerInfo
	2, // 2: reg_ru.GetResourcesResponse.ResourceDataEntry.value:type_name -> reg_ru.GetResourcesResponse.ResourceList
	0, // 3: reg_ru.RegRuService.GetResources:input_type -> reg_ru.GetResourcesRequest
	1, // 4: reg_ru.RegRuService.GetResources:output_type -> reg_ru.GetResourcesResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_reg_ru_reg_ru_proto_init() }
func file_reg_ru_reg_ru_proto_init() {
	if File_reg_ru_reg_ru_proto != nil {
		return
	}
	file_common_server_info_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_reg_ru_reg_ru_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reg_ru_reg_ru_proto_goTypes,
		DependencyIndexes: file_reg_ru_reg_ru_proto_depIdxs,
		MessageInfos:      file_reg_ru_reg_ru_proto_msgTypes,
	}.Build()
	File_reg_ru_reg_ru_proto = out.File
	file_reg_ru_reg_ru_proto_rawDesc = nil
	file_reg_ru_reg_ru_proto_goTypes = nil
	file_reg_ru_reg_ru_proto_depIdxs = nil
}
