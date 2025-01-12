// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: cloudflare/cloudflare.proto

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

type GetZonesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens []string `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"` // Список токенов Cloudflare
	Labels []string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"` // Список лейблов
}

func (x *GetZonesRequest) Reset() {
	*x = GetZonesRequest{}
	mi := &file_cloudflare_cloudflare_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetZonesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetZonesRequest) ProtoMessage() {}

func (x *GetZonesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloudflare_cloudflare_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetZonesRequest.ProtoReflect.Descriptor instead.
func (*GetZonesRequest) Descriptor() ([]byte, []int) {
	return file_cloudflare_cloudflare_proto_rawDescGZIP(), []int{0}
}

func (x *GetZonesRequest) GetTokens() []string {
	if x != nil {
		return x.Tokens
	}
	return nil
}

func (x *GetZonesRequest) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type GetZonesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ZoneData map[string]*GetZonesResponse_ZoneList `protobuf:"bytes,1,rep,name=zone_data,json=zoneData,proto3" json:"zone_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Сопоставление лейблов и зон
}

func (x *GetZonesResponse) Reset() {
	*x = GetZonesResponse{}
	mi := &file_cloudflare_cloudflare_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetZonesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetZonesResponse) ProtoMessage() {}

func (x *GetZonesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloudflare_cloudflare_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetZonesResponse.ProtoReflect.Descriptor instead.
func (*GetZonesResponse) Descriptor() ([]byte, []int) {
	return file_cloudflare_cloudflare_proto_rawDescGZIP(), []int{1}
}

func (x *GetZonesResponse) GetZoneData() map[string]*GetZonesResponse_ZoneList {
	if x != nil {
		return x.ZoneData
	}
	return nil
}

type GetZonesResponse_ZoneList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Zones []*ServerInfo `protobuf:"bytes,1,rep,name=zones,proto3" json:"zones,omitempty"`
}

func (x *GetZonesResponse_ZoneList) Reset() {
	*x = GetZonesResponse_ZoneList{}
	mi := &file_cloudflare_cloudflare_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetZonesResponse_ZoneList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetZonesResponse_ZoneList) ProtoMessage() {}

func (x *GetZonesResponse_ZoneList) ProtoReflect() protoreflect.Message {
	mi := &file_cloudflare_cloudflare_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetZonesResponse_ZoneList.ProtoReflect.Descriptor instead.
func (*GetZonesResponse_ZoneList) Descriptor() ([]byte, []int) {
	return file_cloudflare_cloudflare_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetZonesResponse_ZoneList) GetZones() []*ServerInfo {
	if x != nil {
		return x.Zones
	}
	return nil
}

var File_cloudflare_cloudflare_proto protoreflect.FileDescriptor

var file_cloudflare_cloudflare_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6c, 0x61, 0x72, 0x65, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x66, 0x6c, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6c, 0x61, 0x72, 0x65, 0x1a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0xf5, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x5a, 0x6f,
	0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x09, 0x7a,
	0x6f, 0x6e, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6c, 0x61, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x5a,
	0x6f, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x5a, 0x6f, 0x6e,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x7a, 0x6f, 0x6e, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x1a, 0x34, 0x0a, 0x08, 0x5a, 0x6f, 0x6e, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x28, 0x0a, 0x05, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x05, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x1a, 0x62, 0x0a, 0x0d, 0x5a, 0x6f,
	0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3b, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6c, 0x61, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x5a, 0x6f, 0x6e,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x5a,
	0x0a, 0x11, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6c, 0x61, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x73, 0x12,
	0x1b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6c, 0x61, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x5a, 0x6f, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6c, 0x61, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x5a, 0x6f, 0x6e,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_cloudflare_cloudflare_proto_rawDescOnce sync.Once
	file_cloudflare_cloudflare_proto_rawDescData = file_cloudflare_cloudflare_proto_rawDesc
)

func file_cloudflare_cloudflare_proto_rawDescGZIP() []byte {
	file_cloudflare_cloudflare_proto_rawDescOnce.Do(func() {
		file_cloudflare_cloudflare_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloudflare_cloudflare_proto_rawDescData)
	})
	return file_cloudflare_cloudflare_proto_rawDescData
}

var file_cloudflare_cloudflare_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cloudflare_cloudflare_proto_goTypes = []any{
	(*GetZonesRequest)(nil),           // 0: cloudflare.GetZonesRequest
	(*GetZonesResponse)(nil),          // 1: cloudflare.GetZonesResponse
	(*GetZonesResponse_ZoneList)(nil), // 2: cloudflare.GetZonesResponse.ZoneList
	nil,                               // 3: cloudflare.GetZonesResponse.ZoneDataEntry
	(*ServerInfo)(nil),                // 4: common.ServerInfo
}
var file_cloudflare_cloudflare_proto_depIdxs = []int32{
	3, // 0: cloudflare.GetZonesResponse.zone_data:type_name -> cloudflare.GetZonesResponse.ZoneDataEntry
	4, // 1: cloudflare.GetZonesResponse.ZoneList.zones:type_name -> common.ServerInfo
	2, // 2: cloudflare.GetZonesResponse.ZoneDataEntry.value:type_name -> cloudflare.GetZonesResponse.ZoneList
	0, // 3: cloudflare.CloudflareService.GetZones:input_type -> cloudflare.GetZonesRequest
	1, // 4: cloudflare.CloudflareService.GetZones:output_type -> cloudflare.GetZonesResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_cloudflare_cloudflare_proto_init() }
func file_cloudflare_cloudflare_proto_init() {
	if File_cloudflare_cloudflare_proto != nil {
		return
	}
	file_common_server_info_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cloudflare_cloudflare_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cloudflare_cloudflare_proto_goTypes,
		DependencyIndexes: file_cloudflare_cloudflare_proto_depIdxs,
		MessageInfos:      file_cloudflare_cloudflare_proto_msgTypes,
	}.Build()
	File_cloudflare_cloudflare_proto = out.File
	file_cloudflare_cloudflare_proto_rawDesc = nil
	file_cloudflare_cloudflare_proto_goTypes = nil
	file_cloudflare_cloudflare_proto_depIdxs = nil
}
