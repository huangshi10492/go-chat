// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: websocket.proto

package ws_pb

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

type WsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type int32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *WsRequest) Reset() {
	*x = WsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_websocket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WsRequest) ProtoMessage() {}

func (x *WsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_websocket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WsRequest.ProtoReflect.Descriptor instead.
func (*WsRequest) Descriptor() ([]byte, []int) {
	return file_websocket_proto_rawDescGZIP(), []int{0}
}

func (x *WsRequest) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type WsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type int32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	// Types that are assignable to Data:
	//
	//	*WsResponse_MessageResponse
	//	*WsResponse_MailboxResponse
	Data isWsResponse_Data `protobuf_oneof:"Data"`
}

func (x *WsResponse) Reset() {
	*x = WsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_websocket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WsResponse) ProtoMessage() {}

func (x *WsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_websocket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WsResponse.ProtoReflect.Descriptor instead.
func (*WsResponse) Descriptor() ([]byte, []int) {
	return file_websocket_proto_rawDescGZIP(), []int{1}
}

func (x *WsResponse) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (m *WsResponse) GetData() isWsResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *WsResponse) GetMessageResponse() *MessageResponse {
	if x, ok := x.GetData().(*WsResponse_MessageResponse); ok {
		return x.MessageResponse
	}
	return nil
}

func (x *WsResponse) GetMailboxResponse() *MailboxResponse {
	if x, ok := x.GetData().(*WsResponse_MailboxResponse); ok {
		return x.MailboxResponse
	}
	return nil
}

type isWsResponse_Data interface {
	isWsResponse_Data()
}

type WsResponse_MessageResponse struct {
	MessageResponse *MessageResponse `protobuf:"bytes,2,opt,name=MessageResponse,proto3,oneof"`
}

type WsResponse_MailboxResponse struct {
	MailboxResponse *MailboxResponse `protobuf:"bytes,3,opt,name=MailboxResponse,proto3,oneof"`
}

func (*WsResponse_MessageResponse) isWsResponse_Data() {}

func (*WsResponse_MailboxResponse) isWsResponse_Data() {}

type MessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *Message `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Mailbox *Mailbox `protobuf:"bytes,2,opt,name=Mailbox,proto3" json:"Mailbox,omitempty"`
}

func (x *MessageResponse) Reset() {
	*x = MessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_websocket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageResponse) ProtoMessage() {}

func (x *MessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_websocket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageResponse.ProtoReflect.Descriptor instead.
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return file_websocket_proto_rawDescGZIP(), []int{2}
}

func (x *MessageResponse) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *MessageResponse) GetMailbox() *Mailbox {
	if x != nil {
		return x.Mailbox
	}
	return nil
}

type MailboxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mailbox *Mailbox `protobuf:"bytes,1,opt,name=Mailbox,proto3" json:"Mailbox,omitempty"`
}

func (x *MailboxResponse) Reset() {
	*x = MailboxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_websocket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailboxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailboxResponse) ProtoMessage() {}

func (x *MailboxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_websocket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailboxResponse.ProtoReflect.Descriptor instead.
func (*MailboxResponse) Descriptor() ([]byte, []int) {
	return file_websocket_proto_rawDescGZIP(), []int{3}
}

func (x *MailboxResponse) GetMailbox() *Mailbox {
	if x != nil {
		return x.Mailbox
	}
	return nil
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
		mi := &file_websocket_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_websocket_proto_msgTypes[4]
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
	return file_websocket_proto_rawDescGZIP(), []int{4}
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

type Mailbox struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderId    int64  `protobuf:"varint,1,opt,name=SenderId,proto3" json:"SenderId,omitempty"`
	ReceiverId  int64  `protobuf:"varint,2,opt,name=ReceiverId,proto3" json:"ReceiverId,omitempty"`
	Seq         int64  `protobuf:"varint,3,opt,name=Seq,proto3" json:"Seq,omitempty"`
	ContentType int32  `protobuf:"varint,4,opt,name=ContentType,proto3" json:"ContentType,omitempty"`
	Content     string `protobuf:"bytes,5,opt,name=Content,proto3" json:"Content,omitempty"`
	CreateAt    int64  `protobuf:"varint,6,opt,name=CreateAt,proto3" json:"CreateAt,omitempty"`
	ClientId    int64  `protobuf:"varint,7,opt,name=ClientId,proto3" json:"ClientId,omitempty"`
}

func (x *Mailbox) Reset() {
	*x = Mailbox{}
	if protoimpl.UnsafeEnabled {
		mi := &file_websocket_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mailbox) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mailbox) ProtoMessage() {}

func (x *Mailbox) ProtoReflect() protoreflect.Message {
	mi := &file_websocket_proto_msgTypes[5]
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
	return file_websocket_proto_rawDescGZIP(), []int{5}
}

func (x *Mailbox) GetSenderId() int64 {
	if x != nil {
		return x.SenderId
	}
	return 0
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

func (x *Mailbox) GetContentType() int32 {
	if x != nil {
		return x.ContentType
	}
	return 0
}

func (x *Mailbox) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Mailbox) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Mailbox) GetClientId() int64 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

var File_websocket_proto protoreflect.FileDescriptor

var file_websocket_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x1f, 0x0a, 0x09, 0x57,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0xb6, 0x01, 0x0a,
	0x0a, 0x57, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x45, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0f, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f,
	0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x62,
	0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x4d, 0x61,
	0x69, 0x6c, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x6b, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x52, 0x07, 0x4d, 0x61, 0x69, 0x6c, 0x62,
	0x6f, 0x78, 0x22, 0x3e, 0x0a, 0x0f, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x52, 0x07, 0x4d, 0x61, 0x69, 0x6c, 0x62,
	0x6f, 0x78, 0x22, 0x77, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x18, 0x0a,
	0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0xcb, 0x01, 0x0a, 0x07,
	0x4d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x65, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x53, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x77,
	0x73, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_websocket_proto_rawDescOnce sync.Once
	file_websocket_proto_rawDescData = file_websocket_proto_rawDesc
)

func file_websocket_proto_rawDescGZIP() []byte {
	file_websocket_proto_rawDescOnce.Do(func() {
		file_websocket_proto_rawDescData = protoimpl.X.CompressGZIP(file_websocket_proto_rawDescData)
	})
	return file_websocket_proto_rawDescData
}

var file_websocket_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_websocket_proto_goTypes = []interface{}{
	(*WsRequest)(nil),       // 0: protobuf.WsRequest
	(*WsResponse)(nil),      // 1: protobuf.WsResponse
	(*MessageResponse)(nil), // 2: protobuf.MessageResponse
	(*MailboxResponse)(nil), // 3: protobuf.MailboxResponse
	(*Message)(nil),         // 4: protobuf.Message
	(*Mailbox)(nil),         // 5: protobuf.Mailbox
}
var file_websocket_proto_depIdxs = []int32{
	2, // 0: protobuf.WsResponse.MessageResponse:type_name -> protobuf.MessageResponse
	3, // 1: protobuf.WsResponse.MailboxResponse:type_name -> protobuf.MailboxResponse
	4, // 2: protobuf.MessageResponse.Message:type_name -> protobuf.Message
	5, // 3: protobuf.MessageResponse.Mailbox:type_name -> protobuf.Mailbox
	5, // 4: protobuf.MailboxResponse.Mailbox:type_name -> protobuf.Mailbox
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_websocket_proto_init() }
func file_websocket_proto_init() {
	if File_websocket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_websocket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WsRequest); i {
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
		file_websocket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WsResponse); i {
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
		file_websocket_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageResponse); i {
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
		file_websocket_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailboxResponse); i {
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
		file_websocket_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_websocket_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
	}
	file_websocket_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*WsResponse_MessageResponse)(nil),
		(*WsResponse_MailboxResponse)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_websocket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_websocket_proto_goTypes,
		DependencyIndexes: file_websocket_proto_depIdxs,
		MessageInfos:      file_websocket_proto_msgTypes,
	}.Build()
	File_websocket_proto = out.File
	file_websocket_proto_rawDesc = nil
	file_websocket_proto_goTypes = nil
	file_websocket_proto_depIdxs = nil
}