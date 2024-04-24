// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.1
// source: query.proto

package query

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// trpc create -f --protofile=query.proto --protocol=trpc --rpconly --nogomod --mock=false
type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryText string `protobuf:"bytes,1,opt,name=queryText,proto3" json:"queryText,omitempty"`
	Page      int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`       // 新增分页参数：页码
	PerPage   int32  `protobuf:"varint,3,opt,name=perPage,proto3" json:"perPage,omitempty"` // 新增分页参数：每页数量
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{0}
}

func (x *QueryRequest) GetQueryText() string {
	if x != nil {
		return x.QueryText
	}
	return ""
}

func (x *QueryRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *QueryRequest) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

type RowData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AskText   string `protobuf:"bytes,2,opt,name=askText,proto3" json:"askText,omitempty"`
	TitleText string `protobuf:"bytes,3,opt,name=titleText,proto3" json:"titleText,omitempty"`
	ReplyText string `protobuf:"bytes,4,opt,name=replyText,proto3" json:"replyText,omitempty"`
}

func (x *RowData) Reset() {
	*x = RowData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RowData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RowData) ProtoMessage() {}

func (x *RowData) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RowData.ProtoReflect.Descriptor instead.
func (*RowData) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{1}
}

func (x *RowData) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RowData) GetAskText() string {
	if x != nil {
		return x.AskText
	}
	return ""
}

func (x *RowData) GetTitleText() string {
	if x != nil {
		return x.TitleText
	}
	return ""
}

func (x *RowData) GetReplyText() string {
	if x != nil {
		return x.ReplyText
	}
	return ""
}

type QueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Duration string     `protobuf:"bytes,1,opt,name=duration,proto3" json:"duration,omitempty"`
	Count    int32      `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Rows     []*RowData `protobuf:"bytes,3,rep,name=rows,proto3" json:"rows,omitempty"`
	Error    string     `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *QueryResponse) Reset() {
	*x = QueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResponse) ProtoMessage() {}

func (x *QueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResponse.ProtoReflect.Descriptor instead.
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{2}
}

func (x *QueryResponse) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *QueryResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *QueryResponse) GetRows() []*RowData {
	if x != nil {
		return x.Rows
	}
	return nil
}

func (x *QueryResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_query_proto protoreflect.FileDescriptor

var file_query_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x64,
	0x65, 0x6d, 0x6f, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x5a, 0x0a, 0x0c, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x54, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x54, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x65,
	0x72, 0x50, 0x61, 0x67, 0x65, 0x22, 0x6f, 0x0a, 0x07, 0x52, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x73, 0x6b, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x73, 0x6b, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x54, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6c,
	0x79, 0x54, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x70,
	0x6c, 0x79, 0x54, 0x65, 0x78, 0x74, 0x22, 0x80, 0x01, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x72, 0x6f,
	0x77, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x52, 0x6f, 0x77, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x72,
	0x6f, 0x77, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x4c, 0x0a, 0x0c, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x05, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x18, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64,
	0x65, 0x6d, 0x6f, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1d, 0x5a, 0x1b, 0x74, 0x72, 0x70, 0x63, 0x2d,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_query_proto_rawDescOnce sync.Once
	file_query_proto_rawDescData = file_query_proto_rawDesc
)

func file_query_proto_rawDescGZIP() []byte {
	file_query_proto_rawDescOnce.Do(func() {
		file_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_query_proto_rawDescData)
	})
	return file_query_proto_rawDescData
}

var file_query_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_query_proto_goTypes = []interface{}{
	(*QueryRequest)(nil),  // 0: demo.query.QueryRequest
	(*RowData)(nil),       // 1: demo.query.RowData
	(*QueryResponse)(nil), // 2: demo.query.QueryResponse
}
var file_query_proto_depIdxs = []int32{
	1, // 0: demo.query.QueryResponse.rows:type_name -> demo.query.RowData
	0, // 1: demo.query.QueryService.Query:input_type -> demo.query.QueryRequest
	2, // 2: demo.query.QueryService.Query:output_type -> demo.query.QueryResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_query_proto_init() }
func file_query_proto_init() {
	if File_query_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRequest); i {
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
		file_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RowData); i {
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
		file_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryResponse); i {
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
			RawDescriptor: file_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_query_proto_goTypes,
		DependencyIndexes: file_query_proto_depIdxs,
		MessageInfos:      file_query_proto_msgTypes,
	}.Build()
	File_query_proto = out.File
	file_query_proto_rawDesc = nil
	file_query_proto_goTypes = nil
	file_query_proto_depIdxs = nil
}
