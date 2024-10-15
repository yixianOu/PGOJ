// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: problem.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ProblemService_AddProblem_FullMethodName          = "/problem.problem_service/AddProblem"
	ProblemService_UpdateProblem_FullMethodName       = "/problem.problem_service/UpdateProblem"
	ProblemService_DelProblem_FullMethodName          = "/problem.problem_service/DelProblem"
	ProblemService_GetProblemById_FullMethodName      = "/problem.problem_service/GetProblemById"
	ProblemService_SearchProblem_FullMethodName       = "/problem.problem_service/SearchProblem"
	ProblemService_ListProblemsByTagId_FullMethodName = "/problem.problem_service/ListProblemsByTagId"
	ProblemService_UpdateProblemdata_FullMethodName   = "/problem.problem_service/UpdateProblemdata"
	ProblemService_GetProblemdataById_FullMethodName  = "/problem.problem_service/GetProblemdataById"
	ProblemService_SearchProblemdata_FullMethodName   = "/problem.problem_service/SearchProblemdata"
	ProblemService_AddTag_FullMethodName              = "/problem.problem_service/AddTag"
	ProblemService_UpdateTag_FullMethodName           = "/problem.problem_service/UpdateTag"
	ProblemService_DelTag_FullMethodName              = "/problem.problem_service/DelTag"
	ProblemService_GetTagById_FullMethodName          = "/problem.problem_service/GetTagById"
	ProblemService_SearchTag_FullMethodName           = "/problem.problem_service/SearchTag"
	ProblemService_ListTagsByProblemId_FullMethodName = "/problem.problem_service/ListTagsByProblemId"
	ProblemService_AddTestcases_FullMethodName        = "/problem.problem_service/AddTestcases"
	ProblemService_UpdateTestcases_FullMethodName     = "/problem.problem_service/UpdateTestcases"
	ProblemService_DelTestcases_FullMethodName        = "/problem.problem_service/DelTestcases"
	ProblemService_GetTestcasesById_FullMethodName    = "/problem.problem_service/GetTestcasesById"
	ProblemService_SearchTestcases_FullMethodName     = "/problem.problem_service/SearchTestcases"
)

// ProblemServiceClient is the client API for ProblemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProblemServiceClient interface {
	AddProblem(ctx context.Context, in *AddProblemReq, opts ...grpc.CallOption) (*AddProblemResp, error)
	UpdateProblem(ctx context.Context, in *UpdateProblemReq, opts ...grpc.CallOption) (*UpdateProblemResp, error)
	DelProblem(ctx context.Context, in *DelProblemReq, opts ...grpc.CallOption) (*DelProblemResp, error)
	GetProblemById(ctx context.Context, in *GetProblemByIdReq, opts ...grpc.CallOption) (*GetProblemByIdResp, error)
	SearchProblem(ctx context.Context, in *SearchProblemReq, opts ...grpc.CallOption) (*SearchProblemResp, error)
	ListProblemsByTagId(ctx context.Context, in *ListProblemsByTagIdReq, opts ...grpc.CallOption) (*ListProblemsByTagIdResp, error)
	// -----------------------problemTag-----------------------
	//
	//	rpc AddProblemTag(AddProblemTagReq) returns (AddProblemTagResp);
	//	rpc UpdateProblemTag(UpdateProblemTagReq) returns (UpdateProblemTagResp);
	//	rpc DelProblemTag(DelProblemTagReq) returns (DelProblemTagResp);
	//	rpc GetProblemTagById(GetProblemTagByIdReq) returns (GetProblemTagByIdResp);
	//	rpc SearchProblemTag(SearchProblemTagReq) returns (SearchProblemTagResp);
	//
	// -----------------------problemdata-----------------------
	//
	//	rpc AddProblemdata(AddProblemdataReq) returns (AddProblemdataResp);
	UpdateProblemdata(ctx context.Context, in *UpdateProblemdataReq, opts ...grpc.CallOption) (*UpdateProblemdataResp, error)
	// rpc DelProblemdata(DelProblemdataReq) returns (DelProblemdataResp);
	GetProblemdataById(ctx context.Context, in *GetProblemdataByIdReq, opts ...grpc.CallOption) (*GetProblemdataByIdResp, error)
	SearchProblemdata(ctx context.Context, in *SearchProblemdataReq, opts ...grpc.CallOption) (*SearchProblemdataResp, error)
	AddTag(ctx context.Context, in *AddTagReq, opts ...grpc.CallOption) (*AddTagResp, error)
	UpdateTag(ctx context.Context, in *UpdateTagReq, opts ...grpc.CallOption) (*UpdateTagResp, error)
	DelTag(ctx context.Context, in *DelTagReq, opts ...grpc.CallOption) (*DelTagResp, error)
	GetTagById(ctx context.Context, in *GetTagByIdReq, opts ...grpc.CallOption) (*GetTagByIdResp, error)
	SearchTag(ctx context.Context, in *SearchTagReq, opts ...grpc.CallOption) (*SearchTagResp, error)
	ListTagsByProblemId(ctx context.Context, in *ListTagsByProblemIdReq, opts ...grpc.CallOption) (*ListTagsByProblemIdResp, error)
	AddTestcases(ctx context.Context, in *AddTestcasesReq, opts ...grpc.CallOption) (*AddTestcasesResp, error)
	UpdateTestcases(ctx context.Context, in *UpdateTestcasesReq, opts ...grpc.CallOption) (*UpdateTestcasesResp, error)
	DelTestcases(ctx context.Context, in *DelTestcasesReq, opts ...grpc.CallOption) (*DelTestcasesResp, error)
	GetTestcasesById(ctx context.Context, in *GetTestcasesByIdReq, opts ...grpc.CallOption) (*GetTestcasesByIdResp, error)
	SearchTestcases(ctx context.Context, in *SearchTestcasesReq, opts ...grpc.CallOption) (*SearchTestcasesResp, error)
}

type problemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProblemServiceClient(cc grpc.ClientConnInterface) ProblemServiceClient {
	return &problemServiceClient{cc}
}

func (c *problemServiceClient) AddProblem(ctx context.Context, in *AddProblemReq, opts ...grpc.CallOption) (*AddProblemResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddProblemResp)
	err := c.cc.Invoke(ctx, ProblemService_AddProblem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) UpdateProblem(ctx context.Context, in *UpdateProblemReq, opts ...grpc.CallOption) (*UpdateProblemResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateProblemResp)
	err := c.cc.Invoke(ctx, ProblemService_UpdateProblem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) DelProblem(ctx context.Context, in *DelProblemReq, opts ...grpc.CallOption) (*DelProblemResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DelProblemResp)
	err := c.cc.Invoke(ctx, ProblemService_DelProblem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) GetProblemById(ctx context.Context, in *GetProblemByIdReq, opts ...grpc.CallOption) (*GetProblemByIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProblemByIdResp)
	err := c.cc.Invoke(ctx, ProblemService_GetProblemById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) SearchProblem(ctx context.Context, in *SearchProblemReq, opts ...grpc.CallOption) (*SearchProblemResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchProblemResp)
	err := c.cc.Invoke(ctx, ProblemService_SearchProblem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) ListProblemsByTagId(ctx context.Context, in *ListProblemsByTagIdReq, opts ...grpc.CallOption) (*ListProblemsByTagIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListProblemsByTagIdResp)
	err := c.cc.Invoke(ctx, ProblemService_ListProblemsByTagId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) UpdateProblemdata(ctx context.Context, in *UpdateProblemdataReq, opts ...grpc.CallOption) (*UpdateProblemdataResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateProblemdataResp)
	err := c.cc.Invoke(ctx, ProblemService_UpdateProblemdata_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) GetProblemdataById(ctx context.Context, in *GetProblemdataByIdReq, opts ...grpc.CallOption) (*GetProblemdataByIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProblemdataByIdResp)
	err := c.cc.Invoke(ctx, ProblemService_GetProblemdataById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) SearchProblemdata(ctx context.Context, in *SearchProblemdataReq, opts ...grpc.CallOption) (*SearchProblemdataResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchProblemdataResp)
	err := c.cc.Invoke(ctx, ProblemService_SearchProblemdata_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) AddTag(ctx context.Context, in *AddTagReq, opts ...grpc.CallOption) (*AddTagResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddTagResp)
	err := c.cc.Invoke(ctx, ProblemService_AddTag_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) UpdateTag(ctx context.Context, in *UpdateTagReq, opts ...grpc.CallOption) (*UpdateTagResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTagResp)
	err := c.cc.Invoke(ctx, ProblemService_UpdateTag_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) DelTag(ctx context.Context, in *DelTagReq, opts ...grpc.CallOption) (*DelTagResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DelTagResp)
	err := c.cc.Invoke(ctx, ProblemService_DelTag_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) GetTagById(ctx context.Context, in *GetTagByIdReq, opts ...grpc.CallOption) (*GetTagByIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTagByIdResp)
	err := c.cc.Invoke(ctx, ProblemService_GetTagById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) SearchTag(ctx context.Context, in *SearchTagReq, opts ...grpc.CallOption) (*SearchTagResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchTagResp)
	err := c.cc.Invoke(ctx, ProblemService_SearchTag_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) ListTagsByProblemId(ctx context.Context, in *ListTagsByProblemIdReq, opts ...grpc.CallOption) (*ListTagsByProblemIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTagsByProblemIdResp)
	err := c.cc.Invoke(ctx, ProblemService_ListTagsByProblemId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) AddTestcases(ctx context.Context, in *AddTestcasesReq, opts ...grpc.CallOption) (*AddTestcasesResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddTestcasesResp)
	err := c.cc.Invoke(ctx, ProblemService_AddTestcases_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) UpdateTestcases(ctx context.Context, in *UpdateTestcasesReq, opts ...grpc.CallOption) (*UpdateTestcasesResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTestcasesResp)
	err := c.cc.Invoke(ctx, ProblemService_UpdateTestcases_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) DelTestcases(ctx context.Context, in *DelTestcasesReq, opts ...grpc.CallOption) (*DelTestcasesResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DelTestcasesResp)
	err := c.cc.Invoke(ctx, ProblemService_DelTestcases_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) GetTestcasesById(ctx context.Context, in *GetTestcasesByIdReq, opts ...grpc.CallOption) (*GetTestcasesByIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTestcasesByIdResp)
	err := c.cc.Invoke(ctx, ProblemService_GetTestcasesById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *problemServiceClient) SearchTestcases(ctx context.Context, in *SearchTestcasesReq, opts ...grpc.CallOption) (*SearchTestcasesResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchTestcasesResp)
	err := c.cc.Invoke(ctx, ProblemService_SearchTestcases_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProblemServiceServer is the server API for ProblemService service.
// All implementations must embed UnimplementedProblemServiceServer
// for forward compatibility
type ProblemServiceServer interface {
	AddProblem(context.Context, *AddProblemReq) (*AddProblemResp, error)
	UpdateProblem(context.Context, *UpdateProblemReq) (*UpdateProblemResp, error)
	DelProblem(context.Context, *DelProblemReq) (*DelProblemResp, error)
	GetProblemById(context.Context, *GetProblemByIdReq) (*GetProblemByIdResp, error)
	SearchProblem(context.Context, *SearchProblemReq) (*SearchProblemResp, error)
	ListProblemsByTagId(context.Context, *ListProblemsByTagIdReq) (*ListProblemsByTagIdResp, error)
	// -----------------------problemTag-----------------------
	//
	//	rpc AddProblemTag(AddProblemTagReq) returns (AddProblemTagResp);
	//	rpc UpdateProblemTag(UpdateProblemTagReq) returns (UpdateProblemTagResp);
	//	rpc DelProblemTag(DelProblemTagReq) returns (DelProblemTagResp);
	//	rpc GetProblemTagById(GetProblemTagByIdReq) returns (GetProblemTagByIdResp);
	//	rpc SearchProblemTag(SearchProblemTagReq) returns (SearchProblemTagResp);
	//
	// -----------------------problemdata-----------------------
	//
	//	rpc AddProblemdata(AddProblemdataReq) returns (AddProblemdataResp);
	UpdateProblemdata(context.Context, *UpdateProblemdataReq) (*UpdateProblemdataResp, error)
	// rpc DelProblemdata(DelProblemdataReq) returns (DelProblemdataResp);
	GetProblemdataById(context.Context, *GetProblemdataByIdReq) (*GetProblemdataByIdResp, error)
	SearchProblemdata(context.Context, *SearchProblemdataReq) (*SearchProblemdataResp, error)
	AddTag(context.Context, *AddTagReq) (*AddTagResp, error)
	UpdateTag(context.Context, *UpdateTagReq) (*UpdateTagResp, error)
	DelTag(context.Context, *DelTagReq) (*DelTagResp, error)
	GetTagById(context.Context, *GetTagByIdReq) (*GetTagByIdResp, error)
	SearchTag(context.Context, *SearchTagReq) (*SearchTagResp, error)
	ListTagsByProblemId(context.Context, *ListTagsByProblemIdReq) (*ListTagsByProblemIdResp, error)
	AddTestcases(context.Context, *AddTestcasesReq) (*AddTestcasesResp, error)
	UpdateTestcases(context.Context, *UpdateTestcasesReq) (*UpdateTestcasesResp, error)
	DelTestcases(context.Context, *DelTestcasesReq) (*DelTestcasesResp, error)
	GetTestcasesById(context.Context, *GetTestcasesByIdReq) (*GetTestcasesByIdResp, error)
	SearchTestcases(context.Context, *SearchTestcasesReq) (*SearchTestcasesResp, error)
	mustEmbedUnimplementedProblemServiceServer()
}

// UnimplementedProblemServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProblemServiceServer struct {
}

func (UnimplementedProblemServiceServer) AddProblem(context.Context, *AddProblemReq) (*AddProblemResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProblem not implemented")
}
func (UnimplementedProblemServiceServer) UpdateProblem(context.Context, *UpdateProblemReq) (*UpdateProblemResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProblem not implemented")
}
func (UnimplementedProblemServiceServer) DelProblem(context.Context, *DelProblemReq) (*DelProblemResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelProblem not implemented")
}
func (UnimplementedProblemServiceServer) GetProblemById(context.Context, *GetProblemByIdReq) (*GetProblemByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProblemById not implemented")
}
func (UnimplementedProblemServiceServer) SearchProblem(context.Context, *SearchProblemReq) (*SearchProblemResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchProblem not implemented")
}
func (UnimplementedProblemServiceServer) ListProblemsByTagId(context.Context, *ListProblemsByTagIdReq) (*ListProblemsByTagIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProblemsByTagId not implemented")
}
func (UnimplementedProblemServiceServer) UpdateProblemdata(context.Context, *UpdateProblemdataReq) (*UpdateProblemdataResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProblemdata not implemented")
}
func (UnimplementedProblemServiceServer) GetProblemdataById(context.Context, *GetProblemdataByIdReq) (*GetProblemdataByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProblemdataById not implemented")
}
func (UnimplementedProblemServiceServer) SearchProblemdata(context.Context, *SearchProblemdataReq) (*SearchProblemdataResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchProblemdata not implemented")
}
func (UnimplementedProblemServiceServer) AddTag(context.Context, *AddTagReq) (*AddTagResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTag not implemented")
}
func (UnimplementedProblemServiceServer) UpdateTag(context.Context, *UpdateTagReq) (*UpdateTagResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTag not implemented")
}
func (UnimplementedProblemServiceServer) DelTag(context.Context, *DelTagReq) (*DelTagResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelTag not implemented")
}
func (UnimplementedProblemServiceServer) GetTagById(context.Context, *GetTagByIdReq) (*GetTagByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTagById not implemented")
}
func (UnimplementedProblemServiceServer) SearchTag(context.Context, *SearchTagReq) (*SearchTagResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchTag not implemented")
}
func (UnimplementedProblemServiceServer) ListTagsByProblemId(context.Context, *ListTagsByProblemIdReq) (*ListTagsByProblemIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTagsByProblemId not implemented")
}
func (UnimplementedProblemServiceServer) AddTestcases(context.Context, *AddTestcasesReq) (*AddTestcasesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTestcases not implemented")
}
func (UnimplementedProblemServiceServer) UpdateTestcases(context.Context, *UpdateTestcasesReq) (*UpdateTestcasesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTestcases not implemented")
}
func (UnimplementedProblemServiceServer) DelTestcases(context.Context, *DelTestcasesReq) (*DelTestcasesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelTestcases not implemented")
}
func (UnimplementedProblemServiceServer) GetTestcasesById(context.Context, *GetTestcasesByIdReq) (*GetTestcasesByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTestcasesById not implemented")
}
func (UnimplementedProblemServiceServer) SearchTestcases(context.Context, *SearchTestcasesReq) (*SearchTestcasesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchTestcases not implemented")
}
func (UnimplementedProblemServiceServer) mustEmbedUnimplementedProblemServiceServer() {}

// UnsafeProblemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProblemServiceServer will
// result in compilation errors.
type UnsafeProblemServiceServer interface {
	mustEmbedUnimplementedProblemServiceServer()
}

func RegisterProblemServiceServer(s grpc.ServiceRegistrar, srv ProblemServiceServer) {
	s.RegisterService(&ProblemService_ServiceDesc, srv)
}

func _ProblemService_AddProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProblemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).AddProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_AddProblem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).AddProblem(ctx, req.(*AddProblemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_UpdateProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProblemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).UpdateProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_UpdateProblem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).UpdateProblem(ctx, req.(*UpdateProblemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_DelProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelProblemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).DelProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_DelProblem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).DelProblem(ctx, req.(*DelProblemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_GetProblemById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProblemByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).GetProblemById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_GetProblemById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).GetProblemById(ctx, req.(*GetProblemByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_SearchProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchProblemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).SearchProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_SearchProblem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).SearchProblem(ctx, req.(*SearchProblemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_ListProblemsByTagId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProblemsByTagIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).ListProblemsByTagId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_ListProblemsByTagId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).ListProblemsByTagId(ctx, req.(*ListProblemsByTagIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_UpdateProblemdata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProblemdataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).UpdateProblemdata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_UpdateProblemdata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).UpdateProblemdata(ctx, req.(*UpdateProblemdataReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_GetProblemdataById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProblemdataByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).GetProblemdataById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_GetProblemdataById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).GetProblemdataById(ctx, req.(*GetProblemdataByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_SearchProblemdata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchProblemdataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).SearchProblemdata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_SearchProblemdata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).SearchProblemdata(ctx, req.(*SearchProblemdataReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_AddTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTagReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).AddTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_AddTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).AddTag(ctx, req.(*AddTagReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_UpdateTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTagReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).UpdateTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_UpdateTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).UpdateTag(ctx, req.(*UpdateTagReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_DelTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelTagReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).DelTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_DelTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).DelTag(ctx, req.(*DelTagReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_GetTagById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).GetTagById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_GetTagById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).GetTagById(ctx, req.(*GetTagByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_SearchTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchTagReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).SearchTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_SearchTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).SearchTag(ctx, req.(*SearchTagReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_ListTagsByProblemId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTagsByProblemIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).ListTagsByProblemId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_ListTagsByProblemId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).ListTagsByProblemId(ctx, req.(*ListTagsByProblemIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_AddTestcases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTestcasesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).AddTestcases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_AddTestcases_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).AddTestcases(ctx, req.(*AddTestcasesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_UpdateTestcases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTestcasesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).UpdateTestcases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_UpdateTestcases_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).UpdateTestcases(ctx, req.(*UpdateTestcasesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_DelTestcases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelTestcasesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).DelTestcases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_DelTestcases_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).DelTestcases(ctx, req.(*DelTestcasesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_GetTestcasesById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTestcasesByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).GetTestcasesById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_GetTestcasesById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).GetTestcasesById(ctx, req.(*GetTestcasesByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProblemService_SearchTestcases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchTestcasesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProblemServiceServer).SearchTestcases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProblemService_SearchTestcases_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProblemServiceServer).SearchTestcases(ctx, req.(*SearchTestcasesReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ProblemService_ServiceDesc is the grpc.ServiceDesc for ProblemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProblemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "problem.problem_service",
	HandlerType: (*ProblemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddProblem",
			Handler:    _ProblemService_AddProblem_Handler,
		},
		{
			MethodName: "UpdateProblem",
			Handler:    _ProblemService_UpdateProblem_Handler,
		},
		{
			MethodName: "DelProblem",
			Handler:    _ProblemService_DelProblem_Handler,
		},
		{
			MethodName: "GetProblemById",
			Handler:    _ProblemService_GetProblemById_Handler,
		},
		{
			MethodName: "SearchProblem",
			Handler:    _ProblemService_SearchProblem_Handler,
		},
		{
			MethodName: "ListProblemsByTagId",
			Handler:    _ProblemService_ListProblemsByTagId_Handler,
		},
		{
			MethodName: "UpdateProblemdata",
			Handler:    _ProblemService_UpdateProblemdata_Handler,
		},
		{
			MethodName: "GetProblemdataById",
			Handler:    _ProblemService_GetProblemdataById_Handler,
		},
		{
			MethodName: "SearchProblemdata",
			Handler:    _ProblemService_SearchProblemdata_Handler,
		},
		{
			MethodName: "AddTag",
			Handler:    _ProblemService_AddTag_Handler,
		},
		{
			MethodName: "UpdateTag",
			Handler:    _ProblemService_UpdateTag_Handler,
		},
		{
			MethodName: "DelTag",
			Handler:    _ProblemService_DelTag_Handler,
		},
		{
			MethodName: "GetTagById",
			Handler:    _ProblemService_GetTagById_Handler,
		},
		{
			MethodName: "SearchTag",
			Handler:    _ProblemService_SearchTag_Handler,
		},
		{
			MethodName: "ListTagsByProblemId",
			Handler:    _ProblemService_ListTagsByProblemId_Handler,
		},
		{
			MethodName: "AddTestcases",
			Handler:    _ProblemService_AddTestcases_Handler,
		},
		{
			MethodName: "UpdateTestcases",
			Handler:    _ProblemService_UpdateTestcases_Handler,
		},
		{
			MethodName: "DelTestcases",
			Handler:    _ProblemService_DelTestcases_Handler,
		},
		{
			MethodName: "GetTestcasesById",
			Handler:    _ProblemService_GetTestcasesById_Handler,
		},
		{
			MethodName: "SearchTestcases",
			Handler:    _ProblemService_SearchTestcases_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "problem.proto",
}
