// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: proto/Calamardo.proto

package proto

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

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Calamardo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Calamardo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_proto_Calamardo_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Calamardo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Calamardo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_proto_Calamardo_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type JoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *JoinRequest) Reset() {
	*x = JoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Calamardo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRequest) ProtoMessage() {}

func (x *JoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Calamardo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRequest.ProtoReflect.Descriptor instead.
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return file_proto_Calamardo_proto_rawDescGZIP(), []int{2}
}

func (x *JoinRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type JoinReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdJugador int32 `protobuf:"varint,1,opt,name=id_jugador,json=idJugador,proto3" json:"id_jugador,omitempty"`
	Alive     bool  `protobuf:"varint,2,opt,name=alive,proto3" json:"alive,omitempty"`
	Round     int32 `protobuf:"varint,3,opt,name=round,proto3" json:"round,omitempty"`
}

func (x *JoinReply) Reset() {
	*x = JoinReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Calamardo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinReply) ProtoMessage() {}

func (x *JoinReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Calamardo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinReply.ProtoReflect.Descriptor instead.
func (*JoinReply) Descriptor() ([]byte, []int) {
	return file_proto_Calamardo_proto_rawDescGZIP(), []int{3}
}

func (x *JoinReply) GetIdJugador() int32 {
	if x != nil {
		return x.IdJugador
	}
	return 0
}

func (x *JoinReply) GetAlive() bool {
	if x != nil {
		return x.Alive
	}
	return false
}

func (x *JoinReply) GetRound() int32 {
	if x != nil {
		return x.Round
	}
	return 0
}

type StartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StartRequest) Reset() {
	*x = StartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Calamardo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRequest) ProtoMessage() {}

func (x *StartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Calamardo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRequest.ProtoReflect.Descriptor instead.
func (*StartRequest) Descriptor() ([]byte, []int) {
	return file_proto_Calamardo_proto_rawDescGZIP(), []int{4}
}

func (x *StartRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type StartReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Started bool `protobuf:"varint,1,opt,name=started,proto3" json:"started,omitempty"`
}

func (x *StartReply) Reset() {
	*x = StartReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Calamardo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartReply) ProtoMessage() {}

func (x *StartReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Calamardo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartReply.ProtoReflect.Descriptor instead.
func (*StartReply) Descriptor() ([]byte, []int) {
	return file_proto_Calamardo_proto_rawDescGZIP(), []int{5}
}

func (x *StartReply) GetStarted() bool {
	if x != nil {
		return x.Started
	}
	return false
}

type JuegoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Jugada int32 `protobuf:"varint,2,opt,name=jugada,proto3" json:"jugada,omitempty"`
	Round  int32 `protobuf:"varint,3,opt,name=round,proto3" json:"round,omitempty"`
	Score  int32 `protobuf:"varint,4,opt,name=score,proto3" json:"score,omitempty"`
	Etapa  int32 `protobuf:"varint,5,opt,name=etapa,proto3" json:"etapa,omitempty"`
}

func (x *JuegoRequest) Reset() {
	*x = JuegoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Calamardo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JuegoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JuegoRequest) ProtoMessage() {}

func (x *JuegoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Calamardo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JuegoRequest.ProtoReflect.Descriptor instead.
func (*JuegoRequest) Descriptor() ([]byte, []int) {
	return file_proto_Calamardo_proto_rawDescGZIP(), []int{6}
}

func (x *JuegoRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *JuegoRequest) GetJugada() int32 {
	if x != nil {
		return x.Jugada
	}
	return 0
}

func (x *JuegoRequest) GetRound() int32 {
	if x != nil {
		return x.Round
	}
	return 0
}

func (x *JuegoRequest) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *JuegoRequest) GetEtapa() int32 {
	if x != nil {
		return x.Etapa
	}
	return 0
}

type JuegoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alive bool  `protobuf:"varint,1,opt,name=alive,proto3" json:"alive,omitempty"`
	Round int32 `protobuf:"varint,2,opt,name=round,proto3" json:"round,omitempty"`
	Score int32 `protobuf:"varint,3,opt,name=score,proto3" json:"score,omitempty"`
	Etapa int32 `protobuf:"varint,4,opt,name=etapa,proto3" json:"etapa,omitempty"`
}

func (x *JuegoReply) Reset() {
	*x = JuegoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Calamardo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JuegoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JuegoReply) ProtoMessage() {}

func (x *JuegoReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Calamardo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JuegoReply.ProtoReflect.Descriptor instead.
func (*JuegoReply) Descriptor() ([]byte, []int) {
	return file_proto_Calamardo_proto_rawDescGZIP(), []int{7}
}

func (x *JuegoReply) GetAlive() bool {
	if x != nil {
		return x.Alive
	}
	return false
}

func (x *JuegoReply) GetRound() int32 {
	if x != nil {
		return x.Round
	}
	return 0
}

func (x *JuegoReply) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *JuegoReply) GetEtapa() int32 {
	if x != nil {
		return x.Etapa
	}
	return 0
}

var File_proto_Calamardo_proto protoreflect.FileDescriptor

var file_proto_Calamardo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x61, 0x6c, 0x61, 0x6d, 0x61, 0x72, 0x64,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22,
	0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x27, 0x0a, 0x0b, 0x4a, 0x6f,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x56, 0x0a, 0x09, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x64, 0x5f, 0x6a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x69, 0x64, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x61, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x28, 0x0a, 0x0c, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x22, 0x78, 0x0a,
	0x0c, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x6a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6a,
	0x75, 0x67, 0x61, 0x64, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x74, 0x61, 0x70, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x65, 0x74, 0x61, 0x70, 0x61, 0x22, 0x64, 0x0a, 0x0a, 0x4a, 0x75, 0x65, 0x67, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x74, 0x61, 0x70, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x65, 0x74, 0x61, 0x70, 0x61, 0x32, 0xe6, 0x01,
	0x0a, 0x0d, 0x43, 0x61, 0x6c, 0x61, 0x6d, 0x61, 0x72, 0x64, 0x6f, 0x47, 0x61, 0x6d, 0x65, 0x12,
	0x34, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x08, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x6f,
	0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x09, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x34, 0x0a, 0x08, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x4d, 0x73, 0x67, 0x12, 0x13, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x6c, 0x61, 0x62, 0x2f, 0x67, 0x61,
	0x6d, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_Calamardo_proto_rawDescOnce sync.Once
	file_proto_Calamardo_proto_rawDescData = file_proto_Calamardo_proto_rawDesc
)

func file_proto_Calamardo_proto_rawDescGZIP() []byte {
	file_proto_Calamardo_proto_rawDescOnce.Do(func() {
		file_proto_Calamardo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_Calamardo_proto_rawDescData)
	})
	return file_proto_Calamardo_proto_rawDescData
}

var file_proto_Calamardo_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_Calamardo_proto_goTypes = []interface{}{
	(*HelloRequest)(nil), // 0: proto.HelloRequest
	(*HelloReply)(nil),   // 1: proto.HelloReply
	(*JoinRequest)(nil),  // 2: proto.JoinRequest
	(*JoinReply)(nil),    // 3: proto.JoinReply
	(*StartRequest)(nil), // 4: proto.StartRequest
	(*StartReply)(nil),   // 5: proto.StartReply
	(*JuegoRequest)(nil), // 6: proto.JuegoRequest
	(*JuegoReply)(nil),   // 7: proto.JuegoReply
}
var file_proto_Calamardo_proto_depIdxs = []int32{
	0, // 0: proto.CalamardoGame.SayHello:input_type -> proto.HelloRequest
	2, // 1: proto.CalamardoGame.JoinGame:input_type -> proto.JoinRequest
	4, // 2: proto.CalamardoGame.StartGame:input_type -> proto.StartRequest
	6, // 3: proto.CalamardoGame.JuegoMsg:input_type -> proto.JuegoRequest
	1, // 4: proto.CalamardoGame.SayHello:output_type -> proto.HelloReply
	3, // 5: proto.CalamardoGame.JoinGame:output_type -> proto.JoinReply
	5, // 6: proto.CalamardoGame.StartGame:output_type -> proto.StartReply
	7, // 7: proto.CalamardoGame.JuegoMsg:output_type -> proto.JuegoReply
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_Calamardo_proto_init() }
func file_proto_Calamardo_proto_init() {
	if File_proto_Calamardo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_Calamardo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_proto_Calamardo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
		file_proto_Calamardo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRequest); i {
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
		file_proto_Calamardo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinReply); i {
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
		file_proto_Calamardo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRequest); i {
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
		file_proto_Calamardo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartReply); i {
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
		file_proto_Calamardo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JuegoRequest); i {
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
		file_proto_Calamardo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JuegoReply); i {
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
			RawDescriptor: file_proto_Calamardo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_Calamardo_proto_goTypes,
		DependencyIndexes: file_proto_Calamardo_proto_depIdxs,
		MessageInfos:      file_proto_Calamardo_proto_msgTypes,
	}.Build()
	File_proto_Calamardo_proto = out.File
	file_proto_Calamardo_proto_rawDesc = nil
	file_proto_Calamardo_proto_goTypes = nil
	file_proto_Calamardo_proto_depIdxs = nil
}
