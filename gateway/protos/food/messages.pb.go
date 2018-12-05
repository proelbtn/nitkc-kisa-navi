// Code generated by protoc-gen-go. DO NOT EDIT.
// source: food/messages.proto

package food

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

type FoodGenre int32

const (
	FoodGenre_Invalid FoodGenre = 0
	FoodGenre_GenreA  FoodGenre = 1
	FoodGenre_GenreB  FoodGenre = 2
)

var FoodGenre_name = map[int32]string{
	0: "Invalid",
	1: "GenreA",
	2: "GenreB",
}

var FoodGenre_value = map[string]int32{
	"Invalid": 0,
	"GenreA":  1,
	"GenreB":  2,
}

func (x FoodGenre) String() string {
	return proto.EnumName(FoodGenre_name, int32(x))
}

func (FoodGenre) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b52ab5d0a6e96024, []int{0}
}

type FoodRecord struct {
	Id                   int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Data                 *FoodData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *FoodRecord) Reset()         { *m = FoodRecord{} }
func (m *FoodRecord) String() string { return proto.CompactTextString(m) }
func (*FoodRecord) ProtoMessage()    {}
func (*FoodRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52ab5d0a6e96024, []int{0}
}

func (m *FoodRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodRecord.Unmarshal(m, b)
}
func (m *FoodRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodRecord.Marshal(b, m, deterministic)
}
func (m *FoodRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodRecord.Merge(m, src)
}
func (m *FoodRecord) XXX_Size() int {
	return xxx_messageInfo_FoodRecord.Size(m)
}
func (m *FoodRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodRecord.DiscardUnknown(m)
}

var xxx_messageInfo_FoodRecord proto.InternalMessageInfo

func (m *FoodRecord) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FoodRecord) GetData() *FoodData {
	if m != nil {
		return m.Data
	}
	return nil
}

type FoodData struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address              string    `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Latitude             float64   `protobuf:"fixed64,3,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float64   `protobuf:"fixed64,4,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Genre                FoodGenre `protobuf:"varint,5,opt,name=genre,proto3,enum=food.FoodGenre" json:"genre,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *FoodData) Reset()         { *m = FoodData{} }
func (m *FoodData) String() string { return proto.CompactTextString(m) }
func (*FoodData) ProtoMessage()    {}
func (*FoodData) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52ab5d0a6e96024, []int{1}
}

func (m *FoodData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodData.Unmarshal(m, b)
}
func (m *FoodData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodData.Marshal(b, m, deterministic)
}
func (m *FoodData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodData.Merge(m, src)
}
func (m *FoodData) XXX_Size() int {
	return xxx_messageInfo_FoodData.Size(m)
}
func (m *FoodData) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodData.DiscardUnknown(m)
}

var xxx_messageInfo_FoodData proto.InternalMessageInfo

func (m *FoodData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FoodData) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *FoodData) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *FoodData) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *FoodData) GetGenre() FoodGenre {
	if m != nil {
		return m.Genre
	}
	return FoodGenre_Invalid
}

type FoodCreateQuery struct {
	Data                 *FoodData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *FoodCreateQuery) Reset()         { *m = FoodCreateQuery{} }
func (m *FoodCreateQuery) String() string { return proto.CompactTextString(m) }
func (*FoodCreateQuery) ProtoMessage()    {}
func (*FoodCreateQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52ab5d0a6e96024, []int{2}
}

func (m *FoodCreateQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodCreateQuery.Unmarshal(m, b)
}
func (m *FoodCreateQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodCreateQuery.Marshal(b, m, deterministic)
}
func (m *FoodCreateQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodCreateQuery.Merge(m, src)
}
func (m *FoodCreateQuery) XXX_Size() int {
	return xxx_messageInfo_FoodCreateQuery.Size(m)
}
func (m *FoodCreateQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodCreateQuery.DiscardUnknown(m)
}

var xxx_messageInfo_FoodCreateQuery proto.InternalMessageInfo

func (m *FoodCreateQuery) GetData() *FoodData {
	if m != nil {
		return m.Data
	}
	return nil
}

type FoodCreateResult struct {
	Record               *FoodRecord `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *FoodCreateResult) Reset()         { *m = FoodCreateResult{} }
func (m *FoodCreateResult) String() string { return proto.CompactTextString(m) }
func (*FoodCreateResult) ProtoMessage()    {}
func (*FoodCreateResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52ab5d0a6e96024, []int{3}
}

func (m *FoodCreateResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodCreateResult.Unmarshal(m, b)
}
func (m *FoodCreateResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodCreateResult.Marshal(b, m, deterministic)
}
func (m *FoodCreateResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodCreateResult.Merge(m, src)
}
func (m *FoodCreateResult) XXX_Size() int {
	return xxx_messageInfo_FoodCreateResult.Size(m)
}
func (m *FoodCreateResult) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodCreateResult.DiscardUnknown(m)
}

var xxx_messageInfo_FoodCreateResult proto.InternalMessageInfo

func (m *FoodCreateResult) GetRecord() *FoodRecord {
	if m != nil {
		return m.Record
	}
	return nil
}

type FoodDeleteQuery struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FoodDeleteQuery) Reset()         { *m = FoodDeleteQuery{} }
func (m *FoodDeleteQuery) String() string { return proto.CompactTextString(m) }
func (*FoodDeleteQuery) ProtoMessage()    {}
func (*FoodDeleteQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52ab5d0a6e96024, []int{4}
}

func (m *FoodDeleteQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodDeleteQuery.Unmarshal(m, b)
}
func (m *FoodDeleteQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodDeleteQuery.Marshal(b, m, deterministic)
}
func (m *FoodDeleteQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodDeleteQuery.Merge(m, src)
}
func (m *FoodDeleteQuery) XXX_Size() int {
	return xxx_messageInfo_FoodDeleteQuery.Size(m)
}
func (m *FoodDeleteQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodDeleteQuery.DiscardUnknown(m)
}

var xxx_messageInfo_FoodDeleteQuery proto.InternalMessageInfo

func (m *FoodDeleteQuery) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type FoodDeleteResult struct {
	Affected             int64    `protobuf:"varint,1,opt,name=affected,proto3" json:"affected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FoodDeleteResult) Reset()         { *m = FoodDeleteResult{} }
func (m *FoodDeleteResult) String() string { return proto.CompactTextString(m) }
func (*FoodDeleteResult) ProtoMessage()    {}
func (*FoodDeleteResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52ab5d0a6e96024, []int{5}
}

func (m *FoodDeleteResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodDeleteResult.Unmarshal(m, b)
}
func (m *FoodDeleteResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodDeleteResult.Marshal(b, m, deterministic)
}
func (m *FoodDeleteResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodDeleteResult.Merge(m, src)
}
func (m *FoodDeleteResult) XXX_Size() int {
	return xxx_messageInfo_FoodDeleteResult.Size(m)
}
func (m *FoodDeleteResult) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodDeleteResult.DiscardUnknown(m)
}

var xxx_messageInfo_FoodDeleteResult proto.InternalMessageInfo

func (m *FoodDeleteResult) GetAffected() int64 {
	if m != nil {
		return m.Affected
	}
	return 0
}

type FoodGetQuery struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FoodGetQuery) Reset()         { *m = FoodGetQuery{} }
func (m *FoodGetQuery) String() string { return proto.CompactTextString(m) }
func (*FoodGetQuery) ProtoMessage()    {}
func (*FoodGetQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52ab5d0a6e96024, []int{6}
}

func (m *FoodGetQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodGetQuery.Unmarshal(m, b)
}
func (m *FoodGetQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodGetQuery.Marshal(b, m, deterministic)
}
func (m *FoodGetQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodGetQuery.Merge(m, src)
}
func (m *FoodGetQuery) XXX_Size() int {
	return xxx_messageInfo_FoodGetQuery.Size(m)
}
func (m *FoodGetQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodGetQuery.DiscardUnknown(m)
}

var xxx_messageInfo_FoodGetQuery proto.InternalMessageInfo

func (m *FoodGetQuery) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type FoodGetResult struct {
	Record               *FoodRecord `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *FoodGetResult) Reset()         { *m = FoodGetResult{} }
func (m *FoodGetResult) String() string { return proto.CompactTextString(m) }
func (*FoodGetResult) ProtoMessage()    {}
func (*FoodGetResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52ab5d0a6e96024, []int{7}
}

func (m *FoodGetResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodGetResult.Unmarshal(m, b)
}
func (m *FoodGetResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodGetResult.Marshal(b, m, deterministic)
}
func (m *FoodGetResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodGetResult.Merge(m, src)
}
func (m *FoodGetResult) XXX_Size() int {
	return xxx_messageInfo_FoodGetResult.Size(m)
}
func (m *FoodGetResult) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodGetResult.DiscardUnknown(m)
}

var xxx_messageInfo_FoodGetResult proto.InternalMessageInfo

func (m *FoodGetResult) GetRecord() *FoodRecord {
	if m != nil {
		return m.Record
	}
	return nil
}

type FoodSearchQuery struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Latitude             float64   `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float64   `protobuf:"fixed64,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Genre                FoodGenre `protobuf:"varint,4,opt,name=genre,proto3,enum=food.FoodGenre" json:"genre,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *FoodSearchQuery) Reset()         { *m = FoodSearchQuery{} }
func (m *FoodSearchQuery) String() string { return proto.CompactTextString(m) }
func (*FoodSearchQuery) ProtoMessage()    {}
func (*FoodSearchQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52ab5d0a6e96024, []int{8}
}

func (m *FoodSearchQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodSearchQuery.Unmarshal(m, b)
}
func (m *FoodSearchQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodSearchQuery.Marshal(b, m, deterministic)
}
func (m *FoodSearchQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodSearchQuery.Merge(m, src)
}
func (m *FoodSearchQuery) XXX_Size() int {
	return xxx_messageInfo_FoodSearchQuery.Size(m)
}
func (m *FoodSearchQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodSearchQuery.DiscardUnknown(m)
}

var xxx_messageInfo_FoodSearchQuery proto.InternalMessageInfo

func (m *FoodSearchQuery) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FoodSearchQuery) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *FoodSearchQuery) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *FoodSearchQuery) GetGenre() FoodGenre {
	if m != nil {
		return m.Genre
	}
	return FoodGenre_Invalid
}

type FoodSearchResult struct {
	Records              []*FoodRecord `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FoodSearchResult) Reset()         { *m = FoodSearchResult{} }
func (m *FoodSearchResult) String() string { return proto.CompactTextString(m) }
func (*FoodSearchResult) ProtoMessage()    {}
func (*FoodSearchResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b52ab5d0a6e96024, []int{9}
}

func (m *FoodSearchResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodSearchResult.Unmarshal(m, b)
}
func (m *FoodSearchResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodSearchResult.Marshal(b, m, deterministic)
}
func (m *FoodSearchResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodSearchResult.Merge(m, src)
}
func (m *FoodSearchResult) XXX_Size() int {
	return xxx_messageInfo_FoodSearchResult.Size(m)
}
func (m *FoodSearchResult) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodSearchResult.DiscardUnknown(m)
}

var xxx_messageInfo_FoodSearchResult proto.InternalMessageInfo

func (m *FoodSearchResult) GetRecords() []*FoodRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

func init() {
	proto.RegisterEnum("food.FoodGenre", FoodGenre_name, FoodGenre_value)
	proto.RegisterType((*FoodRecord)(nil), "food.FoodRecord")
	proto.RegisterType((*FoodData)(nil), "food.FoodData")
	proto.RegisterType((*FoodCreateQuery)(nil), "food.FoodCreateQuery")
	proto.RegisterType((*FoodCreateResult)(nil), "food.FoodCreateResult")
	proto.RegisterType((*FoodDeleteQuery)(nil), "food.FoodDeleteQuery")
	proto.RegisterType((*FoodDeleteResult)(nil), "food.FoodDeleteResult")
	proto.RegisterType((*FoodGetQuery)(nil), "food.FoodGetQuery")
	proto.RegisterType((*FoodGetResult)(nil), "food.FoodGetResult")
	proto.RegisterType((*FoodSearchQuery)(nil), "food.FoodSearchQuery")
	proto.RegisterType((*FoodSearchResult)(nil), "food.FoodSearchResult")
}

func init() { proto.RegisterFile("food/messages.proto", fileDescriptor_b52ab5d0a6e96024) }

var fileDescriptor_b52ab5d0a6e96024 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x5d, 0x6b, 0xd4, 0x40,
	0x14, 0x75, 0xb2, 0xe9, 0xee, 0xe6, 0x56, 0xb7, 0x61, 0x7c, 0x09, 0x45, 0x24, 0x06, 0x84, 0x50,
	0x68, 0x22, 0x15, 0x91, 0x82, 0x88, 0xd6, 0xa2, 0xf8, 0xe8, 0xf8, 0xe6, 0xdb, 0x6c, 0xe6, 0x6e,
	0x36, 0x90, 0xcd, 0x94, 0x99, 0xc9, 0x4a, 0xff, 0x80, 0xff, 0xc1, 0x7f, 0x2b, 0x99, 0xc9, 0x87,
	0xc8, 0xba, 0xd0, 0xb7, 0xb9, 0xe7, 0xdc, 0x8f, 0x73, 0xef, 0x61, 0xe0, 0xe9, 0x46, 0x4a, 0x91,
	0xef, 0x50, 0x6b, 0x5e, 0xa2, 0xce, 0xee, 0x94, 0x34, 0x92, 0xfa, 0x1d, 0x98, 0x7c, 0x00, 0xf8,
	0x2c, 0xa5, 0x60, 0x58, 0x48, 0x25, 0xe8, 0x0a, 0xbc, 0x4a, 0x44, 0x24, 0x26, 0xe9, 0x8c, 0x79,
	0x95, 0xa0, 0x09, 0xf8, 0x82, 0x1b, 0x1e, 0x79, 0x31, 0x49, 0x4f, 0xaf, 0x56, 0x59, 0x57, 0x92,
	0x75, 0xf9, 0xb7, 0xdc, 0x70, 0x66, 0xb9, 0xe4, 0x37, 0x81, 0xe5, 0x00, 0x51, 0x0a, 0x7e, 0xc3,
	0x77, 0x68, 0x5b, 0x04, 0xcc, 0xbe, 0x69, 0x04, 0x0b, 0x2e, 0x84, 0x42, 0xad, 0x6d, 0x9f, 0x80,
	0x0d, 0x21, 0x3d, 0x87, 0x65, 0xcd, 0x4d, 0x65, 0x5a, 0x81, 0xd1, 0x2c, 0x26, 0x29, 0x61, 0x63,
	0x4c, 0x9f, 0x41, 0x50, 0xcb, 0xa6, 0x74, 0xa4, 0x6f, 0xc9, 0x09, 0xa0, 0x2f, 0xe1, 0xa4, 0xc4,
	0x46, 0x61, 0x74, 0x12, 0x93, 0x74, 0x75, 0x75, 0x36, 0x29, 0xfb, 0xd2, 0xc1, 0xcc, 0xb1, 0xc9,
	0x1b, 0x38, 0xeb, 0xb0, 0x4f, 0x0a, 0xb9, 0xc1, 0x6f, 0x2d, 0xaa, 0xfb, 0x71, 0x25, 0x72, 0x64,
	0xa5, 0x77, 0x10, 0x4e, 0x65, 0x0c, 0x75, 0x5b, 0x1b, 0x9a, 0xc2, 0x5c, 0xd9, 0x23, 0xf5, 0x95,
	0xe1, 0x54, 0xe9, 0x8e, 0xc7, 0x7a, 0x3e, 0x79, 0xe1, 0x86, 0xde, 0x62, 0x8d, 0xc3, 0xd0, 0x7f,
	0xee, 0x9a, 0x64, 0x6e, 0x80, 0x4b, 0xe9, 0x07, 0x9c, 0xc3, 0x92, 0x6f, 0x36, 0x58, 0x18, 0x1c,
	0x32, 0xc7, 0x38, 0x79, 0x0e, 0x8f, 0xdd, 0x6e, 0xe6, 0x70, 0xbf, 0x6b, 0x78, 0xd2, 0xf3, 0x0f,
	0x56, 0xfb, 0x8b, 0x38, 0xb9, 0xdf, 0x91, 0xab, 0x62, 0xeb, 0xda, 0x1f, 0x72, 0xf1, 0x6f, 0xaf,
	0xbc, 0x63, 0x5e, 0xcd, 0xfe, 0xeb, 0x95, 0x7f, 0xd4, 0xab, 0xf7, 0xee, 0x26, 0x4e, 0x47, 0xbf,
	0xc6, 0x05, 0x2c, 0x9c, 0x4c, 0x1d, 0x91, 0x78, 0x76, 0x70, 0x8f, 0x21, 0xe1, 0xe2, 0x15, 0x04,
	0x63, 0x4f, 0x7a, 0x0a, 0x8b, 0xaf, 0xcd, 0x9e, 0xd7, 0x95, 0x08, 0x1f, 0x51, 0x80, 0xb9, 0x45,
	0x3f, 0x86, 0x64, 0x7c, 0xdf, 0x84, 0xde, 0xcd, 0xf5, 0x8f, 0xb7, 0x65, 0x65, 0xb6, 0xed, 0x3a,
	0x2b, 0xe4, 0x2e, 0xbf, 0x53, 0x12, 0xeb, 0xb5, 0x69, 0x72, 0x5d, 0x6c, 0xa5, 0xac, 0x2f, 0x71,
	0x8f, 0x97, 0x0d, 0xdf, 0x57, 0x79, 0xc9, 0x0d, 0xfe, 0xe4, 0xf7, 0xb9, 0xfd, 0x36, 0x3a, 0xef,
	0x04, 0xac, 0xe7, 0x36, 0x78, 0xfd, 0x27, 0x00, 0x00, 0xff, 0xff, 0x11, 0x2f, 0x4c, 0xf8, 0x5a,
	0x03, 0x00, 0x00,
}
