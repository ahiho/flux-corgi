// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: sandbox/sandbox.proto

package sandboxproto

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConfigSandboxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repo          string `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
	ConfigEncoded string `protobuf:"bytes,2,opt,name=config_encoded,json=configEncoded,proto3" json:"config_encoded,omitempty"`
}

func (x *ConfigSandboxRequest) Reset() {
	*x = ConfigSandboxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sandbox_sandbox_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigSandboxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigSandboxRequest) ProtoMessage() {}

func (x *ConfigSandboxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sandbox_sandbox_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigSandboxRequest.ProtoReflect.Descriptor instead.
func (*ConfigSandboxRequest) Descriptor() ([]byte, []int) {
	return file_sandbox_sandbox_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigSandboxRequest) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *ConfigSandboxRequest) GetConfigEncoded() string {
	if x != nil {
		return x.ConfigEncoded
	}
	return ""
}

type ConfigSandboxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConfigSandboxResponse) Reset() {
	*x = ConfigSandboxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sandbox_sandbox_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigSandboxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigSandboxResponse) ProtoMessage() {}

func (x *ConfigSandboxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sandbox_sandbox_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigSandboxResponse.ProtoReflect.Descriptor instead.
func (*ConfigSandboxResponse) Descriptor() ([]byte, []int) {
	return file_sandbox_sandbox_proto_rawDescGZIP(), []int{1}
}

type DeploySandboxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repo          string `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
	PullRequestId int32  `protobuf:"varint,2,opt,name=pull_request_id,json=pullRequestId,proto3" json:"pull_request_id,omitempty"`
	Image         string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *DeploySandboxRequest) Reset() {
	*x = DeploySandboxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sandbox_sandbox_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploySandboxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploySandboxRequest) ProtoMessage() {}

func (x *DeploySandboxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sandbox_sandbox_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploySandboxRequest.ProtoReflect.Descriptor instead.
func (*DeploySandboxRequest) Descriptor() ([]byte, []int) {
	return file_sandbox_sandbox_proto_rawDescGZIP(), []int{2}
}

func (x *DeploySandboxRequest) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *DeploySandboxRequest) GetPullRequestId() int32 {
	if x != nil {
		return x.PullRequestId
	}
	return 0
}

func (x *DeploySandboxRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type DeploySandboxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url   string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	IsNew bool   `protobuf:"varint,2,opt,name=is_new,json=isNew,proto3" json:"is_new,omitempty"`
}

func (x *DeploySandboxResponse) Reset() {
	*x = DeploySandboxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sandbox_sandbox_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploySandboxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploySandboxResponse) ProtoMessage() {}

func (x *DeploySandboxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sandbox_sandbox_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploySandboxResponse.ProtoReflect.Descriptor instead.
func (*DeploySandboxResponse) Descriptor() ([]byte, []int) {
	return file_sandbox_sandbox_proto_rawDescGZIP(), []int{3}
}

func (x *DeploySandboxResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *DeploySandboxResponse) GetIsNew() bool {
	if x != nil {
		return x.IsNew
	}
	return false
}

var File_sandbox_sandbox_proto protoreflect.FileDescriptor

var file_sandbox_sandbox_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x2f, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f,
	0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73,
	0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x72, 0x65, 0x70, 0x6f, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x65,
	0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x68, 0x0a, 0x14, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x53, 0x61,
	0x6e, 0x64, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x65, 0x70, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f,
	0x12, 0x26, 0x0a, 0x0f, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x75, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x40,
	0x0a, 0x15, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f,
	0x6e, 0x65, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x4e, 0x65, 0x77,
	0x32, 0xbe, 0x02, 0x0a, 0x0e, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x94, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x61,
	0x6e, 0x64, 0x62, 0x6f, 0x78, 0x12, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x61,
	0x6e, 0x64, 0x62, 0x6f, 0x78, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x61, 0x6e, 0x64, 0x62,
	0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x92, 0x41, 0x11, 0x62,
	0x0f, 0x0a, 0x0d, 0x0a, 0x09, 0x78, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x6b, 0x65, 0x79, 0x12, 0x00,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x73, 0x61, 0x6e, 0x64,
	0x62, 0x6f, 0x78, 0x3a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x94, 0x01, 0x0a, 0x0d, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x12, 0x28, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73,
	0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x2e, 0x92, 0x41, 0x11, 0x62, 0x0f, 0x0a, 0x0d, 0x0a, 0x09, 0x78, 0x2d, 0x61, 0x70,
	0x69, 0x2d, 0x6b, 0x65, 0x79, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a,
	0x22, 0x0f, 0x2f, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x3a, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x42, 0x2a, 0x5a, 0x28, 0x66, 0x6c, 0x75, 0x78, 0x63, 0x6f, 0x72, 0x67, 0x69, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78,
	0x3b, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sandbox_sandbox_proto_rawDescOnce sync.Once
	file_sandbox_sandbox_proto_rawDescData = file_sandbox_sandbox_proto_rawDesc
)

func file_sandbox_sandbox_proto_rawDescGZIP() []byte {
	file_sandbox_sandbox_proto_rawDescOnce.Do(func() {
		file_sandbox_sandbox_proto_rawDescData = protoimpl.X.CompressGZIP(file_sandbox_sandbox_proto_rawDescData)
	})
	return file_sandbox_sandbox_proto_rawDescData
}

var file_sandbox_sandbox_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sandbox_sandbox_proto_goTypes = []interface{}{
	(*ConfigSandboxRequest)(nil),  // 0: proto.sandboxproto.ConfigSandboxRequest
	(*ConfigSandboxResponse)(nil), // 1: proto.sandboxproto.ConfigSandboxResponse
	(*DeploySandboxRequest)(nil),  // 2: proto.sandboxproto.DeploySandboxRequest
	(*DeploySandboxResponse)(nil), // 3: proto.sandboxproto.DeploySandboxResponse
}
var file_sandbox_sandbox_proto_depIdxs = []int32{
	0, // 0: proto.sandboxproto.SandboxService.ConfigSandbox:input_type -> proto.sandboxproto.ConfigSandboxRequest
	2, // 1: proto.sandboxproto.SandboxService.DeploySandbox:input_type -> proto.sandboxproto.DeploySandboxRequest
	1, // 2: proto.sandboxproto.SandboxService.ConfigSandbox:output_type -> proto.sandboxproto.ConfigSandboxResponse
	3, // 3: proto.sandboxproto.SandboxService.DeploySandbox:output_type -> proto.sandboxproto.DeploySandboxResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sandbox_sandbox_proto_init() }
func file_sandbox_sandbox_proto_init() {
	if File_sandbox_sandbox_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sandbox_sandbox_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigSandboxRequest); i {
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
		file_sandbox_sandbox_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigSandboxResponse); i {
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
		file_sandbox_sandbox_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploySandboxRequest); i {
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
		file_sandbox_sandbox_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploySandboxResponse); i {
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
			RawDescriptor: file_sandbox_sandbox_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sandbox_sandbox_proto_goTypes,
		DependencyIndexes: file_sandbox_sandbox_proto_depIdxs,
		MessageInfos:      file_sandbox_sandbox_proto_msgTypes,
	}.Build()
	File_sandbox_sandbox_proto = out.File
	file_sandbox_sandbox_proto_rawDesc = nil
	file_sandbox_sandbox_proto_goTypes = nil
	file_sandbox_sandbox_proto_depIdxs = nil
}
