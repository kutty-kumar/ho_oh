// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ho-oh/pkg/core_v1/core.proto

package core_v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
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

type Gender int32

const (
	Gender_unknown_gender Gender = 0
	Gender_male           Gender = 1
	Gender_female         Gender = 2
)

var Gender_name = map[int32]string{
	0: "unknown_gender",
	1: "male",
	2: "female",
}

var Gender_value = map[string]int32{
	"unknown_gender": 0,
	"male":           1,
	"female":         2,
}

func (x Gender) String() string {
	return proto.EnumName(Gender_name, int32(x))
}

func (Gender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e78aff46d42779ed, []int{0}
}

type Status int32

const (
	Status_unknown_status Status = 0
	Status_active         Status = 1
	Status_inactive       Status = 2
)

var Status_name = map[int32]string{
	0: "unknown_status",
	1: "active",
	2: "inactive",
}

var Status_value = map[string]int32{
	"unknown_status": 0,
	"active":         1,
	"inactive":       2,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e78aff46d42779ed, []int{1}
}

type TimeUnit int32

const (
	TimeUnit_unknown_time_unit TimeUnit = 0
	TimeUnit_second            TimeUnit = 2
	TimeUnit_millisecond       TimeUnit = 1
	TimeUnit_minute            TimeUnit = 3
	TimeUnit_hour              TimeUnit = 4
	TimeUnit_day               TimeUnit = 5
	TimeUnit_week              TimeUnit = 6
	TimeUnit_month             TimeUnit = 7
	TimeUnit_year              TimeUnit = 8
)

var TimeUnit_name = map[int32]string{
	0: "unknown_time_unit",
	2: "second",
	1: "millisecond",
	3: "minute",
	4: "hour",
	5: "day",
	6: "week",
	7: "month",
	8: "year",
}

var TimeUnit_value = map[string]int32{
	"unknown_time_unit": 0,
	"second":            2,
	"millisecond":       1,
	"minute":            3,
	"hour":              4,
	"day":               5,
	"week":              6,
	"month":             7,
	"year":              8,
}

func (x TimeUnit) String() string {
	return proto.EnumName(TimeUnit_name, int32(x))
}

func (TimeUnit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e78aff46d42779ed, []int{2}
}

type IdentityType int32

const (
	IdentityType_unknown_identity_type IdentityType = 0
	IdentityType_email                 IdentityType = 2
	IdentityType_phone                 IdentityType = 1
)

var IdentityType_name = map[int32]string{
	0: "unknown_identity_type",
	2: "email",
	1: "phone",
}

var IdentityType_value = map[string]int32{
	"unknown_identity_type": 0,
	"email":                 2,
	"phone":                 1,
}

func (x IdentityType) String() string {
	return proto.EnumName(IdentityType_name, int32(x))
}

func (IdentityType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e78aff46d42779ed, []int{3}
}

type Relation int32

const (
	Relation_unknown_relation Relation = 0
	Relation_husband          Relation = 6
	Relation_father           Relation = 1
	Relation_mother           Relation = 2
	Relation_son              Relation = 3
	Relation_daughter         Relation = 4
	Relation_wife             Relation = 5
)

var Relation_name = map[int32]string{
	0: "unknown_relation",
	6: "husband",
	1: "father",
	2: "mother",
	3: "son",
	4: "daughter",
	5: "wife",
}

var Relation_value = map[string]int32{
	"unknown_relation": 0,
	"husband":          6,
	"father":           1,
	"mother":           2,
	"son":              3,
	"daughter":         4,
	"wife":             5,
}

func (x Relation) String() string {
	return proto.EnumName(Relation_name, int32(x))
}

func (Relation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e78aff46d42779ed, []int{4}
}

func init() {
	proto.RegisterEnum("core.Gender", Gender_name, Gender_value)
	proto.RegisterEnum("core.Status", Status_name, Status_value)
	proto.RegisterEnum("core.TimeUnit", TimeUnit_name, TimeUnit_value)
	proto.RegisterEnum("core.IdentityType", IdentityType_name, IdentityType_value)
	proto.RegisterEnum("core.Relation", Relation_name, Relation_value)
}

func init() { proto.RegisterFile("ho-oh/pkg/core_v1/core.proto", fileDescriptor_e78aff46d42779ed) }

var fileDescriptor_e78aff46d42779ed = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x4d, 0x8b, 0xd4, 0x40,
	0x10, 0x86, 0x67, 0x33, 0x33, 0x99, 0x58, 0xbb, 0x68, 0xd9, 0x38, 0x88, 0x1f, 0xbf, 0x20, 0xa0,
	0x41, 0x05, 0x2f, 0x1e, 0x04, 0x2f, 0xe2, 0x55, 0xd7, 0x8b, 0x97, 0xd0, 0x33, 0xa9, 0xa4, 0x9b,
	0x49, 0x77, 0x85, 0xa4, 0x7a, 0x97, 0x80, 0x3f, 0x5e, 0xba, 0x27, 0x01, 0x4f, 0xfd, 0xe6, 0xa1,
	0x5e, 0x78, 0x52, 0x05, 0x6f, 0x0d, 0xbf, 0x63, 0x53, 0x0d, 0x97, 0xae, 0x3a, 0xf3, 0x48, 0xf5,
	0xc3, 0x87, 0xf4, 0xbe, 0x1f, 0x46, 0x16, 0x56, 0xbb, 0x98, 0x5f, 0xbf, 0xe9, 0x98, 0xbb, 0x9e,
	0xaa, 0xc4, 0x4e, 0xa1, 0xad, 0xc8, 0x0d, 0x32, 0x5f, 0x47, 0xca, 0x8f, 0x90, 0x7f, 0x27, 0xdf,
	0xd0, 0xa8, 0x14, 0x3c, 0x0d, 0xfe, 0xe2, 0xf9, 0xd1, 0xd7, 0x5d, 0x22, 0xb8, 0x51, 0x05, 0xec,
	0x9c, 0xee, 0x09, 0x6f, 0x14, 0x40, 0xde, 0x52, 0xca, 0x59, 0xf9, 0x19, 0xf2, 0x5f, 0xa2, 0x25,
	0x4c, 0xff, 0x77, 0xa6, 0x44, 0x70, 0x13, 0x27, 0xf5, 0x59, 0xec, 0x43, 0x6c, 0xdd, 0x41, 0x61,
	0xfd, 0xf2, 0x95, 0x95, 0x7f, 0xa1, 0xb8, 0xb7, 0x8e, 0x7e, 0x7b, 0x2b, 0xea, 0x08, 0xcf, 0xd7,
	0xa6, 0x58, 0x47, 0x75, 0xf0, 0x56, 0xae, 0xe5, 0x89, 0xce, 0xec, 0x1b, 0xcc, 0xd4, 0x33, 0xb8,
	0x75, 0xb6, 0xef, 0xed, 0x02, 0x92, 0x83, 0xb3, 0x3e, 0x08, 0xe1, 0x36, 0x9a, 0x19, 0x0e, 0x23,
	0xee, 0xd4, 0x01, 0xb6, 0x8d, 0x9e, 0x71, 0x1f, 0xd1, 0x23, 0xd1, 0x05, 0x73, 0xf5, 0x04, 0xf6,
	0x8e, 0xbd, 0x18, 0x3c, 0x44, 0x38, 0x93, 0x1e, 0xb1, 0x28, 0xbf, 0xc2, 0xdd, 0x8f, 0x86, 0xbc,
	0x58, 0x99, 0xef, 0xe7, 0x81, 0xd4, 0x2b, 0x38, 0xae, 0x06, 0x76, 0xe1, 0xb5, 0xcc, 0x03, 0xe1,
	0x26, 0xf6, 0xc9, 0x69, 0xdb, 0x63, 0x16, 0xe3, 0x60, 0xd8, 0x13, 0xde, 0x94, 0x2d, 0x14, 0x3f,
	0xa9, 0xd7, 0x62, 0xd9, 0xab, 0x17, 0x80, 0x6b, 0x79, 0x5c, 0x18, 0x6e, 0xd4, 0x2d, 0x1c, 0x4c,
	0x98, 0x4e, 0xda, 0x37, 0x98, 0xa7, 0x8d, 0x69, 0x31, 0x34, 0x2e, 0xe6, 0x9c, 0x72, 0x16, 0x7d,
	0x27, 0xf6, 0xb8, 0x8d, 0xcb, 0x69, 0x74, 0xe8, 0x8c, 0x50, 0xfc, 0x8d, 0x68, 0x6f, 0x5b, 0xc2,
	0xfd, 0xb7, 0x97, 0x7f, 0x8e, 0x86, 0x6b, 0x36, 0xeb, 0x45, 0xbf, 0x2c, 0xef, 0x29, 0x4f, 0x27,
	0xfb, 0xf4, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x60, 0xd2, 0x64, 0xf5, 0x01, 0x00, 0x00,
}
