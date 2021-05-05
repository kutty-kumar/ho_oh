// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ho-oh/snorlax_v1/snorlax.proto

package snorlax_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	core_v1 "github.com/kutty-kumar/ho_oh/core_v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type UserDto struct {
	FirstName            string         `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string         `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	DateOfBirth          string         `protobuf:"bytes,3,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	Gender               core_v1.Gender `protobuf:"varint,4,opt,name=gender,proto3,enum=core.Gender" json:"gender,omitempty"`
	Status               core_v1.Status `protobuf:"varint,5,opt,name=status,proto3,enum=core.Status" json:"status,omitempty"`
	UserId               string         `protobuf:"bytes,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Email                string         `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	Password             string         `protobuf:"bytes,8,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UserDto) Reset()         { *m = UserDto{} }
func (m *UserDto) String() string { return proto.CompactTextString(m) }
func (*UserDto) ProtoMessage()    {}
func (*UserDto) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc85baa20be4828d, []int{0}
}

func (m *UserDto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDto.Unmarshal(m, b)
}
func (m *UserDto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDto.Marshal(b, m, deterministic)
}
func (m *UserDto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDto.Merge(m, src)
}
func (m *UserDto) XXX_Size() int {
	return xxx_messageInfo_UserDto.Size(m)
}
func (m *UserDto) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDto.DiscardUnknown(m)
}

var xxx_messageInfo_UserDto proto.InternalMessageInfo

func (m *UserDto) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UserDto) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *UserDto) GetDateOfBirth() string {
	if m != nil {
		return m.DateOfBirth
	}
	return ""
}

func (m *UserDto) GetGender() core_v1.Gender {
	if m != nil {
		return m.Gender
	}
	return core_v1.Gender_unknown_gender
}

func (m *UserDto) GetStatus() core_v1.Status {
	if m != nil {
		return m.Status
	}
	return core_v1.Status_unknown_status
}

func (m *UserDto) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserDto) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserDto) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UserOperationResponse struct {
	Response             *UserDto `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserOperationResponse) Reset()         { *m = UserOperationResponse{} }
func (m *UserOperationResponse) String() string { return proto.CompactTextString(m) }
func (*UserOperationResponse) ProtoMessage()    {}
func (*UserOperationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc85baa20be4828d, []int{1}
}

func (m *UserOperationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserOperationResponse.Unmarshal(m, b)
}
func (m *UserOperationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserOperationResponse.Marshal(b, m, deterministic)
}
func (m *UserOperationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserOperationResponse.Merge(m, src)
}
func (m *UserOperationResponse) XXX_Size() int {
	return xxx_messageInfo_UserOperationResponse.Size(m)
}
func (m *UserOperationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserOperationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserOperationResponse proto.InternalMessageInfo

func (m *UserOperationResponse) GetResponse() *UserDto {
	if m != nil {
		return m.Response
	}
	return nil
}

type CreateUserRequest struct {
	Payload              *UserDto `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc85baa20be4828d, []int{2}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetPayload() *UserDto {
	if m != nil {
		return m.Payload
	}
	return nil
}

type UpdateUserRequest struct {
	Payload              *UserDto `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc85baa20be4828d, []int{3}
}

func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRequest.Unmarshal(m, b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(m, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRequest.Size(m)
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetPayload() *UserDto {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *UpdateUserRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetUserByEmailAndPasswordRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserByEmailAndPasswordRequest) Reset()         { *m = GetUserByEmailAndPasswordRequest{} }
func (m *GetUserByEmailAndPasswordRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserByEmailAndPasswordRequest) ProtoMessage()    {}
func (*GetUserByEmailAndPasswordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc85baa20be4828d, []int{4}
}

func (m *GetUserByEmailAndPasswordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserByEmailAndPasswordRequest.Unmarshal(m, b)
}
func (m *GetUserByEmailAndPasswordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserByEmailAndPasswordRequest.Marshal(b, m, deterministic)
}
func (m *GetUserByEmailAndPasswordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserByEmailAndPasswordRequest.Merge(m, src)
}
func (m *GetUserByEmailAndPasswordRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserByEmailAndPasswordRequest.Size(m)
}
func (m *GetUserByEmailAndPasswordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserByEmailAndPasswordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserByEmailAndPasswordRequest proto.InternalMessageInfo

func (m *GetUserByEmailAndPasswordRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *GetUserByEmailAndPasswordRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type GetUserByExternalIdRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserByExternalIdRequest) Reset()         { *m = GetUserByExternalIdRequest{} }
func (m *GetUserByExternalIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserByExternalIdRequest) ProtoMessage()    {}
func (*GetUserByExternalIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc85baa20be4828d, []int{5}
}

func (m *GetUserByExternalIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserByExternalIdRequest.Unmarshal(m, b)
}
func (m *GetUserByExternalIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserByExternalIdRequest.Marshal(b, m, deterministic)
}
func (m *GetUserByExternalIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserByExternalIdRequest.Merge(m, src)
}
func (m *GetUserByExternalIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserByExternalIdRequest.Size(m)
}
func (m *GetUserByExternalIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserByExternalIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserByExternalIdRequest proto.InternalMessageInfo

func (m *GetUserByExternalIdRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func init() {
	proto.RegisterType((*UserDto)(nil), "snorlax_v1.UserDto")
	proto.RegisterType((*UserOperationResponse)(nil), "snorlax_v1.UserOperationResponse")
	proto.RegisterType((*CreateUserRequest)(nil), "snorlax_v1.CreateUserRequest")
	proto.RegisterType((*UpdateUserRequest)(nil), "snorlax_v1.UpdateUserRequest")
	proto.RegisterType((*GetUserByEmailAndPasswordRequest)(nil), "snorlax_v1.GetUserByEmailAndPasswordRequest")
	proto.RegisterType((*GetUserByExternalIdRequest)(nil), "snorlax_v1.GetUserByExternalIdRequest")
}

func init() { proto.RegisterFile("ho-oh/snorlax_v1/snorlax.proto", fileDescriptor_fc85baa20be4828d) }

var fileDescriptor_fc85baa20be4828d = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x95, 0x8c, 0xf5, 0xcf, 0x29, 0x4c, 0x9a, 0xbb, 0x69, 0x21, 0xa3, 0x53, 0x89, 0x10,
	0x9a, 0x26, 0xd6, 0x68, 0x45, 0xdc, 0xec, 0x8e, 0x02, 0x1a, 0xbb, 0x61, 0x28, 0x63, 0x57, 0x5c,
	0x44, 0xde, 0x72, 0xda, 0x46, 0x4a, 0xed, 0x60, 0xbb, 0x65, 0x13, 0x9a, 0x84, 0xb8, 0xe5, 0x92,
	0x47, 0xe3, 0x15, 0x78, 0x00, 0x1e, 0x01, 0xd9, 0x4d, 0x9a, 0xb0, 0xb6, 0xa8, 0x12, 0x57, 0xb1,
	0xbf, 0x73, 0x72, 0xbe, 0xcf, 0x3f, 0x47, 0x81, 0xbd, 0x21, 0x3f, 0xe4, 0x43, 0x5f, 0x32, 0x2e,
	0x12, 0x7a, 0x1d, 0x4e, 0x8e, 0xf2, 0x65, 0x27, 0x15, 0x5c, 0x71, 0x02, 0x45, 0xc5, 0x7d, 0x34,
	0xe0, 0x7c, 0x90, 0xa0, 0x4f, 0xd3, 0xd8, 0xa7, 0x8c, 0x71, 0x45, 0x55, 0xcc, 0x99, 0x9c, 0x76,
	0xba, 0xe4, 0x8a, 0x0b, 0xd4, 0x03, 0xf4, 0x73, 0xaa, 0x79, 0x5f, 0x6d, 0xa8, 0x5e, 0x48, 0x14,
	0xaf, 0x15, 0x27, 0x2d, 0x80, 0x7e, 0x2c, 0xa4, 0x0a, 0x19, 0x1d, 0xa1, 0x63, 0xb5, 0xad, 0xfd,
	0x7a, 0x50, 0x37, 0xca, 0x3b, 0x3a, 0x42, 0xb2, 0x0b, 0xf5, 0x84, 0xe6, 0x55, 0xdb, 0x54, 0x6b,
	0x5a, 0x30, 0x45, 0x0f, 0x1e, 0x44, 0x54, 0x61, 0xc8, 0xfb, 0xe1, 0x65, 0x2c, 0xd4, 0xd0, 0x59,
	0x33, 0x0d, 0x0d, 0x2d, 0x9e, 0xf5, 0x7b, 0x5a, 0x22, 0x4f, 0xa0, 0x32, 0x40, 0x16, 0xa1, 0x70,
	0xee, 0xb5, 0xad, 0xfd, 0x8d, 0xee, 0xfd, 0x8e, 0x09, 0x72, 0x62, 0xb4, 0x20, 0xab, 0xe9, 0x2e,
	0xa9, 0xa8, 0x1a, 0x4b, 0x67, 0xbd, 0xdc, 0x75, 0x6e, 0xb4, 0x20, 0xab, 0x91, 0x1d, 0xa8, 0x8e,
	0x25, 0x8a, 0x30, 0x8e, 0x9c, 0x8a, 0x71, 0xaa, 0xe8, 0xed, 0x69, 0x44, 0xb6, 0x60, 0x1d, 0x47,
	0x34, 0x4e, 0x9c, 0xaa, 0x91, 0xa7, 0x1b, 0xe2, 0x42, 0x2d, 0xa5, 0x52, 0x7e, 0xe6, 0x22, 0x72,
	0x6a, 0xd3, 0xe8, 0xf9, 0xde, 0x7b, 0x0b, 0xdb, 0x9a, 0xc0, 0x59, 0x8a, 0xc2, 0xe0, 0x0a, 0x50,
	0xa6, 0x9c, 0x49, 0x24, 0x3e, 0xd4, 0x44, 0xb6, 0x36, 0x34, 0x1a, 0xdd, 0x66, 0xa7, 0x80, 0xdd,
	0xc9, 0xb0, 0x05, 0xb3, 0x26, 0xaf, 0x07, 0x9b, 0xaf, 0x04, 0x52, 0x85, 0xba, 0x14, 0xe0, 0xa7,
	0x31, 0x4a, 0x45, 0x0e, 0xa1, 0x9a, 0xd2, 0x9b, 0x84, 0xd3, 0xe8, 0x5f, 0x43, 0xf2, 0x1e, 0xef,
	0x23, 0x6c, 0x5e, 0xa4, 0xd1, 0x7f, 0xcd, 0x28, 0xc3, 0xb1, 0xcb, 0x70, 0xbc, 0x0f, 0xd0, 0x3e,
	0x41, 0xa5, 0xfb, 0x7b, 0x37, 0x6f, 0x34, 0x98, 0x97, 0x2c, 0x7a, 0x9f, 0x71, 0xc8, 0xbd, 0x66,
	0x00, 0xad, 0x65, 0x00, 0xed, 0x3b, 0x00, 0x5f, 0x80, 0x5b, 0x4c, 0xbd, 0x56, 0x28, 0x18, 0x4d,
	0x4e, 0x67, 0xf3, 0x4a, 0x61, 0xac, 0x72, 0x98, 0xee, 0xef, 0x35, 0x68, 0xe8, 0x97, 0xce, 0x51,
	0x4c, 0xe2, 0x2b, 0x24, 0x11, 0x40, 0x41, 0x8f, 0xb4, 0xca, 0x27, 0x9c, 0xa3, 0xea, 0x3e, 0xbe,
	0x0b, 0x60, 0xee, 0xfa, 0xbc, 0xad, 0x6f, 0x3f, 0x7f, 0xfd, 0xb0, 0x37, 0xbc, 0xba, 0x3f, 0x39,
	0xf2, 0xb5, 0xa7, 0x3c, 0xb6, 0x0e, 0x08, 0x03, 0x28, 0xf8, 0xfe, 0xed, 0x32, 0xc7, 0x7d, 0x15,
	0x97, 0x3d, 0xe3, 0xe2, 0x74, 0x9b, 0x33, 0x17, 0xff, 0x4b, 0x76, 0xde, 0x5b, 0xed, 0x77, 0x0b,
	0xcd, 0x05, 0x70, 0xc8, 0xd3, 0xf2, 0xe4, 0xe5, 0xf4, 0x56, 0x49, 0xb0, 0x6b, 0x12, 0x6c, 0x93,
	0x45, 0x09, 0xc8, 0x77, 0x0b, 0x1e, 0x2e, 0xbd, 0x72, 0xf2, 0x6c, 0x71, 0x8a, 0xc5, 0x5f, 0xc6,
	0x2a, 0x59, 0x5a, 0x26, 0xcb, 0x8e, 0x4b, 0x8a, 0x2c, 0x13, 0x9a, 0xc4, 0x9a, 0xeb, 0xb1, 0x75,
	0x70, 0x59, 0x31, 0x3f, 0x9d, 0xe7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xc9, 0x76, 0xc2,
	0xd4, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserOperationResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserOperationResponse, error)
	GetUserByExternalId(ctx context.Context, in *GetUserByExternalIdRequest, opts ...grpc.CallOption) (*UserOperationResponse, error)
	GetUserByEmailAndPassword(ctx context.Context, in *GetUserByEmailAndPasswordRequest, opts ...grpc.CallOption) (*UserOperationResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserOperationResponse, error) {
	out := new(UserOperationResponse)
	err := c.cc.Invoke(ctx, "/snorlax_v1.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserOperationResponse, error) {
	out := new(UserOperationResponse)
	err := c.cc.Invoke(ctx, "/snorlax_v1.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByExternalId(ctx context.Context, in *GetUserByExternalIdRequest, opts ...grpc.CallOption) (*UserOperationResponse, error) {
	out := new(UserOperationResponse)
	err := c.cc.Invoke(ctx, "/snorlax_v1.UserService/GetUserByExternalId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByEmailAndPassword(ctx context.Context, in *GetUserByEmailAndPasswordRequest, opts ...grpc.CallOption) (*UserOperationResponse, error) {
	out := new(UserOperationResponse)
	err := c.cc.Invoke(ctx, "/snorlax_v1.UserService/GetUserByEmailAndPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*UserOperationResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UserOperationResponse, error)
	GetUserByExternalId(context.Context, *GetUserByExternalIdRequest) (*UserOperationResponse, error)
	GetUserByEmailAndPassword(context.Context, *GetUserByEmailAndPasswordRequest) (*UserOperationResponse, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snorlax_v1.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snorlax_v1.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByExternalId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByExternalIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByExternalId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snorlax_v1.UserService/GetUserByExternalId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByExternalId(ctx, req.(*GetUserByExternalIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByEmailAndPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByEmailAndPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByEmailAndPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snorlax_v1.UserService/GetUserByEmailAndPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByEmailAndPassword(ctx, req.(*GetUserByEmailAndPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "snorlax_v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "GetUserByExternalId",
			Handler:    _UserService_GetUserByExternalId_Handler,
		},
		{
			MethodName: "GetUserByEmailAndPassword",
			Handler:    _UserService_GetUserByEmailAndPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ho-oh/snorlax_v1/snorlax.proto",
}