// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: telegram_backend/telegram_backend.proto

package telegram_backend

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type Bot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	Username  string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *Bot) Reset() {
	*x = Bot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telegram_backend_telegram_backend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bot) ProtoMessage() {}

func (x *Bot) ProtoReflect() protoreflect.Message {
	mi := &file_telegram_backend_telegram_backend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bot.ProtoReflect.Descriptor instead.
func (*Bot) Descriptor() ([]byte, []int) {
	return file_telegram_backend_telegram_backend_proto_rawDescGZIP(), []int{0}
}

func (x *Bot) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Bot) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Bot) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

// AuthByInitData --------------------------------------------------------------------
type AuthByInitDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InitData string `protobuf:"bytes,1,opt,name=init_data,json=initData,proto3" json:"init_data,omitempty"`
}

func (x *AuthByInitDataRequest) Reset() {
	*x = AuthByInitDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telegram_backend_telegram_backend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthByInitDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthByInitDataRequest) ProtoMessage() {}

func (x *AuthByInitDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_telegram_backend_telegram_backend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthByInitDataRequest.ProtoReflect.Descriptor instead.
func (*AuthByInitDataRequest) Descriptor() ([]byte, []int) {
	return file_telegram_backend_telegram_backend_proto_rawDescGZIP(), []int{1}
}

func (x *AuthByInitDataRequest) GetInitData() string {
	if x != nil {
		return x.InitData
	}
	return ""
}

type AuthByInitDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwt string `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
}

func (x *AuthByInitDataResponse) Reset() {
	*x = AuthByInitDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telegram_backend_telegram_backend_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthByInitDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthByInitDataResponse) ProtoMessage() {}

func (x *AuthByInitDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_telegram_backend_telegram_backend_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthByInitDataResponse.ProtoReflect.Descriptor instead.
func (*AuthByInitDataResponse) Descriptor() ([]byte, []int) {
	return file_telegram_backend_telegram_backend_proto_rawDescGZIP(), []int{2}
}

func (x *AuthByInitDataResponse) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

// AddBot ----------------------------------------------------------------------------
type AddBotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AddBotRequest) Reset() {
	*x = AddBotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telegram_backend_telegram_backend_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBotRequest) ProtoMessage() {}

func (x *AddBotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_telegram_backend_telegram_backend_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBotRequest.ProtoReflect.Descriptor instead.
func (*AddBotRequest) Descriptor() ([]byte, []int) {
	return file_telegram_backend_telegram_backend_proto_rawDescGZIP(), []int{3}
}

func (x *AddBotRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AddBotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *Bot `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AddBotResponse) Reset() {
	*x = AddBotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telegram_backend_telegram_backend_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBotResponse) ProtoMessage() {}

func (x *AddBotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_telegram_backend_telegram_backend_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBotResponse.ProtoReflect.Descriptor instead.
func (*AddBotResponse) Descriptor() ([]byte, []int) {
	return file_telegram_backend_telegram_backend_proto_rawDescGZIP(), []int{4}
}

func (x *AddBotResponse) GetResult() *Bot {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_telegram_backend_telegram_backend_proto protoreflect.FileDescriptor

var file_telegram_backend_telegram_backend_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x74, 0x65, 0x6c, 0x65, 0x67,
	0x72, 0x61, 0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x1a, 0x1b, 0x62, 0x75, 0x66,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x03, 0x42, 0x6f, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x34, 0x0a, 0x15, 0x41, 0x75,
	0x74, 0x68, 0x42, 0x79, 0x49, 0x6e, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x2a, 0x0a, 0x16, 0x41, 0x75, 0x74, 0x68, 0x42, 0x79, 0x49, 0x6e, 0x69, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x77,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x22, 0x2f, 0x0a, 0x0d,
	0x41, 0x64, 0x64, 0x42, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48,
	0x05, 0x72, 0x03, 0x98, 0x01, 0x2e, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3f, 0x0a,
	0x0e, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2d, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2e, 0x42, 0x6f, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xca,
	0x01, 0x0a, 0x16, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x42, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x0e, 0x41, 0x75, 0x74,
	0x68, 0x42, 0x79, 0x49, 0x6e, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x27, 0x2e, 0x74, 0x65,
	0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x42, 0x79, 0x49, 0x6e, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x5f,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x42, 0x79, 0x49, 0x6e,
	0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b,
	0x0a, 0x06, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x74, 0x12, 0x1f, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x67,
	0x72, 0x61, 0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x41, 0x64, 0x64, 0x42,
	0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x67, 0x72, 0x61, 0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x41, 0x64, 0x64,
	0x42, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xbb, 0x01, 0x0a, 0x14,
	0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x42, 0x14, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x42, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x75, 0x69, 0x6e, 0x74, 0x61, 0x2d,
	0x6e, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0xa2,
	0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x0f, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d,
	0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0xca, 0x02, 0x0f, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72,
	0x61, 0x6d, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0xe2, 0x02, 0x1b, 0x54, 0x65, 0x6c, 0x65,
	0x67, 0x72, 0x61, 0x6d, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72,
	0x61, 0x6d, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_telegram_backend_telegram_backend_proto_rawDescOnce sync.Once
	file_telegram_backend_telegram_backend_proto_rawDescData = file_telegram_backend_telegram_backend_proto_rawDesc
)

func file_telegram_backend_telegram_backend_proto_rawDescGZIP() []byte {
	file_telegram_backend_telegram_backend_proto_rawDescOnce.Do(func() {
		file_telegram_backend_telegram_backend_proto_rawDescData = protoimpl.X.CompressGZIP(file_telegram_backend_telegram_backend_proto_rawDescData)
	})
	return file_telegram_backend_telegram_backend_proto_rawDescData
}

var file_telegram_backend_telegram_backend_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_telegram_backend_telegram_backend_proto_goTypes = []any{
	(*Bot)(nil),                    // 0: telegram_backend.Bot
	(*AuthByInitDataRequest)(nil),  // 1: telegram_backend.AuthByInitDataRequest
	(*AuthByInitDataResponse)(nil), // 2: telegram_backend.AuthByInitDataResponse
	(*AddBotRequest)(nil),          // 3: telegram_backend.AddBotRequest
	(*AddBotResponse)(nil),         // 4: telegram_backend.AddBotResponse
}
var file_telegram_backend_telegram_backend_proto_depIdxs = []int32{
	0, // 0: telegram_backend.AddBotResponse.result:type_name -> telegram_backend.Bot
	1, // 1: telegram_backend.TelegramBackendService.AuthByInitData:input_type -> telegram_backend.AuthByInitDataRequest
	3, // 2: telegram_backend.TelegramBackendService.AddBot:input_type -> telegram_backend.AddBotRequest
	2, // 3: telegram_backend.TelegramBackendService.AuthByInitData:output_type -> telegram_backend.AuthByInitDataResponse
	4, // 4: telegram_backend.TelegramBackendService.AddBot:output_type -> telegram_backend.AddBotResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_telegram_backend_telegram_backend_proto_init() }
func file_telegram_backend_telegram_backend_proto_init() {
	if File_telegram_backend_telegram_backend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_telegram_backend_telegram_backend_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Bot); i {
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
		file_telegram_backend_telegram_backend_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AuthByInitDataRequest); i {
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
		file_telegram_backend_telegram_backend_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AuthByInitDataResponse); i {
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
		file_telegram_backend_telegram_backend_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AddBotRequest); i {
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
		file_telegram_backend_telegram_backend_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AddBotResponse); i {
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
			RawDescriptor: file_telegram_backend_telegram_backend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_telegram_backend_telegram_backend_proto_goTypes,
		DependencyIndexes: file_telegram_backend_telegram_backend_proto_depIdxs,
		MessageInfos:      file_telegram_backend_telegram_backend_proto_msgTypes,
	}.Build()
	File_telegram_backend_telegram_backend_proto = out.File
	file_telegram_backend_telegram_backend_proto_rawDesc = nil
	file_telegram_backend_telegram_backend_proto_goTypes = nil
	file_telegram_backend_telegram_backend_proto_depIdxs = nil
}
