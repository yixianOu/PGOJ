package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"oj-micro/problems/cmd/rpc/internal/config"
	"oj-micro/problems/model"
)

type ServiceContext struct {
	Config config.Config

	ProblemModel     model.ProblemModel
	ProblemdataModel model.ProblemdataModel
	TagModel         model.TagModel
	ProblemTagModel  model.ProblemTagModel
	TestCasesModel   model.TestcasesModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config: c,

		ProblemModel:     model.NewProblemModel(dbConn, c.Cache),
		ProblemdataModel: model.NewProblemdataModel(dbConn, c.Cache),
		TagModel:         model.NewTagModel(dbConn, c.Cache),
		ProblemTagModel:  model.NewProblemTagModel(dbConn, c.Cache),
		TestCasesModel:   model.NewTestcasesModel(dbConn, c.Cache),
	}
}
