// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: message.proto

package message

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

type MessageHandlerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderId    int64  `protobuf:"varint,1,opt,name=SenderId,proto3" json:"SenderId,omitempty"`
	SenderType  int32  `protobuf:"varint,2,opt,name=SenderType,proto3" json:"SenderType,omitempty"` //1:个人 2:群
	GroupId     int64  `protobuf:"varint,3,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
	ReceiverId  int64  `protobuf:"varint,4,opt,name=ReceiverId,proto3" json:"ReceiverId,omitempty"`
	ContentType int32  `protobuf:"varint,5,opt,name=ContentType,proto3" json:"ContentType,omitempty"` //1:文本 2:图片 3:文件
	Content     string `protobuf:"bytes,6,opt,name=Content,proto3" json:"Content,omitempty"`
	ClientId    int64  `protobuf:"varint,7,opt,name=ClientId,proto3" json:"ClientId,omitempty"`
}

func (x *MessageHandlerRequest) Reset() {
	*x = MessageHandlerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageHandlerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageHandlerRequest) ProtoMessage() {}

func (x *MessageHandlerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageHandlerRequest.ProtoReflect.Descriptor instead.
func (*MessageHandlerRequest) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *MessageHandlerRequest) GetSenderId() int64 {
	if x != nil {
		return x.SenderId
	}
	return 0
}

func (x *MessageHandlerRequest) GetSenderType() int32 {
	if x != nil {
		return x.SenderType
	}
	return 0
}

func (x *MessageHandlerRequest) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *MessageHandlerRequest) GetReceiverId() int64 {
	if x != nil {
		return x.ReceiverId
	}
	return 0
}

func (x *MessageHandlerRequest) GetContentType() int32 {
	if x != nil {
		return x.ContentType
	}
	return 0
}

func (x *MessageHandlerRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *MessageHandlerRequest) GetClientId() int64 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

type ReadHandlerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderId  int64  `protobuf:"varint,1,opt,name=SenderId,proto3" json:"SenderId,omitempty"`
	MessageId string `protobuf:"bytes,2,opt,name=MessageId,proto3" json:"MessageId,omitempty"`
	ClientId  int64  `protobuf:"varint,3,opt,name=ClientId,proto3" json:"ClientId,omitempty"`
}

func (x *ReadHandlerRequest) Reset() {
	*x = ReadHandlerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadHandlerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadHandlerRequest) ProtoMessage() {}

func (x *ReadHandlerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadHandlerRequest.ProtoReflect.Descriptor instead.
func (*ReadHandlerRequest) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *ReadHandlerRequest) GetSenderId() int64 {
	if x != nil {
		return x.SenderId
	}
	return 0
}

func (x *ReadHandlerRequest) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *ReadHandlerRequest) GetClientId() int64 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

type WithdrawHandlerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderId   int64  `protobuf:"varint,1,opt,name=SenderId,proto3" json:"SenderId,omitempty"`
	SenderType int32  `protobuf:"varint,2,opt,name=SenderType,proto3" json:"SenderType,omitempty"` //1:个人 2:群
	GroupId    int64  `protobuf:"varint,3,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
	ReceiverId int64  `protobuf:"varint,4,opt,name=ReceiverId,proto3" json:"ReceiverId,omitempty"`
	MessageId  string `protobuf:"bytes,5,opt,name=MessageId,proto3" json:"MessageId,omitempty"`
	ClientId   int64  `protobuf:"varint,6,opt,name=ClientId,proto3" json:"ClientId,omitempty"`
}

func (x *WithdrawHandlerRequest) Reset() {
	*x = WithdrawHandlerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawHandlerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawHandlerRequest) ProtoMessage() {}

func (x *WithdrawHandlerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawHandlerRequest.ProtoReflect.Descriptor instead.
func (*WithdrawHandlerRequest) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *WithdrawHandlerRequest) GetSenderId() int64 {
	if x != nil {
		return x.SenderId
	}
	return 0
}

func (x *WithdrawHandlerRequest) GetSenderType() int32 {
	if x != nil {
		return x.SenderType
	}
	return 0
}

func (x *WithdrawHandlerRequest) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *WithdrawHandlerRequest) GetReceiverId() int64 {
	if x != nil {
		return x.ReceiverId
	}
	return 0
}

func (x *WithdrawHandlerRequest) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *WithdrawHandlerRequest) GetClientId() int64 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xe5, 0x01, 0x0a, 0x15, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x22, 0x6a, 0x0a, 0x12, 0x52, 0x65, 0x61, 0x64, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xc8, 0x01, 0x0a,
	0x16, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x32, 0xcb, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x40, 0x0a, 0x0e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x1e,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3a,
	0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x1b, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x42, 0x0a, 0x0f, 0x57, 0x69,
	0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x1f, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0b,
	0x5a, 0x09, 0x2e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_message_proto_goTypes = []interface{}{
	(*MessageHandlerRequest)(nil),  // 0: message.MessageHandlerRequest
	(*ReadHandlerRequest)(nil),     // 1: message.ReadHandlerRequest
	(*WithdrawHandlerRequest)(nil), // 2: message.WithdrawHandlerRequest
	(*Empty)(nil),                  // 3: message.Empty
}
var file_message_proto_depIdxs = []int32{
	0, // 0: message.Message.MessageHandler:input_type -> message.MessageHandlerRequest
	1, // 1: message.Message.ReadHandler:input_type -> message.ReadHandlerRequest
	2, // 2: message.Message.WithdrawHandler:input_type -> message.WithdrawHandlerRequest
	3, // 3: message.Message.MessageHandler:output_type -> message.Empty
	3, // 4: message.Message.ReadHandler:output_type -> message.Empty
	3, // 5: message.Message.WithdrawHandler:output_type -> message.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageHandlerRequest); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadHandlerRequest); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawHandlerRequest); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
