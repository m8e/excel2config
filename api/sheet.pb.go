// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sheet.proto

// package 命名使用 {appid}.{version} 的方式, version 形如 v1, v2 ..

package api

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type LoadExcelReq struct {
	GridKey              string   `protobuf:"bytes,1,opt,name=gridKey,proto3" json:"gridKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadExcelReq) Reset()         { *m = LoadExcelReq{} }
func (m *LoadExcelReq) String() string { return proto.CompactTextString(m) }
func (*LoadExcelReq) ProtoMessage()    {}
func (*LoadExcelReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2edb8d084c32530c, []int{0}
}
func (m *LoadExcelReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoadExcelReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoadExcelReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LoadExcelReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadExcelReq.Merge(m, src)
}
func (m *LoadExcelReq) XXX_Size() int {
	return m.Size()
}
func (m *LoadExcelReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadExcelReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoadExcelReq proto.InternalMessageInfo

type LoadExcelResp struct {
	Jsonstr              string   `protobuf:"bytes,1,opt,name=jsonstr,proto3" json:"jsonstr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadExcelResp) Reset()         { *m = LoadExcelResp{} }
func (m *LoadExcelResp) String() string { return proto.CompactTextString(m) }
func (*LoadExcelResp) ProtoMessage()    {}
func (*LoadExcelResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_2edb8d084c32530c, []int{1}
}
func (m *LoadExcelResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoadExcelResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoadExcelResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LoadExcelResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadExcelResp.Merge(m, src)
}
func (m *LoadExcelResp) XXX_Size() int {
	return m.Size()
}
func (m *LoadExcelResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadExcelResp.DiscardUnknown(m)
}

var xxx_messageInfo_LoadExcelResp proto.InternalMessageInfo

type LoadExcelSheetReq struct {
	GridKey              string   `protobuf:"bytes,1,opt,name=gridKey,proto3" json:"gridKey,omitempty"`
	Indexs               []string `protobuf:"bytes,2,rep,name=indexs,proto3" json:"indexs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadExcelSheetReq) Reset()         { *m = LoadExcelSheetReq{} }
func (m *LoadExcelSheetReq) String() string { return proto.CompactTextString(m) }
func (*LoadExcelSheetReq) ProtoMessage()    {}
func (*LoadExcelSheetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2edb8d084c32530c, []int{2}
}
func (m *LoadExcelSheetReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoadExcelSheetReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoadExcelSheetReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LoadExcelSheetReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadExcelSheetReq.Merge(m, src)
}
func (m *LoadExcelSheetReq) XXX_Size() int {
	return m.Size()
}
func (m *LoadExcelSheetReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadExcelSheetReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoadExcelSheetReq proto.InternalMessageInfo

type LoadExcelSheetResp struct {
	Jsonstr              string   `protobuf:"bytes,1,opt,name=jsonstr,proto3" json:"jsonstr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadExcelSheetResp) Reset()         { *m = LoadExcelSheetResp{} }
func (m *LoadExcelSheetResp) String() string { return proto.CompactTextString(m) }
func (*LoadExcelSheetResp) ProtoMessage()    {}
func (*LoadExcelSheetResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_2edb8d084c32530c, []int{3}
}
func (m *LoadExcelSheetResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoadExcelSheetResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoadExcelSheetResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LoadExcelSheetResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadExcelSheetResp.Merge(m, src)
}
func (m *LoadExcelSheetResp) XXX_Size() int {
	return m.Size()
}
func (m *LoadExcelSheetResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadExcelSheetResp.DiscardUnknown(m)
}

var xxx_messageInfo_LoadExcelSheetResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LoadExcelReq)(nil), "sheet.service.v1.LoadExcelReq")
	proto.RegisterType((*LoadExcelResp)(nil), "sheet.service.v1.LoadExcelResp")
	proto.RegisterType((*LoadExcelSheetReq)(nil), "sheet.service.v1.LoadExcelSheetReq")
	proto.RegisterType((*LoadExcelSheetResp)(nil), "sheet.service.v1.LoadExcelSheetResp")
}

func init() { proto.RegisterFile("sheet.proto", fileDescriptor_2edb8d084c32530c) }

var fileDescriptor_2edb8d084c32530c = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xbb, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0xeb, 0x56, 0x14, 0xd5, 0xbd, 0x08, 0x0c, 0x42, 0x25, 0x42, 0xa6, 0x32, 0x0c, 0x65,
	0xc0, 0x11, 0xf0, 0x06, 0x48, 0x9d, 0x60, 0x0a, 0x1b, 0x5b, 0x2e, 0xc6, 0x35, 0x2a, 0xb1, 0x89,
	0xdd, 0xaa, 0xac, 0xbc, 0x02, 0x0b, 0x8f, 0xd4, 0x11, 0x89, 0x19, 0x09, 0x02, 0x0f, 0x82, 0xe2,
	0x34, 0x55, 0xe8, 0x10, 0xb6, 0x7c, 0x39, 0x9f, 0xff, 0x73, 0x7c, 0x0c, 0xdb, 0x7a, 0xcc, 0x98,
	0xa1, 0x2a, 0x91, 0x46, 0xa2, 0xad, 0x1c, 0x34, 0x4b, 0x66, 0x22, 0x64, 0x74, 0x76, 0xe6, 0x9c,
	0x72, 0x61, 0xc6, 0xd3, 0x80, 0x86, 0xf2, 0xc1, 0xe5, 0x92, 0x4b, 0xd7, 0x8a, 0xc1, 0xf4, 0xce,
	0x92, 0x05, 0xfb, 0x95, 0x07, 0x38, 0x07, 0x5c, 0x4a, 0x3e, 0x61, 0xae, 0xaf, 0x84, 0xeb, 0xc7,
	0xb1, 0x34, 0xbe, 0x11, 0x32, 0xd6, 0x79, 0x95, 0x0c, 0x61, 0xe7, 0x5a, 0xfa, 0xd1, 0x68, 0x1e,
	0xb2, 0x89, 0xc7, 0x1e, 0x51, 0x1f, 0x6e, 0xf2, 0x44, 0x44, 0x57, 0xec, 0xa9, 0x0f, 0x06, 0x60,
	0xd8, 0xf2, 0x0a, 0x24, 0x27, 0xb0, 0x5b, 0x32, 0xb5, 0xca, 0xd4, 0x7b, 0x2d, 0x63, 0x6d, 0x92,
	0x42, 0x5d, 0x22, 0x19, 0xc1, 0xed, 0x95, 0x7a, 0x93, 0x8d, 0x5f, 0x99, 0x8c, 0xf6, 0x60, 0x53,
	0xc4, 0x11, 0x9b, 0xeb, 0x7e, 0x7d, 0xd0, 0x18, 0xb6, 0xbc, 0x25, 0x11, 0x0a, 0xd1, 0x7a, 0x4c,
	0x55, 0xdb, 0xf3, 0x0f, 0x00, 0x37, 0xac, 0x87, 0x02, 0xd8, 0x5a, 0x9d, 0x44, 0x98, 0xae, 0xaf,
	0x90, 0x96, 0xaf, 0xec, 0x1c, 0x56, 0xd6, 0xb5, 0x22, 0xbb, 0xcf, 0xef, 0x3f, 0x2f, 0xf5, 0x1e,
	0xe9, 0xb0, 0xec, 0x97, 0x6b, 0x6d, 0x8d, 0x14, 0xec, 0xfd, 0x9d, 0x0e, 0x1d, 0x55, 0x04, 0x15,
	0x6b, 0x70, 0x8e, 0xff, 0x97, 0xb4, 0x22, 0x3b, 0xb6, 0x65, 0x97, 0xb4, 0x4b, 0x2d, 0x2f, 0xf7,
	0x17, 0x5f, 0xb8, 0xb6, 0x48, 0x31, 0x78, 0x4b, 0x31, 0xf8, 0x4c, 0x31, 0x78, 0xfd, 0xc6, 0xb5,
	0xdb, 0x86, 0xaf, 0x44, 0xd0, 0xb4, 0xaf, 0x79, 0xf1, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x12,
	0xa0, 0x65, 0x3b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SheetClient is the client API for Sheet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SheetClient interface {
	LoadExcel(ctx context.Context, in *LoadExcelReq, opts ...grpc.CallOption) (*LoadExcelResp, error)
	LoadExcelSheet(ctx context.Context, in *LoadExcelSheetReq, opts ...grpc.CallOption) (*LoadExcelSheetResp, error)
}

type sheetClient struct {
	cc *grpc.ClientConn
}

func NewSheetClient(cc *grpc.ClientConn) SheetClient {
	return &sheetClient{cc}
}

func (c *sheetClient) LoadExcel(ctx context.Context, in *LoadExcelReq, opts ...grpc.CallOption) (*LoadExcelResp, error) {
	out := new(LoadExcelResp)
	err := c.cc.Invoke(ctx, "/sheet.service.v1.Sheet/LoadExcel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sheetClient) LoadExcelSheet(ctx context.Context, in *LoadExcelSheetReq, opts ...grpc.CallOption) (*LoadExcelSheetResp, error) {
	out := new(LoadExcelSheetResp)
	err := c.cc.Invoke(ctx, "/sheet.service.v1.Sheet/LoadExcelSheet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SheetServer is the server API for Sheet service.
type SheetServer interface {
	LoadExcel(context.Context, *LoadExcelReq) (*LoadExcelResp, error)
	LoadExcelSheet(context.Context, *LoadExcelSheetReq) (*LoadExcelSheetResp, error)
}

// UnimplementedSheetServer can be embedded to have forward compatible implementations.
type UnimplementedSheetServer struct {
}

func (*UnimplementedSheetServer) LoadExcel(ctx context.Context, req *LoadExcelReq) (*LoadExcelResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadExcel not implemented")
}
func (*UnimplementedSheetServer) LoadExcelSheet(ctx context.Context, req *LoadExcelSheetReq) (*LoadExcelSheetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadExcelSheet not implemented")
}

func RegisterSheetServer(s *grpc.Server, srv SheetServer) {
	s.RegisterService(&_Sheet_serviceDesc, srv)
}

func _Sheet_LoadExcel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadExcelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SheetServer).LoadExcel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sheet.service.v1.Sheet/LoadExcel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SheetServer).LoadExcel(ctx, req.(*LoadExcelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sheet_LoadExcelSheet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadExcelSheetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SheetServer).LoadExcelSheet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sheet.service.v1.Sheet/LoadExcelSheet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SheetServer).LoadExcelSheet(ctx, req.(*LoadExcelSheetReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sheet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sheet.service.v1.Sheet",
	HandlerType: (*SheetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoadExcel",
			Handler:    _Sheet_LoadExcel_Handler,
		},
		{
			MethodName: "LoadExcelSheet",
			Handler:    _Sheet_LoadExcelSheet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sheet.proto",
}

func (m *LoadExcelReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoadExcelReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LoadExcelReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.GridKey) > 0 {
		i -= len(m.GridKey)
		copy(dAtA[i:], m.GridKey)
		i = encodeVarintSheet(dAtA, i, uint64(len(m.GridKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LoadExcelResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoadExcelResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LoadExcelResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Jsonstr) > 0 {
		i -= len(m.Jsonstr)
		copy(dAtA[i:], m.Jsonstr)
		i = encodeVarintSheet(dAtA, i, uint64(len(m.Jsonstr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LoadExcelSheetReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoadExcelSheetReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LoadExcelSheetReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Indexs) > 0 {
		for iNdEx := len(m.Indexs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Indexs[iNdEx])
			copy(dAtA[i:], m.Indexs[iNdEx])
			i = encodeVarintSheet(dAtA, i, uint64(len(m.Indexs[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.GridKey) > 0 {
		i -= len(m.GridKey)
		copy(dAtA[i:], m.GridKey)
		i = encodeVarintSheet(dAtA, i, uint64(len(m.GridKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LoadExcelSheetResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoadExcelSheetResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LoadExcelSheetResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Jsonstr) > 0 {
		i -= len(m.Jsonstr)
		copy(dAtA[i:], m.Jsonstr)
		i = encodeVarintSheet(dAtA, i, uint64(len(m.Jsonstr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSheet(dAtA []byte, offset int, v uint64) int {
	offset -= sovSheet(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LoadExcelReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.GridKey)
	if l > 0 {
		n += 1 + l + sovSheet(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LoadExcelResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Jsonstr)
	if l > 0 {
		n += 1 + l + sovSheet(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LoadExcelSheetReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.GridKey)
	if l > 0 {
		n += 1 + l + sovSheet(uint64(l))
	}
	if len(m.Indexs) > 0 {
		for _, s := range m.Indexs {
			l = len(s)
			n += 1 + l + sovSheet(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LoadExcelSheetResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Jsonstr)
	if l > 0 {
		n += 1 + l + sovSheet(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSheet(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSheet(x uint64) (n int) {
	return sovSheet(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LoadExcelReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSheet
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LoadExcelReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoadExcelReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GridKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSheet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSheet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSheet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GridKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSheet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSheet
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSheet
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LoadExcelResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSheet
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LoadExcelResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoadExcelResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Jsonstr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSheet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSheet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSheet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Jsonstr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSheet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSheet
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSheet
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LoadExcelSheetReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSheet
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LoadExcelSheetReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoadExcelSheetReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GridKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSheet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSheet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSheet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GridKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Indexs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSheet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSheet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSheet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Indexs = append(m.Indexs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSheet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSheet
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSheet
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LoadExcelSheetResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSheet
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LoadExcelSheetResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoadExcelSheetResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Jsonstr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSheet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSheet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSheet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Jsonstr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSheet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSheet
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSheet
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSheet(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSheet
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSheet
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSheet
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSheet
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSheet
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSheet
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSheet        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSheet          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSheet = fmt.Errorf("proto: unexpected end of group")
)
