// Code generated by goctl. DO NOT EDIT.
// Source: problem.proto

package problemservice

import (
	"context"

	"oj-micro/problems/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddProblemReq                           = pb.AddProblemReq
	AddProblemResp                          = pb.AddProblemResp
	AddTagReq                               = pb.AddTagReq
	AddTagResp                              = pb.AddTagResp
	AddTestcasesReq                         = pb.AddTestcasesReq
	AddTestcasesResp                        = pb.AddTestcasesResp
	DelProblemReq                           = pb.DelProblemReq
	DelProblemResp                          = pb.DelProblemResp
	DelTagReq                               = pb.DelTagReq
	DelTagResp                              = pb.DelTagResp
	DelTestcasesReq                         = pb.DelTestcasesReq
	DelTestcasesResp                        = pb.DelTestcasesResp
	GetProblemByIdReq                       = pb.GetProblemByIdReq
	GetProblemByIdResp                      = pb.GetProblemByIdResp
	GetProblemdataByIdReq                   = pb.GetProblemdataByIdReq
	GetProblemdataByIdResp                  = pb.GetProblemdataByIdResp
	GetProblemdataByProblemIdReq            = pb.GetProblemdataByProblemIdReq
	GetProblemdataByProblemIdResp           = pb.GetProblemdataByProblemIdResp
	GetTagByIdReq                           = pb.GetTagByIdReq
	GetTagByIdResp                          = pb.GetTagByIdResp
	GetTestcasesByIdReq                     = pb.GetTestcasesByIdReq
	GetTestcasesByIdResp                    = pb.GetTestcasesByIdResp
	GetTestcasesByProblemIdAndTestGroupReq  = pb.GetTestcasesByProblemIdAndTestGroupReq
	GetTestcasesByProblemIdAndTestGroupResp = pb.GetTestcasesByProblemIdAndTestGroupResp
	ListProblemsByTagIdReq                  = pb.ListProblemsByTagIdReq
	ListProblemsByTagIdResp                 = pb.ListProblemsByTagIdResp
	ListTagsByProblemIdReq                  = pb.ListTagsByProblemIdReq
	ListTagsByProblemIdResp                 = pb.ListTagsByProblemIdResp
	Problem                                 = pb.Problem
	Problemdata                             = pb.Problemdata
	SearchProblemReq                        = pb.SearchProblemReq
	SearchProblemResp                       = pb.SearchProblemResp
	SearchProblemdataReq                    = pb.SearchProblemdataReq
	SearchProblemdataResp                   = pb.SearchProblemdataResp
	SearchTagReq                            = pb.SearchTagReq
	SearchTagResp                           = pb.SearchTagResp
	SearchTestcasesReq                      = pb.SearchTestcasesReq
	SearchTestcasesResp                     = pb.SearchTestcasesResp
	Tag                                     = pb.Tag
	Testcases                               = pb.Testcases
	UpdateProblemReq                        = pb.UpdateProblemReq
	UpdateProblemResp                       = pb.UpdateProblemResp
	UpdateProblemdataReq                    = pb.UpdateProblemdataReq
	UpdateProblemdataResp                   = pb.UpdateProblemdataResp
	UpdateTagReq                            = pb.UpdateTagReq
	UpdateTagResp                           = pb.UpdateTagResp

	ProblemService interface {
		AddProblem(ctx context.Context, in *AddProblemReq, opts ...grpc.CallOption) (*AddProblemResp, error)
		UpdateProblem(ctx context.Context, in *UpdateProblemReq, opts ...grpc.CallOption) (*UpdateProblemResp, error)
		DelProblem(ctx context.Context, in *DelProblemReq, opts ...grpc.CallOption) (*DelProblemResp, error)
		GetProblemById(ctx context.Context, in *GetProblemByIdReq, opts ...grpc.CallOption) (*GetProblemByIdResp, error)
		SearchProblem(ctx context.Context, in *SearchProblemReq, opts ...grpc.CallOption) (*SearchProblemResp, error)
		ListProblemsByTagId(ctx context.Context, in *ListProblemsByTagIdReq, opts ...grpc.CallOption) (*ListProblemsByTagIdResp, error)
		// -----------------------problemTag-----------------------
		UpdateProblemdata(ctx context.Context, in *UpdateProblemdataReq, opts ...grpc.CallOption) (*UpdateProblemdataResp, error)
		// rpc DelProblemdata(DelProblemdataReq) returns (DelProblemdataResp);
		GetProblemdataById(ctx context.Context, in *GetProblemdataByIdReq, opts ...grpc.CallOption) (*GetProblemdataByIdResp, error)
		GetProblemdataByProblemId(ctx context.Context, in *GetProblemdataByProblemIdReq, opts ...grpc.CallOption) (*GetProblemdataByProblemIdResp, error)
		SearchProblemdata(ctx context.Context, in *SearchProblemdataReq, opts ...grpc.CallOption) (*SearchProblemdataResp, error)
		AddTag(ctx context.Context, in *AddTagReq, opts ...grpc.CallOption) (*AddTagResp, error)
		UpdateTag(ctx context.Context, in *UpdateTagReq, opts ...grpc.CallOption) (*UpdateTagResp, error)
		DelTag(ctx context.Context, in *DelTagReq, opts ...grpc.CallOption) (*DelTagResp, error)
		GetTagById(ctx context.Context, in *GetTagByIdReq, opts ...grpc.CallOption) (*GetTagByIdResp, error)
		SearchTag(ctx context.Context, in *SearchTagReq, opts ...grpc.CallOption) (*SearchTagResp, error)
		ListTagsByProblemId(ctx context.Context, in *ListTagsByProblemIdReq, opts ...grpc.CallOption) (*ListTagsByProblemIdResp, error)
		AddTestcases(ctx context.Context, in *AddTestcasesReq, opts ...grpc.CallOption) (*AddTestcasesResp, error)
		DelTestcases(ctx context.Context, in *DelTestcasesReq, opts ...grpc.CallOption) (*DelTestcasesResp, error)
		GetTestcasesById(ctx context.Context, in *GetTestcasesByIdReq, opts ...grpc.CallOption) (*GetTestcasesByIdResp, error)
		GetTestcasesByProblemIdAndTestGroup(ctx context.Context, in *GetTestcasesByProblemIdAndTestGroupReq, opts ...grpc.CallOption) (*GetTestcasesByProblemIdAndTestGroupResp, error)
		SearchTestcases(ctx context.Context, in *SearchTestcasesReq, opts ...grpc.CallOption) (*SearchTestcasesResp, error)
	}

	defaultProblemService struct {
		cli zrpc.Client
	}
)

func NewProblemService(cli zrpc.Client) ProblemService {
	return &defaultProblemService{
		cli: cli,
	}
}

func (m *defaultProblemService) AddProblem(ctx context.Context, in *AddProblemReq, opts ...grpc.CallOption) (*AddProblemResp, error) {
	client := pb.NewProblemServiceClient(m.cli.Conn())
	return client.AddProblem(ctx, in, opts...)
}

func (m *defaultProblemService) UpdateProblem(ctx context.Context, in *UpdateProblemReq, opts ...grpc.CallOption) (*UpdateProblemResp, error) {
	client := pb.NewProblemServiceClient(m.cli.Conn())
	return client.UpdateProblem(ctx, in, opts...)
}

func (m *defaultProblemService) DelProblem(ctx context.Context, in *DelProblemReq, opts ...grpc.CallOption) (*DelProblemResp, error) {
	client := pb.NewProblemServiceClient(m.cli.Conn())
	return client.DelProblem(ctx, in, opts...)
}

func (m *defaultProblemService) GetProblemById(ctx context.Context, in *GetProblemByIdReq, opts ...grpc.CallOption) (*GetProblemByIdResp, error) {
	client := pb.NewProblemServiceClient(m.cli.Conn())
	return client.GetProblemById(ctx, in, opts...)
}

func (m *defaultProblemService) SearchProblem(ctx context.Context, in *SearchProblemReq, opts ...grpc.CallOption) (*SearchProblemResp, error) {
	client := pb.NewProblemServiceClient(m.cli.Conn())
	return client.SearchProblem(ctx, in, opts...)
}

func (m *defaultProblemService) ListProblemsByTagId(ctx context.Context, in *ListProblemsByTagIdReq, opts ...grpc.CallOption) (*ListProblemsByTagIdResp, error) {
	client := pb.NewProblemServiceClient(m.cli.Conn())
	return client.ListProblemsByTagId(ctx, in, opts...)
}

// -----------------------problemTag-----------------------
func (m *defaultProblemService) UpdateProblemdata(ctx context.Context, in *UpdateProblemdataReq, opts ...grpc.CallOption) (*UpdateProblemdataResp, error) {
	client := pb.NewProblemServiceClient(m.cli.Conn())
	return client.UpdateProblemdata(ctx, in, opts...)
}

// rpc DelProblemdata(DelProblemdataReq) returns (DelProblemdataResp);
func (m *defaultProblemService) GetProblemdataById(ctx context.Context, in *GetProblemdataByIdReq, opts ...grpc.CallOption) (*GetProblemdataByIdResp, error) {
	client := pb.NewProblemServiceClient(m.cli.Conn())
	return client.GetProblemdataById(ctx, in, opts...)
}

func (m *defaultProblemService) GetProblemdataByProblemId(ctx context.Context, in *GetProblemdataByProblemIdReq, opts ...grpc.CallOption) (*GetProblemdataByProblemIdResp, error) {
	client := pb.NewProblemServiceClient(m.cli.Conn())
	return client.GetProblemdataByProblemId(ctx, in, opts...)
}

func (m *defaultProblemService) SearchProblemdata(ctx context.Context, in *SearchProblemdataReq, opts ...grpc.CallOption) (*SearchProblemdataResp, error) {
	client := pb.NewProblemServiceClient(m.cli.Conn())
	return client.SearchProblemdata(ctx, in, opts...)
}

func (m *defaultProblemService) AddTag(ctx context.Context, in *AddTagReq, opts ...grpc.CallOption) (*AddTagResp, error) {
	client := pb.NewProblemServiceClient(m.cli.Conn())
	return client.AddTag(ctx, in, opts...)
}

func (m *defaultProblemService) UpdateTag(ctx context.Context, in *UpdateTagReq, opts ...grpc.CallOption) (*UpdateTagResp, error) {
	client := pb.NewProblemServiceClient(m.cli.Conn())
	return client.UpdateTag(ctx, in, opts...)
}

func (m *defaultProblemService) DelTag(ctx context.Context, in *DelTagReq, opts ...grpc.CallOption) (*DelTagResp, error) {
	client := pb.NewProblemServiceClient(m.cli.Conn())
	return client.DelTag(ctx, in, opts...)
}

func (m *defaultProblemService) GetTagById(ctx context.Context, in *GetTagByIdReq, opts ...grpc.CallOption) (*GetTagByIdResp, error) {
	client := pb.NewProblemServiceClient(m.cli.Conn())
	return client.GetTagById(ctx, in, opts...)
}

func (m *defaultProblemService) SearchTag(ctx context.Context, in *SearchTagReq, opts ...grpc.CallOption) (*SearchTagResp, error) {
	client := pb.NewProblemServiceClient(m.cli.Conn())
	return client.SearchTag(ctx, in, opts...)
}

func (m *defaultProblemService) ListTagsByProblemId(ctx context.Context, in *ListTagsByProblemIdReq, opts ...grpc.CallOption) (*ListTagsByProblemIdResp, error) {
	client := pb.NewProblemServiceClient(m.cli.Conn())
	return client.ListTagsByProblemId(ctx, in, opts...)
}

func (m *defaultProblemService) AddTestcases(ctx context.Context, in *AddTestcasesReq, opts ...grpc.CallOption) (*AddTestcasesResp, error) {
	client := pb.NewProblemServiceClient(m.cli.Conn())
	return client.AddTestcases(ctx, in, opts...)
}

func (m *defaultProblemService) DelTestcases(ctx context.Context, in *DelTestcasesReq, opts ...grpc.CallOption) (*DelTestcasesResp, error) {
	client := pb.NewProblemServiceClient(m.cli.Conn())
	return client.DelTestcases(ctx, in, opts...)
}

func (m *defaultProblemService) GetTestcasesById(ctx context.Context, in *GetTestcasesByIdReq, opts ...grpc.CallOption) (*GetTestcasesByIdResp, error) {
	client := pb.NewProblemServiceClient(m.cli.Conn())
	return client.GetTestcasesById(ctx, in, opts...)
}

func (m *defaultProblemService) GetTestcasesByProblemIdAndTestGroup(ctx context.Context, in *GetTestcasesByProblemIdAndTestGroupReq, opts ...grpc.CallOption) (*GetTestcasesByProblemIdAndTestGroupResp, error) {
	client := pb.NewProblemServiceClient(m.cli.Conn())
	return client.GetTestcasesByProblemIdAndTestGroup(ctx, in, opts...)
}

func (m *defaultProblemService) SearchTestcases(ctx context.Context, in *SearchTestcasesReq, opts ...grpc.CallOption) (*SearchTestcasesResp, error) {
	client := pb.NewProblemServiceClient(m.cli.Conn())
	return client.SearchTestcases(ctx, in, opts...)
}
