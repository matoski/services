// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/m3o/services/users/service/proto/users.proto

package go_micro_srv_users

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

type EventType int32

const (
	EventType_Unknown     EventType = 0
	EventType_UserCreated EventType = 1
	EventType_UserUpdated EventType = 2
	EventType_UserDeleted EventType = 3
)

var EventType_name = map[int32]string{
	0: "Unknown",
	1: "UserCreated",
	2: "UserUpdated",
	3: "UserDeleted",
}

var EventType_value = map[string]int32{
	"Unknown":     0,
	"UserCreated": 1,
	"UserUpdated": 2,
	"UserDeleted": 3,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4d5e149453dfbc06, []int{0}
}

type Event struct {
	Type                 EventType `protobuf:"varint,1,opt,name=type,proto3,enum=go.micro.srv.users.EventType" json:"type,omitempty"`
	User                 *User     `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d5e149453dfbc06, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetType() EventType {
	if m != nil {
		return m.Type
	}
	return EventType_Unknown
}

func (m *Event) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type User struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64             `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64             `protobuf:"varint,3,opt,name=updated,proto3" json:"updated,omitempty"`
	FirstName            string            `protobuf:"bytes,4,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string            `protobuf:"bytes,5,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email                string            `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,8,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ProfilePictureUrl    string            `protobuf:"bytes,9,opt,name=profile_picture_url,json=profilePictureUrl,proto3" json:"profile_picture_url,omitempty"`
	InviteVerified       bool              `protobuf:"varint,10,opt,name=invite_verified,json=inviteVerified,proto3" json:"invite_verified,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d5e149453dfbc06, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *User) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *User) GetProfilePictureUrl() string {
	if m != nil {
		return m.ProfilePictureUrl
	}
	return ""
}

func (m *User) GetInviteVerified() bool {
	if m != nil {
		return m.InviteVerified
	}
	return false
}

type CreateRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	ValidateOnly         bool     `protobuf:"varint,2,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d5e149453dfbc06, []int{2}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CreateRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

type CreateResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d5e149453dfbc06, []int{3}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type DeleteRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d5e149453dfbc06, []int{4}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d5e149453dfbc06, []int{5}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

type ReadRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Ids                  []string `protobuf:"bytes,3,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d5e149453dfbc06, []int{6}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReadRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ReadRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type ReadResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Users                []*User  `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d5e149453dfbc06, []int{7}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *ReadResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type UpdateRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d5e149453dfbc06, []int{8}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d5e149453dfbc06, []int{9}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type SearchRequest struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d5e149453dfbc06, []int{10}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type SearchResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d5e149453dfbc06, []int{11}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterEnum("go.micro.srv.users.EventType", EventType_name, EventType_value)
	proto.RegisterType((*Event)(nil), "go.micro.srv.users.Event")
	proto.RegisterType((*User)(nil), "go.micro.srv.users.User")
	proto.RegisterMapType((map[string]string)(nil), "go.micro.srv.users.User.MetadataEntry")
	proto.RegisterType((*CreateRequest)(nil), "go.micro.srv.users.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "go.micro.srv.users.CreateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "go.micro.srv.users.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "go.micro.srv.users.DeleteResponse")
	proto.RegisterType((*ReadRequest)(nil), "go.micro.srv.users.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "go.micro.srv.users.ReadResponse")
	proto.RegisterType((*UpdateRequest)(nil), "go.micro.srv.users.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "go.micro.srv.users.UpdateResponse")
	proto.RegisterType((*SearchRequest)(nil), "go.micro.srv.users.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "go.micro.srv.users.SearchResponse")
}

func init() {
	proto.RegisterFile("github.com/m3o/services/users/service/proto/users.proto", fileDescriptor_4d5e149453dfbc06)
}

var fileDescriptor_4d5e149453dfbc06 = []byte{
	// 647 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4b, 0x6f, 0xd3, 0x40,
	0x10, 0xae, 0xed, 0x24, 0x4d, 0x26, 0x4d, 0x6a, 0x96, 0x1e, 0xac, 0xa2, 0xaa, 0xc1, 0x08, 0x88,
	0x10, 0x72, 0x44, 0xb8, 0xf0, 0x10, 0x08, 0x01, 0x3d, 0x55, 0x3c, 0x6a, 0x08, 0xd7, 0xc8, 0x8d,
	0xa7, 0xed, 0xaa, 0x8e, 0xed, 0xee, 0xda, 0x46, 0xfe, 0x37, 0xfc, 0x4a, 0xce, 0x68, 0x77, 0xbd,
	0x69, 0x5d, 0x1c, 0x54, 0x7a, 0xf3, 0xcc, 0xf7, 0xcd, 0xfb, 0xdb, 0x04, 0x5e, 0x9e, 0xd2, 0xec,
	0x2c, 0x3f, 0xf6, 0x16, 0xc9, 0x72, 0xb2, 0xa4, 0x0b, 0x96, 0x4c, 0x38, 0xb2, 0x82, 0x2e, 0x90,
	0x4f, 0x72, 0x8e, 0x8c, 0x6b, 0x73, 0x92, 0xb2, 0x24, 0x4b, 0x94, 0xcf, 0x93, 0xdf, 0x84, 0x9c,
	0x26, 0x9e, 0x0c, 0xf1, 0x38, 0x2b, 0x3c, 0x89, 0xb8, 0x67, 0xd0, 0x3e, 0x28, 0x30, 0xce, 0xc8,
	0x33, 0x68, 0x65, 0x65, 0x8a, 0x8e, 0x31, 0x32, 0xc6, 0xc3, 0xe9, 0x9e, 0xf7, 0x37, 0xd7, 0x93,
	0xc4, 0xef, 0x65, 0x8a, 0xbe, 0xa4, 0x92, 0xa7, 0xd0, 0x12, 0x80, 0x63, 0x8e, 0x8c, 0x71, 0x7f,
	0xea, 0x34, 0x85, 0xcc, 0x38, 0x32, 0x5f, 0xb2, 0xdc, 0xdf, 0x26, 0xb4, 0x84, 0x49, 0x86, 0x60,
	0xd2, 0x50, 0xd6, 0xe9, 0xf9, 0x26, 0x0d, 0x89, 0x03, 0x9b, 0x0b, 0x86, 0x41, 0x86, 0xa1, 0xcc,
	0x64, 0xf9, 0xda, 0x14, 0x48, 0x9e, 0x86, 0x12, 0xb1, 0x14, 0x52, 0x99, 0x64, 0x0f, 0xe0, 0x84,
	0x32, 0x9e, 0xcd, 0xe3, 0x60, 0x89, 0x4e, 0x4b, 0xe6, 0xea, 0x49, 0xcf, 0xe7, 0x60, 0x89, 0xe4,
	0x1e, 0xf4, 0xa2, 0x40, 0xa3, 0x6d, 0x89, 0x76, 0x85, 0x43, 0x82, 0x3b, 0xd0, 0xc6, 0x65, 0x40,
	0x23, 0xa7, 0x23, 0x01, 0x65, 0x90, 0xf7, 0xd0, 0x5d, 0x62, 0x16, 0x84, 0x41, 0x16, 0x38, 0xdd,
	0x91, 0x35, 0xee, 0x4f, 0x1f, 0xad, 0x1b, 0xc8, 0xfb, 0x54, 0x11, 0x0f, 0xe2, 0x8c, 0x95, 0xfe,
	0x2a, 0x8e, 0x78, 0x70, 0x37, 0x65, 0xc9, 0x09, 0x8d, 0x70, 0x9e, 0xd2, 0x45, 0x96, 0x33, 0x9c,
	0xe7, 0x2c, 0x72, 0x7a, 0xb2, 0xce, 0x9d, 0x0a, 0xfa, 0xaa, 0x90, 0x19, 0x8b, 0xc8, 0x63, 0xd8,
	0xa6, 0x71, 0x41, 0x33, 0x9c, 0x17, 0xc8, 0xe8, 0x09, 0xc5, 0xd0, 0x81, 0x91, 0x31, 0xee, 0xfa,
	0x43, 0xe5, 0xfe, 0x51, 0x79, 0x77, 0x5f, 0xc3, 0xa0, 0x56, 0x93, 0xd8, 0x60, 0x9d, 0x63, 0x59,
	0x2d, 0x51, 0x7c, 0x8a, 0xa9, 0x8a, 0x20, 0xca, 0x51, 0xee, 0xb0, 0xe7, 0x2b, 0xe3, 0x95, 0xf9,
	0xc2, 0x70, 0x8f, 0x61, 0xf0, 0x41, 0x2e, 0xd4, 0xc7, 0x8b, 0x1c, 0x79, 0xb6, 0xba, 0x9b, 0x71,
	0x93, 0xbb, 0x91, 0x07, 0x30, 0x28, 0x82, 0x88, 0x8a, 0xbd, 0xcf, 0x93, 0x38, 0x2a, 0x65, 0x81,
	0xae, 0xbf, 0xa5, 0x9d, 0x5f, 0xe2, 0xa8, 0x74, 0xdf, 0xc2, 0x50, 0xd7, 0xe0, 0x69, 0x12, 0x73,
	0xfc, 0xbf, 0x22, 0xee, 0x3e, 0x0c, 0x3e, 0x62, 0x84, 0x97, 0x3d, 0x5e, 0x13, 0x89, 0x6b, 0xc3,
	0x50, 0x13, 0x54, 0x01, 0xf7, 0x00, 0xfa, 0x3e, 0x06, 0xe1, 0x9a, 0x80, 0xcb, 0x2b, 0x9b, 0x57,
	0xaf, 0x6c, 0x83, 0x45, 0x43, 0xee, 0x58, 0x23, 0x4b, 0xec, 0x8d, 0x86, 0xdc, 0x8d, 0x60, 0x4b,
	0xa5, 0xb9, 0x4d, 0xdf, 0xc4, 0x83, 0xb6, 0xf4, 0x39, 0xa6, 0x94, 0xcc, 0x7a, 0xba, 0xa2, 0xb9,
	0x6f, 0x60, 0x30, 0x93, 0x12, 0xbe, 0xd5, 0x2d, 0xc4, 0x9a, 0x75, 0xf8, 0xad, 0xd6, 0xfc, 0x10,
	0x06, 0xdf, 0x30, 0x60, 0x8b, 0x33, 0x5d, 0x7e, 0x07, 0xda, 0x17, 0x39, 0x32, 0xad, 0x24, 0x65,
	0xb8, 0xef, 0x60, 0xa8, 0x69, 0x55, 0x99, 0xd5, 0x9c, 0xc6, 0x8d, 0xe6, 0x7c, 0x72, 0x08, 0xbd,
	0xd5, 0xaf, 0x05, 0xe9, 0xc3, 0xe6, 0x2c, 0x3e, 0x8f, 0x93, 0x9f, 0xb1, 0xbd, 0x41, 0xb6, 0xa1,
	0x2f, 0x88, 0x4a, 0x2d, 0xa1, 0x6d, 0x68, 0x87, 0x9a, 0x2b, 0xb4, 0x4d, 0xed, 0x50, 0xe7, 0x0e,
	0x6d, 0x6b, 0xfa, 0xcb, 0x82, 0xb6, 0xf0, 0x70, 0x72, 0x04, 0x1d, 0x15, 0x48, 0xee, 0x37, 0x75,
	0x50, 0x93, 0xf9, 0xae, 0xfb, 0x2f, 0x4a, 0x25, 0xa2, 0x0d, 0x72, 0x08, 0x2d, 0x71, 0x7f, 0xb2,
	0xdf, 0xc4, 0xbe, 0x22, 0xb0, 0xdd, 0xd1, 0x7a, 0xc2, 0x2a, 0xd9, 0x11, 0x74, 0xd4, 0x1c, 0xcd,
	0xfd, 0xd5, 0x4e, 0xdf, 0xdc, 0x5f, 0xfd, 0xbc, 0x2a, 0xa5, 0xda, 0x44, 0x73, 0xca, 0xda, 0xab,
	0x69, 0x4e, 0x79, 0xed, 0xdd, 0xc8, 0x94, 0xea, 0xbc, 0xcd, 0x29, 0x6b, 0x0a, 0x69, 0x4e, 0x59,
	0x57, 0x87, 0xbb, 0x71, 0xdc, 0x91, 0xff, 0x30, 0xcf, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc3,
	0x5e, 0xe4, 0x0a, 0x9e, 0x06, 0x00, 0x00,
}
