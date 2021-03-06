// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// SphynxClient is the client API for Sphynx service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SphynxClient interface {
	CanCompute(ctx context.Context, in *CanComputeRequest, opts ...grpc.CallOption) (*CanComputeReply, error)
	Compute(ctx context.Context, in *ComputeRequest, opts ...grpc.CallOption) (*ComputeReply, error)
	GetScalar(ctx context.Context, in *GetScalarRequest, opts ...grpc.CallOption) (*GetScalarReply, error)
	WriteToUnorderedDisk(ctx context.Context, in *WriteToUnorderedDiskRequest, opts ...grpc.CallOption) (*WriteToUnorderedDiskReply, error)
	ReadFromUnorderedDisk(ctx context.Context, in *ReadFromUnorderedDiskRequest, opts ...grpc.CallOption) (*ReadFromUnorderedDiskReply, error)
	WriteToOrderedDisk(ctx context.Context, in *WriteToOrderedDiskRequest, opts ...grpc.CallOption) (*WriteToOrderedDiskReply, error)
	ReadFromOrderedSphynxDisk(ctx context.Context, in *ReadFromOrderedSphynxDiskRequest, opts ...grpc.CallOption) (*ReadFromOrderedSphynxDiskReply, error)
	HasOnOrderedSphynxDisk(ctx context.Context, in *HasOnOrderedSphynxDiskRequest, opts ...grpc.CallOption) (*HasOnOrderedSphynxDiskReply, error)
	HasInSphynxMemory(ctx context.Context, in *HasInSphynxMemoryRequest, opts ...grpc.CallOption) (*HasInSphynxMemoryReply, error)
	Clear(ctx context.Context, in *ClearRequest, opts ...grpc.CallOption) (*ClearReply, error)
}

type sphynxClient struct {
	cc grpc.ClientConnInterface
}

func NewSphynxClient(cc grpc.ClientConnInterface) SphynxClient {
	return &sphynxClient{cc}
}

var sphynxCanComputeStreamDesc = &grpc.StreamDesc{
	StreamName: "CanCompute",
}

func (c *sphynxClient) CanCompute(ctx context.Context, in *CanComputeRequest, opts ...grpc.CallOption) (*CanComputeReply, error) {
	out := new(CanComputeReply)
	err := c.cc.Invoke(ctx, "/sphynx.Sphynx/CanCompute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var sphynxComputeStreamDesc = &grpc.StreamDesc{
	StreamName: "Compute",
}

func (c *sphynxClient) Compute(ctx context.Context, in *ComputeRequest, opts ...grpc.CallOption) (*ComputeReply, error) {
	out := new(ComputeReply)
	err := c.cc.Invoke(ctx, "/sphynx.Sphynx/Compute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var sphynxGetScalarStreamDesc = &grpc.StreamDesc{
	StreamName: "GetScalar",
}

func (c *sphynxClient) GetScalar(ctx context.Context, in *GetScalarRequest, opts ...grpc.CallOption) (*GetScalarReply, error) {
	out := new(GetScalarReply)
	err := c.cc.Invoke(ctx, "/sphynx.Sphynx/GetScalar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var sphynxWriteToUnorderedDiskStreamDesc = &grpc.StreamDesc{
	StreamName: "WriteToUnorderedDisk",
}

func (c *sphynxClient) WriteToUnorderedDisk(ctx context.Context, in *WriteToUnorderedDiskRequest, opts ...grpc.CallOption) (*WriteToUnorderedDiskReply, error) {
	out := new(WriteToUnorderedDiskReply)
	err := c.cc.Invoke(ctx, "/sphynx.Sphynx/WriteToUnorderedDisk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var sphynxReadFromUnorderedDiskStreamDesc = &grpc.StreamDesc{
	StreamName: "ReadFromUnorderedDisk",
}

func (c *sphynxClient) ReadFromUnorderedDisk(ctx context.Context, in *ReadFromUnorderedDiskRequest, opts ...grpc.CallOption) (*ReadFromUnorderedDiskReply, error) {
	out := new(ReadFromUnorderedDiskReply)
	err := c.cc.Invoke(ctx, "/sphynx.Sphynx/ReadFromUnorderedDisk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var sphynxWriteToOrderedDiskStreamDesc = &grpc.StreamDesc{
	StreamName: "WriteToOrderedDisk",
}

func (c *sphynxClient) WriteToOrderedDisk(ctx context.Context, in *WriteToOrderedDiskRequest, opts ...grpc.CallOption) (*WriteToOrderedDiskReply, error) {
	out := new(WriteToOrderedDiskReply)
	err := c.cc.Invoke(ctx, "/sphynx.Sphynx/WriteToOrderedDisk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var sphynxReadFromOrderedSphynxDiskStreamDesc = &grpc.StreamDesc{
	StreamName: "ReadFromOrderedSphynxDisk",
}

func (c *sphynxClient) ReadFromOrderedSphynxDisk(ctx context.Context, in *ReadFromOrderedSphynxDiskRequest, opts ...grpc.CallOption) (*ReadFromOrderedSphynxDiskReply, error) {
	out := new(ReadFromOrderedSphynxDiskReply)
	err := c.cc.Invoke(ctx, "/sphynx.Sphynx/ReadFromOrderedSphynxDisk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var sphynxHasOnOrderedSphynxDiskStreamDesc = &grpc.StreamDesc{
	StreamName: "HasOnOrderedSphynxDisk",
}

func (c *sphynxClient) HasOnOrderedSphynxDisk(ctx context.Context, in *HasOnOrderedSphynxDiskRequest, opts ...grpc.CallOption) (*HasOnOrderedSphynxDiskReply, error) {
	out := new(HasOnOrderedSphynxDiskReply)
	err := c.cc.Invoke(ctx, "/sphynx.Sphynx/HasOnOrderedSphynxDisk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var sphynxHasInSphynxMemoryStreamDesc = &grpc.StreamDesc{
	StreamName: "HasInSphynxMemory",
}

func (c *sphynxClient) HasInSphynxMemory(ctx context.Context, in *HasInSphynxMemoryRequest, opts ...grpc.CallOption) (*HasInSphynxMemoryReply, error) {
	out := new(HasInSphynxMemoryReply)
	err := c.cc.Invoke(ctx, "/sphynx.Sphynx/HasInSphynxMemory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var sphynxClearStreamDesc = &grpc.StreamDesc{
	StreamName: "Clear",
}

func (c *sphynxClient) Clear(ctx context.Context, in *ClearRequest, opts ...grpc.CallOption) (*ClearReply, error) {
	out := new(ClearReply)
	err := c.cc.Invoke(ctx, "/sphynx.Sphynx/Clear", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SphynxService is the service API for Sphynx service.
// Fields should be assigned to their respective handler implementations only before
// RegisterSphynxService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type SphynxService struct {
	CanCompute                func(context.Context, *CanComputeRequest) (*CanComputeReply, error)
	Compute                   func(context.Context, *ComputeRequest) (*ComputeReply, error)
	GetScalar                 func(context.Context, *GetScalarRequest) (*GetScalarReply, error)
	WriteToUnorderedDisk      func(context.Context, *WriteToUnorderedDiskRequest) (*WriteToUnorderedDiskReply, error)
	ReadFromUnorderedDisk     func(context.Context, *ReadFromUnorderedDiskRequest) (*ReadFromUnorderedDiskReply, error)
	WriteToOrderedDisk        func(context.Context, *WriteToOrderedDiskRequest) (*WriteToOrderedDiskReply, error)
	ReadFromOrderedSphynxDisk func(context.Context, *ReadFromOrderedSphynxDiskRequest) (*ReadFromOrderedSphynxDiskReply, error)
	HasOnOrderedSphynxDisk    func(context.Context, *HasOnOrderedSphynxDiskRequest) (*HasOnOrderedSphynxDiskReply, error)
	HasInSphynxMemory         func(context.Context, *HasInSphynxMemoryRequest) (*HasInSphynxMemoryReply, error)
	Clear                     func(context.Context, *ClearRequest) (*ClearReply, error)
}

func (s *SphynxService) canCompute(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CanComputeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.CanCompute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/sphynx.Sphynx/CanCompute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.CanCompute(ctx, req.(*CanComputeRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SphynxService) compute(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComputeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Compute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/sphynx.Sphynx/Compute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Compute(ctx, req.(*ComputeRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SphynxService) getScalar(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScalarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.GetScalar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/sphynx.Sphynx/GetScalar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetScalar(ctx, req.(*GetScalarRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SphynxService) writeToUnorderedDisk(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteToUnorderedDiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.WriteToUnorderedDisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/sphynx.Sphynx/WriteToUnorderedDisk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.WriteToUnorderedDisk(ctx, req.(*WriteToUnorderedDiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SphynxService) readFromUnorderedDisk(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadFromUnorderedDiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.ReadFromUnorderedDisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/sphynx.Sphynx/ReadFromUnorderedDisk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ReadFromUnorderedDisk(ctx, req.(*ReadFromUnorderedDiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SphynxService) writeToOrderedDisk(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteToOrderedDiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.WriteToOrderedDisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/sphynx.Sphynx/WriteToOrderedDisk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.WriteToOrderedDisk(ctx, req.(*WriteToOrderedDiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SphynxService) readFromOrderedSphynxDisk(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadFromOrderedSphynxDiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.ReadFromOrderedSphynxDisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/sphynx.Sphynx/ReadFromOrderedSphynxDisk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ReadFromOrderedSphynxDisk(ctx, req.(*ReadFromOrderedSphynxDiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SphynxService) hasOnOrderedSphynxDisk(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasOnOrderedSphynxDiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.HasOnOrderedSphynxDisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/sphynx.Sphynx/HasOnOrderedSphynxDisk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.HasOnOrderedSphynxDisk(ctx, req.(*HasOnOrderedSphynxDiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SphynxService) hasInSphynxMemory(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasInSphynxMemoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.HasInSphynxMemory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/sphynx.Sphynx/HasInSphynxMemory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.HasInSphynxMemory(ctx, req.(*HasInSphynxMemoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SphynxService) clear(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Clear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/sphynx.Sphynx/Clear",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Clear(ctx, req.(*ClearRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterSphynxService registers a service implementation with a gRPC server.
func RegisterSphynxService(s grpc.ServiceRegistrar, srv *SphynxService) {
	srvCopy := *srv
	if srvCopy.CanCompute == nil {
		srvCopy.CanCompute = func(context.Context, *CanComputeRequest) (*CanComputeReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method CanCompute not implemented")
		}
	}
	if srvCopy.Compute == nil {
		srvCopy.Compute = func(context.Context, *ComputeRequest) (*ComputeReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Compute not implemented")
		}
	}
	if srvCopy.GetScalar == nil {
		srvCopy.GetScalar = func(context.Context, *GetScalarRequest) (*GetScalarReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method GetScalar not implemented")
		}
	}
	if srvCopy.WriteToUnorderedDisk == nil {
		srvCopy.WriteToUnorderedDisk = func(context.Context, *WriteToUnorderedDiskRequest) (*WriteToUnorderedDiskReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method WriteToUnorderedDisk not implemented")
		}
	}
	if srvCopy.ReadFromUnorderedDisk == nil {
		srvCopy.ReadFromUnorderedDisk = func(context.Context, *ReadFromUnorderedDiskRequest) (*ReadFromUnorderedDiskReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method ReadFromUnorderedDisk not implemented")
		}
	}
	if srvCopy.WriteToOrderedDisk == nil {
		srvCopy.WriteToOrderedDisk = func(context.Context, *WriteToOrderedDiskRequest) (*WriteToOrderedDiskReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method WriteToOrderedDisk not implemented")
		}
	}
	if srvCopy.ReadFromOrderedSphynxDisk == nil {
		srvCopy.ReadFromOrderedSphynxDisk = func(context.Context, *ReadFromOrderedSphynxDiskRequest) (*ReadFromOrderedSphynxDiskReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method ReadFromOrderedSphynxDisk not implemented")
		}
	}
	if srvCopy.HasOnOrderedSphynxDisk == nil {
		srvCopy.HasOnOrderedSphynxDisk = func(context.Context, *HasOnOrderedSphynxDiskRequest) (*HasOnOrderedSphynxDiskReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method HasOnOrderedSphynxDisk not implemented")
		}
	}
	if srvCopy.HasInSphynxMemory == nil {
		srvCopy.HasInSphynxMemory = func(context.Context, *HasInSphynxMemoryRequest) (*HasInSphynxMemoryReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method HasInSphynxMemory not implemented")
		}
	}
	if srvCopy.Clear == nil {
		srvCopy.Clear = func(context.Context, *ClearRequest) (*ClearReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Clear not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "sphynx.Sphynx",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "CanCompute",
				Handler:    srvCopy.canCompute,
			},
			{
				MethodName: "Compute",
				Handler:    srvCopy.compute,
			},
			{
				MethodName: "GetScalar",
				Handler:    srvCopy.getScalar,
			},
			{
				MethodName: "WriteToUnorderedDisk",
				Handler:    srvCopy.writeToUnorderedDisk,
			},
			{
				MethodName: "ReadFromUnorderedDisk",
				Handler:    srvCopy.readFromUnorderedDisk,
			},
			{
				MethodName: "WriteToOrderedDisk",
				Handler:    srvCopy.writeToOrderedDisk,
			},
			{
				MethodName: "ReadFromOrderedSphynxDisk",
				Handler:    srvCopy.readFromOrderedSphynxDisk,
			},
			{
				MethodName: "HasOnOrderedSphynxDisk",
				Handler:    srvCopy.hasOnOrderedSphynxDisk,
			},
			{
				MethodName: "HasInSphynxMemory",
				Handler:    srvCopy.hasInSphynxMemory,
			},
			{
				MethodName: "Clear",
				Handler:    srvCopy.clear,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "proto/sphynx.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewSphynxService creates a new SphynxService containing the
// implemented methods of the Sphynx service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewSphynxService(s interface{}) *SphynxService {
	ns := &SphynxService{}
	if h, ok := s.(interface {
		CanCompute(context.Context, *CanComputeRequest) (*CanComputeReply, error)
	}); ok {
		ns.CanCompute = h.CanCompute
	}
	if h, ok := s.(interface {
		Compute(context.Context, *ComputeRequest) (*ComputeReply, error)
	}); ok {
		ns.Compute = h.Compute
	}
	if h, ok := s.(interface {
		GetScalar(context.Context, *GetScalarRequest) (*GetScalarReply, error)
	}); ok {
		ns.GetScalar = h.GetScalar
	}
	if h, ok := s.(interface {
		WriteToUnorderedDisk(context.Context, *WriteToUnorderedDiskRequest) (*WriteToUnorderedDiskReply, error)
	}); ok {
		ns.WriteToUnorderedDisk = h.WriteToUnorderedDisk
	}
	if h, ok := s.(interface {
		ReadFromUnorderedDisk(context.Context, *ReadFromUnorderedDiskRequest) (*ReadFromUnorderedDiskReply, error)
	}); ok {
		ns.ReadFromUnorderedDisk = h.ReadFromUnorderedDisk
	}
	if h, ok := s.(interface {
		WriteToOrderedDisk(context.Context, *WriteToOrderedDiskRequest) (*WriteToOrderedDiskReply, error)
	}); ok {
		ns.WriteToOrderedDisk = h.WriteToOrderedDisk
	}
	if h, ok := s.(interface {
		ReadFromOrderedSphynxDisk(context.Context, *ReadFromOrderedSphynxDiskRequest) (*ReadFromOrderedSphynxDiskReply, error)
	}); ok {
		ns.ReadFromOrderedSphynxDisk = h.ReadFromOrderedSphynxDisk
	}
	if h, ok := s.(interface {
		HasOnOrderedSphynxDisk(context.Context, *HasOnOrderedSphynxDiskRequest) (*HasOnOrderedSphynxDiskReply, error)
	}); ok {
		ns.HasOnOrderedSphynxDisk = h.HasOnOrderedSphynxDisk
	}
	if h, ok := s.(interface {
		HasInSphynxMemory(context.Context, *HasInSphynxMemoryRequest) (*HasInSphynxMemoryReply, error)
	}); ok {
		ns.HasInSphynxMemory = h.HasInSphynxMemory
	}
	if h, ok := s.(interface {
		Clear(context.Context, *ClearRequest) (*ClearReply, error)
	}); ok {
		ns.Clear = h.Clear
	}
	return ns
}

// UnstableSphynxService is the service API for Sphynx service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableSphynxService interface {
	CanCompute(context.Context, *CanComputeRequest) (*CanComputeReply, error)
	Compute(context.Context, *ComputeRequest) (*ComputeReply, error)
	GetScalar(context.Context, *GetScalarRequest) (*GetScalarReply, error)
	WriteToUnorderedDisk(context.Context, *WriteToUnorderedDiskRequest) (*WriteToUnorderedDiskReply, error)
	ReadFromUnorderedDisk(context.Context, *ReadFromUnorderedDiskRequest) (*ReadFromUnorderedDiskReply, error)
	WriteToOrderedDisk(context.Context, *WriteToOrderedDiskRequest) (*WriteToOrderedDiskReply, error)
	ReadFromOrderedSphynxDisk(context.Context, *ReadFromOrderedSphynxDiskRequest) (*ReadFromOrderedSphynxDiskReply, error)
	HasOnOrderedSphynxDisk(context.Context, *HasOnOrderedSphynxDiskRequest) (*HasOnOrderedSphynxDiskReply, error)
	HasInSphynxMemory(context.Context, *HasInSphynxMemoryRequest) (*HasInSphynxMemoryReply, error)
	Clear(context.Context, *ClearRequest) (*ClearReply, error)
}
