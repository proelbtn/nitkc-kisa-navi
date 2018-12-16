// Code generated by protoc-gen-go. DO NOT EDIT.
// source: souvenir/messages.proto

package souvenir

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

type SouvenirGenre struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SouvenirGenre) Reset()         { *m = SouvenirGenre{} }
func (m *SouvenirGenre) String() string { return proto.CompactTextString(m) }
func (*SouvenirGenre) ProtoMessage()    {}
func (*SouvenirGenre) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8f236c690c7bcc2, []int{0}
}

func (m *SouvenirGenre) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SouvenirGenre.Unmarshal(m, b)
}
func (m *SouvenirGenre) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SouvenirGenre.Marshal(b, m, deterministic)
}
func (m *SouvenirGenre) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SouvenirGenre.Merge(m, src)
}
func (m *SouvenirGenre) XXX_Size() int {
	return xxx_messageInfo_SouvenirGenre.Size(m)
}
func (m *SouvenirGenre) XXX_DiscardUnknown() {
	xxx_messageInfo_SouvenirGenre.DiscardUnknown(m)
}

var xxx_messageInfo_SouvenirGenre proto.InternalMessageInfo

func (m *SouvenirGenre) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SouvenirGenre) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SouvenirRecord struct {
	Id                   uint64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Data                 *SouvenirData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SouvenirRecord) Reset()         { *m = SouvenirRecord{} }
func (m *SouvenirRecord) String() string { return proto.CompactTextString(m) }
func (*SouvenirRecord) ProtoMessage()    {}
func (*SouvenirRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8f236c690c7bcc2, []int{1}
}

func (m *SouvenirRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SouvenirRecord.Unmarshal(m, b)
}
func (m *SouvenirRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SouvenirRecord.Marshal(b, m, deterministic)
}
func (m *SouvenirRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SouvenirRecord.Merge(m, src)
}
func (m *SouvenirRecord) XXX_Size() int {
	return xxx_messageInfo_SouvenirRecord.Size(m)
}
func (m *SouvenirRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_SouvenirRecord.DiscardUnknown(m)
}

var xxx_messageInfo_SouvenirRecord proto.InternalMessageInfo

func (m *SouvenirRecord) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SouvenirRecord) GetData() *SouvenirData {
	if m != nil {
		return m.Data
	}
	return nil
}

type SouvenirData struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price                uint64   `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	GenreId              uint64   `protobuf:"varint,3,opt,name=genre_id,json=genreId,proto3" json:"genre_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SouvenirData) Reset()         { *m = SouvenirData{} }
func (m *SouvenirData) String() string { return proto.CompactTextString(m) }
func (*SouvenirData) ProtoMessage()    {}
func (*SouvenirData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8f236c690c7bcc2, []int{2}
}

func (m *SouvenirData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SouvenirData.Unmarshal(m, b)
}
func (m *SouvenirData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SouvenirData.Marshal(b, m, deterministic)
}
func (m *SouvenirData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SouvenirData.Merge(m, src)
}
func (m *SouvenirData) XXX_Size() int {
	return xxx_messageInfo_SouvenirData.Size(m)
}
func (m *SouvenirData) XXX_DiscardUnknown() {
	xxx_messageInfo_SouvenirData.DiscardUnknown(m)
}

var xxx_messageInfo_SouvenirData proto.InternalMessageInfo

func (m *SouvenirData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SouvenirData) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *SouvenirData) GetGenreId() uint64 {
	if m != nil {
		return m.GenreId
	}
	return 0
}

type SouvenirCreateQuery struct {
	Data                 *SouvenirData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SouvenirCreateQuery) Reset()         { *m = SouvenirCreateQuery{} }
func (m *SouvenirCreateQuery) String() string { return proto.CompactTextString(m) }
func (*SouvenirCreateQuery) ProtoMessage()    {}
func (*SouvenirCreateQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8f236c690c7bcc2, []int{3}
}

func (m *SouvenirCreateQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SouvenirCreateQuery.Unmarshal(m, b)
}
func (m *SouvenirCreateQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SouvenirCreateQuery.Marshal(b, m, deterministic)
}
func (m *SouvenirCreateQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SouvenirCreateQuery.Merge(m, src)
}
func (m *SouvenirCreateQuery) XXX_Size() int {
	return xxx_messageInfo_SouvenirCreateQuery.Size(m)
}
func (m *SouvenirCreateQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_SouvenirCreateQuery.DiscardUnknown(m)
}

var xxx_messageInfo_SouvenirCreateQuery proto.InternalMessageInfo

func (m *SouvenirCreateQuery) GetData() *SouvenirData {
	if m != nil {
		return m.Data
	}
	return nil
}

type SouvenirCreateResult struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SouvenirCreateResult) Reset()         { *m = SouvenirCreateResult{} }
func (m *SouvenirCreateResult) String() string { return proto.CompactTextString(m) }
func (*SouvenirCreateResult) ProtoMessage()    {}
func (*SouvenirCreateResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8f236c690c7bcc2, []int{4}
}

func (m *SouvenirCreateResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SouvenirCreateResult.Unmarshal(m, b)
}
func (m *SouvenirCreateResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SouvenirCreateResult.Marshal(b, m, deterministic)
}
func (m *SouvenirCreateResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SouvenirCreateResult.Merge(m, src)
}
func (m *SouvenirCreateResult) XXX_Size() int {
	return xxx_messageInfo_SouvenirCreateResult.Size(m)
}
func (m *SouvenirCreateResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SouvenirCreateResult.DiscardUnknown(m)
}

var xxx_messageInfo_SouvenirCreateResult proto.InternalMessageInfo

func (m *SouvenirCreateResult) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type SouvenirDeleteQuery struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SouvenirDeleteQuery) Reset()         { *m = SouvenirDeleteQuery{} }
func (m *SouvenirDeleteQuery) String() string { return proto.CompactTextString(m) }
func (*SouvenirDeleteQuery) ProtoMessage()    {}
func (*SouvenirDeleteQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8f236c690c7bcc2, []int{5}
}

func (m *SouvenirDeleteQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SouvenirDeleteQuery.Unmarshal(m, b)
}
func (m *SouvenirDeleteQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SouvenirDeleteQuery.Marshal(b, m, deterministic)
}
func (m *SouvenirDeleteQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SouvenirDeleteQuery.Merge(m, src)
}
func (m *SouvenirDeleteQuery) XXX_Size() int {
	return xxx_messageInfo_SouvenirDeleteQuery.Size(m)
}
func (m *SouvenirDeleteQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_SouvenirDeleteQuery.DiscardUnknown(m)
}

var xxx_messageInfo_SouvenirDeleteQuery proto.InternalMessageInfo

func (m *SouvenirDeleteQuery) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SouvenirDeleteResult struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SouvenirDeleteResult) Reset()         { *m = SouvenirDeleteResult{} }
func (m *SouvenirDeleteResult) String() string { return proto.CompactTextString(m) }
func (*SouvenirDeleteResult) ProtoMessage()    {}
func (*SouvenirDeleteResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8f236c690c7bcc2, []int{6}
}

func (m *SouvenirDeleteResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SouvenirDeleteResult.Unmarshal(m, b)
}
func (m *SouvenirDeleteResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SouvenirDeleteResult.Marshal(b, m, deterministic)
}
func (m *SouvenirDeleteResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SouvenirDeleteResult.Merge(m, src)
}
func (m *SouvenirDeleteResult) XXX_Size() int {
	return xxx_messageInfo_SouvenirDeleteResult.Size(m)
}
func (m *SouvenirDeleteResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SouvenirDeleteResult.DiscardUnknown(m)
}

var xxx_messageInfo_SouvenirDeleteResult proto.InternalMessageInfo

func (m *SouvenirDeleteResult) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type SouvenirGetQuery struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SouvenirGetQuery) Reset()         { *m = SouvenirGetQuery{} }
func (m *SouvenirGetQuery) String() string { return proto.CompactTextString(m) }
func (*SouvenirGetQuery) ProtoMessage()    {}
func (*SouvenirGetQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8f236c690c7bcc2, []int{7}
}

func (m *SouvenirGetQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SouvenirGetQuery.Unmarshal(m, b)
}
func (m *SouvenirGetQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SouvenirGetQuery.Marshal(b, m, deterministic)
}
func (m *SouvenirGetQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SouvenirGetQuery.Merge(m, src)
}
func (m *SouvenirGetQuery) XXX_Size() int {
	return xxx_messageInfo_SouvenirGetQuery.Size(m)
}
func (m *SouvenirGetQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_SouvenirGetQuery.DiscardUnknown(m)
}

var xxx_messageInfo_SouvenirGetQuery proto.InternalMessageInfo

func (m *SouvenirGetQuery) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SouvenirGetResult struct {
	Record               *SouvenirRecord `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SouvenirGetResult) Reset()         { *m = SouvenirGetResult{} }
func (m *SouvenirGetResult) String() string { return proto.CompactTextString(m) }
func (*SouvenirGetResult) ProtoMessage()    {}
func (*SouvenirGetResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8f236c690c7bcc2, []int{8}
}

func (m *SouvenirGetResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SouvenirGetResult.Unmarshal(m, b)
}
func (m *SouvenirGetResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SouvenirGetResult.Marshal(b, m, deterministic)
}
func (m *SouvenirGetResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SouvenirGetResult.Merge(m, src)
}
func (m *SouvenirGetResult) XXX_Size() int {
	return xxx_messageInfo_SouvenirGetResult.Size(m)
}
func (m *SouvenirGetResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SouvenirGetResult.DiscardUnknown(m)
}

var xxx_messageInfo_SouvenirGetResult proto.InternalMessageInfo

func (m *SouvenirGetResult) GetRecord() *SouvenirRecord {
	if m != nil {
		return m.Record
	}
	return nil
}

type SouvenirSearchQuery struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	GenreId              uint64   `protobuf:"varint,4,opt,name=genre_id,json=genreId,proto3" json:"genre_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SouvenirSearchQuery) Reset()         { *m = SouvenirSearchQuery{} }
func (m *SouvenirSearchQuery) String() string { return proto.CompactTextString(m) }
func (*SouvenirSearchQuery) ProtoMessage()    {}
func (*SouvenirSearchQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8f236c690c7bcc2, []int{9}
}

func (m *SouvenirSearchQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SouvenirSearchQuery.Unmarshal(m, b)
}
func (m *SouvenirSearchQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SouvenirSearchQuery.Marshal(b, m, deterministic)
}
func (m *SouvenirSearchQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SouvenirSearchQuery.Merge(m, src)
}
func (m *SouvenirSearchQuery) XXX_Size() int {
	return xxx_messageInfo_SouvenirSearchQuery.Size(m)
}
func (m *SouvenirSearchQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_SouvenirSearchQuery.DiscardUnknown(m)
}

var xxx_messageInfo_SouvenirSearchQuery proto.InternalMessageInfo

func (m *SouvenirSearchQuery) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SouvenirSearchQuery) GetGenreId() uint64 {
	if m != nil {
		return m.GenreId
	}
	return 0
}

type SouvenirSearchResult struct {
	Records              []*SouvenirRecord `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SouvenirSearchResult) Reset()         { *m = SouvenirSearchResult{} }
func (m *SouvenirSearchResult) String() string { return proto.CompactTextString(m) }
func (*SouvenirSearchResult) ProtoMessage()    {}
func (*SouvenirSearchResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8f236c690c7bcc2, []int{10}
}

func (m *SouvenirSearchResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SouvenirSearchResult.Unmarshal(m, b)
}
func (m *SouvenirSearchResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SouvenirSearchResult.Marshal(b, m, deterministic)
}
func (m *SouvenirSearchResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SouvenirSearchResult.Merge(m, src)
}
func (m *SouvenirSearchResult) XXX_Size() int {
	return xxx_messageInfo_SouvenirSearchResult.Size(m)
}
func (m *SouvenirSearchResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SouvenirSearchResult.DiscardUnknown(m)
}

var xxx_messageInfo_SouvenirSearchResult proto.InternalMessageInfo

func (m *SouvenirSearchResult) GetRecords() []*SouvenirRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

type SouvenirEmptyQuery struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SouvenirEmptyQuery) Reset()         { *m = SouvenirEmptyQuery{} }
func (m *SouvenirEmptyQuery) String() string { return proto.CompactTextString(m) }
func (*SouvenirEmptyQuery) ProtoMessage()    {}
func (*SouvenirEmptyQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8f236c690c7bcc2, []int{11}
}

func (m *SouvenirEmptyQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SouvenirEmptyQuery.Unmarshal(m, b)
}
func (m *SouvenirEmptyQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SouvenirEmptyQuery.Marshal(b, m, deterministic)
}
func (m *SouvenirEmptyQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SouvenirEmptyQuery.Merge(m, src)
}
func (m *SouvenirEmptyQuery) XXX_Size() int {
	return xxx_messageInfo_SouvenirEmptyQuery.Size(m)
}
func (m *SouvenirEmptyQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_SouvenirEmptyQuery.DiscardUnknown(m)
}

var xxx_messageInfo_SouvenirEmptyQuery proto.InternalMessageInfo

type SouvenirGetGenresResult struct {
	Genres               []*SouvenirGenre `protobuf:"bytes,1,rep,name=genres,proto3" json:"genres,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SouvenirGetGenresResult) Reset()         { *m = SouvenirGetGenresResult{} }
func (m *SouvenirGetGenresResult) String() string { return proto.CompactTextString(m) }
func (*SouvenirGetGenresResult) ProtoMessage()    {}
func (*SouvenirGetGenresResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8f236c690c7bcc2, []int{12}
}

func (m *SouvenirGetGenresResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SouvenirGetGenresResult.Unmarshal(m, b)
}
func (m *SouvenirGetGenresResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SouvenirGetGenresResult.Marshal(b, m, deterministic)
}
func (m *SouvenirGetGenresResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SouvenirGetGenresResult.Merge(m, src)
}
func (m *SouvenirGetGenresResult) XXX_Size() int {
	return xxx_messageInfo_SouvenirGetGenresResult.Size(m)
}
func (m *SouvenirGetGenresResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SouvenirGetGenresResult.DiscardUnknown(m)
}

var xxx_messageInfo_SouvenirGetGenresResult proto.InternalMessageInfo

func (m *SouvenirGetGenresResult) GetGenres() []*SouvenirGenre {
	if m != nil {
		return m.Genres
	}
	return nil
}

func init() {
	proto.RegisterType((*SouvenirGenre)(nil), "souvenir.SouvenirGenre")
	proto.RegisterType((*SouvenirRecord)(nil), "souvenir.SouvenirRecord")
	proto.RegisterType((*SouvenirData)(nil), "souvenir.SouvenirData")
	proto.RegisterType((*SouvenirCreateQuery)(nil), "souvenir.SouvenirCreateQuery")
	proto.RegisterType((*SouvenirCreateResult)(nil), "souvenir.SouvenirCreateResult")
	proto.RegisterType((*SouvenirDeleteQuery)(nil), "souvenir.SouvenirDeleteQuery")
	proto.RegisterType((*SouvenirDeleteResult)(nil), "souvenir.SouvenirDeleteResult")
	proto.RegisterType((*SouvenirGetQuery)(nil), "souvenir.SouvenirGetQuery")
	proto.RegisterType((*SouvenirGetResult)(nil), "souvenir.SouvenirGetResult")
	proto.RegisterType((*SouvenirSearchQuery)(nil), "souvenir.SouvenirSearchQuery")
	proto.RegisterType((*SouvenirSearchResult)(nil), "souvenir.SouvenirSearchResult")
	proto.RegisterType((*SouvenirEmptyQuery)(nil), "souvenir.SouvenirEmptyQuery")
	proto.RegisterType((*SouvenirGetGenresResult)(nil), "souvenir.SouvenirGetGenresResult")
}

func init() { proto.RegisterFile("souvenir/messages.proto", fileDescriptor_a8f236c690c7bcc2) }

var fileDescriptor_a8f236c690c7bcc2 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcf, 0x8b, 0xda, 0x40,
	0x14, 0x26, 0x9a, 0xaa, 0x7d, 0x6d, 0xa5, 0x9d, 0x4a, 0x4d, 0x6f, 0x32, 0x50, 0x90, 0x82, 0x89,
	0xe8, 0xb1, 0xf4, 0xd0, 0x56, 0x29, 0x95, 0x5e, 0x3a, 0xde, 0x7a, 0x59, 0xc6, 0xe4, 0x11, 0x03,
	0x49, 0x26, 0xcc, 0x4c, 0x5c, 0xfc, 0xef, 0x17, 0x27, 0x19, 0x13, 0x57, 0xd9, 0xdd, 0xdb, 0xcc,
	0xbc, 0xef, 0x7d, 0x3f, 0xde, 0x4b, 0x60, 0xac, 0x44, 0x79, 0xc0, 0x3c, 0x91, 0x41, 0x86, 0x4a,
	0xf1, 0x18, 0x95, 0x5f, 0x48, 0xa1, 0x05, 0x19, 0xd8, 0x02, 0x5d, 0xc2, 0xbb, 0x6d, 0x7d, 0xfe,
	0x8d, 0xb9, 0x44, 0x32, 0x84, 0x4e, 0x12, 0x79, 0xce, 0xc4, 0x99, 0xba, 0xac, 0x93, 0x44, 0x84,
	0x80, 0x9b, 0xf3, 0x0c, 0xbd, 0xce, 0xc4, 0x99, 0xbe, 0x66, 0xe6, 0x4c, 0xff, 0xc2, 0xd0, 0x36,
	0x31, 0x0c, 0x85, 0x8c, 0xae, 0xba, 0xbe, 0x82, 0x1b, 0x71, 0xcd, 0x4d, 0xd7, 0x9b, 0xc5, 0x27,
	0xdf, 0xea, 0xf9, 0xb6, 0x6f, 0xc5, 0x35, 0x67, 0x06, 0x43, 0xb7, 0xf0, 0xb6, 0xfd, 0x7a, 0x56,
	0x74, 0x1a, 0x45, 0x32, 0x82, 0x57, 0x85, 0x4c, 0xc2, 0xca, 0x86, 0xcb, 0xaa, 0x0b, 0xf9, 0x0c,
	0x83, 0xf8, 0x64, 0xfa, 0x2e, 0x89, 0xbc, 0xae, 0x29, 0xf4, 0xcd, 0xfd, 0x4f, 0x44, 0x7f, 0xc0,
	0x47, 0x4b, 0xfa, 0x4b, 0x22, 0xd7, 0xf8, 0xaf, 0x44, 0x79, 0x3c, 0xfb, 0x72, 0x5e, 0xe0, 0x6b,
	0x0e, 0xa3, 0x4b, 0x0a, 0x86, 0xaa, 0x4c, 0x35, 0xf1, 0xa0, 0xaf, 0xca, 0x30, 0x44, 0xa5, 0x0c,
	0xcd, 0x80, 0xd9, 0x2b, 0xfd, 0xd2, 0x88, 0xae, 0x30, 0x45, 0x2b, 0xfa, 0x68, 0x38, 0x6d, 0xe2,
	0x0a, 0xf6, 0x2c, 0x31, 0x85, 0xf7, 0xcd, 0x96, 0xf4, 0x6d, 0xd6, 0x35, 0x7c, 0x68, 0x61, 0x6a,
	0xca, 0x39, 0xf4, 0xa4, 0xd9, 0x50, 0x9d, 0xd8, 0xbb, 0x4e, 0x5c, 0x6d, 0x90, 0xd5, 0x38, 0xba,
	0x6a, 0x32, 0x6c, 0x91, 0xcb, 0x70, 0x5f, 0xa9, 0xdd, 0x5a, 0x4a, 0x7b, 0xfc, 0xee, 0xe5, 0xf8,
	0x37, 0x4d, 0xc4, 0x8a, 0xa5, 0xf6, 0xb3, 0x80, 0x7e, 0xa5, 0x73, 0x8a, 0xd8, 0x7d, 0xd2, 0x90,
	0x05, 0xd2, 0x11, 0x10, 0x5b, 0x5a, 0x67, 0x85, 0x3e, 0x1a, 0x43, 0x74, 0x03, 0xe3, 0x56, 0x5c,
	0xf3, 0xed, 0xaa, 0x5a, 0x24, 0x80, 0x9e, 0xf1, 0x61, 0x35, 0xc6, 0xd7, 0x1a, 0x06, 0xcf, 0x6a,
	0xd8, 0xcf, 0xef, 0xff, 0xbf, 0xc5, 0x89, 0xde, 0x97, 0x3b, 0x3f, 0x14, 0x59, 0x50, 0x48, 0x81,
	0xe9, 0x4e, 0xe7, 0x81, 0x0a, 0xf7, 0x42, 0xa4, 0x33, 0x3c, 0xe0, 0x2c, 0xe7, 0x87, 0x24, 0x88,
	0xb9, 0xc6, 0x7b, 0x7e, 0x0c, 0xcc, 0x3f, 0xa4, 0x02, 0x4b, 0xba, 0xeb, 0x99, 0x87, 0xe5, 0x43,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x01, 0x69, 0x9d, 0x4b, 0x6f, 0x03, 0x00, 0x00,
}
