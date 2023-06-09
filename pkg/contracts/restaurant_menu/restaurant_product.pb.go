// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: restaurant_product.proto

package restaurant

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ProductType int32

const (
	ProductType_PRODUCT_TYPE_UNSPECIFIED ProductType = 0
	ProductType_PRODUCT_TYPE_SALAD       ProductType = 1
	ProductType_PRODUCT_TYPE_GARNISH     ProductType = 2
	ProductType_PRODUCT_TYPE_MEAT        ProductType = 3
	ProductType_PRODUCT_TYPE_SOUP        ProductType = 4
	ProductType_PRODUCT_TYPE_DRINK       ProductType = 5
	ProductType_PRODUCT_TYPE_DESSERT     ProductType = 6
)

// Enum value maps for ProductType.
var (
	ProductType_name = map[int32]string{
		0: "PRODUCT_TYPE_UNSPECIFIED",
		1: "PRODUCT_TYPE_SALAD",
		2: "PRODUCT_TYPE_GARNISH",
		3: "PRODUCT_TYPE_MEAT",
		4: "PRODUCT_TYPE_SOUP",
		5: "PRODUCT_TYPE_DRINK",
		6: "PRODUCT_TYPE_DESSERT",
	}
	ProductType_value = map[string]int32{
		"PRODUCT_TYPE_UNSPECIFIED": 0,
		"PRODUCT_TYPE_SALAD":       1,
		"PRODUCT_TYPE_GARNISH":     2,
		"PRODUCT_TYPE_MEAT":        3,
		"PRODUCT_TYPE_SOUP":        4,
		"PRODUCT_TYPE_DRINK":       5,
		"PRODUCT_TYPE_DESSERT":     6,
	}
)

func (x ProductType) Enum() *ProductType {
	p := new(ProductType)
	*p = x
	return p
}

func (x ProductType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProductType) Descriptor() protoreflect.EnumDescriptor {
	return file_restaurant_product_proto_enumTypes[0].Descriptor()
}

func (ProductType) Type() protoreflect.EnumType {
	return &file_restaurant_product_proto_enumTypes[0]
}

func (x ProductType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProductType.Descriptor instead.
func (ProductType) EnumDescriptor() ([]byte, []int) {
	return file_restaurant_product_proto_rawDescGZIP(), []int{0}
}

type CreateProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string      `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Type        ProductType `protobuf:"varint,3,opt,name=type,proto3,enum=mediasoft.internship.final.task.contracts.restaurant.ProductType" json:"type,omitempty"`
	Weight      int32       `protobuf:"varint,4,opt,name=weight,proto3" json:"weight,omitempty"`
	Price       float64     `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *CreateProductRequest) Reset() {
	*x = CreateProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductRequest) ProtoMessage() {}

func (x *CreateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductRequest.ProtoReflect.Descriptor instead.
func (*CreateProductRequest) Descriptor() ([]byte, []int) {
	return file_restaurant_product_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateProductRequest) GetType() ProductType {
	if x != nil {
		return x.Type
	}
	return ProductType_PRODUCT_TYPE_UNSPECIFIED
}

func (x *CreateProductRequest) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *CreateProductRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type CreateProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateProductResponse) Reset() {
	*x = CreateProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductResponse) ProtoMessage() {}

func (x *CreateProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductResponse.ProtoReflect.Descriptor instead.
func (*CreateProductResponse) Descriptor() ([]byte, []int) {
	return file_restaurant_product_proto_rawDescGZIP(), []int{1}
}

type GetProductListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetProductListRequest) Reset() {
	*x = GetProductListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductListRequest) ProtoMessage() {}

func (x *GetProductListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductListRequest.ProtoReflect.Descriptor instead.
func (*GetProductListRequest) Descriptor() ([]byte, []int) {
	return file_restaurant_product_proto_rawDescGZIP(), []int{2}
}

type GetProductListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*Product `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *GetProductListResponse) Reset() {
	*x = GetProductListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductListResponse) ProtoMessage() {}

func (x *GetProductListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductListResponse.ProtoReflect.Descriptor instead.
func (*GetProductListResponse) Descriptor() ([]byte, []int) {
	return file_restaurant_product_proto_rawDescGZIP(), []int{3}
}

func (x *GetProductListResponse) GetResult() []*Product {
	if x != nil {
		return x.Result
	}
	return nil
}

type GetProductListByUuidsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuids []string `protobuf:"bytes,1,rep,name=uuids,proto3" json:"uuids,omitempty"`
}

func (x *GetProductListByUuidsRequest) Reset() {
	*x = GetProductListByUuidsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductListByUuidsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductListByUuidsRequest) ProtoMessage() {}

func (x *GetProductListByUuidsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductListByUuidsRequest.ProtoReflect.Descriptor instead.
func (*GetProductListByUuidsRequest) Descriptor() ([]byte, []int) {
	return file_restaurant_product_proto_rawDescGZIP(), []int{4}
}

func (x *GetProductListByUuidsRequest) GetUuids() []string {
	if x != nil {
		return x.Uuids
	}
	return nil
}

type GetProductListByUuidsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*Product `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *GetProductListByUuidsResponse) Reset() {
	*x = GetProductListByUuidsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductListByUuidsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductListByUuidsResponse) ProtoMessage() {}

func (x *GetProductListByUuidsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductListByUuidsResponse.ProtoReflect.Descriptor instead.
func (*GetProductListByUuidsResponse) Descriptor() ([]byte, []int) {
	return file_restaurant_product_proto_rawDescGZIP(), []int{5}
}

func (x *GetProductListByUuidsResponse) GetResult() []*Product {
	if x != nil {
		return x.Result
	}
	return nil
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid        string               `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name        string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Type        ProductType          `protobuf:"varint,4,opt,name=type,proto3,enum=mediasoft.internship.final.task.contracts.restaurant.ProductType" json:"type,omitempty"`
	Weight      int32                `protobuf:"varint,5,opt,name=weight,proto3" json:"weight,omitempty"`
	Price       float64              `protobuf:"fixed64,6,opt,name=price,proto3" json:"price,omitempty"`
	CreatedAt   *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_product_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_product_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_restaurant_product_proto_rawDescGZIP(), []int{6}
}

func (x *Product) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetType() ProductType {
	if x != nil {
		return x.Type
	}
	return ProductType_PRODUCT_TYPE_UNSPECIFIED
}

func (x *Product) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Product) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_restaurant_product_proto protoreflect.FileDescriptor

var file_restaurant_product_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x34, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x73, 0x6f, 0x66, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xea, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5a, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x41, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73,
	0x6f, 0x66, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x66,
	0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x01, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x17, 0x0a,
	0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x6f, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x73, 0x6f, 0x66, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x34, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x79, 0x55, 0x75, 0x69, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x75, 0x75, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x75, 0x75, 0x69, 0x64, 0x73, 0x22, 0x76, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x55, 0x75, 0x69, 0x64, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73,
	0x6f, 0x66, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x66,
	0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x93,
	0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x55, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x41, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x6f, 0x66, 0x74, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x72,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x2a, 0xbd, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x53, 0x41, 0x4c, 0x41, 0x44, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x52,
	0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x41, 0x52, 0x4e, 0x49,
	0x53, 0x48, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x41, 0x54, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x50,
	0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x4f, 0x55, 0x50,
	0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x44, 0x52, 0x49, 0x4e, 0x4b, 0x10, 0x05, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x52,
	0x4f, 0x44, 0x55, 0x43, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x53, 0x53, 0x45,
	0x52, 0x54, 0x10, 0x06, 0x32, 0x8f, 0x05, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc9, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x4a, 0x2e, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x73, 0x6f, 0x66, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4b, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x6f, 0x66,
	0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x66, 0x69, 0x6e,
	0x61, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x73, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f,
	0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x12, 0xc9, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4b, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x6f,
	0x66, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x66, 0x69,
	0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x4c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x6f, 0x66, 0x74, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6c,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e,
	0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x72, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12,
	0xe4, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x79, 0x55, 0x75, 0x69, 0x64, 0x73, 0x12, 0x52, 0x2e, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x73, 0x6f, 0x66, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x79, 0x55, 0x75, 0x69, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x53, 0x2e,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x6f, 0x66, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x73, 0x68, 0x69, 0x70, 0x2e, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x79, 0x55, 0x75, 0x69, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x72, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x2f, 0x75, 0x75, 0x69, 0x64, 0x73, 0x42, 0x4f, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x6f, 0x66, 0x74, 0x2d, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x2f, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x2d,
	0x74, 0x61, 0x73, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x72, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_restaurant_product_proto_rawDescOnce sync.Once
	file_restaurant_product_proto_rawDescData = file_restaurant_product_proto_rawDesc
)

func file_restaurant_product_proto_rawDescGZIP() []byte {
	file_restaurant_product_proto_rawDescOnce.Do(func() {
		file_restaurant_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_restaurant_product_proto_rawDescData)
	})
	return file_restaurant_product_proto_rawDescData
}

var file_restaurant_product_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_restaurant_product_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_restaurant_product_proto_goTypes = []interface{}{
	(ProductType)(0),                      // 0: mediasoft.internship.final.task.contracts.restaurant.ProductType
	(*CreateProductRequest)(nil),          // 1: mediasoft.internship.final.task.contracts.restaurant.CreateProductRequest
	(*CreateProductResponse)(nil),         // 2: mediasoft.internship.final.task.contracts.restaurant.CreateProductResponse
	(*GetProductListRequest)(nil),         // 3: mediasoft.internship.final.task.contracts.restaurant.GetProductListRequest
	(*GetProductListResponse)(nil),        // 4: mediasoft.internship.final.task.contracts.restaurant.GetProductListResponse
	(*GetProductListByUuidsRequest)(nil),  // 5: mediasoft.internship.final.task.contracts.restaurant.GetProductListByUuidsRequest
	(*GetProductListByUuidsResponse)(nil), // 6: mediasoft.internship.final.task.contracts.restaurant.GetProductListByUuidsResponse
	(*Product)(nil),                       // 7: mediasoft.internship.final.task.contracts.restaurant.Product
	(*timestamp.Timestamp)(nil),           // 8: google.protobuf.Timestamp
}
var file_restaurant_product_proto_depIdxs = []int32{
	0, // 0: mediasoft.internship.final.task.contracts.restaurant.CreateProductRequest.type:type_name -> mediasoft.internship.final.task.contracts.restaurant.ProductType
	7, // 1: mediasoft.internship.final.task.contracts.restaurant.GetProductListResponse.result:type_name -> mediasoft.internship.final.task.contracts.restaurant.Product
	7, // 2: mediasoft.internship.final.task.contracts.restaurant.GetProductListByUuidsResponse.result:type_name -> mediasoft.internship.final.task.contracts.restaurant.Product
	0, // 3: mediasoft.internship.final.task.contracts.restaurant.Product.type:type_name -> mediasoft.internship.final.task.contracts.restaurant.ProductType
	8, // 4: mediasoft.internship.final.task.contracts.restaurant.Product.created_at:type_name -> google.protobuf.Timestamp
	1, // 5: mediasoft.internship.final.task.contracts.restaurant.ProductService.CreateProduct:input_type -> mediasoft.internship.final.task.contracts.restaurant.CreateProductRequest
	3, // 6: mediasoft.internship.final.task.contracts.restaurant.ProductService.GetProductList:input_type -> mediasoft.internship.final.task.contracts.restaurant.GetProductListRequest
	5, // 7: mediasoft.internship.final.task.contracts.restaurant.ProductService.GetProductListByUuids:input_type -> mediasoft.internship.final.task.contracts.restaurant.GetProductListByUuidsRequest
	2, // 8: mediasoft.internship.final.task.contracts.restaurant.ProductService.CreateProduct:output_type -> mediasoft.internship.final.task.contracts.restaurant.CreateProductResponse
	4, // 9: mediasoft.internship.final.task.contracts.restaurant.ProductService.GetProductList:output_type -> mediasoft.internship.final.task.contracts.restaurant.GetProductListResponse
	6, // 10: mediasoft.internship.final.task.contracts.restaurant.ProductService.GetProductListByUuids:output_type -> mediasoft.internship.final.task.contracts.restaurant.GetProductListByUuidsResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_restaurant_product_proto_init() }
func file_restaurant_product_proto_init() {
	if File_restaurant_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_restaurant_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductRequest); i {
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
		file_restaurant_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductResponse); i {
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
		file_restaurant_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductListRequest); i {
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
		file_restaurant_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductListResponse); i {
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
		file_restaurant_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductListByUuidsRequest); i {
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
		file_restaurant_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductListByUuidsResponse); i {
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
		file_restaurant_product_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
			RawDescriptor: file_restaurant_product_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_restaurant_product_proto_goTypes,
		DependencyIndexes: file_restaurant_product_proto_depIdxs,
		EnumInfos:         file_restaurant_product_proto_enumTypes,
		MessageInfos:      file_restaurant_product_proto_msgTypes,
	}.Build()
	File_restaurant_product_proto = out.File
	file_restaurant_product_proto_rawDesc = nil
	file_restaurant_product_proto_goTypes = nil
	file_restaurant_product_proto_depIdxs = nil
}
