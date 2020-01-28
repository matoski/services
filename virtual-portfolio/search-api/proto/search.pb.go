// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/search.proto

package search

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

type Query struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_47e70f09e802673d, []int{0}
}

func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *Query) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *Query) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type Response struct {
	Results              []*Result `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_47e70f09e802673d, []int{1}
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

func (m *Response) GetResults() []*Result {
	if m != nil {
		return m.Results
	}
	return nil
}

type Result struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Uuid                 string   `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Subtitle             string   `protobuf:"bytes,4,opt,name=subtitle,proto3" json:"subtitle,omitempty"`
	ProfilePictureUrl    string   `protobuf:"bytes,5,opt,name=profile_picture_url,json=profilePictureUrl,proto3" json:"profile_picture_url,omitempty"`
	Color                string   `protobuf:"bytes,6,opt,name=color,proto3" json:"color,omitempty"`
	Description          string   `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Following            bool     `protobuf:"varint,8,opt,name=following,proto3" json:"following,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_47e70f09e802673d, []int{2}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Result) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Result) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Result) GetSubtitle() string {
	if m != nil {
		return m.Subtitle
	}
	return ""
}

func (m *Result) GetProfilePictureUrl() string {
	if m != nil {
		return m.ProfilePictureUrl
	}
	return ""
}

func (m *Result) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *Result) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Result) GetFollowing() bool {
	if m != nil {
		return m.Following
	}
	return false
}

func init() {
	proto.RegisterType((*Query)(nil), "Query")
	proto.RegisterType((*Response)(nil), "Response")
	proto.RegisterType((*Result)(nil), "Result")
}

func init() { proto.RegisterFile("proto/search.proto", fileDescriptor_47e70f09e802673d) }

var fileDescriptor_47e70f09e802673d = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xcf, 0x4e, 0xf3, 0x30,
	0x10, 0xc4, 0x3f, 0x7f, 0x6d, 0xfe, 0x6d, 0x4f, 0x18, 0x0e, 0x56, 0x05, 0x52, 0xc8, 0x85, 0x5c,
	0x08, 0x52, 0x79, 0x08, 0xae, 0x60, 0xc4, 0xb9, 0xa2, 0xa9, 0x0b, 0x96, 0x4c, 0x6c, 0xd6, 0xb6,
	0x50, 0xdf, 0x98, 0xc7, 0x40, 0x59, 0x93, 0x94, 0xdb, 0xfc, 0x66, 0x56, 0x3b, 0xd6, 0x1a, 0xb8,
	0x43, 0x1b, 0xec, 0x9d, 0x57, 0xaf, 0xd8, 0xbf, 0x77, 0x04, 0xcd, 0x03, 0x64, 0x4f, 0x51, 0xe1,
	0x91, 0x5f, 0x40, 0xf6, 0x39, 0x0a, 0xc1, 0x6a, 0xd6, 0x56, 0x32, 0xc1, 0xe8, 0x1a, 0xfd, 0xa1,
	0x83, 0xf8, 0x5f, 0xb3, 0x36, 0x93, 0x09, 0x38, 0x87, 0x65, 0x38, 0x3a, 0x25, 0x16, 0x34, 0x4a,
	0xba, 0xb9, 0x85, 0x52, 0x2a, 0xef, 0xec, 0xe0, 0x15, 0xbf, 0x86, 0x02, 0x95, 0x8f, 0x26, 0x78,
	0xc1, 0xea, 0x45, 0xbb, 0xda, 0x14, 0x9d, 0x24, 0x96, 0x93, 0xdf, 0x7c, 0x33, 0xc8, 0x93, 0x37,
	0x6f, 0x63, 0xa7, 0x6d, 0xa3, 0x17, 0xa3, 0xde, 0x53, 0x6d, 0x25, 0x49, 0x8f, 0x6f, 0x09, 0x3a,
	0x98, 0xa9, 0x36, 0x01, 0x5f, 0x43, 0xe9, 0xe3, 0x2e, 0x05, 0x4b, 0x0a, 0x66, 0xe6, 0x1d, 0x9c,
	0x3b, 0xb4, 0x07, 0x6d, 0xd4, 0xd6, 0xe9, 0x3e, 0x44, 0x54, 0xdb, 0x88, 0x46, 0x64, 0x34, 0x76,
	0xf6, 0x1b, 0x3d, 0xa6, 0xe4, 0x05, 0xcd, 0xd8, 0xd0, 0x5b, 0x63, 0x51, 0xe4, 0xa9, 0x81, 0x80,
	0xd7, 0xb0, 0xda, 0x2b, 0xdf, 0xa3, 0x76, 0x41, 0xdb, 0x41, 0x14, 0x94, 0xfd, 0xb5, 0xf8, 0x25,
	0x54, 0x07, 0x6b, 0x8c, 0xfd, 0xd2, 0xc3, 0x9b, 0x28, 0x6b, 0xd6, 0x96, 0xf2, 0x64, 0x6c, 0x6e,
	0x20, 0x7f, 0xa6, 0x93, 0xf3, 0xab, 0x59, 0xe5, 0x1d, 0x5d, 0x7d, 0x5d, 0x75, 0xd3, 0xd1, 0x9a,
	0x7f, 0xbb, 0x9c, 0xbe, 0xe4, 0xfe, 0x27, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x35, 0x8e, 0xf5, 0xa8,
	0x01, 0x00, 0x00,
}