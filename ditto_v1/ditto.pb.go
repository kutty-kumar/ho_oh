// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ho-oh/ditto_v1/ditto.proto

package ditto_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type PrinterDto struct {
	UserId               string               `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	SerialNumber         string               `protobuf:"bytes,4,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	Status               core_v1.Status       `protobuf:"varint,5,opt,name=status,proto3,enum=core.Status" json:"status,omitempty"`
	ExternalId           string               `protobuf:"bytes,6,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	ProductNumber        string               `protobuf:"bytes,7,opt,name=product_number,json=productNumber,proto3" json:"product_number,omitempty"`
	FromDate             *timestamp.Timestamp `protobuf:"bytes,8,opt,name=from_date,json=fromDate,proto3" json:"from_date,omitempty"`
	ToDate               *timestamp.Timestamp `protobuf:"bytes,9,opt,name=to_date,json=toDate,proto3" json:"to_date,omitempty"`
	FromIndex            uint64               `protobuf:"varint,10,opt,name=from_index,json=fromIndex,proto3" json:"from_index,omitempty"`
	ToIndex              uint64               `protobuf:"varint,11,opt,name=to_index,json=toIndex,proto3" json:"to_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PrinterDto) Reset()         { *m = PrinterDto{} }
func (m *PrinterDto) String() string { return proto.CompactTextString(m) }
func (*PrinterDto) ProtoMessage()    {}
func (*PrinterDto) Descriptor() ([]byte, []int) {
	return fileDescriptor_abf6bd2e869b264a, []int{0}
}

func (m *PrinterDto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrinterDto.Unmarshal(m, b)
}
func (m *PrinterDto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrinterDto.Marshal(b, m, deterministic)
}
func (m *PrinterDto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrinterDto.Merge(m, src)
}
func (m *PrinterDto) XXX_Size() int {
	return xxx_messageInfo_PrinterDto.Size(m)
}
func (m *PrinterDto) XXX_DiscardUnknown() {
	xxx_messageInfo_PrinterDto.DiscardUnknown(m)
}

var xxx_messageInfo_PrinterDto proto.InternalMessageInfo

func (m *PrinterDto) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *PrinterDto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PrinterDto) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PrinterDto) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *PrinterDto) GetStatus() core_v1.Status {
	if m != nil {
		return m.Status
	}
	return core_v1.Status_unknown_status
}

func (m *PrinterDto) GetExternalId() string {
	if m != nil {
		return m.ExternalId
	}
	return ""
}

func (m *PrinterDto) GetProductNumber() string {
	if m != nil {
		return m.ProductNumber
	}
	return ""
}

func (m *PrinterDto) GetFromDate() *timestamp.Timestamp {
	if m != nil {
		return m.FromDate
	}
	return nil
}

func (m *PrinterDto) GetToDate() *timestamp.Timestamp {
	if m != nil {
		return m.ToDate
	}
	return nil
}

func (m *PrinterDto) GetFromIndex() uint64 {
	if m != nil {
		return m.FromIndex
	}
	return 0
}

func (m *PrinterDto) GetToIndex() uint64 {
	if m != nil {
		return m.ToIndex
	}
	return 0
}

type CreatePrinterRequest struct {
	Request              *PrinterDto `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreatePrinterRequest) Reset()         { *m = CreatePrinterRequest{} }
func (m *CreatePrinterRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePrinterRequest) ProtoMessage()    {}
func (*CreatePrinterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_abf6bd2e869b264a, []int{1}
}

func (m *CreatePrinterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePrinterRequest.Unmarshal(m, b)
}
func (m *CreatePrinterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePrinterRequest.Marshal(b, m, deterministic)
}
func (m *CreatePrinterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePrinterRequest.Merge(m, src)
}
func (m *CreatePrinterRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePrinterRequest.Size(m)
}
func (m *CreatePrinterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePrinterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePrinterRequest proto.InternalMessageInfo

func (m *CreatePrinterRequest) GetRequest() *PrinterDto {
	if m != nil {
		return m.Request
	}
	return nil
}

type CreatePrinterResponse struct {
	Response             *PrinterDto `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreatePrinterResponse) Reset()         { *m = CreatePrinterResponse{} }
func (m *CreatePrinterResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePrinterResponse) ProtoMessage()    {}
func (*CreatePrinterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_abf6bd2e869b264a, []int{2}
}

func (m *CreatePrinterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePrinterResponse.Unmarshal(m, b)
}
func (m *CreatePrinterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePrinterResponse.Marshal(b, m, deterministic)
}
func (m *CreatePrinterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePrinterResponse.Merge(m, src)
}
func (m *CreatePrinterResponse) XXX_Size() int {
	return xxx_messageInfo_CreatePrinterResponse.Size(m)
}
func (m *CreatePrinterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePrinterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePrinterResponse proto.InternalMessageInfo

func (m *CreatePrinterResponse) GetResponse() *PrinterDto {
	if m != nil {
		return m.Response
	}
	return nil
}

type UpdatePrinterRequest struct {
	PrinterId            string      `protobuf:"bytes,1,opt,name=printer_id,json=printerId,proto3" json:"printer_id,omitempty"`
	Request              *PrinterDto `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpdatePrinterRequest) Reset()         { *m = UpdatePrinterRequest{} }
func (m *UpdatePrinterRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePrinterRequest) ProtoMessage()    {}
func (*UpdatePrinterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_abf6bd2e869b264a, []int{3}
}

func (m *UpdatePrinterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePrinterRequest.Unmarshal(m, b)
}
func (m *UpdatePrinterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePrinterRequest.Marshal(b, m, deterministic)
}
func (m *UpdatePrinterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePrinterRequest.Merge(m, src)
}
func (m *UpdatePrinterRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePrinterRequest.Size(m)
}
func (m *UpdatePrinterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePrinterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePrinterRequest proto.InternalMessageInfo

func (m *UpdatePrinterRequest) GetPrinterId() string {
	if m != nil {
		return m.PrinterId
	}
	return ""
}

func (m *UpdatePrinterRequest) GetRequest() *PrinterDto {
	if m != nil {
		return m.Request
	}
	return nil
}

type UpdatePrinterResponse struct {
	Response             *PrinterDto `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpdatePrinterResponse) Reset()         { *m = UpdatePrinterResponse{} }
func (m *UpdatePrinterResponse) String() string { return proto.CompactTextString(m) }
func (*UpdatePrinterResponse) ProtoMessage()    {}
func (*UpdatePrinterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_abf6bd2e869b264a, []int{4}
}

func (m *UpdatePrinterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePrinterResponse.Unmarshal(m, b)
}
func (m *UpdatePrinterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePrinterResponse.Marshal(b, m, deterministic)
}
func (m *UpdatePrinterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePrinterResponse.Merge(m, src)
}
func (m *UpdatePrinterResponse) XXX_Size() int {
	return xxx_messageInfo_UpdatePrinterResponse.Size(m)
}
func (m *UpdatePrinterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePrinterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePrinterResponse proto.InternalMessageInfo

func (m *UpdatePrinterResponse) GetResponse() *PrinterDto {
	if m != nil {
		return m.Response
	}
	return nil
}

type GetPrinterByExternalIdRequest struct {
	PrinterId            string   `protobuf:"bytes,1,opt,name=printer_id,json=printerId,proto3" json:"printer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPrinterByExternalIdRequest) Reset()         { *m = GetPrinterByExternalIdRequest{} }
func (m *GetPrinterByExternalIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetPrinterByExternalIdRequest) ProtoMessage()    {}
func (*GetPrinterByExternalIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_abf6bd2e869b264a, []int{5}
}

func (m *GetPrinterByExternalIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPrinterByExternalIdRequest.Unmarshal(m, b)
}
func (m *GetPrinterByExternalIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPrinterByExternalIdRequest.Marshal(b, m, deterministic)
}
func (m *GetPrinterByExternalIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPrinterByExternalIdRequest.Merge(m, src)
}
func (m *GetPrinterByExternalIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetPrinterByExternalIdRequest.Size(m)
}
func (m *GetPrinterByExternalIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPrinterByExternalIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPrinterByExternalIdRequest proto.InternalMessageInfo

func (m *GetPrinterByExternalIdRequest) GetPrinterId() string {
	if m != nil {
		return m.PrinterId
	}
	return ""
}

type GetPrinterByExternalIdResponse struct {
	Response             *PrinterDto `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetPrinterByExternalIdResponse) Reset()         { *m = GetPrinterByExternalIdResponse{} }
func (m *GetPrinterByExternalIdResponse) String() string { return proto.CompactTextString(m) }
func (*GetPrinterByExternalIdResponse) ProtoMessage()    {}
func (*GetPrinterByExternalIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_abf6bd2e869b264a, []int{6}
}

func (m *GetPrinterByExternalIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPrinterByExternalIdResponse.Unmarshal(m, b)
}
func (m *GetPrinterByExternalIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPrinterByExternalIdResponse.Marshal(b, m, deterministic)
}
func (m *GetPrinterByExternalIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPrinterByExternalIdResponse.Merge(m, src)
}
func (m *GetPrinterByExternalIdResponse) XXX_Size() int {
	return xxx_messageInfo_GetPrinterByExternalIdResponse.Size(m)
}
func (m *GetPrinterByExternalIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPrinterByExternalIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPrinterByExternalIdResponse proto.InternalMessageInfo

func (m *GetPrinterByExternalIdResponse) GetResponse() *PrinterDto {
	if m != nil {
		return m.Response
	}
	return nil
}

type MultiGetPrintersByExternalIdRequest struct {
	PrinterIds           []string `protobuf:"bytes,1,rep,name=printer_ids,json=printerIds,proto3" json:"printer_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiGetPrintersByExternalIdRequest) Reset()         { *m = MultiGetPrintersByExternalIdRequest{} }
func (m *MultiGetPrintersByExternalIdRequest) String() string { return proto.CompactTextString(m) }
func (*MultiGetPrintersByExternalIdRequest) ProtoMessage()    {}
func (*MultiGetPrintersByExternalIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_abf6bd2e869b264a, []int{7}
}

func (m *MultiGetPrintersByExternalIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiGetPrintersByExternalIdRequest.Unmarshal(m, b)
}
func (m *MultiGetPrintersByExternalIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiGetPrintersByExternalIdRequest.Marshal(b, m, deterministic)
}
func (m *MultiGetPrintersByExternalIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiGetPrintersByExternalIdRequest.Merge(m, src)
}
func (m *MultiGetPrintersByExternalIdRequest) XXX_Size() int {
	return xxx_messageInfo_MultiGetPrintersByExternalIdRequest.Size(m)
}
func (m *MultiGetPrintersByExternalIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiGetPrintersByExternalIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MultiGetPrintersByExternalIdRequest proto.InternalMessageInfo

func (m *MultiGetPrintersByExternalIdRequest) GetPrinterIds() []string {
	if m != nil {
		return m.PrinterIds
	}
	return nil
}

type MultiGetPrintersByExternalIdResponse struct {
	Result               []*PrinterDto `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MultiGetPrintersByExternalIdResponse) Reset()         { *m = MultiGetPrintersByExternalIdResponse{} }
func (m *MultiGetPrintersByExternalIdResponse) String() string { return proto.CompactTextString(m) }
func (*MultiGetPrintersByExternalIdResponse) ProtoMessage()    {}
func (*MultiGetPrintersByExternalIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_abf6bd2e869b264a, []int{8}
}

func (m *MultiGetPrintersByExternalIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiGetPrintersByExternalIdResponse.Unmarshal(m, b)
}
func (m *MultiGetPrintersByExternalIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiGetPrintersByExternalIdResponse.Marshal(b, m, deterministic)
}
func (m *MultiGetPrintersByExternalIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiGetPrintersByExternalIdResponse.Merge(m, src)
}
func (m *MultiGetPrintersByExternalIdResponse) XXX_Size() int {
	return xxx_messageInfo_MultiGetPrintersByExternalIdResponse.Size(m)
}
func (m *MultiGetPrintersByExternalIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiGetPrintersByExternalIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MultiGetPrintersByExternalIdResponse proto.InternalMessageInfo

func (m *MultiGetPrintersByExternalIdResponse) GetResult() []*PrinterDto {
	if m != nil {
		return m.Result
	}
	return nil
}

type NoOpRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NoOpRequest) Reset()         { *m = NoOpRequest{} }
func (m *NoOpRequest) String() string { return proto.CompactTextString(m) }
func (*NoOpRequest) ProtoMessage()    {}
func (*NoOpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_abf6bd2e869b264a, []int{9}
}

func (m *NoOpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NoOpRequest.Unmarshal(m, b)
}
func (m *NoOpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NoOpRequest.Marshal(b, m, deterministic)
}
func (m *NoOpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoOpRequest.Merge(m, src)
}
func (m *NoOpRequest) XXX_Size() int {
	return xxx_messageInfo_NoOpRequest.Size(m)
}
func (m *NoOpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NoOpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NoOpRequest proto.InternalMessageInfo

type DeletePrinterRequest struct {
	PrinterId            string   `protobuf:"bytes,1,opt,name=printer_id,json=printerId,proto3" json:"printer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePrinterRequest) Reset()         { *m = DeletePrinterRequest{} }
func (m *DeletePrinterRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePrinterRequest) ProtoMessage()    {}
func (*DeletePrinterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_abf6bd2e869b264a, []int{10}
}

func (m *DeletePrinterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePrinterRequest.Unmarshal(m, b)
}
func (m *DeletePrinterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePrinterRequest.Marshal(b, m, deterministic)
}
func (m *DeletePrinterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePrinterRequest.Merge(m, src)
}
func (m *DeletePrinterRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePrinterRequest.Size(m)
}
func (m *DeletePrinterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePrinterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePrinterRequest proto.InternalMessageInfo

func (m *DeletePrinterRequest) GetPrinterId() string {
	if m != nil {
		return m.PrinterId
	}
	return ""
}

func init() {
	proto.RegisterType((*PrinterDto)(nil), "ditto_v1.PrinterDto")
	proto.RegisterType((*CreatePrinterRequest)(nil), "ditto_v1.CreatePrinterRequest")
	proto.RegisterType((*CreatePrinterResponse)(nil), "ditto_v1.CreatePrinterResponse")
	proto.RegisterType((*UpdatePrinterRequest)(nil), "ditto_v1.UpdatePrinterRequest")
	proto.RegisterType((*UpdatePrinterResponse)(nil), "ditto_v1.UpdatePrinterResponse")
	proto.RegisterType((*GetPrinterByExternalIdRequest)(nil), "ditto_v1.GetPrinterByExternalIdRequest")
	proto.RegisterType((*GetPrinterByExternalIdResponse)(nil), "ditto_v1.GetPrinterByExternalIdResponse")
	proto.RegisterType((*MultiGetPrintersByExternalIdRequest)(nil), "ditto_v1.MultiGetPrintersByExternalIdRequest")
	proto.RegisterType((*MultiGetPrintersByExternalIdResponse)(nil), "ditto_v1.MultiGetPrintersByExternalIdResponse")
	proto.RegisterType((*NoOpRequest)(nil), "ditto_v1.NoOpRequest")
	proto.RegisterType((*DeletePrinterRequest)(nil), "ditto_v1.DeletePrinterRequest")
}

func init() { proto.RegisterFile("ho-oh/ditto_v1/ditto.proto", fileDescriptor_abf6bd2e869b264a) }

var fileDescriptor_abf6bd2e869b264a = []byte{
	// 721 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0x4d, 0x53, 0xd3, 0x40,
	0x18, 0xc7, 0x27, 0x80, 0x7d, 0x79, 0x4a, 0x39, 0xec, 0x14, 0x08, 0x81, 0xd2, 0x1a, 0x50, 0x3b,
	0x8c, 0x4d, 0xa5, 0x8c, 0xe3, 0x0c, 0x07, 0x0f, 0x8a, 0x38, 0x3d, 0x88, 0x4e, 0x80, 0x73, 0x27,
	0x34, 0x0f, 0x90, 0x99, 0x36, 0x1b, 0x77, 0x37, 0x0c, 0x8e, 0x7a, 0xf1, 0xe4, 0xc9, 0x8b, 0x33,
	0x9e, 0xfd, 0x0e, 0x7e, 0x14, 0xbf, 0x82, 0x1f, 0xc4, 0xc9, 0x66, 0x43, 0x52, 0x4c, 0xa1, 0x78,
	0xea, 0xe6, 0x79, 0xfb, 0xff, 0xb2, 0xcf, 0xbf, 0x13, 0x30, 0xce, 0x69, 0x9b, 0x9e, 0x77, 0x5c,
	0x4f, 0x08, 0xda, 0xbf, 0xd8, 0x8e, 0x0f, 0x56, 0xc0, 0xa8, 0xa0, 0xa4, 0x94, 0x44, 0x0d, 0x32,
	0xa0, 0x0c, 0xa3, 0x74, 0xf4, 0x1b, 0x67, 0x8d, 0xb5, 0x33, 0x4a, 0xcf, 0x86, 0xd8, 0x71, 0x02,
	0xaf, 0xe3, 0xf8, 0x3e, 0x15, 0x8e, 0xf0, 0xa8, 0xcf, 0x55, 0xb6, 0xa1, 0xb2, 0xf2, 0xe9, 0x24,
	0x3c, 0xed, 0x08, 0x6f, 0x84, 0x5c, 0x38, 0xa3, 0x20, 0x2e, 0x30, 0x7f, 0xcc, 0x02, 0xbc, 0x63,
	0x9e, 0x2f, 0x90, 0xed, 0x09, 0x4a, 0x96, 0xa1, 0x18, 0x72, 0x64, 0x7d, 0xcf, 0xd5, 0xb5, 0xa6,
	0xd6, 0x2a, 0xdb, 0x85, 0xe8, 0xb1, 0xe7, 0x12, 0x02, 0x73, 0xbe, 0x33, 0x42, 0x7d, 0x46, 0x46,
	0xe5, 0x99, 0x34, 0xa1, 0xe2, 0x22, 0x1f, 0x30, 0x2f, 0x88, 0x24, 0xf5, 0x59, 0x99, 0xca, 0x86,
	0xc8, 0x06, 0x54, 0x39, 0x32, 0xcf, 0x19, 0xf6, 0xfd, 0x70, 0x74, 0x82, 0x4c, 0x9f, 0x93, 0x35,
	0xf3, 0x71, 0xf0, 0x40, 0xc6, 0xc8, 0x26, 0x14, 0xb8, 0x70, 0x44, 0xc8, 0xf5, 0x7b, 0x4d, 0xad,
	0xb5, 0xd0, 0x9d, 0xb7, 0xe4, 0xeb, 0x1d, 0xca, 0x98, 0xad, 0x72, 0xa4, 0x01, 0x15, 0xbc, 0x14,
	0xc8, 0x7c, 0x67, 0x18, 0xd1, 0x15, 0xe4, 0x20, 0x48, 0x42, 0x3d, 0x97, 0x3c, 0x80, 0x85, 0x80,
	0x51, 0x37, 0x1c, 0x88, 0x44, 0xac, 0x28, 0x6b, 0xaa, 0x2a, 0xaa, 0xd4, 0x9e, 0x41, 0xf9, 0x94,
	0xd1, 0x51, 0xdf, 0x75, 0x04, 0xea, 0xa5, 0xa6, 0xd6, 0xaa, 0x74, 0x0d, 0x2b, 0xbe, 0x25, 0x2b,
	0xb9, 0x25, 0xeb, 0x28, 0xb9, 0x25, 0xbb, 0x14, 0x15, 0xef, 0x39, 0x02, 0xc9, 0x0e, 0x14, 0x05,
	0x8d, 0xdb, 0xca, 0xb7, 0xb6, 0x15, 0x04, 0x95, 0x4d, 0x75, 0x00, 0xa9, 0xe6, 0xf9, 0x2e, 0x5e,
	0xea, 0xd0, 0xd4, 0x5a, 0x73, 0xb6, 0xd4, 0xef, 0x45, 0x01, 0xb2, 0x02, 0x25, 0x41, 0x55, 0xb2,
	0x22, 0x93, 0x45, 0x41, 0x65, 0xca, 0xdc, 0x87, 0xda, 0x4b, 0x86, 0x8e, 0x40, 0xb5, 0x1d, 0x1b,
	0xdf, 0x87, 0xc8, 0x05, 0xb1, 0xa0, 0xc8, 0xe2, 0xa3, 0xdc, 0x50, 0xa5, 0x5b, 0xb3, 0x12, 0x7f,
	0x58, 0xe9, 0x22, 0xed, 0xa4, 0xc8, 0xec, 0xc1, 0xe2, 0xb5, 0x39, 0x3c, 0xa0, 0x3e, 0x47, 0xf2,
	0x04, 0x4a, 0x4c, 0x9d, 0xe5, 0x56, 0x27, 0x4d, 0xba, 0xaa, 0x32, 0x11, 0x6a, 0xc7, 0x81, 0xfb,
	0x2f, 0x52, 0x1d, 0x20, 0x88, 0x23, 0xa9, 0x6f, 0xca, 0x2a, 0xd2, 0x73, 0xb3, 0xc4, 0x33, 0x53,
	0x12, 0x5f, 0x93, 0xc9, 0x21, 0xd6, 0xa6, 0x22, 0x7e, 0x0e, 0xf5, 0xd7, 0x28, 0x54, 0xea, 0xc5,
	0x87, 0x57, 0x57, 0x6e, 0x99, 0x0e, 0xdd, 0xb4, 0x61, 0x7d, 0x52, 0xff, 0x7f, 0x33, 0xed, 0xc3,
	0xc6, 0x9b, 0x70, 0x28, 0xbc, 0x74, 0x30, 0xcf, 0x23, 0x6b, 0x40, 0x25, 0x25, 0xe3, 0xba, 0xd6,
	0x9c, 0x8d, 0xfc, 0x7e, 0x85, 0xc6, 0xcd, 0x23, 0xd8, 0xbc, 0x79, 0x8e, 0x22, 0x7c, 0x0c, 0x05,
	0x86, 0x3c, 0x1c, 0x0a, 0x39, 0x63, 0x12, 0x9f, 0xaa, 0x31, 0xab, 0x50, 0x39, 0xa0, 0x6f, 0x03,
	0x45, 0x61, 0x3e, 0x85, 0xda, 0x1e, 0x0e, 0xf1, 0x8e, 0x2b, 0xef, 0xfe, 0x2a, 0xc0, 0x82, 0xea,
	0x38, 0x44, 0x76, 0xe1, 0x0d, 0x90, 0x50, 0xa8, 0x8e, 0xf9, 0x90, 0xac, 0xa7, 0x1c, 0x79, 0x46,
	0x37, 0x1a, 0x13, 0xf3, 0xea, 0x22, 0xeb, 0x5f, 0x7e, 0xff, 0xf9, 0x3e, 0xb3, 0x6c, 0xce, 0x77,
	0x2e, 0xb6, 0x3b, 0x4a, 0x9b, 0xef, 0x26, 0x36, 0x22, 0x9f, 0xa0, 0x3a, 0x66, 0xa3, 0xac, 0x60,
	0x9e, 0x8d, 0xb3, 0x82, 0xb9, 0xfe, 0x33, 0xb7, 0xa4, 0xe0, 0x66, 0x77, 0x25, 0x2b, 0xd8, 0xf9,
	0x98, 0x5e, 0xc4, 0xe7, 0x54, 0xfd, 0x9b, 0x06, 0x4b, 0xf9, 0xd6, 0x21, 0x8f, 0x52, 0x9d, 0x1b,
	0xcd, 0x69, 0xb4, 0x6e, 0x2f, 0x54, 0x64, 0xf7, 0x25, 0xd9, 0x2a, 0x99, 0x4c, 0x46, 0x7e, 0x6a,
	0xb0, 0x76, 0x93, 0x5f, 0x48, 0x3b, 0x55, 0x9b, 0xc2, 0x9f, 0x86, 0x35, 0x6d, 0xf9, 0x38, 0xa2,
	0xb1, 0x34, 0x86, 0x38, 0x8a, 0x5a, 0xdb, 0x67, 0x28, 0x76, 0xb5, 0x2d, 0xf2, 0x55, 0x83, 0xe5,
	0xeb, 0xb3, 0xf6, 0x29, 0x3b, 0xe6, 0xc8, 0xc8, 0x62, 0x2a, 0x97, 0xf1, 0xe7, 0x9d, 0x29, 0x1e,
	0x4a, 0x8a, 0xa6, 0xb1, 0x1a, 0x51, 0x44, 0x9f, 0xb6, 0x76, 0x3e, 0x0a, 0x87, 0xea, 0x98, 0xef,
	0xb3, 0xe6, 0xc9, 0xfb, 0x43, 0xdc, 0x6e, 0x1e, 0xf5, 0xfe, 0x5b, 0x93, 0x57, 0x74, 0x52, 0x90,
	0x1f, 0x92, 0x9d, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x82, 0x7b, 0xd8, 0xb2, 0x0d, 0x08, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PrinterServiceClient is the client API for PrinterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PrinterServiceClient interface {
	CreatePrinter(ctx context.Context, in *CreatePrinterRequest, opts ...grpc.CallOption) (*CreatePrinterResponse, error)
	UpdatePrinter(ctx context.Context, in *UpdatePrinterRequest, opts ...grpc.CallOption) (*UpdatePrinterResponse, error)
	GetPrinterByExternalId(ctx context.Context, in *GetPrinterByExternalIdRequest, opts ...grpc.CallOption) (*GetPrinterByExternalIdResponse, error)
	MultiGetPrintersByExternalId(ctx context.Context, in *MultiGetPrintersByExternalIdRequest, opts ...grpc.CallOption) (*MultiGetPrintersByExternalIdResponse, error)
	MultiGetPrintersForUser(ctx context.Context, in *NoOpRequest, opts ...grpc.CallOption) (*MultiGetPrintersByExternalIdResponse, error)
	DeletePrinter(ctx context.Context, in *DeletePrinterRequest, opts ...grpc.CallOption) (*UpdatePrinterResponse, error)
}

type printerServiceClient struct {
	cc *grpc.ClientConn
}

func NewPrinterServiceClient(cc *grpc.ClientConn) PrinterServiceClient {
	return &printerServiceClient{cc}
}

func (c *printerServiceClient) CreatePrinter(ctx context.Context, in *CreatePrinterRequest, opts ...grpc.CallOption) (*CreatePrinterResponse, error) {
	out := new(CreatePrinterResponse)
	err := c.cc.Invoke(ctx, "/ditto_v1.PrinterService/CreatePrinter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *printerServiceClient) UpdatePrinter(ctx context.Context, in *UpdatePrinterRequest, opts ...grpc.CallOption) (*UpdatePrinterResponse, error) {
	out := new(UpdatePrinterResponse)
	err := c.cc.Invoke(ctx, "/ditto_v1.PrinterService/UpdatePrinter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *printerServiceClient) GetPrinterByExternalId(ctx context.Context, in *GetPrinterByExternalIdRequest, opts ...grpc.CallOption) (*GetPrinterByExternalIdResponse, error) {
	out := new(GetPrinterByExternalIdResponse)
	err := c.cc.Invoke(ctx, "/ditto_v1.PrinterService/GetPrinterByExternalId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *printerServiceClient) MultiGetPrintersByExternalId(ctx context.Context, in *MultiGetPrintersByExternalIdRequest, opts ...grpc.CallOption) (*MultiGetPrintersByExternalIdResponse, error) {
	out := new(MultiGetPrintersByExternalIdResponse)
	err := c.cc.Invoke(ctx, "/ditto_v1.PrinterService/MultiGetPrintersByExternalId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *printerServiceClient) MultiGetPrintersForUser(ctx context.Context, in *NoOpRequest, opts ...grpc.CallOption) (*MultiGetPrintersByExternalIdResponse, error) {
	out := new(MultiGetPrintersByExternalIdResponse)
	err := c.cc.Invoke(ctx, "/ditto_v1.PrinterService/MultiGetPrintersForUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *printerServiceClient) DeletePrinter(ctx context.Context, in *DeletePrinterRequest, opts ...grpc.CallOption) (*UpdatePrinterResponse, error) {
	out := new(UpdatePrinterResponse)
	err := c.cc.Invoke(ctx, "/ditto_v1.PrinterService/DeletePrinter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrinterServiceServer is the server API for PrinterService service.
type PrinterServiceServer interface {
	CreatePrinter(context.Context, *CreatePrinterRequest) (*CreatePrinterResponse, error)
	UpdatePrinter(context.Context, *UpdatePrinterRequest) (*UpdatePrinterResponse, error)
	GetPrinterByExternalId(context.Context, *GetPrinterByExternalIdRequest) (*GetPrinterByExternalIdResponse, error)
	MultiGetPrintersByExternalId(context.Context, *MultiGetPrintersByExternalIdRequest) (*MultiGetPrintersByExternalIdResponse, error)
	MultiGetPrintersForUser(context.Context, *NoOpRequest) (*MultiGetPrintersByExternalIdResponse, error)
	DeletePrinter(context.Context, *DeletePrinterRequest) (*UpdatePrinterResponse, error)
}

func RegisterPrinterServiceServer(s *grpc.Server, srv PrinterServiceServer) {
	s.RegisterService(&_PrinterService_serviceDesc, srv)
}

func _PrinterService_CreatePrinter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePrinterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrinterServiceServer).CreatePrinter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ditto_v1.PrinterService/CreatePrinter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrinterServiceServer).CreatePrinter(ctx, req.(*CreatePrinterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrinterService_UpdatePrinter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePrinterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrinterServiceServer).UpdatePrinter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ditto_v1.PrinterService/UpdatePrinter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrinterServiceServer).UpdatePrinter(ctx, req.(*UpdatePrinterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrinterService_GetPrinterByExternalId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrinterByExternalIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrinterServiceServer).GetPrinterByExternalId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ditto_v1.PrinterService/GetPrinterByExternalId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrinterServiceServer).GetPrinterByExternalId(ctx, req.(*GetPrinterByExternalIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrinterService_MultiGetPrintersByExternalId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiGetPrintersByExternalIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrinterServiceServer).MultiGetPrintersByExternalId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ditto_v1.PrinterService/MultiGetPrintersByExternalId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrinterServiceServer).MultiGetPrintersByExternalId(ctx, req.(*MultiGetPrintersByExternalIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrinterService_MultiGetPrintersForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoOpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrinterServiceServer).MultiGetPrintersForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ditto_v1.PrinterService/MultiGetPrintersForUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrinterServiceServer).MultiGetPrintersForUser(ctx, req.(*NoOpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrinterService_DeletePrinter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePrinterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrinterServiceServer).DeletePrinter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ditto_v1.PrinterService/DeletePrinter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrinterServiceServer).DeletePrinter(ctx, req.(*DeletePrinterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PrinterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ditto_v1.PrinterService",
	HandlerType: (*PrinterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePrinter",
			Handler:    _PrinterService_CreatePrinter_Handler,
		},
		{
			MethodName: "UpdatePrinter",
			Handler:    _PrinterService_UpdatePrinter_Handler,
		},
		{
			MethodName: "GetPrinterByExternalId",
			Handler:    _PrinterService_GetPrinterByExternalId_Handler,
		},
		{
			MethodName: "MultiGetPrintersByExternalId",
			Handler:    _PrinterService_MultiGetPrintersByExternalId_Handler,
		},
		{
			MethodName: "MultiGetPrintersForUser",
			Handler:    _PrinterService_MultiGetPrintersForUser_Handler,
		},
		{
			MethodName: "DeletePrinter",
			Handler:    _PrinterService_DeletePrinter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ho-oh/ditto_v1/ditto.proto",
}
