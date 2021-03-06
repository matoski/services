// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/m3o/services/kubernetes/service/proto/kubernetes.proto

package go_micro_service_kubernetes

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

type CreateNamespaceRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateNamespaceRequest) Reset()         { *m = CreateNamespaceRequest{} }
func (m *CreateNamespaceRequest) String() string { return proto.CompactTextString(m) }
func (*CreateNamespaceRequest) ProtoMessage()    {}
func (*CreateNamespaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_09bf30676138af59, []int{0}
}

func (m *CreateNamespaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateNamespaceRequest.Unmarshal(m, b)
}
func (m *CreateNamespaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateNamespaceRequest.Marshal(b, m, deterministic)
}
func (m *CreateNamespaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNamespaceRequest.Merge(m, src)
}
func (m *CreateNamespaceRequest) XXX_Size() int {
	return xxx_messageInfo_CreateNamespaceRequest.Size(m)
}
func (m *CreateNamespaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNamespaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNamespaceRequest proto.InternalMessageInfo

func (m *CreateNamespaceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateNamespaceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateNamespaceResponse) Reset()         { *m = CreateNamespaceResponse{} }
func (m *CreateNamespaceResponse) String() string { return proto.CompactTextString(m) }
func (*CreateNamespaceResponse) ProtoMessage()    {}
func (*CreateNamespaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_09bf30676138af59, []int{1}
}

func (m *CreateNamespaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateNamespaceResponse.Unmarshal(m, b)
}
func (m *CreateNamespaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateNamespaceResponse.Marshal(b, m, deterministic)
}
func (m *CreateNamespaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNamespaceResponse.Merge(m, src)
}
func (m *CreateNamespaceResponse) XXX_Size() int {
	return xxx_messageInfo_CreateNamespaceResponse.Size(m)
}
func (m *CreateNamespaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNamespaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNamespaceResponse proto.InternalMessageInfo

type DeleteNamespaceRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteNamespaceRequest) Reset()         { *m = DeleteNamespaceRequest{} }
func (m *DeleteNamespaceRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteNamespaceRequest) ProtoMessage()    {}
func (*DeleteNamespaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_09bf30676138af59, []int{2}
}

func (m *DeleteNamespaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteNamespaceRequest.Unmarshal(m, b)
}
func (m *DeleteNamespaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteNamespaceRequest.Marshal(b, m, deterministic)
}
func (m *DeleteNamespaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteNamespaceRequest.Merge(m, src)
}
func (m *DeleteNamespaceRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteNamespaceRequest.Size(m)
}
func (m *DeleteNamespaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteNamespaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteNamespaceRequest proto.InternalMessageInfo

func (m *DeleteNamespaceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeleteNamespaceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteNamespaceResponse) Reset()         { *m = DeleteNamespaceResponse{} }
func (m *DeleteNamespaceResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteNamespaceResponse) ProtoMessage()    {}
func (*DeleteNamespaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_09bf30676138af59, []int{3}
}

func (m *DeleteNamespaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteNamespaceResponse.Unmarshal(m, b)
}
func (m *DeleteNamespaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteNamespaceResponse.Marshal(b, m, deterministic)
}
func (m *DeleteNamespaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteNamespaceResponse.Merge(m, src)
}
func (m *DeleteNamespaceResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteNamespaceResponse.Size(m)
}
func (m *DeleteNamespaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteNamespaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteNamespaceResponse proto.InternalMessageInfo

type CreateImagePullSecretRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateImagePullSecretRequest) Reset()         { *m = CreateImagePullSecretRequest{} }
func (m *CreateImagePullSecretRequest) String() string { return proto.CompactTextString(m) }
func (*CreateImagePullSecretRequest) ProtoMessage()    {}
func (*CreateImagePullSecretRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_09bf30676138af59, []int{4}
}

func (m *CreateImagePullSecretRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateImagePullSecretRequest.Unmarshal(m, b)
}
func (m *CreateImagePullSecretRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateImagePullSecretRequest.Marshal(b, m, deterministic)
}
func (m *CreateImagePullSecretRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateImagePullSecretRequest.Merge(m, src)
}
func (m *CreateImagePullSecretRequest) XXX_Size() int {
	return xxx_messageInfo_CreateImagePullSecretRequest.Size(m)
}
func (m *CreateImagePullSecretRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateImagePullSecretRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateImagePullSecretRequest proto.InternalMessageInfo

func (m *CreateImagePullSecretRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateImagePullSecretRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *CreateImagePullSecretRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CreateImagePullSecretResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateImagePullSecretResponse) Reset()         { *m = CreateImagePullSecretResponse{} }
func (m *CreateImagePullSecretResponse) String() string { return proto.CompactTextString(m) }
func (*CreateImagePullSecretResponse) ProtoMessage()    {}
func (*CreateImagePullSecretResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_09bf30676138af59, []int{5}
}

func (m *CreateImagePullSecretResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateImagePullSecretResponse.Unmarshal(m, b)
}
func (m *CreateImagePullSecretResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateImagePullSecretResponse.Marshal(b, m, deterministic)
}
func (m *CreateImagePullSecretResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateImagePullSecretResponse.Merge(m, src)
}
func (m *CreateImagePullSecretResponse) XXX_Size() int {
	return xxx_messageInfo_CreateImagePullSecretResponse.Size(m)
}
func (m *CreateImagePullSecretResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateImagePullSecretResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateImagePullSecretResponse proto.InternalMessageInfo

type DeleteImagePullSecretRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteImagePullSecretRequest) Reset()         { *m = DeleteImagePullSecretRequest{} }
func (m *DeleteImagePullSecretRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteImagePullSecretRequest) ProtoMessage()    {}
func (*DeleteImagePullSecretRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_09bf30676138af59, []int{6}
}

func (m *DeleteImagePullSecretRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteImagePullSecretRequest.Unmarshal(m, b)
}
func (m *DeleteImagePullSecretRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteImagePullSecretRequest.Marshal(b, m, deterministic)
}
func (m *DeleteImagePullSecretRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteImagePullSecretRequest.Merge(m, src)
}
func (m *DeleteImagePullSecretRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteImagePullSecretRequest.Size(m)
}
func (m *DeleteImagePullSecretRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteImagePullSecretRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteImagePullSecretRequest proto.InternalMessageInfo

func (m *DeleteImagePullSecretRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeleteImagePullSecretRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type DeleteImagePullSecretResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteImagePullSecretResponse) Reset()         { *m = DeleteImagePullSecretResponse{} }
func (m *DeleteImagePullSecretResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteImagePullSecretResponse) ProtoMessage()    {}
func (*DeleteImagePullSecretResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_09bf30676138af59, []int{7}
}

func (m *DeleteImagePullSecretResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteImagePullSecretResponse.Unmarshal(m, b)
}
func (m *DeleteImagePullSecretResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteImagePullSecretResponse.Marshal(b, m, deterministic)
}
func (m *DeleteImagePullSecretResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteImagePullSecretResponse.Merge(m, src)
}
func (m *DeleteImagePullSecretResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteImagePullSecretResponse.Size(m)
}
func (m *DeleteImagePullSecretResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteImagePullSecretResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteImagePullSecretResponse proto.InternalMessageInfo

type CreateServiceAccountRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ImagePullSecrets     []string `protobuf:"bytes,2,rep,name=image_pull_secrets,json=imagePullSecrets,proto3" json:"image_pull_secrets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateServiceAccountRequest) Reset()         { *m = CreateServiceAccountRequest{} }
func (m *CreateServiceAccountRequest) String() string { return proto.CompactTextString(m) }
func (*CreateServiceAccountRequest) ProtoMessage()    {}
func (*CreateServiceAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_09bf30676138af59, []int{8}
}

func (m *CreateServiceAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServiceAccountRequest.Unmarshal(m, b)
}
func (m *CreateServiceAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServiceAccountRequest.Marshal(b, m, deterministic)
}
func (m *CreateServiceAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServiceAccountRequest.Merge(m, src)
}
func (m *CreateServiceAccountRequest) XXX_Size() int {
	return xxx_messageInfo_CreateServiceAccountRequest.Size(m)
}
func (m *CreateServiceAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServiceAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServiceAccountRequest proto.InternalMessageInfo

func (m *CreateServiceAccountRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *CreateServiceAccountRequest) GetImagePullSecrets() []string {
	if m != nil {
		return m.ImagePullSecrets
	}
	return nil
}

type CreateServiceAccountResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateServiceAccountResponse) Reset()         { *m = CreateServiceAccountResponse{} }
func (m *CreateServiceAccountResponse) String() string { return proto.CompactTextString(m) }
func (*CreateServiceAccountResponse) ProtoMessage()    {}
func (*CreateServiceAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_09bf30676138af59, []int{9}
}

func (m *CreateServiceAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServiceAccountResponse.Unmarshal(m, b)
}
func (m *CreateServiceAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServiceAccountResponse.Marshal(b, m, deterministic)
}
func (m *CreateServiceAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServiceAccountResponse.Merge(m, src)
}
func (m *CreateServiceAccountResponse) XXX_Size() int {
	return xxx_messageInfo_CreateServiceAccountResponse.Size(m)
}
func (m *CreateServiceAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServiceAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServiceAccountResponse proto.InternalMessageInfo

type DeleteServiceAccountRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteServiceAccountRequest) Reset()         { *m = DeleteServiceAccountRequest{} }
func (m *DeleteServiceAccountRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteServiceAccountRequest) ProtoMessage()    {}
func (*DeleteServiceAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_09bf30676138af59, []int{10}
}

func (m *DeleteServiceAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteServiceAccountRequest.Unmarshal(m, b)
}
func (m *DeleteServiceAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteServiceAccountRequest.Marshal(b, m, deterministic)
}
func (m *DeleteServiceAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteServiceAccountRequest.Merge(m, src)
}
func (m *DeleteServiceAccountRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteServiceAccountRequest.Size(m)
}
func (m *DeleteServiceAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteServiceAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteServiceAccountRequest proto.InternalMessageInfo

func (m *DeleteServiceAccountRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type DeleteServiceAccountResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteServiceAccountResponse) Reset()         { *m = DeleteServiceAccountResponse{} }
func (m *DeleteServiceAccountResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteServiceAccountResponse) ProtoMessage()    {}
func (*DeleteServiceAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_09bf30676138af59, []int{11}
}

func (m *DeleteServiceAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteServiceAccountResponse.Unmarshal(m, b)
}
func (m *DeleteServiceAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteServiceAccountResponse.Marshal(b, m, deterministic)
}
func (m *DeleteServiceAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteServiceAccountResponse.Merge(m, src)
}
func (m *DeleteServiceAccountResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteServiceAccountResponse.Size(m)
}
func (m *DeleteServiceAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteServiceAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteServiceAccountResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateNamespaceRequest)(nil), "go.micro.service.kubernetes.CreateNamespaceRequest")
	proto.RegisterType((*CreateNamespaceResponse)(nil), "go.micro.service.kubernetes.CreateNamespaceResponse")
	proto.RegisterType((*DeleteNamespaceRequest)(nil), "go.micro.service.kubernetes.DeleteNamespaceRequest")
	proto.RegisterType((*DeleteNamespaceResponse)(nil), "go.micro.service.kubernetes.DeleteNamespaceResponse")
	proto.RegisterType((*CreateImagePullSecretRequest)(nil), "go.micro.service.kubernetes.CreateImagePullSecretRequest")
	proto.RegisterType((*CreateImagePullSecretResponse)(nil), "go.micro.service.kubernetes.CreateImagePullSecretResponse")
	proto.RegisterType((*DeleteImagePullSecretRequest)(nil), "go.micro.service.kubernetes.DeleteImagePullSecretRequest")
	proto.RegisterType((*DeleteImagePullSecretResponse)(nil), "go.micro.service.kubernetes.DeleteImagePullSecretResponse")
	proto.RegisterType((*CreateServiceAccountRequest)(nil), "go.micro.service.kubernetes.CreateServiceAccountRequest")
	proto.RegisterType((*CreateServiceAccountResponse)(nil), "go.micro.service.kubernetes.CreateServiceAccountResponse")
	proto.RegisterType((*DeleteServiceAccountRequest)(nil), "go.micro.service.kubernetes.DeleteServiceAccountRequest")
	proto.RegisterType((*DeleteServiceAccountResponse)(nil), "go.micro.service.kubernetes.DeleteServiceAccountResponse")
}

func init() {
	proto.RegisterFile("github.com/m3o/services/kubernetes/service/proto/kubernetes.proto", fileDescriptor_09bf30676138af59)
}

var fileDescriptor_09bf30676138af59 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xc1, 0x4e, 0xfa, 0x40,
	0x10, 0xc6, 0x29, 0xfc, 0xff, 0x24, 0xcc, 0x45, 0xb3, 0x41, 0xad, 0x05, 0x95, 0xf4, 0xc4, 0x81,
	0xb4, 0x89, 0x78, 0x10, 0x3d, 0x19, 0xbc, 0x18, 0x13, 0x43, 0xe0, 0x01, 0x48, 0x69, 0x46, 0x6c,
	0x68, 0xbb, 0xb5, 0xdb, 0x7a, 0xf4, 0x0d, 0x4c, 0x7c, 0x36, 0x9f, 0xc8, 0xb4, 0xdb, 0x82, 0x2d,
	0x4b, 0x8b, 0x8d, 0x37, 0xd8, 0x9d, 0xf9, 0xbe, 0x6f, 0xda, 0xdf, 0x14, 0xc6, 0x4b, 0x2b, 0x78,
	0x09, 0x17, 0x9a, 0x49, 0x1d, 0xdd, 0xb1, 0x4c, 0x9f, 0xea, 0x0c, 0xfd, 0x37, 0xcb, 0x44, 0xa6,
	0xaf, 0xc2, 0x05, 0xfa, 0x2e, 0x06, 0xc8, 0xd2, 0x33, 0xdd, 0xf3, 0x69, 0x40, 0x7f, 0x5c, 0x68,
	0xf1, 0x01, 0xe9, 0x2c, 0xa9, 0x16, 0x37, 0x6b, 0x49, 0xa1, 0xb6, 0x29, 0x51, 0x07, 0x70, 0x3c,
	0xf6, 0xd1, 0x08, 0xf0, 0xc9, 0x70, 0x90, 0x79, 0x86, 0x89, 0x53, 0x7c, 0x0d, 0x91, 0x05, 0x84,
	0xc0, 0x3f, 0xd7, 0x70, 0x50, 0x96, 0x7a, 0x52, 0xbf, 0x35, 0x8d, 0x7f, 0xab, 0xa7, 0x70, 0xb2,
	0x55, 0xcd, 0x3c, 0xea, 0x32, 0x8c, 0x84, 0xee, 0xd1, 0xc6, 0xfd, 0x85, 0xb6, 0xaa, 0x13, 0xa1,
	0x67, 0xe8, 0x72, 0x8f, 0x07, 0xc7, 0x58, 0xe2, 0x24, 0xb4, 0xed, 0x19, 0x9a, 0x3e, 0x06, 0x05,
	0x72, 0xa4, 0x0b, 0x2d, 0x37, 0x15, 0x92, 0xeb, 0xf1, 0xc5, 0xe6, 0x80, 0xb4, 0xe1, 0x7f, 0x40,
	0x57, 0xe8, 0xca, 0x8d, 0xf8, 0x86, 0xff, 0x51, 0x2f, 0xe0, 0x6c, 0x87, 0x4f, 0x12, 0x64, 0x02,
	0x5d, 0x9e, 0xf1, 0xaf, 0x82, 0x44, 0x96, 0x3b, 0x14, 0x13, 0x4b, 0x0b, 0x3a, 0x3c, 0xd3, 0x8c,
	0xbf, 0xa9, 0x3b, 0xd3, 0xa4, 0xa1, 0xbb, 0x76, 0xcc, 0xa8, 0x4b, 0xf9, 0x31, 0x07, 0x40, 0xac,
	0x48, 0x77, 0xee, 0x85, 0xb6, 0x3d, 0x67, 0xb1, 0x32, 0x93, 0xeb, 0xbd, 0x46, 0xbf, 0x35, 0x3d,
	0xb4, 0xb2, 0x8e, 0x4c, 0x3d, 0x4f, 0x1f, 0x73, 0xde, 0x2a, 0x89, 0x72, 0x0b, 0x1d, 0x9e, 0xb5,
	0x42, 0x94, 0x48, 0x5c, 0xdc, 0xcc, 0xc5, 0x2f, 0xbf, 0x9a, 0x00, 0x8f, 0x6b, 0x08, 0xc9, 0x3b,
	0x1c, 0xe4, 0xb0, 0x22, 0x43, 0xad, 0x80, 0x5a, 0x4d, 0x8c, 0xac, 0x72, 0xf5, 0xbb, 0xa6, 0x64,
	0xd2, 0x5a, 0xe4, 0x9f, 0xa3, 0xb1, 0xc4, 0x5f, 0x4c, 0x7a, 0x89, 0xff, 0x2e, 0xe0, 0x6b, 0xe4,
	0x53, 0x82, 0x23, 0x21, 0x8b, 0x64, 0xb4, 0xc7, 0x44, 0x62, 0x3c, 0x95, 0x9b, 0x2a, 0xad, 0x99,
	0x48, 0x42, 0x56, 0x4b, 0x22, 0x15, 0x6d, 0x4c, 0x49, 0xa4, 0xe2, 0xd5, 0xa8, 0x91, 0x0f, 0x09,
	0xda, 0x22, 0x64, 0xc9, 0xf5, 0x1e, 0x93, 0x0a, 0x29, 0x56, 0x46, 0x15, 0x3a, 0x33, 0x79, 0x44,
	0x94, 0x97, 0xe4, 0x29, 0xd8, 0x2a, 0x65, 0x54, 0xa1, 0x33, 0xcd, 0xb3, 0x68, 0xc6, 0x9f, 0xfb,
	0xe1, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x14, 0x60, 0x69, 0xd9, 0x35, 0x06, 0x00, 0x00,
}
