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

	Accion  string `protobuf:"bytes,1,opt,name=Accion,proto3" json:"Accion,omitempty"`
	Planeta string `protobuf:"bytes,2,opt,name=Planeta,proto3" json:"Planeta,omitempty"`
	Ciudad  string `protobuf:"bytes,3,opt,name=Ciudad,proto3" json:"Ciudad,omitempty"`
	Valor   int32  `protobuf:"varint,4,opt,name=Valor,proto3" json:"Valor,omitempty"`
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

func (x *Operacion) GetValor() int32 {
	if x != nil {
		return x.Valor
	}
	return 0
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
		mi := &file_lab3_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_lab3_proto_rawDescGZIP(), []int{2}
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
	0x53, 0x6f, 0x6c, 0x64, 0x61, 0x64, 0x6f, 0x73, 0x22, 0x6b, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x63, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x63, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x69, 0x75, 0x64, 0x61,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x69, 0x75, 0x64, 0x61, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x56, 0x61, 0x6c, 0x6f, 0x72, 0x22, 0x23, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x4e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x4e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x32, 0x60, 0x0a, 0x08, 0x53, 0x74,
	0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x12, 0x24, 0x0a, 0x0a, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72,
	0x69, 0x6e, 0x66, 0x6f, 0x12, 0x0a, 0x2e, 0x6c, 0x61, 0x62, 0x33, 0x2e, 0x49, 0x6e, 0x66, 0x6f,
	0x1a, 0x0a, 0x2e, 0x6c, 0x61, 0x62, 0x33, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x0c,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x61, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0f, 0x2e, 0x6c,
	0x61, 0x62, 0x33, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x1a, 0x0d, 0x2e,
	0x6c, 0x61, 0x62, 0x33, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x07, 0x5a, 0x05,
	0x2f, 0x6c, 0x61, 0x62, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_lab3_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_lab3_proto_goTypes = []interface{}{
	(*Info)(nil),      // 0: lab3.Info
	(*Operacion)(nil), // 1: lab3.Operacion
	(*Message)(nil),   // 2: lab3.Message
}
var file_lab3_proto_depIdxs = []int32{
	0, // 0: lab3.Starwars.Enviarinfo:input_type -> lab3.Info
	1, // 1: lab3.Starwars.Alertabroken:input_type -> lab3.Operacion
	0, // 2: lab3.Starwars.Enviarinfo:output_type -> lab3.Info
	2, // 3: lab3.Starwars.Alertabroken:output_type -> lab3.Message
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
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
			NumMessages:   3,
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

// StarwarsServer is the server API for Starwars service.
type StarwarsServer interface {
	Enviarinfo(context.Context, *Info) (*Info, error)
	Alertabroken(context.Context, *Operacion) (*Message, error)
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lab3.proto",
}
