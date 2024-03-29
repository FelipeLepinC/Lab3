// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: lab3.proto

package lab3

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planeta  string `protobuf:"bytes,1,opt,name=Planeta,proto3" json:"Planeta,omitempty"`
	Ciudad   string `protobuf:"bytes,2,opt,name=Ciudad,proto3" json:"Ciudad,omitempty"`
	Soldados int32  `protobuf:"varint,3,opt,name=Soldados,proto3" json:"Soldados,omitempty"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lab3_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_lab3_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Info.ProtoReflect.Descriptor instead.
func (*Info) Descriptor() ([]byte, []int) {
	return file_lab3_proto_rawDescGZIP(), []int{0}
}

func (x *Info) GetPlaneta() string {
	if x != nil {
		return x.Planeta
	}
	return ""
}

func (x *Info) GetCiudad() string {
	if x != nil {
		return x.Ciudad
	}
	return ""
}

func (x *Info) GetSoldados() int32 {
	if x != nil {
		return x.Soldados
	}
	return 0
}

type Operacion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accion   string `protobuf:"bytes,1,opt,name=Accion,proto3" json:"Accion,omitempty"`
	Planeta  string `protobuf:"bytes,2,opt,name=Planeta,proto3" json:"Planeta,omitempty"`
	Ciudad   string `protobuf:"bytes,3,opt,name=Ciudad,proto3" json:"Ciudad,omitempty"`
	Intvalue int32  `protobuf:"varint,4,opt,name=Intvalue,proto3" json:"Intvalue,omitempty"`
	Svalue   string `protobuf:"bytes,5,opt,name=Svalue,proto3" json:"Svalue,omitempty"`
	Servidor int32  `protobuf:"varint,6,opt,name=Servidor,proto3" json:"Servidor,omitempty"`
}

func (x *Operacion) Reset() {
	*x = Operacion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lab3_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operacion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operacion) ProtoMessage() {}

func (x *Operacion) ProtoReflect() protoreflect.Message {
	mi := &file_lab3_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operacion.ProtoReflect.Descriptor instead.
func (*Operacion) Descriptor() ([]byte, []int) {
	return file_lab3_proto_rawDescGZIP(), []int{1}
}

func (x *Operacion) GetAccion() string {
	if x != nil {
		return x.Accion
	}
	return ""
}

func (x *Operacion) GetPlaneta() string {
	if x != nil {
		return x.Planeta
	}
	return ""
}

func (x *Operacion) GetCiudad() string {
	if x != nil {
		return x.Ciudad
	}
	return ""
}

func (x *Operacion) GetIntvalue() int32 {
	if x != nil {
		return x.Intvalue
	}
	return 0
}

func (x *Operacion) GetSvalue() string {
	if x != nil {
		return x.Svalue
	}
	return ""
}

func (x *Operacion) GetServidor() int32 {
	if x != nil {
		return x.Servidor
	}
	return 0
}

type L struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planeta  string `protobuf:"bytes,1,opt,name=Planeta,proto3" json:"Planeta,omitempty"`
	Ciudad   string `protobuf:"bytes,2,opt,name=Ciudad,proto3" json:"Ciudad,omitempty"`
	Servidor int32  `protobuf:"varint,3,opt,name=Servidor,proto3" json:"Servidor,omitempty"`
}

func (x *L) Reset() {
	*x = L{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lab3_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *L) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*L) ProtoMessage() {}

func (x *L) ProtoReflect() protoreflect.Message {
	mi := &file_lab3_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use L.ProtoReflect.Descriptor instead.
func (*L) Descriptor() ([]byte, []int) {
	return file_lab3_proto_rawDescGZIP(), []int{2}
}

func (x *L) GetPlaneta() string {
	if x != nil {
		return x.Planeta
	}
	return ""
}

func (x *L) GetCiudad() string {
	if x != nil {
		return x.Ciudad
	}
	return ""
}

func (x *L) GetServidor() int32 {
	if x != nil {
		return x.Servidor
	}
	return 0
}

type Lresponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planeta  string `protobuf:"bytes,1,opt,name=Planeta,proto3" json:"Planeta,omitempty"`
	Valor    int32  `protobuf:"varint,2,opt,name=Valor,proto3" json:"Valor,omitempty"`
	X        int32  `protobuf:"varint,3,opt,name=X,proto3" json:"X,omitempty"`
	Y        int32  `protobuf:"varint,4,opt,name=Y,proto3" json:"Y,omitempty"`
	Z        int32  `protobuf:"varint,5,opt,name=Z,proto3" json:"Z,omitempty"`
	Servidor int32  `protobuf:"varint,6,opt,name=Servidor,proto3" json:"Servidor,omitempty"`
	Merge    bool   `protobuf:"varint,7,opt,name=Merge,proto3" json:"Merge,omitempty"`
}

func (x *Lresponse) Reset() {
	*x = Lresponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lab3_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lresponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lresponse) ProtoMessage() {}

func (x *Lresponse) ProtoReflect() protoreflect.Message {
	mi := &file_lab3_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lresponse.ProtoReflect.Descriptor instead.
func (*Lresponse) Descriptor() ([]byte, []int) {
	return file_lab3_proto_rawDescGZIP(), []int{3}
}

func (x *Lresponse) GetPlaneta() string {
	if x != nil {
		return x.Planeta
	}
	return ""
}

func (x *Lresponse) GetValor() int32 {
	if x != nil {
		return x.Valor
	}
	return 0
}

func (x *Lresponse) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Lresponse) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Lresponse) GetZ() int32 {
	if x != nil {
		return x.Z
	}
	return 0
}

func (x *Lresponse) GetServidor() int32 {
	if x != nil {
		return x.Servidor
	}
	return 0
}

func (x *Lresponse) GetMerge() bool {
	if x != nil {
		return x.Merge
	}
	return false
}

type Reloj struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planeta string `protobuf:"bytes,1,opt,name=Planeta,proto3" json:"Planeta,omitempty"`
	X       int32  `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
	Y       int32  `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`
	Z       int32  `protobuf:"varint,4,opt,name=z,proto3" json:"z,omitempty"`
	Merge   bool   `protobuf:"varint,5,opt,name=merge,proto3" json:"merge,omitempty"`
}

func (x *Reloj) Reset() {
	*x = Reloj{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lab3_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reloj) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reloj) ProtoMessage() {}

func (x *Reloj) ProtoReflect() protoreflect.Message {
	mi := &file_lab3_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reloj.ProtoReflect.Descriptor instead.
func (*Reloj) Descriptor() ([]byte, []int) {
	return file_lab3_proto_rawDescGZIP(), []int{4}
}

func (x *Reloj) GetPlaneta() string {
	if x != nil {
		return x.Planeta
	}
	return ""
}

func (x *Reloj) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Reloj) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Reloj) GetZ() int32 {
	if x != nil {
		return x.Z
	}
	return 0
}

func (x *Reloj) GetMerge() bool {
	if x != nil {
		return x.Merge
	}
	return false
}

type Cont struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=X,proto3" json:"X,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=Y,proto3" json:"Y,omitempty"`
	Z int32 `protobuf:"varint,3,opt,name=Z,proto3" json:"Z,omitempty"`
}

func (x *Cont) Reset() {
	*x = Cont{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lab3_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cont) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cont) ProtoMessage() {}

func (x *Cont) ProtoReflect() protoreflect.Message {
	mi := &file_lab3_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cont.ProtoReflect.Descriptor instead.
func (*Cont) Descriptor() ([]byte, []int) {
	return file_lab3_proto_rawDescGZIP(), []int{5}
}

func (x *Cont) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Cont) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Cont) GetZ() int32 {
	if x != nil {
		return x.Z
	}
	return 0
}

type Merge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	M map[string]string `protobuf:"bytes,1,rep,name=m,proto3" json:"m,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Merge) Reset() {
	*x = Merge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lab3_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Merge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Merge) ProtoMessage() {}

func (x *Merge) ProtoReflect() protoreflect.Message {
	mi := &file_lab3_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Merge.ProtoReflect.Descriptor instead.
func (*Merge) Descriptor() ([]byte, []int) {
	return file_lab3_proto_rawDescGZIP(), []int{6}
}

func (x *Merge) GetM() map[string]string {
	if x != nil {
		return x.M
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nserver int32 `protobuf:"varint,1,opt,name=Nserver,proto3" json:"Nserver,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lab3_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_lab3_proto_msgTypes[7]
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
	return file_lab3_proto_rawDescGZIP(), []int{7}
}

func (x *Message) GetNserver() int32 {
	if x != nil {
		return x.Nserver
	}
	return 0
}

var File_lab3_proto protoreflect.FileDescriptor

var file_lab3_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x61, 0x62, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6c, 0x61,
	0x62, 0x33, 0x22, 0x54, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x6c,
	0x61, 0x6e, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x6c, 0x61,
	0x6e, 0x65, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x69, 0x75, 0x64, 0x61, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x69, 0x75, 0x64, 0x61, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x53, 0x6f, 0x6c, 0x64, 0x61, 0x64, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x53, 0x6f, 0x6c, 0x64, 0x61, 0x64, 0x6f, 0x73, 0x22, 0xa5, 0x01, 0x0a, 0x09, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x63, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x63, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x69, 0x75, 0x64,
	0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x69, 0x75, 0x64, 0x61, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x49, 0x6e, 0x74, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x53, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72,
	0x22, 0x51, 0x0a, 0x01, 0x4c, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x12,
	0x16, 0x0a, 0x06, 0x43, 0x69, 0x75, 0x64, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x43, 0x69, 0x75, 0x64, 0x61, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x64, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x64, 0x6f, 0x72, 0x22, 0x97, 0x01, 0x0a, 0x09, 0x4c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x56,
	0x61, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x6f,
	0x72, 0x12, 0x0c, 0x0a, 0x01, 0x58, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x58, 0x12,
	0x0c, 0x0a, 0x01, 0x59, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x59, 0x12, 0x0c, 0x0a,
	0x01, 0x5a, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x5a, 0x12, 0x1a, 0x0a, 0x08, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x65, 0x72, 0x67, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x22, 0x61, 0x0a,
	0x05, 0x52, 0x65, 0x6c, 0x6f, 0x6a, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61,
	0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c,
	0x0a, 0x01, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01,
	0x7a, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x7a, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x65,
	0x72, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6d, 0x65, 0x72, 0x67, 0x65,
	0x22, 0x30, 0x0a, 0x04, 0x43, 0x6f, 0x6e, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x58, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x01, 0x58, 0x12, 0x0c, 0x0a, 0x01, 0x59, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x01, 0x59, 0x12, 0x0c, 0x0a, 0x01, 0x5a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x01, 0x5a, 0x22, 0x5f, 0x0a, 0x05, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x01, 0x6d,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x61, 0x62, 0x33, 0x2e, 0x4d, 0x65,
	0x72, 0x67, 0x65, 0x2e, 0x4d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x01, 0x6d, 0x1a, 0x34, 0x0a,
	0x06, 0x4d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x23, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x4e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x4e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x32, 0xfb, 0x02, 0x0a, 0x08, 0x53, 0x74, 0x61,
	0x72, 0x77, 0x61, 0x72, 0x73, 0x12, 0x24, 0x0a, 0x0a, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x69,
	0x6e, 0x66, 0x6f, 0x12, 0x0a, 0x2e, 0x6c, 0x61, 0x62, 0x33, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x1a,
	0x0a, 0x2e, 0x6c, 0x61, 0x62, 0x33, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x0c, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x61, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0f, 0x2e, 0x6c, 0x61,
	0x62, 0x33, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x1a, 0x0d, 0x2e, 0x6c,
	0x61, 0x62, 0x33, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x46,
	0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x12, 0x0f, 0x2e, 0x6c, 0x61, 0x62, 0x33, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x1a, 0x0b, 0x2e, 0x6c, 0x61, 0x62, 0x33, 0x2e, 0x52,
	0x65, 0x6c, 0x6f, 0x6a, 0x12, 0x2d, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x72, 0x69, 0x6f, 0x12, 0x0d, 0x2e, 0x6c, 0x61, 0x62, 0x33, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x0d, 0x2e, 0x6c, 0x61, 0x62, 0x33, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x4c, 0x65, 0x69, 0x61, 0x12, 0x07, 0x2e, 0x6c, 0x61,
	0x62, 0x33, 0x2e, 0x4c, 0x1a, 0x0f, 0x2e, 0x6c, 0x61, 0x62, 0x33, 0x2e, 0x4c, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x09, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6c, 0x65,
	0x69, 0x61, 0x12, 0x0f, 0x2e, 0x6c, 0x61, 0x62, 0x33, 0x2e, 0x4c, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x1a, 0x07, 0x2e, 0x6c, 0x61, 0x62, 0x33, 0x2e, 0x4c, 0x12, 0x27, 0x0a, 0x0b,
	0x4c, 0x65, 0x69, 0x61, 0x66, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x12, 0x07, 0x2e, 0x6c, 0x61,
	0x62, 0x33, 0x2e, 0x4c, 0x1a, 0x0f, 0x2e, 0x6c, 0x61, 0x62, 0x33, 0x2e, 0x4c, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x08, 0x50, 0x65, 0x64, 0x69, 0x72, 0x64, 0x69,
	0x63, 0x12, 0x0d, 0x2e, 0x6c, 0x61, 0x62, 0x33, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x0b, 0x2e, 0x6c, 0x61, 0x62, 0x33, 0x2e, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x12, 0x27, 0x0a,
	0x09, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x64, 0x69, 0x63, 0x12, 0x0b, 0x2e, 0x6c, 0x61, 0x62,
	0x33, 0x2e, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x1a, 0x0d, 0x2e, 0x6c, 0x61, 0x62, 0x33, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x6c, 0x61, 0x62, 0x33, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lab3_proto_rawDescOnce sync.Once
	file_lab3_proto_rawDescData = file_lab3_proto_rawDesc
)

func file_lab3_proto_rawDescGZIP() []byte {
	file_lab3_proto_rawDescOnce.Do(func() {
		file_lab3_proto_rawDescData = protoimpl.X.CompressGZIP(file_lab3_proto_rawDescData)
	})
	return file_lab3_proto_rawDescData
}

var file_lab3_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_lab3_proto_goTypes = []interface{}{
	(*Info)(nil),      // 0: lab3.Info
	(*Operacion)(nil), // 1: lab3.Operacion
	(*L)(nil),         // 2: lab3.L
	(*Lresponse)(nil), // 3: lab3.Lresponse
	(*Reloj)(nil),     // 4: lab3.Reloj
	(*Cont)(nil),      // 5: lab3.Cont
	(*Merge)(nil),     // 6: lab3.Merge
	(*Message)(nil),   // 7: lab3.Message
	nil,               // 8: lab3.Merge.MEntry
}
var file_lab3_proto_depIdxs = []int32{
	8,  // 0: lab3.Merge.m:type_name -> lab3.Merge.MEntry
	0,  // 1: lab3.Starwars.Enviarinfo:input_type -> lab3.Info
	1,  // 2: lab3.Starwars.Alertabroken:input_type -> lab3.Operacion
	1,  // 3: lab3.Starwars.Fulcrum:input_type -> lab3.Operacion
	7,  // 4: lab3.Starwars.Intermediario:input_type -> lab3.Message
	2,  // 5: lab3.Starwars.Leia:input_type -> lab3.L
	3,  // 6: lab3.Starwars.Interleia:input_type -> lab3.Lresponse
	2,  // 7: lab3.Starwars.Leiafulcrum:input_type -> lab3.L
	7,  // 8: lab3.Starwars.Pedirdic:input_type -> lab3.Message
	6,  // 9: lab3.Starwars.Enviardic:input_type -> lab3.Merge
	0,  // 10: lab3.Starwars.Enviarinfo:output_type -> lab3.Info
	7,  // 11: lab3.Starwars.Alertabroken:output_type -> lab3.Message
	4,  // 12: lab3.Starwars.Fulcrum:output_type -> lab3.Reloj
	7,  // 13: lab3.Starwars.Intermediario:output_type -> lab3.Message
	3,  // 14: lab3.Starwars.Leia:output_type -> lab3.Lresponse
	2,  // 15: lab3.Starwars.Interleia:output_type -> lab3.L
	3,  // 16: lab3.Starwars.Leiafulcrum:output_type -> lab3.Lresponse
	6,  // 17: lab3.Starwars.Pedirdic:output_type -> lab3.Merge
	7,  // 18: lab3.Starwars.Enviardic:output_type -> lab3.Message
	10, // [10:19] is the sub-list for method output_type
	1,  // [1:10] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_lab3_proto_init() }
func file_lab3_proto_init() {
	if File_lab3_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lab3_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Info); i {
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
		file_lab3_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operacion); i {
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
		file_lab3_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*L); i {
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
		file_lab3_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lresponse); i {
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
		file_lab3_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reloj); i {
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
		file_lab3_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cont); i {
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
		file_lab3_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Merge); i {
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
		file_lab3_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_lab3_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lab3_proto_goTypes,
		DependencyIndexes: file_lab3_proto_depIdxs,
		MessageInfos:      file_lab3_proto_msgTypes,
	}.Build()
	File_lab3_proto = out.File
	file_lab3_proto_rawDesc = nil
	file_lab3_proto_goTypes = nil
	file_lab3_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StarwarsClient is the client API for Starwars service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StarwarsClient interface {
	Enviarinfo(ctx context.Context, in *Info, opts ...grpc.CallOption) (*Info, error)
	Alertabroken(ctx context.Context, in *Operacion, opts ...grpc.CallOption) (*Message, error)
	Fulcrum(ctx context.Context, in *Operacion, opts ...grpc.CallOption) (*Reloj, error)
	Intermediario(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Leia(ctx context.Context, in *L, opts ...grpc.CallOption) (*Lresponse, error)
	Interleia(ctx context.Context, in *Lresponse, opts ...grpc.CallOption) (*L, error)
	Leiafulcrum(ctx context.Context, in *L, opts ...grpc.CallOption) (*Lresponse, error)
	Pedirdic(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Merge, error)
	Enviardic(ctx context.Context, in *Merge, opts ...grpc.CallOption) (*Message, error)
}

type starwarsClient struct {
	cc grpc.ClientConnInterface
}

func NewStarwarsClient(cc grpc.ClientConnInterface) StarwarsClient {
	return &starwarsClient{cc}
}

func (c *starwarsClient) Enviarinfo(ctx context.Context, in *Info, opts ...grpc.CallOption) (*Info, error) {
	out := new(Info)
	err := c.cc.Invoke(ctx, "/lab3.Starwars/Enviarinfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *starwarsClient) Alertabroken(ctx context.Context, in *Operacion, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/lab3.Starwars/Alertabroken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *starwarsClient) Fulcrum(ctx context.Context, in *Operacion, opts ...grpc.CallOption) (*Reloj, error) {
	out := new(Reloj)
	err := c.cc.Invoke(ctx, "/lab3.Starwars/Fulcrum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *starwarsClient) Intermediario(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/lab3.Starwars/Intermediario", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *starwarsClient) Leia(ctx context.Context, in *L, opts ...grpc.CallOption) (*Lresponse, error) {
	out := new(Lresponse)
	err := c.cc.Invoke(ctx, "/lab3.Starwars/Leia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *starwarsClient) Interleia(ctx context.Context, in *Lresponse, opts ...grpc.CallOption) (*L, error) {
	out := new(L)
	err := c.cc.Invoke(ctx, "/lab3.Starwars/Interleia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *starwarsClient) Leiafulcrum(ctx context.Context, in *L, opts ...grpc.CallOption) (*Lresponse, error) {
	out := new(Lresponse)
	err := c.cc.Invoke(ctx, "/lab3.Starwars/Leiafulcrum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *starwarsClient) Pedirdic(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Merge, error) {
	out := new(Merge)
	err := c.cc.Invoke(ctx, "/lab3.Starwars/Pedirdic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *starwarsClient) Enviardic(ctx context.Context, in *Merge, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/lab3.Starwars/Enviardic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StarwarsServer is the server API for Starwars service.
type StarwarsServer interface {
	Enviarinfo(context.Context, *Info) (*Info, error)
	Alertabroken(context.Context, *Operacion) (*Message, error)
	Fulcrum(context.Context, *Operacion) (*Reloj, error)
	Intermediario(context.Context, *Message) (*Message, error)
	Leia(context.Context, *L) (*Lresponse, error)
	Interleia(context.Context, *Lresponse) (*L, error)
	Leiafulcrum(context.Context, *L) (*Lresponse, error)
	Pedirdic(context.Context, *Message) (*Merge, error)
	Enviardic(context.Context, *Merge) (*Message, error)
}

// UnimplementedStarwarsServer can be embedded to have forward compatible implementations.
type UnimplementedStarwarsServer struct {
}

func (*UnimplementedStarwarsServer) Enviarinfo(context.Context, *Info) (*Info, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enviarinfo not implemented")
}
func (*UnimplementedStarwarsServer) Alertabroken(context.Context, *Operacion) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Alertabroken not implemented")
}
func (*UnimplementedStarwarsServer) Fulcrum(context.Context, *Operacion) (*Reloj, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fulcrum not implemented")
}
func (*UnimplementedStarwarsServer) Intermediario(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Intermediario not implemented")
}
func (*UnimplementedStarwarsServer) Leia(context.Context, *L) (*Lresponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Leia not implemented")
}
func (*UnimplementedStarwarsServer) Interleia(context.Context, *Lresponse) (*L, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Interleia not implemented")
}
func (*UnimplementedStarwarsServer) Leiafulcrum(context.Context, *L) (*Lresponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Leiafulcrum not implemented")
}
func (*UnimplementedStarwarsServer) Pedirdic(context.Context, *Message) (*Merge, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pedirdic not implemented")
}
func (*UnimplementedStarwarsServer) Enviardic(context.Context, *Merge) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enviardic not implemented")
}

func RegisterStarwarsServer(s *grpc.Server, srv StarwarsServer) {
	s.RegisterService(&_Starwars_serviceDesc, srv)
}

func _Starwars_Enviarinfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Info)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarwarsServer).Enviarinfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lab3.Starwars/Enviarinfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarwarsServer).Enviarinfo(ctx, req.(*Info))
	}
	return interceptor(ctx, in, info, handler)
}

func _Starwars_Alertabroken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Operacion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarwarsServer).Alertabroken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lab3.Starwars/Alertabroken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarwarsServer).Alertabroken(ctx, req.(*Operacion))
	}
	return interceptor(ctx, in, info, handler)
}

func _Starwars_Fulcrum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Operacion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarwarsServer).Fulcrum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lab3.Starwars/Fulcrum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarwarsServer).Fulcrum(ctx, req.(*Operacion))
	}
	return interceptor(ctx, in, info, handler)
}

func _Starwars_Intermediario_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarwarsServer).Intermediario(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lab3.Starwars/Intermediario",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarwarsServer).Intermediario(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Starwars_Leia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(L)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarwarsServer).Leia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lab3.Starwars/Leia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarwarsServer).Leia(ctx, req.(*L))
	}
	return interceptor(ctx, in, info, handler)
}

func _Starwars_Interleia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Lresponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarwarsServer).Interleia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lab3.Starwars/Interleia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarwarsServer).Interleia(ctx, req.(*Lresponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _Starwars_Leiafulcrum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(L)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarwarsServer).Leiafulcrum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lab3.Starwars/Leiafulcrum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarwarsServer).Leiafulcrum(ctx, req.(*L))
	}
	return interceptor(ctx, in, info, handler)
}

func _Starwars_Pedirdic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarwarsServer).Pedirdic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lab3.Starwars/Pedirdic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarwarsServer).Pedirdic(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Starwars_Enviardic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Merge)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarwarsServer).Enviardic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lab3.Starwars/Enviardic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarwarsServer).Enviardic(ctx, req.(*Merge))
	}
	return interceptor(ctx, in, info, handler)
}

var _Starwars_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lab3.Starwars",
	HandlerType: (*StarwarsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Enviarinfo",
			Handler:    _Starwars_Enviarinfo_Handler,
		},
		{
			MethodName: "Alertabroken",
			Handler:    _Starwars_Alertabroken_Handler,
		},
		{
			MethodName: "Fulcrum",
			Handler:    _Starwars_Fulcrum_Handler,
		},
		{
			MethodName: "Intermediario",
			Handler:    _Starwars_Intermediario_Handler,
		},
		{
			MethodName: "Leia",
			Handler:    _Starwars_Leia_Handler,
		},
		{
			MethodName: "Interleia",
			Handler:    _Starwars_Interleia_Handler,
		},
		{
			MethodName: "Leiafulcrum",
			Handler:    _Starwars_Leiafulcrum_Handler,
		},
		{
			MethodName: "Pedirdic",
			Handler:    _Starwars_Pedirdic_Handler,
		},
		{
			MethodName: "Enviardic",
			Handler:    _Starwars_Enviardic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lab3.proto",
}
