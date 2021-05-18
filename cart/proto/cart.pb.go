// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cart.proto

package cart

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CartInfo struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProductId            int64    `protobuf:"varint,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	SizeId               int64    `protobuf:"varint,4,opt,name=size_id,json=sizeId,proto3" json:"size_id,omitempty"`
	Num                  int64    `protobuf:"varint,5,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CartInfo) Reset()         { *m = CartInfo{} }
func (m *CartInfo) String() string { return proto.CompactTextString(m) }
func (*CartInfo) ProtoMessage()    {}
func (*CartInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf731a5c8f9a516f, []int{0}
}

func (m *CartInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CartInfo.Unmarshal(m, b)
}
func (m *CartInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CartInfo.Marshal(b, m, deterministic)
}
func (m *CartInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CartInfo.Merge(m, src)
}
func (m *CartInfo) XXX_Size() int {
	return xxx_messageInfo_CartInfo.Size(m)
}
func (m *CartInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CartInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CartInfo proto.InternalMessageInfo

func (m *CartInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CartInfo) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CartInfo) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *CartInfo) GetSizeId() int64 {
	if m != nil {
		return m.SizeId
	}
	return 0
}

func (m *CartInfo) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

type ResponseAdd struct {
	CartId               int64    `protobuf:"varint,1,opt,name=cart_id,json=cartId,proto3" json:"cart_id,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseAdd) Reset()         { *m = ResponseAdd{} }
func (m *ResponseAdd) String() string { return proto.CompactTextString(m) }
func (*ResponseAdd) ProtoMessage()    {}
func (*ResponseAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf731a5c8f9a516f, []int{1}
}

func (m *ResponseAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseAdd.Unmarshal(m, b)
}
func (m *ResponseAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseAdd.Marshal(b, m, deterministic)
}
func (m *ResponseAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseAdd.Merge(m, src)
}
func (m *ResponseAdd) XXX_Size() int {
	return xxx_messageInfo_ResponseAdd.Size(m)
}
func (m *ResponseAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseAdd proto.InternalMessageInfo

func (m *ResponseAdd) GetCartId() int64 {
	if m != nil {
		return m.CartId
	}
	return 0
}

func (m *ResponseAdd) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Clean struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Clean) Reset()         { *m = Clean{} }
func (m *Clean) String() string { return proto.CompactTextString(m) }
func (*Clean) ProtoMessage()    {}
func (*Clean) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf731a5c8f9a516f, []int{2}
}

func (m *Clean) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Clean.Unmarshal(m, b)
}
func (m *Clean) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Clean.Marshal(b, m, deterministic)
}
func (m *Clean) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Clean.Merge(m, src)
}
func (m *Clean) XXX_Size() int {
	return xxx_messageInfo_Clean.Size(m)
}
func (m *Clean) XXX_DiscardUnknown() {
	xxx_messageInfo_Clean.DiscardUnknown(m)
}

var xxx_messageInfo_Clean proto.InternalMessageInfo

func (m *Clean) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type Response struct {
	Meg                  string   `protobuf:"bytes,1,opt,name=meg,proto3" json:"meg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf731a5c8f9a516f, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMeg() string {
	if m != nil {
		return m.Meg
	}
	return ""
}

type Item struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ChangeNum            int64    `protobuf:"varint,2,opt,name=change_num,json=changeNum,proto3" json:"change_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf731a5c8f9a516f, []int{4}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Item) GetChangeNum() int64 {
	if m != nil {
		return m.ChangeNum
	}
	return 0
}

type CartID struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CartID) Reset()         { *m = CartID{} }
func (m *CartID) String() string { return proto.CompactTextString(m) }
func (*CartID) ProtoMessage()    {}
func (*CartID) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf731a5c8f9a516f, []int{5}
}

func (m *CartID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CartID.Unmarshal(m, b)
}
func (m *CartID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CartID.Marshal(b, m, deterministic)
}
func (m *CartID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CartID.Merge(m, src)
}
func (m *CartID) XXX_Size() int {
	return xxx_messageInfo_CartID.Size(m)
}
func (m *CartID) XXX_DiscardUnknown() {
	xxx_messageInfo_CartID.DiscardUnknown(m)
}

var xxx_messageInfo_CartID proto.InternalMessageInfo

func (m *CartID) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CartFindAll struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CartFindAll) Reset()         { *m = CartFindAll{} }
func (m *CartFindAll) String() string { return proto.CompactTextString(m) }
func (*CartFindAll) ProtoMessage()    {}
func (*CartFindAll) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf731a5c8f9a516f, []int{6}
}

func (m *CartFindAll) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CartFindAll.Unmarshal(m, b)
}
func (m *CartFindAll) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CartFindAll.Marshal(b, m, deterministic)
}
func (m *CartFindAll) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CartFindAll.Merge(m, src)
}
func (m *CartFindAll) XXX_Size() int {
	return xxx_messageInfo_CartFindAll.Size(m)
}
func (m *CartFindAll) XXX_DiscardUnknown() {
	xxx_messageInfo_CartFindAll.DiscardUnknown(m)
}

var xxx_messageInfo_CartFindAll proto.InternalMessageInfo

func (m *CartFindAll) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type CartAll struct {
	CartInfo             []*CartInfo `protobuf:"bytes,1,rep,name=cart_info,json=cartInfo,proto3" json:"cart_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CartAll) Reset()         { *m = CartAll{} }
func (m *CartAll) String() string { return proto.CompactTextString(m) }
func (*CartAll) ProtoMessage()    {}
func (*CartAll) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf731a5c8f9a516f, []int{7}
}

func (m *CartAll) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CartAll.Unmarshal(m, b)
}
func (m *CartAll) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CartAll.Marshal(b, m, deterministic)
}
func (m *CartAll) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CartAll.Merge(m, src)
}
func (m *CartAll) XXX_Size() int {
	return xxx_messageInfo_CartAll.Size(m)
}
func (m *CartAll) XXX_DiscardUnknown() {
	xxx_messageInfo_CartAll.DiscardUnknown(m)
}

var xxx_messageInfo_CartAll proto.InternalMessageInfo

func (m *CartAll) GetCartInfo() []*CartInfo {
	if m != nil {
		return m.CartInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*CartInfo)(nil), "go.micro.service.cart.CartInfo")
	proto.RegisterType((*ResponseAdd)(nil), "go.micro.service.cart.ResponseAdd")
	proto.RegisterType((*Clean)(nil), "go.micro.service.cart.Clean")
	proto.RegisterType((*Response)(nil), "go.micro.service.cart.Response")
	proto.RegisterType((*Item)(nil), "go.micro.service.cart.Item")
	proto.RegisterType((*CartID)(nil), "go.micro.service.cart.CartID")
	proto.RegisterType((*CartFindAll)(nil), "go.micro.service.cart.CartFindAll")
	proto.RegisterType((*CartAll)(nil), "go.micro.service.cart.CartAll")
}

func init() { proto.RegisterFile("cart.proto", fileDescriptor_bf731a5c8f9a516f) }

var fileDescriptor_bf731a5c8f9a516f = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x51, 0x8b, 0xd3, 0x40,
	0x10, 0x6e, 0x9a, 0x5c, 0xd2, 0x4c, 0xe1, 0x90, 0x05, 0x31, 0xd4, 0xab, 0x57, 0xf6, 0x41, 0xee,
	0x29, 0xc2, 0x89, 0x20, 0xe8, 0x4b, 0xee, 0x42, 0x4b, 0x40, 0x8a, 0xe4, 0xd1, 0x97, 0x12, 0xb3,
	0xd3, 0x18, 0x48, 0x76, 0xcb, 0x26, 0x11, 0x14, 0xfc, 0xa3, 0xfe, 0x1a, 0x99, 0x4d, 0x2b, 0x45,
	0x4d, 0xdb, 0x07, 0xdf, 0x66, 0xbf, 0xf9, 0xbe, 0x8f, 0x6f, 0x66, 0x12, 0x80, 0x3c, 0xd3, 0x6d,
	0xb8, 0xd3, 0xaa, 0x55, 0xec, 0x69, 0xa1, 0xc2, 0xba, 0xcc, 0xb5, 0x0a, 0x1b, 0xd4, 0x5f, 0xcb,
	0x1c, 0x43, 0x6a, 0xf2, 0x1f, 0x30, 0x79, 0xcc, 0x74, 0x9b, 0xc8, 0xad, 0x62, 0xd7, 0x30, 0x2e,
	0x45, 0x60, 0x2d, 0xac, 0x3b, 0x3b, 0x1d, 0x97, 0x82, 0x3d, 0x03, 0xaf, 0x6b, 0x50, 0x6f, 0x4a,
	0x11, 0x8c, 0x0d, 0xe8, 0xd2, 0x33, 0x11, 0x6c, 0x0e, 0xb0, 0xd3, 0x4a, 0x74, 0x79, 0x4b, 0x3d,
	0xdb, 0xf4, 0xfc, 0x3d, 0x92, 0x18, 0x5d, 0x53, 0x7e, 0x47, 0xea, 0x39, 0xbd, 0x8e, 0x9e, 0x89,
	0x60, 0x4f, 0xc0, 0x96, 0x5d, 0x1d, 0x5c, 0x19, 0x90, 0x4a, 0xfe, 0x16, 0xa6, 0x29, 0x36, 0x3b,
	0x25, 0x1b, 0x8c, 0x84, 0x51, 0x52, 0xaa, 0xcd, 0xef, 0x18, 0x2e, 0x3d, 0x7b, 0x65, 0xdd, 0x14,
	0x26, 0x86, 0x9f, 0x52, 0xc9, 0x17, 0x70, 0xf5, 0x58, 0x61, 0x26, 0x8f, 0x53, 0x5a, 0xc7, 0x29,
	0xf9, 0x0d, 0x4c, 0x0e, 0xde, 0x46, 0x8f, 0x85, 0x21, 0x90, 0x1e, 0x0b, 0xfe, 0x06, 0x9c, 0xa4,
	0xc5, 0xfa, 0xaf, 0xa1, 0xe7, 0x00, 0xf9, 0x97, 0x4c, 0x16, 0xb8, 0xa1, 0xa8, 0xfd, 0xdc, 0x7e,
	0x8f, 0xac, 0xbb, 0x9a, 0x07, 0xe0, 0x9a, 0x7d, 0xc5, 0x7f, 0x0a, 0xf9, 0x4b, 0x98, 0x52, 0x67,
	0x59, 0x4a, 0x11, 0x55, 0xd5, 0x70, 0xac, 0x15, 0x78, 0xc4, 0x23, 0xce, 0x7b, 0xf0, 0xfb, 0x71,
	0xe5, 0x56, 0x05, 0xd6, 0xc2, 0xbe, 0x9b, 0xde, 0xdf, 0x86, 0xff, 0xbc, 0x53, 0x78, 0x38, 0x52,
	0x3a, 0xc9, 0xf7, 0xd5, 0xfd, 0x4f, 0x1b, 0x1c, 0x82, 0xd9, 0x47, 0xf0, 0x22, 0x21, 0x4c, 0x79,
	0x4e, 0x3e, 0xe3, 0x03, 0x84, 0xa3, 0x2b, 0xf0, 0x11, 0xfb, 0x00, 0xbe, 0x59, 0xae, 0xf1, 0xbc,
	0x19, 0xf2, 0x24, 0xc6, 0xec, 0xf6, 0x8c, 0x21, 0x1f, 0xb1, 0x25, 0x38, 0x89, 0xcc, 0x35, 0x7b,
	0x3e, 0x40, 0xa5, 0x3b, 0x5c, 0xe8, 0x13, 0xe3, 0x7f, 0xf0, 0x49, 0xe1, 0x3a, 0xc6, 0x0a, 0x5b,
	0x24, 0xc1, 0xc3, 0xb7, 0x24, 0x66, 0xf3, 0x53, 0x6b, 0x8b, 0x2f, 0xf1, 0x5c, 0x83, 0xbb, 0x42,
	0x73, 0x54, 0x7e, 0xc2, 0x6b, 0xff, 0x71, 0xcc, 0x5e, 0x9c, 0xe0, 0x44, 0x55, 0xc5, 0x47, 0x0f,
	0xfe, 0x27, 0x2f, 0x7c, 0xf5, 0x8e, 0xc0, 0xcf, 0xae, 0xf9, 0x81, 0x5f, 0xff, 0x0a, 0x00, 0x00,
	0xff, 0xff, 0x22, 0x0f, 0x2e, 0xfc, 0xce, 0x03, 0x00, 0x00,
}
