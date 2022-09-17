// Code generated by protoc-gen-go. DO NOT EDIT.
// source: response_v2.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SimpleResponse struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Info                 string   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleResponse) Reset()         { *m = SimpleResponse{} }
func (m *SimpleResponse) String() string { return proto.CompactTextString(m) }
func (*SimpleResponse) ProtoMessage()    {}
func (*SimpleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_response_v2_468be3765a7a7f37, []int{0}
}
func (m *SimpleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleResponse.Unmarshal(m, b)
}
func (m *SimpleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleResponse.Marshal(b, m, deterministic)
}
func (dst *SimpleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleResponse.Merge(dst, src)
}
func (m *SimpleResponse) XXX_Size() int {
	return xxx_messageInfo_SimpleResponse.Size(m)
}
func (m *SimpleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleResponse proto.InternalMessageInfo

func (m *SimpleResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SimpleResponse) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

type Response struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Info                 string   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	Data                 *any.Any `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_response_v2_468be3765a7a7f37, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *Response) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

type BatchWriteResponse struct {
	Code                 uint32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Info                 string      `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	Size                 uint32      `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Responses            []*Response `protobuf:"bytes,4,rep,name=responses,proto3" json:"responses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BatchWriteResponse) Reset()         { *m = BatchWriteResponse{} }
func (m *BatchWriteResponse) String() string { return proto.CompactTextString(m) }
func (*BatchWriteResponse) ProtoMessage()    {}
func (*BatchWriteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_response_v2_468be3765a7a7f37, []int{2}
}
func (m *BatchWriteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchWriteResponse.Unmarshal(m, b)
}
func (m *BatchWriteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchWriteResponse.Marshal(b, m, deterministic)
}
func (dst *BatchWriteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchWriteResponse.Merge(dst, src)
}
func (m *BatchWriteResponse) XXX_Size() int {
	return xxx_messageInfo_BatchWriteResponse.Size(m)
}
func (m *BatchWriteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchWriteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BatchWriteResponse proto.InternalMessageInfo

func (m *BatchWriteResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BatchWriteResponse) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *BatchWriteResponse) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *BatchWriteResponse) GetResponses() []*Response {
	if m != nil {
		return m.Responses
	}
	return nil
}

type BatchQueryResponse struct {
	Code                 uint32     `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Info                 string     `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	Amount               uint32     `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Size                 uint32     `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	Data                 []*any.Any `protobuf:"bytes,5,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *BatchQueryResponse) Reset()         { *m = BatchQueryResponse{} }
func (m *BatchQueryResponse) String() string { return proto.CompactTextString(m) }
func (*BatchQueryResponse) ProtoMessage()    {}
func (*BatchQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_response_v2_468be3765a7a7f37, []int{3}
}
func (m *BatchQueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchQueryResponse.Unmarshal(m, b)
}
func (m *BatchQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchQueryResponse.Marshal(b, m, deterministic)
}
func (dst *BatchQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchQueryResponse.Merge(dst, src)
}
func (m *BatchQueryResponse) XXX_Size() int {
	return xxx_messageInfo_BatchQueryResponse.Size(m)
}
func (m *BatchQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BatchQueryResponse proto.InternalMessageInfo

func (m *BatchQueryResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BatchQueryResponse) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *BatchQueryResponse) GetAmount() uint32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *BatchQueryResponse) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *BatchQueryResponse) GetData() []*any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*SimpleResponse)(nil), "v2.SimpleResponse")
	proto.RegisterType((*Response)(nil), "v2.Response")
	proto.RegisterType((*BatchWriteResponse)(nil), "v2.BatchWriteResponse")
	proto.RegisterType((*BatchQueryResponse)(nil), "v2.BatchQueryResponse")
}

func init() { proto.RegisterFile("response_v2.proto", fileDescriptor_response_v2_468be3765a7a7f37) }

var fileDescriptor_response_v2_468be3765a7a7f37 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x4a, 0x2d, 0x2e,
	0xc8, 0xcf, 0x2b, 0x4e, 0x8d, 0x2f, 0x33, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a,
	0x33, 0x92, 0x92, 0x4c, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x24, 0x95, 0xa6, 0xe9,
	0x27, 0xe6, 0x55, 0x42, 0xa4, 0x95, 0x2c, 0xb8, 0xf8, 0x82, 0x33, 0x73, 0x0b, 0x72, 0x52, 0x83,
	0xa0, 0x3a, 0x85, 0x84, 0xb8, 0x58, 0x92, 0xf3, 0x53, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x78,
	0x83, 0xc0, 0x6c, 0x90, 0x58, 0x66, 0x5e, 0x5a, 0xbe, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10,
	0x98, 0xad, 0x94, 0xc1, 0xc5, 0x41, 0xaa, 0x1e, 0x21, 0x0d, 0x2e, 0x96, 0x94, 0xc4, 0x92, 0x44,
	0x09, 0x66, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x11, 0x3d, 0x88, 0xbb, 0xf4, 0x60, 0xee, 0xd2, 0x73,
	0xcc, 0xab, 0x0c, 0x02, 0xab, 0xf0, 0x62, 0xe1, 0xe0, 0x11, 0xe0, 0xf7, 0x62, 0xe1, 0x10, 0x10,
	0x10, 0x56, 0xaa, 0xe1, 0x12, 0x72, 0x4a, 0x2c, 0x49, 0xce, 0x08, 0x2f, 0xca, 0x2c, 0x21, 0xd9,
	0x9d, 0x20, 0xb1, 0xe2, 0xcc, 0xaa, 0x54, 0xb0, 0x9d, 0xbc, 0x41, 0x60, 0xb6, 0x90, 0x16, 0x17,
	0x27, 0x2c, 0xa4, 0x8a, 0x25, 0x58, 0x14, 0x98, 0x35, 0xb8, 0x8d, 0x78, 0xf4, 0xca, 0x8c, 0xf4,
	0x60, 0x86, 0x07, 0x21, 0xa4, 0x95, 0x26, 0x31, 0x42, 0xad, 0x0f, 0x2c, 0x4d, 0x2d, 0xaa, 0x24,
	0xd9, 0x7a, 0x31, 0x2e, 0xb6, 0xc4, 0xdc, 0xfc, 0xd2, 0xbc, 0x12, 0xa8, 0x03, 0xa0, 0x3c, 0xb8,
	0xb3, 0x58, 0x90, 0x9c, 0x05, 0x0b, 0x1e, 0x56, 0xb0, 0x8b, 0xf0, 0x04, 0x4f, 0x12, 0x1b, 0x58,
	0xcc, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x45, 0x44, 0xb8, 0xae, 0xf1, 0x01, 0x00, 0x00,
}