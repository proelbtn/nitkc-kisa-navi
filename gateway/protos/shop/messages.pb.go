// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shop/messages.proto

package shop

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

type ShopGenre struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShopGenre) Reset()         { *m = ShopGenre{} }
func (m *ShopGenre) String() string { return proto.CompactTextString(m) }
func (*ShopGenre) ProtoMessage()    {}
func (*ShopGenre) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa9375cd29e3697d, []int{0}
}

func (m *ShopGenre) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShopGenre.Unmarshal(m, b)
}
func (m *ShopGenre) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShopGenre.Marshal(b, m, deterministic)
}
func (m *ShopGenre) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShopGenre.Merge(m, src)
}
func (m *ShopGenre) XXX_Size() int {
	return xxx_messageInfo_ShopGenre.Size(m)
}
func (m *ShopGenre) XXX_DiscardUnknown() {
	xxx_messageInfo_ShopGenre.DiscardUnknown(m)
}

var xxx_messageInfo_ShopGenre proto.InternalMessageInfo

func (m *ShopGenre) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ShopGenre) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ShopRecord struct {
	Id                   int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Data                 *ShopData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ShopRecord) Reset()         { *m = ShopRecord{} }
func (m *ShopRecord) String() string { return proto.CompactTextString(m) }
func (*ShopRecord) ProtoMessage()    {}
func (*ShopRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa9375cd29e3697d, []int{1}
}

func (m *ShopRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShopRecord.Unmarshal(m, b)
}
func (m *ShopRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShopRecord.Marshal(b, m, deterministic)
}
func (m *ShopRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShopRecord.Merge(m, src)
}
func (m *ShopRecord) XXX_Size() int {
	return xxx_messageInfo_ShopRecord.Size(m)
}
func (m *ShopRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ShopRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ShopRecord proto.InternalMessageInfo

func (m *ShopRecord) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ShopRecord) GetData() *ShopData {
	if m != nil {
		return m.Data
	}
	return nil
}

type ShopData struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price                uint64   `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	Genre                uint64   `protobuf:"varint,3,opt,name=genre,proto3" json:"genre,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShopData) Reset()         { *m = ShopData{} }
func (m *ShopData) String() string { return proto.CompactTextString(m) }
func (*ShopData) ProtoMessage()    {}
func (*ShopData) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa9375cd29e3697d, []int{2}
}

func (m *ShopData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShopData.Unmarshal(m, b)
}
func (m *ShopData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShopData.Marshal(b, m, deterministic)
}
func (m *ShopData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShopData.Merge(m, src)
}
func (m *ShopData) XXX_Size() int {
	return xxx_messageInfo_ShopData.Size(m)
}
func (m *ShopData) XXX_DiscardUnknown() {
	xxx_messageInfo_ShopData.DiscardUnknown(m)
}

var xxx_messageInfo_ShopData proto.InternalMessageInfo

func (m *ShopData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ShopData) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *ShopData) GetGenre() uint64 {
	if m != nil {
		return m.Genre
	}
	return 0
}

type ShopCreateQuery struct {
	Data                 *ShopData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ShopCreateQuery) Reset()         { *m = ShopCreateQuery{} }
func (m *ShopCreateQuery) String() string { return proto.CompactTextString(m) }
func (*ShopCreateQuery) ProtoMessage()    {}
func (*ShopCreateQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa9375cd29e3697d, []int{3}
}

func (m *ShopCreateQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShopCreateQuery.Unmarshal(m, b)
}
func (m *ShopCreateQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShopCreateQuery.Marshal(b, m, deterministic)
}
func (m *ShopCreateQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShopCreateQuery.Merge(m, src)
}
func (m *ShopCreateQuery) XXX_Size() int {
	return xxx_messageInfo_ShopCreateQuery.Size(m)
}
func (m *ShopCreateQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ShopCreateQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ShopCreateQuery proto.InternalMessageInfo

func (m *ShopCreateQuery) GetData() *ShopData {
	if m != nil {
		return m.Data
	}
	return nil
}

type ShopCreateResult struct {
	Record               *ShopRecord `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ShopCreateResult) Reset()         { *m = ShopCreateResult{} }
func (m *ShopCreateResult) String() string { return proto.CompactTextString(m) }
func (*ShopCreateResult) ProtoMessage()    {}
func (*ShopCreateResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa9375cd29e3697d, []int{4}
}

func (m *ShopCreateResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShopCreateResult.Unmarshal(m, b)
}
func (m *ShopCreateResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShopCreateResult.Marshal(b, m, deterministic)
}
func (m *ShopCreateResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShopCreateResult.Merge(m, src)
}
func (m *ShopCreateResult) XXX_Size() int {
	return xxx_messageInfo_ShopCreateResult.Size(m)
}
func (m *ShopCreateResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ShopCreateResult.DiscardUnknown(m)
}

var xxx_messageInfo_ShopCreateResult proto.InternalMessageInfo

func (m *ShopCreateResult) GetRecord() *ShopRecord {
	if m != nil {
		return m.Record
	}
	return nil
}

type ShopDeleteQuery struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShopDeleteQuery) Reset()         { *m = ShopDeleteQuery{} }
func (m *ShopDeleteQuery) String() string { return proto.CompactTextString(m) }
func (*ShopDeleteQuery) ProtoMessage()    {}
func (*ShopDeleteQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa9375cd29e3697d, []int{5}
}

func (m *ShopDeleteQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShopDeleteQuery.Unmarshal(m, b)
}
func (m *ShopDeleteQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShopDeleteQuery.Marshal(b, m, deterministic)
}
func (m *ShopDeleteQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShopDeleteQuery.Merge(m, src)
}
func (m *ShopDeleteQuery) XXX_Size() int {
	return xxx_messageInfo_ShopDeleteQuery.Size(m)
}
func (m *ShopDeleteQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ShopDeleteQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ShopDeleteQuery proto.InternalMessageInfo

func (m *ShopDeleteQuery) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ShopDeleteResult struct {
	Affected             int64    `protobuf:"varint,1,opt,name=affected,proto3" json:"affected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShopDeleteResult) Reset()         { *m = ShopDeleteResult{} }
func (m *ShopDeleteResult) String() string { return proto.CompactTextString(m) }
func (*ShopDeleteResult) ProtoMessage()    {}
func (*ShopDeleteResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa9375cd29e3697d, []int{6}
}

func (m *ShopDeleteResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShopDeleteResult.Unmarshal(m, b)
}
func (m *ShopDeleteResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShopDeleteResult.Marshal(b, m, deterministic)
}
func (m *ShopDeleteResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShopDeleteResult.Merge(m, src)
}
func (m *ShopDeleteResult) XXX_Size() int {
	return xxx_messageInfo_ShopDeleteResult.Size(m)
}
func (m *ShopDeleteResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ShopDeleteResult.DiscardUnknown(m)
}

var xxx_messageInfo_ShopDeleteResult proto.InternalMessageInfo

func (m *ShopDeleteResult) GetAffected() int64 {
	if m != nil {
		return m.Affected
	}
	return 0
}

type ShopGetQuery struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShopGetQuery) Reset()         { *m = ShopGetQuery{} }
func (m *ShopGetQuery) String() string { return proto.CompactTextString(m) }
func (*ShopGetQuery) ProtoMessage()    {}
func (*ShopGetQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa9375cd29e3697d, []int{7}
}

func (m *ShopGetQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShopGetQuery.Unmarshal(m, b)
}
func (m *ShopGetQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShopGetQuery.Marshal(b, m, deterministic)
}
func (m *ShopGetQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShopGetQuery.Merge(m, src)
}
func (m *ShopGetQuery) XXX_Size() int {
	return xxx_messageInfo_ShopGetQuery.Size(m)
}
func (m *ShopGetQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ShopGetQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ShopGetQuery proto.InternalMessageInfo

func (m *ShopGetQuery) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ShopGetResult struct {
	Record               *ShopRecord `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ShopGetResult) Reset()         { *m = ShopGetResult{} }
func (m *ShopGetResult) String() string { return proto.CompactTextString(m) }
func (*ShopGetResult) ProtoMessage()    {}
func (*ShopGetResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa9375cd29e3697d, []int{8}
}

func (m *ShopGetResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShopGetResult.Unmarshal(m, b)
}
func (m *ShopGetResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShopGetResult.Marshal(b, m, deterministic)
}
func (m *ShopGetResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShopGetResult.Merge(m, src)
}
func (m *ShopGetResult) XXX_Size() int {
	return xxx_messageInfo_ShopGetResult.Size(m)
}
func (m *ShopGetResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ShopGetResult.DiscardUnknown(m)
}

var xxx_messageInfo_ShopGetResult proto.InternalMessageInfo

func (m *ShopGetResult) GetRecord() *ShopRecord {
	if m != nil {
		return m.Record
	}
	return nil
}

type ShopSearchQuery struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Latitude             float64  `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float64  `protobuf:"fixed64,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Genre                uint64   `protobuf:"varint,4,opt,name=genre,proto3" json:"genre,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShopSearchQuery) Reset()         { *m = ShopSearchQuery{} }
func (m *ShopSearchQuery) String() string { return proto.CompactTextString(m) }
func (*ShopSearchQuery) ProtoMessage()    {}
func (*ShopSearchQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa9375cd29e3697d, []int{9}
}

func (m *ShopSearchQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShopSearchQuery.Unmarshal(m, b)
}
func (m *ShopSearchQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShopSearchQuery.Marshal(b, m, deterministic)
}
func (m *ShopSearchQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShopSearchQuery.Merge(m, src)
}
func (m *ShopSearchQuery) XXX_Size() int {
	return xxx_messageInfo_ShopSearchQuery.Size(m)
}
func (m *ShopSearchQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ShopSearchQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ShopSearchQuery proto.InternalMessageInfo

func (m *ShopSearchQuery) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ShopSearchQuery) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *ShopSearchQuery) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *ShopSearchQuery) GetGenre() uint64 {
	if m != nil {
		return m.Genre
	}
	return 0
}

type ShopSearchResult struct {
	Records              []*ShopRecord `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ShopSearchResult) Reset()         { *m = ShopSearchResult{} }
func (m *ShopSearchResult) String() string { return proto.CompactTextString(m) }
func (*ShopSearchResult) ProtoMessage()    {}
func (*ShopSearchResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa9375cd29e3697d, []int{10}
}

func (m *ShopSearchResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShopSearchResult.Unmarshal(m, b)
}
func (m *ShopSearchResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShopSearchResult.Marshal(b, m, deterministic)
}
func (m *ShopSearchResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShopSearchResult.Merge(m, src)
}
func (m *ShopSearchResult) XXX_Size() int {
	return xxx_messageInfo_ShopSearchResult.Size(m)
}
func (m *ShopSearchResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ShopSearchResult.DiscardUnknown(m)
}

var xxx_messageInfo_ShopSearchResult proto.InternalMessageInfo

func (m *ShopSearchResult) GetRecords() []*ShopRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

type ShopEmptyQuery struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShopEmptyQuery) Reset()         { *m = ShopEmptyQuery{} }
func (m *ShopEmptyQuery) String() string { return proto.CompactTextString(m) }
func (*ShopEmptyQuery) ProtoMessage()    {}
func (*ShopEmptyQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa9375cd29e3697d, []int{11}
}

func (m *ShopEmptyQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShopEmptyQuery.Unmarshal(m, b)
}
func (m *ShopEmptyQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShopEmptyQuery.Marshal(b, m, deterministic)
}
func (m *ShopEmptyQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShopEmptyQuery.Merge(m, src)
}
func (m *ShopEmptyQuery) XXX_Size() int {
	return xxx_messageInfo_ShopEmptyQuery.Size(m)
}
func (m *ShopEmptyQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ShopEmptyQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ShopEmptyQuery proto.InternalMessageInfo

type ShopGetGenresResult struct {
	Genres               []*ShopGenre `protobuf:"bytes,1,rep,name=genres,proto3" json:"genres,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ShopGetGenresResult) Reset()         { *m = ShopGetGenresResult{} }
func (m *ShopGetGenresResult) String() string { return proto.CompactTextString(m) }
func (*ShopGetGenresResult) ProtoMessage()    {}
func (*ShopGetGenresResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa9375cd29e3697d, []int{12}
}

func (m *ShopGetGenresResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShopGetGenresResult.Unmarshal(m, b)
}
func (m *ShopGetGenresResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShopGetGenresResult.Marshal(b, m, deterministic)
}
func (m *ShopGetGenresResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShopGetGenresResult.Merge(m, src)
}
func (m *ShopGetGenresResult) XXX_Size() int {
	return xxx_messageInfo_ShopGetGenresResult.Size(m)
}
func (m *ShopGetGenresResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ShopGetGenresResult.DiscardUnknown(m)
}

var xxx_messageInfo_ShopGetGenresResult proto.InternalMessageInfo

func (m *ShopGetGenresResult) GetGenres() []*ShopGenre {
	if m != nil {
		return m.Genres
	}
	return nil
}

func init() {
	proto.RegisterType((*ShopGenre)(nil), "shop.ShopGenre")
	proto.RegisterType((*ShopRecord)(nil), "shop.ShopRecord")
	proto.RegisterType((*ShopData)(nil), "shop.ShopData")
	proto.RegisterType((*ShopCreateQuery)(nil), "shop.ShopCreateQuery")
	proto.RegisterType((*ShopCreateResult)(nil), "shop.ShopCreateResult")
	proto.RegisterType((*ShopDeleteQuery)(nil), "shop.ShopDeleteQuery")
	proto.RegisterType((*ShopDeleteResult)(nil), "shop.ShopDeleteResult")
	proto.RegisterType((*ShopGetQuery)(nil), "shop.ShopGetQuery")
	proto.RegisterType((*ShopGetResult)(nil), "shop.ShopGetResult")
	proto.RegisterType((*ShopSearchQuery)(nil), "shop.ShopSearchQuery")
	proto.RegisterType((*ShopSearchResult)(nil), "shop.ShopSearchResult")
	proto.RegisterType((*ShopEmptyQuery)(nil), "shop.ShopEmptyQuery")
	proto.RegisterType((*ShopGetGenresResult)(nil), "shop.ShopGetGenresResult")
}

func init() { proto.RegisterFile("shop/messages.proto", fileDescriptor_fa9375cd29e3697d) }

var fileDescriptor_fa9375cd29e3697d = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0xab, 0xd3, 0x40,
	0x10, 0x27, 0x4d, 0xac, 0xed, 0xa8, 0x7d, 0x25, 0xcf, 0x43, 0x78, 0x88, 0xd4, 0xbd, 0x58, 0x84,
	0x26, 0xa0, 0x88, 0x14, 0xa4, 0x88, 0x56, 0x04, 0x6f, 0x6e, 0x6f, 0xde, 0xb6, 0xc9, 0x34, 0x09,
	0x24, 0xd9, 0xb0, 0xbb, 0xa9, 0xf4, 0xbf, 0x97, 0xfd, 0x48, 0x52, 0xa4, 0x08, 0xef, 0xd6, 0x99,
	0xdf, 0xfc, 0x3e, 0x76, 0xa6, 0x81, 0x7b, 0x59, 0xf0, 0x36, 0xa9, 0x51, 0x4a, 0x96, 0xa3, 0x8c,
	0x5b, 0xc1, 0x15, 0x0f, 0x03, 0xdd, 0x24, 0x09, 0xcc, 0x0f, 0x05, 0x6f, 0x7f, 0x60, 0x23, 0x30,
	0x5c, 0xc0, 0xa4, 0xcc, 0x22, 0x6f, 0xe5, 0xad, 0x03, 0x3a, 0x29, 0xb3, 0x30, 0x84, 0xa0, 0x61,
	0x35, 0x46, 0x93, 0x95, 0xb7, 0x9e, 0x53, 0xf3, 0x9b, 0x7c, 0x01, 0xd0, 0x04, 0x8a, 0x29, 0x17,
	0xd9, 0x15, 0xc3, 0x37, 0x0c, 0x02, 0x41, 0xc6, 0x14, 0x33, 0x8c, 0x67, 0xef, 0x17, 0xb1, 0xf6,
	0x88, 0xf5, 0xfc, 0x9e, 0x29, 0x46, 0x0d, 0x46, 0x7e, 0xc2, 0xac, 0xef, 0x0c, 0x0e, 0xde, 0xe8,
	0x10, 0xbe, 0x84, 0x27, 0xad, 0x28, 0x53, 0x6b, 0x1b, 0x50, 0x5b, 0xe8, 0x6e, 0xae, 0x43, 0x46,
	0xbe, 0xed, 0x9a, 0x82, 0x7c, 0x84, 0x3b, 0xad, 0xf5, 0x4d, 0x20, 0x53, 0xf8, 0xab, 0x43, 0x71,
	0x19, 0x22, 0x78, 0xff, 0x89, 0xf0, 0x19, 0x96, 0x23, 0x8d, 0xa2, 0xec, 0x2a, 0x15, 0xae, 0x61,
	0x2a, 0xcc, 0xa3, 0x1c, 0x73, 0x39, 0x32, 0xed, 0x63, 0xa9, 0xc3, 0xc9, 0x1b, 0x6b, 0xba, 0xc7,
	0x0a, 0x7b, 0xd3, 0x7f, 0xf6, 0x40, 0x62, 0x6b, 0x60, 0x47, 0x9c, 0xc1, 0x03, 0xcc, 0xd8, 0xe9,
	0x84, 0xa9, 0xc2, 0x7e, 0x72, 0xa8, 0xc9, 0x6b, 0x78, 0x6e, 0xcf, 0xa0, 0x6e, 0xeb, 0x6d, 0xe1,
	0x85, 0xc3, 0x1f, 0x9d, 0xb6, 0xb3, 0x69, 0x0f, 0xc8, 0x44, 0x5a, 0x58, 0xf5, 0x5b, 0x5b, 0x7f,
	0x80, 0x59, 0xc5, 0x54, 0xa9, 0xba, 0xcc, 0x2e, 0xde, 0xa3, 0x43, 0x1d, 0xbe, 0x82, 0x79, 0xc5,
	0x9b, 0xdc, 0x82, 0xbe, 0x01, 0xc7, 0xc6, 0x78, 0x99, 0xe0, 0xfa, 0x32, 0x3b, 0xbb, 0x01, 0x6b,
	0xeb, 0x42, 0xbf, 0x83, 0xa7, 0x36, 0x94, 0x8c, 0xbc, 0x95, 0x7f, 0x33, 0x75, 0x3f, 0x40, 0x96,
	0xb0, 0xd0, 0xed, 0xef, 0x75, 0xab, 0x2e, 0x26, 0x35, 0xd9, 0xc1, 0xbd, 0xdb, 0x81, 0xf9, 0xb7,
	0x4a, 0x27, 0xfa, 0x16, 0xa6, 0xc6, 0xb1, 0xd7, 0xbc, 0x1b, 0x35, 0xcd, 0x1c, 0x75, 0xf0, 0xd7,
	0xed, 0xef, 0x4f, 0x79, 0xa9, 0x8a, 0xee, 0x18, 0xa7, 0xbc, 0x4e, 0x5a, 0xc1, 0xb1, 0x3a, 0xaa,
	0x26, 0x91, 0x69, 0xc1, 0x79, 0xb5, 0xc1, 0x33, 0x6e, 0x1a, 0x76, 0x2e, 0x93, 0x9c, 0x29, 0xfc,
	0xc3, 0x2e, 0x89, 0xf9, 0x4a, 0x64, 0xa2, 0xc5, 0x8e, 0x53, 0x53, 0x7c, 0xf8, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0x1b, 0x62, 0x2d, 0xf5, 0x49, 0x03, 0x00, 0x00,
}