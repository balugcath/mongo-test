// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: product_price.proto

package api

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ProductPriceListParams_SortOrder int32

const (
	ProductPriceListParams_NA   ProductPriceListParams_SortOrder = 0
	ProductPriceListParams_ASC  ProductPriceListParams_SortOrder = 1
	ProductPriceListParams_DESC ProductPriceListParams_SortOrder = -1
)

// Enum value maps for ProductPriceListParams_SortOrder.
var (
	ProductPriceListParams_SortOrder_name = map[int32]string{
		0:  "NA",
		1:  "ASC",
		-1: "DESC",
	}
	ProductPriceListParams_SortOrder_value = map[string]int32{
		"NA":   0,
		"ASC":  1,
		"DESC": -1,
	}
)

func (x ProductPriceListParams_SortOrder) Enum() *ProductPriceListParams_SortOrder {
	p := new(ProductPriceListParams_SortOrder)
	*p = x
	return p
}

func (x ProductPriceListParams_SortOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProductPriceListParams_SortOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_product_price_proto_enumTypes[0].Descriptor()
}

func (ProductPriceListParams_SortOrder) Type() protoreflect.EnumType {
	return &file_product_price_proto_enumTypes[0]
}

func (x ProductPriceListParams_SortOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProductPriceListParams_SortOrder.Descriptor instead.
func (ProductPriceListParams_SortOrder) EnumDescriptor() ([]byte, []int) {
	return file_product_price_proto_rawDescGZIP(), []int{0, 0}
}

type ProductPriceListParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SortOrder ProductPriceListParams_SortOrder `protobuf:"varint,1,opt,name=sort_order,json=sortOrder,proto3,enum=api.ProductPriceListParams_SortOrder" json:"sort_order,omitempty"`
	// ../pkg/types/types.go
	SortField string `protobuf:"bytes,2,opt,name=sort_field,json=sortField,proto3" json:"sort_field,omitempty"`
	Page      int64  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	LenPage   int64  `protobuf:"varint,4,opt,name=len_page,json=lenPage,proto3" json:"len_page,omitempty"`
}

func (x *ProductPriceListParams) Reset() {
	*x = ProductPriceListParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_price_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductPriceListParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductPriceListParams) ProtoMessage() {}

func (x *ProductPriceListParams) ProtoReflect() protoreflect.Message {
	mi := &file_product_price_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductPriceListParams.ProtoReflect.Descriptor instead.
func (*ProductPriceListParams) Descriptor() ([]byte, []int) {
	return file_product_price_proto_rawDescGZIP(), []int{0}
}

func (x *ProductPriceListParams) GetSortOrder() ProductPriceListParams_SortOrder {
	if x != nil {
		return x.SortOrder
	}
	return ProductPriceListParams_NA
}

func (x *ProductPriceListParams) GetSortField() string {
	if x != nil {
		return x.SortField
	}
	return ""
}

func (x *ProductPriceListParams) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ProductPriceListParams) GetLenPage() int64 {
	if x != nil {
		return x.LenPage
	}
	return 0
}

type ProductPriceListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductName string               `protobuf:"bytes,1,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	Price       float64              `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
	LastUpdate  *timestamp.Timestamp `protobuf:"bytes,3,opt,name=last_update,json=lastUpdate,proto3" json:"last_update,omitempty"`
	CountUpdate int64                `protobuf:"varint,4,opt,name=count_update,json=countUpdate,proto3" json:"count_update,omitempty"`
}

func (x *ProductPriceListReply) Reset() {
	*x = ProductPriceListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_price_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductPriceListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductPriceListReply) ProtoMessage() {}

func (x *ProductPriceListReply) ProtoReflect() protoreflect.Message {
	mi := &file_product_price_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductPriceListReply.ProtoReflect.Descriptor instead.
func (*ProductPriceListReply) Descriptor() ([]byte, []int) {
	return file_product_price_proto_rawDescGZIP(), []int{1}
}

func (x *ProductPriceListReply) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *ProductPriceListReply) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ProductPriceListReply) GetLastUpdate() *timestamp.Timestamp {
	if x != nil {
		return x.LastUpdate
	}
	return nil
}

func (x *ProductPriceListReply) GetCountUpdate() int64 {
	if x != nil {
		return x.CountUpdate
	}
	return 0
}

type ProductPriceFetchParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ProductPriceFetchParams) Reset() {
	*x = ProductPriceFetchParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_price_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductPriceFetchParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductPriceFetchParams) ProtoMessage() {}

func (x *ProductPriceFetchParams) ProtoReflect() protoreflect.Message {
	mi := &file_product_price_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductPriceFetchParams.ProtoReflect.Descriptor instead.
func (*ProductPriceFetchParams) Descriptor() ([]byte, []int) {
	return file_product_price_proto_rawDescGZIP(), []int{2}
}

func (x *ProductPriceFetchParams) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_product_price_proto protoreflect.FileDescriptor

var file_product_price_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x01, 0x0a, 0x16, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x45, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x09, 0x73, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f,
	0x72, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x6c, 0x65, 0x6e, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x6c, 0x65, 0x6e, 0x50, 0x61, 0x67, 0x65, 0x22, 0x30, 0x0a, 0x0a, 0x53, 0x6f, 0x72, 0x74,
	0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x06, 0x0a, 0x02, 0x4e, 0x41, 0x10, 0x00, 0x12, 0x07,
	0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x04, 0x44, 0x45, 0x53, 0x43, 0x10,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x22, 0xb0, 0x01, 0x0a, 0x15, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a,
	0x0b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x2b, 0x0a,
	0x17, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x32, 0x96, 0x01, 0x0a, 0x0c, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x12, 0x3f, 0x0a, 0x05, 0x46, 0x65, 0x74, 0x63, 0x68, 0x12, 0x1c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x62, 0x61, 0x6c, 0x75, 0x67, 0x63, 0x61, 0x74, 0x68, 0x2f, 0x6d, 0x6f, 0x6e, 0x67,
	0x6f, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_product_price_proto_rawDescOnce sync.Once
	file_product_price_proto_rawDescData = file_product_price_proto_rawDesc
)

func file_product_price_proto_rawDescGZIP() []byte {
	file_product_price_proto_rawDescOnce.Do(func() {
		file_product_price_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_price_proto_rawDescData)
	})
	return file_product_price_proto_rawDescData
}

var file_product_price_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_product_price_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_product_price_proto_goTypes = []interface{}{
	(ProductPriceListParams_SortOrder)(0), // 0: api.ProductPriceListParams.Sort_order
	(*ProductPriceListParams)(nil),        // 1: api.ProductPriceListParams
	(*ProductPriceListReply)(nil),         // 2: api.ProductPriceListReply
	(*ProductPriceFetchParams)(nil),       // 3: api.ProductPriceFetchParams
	(*timestamp.Timestamp)(nil),           // 4: google.protobuf.Timestamp
	(*empty.Empty)(nil),                   // 5: google.protobuf.Empty
}
var file_product_price_proto_depIdxs = []int32{
	0, // 0: api.ProductPriceListParams.sort_order:type_name -> api.ProductPriceListParams.Sort_order
	4, // 1: api.ProductPriceListReply.last_update:type_name -> google.protobuf.Timestamp
	1, // 2: api.ProductPrice.List:input_type -> api.ProductPriceListParams
	3, // 3: api.ProductPrice.Fetch:input_type -> api.ProductPriceFetchParams
	2, // 4: api.ProductPrice.List:output_type -> api.ProductPriceListReply
	5, // 5: api.ProductPrice.Fetch:output_type -> google.protobuf.Empty
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_product_price_proto_init() }
func file_product_price_proto_init() {
	if File_product_price_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_price_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductPriceListParams); i {
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
		file_product_price_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductPriceListReply); i {
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
		file_product_price_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductPriceFetchParams); i {
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
			RawDescriptor: file_product_price_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_price_proto_goTypes,
		DependencyIndexes: file_product_price_proto_depIdxs,
		EnumInfos:         file_product_price_proto_enumTypes,
		MessageInfos:      file_product_price_proto_msgTypes,
	}.Build()
	File_product_price_proto = out.File
	file_product_price_proto_rawDesc = nil
	file_product_price_proto_goTypes = nil
	file_product_price_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProductPriceClient is the client API for ProductPrice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductPriceClient interface {
	List(ctx context.Context, opts ...grpc.CallOption) (ProductPrice_ListClient, error)
	Fetch(ctx context.Context, in *ProductPriceFetchParams, opts ...grpc.CallOption) (*empty.Empty, error)
}

type productPriceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductPriceClient(cc grpc.ClientConnInterface) ProductPriceClient {
	return &productPriceClient{cc}
}

func (c *productPriceClient) List(ctx context.Context, opts ...grpc.CallOption) (ProductPrice_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProductPrice_serviceDesc.Streams[0], "/api.ProductPrice/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &productPriceListClient{stream}
	return x, nil
}

type ProductPrice_ListClient interface {
	Send(*ProductPriceListParams) error
	Recv() (*ProductPriceListReply, error)
	grpc.ClientStream
}

type productPriceListClient struct {
	grpc.ClientStream
}

func (x *productPriceListClient) Send(m *ProductPriceListParams) error {
	return x.ClientStream.SendMsg(m)
}

func (x *productPriceListClient) Recv() (*ProductPriceListReply, error) {
	m := new(ProductPriceListReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *productPriceClient) Fetch(ctx context.Context, in *ProductPriceFetchParams, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.ProductPrice/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductPriceServer is the server API for ProductPrice service.
type ProductPriceServer interface {
	List(ProductPrice_ListServer) error
	Fetch(context.Context, *ProductPriceFetchParams) (*empty.Empty, error)
}

// UnimplementedProductPriceServer can be embedded to have forward compatible implementations.
type UnimplementedProductPriceServer struct {
}

func (*UnimplementedProductPriceServer) List(ProductPrice_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedProductPriceServer) Fetch(context.Context, *ProductPriceFetchParams) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}

func RegisterProductPriceServer(s *grpc.Server, srv ProductPriceServer) {
	s.RegisterService(&_ProductPrice_serviceDesc, srv)
}

func _ProductPrice_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProductPriceServer).List(&productPriceListServer{stream})
}

type ProductPrice_ListServer interface {
	Send(*ProductPriceListReply) error
	Recv() (*ProductPriceListParams, error)
	grpc.ServerStream
}

type productPriceListServer struct {
	grpc.ServerStream
}

func (x *productPriceListServer) Send(m *ProductPriceListReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *productPriceListServer) Recv() (*ProductPriceListParams, error) {
	m := new(ProductPriceListParams)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ProductPrice_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductPriceFetchParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductPriceServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ProductPrice/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductPriceServer).Fetch(ctx, req.(*ProductPriceFetchParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductPrice_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ProductPrice",
	HandlerType: (*ProductPriceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fetch",
			Handler:    _ProductPrice_Fetch_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _ProductPrice_List_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "product_price.proto",
}
