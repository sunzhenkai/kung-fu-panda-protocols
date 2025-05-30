// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: protos/dumper/http.proto

package dumper

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

type URI struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Full          string                 `protobuf:"bytes,1,opt,name=full,proto3" json:"full,omitempty"`
	Schema        string                 `protobuf:"bytes,100,opt,name=schema,proto3" json:"schema,omitempty"`
	Host          string                 `protobuf:"bytes,101,opt,name=host,proto3" json:"host,omitempty"`
	Port          int32                  `protobuf:"varint,102,opt,name=port,proto3" json:"port,omitempty"`
	Path          string                 `protobuf:"bytes,103,opt,name=path,proto3" json:"path,omitempty"`
	Query         string                 `protobuf:"bytes,104,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *URI) Reset() {
	*x = URI{}
	mi := &file_protos_dumper_http_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *URI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URI) ProtoMessage() {}

func (x *URI) ProtoReflect() protoreflect.Message {
	mi := &file_protos_dumper_http_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URI.ProtoReflect.Descriptor instead.
func (*URI) Descriptor() ([]byte, []int) {
	return file_protos_dumper_http_proto_rawDescGZIP(), []int{0}
}

func (x *URI) GetFull() string {
	if x != nil {
		return x.Full
	}
	return ""
}

func (x *URI) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *URI) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *URI) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *URI) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *URI) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

type HttpRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uri           *URI                   `protobuf:"bytes,100,opt,name=uri,proto3" json:"uri,omitempty"`
	Body          []byte                 `protobuf:"bytes,150,opt,name=body,proto3" json:"body,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HttpRequest) Reset() {
	*x = HttpRequest{}
	mi := &file_protos_dumper_http_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HttpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpRequest) ProtoMessage() {}

func (x *HttpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_dumper_http_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpRequest.ProtoReflect.Descriptor instead.
func (*HttpRequest) Descriptor() ([]byte, []int) {
	return file_protos_dumper_http_proto_rawDescGZIP(), []int{1}
}

func (x *HttpRequest) GetUri() *URI {
	if x != nil {
		return x.Uri
	}
	return nil
}

func (x *HttpRequest) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

var File_protos_dumper_http_proto protoreflect.FileDescriptor

const file_protos_dumper_http_proto_rawDesc = "" +
	"\n" +
	"\x18protos/dumper/http.proto\x12\x06dumper\"\x83\x01\n" +
	"\x03URI\x12\x12\n" +
	"\x04full\x18\x01 \x01(\tR\x04full\x12\x16\n" +
	"\x06schema\x18d \x01(\tR\x06schema\x12\x12\n" +
	"\x04host\x18e \x01(\tR\x04host\x12\x12\n" +
	"\x04port\x18f \x01(\x05R\x04port\x12\x12\n" +
	"\x04path\x18g \x01(\tR\x04path\x12\x14\n" +
	"\x05query\x18h \x01(\tR\x05query\"A\n" +
	"\vHttpRequest\x12\x1d\n" +
	"\x03uri\x18d \x01(\v2\v.dumper.URIR\x03uri\x12\x13\n" +
	"\x04body\x18\x96\x01 \x01(\fR\x04bodyB@Z>github.com/sunzhenkai/kung-fu-panda-protocols/go/dumper;dumperb\x06proto3"

var (
	file_protos_dumper_http_proto_rawDescOnce sync.Once
	file_protos_dumper_http_proto_rawDescData []byte
)

func file_protos_dumper_http_proto_rawDescGZIP() []byte {
	file_protos_dumper_http_proto_rawDescOnce.Do(func() {
		file_protos_dumper_http_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_protos_dumper_http_proto_rawDesc), len(file_protos_dumper_http_proto_rawDesc)))
	})
	return file_protos_dumper_http_proto_rawDescData
}

var file_protos_dumper_http_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_dumper_http_proto_goTypes = []any{
	(*URI)(nil),         // 0: dumper.URI
	(*HttpRequest)(nil), // 1: dumper.HttpRequest
}
var file_protos_dumper_http_proto_depIdxs = []int32{
	0, // 0: dumper.HttpRequest.uri:type_name -> dumper.URI
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_dumper_http_proto_init() }
func file_protos_dumper_http_proto_init() {
	if File_protos_dumper_http_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protos_dumper_http_proto_rawDesc), len(file_protos_dumper_http_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_dumper_http_proto_goTypes,
		DependencyIndexes: file_protos_dumper_http_proto_depIdxs,
		MessageInfos:      file_protos_dumper_http_proto_msgTypes,
	}.Build()
	File_protos_dumper_http_proto = out.File
	file_protos_dumper_http_proto_goTypes = nil
	file_protos_dumper_http_proto_depIdxs = nil
}
