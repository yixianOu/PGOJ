package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TestcasesModel = (*customTestcasesModel)(nil)

type (
	// TestcasesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTestcasesModel.
	TestcasesModel interface {
		testcasesModel
		SearchCasesByFields(ctx context.Context, builder squirrel.SelectBuilder, problemId int64, testGroup int64) ([]*Testcases, error)
		SelectBuilder() squirrel.SelectBuilder
		//PartialUpdate(ctx context.Context, newData *Testcases) error
	}

	customTestcasesModel struct {
		*defaultTestcasesModel
	}
)

// NewTestcasesModel returns a model for the database table.
func NewTestcasesModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TestcasesModel {
	return &customTestcasesModel{
		defaultTestcasesModel: newTestcasesModel(conn, c, opts...),
	}
}

func (m *customTestcasesModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *customTestcasesModel) SearchCasesByFields(ctx context.Context, builder squirrel.SelectBuilder, problemId int64, testGroup int64) ([]*Testcases, error) {
	builder = builder.Columns(testcasesRows)

	builder = builder.Where("problem_id = ?", problemId)
	if testGroup != 0 {
		builder = builder.Where("test_group = ?", testGroup)
	}

	builder = builder.OrderBy("test_group ASC")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var list []*Testcases
	err = m.QueryRowsNoCacheCtx(ctx, &list, query, args...)
	if err != nil {
		return nil, err
	}

	return list, nil
}

//func (m *customTestcasesModel) PartialUpdate(ctx context.Context, newData *Testcases) error {
//	data, err := m.FindOne(ctx, newData.TestId)
//	if err != nil {
//		return err
//	}
//
//	if newData.ProblemId == 0 {
//		newData.ProblemId = data.ProblemId
//	}
//	if newData.TestGroup == 0 {
//		newData.TestGroup = data.TestGroup
//	}
//	if newData.InputFilePath == "" {
//		newData.InputFilePath = data.InputFilePath
//	}
//	if newData.OutputFilePath == "" {
//		newData.OutputFilePath = data.OutputFilePath
//	}
//
//	return m.Update(ctx, newData)
//}
