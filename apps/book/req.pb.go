// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: apps/book/pb/req.proto

package book

import (
	request "github.com/infraboard/mcube/http/request"
	request1 "github.com/infraboard/mcube/pb/request"
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

type CreateBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 创建人
	// @gotags: json:"create_by"
	CreateBy string `protobuf:"bytes,1,opt,name=create_by,json=createBy,proto3" json:"create_by"`
	// 名称
	// @gotags: json:"name" validate:"required"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" validate:"required"`
	// 作者
	// @gotags: json:"author" validate:"required"
	Author string `protobuf:"bytes,3,opt,name=author,proto3" json:"author" validate:"required"`
}

func (x *CreateBookRequest) Reset() {
	*x = CreateBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_book_pb_req_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookRequest) ProtoMessage() {}

func (x *CreateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_book_pb_req_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookRequest.ProtoReflect.Descriptor instead.
func (*CreateBookRequest) Descriptor() ([]byte, []int) {
	return file_apps_book_pb_req_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBookRequest) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *CreateBookRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateBookRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type QueryBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数, <package_name>.<message_name>
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// 关键字参数
	// @gotags: json:"keywords"
	Keywords string `protobuf:"bytes,2,opt,name=keywords,proto3" json:"keywords"`
}

func (x *QueryBookRequest) Reset() {
	*x = QueryBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_book_pb_req_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryBookRequest) ProtoMessage() {}

func (x *QueryBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_book_pb_req_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryBookRequest.ProtoReflect.Descriptor instead.
func (*QueryBookRequest) Descriptor() ([]byte, []int) {
	return file_apps_book_pb_req_proto_rawDescGZIP(), []int{1}
}

func (x *QueryBookRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryBookRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

type DescribeBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// book id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DescribeBookRequest) Reset() {
	*x = DescribeBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_book_pb_req_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeBookRequest) ProtoMessage() {}

func (x *DescribeBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_book_pb_req_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeBookRequest.ProtoReflect.Descriptor instead.
func (*DescribeBookRequest) Descriptor() ([]byte, []int) {
	return file_apps_book_pb_req_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeBookRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// book id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 更新模式
	// @gotags: json:"update_mode"
	UpdateMode request1.UpdateMode `protobuf:"varint,2,opt,name=update_mode,json=updateMode,proto3,enum=infraboard.mcube.request.UpdateMode" json:"update_mode"`
	// 更新人
	// @gotags: json:"update_by"
	UpdateBy string `protobuf:"bytes,3,opt,name=update_by,json=updateBy,proto3" json:"update_by"`
	// 更新时间
	// @gotags: json:"update_at"
	UpdateAt int64 `protobuf:"varint,4,opt,name=update_at,json=updateAt,proto3" json:"update_at"`
	// 更新的书本信息
	// @gotags: json:"data"
	Data *CreateBookRequest `protobuf:"bytes,5,opt,name=data,proto3" json:"data"`
}

func (x *UpdateBookRequest) Reset() {
	*x = UpdateBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_book_pb_req_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookRequest) ProtoMessage() {}

func (x *UpdateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_book_pb_req_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookRequest.ProtoReflect.Descriptor instead.
func (*UpdateBookRequest) Descriptor() ([]byte, []int) {
	return file_apps_book_pb_req_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateBookRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateBookRequest) GetUpdateMode() request1.UpdateMode {
	if x != nil {
		return x.UpdateMode
	}
	return request1.UpdateMode(0)
}

func (x *UpdateBookRequest) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *UpdateBookRequest) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *UpdateBookRequest) GetData() *CreateBookRequest {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// book id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DeleteBookRequest) Reset() {
	*x = DeleteBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_book_pb_req_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookRequest) ProtoMessage() {}

func (x *DeleteBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_book_pb_req_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookRequest.ProtoReflect.Descriptor instead.
func (*DeleteBookRequest) Descriptor() ([]byte, []int) {
	return file_apps_book_pb_req_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteBookRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_apps_book_pb_req_proto protoreflect.FileDescriptor

var file_apps_book_pb_req_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x72,
	0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65,
	0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65,
	0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x66, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x22,
	0x25, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xd6, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x45, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x24, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x63, 0x75, 0x62, 0x65, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79,
	0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x30, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x65,
	0x6d, 0x6f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x74,
	0x66, 0x75, 0x6c, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2d, 0x67, 0x37, 0x2f,
	0x61, 0x70, 0x70, 0x73, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_apps_book_pb_req_proto_rawDescOnce sync.Once
	file_apps_book_pb_req_proto_rawDescData = file_apps_book_pb_req_proto_rawDesc
)

func file_apps_book_pb_req_proto_rawDescGZIP() []byte {
	file_apps_book_pb_req_proto_rawDescOnce.Do(func() {
		file_apps_book_pb_req_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_book_pb_req_proto_rawDescData)
	})
	return file_apps_book_pb_req_proto_rawDescData
}

var file_apps_book_pb_req_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_apps_book_pb_req_proto_goTypes = []interface{}{
	(*CreateBookRequest)(nil),   // 0: demo.book.CreateBookRequest
	(*QueryBookRequest)(nil),    // 1: demo.book.QueryBookRequest
	(*DescribeBookRequest)(nil), // 2: demo.book.DescribeBookRequest
	(*UpdateBookRequest)(nil),   // 3: demo.book.UpdateBookRequest
	(*DeleteBookRequest)(nil),   // 4: demo.book.DeleteBookRequest
	(*request.PageRequest)(nil), // 5: infraboard.mcube.page.PageRequest
	(request1.UpdateMode)(0),    // 6: infraboard.mcube.request.UpdateMode
}
var file_apps_book_pb_req_proto_depIdxs = []int32{
	5, // 0: demo.book.QueryBookRequest.page:type_name -> infraboard.mcube.page.PageRequest
	6, // 1: demo.book.UpdateBookRequest.update_mode:type_name -> infraboard.mcube.request.UpdateMode
	0, // 2: demo.book.UpdateBookRequest.data:type_name -> demo.book.CreateBookRequest
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_apps_book_pb_req_proto_init() }
func file_apps_book_pb_req_proto_init() {
	if File_apps_book_pb_req_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_book_pb_req_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookRequest); i {
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
		file_apps_book_pb_req_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryBookRequest); i {
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
		file_apps_book_pb_req_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeBookRequest); i {
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
		file_apps_book_pb_req_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBookRequest); i {
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
		file_apps_book_pb_req_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBookRequest); i {
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
			RawDescriptor: file_apps_book_pb_req_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_book_pb_req_proto_goTypes,
		DependencyIndexes: file_apps_book_pb_req_proto_depIdxs,
		MessageInfos:      file_apps_book_pb_req_proto_msgTypes,
	}.Build()
	File_apps_book_pb_req_proto = out.File
	file_apps_book_pb_req_proto_rawDesc = nil
	file_apps_book_pb_req_proto_goTypes = nil
	file_apps_book_pb_req_proto_depIdxs = nil
}
