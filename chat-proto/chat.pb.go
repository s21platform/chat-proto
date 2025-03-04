// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.1
// source: chat.proto

package chat_proto

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

type ChatEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChatEmpty) Reset() {
	*x = ChatEmpty{}
	mi := &file_chat_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatEmpty) ProtoMessage() {}

func (x *ChatEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatEmpty.ProtoReflect.Descriptor instead.
func (*ChatEmpty) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

type CreatePrivateChatIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanionUuid string `protobuf:"bytes,1,opt,name=companion_uuid,json=companionUuid,proto3" json:"companion_uuid,omitempty"` // uuid второго пользователя, с которым будет идти переписка
}

func (x *CreatePrivateChatIn) Reset() {
	*x = CreatePrivateChatIn{}
	mi := &file_chat_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePrivateChatIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePrivateChatIn) ProtoMessage() {}

func (x *CreatePrivateChatIn) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePrivateChatIn.ProtoReflect.Descriptor instead.
func (*CreatePrivateChatIn) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePrivateChatIn) GetCompanionUuid() string {
	if x != nil {
		return x.CompanionUuid
	}
	return ""
}

type CreatePrivateChatOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewChatUuid string `protobuf:"bytes,1,opt,name=new_chat_uuid,json=newChatUuid,proto3" json:"new_chat_uuid,omitempty"` // uuid созданного чата
}

func (x *CreatePrivateChatOut) Reset() {
	*x = CreatePrivateChatOut{}
	mi := &file_chat_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePrivateChatOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePrivateChatOut) ProtoMessage() {}

func (x *CreatePrivateChatOut) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePrivateChatOut.ProtoReflect.Descriptor instead.
func (*CreatePrivateChatOut) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePrivateChatOut) GetNewChatUuid() string {
	if x != nil {
		return x.NewChatUuid
	}
	return ""
}

type Chat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatUuid  string `protobuf:"bytes,1,opt,name=chat_uuid,json=chatUuid,proto3" json:"chat_uuid,omitempty"`    // UUID чата
	CreatedAt string `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"` // Время создания чата
	DeletedAt string `protobuf:"bytes,3,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"` // Время удаления чата
	DeletedBy string `protobuf:"bytes,4,opt,name=deleted_by,json=deletedBy,proto3" json:"deleted_by,omitempty"` // UUID пользователя, удалившего чат
}

func (x *Chat) Reset() {
	*x = Chat{}
	mi := &file_chat_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chat) ProtoMessage() {}

func (x *Chat) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chat.ProtoReflect.Descriptor instead.
func (*Chat) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{3}
}

func (x *Chat) GetChatUuid() string {
	if x != nil {
		return x.ChatUuid
	}
	return ""
}

func (x *Chat) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Chat) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *Chat) GetDeletedBy() string {
	if x != nil {
		return x.DeletedBy
	}
	return ""
}

type GetChatsOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chats []*Chat `protobuf:"bytes,1,rep,name=chats,proto3" json:"chats,omitempty"` // Список чатов
}

func (x *GetChatsOut) Reset() {
	*x = GetChatsOut{}
	mi := &file_chat_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetChatsOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatsOut) ProtoMessage() {}

func (x *GetChatsOut) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatsOut.ProtoReflect.Descriptor instead.
func (*GetChatsOut) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{4}
}

func (x *GetChatsOut) GetChats() []*Chat {
	if x != nil {
		return x.Chats
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid       string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`                               // uuid пользователя
	Content    string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`                         // само сообщение
	SentAt     string `protobuf:"bytes,3,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`             // время отправки
	UpdatedAt  string `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`    // время обновления
	RootUuid   string `protobuf:"bytes,5,opt,name=root_uuid,json=rootUuid,proto3" json:"root_uuid,omitempty"`       // uuid корневого сообщения
	ParentUuid string `protobuf:"bytes,6,opt,name=parent_uuid,json=parentUuid,proto3" json:"parent_uuid,omitempty"` // uuid сообщения, на которое идет прямой ответ
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_chat_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[5]
	if x != nil {
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
	return file_chat_proto_rawDescGZIP(), []int{5}
}

func (x *Message) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Message) GetSentAt() string {
	if x != nil {
		return x.SentAt
	}
	return ""
}

func (x *Message) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Message) GetRootUuid() string {
	if x != nil {
		return x.RootUuid
	}
	return ""
}

func (x *Message) GetParentUuid() string {
	if x != nil {
		return x.ParentUuid
	}
	return ""
}

type GetPrivateRecentMessagesIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatUuid string `protobuf:"bytes,1,opt,name=chat_uuid,json=chatUuid,proto3" json:"chat_uuid,omitempty"` // uuid чата, из которого достаем сообщения
}

func (x *GetPrivateRecentMessagesIn) Reset() {
	*x = GetPrivateRecentMessagesIn{}
	mi := &file_chat_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPrivateRecentMessagesIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPrivateRecentMessagesIn) ProtoMessage() {}

func (x *GetPrivateRecentMessagesIn) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPrivateRecentMessagesIn.ProtoReflect.Descriptor instead.
func (*GetPrivateRecentMessagesIn) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{6}
}

func (x *GetPrivateRecentMessagesIn) GetChatUuid() string {
	if x != nil {
		return x.ChatUuid
	}
	return ""
}

type GetPrivateRecentMessagesOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"` // список сообщений
}

func (x *GetPrivateRecentMessagesOut) Reset() {
	*x = GetPrivateRecentMessagesOut{}
	mi := &file_chat_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPrivateRecentMessagesOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPrivateRecentMessagesOut) ProtoMessage() {}

func (x *GetPrivateRecentMessagesOut) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPrivateRecentMessagesOut.ProtoReflect.Descriptor instead.
func (*GetPrivateRecentMessagesOut) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{7}
}

func (x *GetPrivateRecentMessagesOut) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

type DeleteMessageIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UuidMessage string `protobuf:"bytes,1,opt,name=uuid_message,json=uuidMessage,proto3" json:"uuid_message,omitempty"` // uuid сообщения
	Mode        string `protobuf:"bytes,2,opt,name=mode,proto3" json:"mode,omitempty"`                                  // тип удаления: у всех или у себя
}

func (x *DeleteMessageIn) Reset() {
	*x = DeleteMessageIn{}
	mi := &file_chat_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteMessageIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMessageIn) ProtoMessage() {}

func (x *DeleteMessageIn) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMessageIn.ProtoReflect.Descriptor instead.
func (*DeleteMessageIn) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteMessageIn) GetUuidMessage() string {
	if x != nil {
		return x.UuidMessage
	}
	return ""
}

func (x *DeleteMessageIn) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

type DeleteMessageOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeletionStatus bool `protobuf:"varint,1,opt,name=deletion_status,json=deletionStatus,proto3" json:"deletion_status,omitempty"` // статус удаления
}

func (x *DeleteMessageOut) Reset() {
	*x = DeleteMessageOut{}
	mi := &file_chat_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteMessageOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMessageOut) ProtoMessage() {}

func (x *DeleteMessageOut) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMessageOut.ProtoReflect.Descriptor instead.
func (*DeleteMessageOut) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteMessageOut) GetDeletionStatus() bool {
	if x != nil {
		return x.DeletionStatus
	}
	return false
}

type EditMessageIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UuidMessage string `protobuf:"bytes,1,opt,name=uuid_message,json=uuidMessage,proto3" json:"uuid_message,omitempty"` // uuid сообщения
	NewContent  string `protobuf:"bytes,2,opt,name=new_content,json=newContent,proto3" json:"new_content,omitempty"`    // новый текст сообщения
}

func (x *EditMessageIn) Reset() {
	*x = EditMessageIn{}
	mi := &file_chat_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EditMessageIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditMessageIn) ProtoMessage() {}

func (x *EditMessageIn) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditMessageIn.ProtoReflect.Descriptor instead.
func (*EditMessageIn) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{10}
}

func (x *EditMessageIn) GetUuidMessage() string {
	if x != nil {
		return x.UuidMessage
	}
	return ""
}

func (x *EditMessageIn) GetNewContent() string {
	if x != nil {
		return x.NewContent
	}
	return ""
}

type EditMessageOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UuidMessage string `protobuf:"bytes,1,opt,name=uuid_message,json=uuidMessage,proto3" json:"uuid_message,omitempty"` // uuid измененного сообщения
	NewContent  string `protobuf:"bytes,2,opt,name=new_content,json=newContent,proto3" json:"new_content,omitempty"`    // новый текст сообщения
}

func (x *EditMessageOut) Reset() {
	*x = EditMessageOut{}
	mi := &file_chat_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EditMessageOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditMessageOut) ProtoMessage() {}

func (x *EditMessageOut) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditMessageOut.ProtoReflect.Descriptor instead.
func (*EditMessageOut) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{11}
}

func (x *EditMessageOut) GetUuidMessage() string {
	if x != nil {
		return x.UuidMessage
	}
	return ""
}

func (x *EditMessageOut) GetNewContent() string {
	if x != nil {
		return x.NewContent
	}
	return ""
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0b, 0x0a, 0x09,
	0x43, 0x68, 0x61, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3c, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e,
	0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x69, 0x6f, 0x6e, 0x55, 0x75, 0x69, 0x64, 0x22, 0x3a, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x4f, 0x75, 0x74, 0x12,
	0x22, 0x0a, 0x0d, 0x6e, 0x65, 0x77, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x43, 0x68, 0x61, 0x74, 0x55,
	0x75, 0x69, 0x64, 0x22, 0x80, 0x01, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x68, 0x61, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x68, 0x61, 0x74, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22, 0x2a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61,
	0x74, 0x73, 0x4f, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x05, 0x63, 0x68, 0x61,
	0x74, 0x73, 0x22, 0xad, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x73, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x6e, 0x74, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x74, 0x55, 0x75, 0x69,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x75,
	0x69, 0x64, 0x22, 0x39, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x49, 0x6e,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x61, 0x74, 0x55, 0x75, 0x69, 0x64, 0x22, 0x43, 0x0a,
	0x1b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x12, 0x24, 0x0a, 0x08,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x22, 0x48, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x75, 0x69, 0x64, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x75, 0x69,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x3b, 0x0a, 0x10,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x75, 0x74,
	0x12, 0x27, 0x0a, 0x0f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x53, 0x0a, 0x0d, 0x45, 0x64, 0x69,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x75,
	0x69, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x75, 0x75, 0x69, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x6e, 0x65, 0x77, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6e, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x54,
	0x0a, 0x0e, 0x45, 0x64, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x75, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x75, 0x75, 0x69, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x75, 0x69, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x65, 0x77, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x32, 0xbc, 0x02, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x1a,
	0x15, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43,
	0x68, 0x61, 0x74, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x74, 0x73, 0x12, 0x0a, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x0c, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x73, 0x4f, 0x75, 0x74, 0x22, 0x00,
	0x12, 0x57, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x63, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1b, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x49, 0x6e, 0x1a, 0x1c, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0d, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x1a, 0x11, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x75, 0x74, 0x22,
	0x00, 0x12, 0x30, 0x0a, 0x0b, 0x45, 0x64, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x0e, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e,
	0x1a, 0x0f, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x75,
	0x74, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_chat_proto_goTypes = []any{
	(*ChatEmpty)(nil),                   // 0: ChatEmpty
	(*CreatePrivateChatIn)(nil),         // 1: CreatePrivateChatIn
	(*CreatePrivateChatOut)(nil),        // 2: CreatePrivateChatOut
	(*Chat)(nil),                        // 3: Chat
	(*GetChatsOut)(nil),                 // 4: GetChatsOut
	(*Message)(nil),                     // 5: Message
	(*GetPrivateRecentMessagesIn)(nil),  // 6: GetPrivateRecentMessagesIn
	(*GetPrivateRecentMessagesOut)(nil), // 7: GetPrivateRecentMessagesOut
	(*DeleteMessageIn)(nil),             // 8: DeleteMessageIn
	(*DeleteMessageOut)(nil),            // 9: DeleteMessageOut
	(*EditMessageIn)(nil),               // 10: EditMessageIn
	(*EditMessageOut)(nil),              // 11: EditMessageOut
}
var file_chat_proto_depIdxs = []int32{
	3,  // 0: GetChatsOut.chats:type_name -> Chat
	5,  // 1: GetPrivateRecentMessagesOut.messages:type_name -> Message
	1,  // 2: ChatService.CreatePrivateChat:input_type -> CreatePrivateChatIn
	0,  // 3: ChatService.GetChats:input_type -> ChatEmpty
	6,  // 4: ChatService.GetPrivateRecentMessages:input_type -> GetPrivateRecentMessagesIn
	8,  // 5: ChatService.DeleteMessage:input_type -> DeleteMessageIn
	10, // 6: ChatService.EditMessage:input_type -> EditMessageIn
	2,  // 7: ChatService.CreatePrivateChat:output_type -> CreatePrivateChatOut
	4,  // 8: ChatService.GetChats:output_type -> GetChatsOut
	7,  // 9: ChatService.GetPrivateRecentMessages:output_type -> GetPrivateRecentMessagesOut
	9,  // 10: ChatService.DeleteMessage:output_type -> DeleteMessageOut
	11, // 11: ChatService.EditMessage:output_type -> EditMessageOut
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}
