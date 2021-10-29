// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package types

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ACEIApplicationClient is the client API for ACEIApplication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ACEIApplicationClient interface {
	Echo(ctx context.Context, in *RequestEcho, opts ...grpc.CallOption) (*ResponseEcho, error)
	Flush(ctx context.Context, in *RequestFlush, opts ...grpc.CallOption) (*ResponseFlush, error)
	Info(ctx context.Context, in *RequestInfo, opts ...grpc.CallOption) (*ResponseInfo, error)
	DeliverTx(ctx context.Context, in *RequestDeliverTx, opts ...grpc.CallOption) (*ResponseDeliverTx, error)
	CheckTx(ctx context.Context, in *RequestCheckTx, opts ...grpc.CallOption) (*ResponseCheckTx, error)
	Query(ctx context.Context, in *RequestQuery, opts ...grpc.CallOption) (*ResponseQuery, error)
	Commit(ctx context.Context, in *RequestCommit, opts ...grpc.CallOption) (*ResponseCommit, error)
	InitLedger(ctx context.Context, in *RequestInitLedger, opts ...grpc.CallOption) (*ResponseInitLedger, error)
	BeginBlock(ctx context.Context, in *RequestBeginBlock, opts ...grpc.CallOption) (*ResponseBeginBlock, error)
	EndBlock(ctx context.Context, in *RequestEndBlock, opts ...grpc.CallOption) (*ResponseEndBlock, error)
	ListSnapshots(ctx context.Context, in *RequestListSnapshots, opts ...grpc.CallOption) (*ResponseListSnapshots, error)
	OfferSnapshot(ctx context.Context, in *RequestOfferSnapshot, opts ...grpc.CallOption) (*ResponseOfferSnapshot, error)
	LoadSnapshotChunk(ctx context.Context, in *RequestLoadSnapshotChunk, opts ...grpc.CallOption) (*ResponseLoadSnapshotChunk, error)
	ApplySnapshotChunk(ctx context.Context, in *RequestApplySnapshotChunk, opts ...grpc.CallOption) (*ResponseApplySnapshotChunk, error)
}

type aCEIApplicationClient struct {
	cc grpc.ClientConnInterface
}

func NewACEIApplicationClient(cc grpc.ClientConnInterface) ACEIApplicationClient {
	return &aCEIApplicationClient{cc}
}

func (c *aCEIApplicationClient) Echo(ctx context.Context, in *RequestEcho, opts ...grpc.CallOption) (*ResponseEcho, error) {
	out := new(ResponseEcho)
	err := c.cc.Invoke(ctx, "/daotl.acei.ACEIApplication/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCEIApplicationClient) Flush(ctx context.Context, in *RequestFlush, opts ...grpc.CallOption) (*ResponseFlush, error) {
	out := new(ResponseFlush)
	err := c.cc.Invoke(ctx, "/daotl.acei.ACEIApplication/Flush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCEIApplicationClient) Info(ctx context.Context, in *RequestInfo, opts ...grpc.CallOption) (*ResponseInfo, error) {
	out := new(ResponseInfo)
	err := c.cc.Invoke(ctx, "/daotl.acei.ACEIApplication/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCEIApplicationClient) DeliverTx(ctx context.Context, in *RequestDeliverTx, opts ...grpc.CallOption) (*ResponseDeliverTx, error) {
	out := new(ResponseDeliverTx)
	err := c.cc.Invoke(ctx, "/daotl.acei.ACEIApplication/DeliverTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCEIApplicationClient) CheckTx(ctx context.Context, in *RequestCheckTx, opts ...grpc.CallOption) (*ResponseCheckTx, error) {
	out := new(ResponseCheckTx)
	err := c.cc.Invoke(ctx, "/daotl.acei.ACEIApplication/CheckTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCEIApplicationClient) Query(ctx context.Context, in *RequestQuery, opts ...grpc.CallOption) (*ResponseQuery, error) {
	out := new(ResponseQuery)
	err := c.cc.Invoke(ctx, "/daotl.acei.ACEIApplication/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCEIApplicationClient) Commit(ctx context.Context, in *RequestCommit, opts ...grpc.CallOption) (*ResponseCommit, error) {
	out := new(ResponseCommit)
	err := c.cc.Invoke(ctx, "/daotl.acei.ACEIApplication/Commit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCEIApplicationClient) InitLedger(ctx context.Context, in *RequestInitLedger, opts ...grpc.CallOption) (*ResponseInitLedger, error) {
	out := new(ResponseInitLedger)
	err := c.cc.Invoke(ctx, "/daotl.acei.ACEIApplication/InitLedger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCEIApplicationClient) BeginBlock(ctx context.Context, in *RequestBeginBlock, opts ...grpc.CallOption) (*ResponseBeginBlock, error) {
	out := new(ResponseBeginBlock)
	err := c.cc.Invoke(ctx, "/daotl.acei.ACEIApplication/BeginBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCEIApplicationClient) EndBlock(ctx context.Context, in *RequestEndBlock, opts ...grpc.CallOption) (*ResponseEndBlock, error) {
	out := new(ResponseEndBlock)
	err := c.cc.Invoke(ctx, "/daotl.acei.ACEIApplication/EndBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCEIApplicationClient) ListSnapshots(ctx context.Context, in *RequestListSnapshots, opts ...grpc.CallOption) (*ResponseListSnapshots, error) {
	out := new(ResponseListSnapshots)
	err := c.cc.Invoke(ctx, "/daotl.acei.ACEIApplication/ListSnapshots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCEIApplicationClient) OfferSnapshot(ctx context.Context, in *RequestOfferSnapshot, opts ...grpc.CallOption) (*ResponseOfferSnapshot, error) {
	out := new(ResponseOfferSnapshot)
	err := c.cc.Invoke(ctx, "/daotl.acei.ACEIApplication/OfferSnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCEIApplicationClient) LoadSnapshotChunk(ctx context.Context, in *RequestLoadSnapshotChunk, opts ...grpc.CallOption) (*ResponseLoadSnapshotChunk, error) {
	out := new(ResponseLoadSnapshotChunk)
	err := c.cc.Invoke(ctx, "/daotl.acei.ACEIApplication/LoadSnapshotChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCEIApplicationClient) ApplySnapshotChunk(ctx context.Context, in *RequestApplySnapshotChunk, opts ...grpc.CallOption) (*ResponseApplySnapshotChunk, error) {
	out := new(ResponseApplySnapshotChunk)
	err := c.cc.Invoke(ctx, "/daotl.acei.ACEIApplication/ApplySnapshotChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ACEIApplicationServer is the server API for ACEIApplication service.
// All implementations must embed UnimplementedACEIApplicationServer
// for forward compatibility
type ACEIApplicationServer interface {
	Echo(context.Context, *RequestEcho) (*ResponseEcho, error)
	Flush(context.Context, *RequestFlush) (*ResponseFlush, error)
	Info(context.Context, *RequestInfo) (*ResponseInfo, error)
	DeliverTx(context.Context, *RequestDeliverTx) (*ResponseDeliverTx, error)
	CheckTx(context.Context, *RequestCheckTx) (*ResponseCheckTx, error)
	Query(context.Context, *RequestQuery) (*ResponseQuery, error)
	Commit(context.Context, *RequestCommit) (*ResponseCommit, error)
	InitLedger(context.Context, *RequestInitLedger) (*ResponseInitLedger, error)
	BeginBlock(context.Context, *RequestBeginBlock) (*ResponseBeginBlock, error)
	EndBlock(context.Context, *RequestEndBlock) (*ResponseEndBlock, error)
	ListSnapshots(context.Context, *RequestListSnapshots) (*ResponseListSnapshots, error)
	OfferSnapshot(context.Context, *RequestOfferSnapshot) (*ResponseOfferSnapshot, error)
	LoadSnapshotChunk(context.Context, *RequestLoadSnapshotChunk) (*ResponseLoadSnapshotChunk, error)
	ApplySnapshotChunk(context.Context, *RequestApplySnapshotChunk) (*ResponseApplySnapshotChunk, error)
	mustEmbedUnimplementedACEIApplicationServer()
}

// UnimplementedACEIApplicationServer must be embedded to have forward compatible implementations.
type UnimplementedACEIApplicationServer struct {
}

func (UnimplementedACEIApplicationServer) Echo(context.Context, *RequestEcho) (*ResponseEcho, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedACEIApplicationServer) Flush(context.Context, *RequestFlush) (*ResponseFlush, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Flush not implemented")
}
func (UnimplementedACEIApplicationServer) Info(context.Context, *RequestInfo) (*ResponseInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedACEIApplicationServer) DeliverTx(context.Context, *RequestDeliverTx) (*ResponseDeliverTx, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeliverTx not implemented")
}
func (UnimplementedACEIApplicationServer) CheckTx(context.Context, *RequestCheckTx) (*ResponseCheckTx, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckTx not implemented")
}
func (UnimplementedACEIApplicationServer) Query(context.Context, *RequestQuery) (*ResponseQuery, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedACEIApplicationServer) Commit(context.Context, *RequestCommit) (*ResponseCommit, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Commit not implemented")
}
func (UnimplementedACEIApplicationServer) InitLedger(context.Context, *RequestInitLedger) (*ResponseInitLedger, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitLedger not implemented")
}
func (UnimplementedACEIApplicationServer) BeginBlock(context.Context, *RequestBeginBlock) (*ResponseBeginBlock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BeginBlock not implemented")
}
func (UnimplementedACEIApplicationServer) EndBlock(context.Context, *RequestEndBlock) (*ResponseEndBlock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndBlock not implemented")
}
func (UnimplementedACEIApplicationServer) ListSnapshots(context.Context, *RequestListSnapshots) (*ResponseListSnapshots, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSnapshots not implemented")
}
func (UnimplementedACEIApplicationServer) OfferSnapshot(context.Context, *RequestOfferSnapshot) (*ResponseOfferSnapshot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OfferSnapshot not implemented")
}
func (UnimplementedACEIApplicationServer) LoadSnapshotChunk(context.Context, *RequestLoadSnapshotChunk) (*ResponseLoadSnapshotChunk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadSnapshotChunk not implemented")
}
func (UnimplementedACEIApplicationServer) ApplySnapshotChunk(context.Context, *RequestApplySnapshotChunk) (*ResponseApplySnapshotChunk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplySnapshotChunk not implemented")
}
func (UnimplementedACEIApplicationServer) mustEmbedUnimplementedACEIApplicationServer() {}

// UnsafeACEIApplicationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ACEIApplicationServer will
// result in compilation errors.
type UnsafeACEIApplicationServer interface {
	mustEmbedUnimplementedACEIApplicationServer()
}

func RegisterACEIApplicationServer(s grpc.ServiceRegistrar, srv ACEIApplicationServer) {
	s.RegisterService(&ACEIApplication_ServiceDesc, srv)
}

func _ACEIApplication_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestEcho)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACEIApplicationServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daotl.acei.ACEIApplication/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACEIApplicationServer).Echo(ctx, req.(*RequestEcho))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACEIApplication_Flush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestFlush)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACEIApplicationServer).Flush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daotl.acei.ACEIApplication/Flush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACEIApplicationServer).Flush(ctx, req.(*RequestFlush))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACEIApplication_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACEIApplicationServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daotl.acei.ACEIApplication/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACEIApplicationServer).Info(ctx, req.(*RequestInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACEIApplication_DeliverTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDeliverTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACEIApplicationServer).DeliverTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daotl.acei.ACEIApplication/DeliverTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACEIApplicationServer).DeliverTx(ctx, req.(*RequestDeliverTx))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACEIApplication_CheckTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCheckTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACEIApplicationServer).CheckTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daotl.acei.ACEIApplication/CheckTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACEIApplicationServer).CheckTx(ctx, req.(*RequestCheckTx))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACEIApplication_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACEIApplicationServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daotl.acei.ACEIApplication/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACEIApplicationServer).Query(ctx, req.(*RequestQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACEIApplication_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCommit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACEIApplicationServer).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daotl.acei.ACEIApplication/Commit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACEIApplicationServer).Commit(ctx, req.(*RequestCommit))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACEIApplication_InitLedger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestInitLedger)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACEIApplicationServer).InitLedger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daotl.acei.ACEIApplication/InitLedger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACEIApplicationServer).InitLedger(ctx, req.(*RequestInitLedger))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACEIApplication_BeginBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestBeginBlock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACEIApplicationServer).BeginBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daotl.acei.ACEIApplication/BeginBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACEIApplicationServer).BeginBlock(ctx, req.(*RequestBeginBlock))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACEIApplication_EndBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestEndBlock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACEIApplicationServer).EndBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daotl.acei.ACEIApplication/EndBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACEIApplicationServer).EndBlock(ctx, req.(*RequestEndBlock))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACEIApplication_ListSnapshots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestListSnapshots)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACEIApplicationServer).ListSnapshots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daotl.acei.ACEIApplication/ListSnapshots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACEIApplicationServer).ListSnapshots(ctx, req.(*RequestListSnapshots))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACEIApplication_OfferSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestOfferSnapshot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACEIApplicationServer).OfferSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daotl.acei.ACEIApplication/OfferSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACEIApplicationServer).OfferSnapshot(ctx, req.(*RequestOfferSnapshot))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACEIApplication_LoadSnapshotChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestLoadSnapshotChunk)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACEIApplicationServer).LoadSnapshotChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daotl.acei.ACEIApplication/LoadSnapshotChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACEIApplicationServer).LoadSnapshotChunk(ctx, req.(*RequestLoadSnapshotChunk))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACEIApplication_ApplySnapshotChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestApplySnapshotChunk)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACEIApplicationServer).ApplySnapshotChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daotl.acei.ACEIApplication/ApplySnapshotChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACEIApplicationServer).ApplySnapshotChunk(ctx, req.(*RequestApplySnapshotChunk))
	}
	return interceptor(ctx, in, info, handler)
}

// ACEIApplication_ServiceDesc is the grpc.ServiceDesc for ACEIApplication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ACEIApplication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "daotl.acei.ACEIApplication",
	HandlerType: (*ACEIApplicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _ACEIApplication_Echo_Handler,
		},
		{
			MethodName: "Flush",
			Handler:    _ACEIApplication_Flush_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _ACEIApplication_Info_Handler,
		},
		{
			MethodName: "DeliverTx",
			Handler:    _ACEIApplication_DeliverTx_Handler,
		},
		{
			MethodName: "CheckTx",
			Handler:    _ACEIApplication_CheckTx_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _ACEIApplication_Query_Handler,
		},
		{
			MethodName: "Commit",
			Handler:    _ACEIApplication_Commit_Handler,
		},
		{
			MethodName: "InitLedger",
			Handler:    _ACEIApplication_InitLedger_Handler,
		},
		{
			MethodName: "BeginBlock",
			Handler:    _ACEIApplication_BeginBlock_Handler,
		},
		{
			MethodName: "EndBlock",
			Handler:    _ACEIApplication_EndBlock_Handler,
		},
		{
			MethodName: "ListSnapshots",
			Handler:    _ACEIApplication_ListSnapshots_Handler,
		},
		{
			MethodName: "OfferSnapshot",
			Handler:    _ACEIApplication_OfferSnapshot_Handler,
		},
		{
			MethodName: "LoadSnapshotChunk",
			Handler:    _ACEIApplication_LoadSnapshotChunk_Handler,
		},
		{
			MethodName: "ApplySnapshotChunk",
			Handler:    _ACEIApplication_ApplySnapshotChunk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "daotl/acei/types.proto",
}
