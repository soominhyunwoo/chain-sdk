// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/transfer/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	query "github.com/soominhyunwoo/chain-sdk/types/query"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryDenomTraceRequest is the request type for the Query/DenomTrace RPC
// method
type QueryDenomTraceRequest struct {
	// hash (in hex format) of the denomination trace information.
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *QueryDenomTraceRequest) Reset()         { *m = QueryDenomTraceRequest{} }
func (m *QueryDenomTraceRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDenomTraceRequest) ProtoMessage()    {}
func (*QueryDenomTraceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a638e2800a01538c, []int{0}
}
func (m *QueryDenomTraceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomTraceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomTraceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomTraceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomTraceRequest.Merge(m, src)
}
func (m *QueryDenomTraceRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomTraceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomTraceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomTraceRequest proto.InternalMessageInfo

func (m *QueryDenomTraceRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

// QueryDenomTraceResponse is the response type for the Query/DenomTrace RPC
// method.
type QueryDenomTraceResponse struct {
	// denom_trace returns the requested denomination trace information.
	DenomTrace *DenomTrace `protobuf:"bytes,1,opt,name=denom_trace,json=denomTrace,proto3" json:"denom_trace,omitempty"`
}

func (m *QueryDenomTraceResponse) Reset()         { *m = QueryDenomTraceResponse{} }
func (m *QueryDenomTraceResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDenomTraceResponse) ProtoMessage()    {}
func (*QueryDenomTraceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a638e2800a01538c, []int{1}
}
func (m *QueryDenomTraceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomTraceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomTraceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomTraceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomTraceResponse.Merge(m, src)
}
func (m *QueryDenomTraceResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomTraceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomTraceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomTraceResponse proto.InternalMessageInfo

func (m *QueryDenomTraceResponse) GetDenomTrace() *DenomTrace {
	if m != nil {
		return m.DenomTrace
	}
	return nil
}

// QueryConnectionsRequest is the request type for the Query/DenomTraces RPC
// method
type QueryDenomTracesRequest struct {
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryDenomTracesRequest) Reset()         { *m = QueryDenomTracesRequest{} }
func (m *QueryDenomTracesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDenomTracesRequest) ProtoMessage()    {}
func (*QueryDenomTracesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a638e2800a01538c, []int{2}
}
func (m *QueryDenomTracesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomTracesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomTracesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomTracesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomTracesRequest.Merge(m, src)
}
func (m *QueryDenomTracesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomTracesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomTracesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomTracesRequest proto.InternalMessageInfo

func (m *QueryDenomTracesRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryConnectionsResponse is the response type for the Query/DenomTraces RPC
// method.
type QueryDenomTracesResponse struct {
	// denom_traces returns all denominations trace information.
	DenomTraces Traces `protobuf:"bytes,1,rep,name=denom_traces,json=denomTraces,proto3,castrepeated=Traces" json:"denom_traces"`
	// pagination defines the pagination in the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryDenomTracesResponse) Reset()         { *m = QueryDenomTracesResponse{} }
func (m *QueryDenomTracesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDenomTracesResponse) ProtoMessage()    {}
func (*QueryDenomTracesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a638e2800a01538c, []int{3}
}
func (m *QueryDenomTracesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomTracesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomTracesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomTracesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomTracesResponse.Merge(m, src)
}
func (m *QueryDenomTracesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomTracesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomTracesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomTracesResponse proto.InternalMessageInfo

func (m *QueryDenomTracesResponse) GetDenomTraces() Traces {
	if m != nil {
		return m.DenomTraces
	}
	return nil
}

func (m *QueryDenomTracesResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a638e2800a01538c, []int{4}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params defines the parameters of the module.
	Params *Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a638e2800a01538c, []int{5}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryDenomTraceRequest)(nil), "ibc.applications.transfer.v1.QueryDenomTraceRequest")
	proto.RegisterType((*QueryDenomTraceResponse)(nil), "ibc.applications.transfer.v1.QueryDenomTraceResponse")
	proto.RegisterType((*QueryDenomTracesRequest)(nil), "ibc.applications.transfer.v1.QueryDenomTracesRequest")
	proto.RegisterType((*QueryDenomTracesResponse)(nil), "ibc.applications.transfer.v1.QueryDenomTracesResponse")
	proto.RegisterType((*QueryParamsRequest)(nil), "ibc.applications.transfer.v1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "ibc.applications.transfer.v1.QueryParamsResponse")
}

func init() {
	proto.RegisterFile("ibc/applications/transfer/v1/query.proto", fileDescriptor_a638e2800a01538c)
}

var fileDescriptor_a638e2800a01538c = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x3f, 0x6f, 0xd3, 0x40,
	0x14, 0xcf, 0x95, 0x12, 0x89, 0x17, 0xc4, 0x70, 0x54, 0x10, 0x59, 0x95, 0x5b, 0x59, 0x08, 0x02,
	0x85, 0x3b, 0x5c, 0xa0, 0x30, 0xa0, 0x0e, 0x15, 0x02, 0xb1, 0x95, 0xc0, 0x80, 0x60, 0x40, 0x67,
	0xe7, 0x70, 0x2c, 0x1a, 0x9f, 0xeb, 0xbb, 0x44, 0x54, 0x88, 0x85, 0x4f, 0x80, 0xc4, 0x8e, 0x98,
	0xd9, 0x19, 0xd8, 0x18, 0x3b, 0x56, 0x62, 0x61, 0x02, 0x94, 0xf0, 0x41, 0x90, 0xef, 0xce, 0x8d,
	0xa3, 0x20, 0x13, 0x4f, 0x39, 0x5d, 0xde, 0xef, 0xfd, 0xfe, 0xbc, 0xe7, 0x83, 0x4e, 0x1c, 0x84,
	0x94, 0xa5, 0xe9, 0x5e, 0x1c, 0x32, 0x15, 0x8b, 0x44, 0x52, 0x95, 0xb1, 0x44, 0xbe, 0xe4, 0x19,
	0x1d, 0xf9, 0x74, 0x7f, 0xc8, 0xb3, 0x03, 0x92, 0x66, 0x42, 0x09, 0xbc, 0x1a, 0x07, 0x21, 0x29,
	0x57, 0x92, 0xa2, 0x92, 0x8c, 0x7c, 0x67, 0x25, 0x12, 0x91, 0xd0, 0x85, 0x34, 0x3f, 0x19, 0x8c,
	0x73, 0x25, 0x14, 0x72, 0x20, 0x24, 0x0d, 0x98, 0xe4, 0xa6, 0x19, 0x1d, 0xf9, 0x01, 0x57, 0xcc,
	0xa7, 0x29, 0x8b, 0xe2, 0x44, 0x37, 0xb2, 0xb5, 0x1b, 0x95, 0x4a, 0x8e, 0xb9, 0x4c, 0xf1, 0x6a,
	0x24, 0x44, 0xb4, 0xc7, 0x29, 0x4b, 0x63, 0xca, 0x92, 0x44, 0x28, 0x2b, 0x49, 0xff, 0xeb, 0x5d,
	0x85, 0x73, 0x8f, 0x72, 0xb2, 0x7b, 0x3c, 0x11, 0x83, 0x27, 0x19, 0x0b, 0x79, 0x97, 0xef, 0x0f,
	0xb9, 0x54, 0x18, 0xc3, 0x72, 0x9f, 0xc9, 0x7e, 0x1b, 0xad, 0xa3, 0xce, 0xa9, 0xae, 0x3e, 0x7b,
	0x3d, 0x38, 0x3f, 0x57, 0x2d, 0x53, 0x91, 0x48, 0x8e, 0x1f, 0x42, 0xab, 0x97, 0xdf, 0xbe, 0x50,
	0xf9, 0xb5, 0x46, 0xb5, 0x36, 0x3b, 0xa4, 0x2a, 0x09, 0x52, 0x6a, 0x03, 0xbd, 0xe3, 0xb3, 0xc7,
	0xe6, 0x58, 0x64, 0x21, 0xea, 0x3e, 0xc0, 0x34, 0x0d, 0x4b, 0x72, 0x91, 0x98, 0xe8, 0x48, 0x1e,
	0x1d, 0x31, 0x73, 0xb0, 0xd1, 0x91, 0x5d, 0x16, 0x15, 0x86, 0xba, 0x25, 0xa4, 0xf7, 0x0d, 0x41,
	0x7b, 0x9e, 0xc3, 0x5a, 0x79, 0x0e, 0xa7, 0x4b, 0x56, 0x64, 0x1b, 0xad, 0x9f, 0xa8, 0xe3, 0x65,
	0xe7, 0xcc, 0xe1, 0xcf, 0xb5, 0xc6, 0xe7, 0x5f, 0x6b, 0x4d, 0xdb, 0xb7, 0x35, 0xf5, 0x26, 0xf1,
	0x83, 0x19, 0x07, 0x4b, 0xda, 0xc1, 0xa5, 0xff, 0x3a, 0x30, 0xca, 0x66, 0x2c, 0xac, 0x00, 0xd6,
	0x0e, 0x76, 0x59, 0xc6, 0x06, 0x45, 0x40, 0xde, 0x63, 0x38, 0x3b, 0x73, 0x6b, 0x2d, 0xdd, 0x85,
	0x66, 0xaa, 0x6f, 0x6c, 0x66, 0x17, 0xaa, 0xcd, 0x58, 0xb4, 0xc5, 0x6c, 0x7e, 0x5c, 0x86, 0x93,
	0xba, 0x2b, 0xfe, 0x8a, 0x00, 0xa6, 0x4e, 0xf1, 0xcd, 0xea, 0x36, 0xff, 0xde, 0x2c, 0xe7, 0x56,
	0x4d, 0x94, 0xf1, 0xe0, 0x6d, 0xbf, 0xfb, 0xfe, 0xe7, 0xc3, 0xd2, 0x1d, 0xbc, 0x45, 0xab, 0xd6,
	0xdf, 0x7c, 0x32, 0xe5, 0xf9, 0xd1, 0x37, 0xf9, 0xee, 0xbe, 0xc5, 0x5f, 0x10, 0xb4, 0x4a, 0xe3,
	0xc6, 0xf5, 0x64, 0x14, 0x09, 0x3b, 0x5b, 0x75, 0x61, 0x56, 0xfe, 0x6d, 0x2d, 0xdf, 0xc7, 0xb4,
	0xa6, 0x7c, 0xfc, 0x09, 0x41, 0xd3, 0x0c, 0x04, 0x5f, 0x5f, 0x80, 0x7b, 0x66, 0x1f, 0x1c, 0xbf,
	0x06, 0xc2, 0x0a, 0xf5, 0xb5, 0xd0, 0x0d, 0x7c, 0x79, 0x01, 0xa1, 0x66, 0x41, 0x76, 0x9e, 0x1e,
	0x8e, 0x5d, 0x74, 0x34, 0x76, 0xd1, 0xef, 0xb1, 0x8b, 0xde, 0x4f, 0xdc, 0xc6, 0xd1, 0xc4, 0x6d,
	0xfc, 0x98, 0xb8, 0x8d, 0x67, 0xdb, 0x51, 0xac, 0xfa, 0xc3, 0x80, 0x84, 0x62, 0x40, 0xed, 0x0b,
	0x67, 0x7e, 0xae, 0xc9, 0xde, 0x2b, 0xfa, 0xba, 0x82, 0x42, 0x1d, 0xa4, 0x5c, 0x06, 0x4d, 0xfd,
	0x4c, 0xdd, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x7f, 0xfe, 0xbd, 0x7d, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// DenomTrace queries a denomination trace information.
	DenomTrace(ctx context.Context, in *QueryDenomTraceRequest, opts ...grpc.CallOption) (*QueryDenomTraceResponse, error)
	// DenomTraces queries all denomination traces.
	DenomTraces(ctx context.Context, in *QueryDenomTracesRequest, opts ...grpc.CallOption) (*QueryDenomTracesResponse, error)
	// Params queries all parameters of the ibc-transfer module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) DenomTrace(ctx context.Context, in *QueryDenomTraceRequest, opts ...grpc.CallOption) (*QueryDenomTraceResponse, error) {
	out := new(QueryDenomTraceResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.transfer.v1.Query/DenomTrace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DenomTraces(ctx context.Context, in *QueryDenomTracesRequest, opts ...grpc.CallOption) (*QueryDenomTracesResponse, error) {
	out := new(QueryDenomTracesResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.transfer.v1.Query/DenomTraces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.transfer.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// DenomTrace queries a denomination trace information.
	DenomTrace(context.Context, *QueryDenomTraceRequest) (*QueryDenomTraceResponse, error)
	// DenomTraces queries all denomination traces.
	DenomTraces(context.Context, *QueryDenomTracesRequest) (*QueryDenomTracesResponse, error)
	// Params queries all parameters of the ibc-transfer module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) DenomTrace(ctx context.Context, req *QueryDenomTraceRequest) (*QueryDenomTraceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomTrace not implemented")
}
func (*UnimplementedQueryServer) DenomTraces(ctx context.Context, req *QueryDenomTracesRequest) (*QueryDenomTracesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomTraces not implemented")
}
func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_DenomTrace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDenomTraceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomTrace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.transfer.v1.Query/DenomTrace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomTrace(ctx, req.(*QueryDenomTraceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DenomTraces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDenomTracesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomTraces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.transfer.v1.Query/DenomTraces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomTraces(ctx, req.(*QueryDenomTracesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.transfer.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.applications.transfer.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DenomTrace",
			Handler:    _Query_DenomTrace_Handler,
		},
		{
			MethodName: "DenomTraces",
			Handler:    _Query_DenomTraces_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/applications/transfer/v1/query.proto",
}

func (m *QueryDenomTraceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomTraceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomTraceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDenomTraceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomTraceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomTraceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DenomTrace != nil {
		{
			size, err := m.DenomTrace.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDenomTracesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomTracesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomTracesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDenomTracesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomTracesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomTracesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.DenomTraces) > 0 {
		for iNdEx := len(m.DenomTraces) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DenomTraces[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Params != nil {
		{
			size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryDenomTraceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDenomTraceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DenomTrace != nil {
		l = m.DenomTrace.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDenomTracesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDenomTracesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DenomTraces) > 0 {
		for _, e := range m.DenomTraces {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Params != nil {
		l = m.Params.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryDenomTraceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryDenomTraceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomTraceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryDenomTraceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryDenomTraceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomTraceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomTrace", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DenomTrace == nil {
				m.DenomTrace = &DenomTrace{}
			}
			if err := m.DenomTrace.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryDenomTracesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryDenomTracesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomTracesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryDenomTracesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryDenomTracesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomTracesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomTraces", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomTraces = append(m.DenomTraces, DenomTrace{})
			if err := m.DenomTraces[len(m.DenomTraces)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Params == nil {
				m.Params = &Params{}
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
