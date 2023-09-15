// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: message2connect.proto

package mq_pb

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

type Notify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotifyType int32         `protobuf:"varint,1,opt,name=NotifyType,proto3" json:"NotifyType,omitempty"` //1:消息 2:已读 3:撤回
	MailBox    []*Mailbox    `protobuf:"bytes,2,rep,name=MailBox,proto3" json:"MailBox,omitempty"`
	Message    *Message      `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`
	Common     *NotifyCommon `protobuf:"bytes,4,opt,name=Common,proto3" json:"Common,omitempty"`
}

func (x *Notify) Reset() {
	*x = Notify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message2connect_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notify) ProtoMessage() {}

func (x *Notify) ProtoReflect() protoreflect.Message {
	mi := &file_message2connect_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notify.ProtoReflect.Descriptor instead.
func (*Notify) Descriptor() ([]byte, []int) {
	return file_message2connect_proto_rawDescGZIP(), []int{0}
}

func (x *Notify) GetNotifyType() int32 {
	if x != nil {
		return x.NotifyType
	}
	return 0
}

func (x *Notify) GetMailBox() []*Mailbox {
	if x != nil {
		return x.MailBox
	}
	return nil
}

func (x *Notify) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Notify) GetCommon() *NotifyCommon {
	if x != nil {
		return x.Common
	}
	return nil
}

type Mailbox struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReceiverId int64 `protobuf:"varint,1,opt,name=ReceiverId,proto3" json:"ReceiverId,omitempty"`
	Seq        int64 `protobuf:"varint,2,opt,name=Seq,proto3" json:"Seq,omitempty"`
}

func (x *Mailbox) Reset() {
	*x = Mailbox{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message2connect_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mailbox) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mailbox) ProtoMessage() {}

func (x *Mailbox) ProtoReflect() protoreflect.Message {
	mi := &file_message2connect_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mailbox.ProtoReflect.Descriptor instead.
func (*Mailbox) Descriptor() ([]byte, []int) {
	return file_message2connect_proto_rawDescGZIP(), []int{1}
}

func (x *Mailbox) GetReceiverId() int64 {
	if x != nil {
		return x.ReceiverId
	}
	return 0
}

func (x *Mailbox) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	GroupId     int64  `protobuf:"varint,2,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
	MessageType int32  `protobuf:"varint,3,opt,name=MessageType,proto3" json:"MessageType,omitempty"`
	MessageData string `protobuf:"bytes,4,opt,name=MessageData,proto3" json:"MessageData,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message2connect_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_message2connect_proto_msgTypes[2]
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
	return file_message2connect_proto_rawDescGZIP(), []int{2}
}

func (x *Message) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Message) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *Message) GetMessageType() int32 {
	if x != nil {
		return x.MessageType
	}
	return 0
}

func (x *Message) GetMessageData() string {
	if x != nil {
		return x.MessageData
	}
	return ""
}

type NotifyCommon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderId    int64  `protobuf:"varint,1,opt,name=SenderId,proto3" json:"SenderId,omitempty"`
	ContentType int32  `protobuf:"varint,2,opt,name=ContentType,proto3" json:"ContentType,omitempty"`
	Content     string `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
	CreateAt    int64  `protobuf:"varint,4,opt,name=CreateAt,proto3" json:"CreateAt,omitempty"`
	ClientId    int64  `protobuf:"varint,5,opt,name=ClientId,proto3" json:"ClientId,omitempty"`
}

func (x *NotifyCommon) Reset() {
	*x = NotifyCommon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message2connect_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyCommon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyCommon) ProtoMessage() {}

func (x *NotifyCommon) ProtoReflect() protoreflect.Message {
	mi := &file_message2connect_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyCommon.ProtoReflect.Descriptor instead.
func (*NotifyCommon) Descriptor() ([]byte, []int) {
	return file_message2connect_proto_rawDescGZIP(), []int{3}
}

func (x *NotifyCommon) GetSenderId() int64 {
	if x != nil {
		return x.SenderId
	}
	return 0
}

func (x *NotifyCommon) GetContentType() int32 {
	if x != nil {
		return x.ContentType
	}
	return 0
}

func (x *NotifyCommon) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *NotifyCommon) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *NotifyCommon) GetClientId() int64 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

var File_message2connect_proto protoreflect.FileDescriptor

var file_message2connect_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x71, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb2, 0x01, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1e, 0x0a, 0x0a,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x07,
	0x4d, 0x61, 0x69, 0x6c, 0x42, 0x6f, 0x78, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x6d, 0x71, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78,
	0x52, 0x07, 0x4d, 0x61, 0x69, 0x6c, 0x42, 0x6f, 0x78, 0x12, 0x2b, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x71, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x71, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x06,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x3b, 0x0a, 0x07, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f,
	0x78, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x53, 0x65, 0x71, 0x22, 0x77, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x18,
	0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x9e, 0x01, 0x0a,
	0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x2f, 0x6d, 0x71, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message2connect_proto_rawDescOnce sync.Once
	file_message2connect_proto_rawDescData = file_message2connect_proto_rawDesc
)

func file_message2connect_proto_rawDescGZIP() []byte {
	file_message2connect_proto_rawDescOnce.Do(func() {
		file_message2connect_proto_rawDescData = protoimpl.X.CompressGZIP(file_message2connect_proto_rawDescData)
	})
	return file_message2connect_proto_rawDescData
}

var file_message2connect_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_message2connect_proto_goTypes = []interface{}{
	(*Notify)(nil),       // 0: mq_proto.Notify
	(*Mailbox)(nil),      // 1: mq_proto.Mailbox
	(*Message)(nil),      // 2: mq_proto.Message
	(*NotifyCommon)(nil), // 3: mq_proto.NotifyCommon
}
var file_message2connect_proto_depIdxs = []int32{
	1, // 0: mq_proto.Notify.MailBox:type_name -> mq_proto.Mailbox
	2, // 1: mq_proto.Notify.Message:type_name -> mq_proto.Message
	3, // 2: mq_proto.Notify.Common:type_name -> mq_proto.NotifyCommon
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_message2connect_proto_init() }
func file_message2connect_proto_init() {
	if File_message2connect_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message2connect_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notify); i {
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
		file_message2connect_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mailbox); i {
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
		file_message2connect_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_message2connect_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyCommon); i {
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
			RawDescriptor: file_message2connect_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message2connect_proto_goTypes,
		DependencyIndexes: file_message2connect_proto_depIdxs,
		MessageInfos:      file_message2connect_proto_msgTypes,
	}.Build()
	File_message2connect_proto = out.File
	file_message2connect_proto_rawDesc = nil
	file_message2connect_proto_goTypes = nil
	file_message2connect_proto_depIdxs = nil
}
