// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: zinc.proto

package protos

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

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zinc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_zinc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_zinc_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type U8 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *U8) Reset() {
	*x = U8{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zinc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *U8) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*U8) ProtoMessage() {}

func (x *U8) ProtoReflect() protoreflect.Message {
	mi := &file_zinc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use U8.ProtoReflect.Descriptor instead.
func (*U8) Descriptor() ([]byte, []int) {
	return file_zinc_proto_rawDescGZIP(), []int{1}
}

func (x *U8) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type U16 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *U16) Reset() {
	*x = U16{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zinc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *U16) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*U16) ProtoMessage() {}

func (x *U16) ProtoReflect() protoreflect.Message {
	mi := &file_zinc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use U16.ProtoReflect.Descriptor instead.
func (*U16) Descriptor() ([]byte, []int) {
	return file_zinc_proto_rawDescGZIP(), []int{2}
}

func (x *U16) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type U248 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *U248) Reset() {
	*x = U248{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zinc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *U248) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*U248) ProtoMessage() {}

func (x *U248) ProtoReflect() protoreflect.Message {
	mi := &file_zinc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use U248.ProtoReflect.Descriptor instead.
func (*U248) Descriptor() ([]byte, []int) {
	return file_zinc_proto_rawDescGZIP(), []int{3}
}

func (x *U248) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender       *Address `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient    *Address `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	TokenAddress *Address `protobuf:"bytes,3,opt,name=token_address,json=tokenAddress,proto3" json:"token_address,omitempty"`
	Amount       *U248    `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zinc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_zinc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_zinc_proto_rawDescGZIP(), []int{4}
}

func (x *Message) GetSender() *Address {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *Message) GetRecipient() *Address {
	if x != nil {
		return x.Recipient
	}
	return nil
}

func (x *Message) GetTokenAddress() *Address {
	if x != nil {
		return x.TokenAddress
	}
	return nil
}

func (x *Message) GetAmount() *U248 {
	if x != nil {
		return x.Amount
	}
	return nil
}

var File_zinc_proto protoreflect.FileDescriptor

var file_zinc_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x7a, 0x69, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x18, 0x0a, 0x02, 0x75,
	0x38, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x19, 0x0a, 0x03, 0x75, 0x31, 0x36, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x1a, 0x0a, 0x04, 0x75, 0x32, 0x34, 0x38, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa1, 0x01, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x09, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65,
	0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1d, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x05, 0x2e, 0x75, 0x32, 0x34, 0x38, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x09, 0x5a, 0x07, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_zinc_proto_rawDescOnce sync.Once
	file_zinc_proto_rawDescData = file_zinc_proto_rawDesc
)

func file_zinc_proto_rawDescGZIP() []byte {
	file_zinc_proto_rawDescOnce.Do(func() {
		file_zinc_proto_rawDescData = protoimpl.X.CompressGZIP(file_zinc_proto_rawDescData)
	})
	return file_zinc_proto_rawDescData
}

var file_zinc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_zinc_proto_goTypes = []interface{}{
	(*Address)(nil), // 0: address
	(*U8)(nil),      // 1: u8
	(*U16)(nil),     // 2: u16
	(*U248)(nil),    // 3: u248
	(*Message)(nil), // 4: Message
}
var file_zinc_proto_depIdxs = []int32{
	0, // 0: Message.sender:type_name -> address
	0, // 1: Message.recipient:type_name -> address
	0, // 2: Message.token_address:type_name -> address
	3, // 3: Message.amount:type_name -> u248
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_zinc_proto_init() }
func file_zinc_proto_init() {
	if File_zinc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_zinc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_zinc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*U8); i {
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
		file_zinc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*U16); i {
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
		file_zinc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*U248); i {
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
		file_zinc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_zinc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_zinc_proto_goTypes,
		DependencyIndexes: file_zinc_proto_depIdxs,
		MessageInfos:      file_zinc_proto_msgTypes,
	}.Build()
	File_zinc_proto = out.File
	file_zinc_proto_rawDesc = nil
	file_zinc_proto_goTypes = nil
	file_zinc_proto_depIdxs = nil
}