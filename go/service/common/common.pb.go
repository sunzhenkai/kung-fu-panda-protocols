// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: protos/service/common/common.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HealthCheckResponse_ServiceStatus int32

const (
	HealthCheckResponse_SERVICE_STATUS_UNSPECIFIED HealthCheckResponse_ServiceStatus = 0
	HealthCheckResponse_SERVICE_STATUS_OK          HealthCheckResponse_ServiceStatus = 1
	HealthCheckResponse_SERVICE_STATUS_UNINIT      HealthCheckResponse_ServiceStatus = 2
	HealthCheckResponse_SERVICE_STATUS_ERROR       HealthCheckResponse_ServiceStatus = 3
)

// Enum value maps for HealthCheckResponse_ServiceStatus.
var (
	HealthCheckResponse_ServiceStatus_name = map[int32]string{
		0: "SERVICE_STATUS_UNSPECIFIED",
		1: "SERVICE_STATUS_OK",
		2: "SERVICE_STATUS_UNINIT",
		3: "SERVICE_STATUS_ERROR",
	}
	HealthCheckResponse_ServiceStatus_value = map[string]int32{
		"SERVICE_STATUS_UNSPECIFIED": 0,
		"SERVICE_STATUS_OK":          1,
		"SERVICE_STATUS_UNINIT":      2,
		"SERVICE_STATUS_ERROR":       3,
	}
)

func (x HealthCheckResponse_ServiceStatus) Enum() *HealthCheckResponse_ServiceStatus {
	p := new(HealthCheckResponse_ServiceStatus)
	*p = x
	return p
}

func (x HealthCheckResponse_ServiceStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HealthCheckResponse_ServiceStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_service_common_common_proto_enumTypes[0].Descriptor()
}

func (HealthCheckResponse_ServiceStatus) Type() protoreflect.EnumType {
	return &file_protos_service_common_common_proto_enumTypes[0]
}

func (x HealthCheckResponse_ServiceStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HealthCheckResponse_ServiceStatus.Descriptor instead.
func (HealthCheckResponse_ServiceStatus) EnumDescriptor() ([]byte, []int) {
	return file_protos_service_common_common_proto_rawDescGZIP(), []int{2, 0}
}

type HealthCheckRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Service       string                 `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	mi := &file_protos_service_common_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_common_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_protos_service_common_common_proto_rawDescGZIP(), []int{0}
}

func (x *HealthCheckRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type NodeState struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	StartTimestampS int32                  `protobuf:"varint,1,opt,name=start_timestamp_s,json=startTimestampS,proto3" json:"start_timestamp_s,omitempty"`
	CpuCores        int32                  `protobuf:"varint,100,opt,name=cpu_cores,json=cpuCores,proto3" json:"cpu_cores,omitempty"`
	CpuCoresUsage   int32                  `protobuf:"varint,101,opt,name=cpu_cores_usage,json=cpuCoresUsage,proto3" json:"cpu_cores_usage,omitempty"`
	MemoryUsage     int32                  `protobuf:"varint,102,opt,name=memory_usage,json=memoryUsage,proto3" json:"memory_usage,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *NodeState) Reset() {
	*x = NodeState{}
	mi := &file_protos_service_common_common_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeState) ProtoMessage() {}

func (x *NodeState) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_common_common_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeState.ProtoReflect.Descriptor instead.
func (*NodeState) Descriptor() ([]byte, []int) {
	return file_protos_service_common_common_proto_rawDescGZIP(), []int{1}
}

func (x *NodeState) GetStartTimestampS() int32 {
	if x != nil {
		return x.StartTimestampS
	}
	return 0
}

func (x *NodeState) GetCpuCores() int32 {
	if x != nil {
		return x.CpuCores
	}
	return 0
}

func (x *NodeState) GetCpuCoresUsage() int32 {
	if x != nil {
		return x.CpuCoresUsage
	}
	return 0
}

func (x *NodeState) GetMemoryUsage() int32 {
	if x != nil {
		return x.MemoryUsage
	}
	return 0
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState            `protogen:"open.v1"`
	ServiceStatus HealthCheckResponse_ServiceStatus `protobuf:"varint,1,opt,name=service_status,json=serviceStatus,proto3,enum=common.HealthCheckResponse_ServiceStatus" json:"service_status,omitempty"`
	Message       string                            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	NodeState     *NodeState                        `protobuf:"bytes,3,opt,name=node_state,json=nodeState,proto3" json:"node_state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	mi := &file_protos_service_common_common_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_common_common_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_protos_service_common_common_proto_rawDescGZIP(), []int{2}
}

func (x *HealthCheckResponse) GetServiceStatus() HealthCheckResponse_ServiceStatus {
	if x != nil {
		return x.ServiceStatus
	}
	return HealthCheckResponse_SERVICE_STATUS_UNSPECIFIED
}

func (x *HealthCheckResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *HealthCheckResponse) GetNodeState() *NodeState {
	if x != nil {
		return x.NodeState
	}
	return nil
}

var File_protos_service_common_common_proto protoreflect.FileDescriptor

const file_protos_service_common_common_proto_rawDesc = "" +
	"\n" +
	"\"protos/service/common/common.proto\x12\x06common\".\n" +
	"\x12HealthCheckRequest\x12\x18\n" +
	"\aservice\x18\x01 \x01(\tR\aservice\"\x9f\x01\n" +
	"\tNodeState\x12*\n" +
	"\x11start_timestamp_s\x18\x01 \x01(\x05R\x0fstartTimestampS\x12\x1b\n" +
	"\tcpu_cores\x18d \x01(\x05R\bcpuCores\x12&\n" +
	"\x0fcpu_cores_usage\x18e \x01(\x05R\rcpuCoresUsage\x12!\n" +
	"\fmemory_usage\x18f \x01(\x05R\vmemoryUsage\"\xb0\x02\n" +
	"\x13HealthCheckResponse\x12P\n" +
	"\x0eservice_status\x18\x01 \x01(\x0e2).common.HealthCheckResponse.ServiceStatusR\rserviceStatus\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x120\n" +
	"\n" +
	"node_state\x18\x03 \x01(\v2\x11.common.NodeStateR\tnodeState\"{\n" +
	"\rServiceStatus\x12\x1e\n" +
	"\x1aSERVICE_STATUS_UNSPECIFIED\x10\x00\x12\x15\n" +
	"\x11SERVICE_STATUS_OK\x10\x01\x12\x19\n" +
	"\x15SERVICE_STATUS_UNINIT\x10\x02\x12\x18\n" +
	"\x14SERVICE_STATUS_ERROR\x10\x032Q\n" +
	"\rHealthService\x12@\n" +
	"\x05Check\x12\x1a.common.HealthCheckRequest\x1a\x1b.common.HealthCheckResponseBKZFgithub.com/sunzhenkai/kung-fu-panda-protocols/go/service/common;common\x80\x01\x01b\x06proto3"

var (
	file_protos_service_common_common_proto_rawDescOnce sync.Once
	file_protos_service_common_common_proto_rawDescData []byte
)

func file_protos_service_common_common_proto_rawDescGZIP() []byte {
	file_protos_service_common_common_proto_rawDescOnce.Do(func() {
		file_protos_service_common_common_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_protos_service_common_common_proto_rawDesc), len(file_protos_service_common_common_proto_rawDesc)))
	})
	return file_protos_service_common_common_proto_rawDescData
}

var file_protos_service_common_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_service_common_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_service_common_common_proto_goTypes = []any{
	(HealthCheckResponse_ServiceStatus)(0), // 0: common.HealthCheckResponse.ServiceStatus
	(*HealthCheckRequest)(nil),             // 1: common.HealthCheckRequest
	(*NodeState)(nil),                      // 2: common.NodeState
	(*HealthCheckResponse)(nil),            // 3: common.HealthCheckResponse
}
var file_protos_service_common_common_proto_depIdxs = []int32{
	0, // 0: common.HealthCheckResponse.service_status:type_name -> common.HealthCheckResponse.ServiceStatus
	2, // 1: common.HealthCheckResponse.node_state:type_name -> common.NodeState
	1, // 2: common.HealthService.Check:input_type -> common.HealthCheckRequest
	3, // 3: common.HealthService.Check:output_type -> common.HealthCheckResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_service_common_common_proto_init() }
func file_protos_service_common_common_proto_init() {
	if File_protos_service_common_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protos_service_common_common_proto_rawDesc), len(file_protos_service_common_common_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_service_common_common_proto_goTypes,
		DependencyIndexes: file_protos_service_common_common_proto_depIdxs,
		EnumInfos:         file_protos_service_common_common_proto_enumTypes,
		MessageInfos:      file_protos_service_common_common_proto_msgTypes,
	}.Build()
	File_protos_service_common_common_proto = out.File
	file_protos_service_common_common_proto_goTypes = nil
	file_protos_service_common_common_proto_depIdxs = nil
}
