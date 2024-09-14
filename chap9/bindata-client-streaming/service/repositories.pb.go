// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: repositories.proto

package service

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

type RepoCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Body:
	//	*RepoCreateRequest_Context
	//	*RepoCreateRequest_Data
	Body isRepoCreateRequest_Body `protobuf_oneof:"body"`
}

func (x *RepoCreateRequest) Reset() {
	*x = RepoCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repositories_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepoCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepoCreateRequest) ProtoMessage() {}

func (x *RepoCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_repositories_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepoCreateRequest.ProtoReflect.Descriptor instead.
func (*RepoCreateRequest) Descriptor() ([]byte, []int) {
	return file_repositories_proto_rawDescGZIP(), []int{0}
}

func (m *RepoCreateRequest) GetBody() isRepoCreateRequest_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *RepoCreateRequest) GetContext() *RepoContext {
	if x, ok := x.GetBody().(*RepoCreateRequest_Context); ok {
		return x.Context
	}
	return nil
}

func (x *RepoCreateRequest) GetData() []byte {
	if x, ok := x.GetBody().(*RepoCreateRequest_Data); ok {
		return x.Data
	}
	return nil
}

type isRepoCreateRequest_Body interface {
	isRepoCreateRequest_Body()
}

type RepoCreateRequest_Context struct {
	Context *RepoContext `protobuf:"bytes,1,opt,name=context,proto3,oneof"`
}

type RepoCreateRequest_Data struct {
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

func (*RepoCreateRequest_Context) isRepoCreateRequest_Body() {}

func (*RepoCreateRequest_Data) isRepoCreateRequest_Body() {}

type RepoContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatorId string `protobuf:"bytes,1,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RepoContext) Reset() {
	*x = RepoContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repositories_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepoContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepoContext) ProtoMessage() {}

func (x *RepoContext) ProtoReflect() protoreflect.Message {
	mi := &file_repositories_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepoContext.ProtoReflect.Descriptor instead.
func (*RepoContext) Descriptor() ([]byte, []int) {
	return file_repositories_proto_rawDescGZIP(), []int{1}
}

func (x *RepoContext) GetCreatorId() string {
	if x != nil {
		return x.CreatorId
	}
	return ""
}

func (x *RepoContext) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Repository struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Url  string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Repository) Reset() {
	*x = Repository{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repositories_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Repository) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Repository) ProtoMessage() {}

func (x *Repository) ProtoReflect() protoreflect.Message {
	mi := &file_repositories_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Repository.ProtoReflect.Descriptor instead.
func (*Repository) Descriptor() ([]byte, []int) {
	return file_repositories_proto_rawDescGZIP(), []int{2}
}

func (x *Repository) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Repository) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Repository) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type RepoCreateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repo *Repository `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
	Size int32       `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *RepoCreateReply) Reset() {
	*x = RepoCreateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repositories_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepoCreateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepoCreateReply) ProtoMessage() {}

func (x *RepoCreateReply) ProtoReflect() protoreflect.Message {
	mi := &file_repositories_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepoCreateReply.ProtoReflect.Descriptor instead.
func (*RepoCreateReply) Descriptor() ([]byte, []int) {
	return file_repositories_proto_rawDescGZIP(), []int{3}
}

func (x *RepoCreateReply) GetRepo() *Repository {
	if x != nil {
		return x.Repo
	}
	return nil
}

func (x *RepoCreateReply) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

var File_repositories_proto protoreflect.FileDescriptor

var file_repositories_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x11, 0x52, 0x65, 0x70, 0x6f, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x52, 0x65, 0x70,
	0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x22, 0x40, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x42, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x46, 0x0a, 0x0f, 0x52, 0x65, 0x70, 0x6f, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1f, 0x0a, 0x04, 0x72, 0x65,
	0x70, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x32,
	0x3e, 0x0a, 0x04, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x36, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x12, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x52, 0x65, 0x70, 0x6f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x28, 0x01, 0x42,
	0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x63,
	0x68, 0x61, 0x70, 0x39, 0x2f, 0x62, 0x69, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2d, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_repositories_proto_rawDescOnce sync.Once
	file_repositories_proto_rawDescData = file_repositories_proto_rawDesc
)

func file_repositories_proto_rawDescGZIP() []byte {
	file_repositories_proto_rawDescOnce.Do(func() {
		file_repositories_proto_rawDescData = protoimpl.X.CompressGZIP(file_repositories_proto_rawDescData)
	})
	return file_repositories_proto_rawDescData
}

var file_repositories_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_repositories_proto_goTypes = []interface{}{
	(*RepoCreateRequest)(nil), // 0: RepoCreateRequest
	(*RepoContext)(nil),       // 1: RepoContext
	(*Repository)(nil),        // 2: Repository
	(*RepoCreateReply)(nil),   // 3: RepoCreateReply
}
var file_repositories_proto_depIdxs = []int32{
	1, // 0: RepoCreateRequest.context:type_name -> RepoContext
	2, // 1: RepoCreateReply.repo:type_name -> Repository
	0, // 2: Repo.CreateRepo:input_type -> RepoCreateRequest
	3, // 3: Repo.CreateRepo:output_type -> RepoCreateReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_repositories_proto_init() }
func file_repositories_proto_init() {
	if File_repositories_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_repositories_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepoCreateRequest); i {
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
		file_repositories_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepoContext); i {
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
		file_repositories_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Repository); i {
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
		file_repositories_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepoCreateReply); i {
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
	file_repositories_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RepoCreateRequest_Context)(nil),
		(*RepoCreateRequest_Data)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_repositories_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_repositories_proto_goTypes,
		DependencyIndexes: file_repositories_proto_depIdxs,
		MessageInfos:      file_repositories_proto_msgTypes,
	}.Build()
	File_repositories_proto = out.File
	file_repositories_proto_rawDesc = nil
	file_repositories_proto_goTypes = nil
	file_repositories_proto_depIdxs = nil
}
