// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: lib/zinctx/src/zinctx.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ValueOneOf:
	//	*Value_Intval
	//	*Value_Stringval
	ValueOneOf isValue_ValueOneOf `protobuf_oneof:"ValueOneOf"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_zinctx_src_zinctx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_lib_zinctx_src_zinctx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_lib_zinctx_src_zinctx_proto_rawDescGZIP(), []int{0}
}

func (m *Value) GetValueOneOf() isValue_ValueOneOf {
	if m != nil {
		return m.ValueOneOf
	}
	return nil
}

func (x *Value) GetIntval() int32 {
	if x, ok := x.GetValueOneOf().(*Value_Intval); ok {
		return x.Intval
	}
	return 0
}

func (x *Value) GetStringval() string {
	if x, ok := x.GetValueOneOf().(*Value_Stringval); ok {
		return x.Stringval
	}
	return ""
}

type isValue_ValueOneOf interface {
	isValue_ValueOneOf()
}

type Value_Intval struct {
	Intval int32 `protobuf:"varint,1,opt,name=intval,proto3,oneof"`
}

type Value_Stringval struct {
	Stringval string `protobuf:"bytes,2,opt,name=stringval,proto3,oneof"`
}

func (*Value_Intval) isValue_ValueOneOf() {}

func (*Value_Stringval) isValue_ValueOneOf() {}

type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address   string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Method    string     `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Arguments *anypb.Any `protobuf:"bytes,3,opt,name=arguments,proto3" json:"arguments,omitempty"`
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_zinctx_src_zinctx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lib_zinctx_src_zinctx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_lib_zinctx_src_zinctx_proto_rawDescGZIP(), []int{1}
}

func (x *QueryRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *QueryRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *QueryRequest) GetArguments() *anypb.Any {
	if x != nil {
		return x.Arguments
	}
	return nil
}

type QueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output *Value `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *QueryResponse) Reset() {
	*x = QueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_zinctx_src_zinctx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResponse) ProtoMessage() {}

func (x *QueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lib_zinctx_src_zinctx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResponse.ProtoReflect.Descriptor instead.
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return file_lib_zinctx_src_zinctx_proto_rawDescGZIP(), []int{2}
}

func (x *QueryResponse) GetOutput() *Value {
	if x != nil {
		return x.Output
	}
	return nil
}

type CallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address   string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Method    string     `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Arguments *anypb.Any `protobuf:"bytes,3,opt,name=arguments,proto3" json:"arguments,omitempty"`
}

func (x *CallRequest) Reset() {
	*x = CallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_zinctx_src_zinctx_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallRequest) ProtoMessage() {}

func (x *CallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lib_zinctx_src_zinctx_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallRequest.ProtoReflect.Descriptor instead.
func (*CallRequest) Descriptor() ([]byte, []int) {
	return file_lib_zinctx_src_zinctx_proto_rawDescGZIP(), []int{3}
}

func (x *CallRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CallRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *CallRequest) GetArguments() *anypb.Any {
	if x != nil {
		return x.Arguments
	}
	return nil
}

type CallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output *Value `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *CallResponse) Reset() {
	*x = CallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_zinctx_src_zinctx_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallResponse) ProtoMessage() {}

func (x *CallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lib_zinctx_src_zinctx_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallResponse.ProtoReflect.Descriptor instead.
func (*CallResponse) Descriptor() ([]byte, []int) {
	return file_lib_zinctx_src_zinctx_proto_rawDescGZIP(), []int{4}
}

func (x *CallResponse) GetOutput() *Value {
	if x != nil {
		return x.Output
	}
	return nil
}

var File_lib_zinctx_src_zinctx_proto protoreflect.FileDescriptor

var file_lib_zinctx_src_zinctx_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6c, 0x69, 0x62, 0x2f, 0x7a, 0x69, 0x6e, 0x63, 0x74, 0x78, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x7a, 0x69, 0x6e, 0x63, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x18, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x00, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x76, 0x61, 0x6c, 0x12, 0x1e, 0x0a, 0x09, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x09, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x76, 0x61, 0x6c, 0x42, 0x0c, 0x0a, 0x0a, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x22, 0x74, 0x0a, 0x0c, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x32, 0x0a, 0x09, 0x61,
	0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0x2f, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1e, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x06, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x22, 0x73, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x32, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x2e, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x09, 0x5a, 0x07, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lib_zinctx_src_zinctx_proto_rawDescOnce sync.Once
	file_lib_zinctx_src_zinctx_proto_rawDescData = file_lib_zinctx_src_zinctx_proto_rawDesc
)

func file_lib_zinctx_src_zinctx_proto_rawDescGZIP() []byte {
	file_lib_zinctx_src_zinctx_proto_rawDescOnce.Do(func() {
		file_lib_zinctx_src_zinctx_proto_rawDescData = protoimpl.X.CompressGZIP(file_lib_zinctx_src_zinctx_proto_rawDescData)
	})
	return file_lib_zinctx_src_zinctx_proto_rawDescData
}

var file_lib_zinctx_src_zinctx_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_lib_zinctx_src_zinctx_proto_goTypes = []interface{}{
	(*Value)(nil),         // 0: Value
	(*QueryRequest)(nil),  // 1: QueryRequest
	(*QueryResponse)(nil), // 2: QueryResponse
	(*CallRequest)(nil),   // 3: CallRequest
	(*CallResponse)(nil),  // 4: CallResponse
	(*anypb.Any)(nil),     // 5: google.protobuf.Any
}
var file_lib_zinctx_src_zinctx_proto_depIdxs = []int32{
	5, // 0: QueryRequest.arguments:type_name -> google.protobuf.Any
	0, // 1: QueryResponse.output:type_name -> Value
	5, // 2: CallRequest.arguments:type_name -> google.protobuf.Any
	0, // 3: CallResponse.output:type_name -> Value
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_lib_zinctx_src_zinctx_proto_init() }
func file_lib_zinctx_src_zinctx_proto_init() {
	if File_lib_zinctx_src_zinctx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lib_zinctx_src_zinctx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_lib_zinctx_src_zinctx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRequest); i {
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
		file_lib_zinctx_src_zinctx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryResponse); i {
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
		file_lib_zinctx_src_zinctx_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallRequest); i {
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
		file_lib_zinctx_src_zinctx_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallResponse); i {
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
	file_lib_zinctx_src_zinctx_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Value_Intval)(nil),
		(*Value_Stringval)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_lib_zinctx_src_zinctx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lib_zinctx_src_zinctx_proto_goTypes,
		DependencyIndexes: file_lib_zinctx_src_zinctx_proto_depIdxs,
		MessageInfos:      file_lib_zinctx_src_zinctx_proto_msgTypes,
	}.Build()
	File_lib_zinctx_src_zinctx_proto = out.File
	file_lib_zinctx_src_zinctx_proto_rawDesc = nil
	file_lib_zinctx_src_zinctx_proto_goTypes = nil
	file_lib_zinctx_src_zinctx_proto_depIdxs = nil
}
