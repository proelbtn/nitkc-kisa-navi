// Code generated by protoc-gen-go. DO NOT EDIT.
// source: souvenir/service.proto

package souvenir

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

func init() { proto.RegisterFile("souvenir/service.proto", fileDescriptor_1739816bdb7dc59f) }

var fileDescriptor_1739816bdb7dc59f = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x3b, 0x4b, 0xc4, 0x40,
	0x14, 0x85, 0x17, 0x85, 0xb0, 0x4e, 0x39, 0x85, 0x42, 0x7c, 0x80, 0xdb, 0x6f, 0x06, 0xb4, 0x14,
	0x41, 0x7c, 0x10, 0xb7, 0x11, 0x74, 0x3b, 0xbb, 0x49, 0x38, 0x24, 0x81, 0x24, 0x13, 0xe6, 0xde,
	0x44, 0xf2, 0xdf, 0x2d, 0x84, 0xbd, 0x4e, 0x9a, 0xcd, 0x6e, 0x7b, 0xbf, 0x73, 0x3e, 0xce, 0x30,
	0xea, 0x9c, 0x5c, 0x3f, 0xa0, 0xad, 0xbc, 0x21, 0xf8, 0xa1, 0xca, 0x91, 0x74, 0xde, 0xb1, 0xd3,
	0xcb, 0x70, 0x8f, 0x2f, 0xa6, 0x44, 0x03, 0x22, 0x5b, 0x80, 0x24, 0x72, 0xf7, 0x7b, 0xa2, 0x96,
	0xdb, 0x7f, 0xa6, 0x37, 0x2a, 0x7a, 0xf1, 0xb0, 0x0c, 0x7d, 0x9d, 0x84, 0x42, 0x12, 0xa8, 0x90,
	0xcf, 0x1e, 0x7e, 0x8c, 0x6f, 0x0e, 0xe1, 0x2f, 0x50, 0x5f, 0xf3, 0x6a, 0xa1, 0xdf, 0x55, 0xf4,
	0x8a, 0x1a, 0xf3, 0x2a, 0x21, 0xa2, 0x3a, 0x8e, 0x57, 0x0b, 0xfd, 0xa4, 0x4e, 0x53, 0xb0, 0x8e,
	0xf7, 0x73, 0x29, 0x58, 0x1c, 0x97, 0xb3, 0x6c, 0xda, 0xb2, 0x51, 0xd1, 0x16, 0xd6, 0xe7, 0xe5,
	0xdc, 0x16, 0x21, 0x07, 0x9f, 0x25, 0x78, 0x52, 0x7d, 0xa8, 0xb3, 0x14, 0x9c, 0xa2, 0xf5, 0x20,
	0x7d, 0xb5, 0x1f, 0x7f, 0x6b, 0x3a, 0x1e, 0x45, 0x76, 0x3b, 0x3b, 0x4a, 0xaa, 0xc1, 0xf7, 0xfc,
	0xf8, 0xfd, 0x50, 0x54, 0x5c, 0xf6, 0x59, 0x92, 0xbb, 0xc6, 0x74, 0xde, 0xa1, 0xce, 0xb8, 0x35,
	0x94, 0x97, 0xce, 0xd5, 0x6b, 0x0c, 0x58, 0xb7, 0x76, 0xa8, 0x4c, 0x61, 0x19, 0x3f, 0x76, 0x34,
	0xbb, 0x3f, 0x23, 0x13, 0xc4, 0x59, 0xb4, 0x3b, 0xdc, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x98,
	0x8e, 0x92, 0xaf, 0x01, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SouvenirClient is the client API for Souvenir service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SouvenirClient interface {
	Create(ctx context.Context, in *SouvenirCreateQuery, opts ...grpc.CallOption) (*SouvenirCreateResult, error)
	Delete(ctx context.Context, in *SouvenirDeleteQuery, opts ...grpc.CallOption) (*SouvenirDeleteQuery, error)
	Get(ctx context.Context, in *SouvenirGetQuery, opts ...grpc.CallOption) (*SouvenirGetResult, error)
	Search(ctx context.Context, in *SouvenirSearchQuery, opts ...grpc.CallOption) (*SouvenirSearchResult, error)
	GetGenres(ctx context.Context, in *SouvenirEmptyQuery, opts ...grpc.CallOption) (*SouvenirGetGenresResult, error)
}

type souvenirClient struct {
	cc *grpc.ClientConn
}

func NewSouvenirClient(cc *grpc.ClientConn) SouvenirClient {
	return &souvenirClient{cc}
}

func (c *souvenirClient) Create(ctx context.Context, in *SouvenirCreateQuery, opts ...grpc.CallOption) (*SouvenirCreateResult, error) {
	out := new(SouvenirCreateResult)
	err := c.cc.Invoke(ctx, "/souvenir.Souvenir/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *souvenirClient) Delete(ctx context.Context, in *SouvenirDeleteQuery, opts ...grpc.CallOption) (*SouvenirDeleteQuery, error) {
	out := new(SouvenirDeleteQuery)
	err := c.cc.Invoke(ctx, "/souvenir.Souvenir/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *souvenirClient) Get(ctx context.Context, in *SouvenirGetQuery, opts ...grpc.CallOption) (*SouvenirGetResult, error) {
	out := new(SouvenirGetResult)
	err := c.cc.Invoke(ctx, "/souvenir.Souvenir/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *souvenirClient) Search(ctx context.Context, in *SouvenirSearchQuery, opts ...grpc.CallOption) (*SouvenirSearchResult, error) {
	out := new(SouvenirSearchResult)
	err := c.cc.Invoke(ctx, "/souvenir.Souvenir/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *souvenirClient) GetGenres(ctx context.Context, in *SouvenirEmptyQuery, opts ...grpc.CallOption) (*SouvenirGetGenresResult, error) {
	out := new(SouvenirGetGenresResult)
	err := c.cc.Invoke(ctx, "/souvenir.Souvenir/GetGenres", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SouvenirServer is the server API for Souvenir service.
type SouvenirServer interface {
	Create(context.Context, *SouvenirCreateQuery) (*SouvenirCreateResult, error)
	Delete(context.Context, *SouvenirDeleteQuery) (*SouvenirDeleteQuery, error)
	Get(context.Context, *SouvenirGetQuery) (*SouvenirGetResult, error)
	Search(context.Context, *SouvenirSearchQuery) (*SouvenirSearchResult, error)
	GetGenres(context.Context, *SouvenirEmptyQuery) (*SouvenirGetGenresResult, error)
}

func RegisterSouvenirServer(s *grpc.Server, srv SouvenirServer) {
	s.RegisterService(&_Souvenir_serviceDesc, srv)
}

func _Souvenir_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SouvenirCreateQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SouvenirServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/souvenir.Souvenir/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SouvenirServer).Create(ctx, req.(*SouvenirCreateQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Souvenir_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SouvenirDeleteQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SouvenirServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/souvenir.Souvenir/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SouvenirServer).Delete(ctx, req.(*SouvenirDeleteQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Souvenir_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SouvenirGetQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SouvenirServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/souvenir.Souvenir/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SouvenirServer).Get(ctx, req.(*SouvenirGetQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Souvenir_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SouvenirSearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SouvenirServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/souvenir.Souvenir/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SouvenirServer).Search(ctx, req.(*SouvenirSearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Souvenir_GetGenres_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SouvenirEmptyQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SouvenirServer).GetGenres(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/souvenir.Souvenir/GetGenres",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SouvenirServer).GetGenres(ctx, req.(*SouvenirEmptyQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _Souvenir_serviceDesc = grpc.ServiceDesc{
	ServiceName: "souvenir.Souvenir",
	HandlerType: (*SouvenirServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Souvenir_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Souvenir_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Souvenir_Get_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Souvenir_Search_Handler,
		},
		{
			MethodName: "GetGenres",
			Handler:    _Souvenir_GetGenres_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "souvenir/service.proto",
}