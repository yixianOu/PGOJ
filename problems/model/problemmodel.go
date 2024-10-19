package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProblemModel = (*customProblemModel)(nil)

type (
	// ProblemModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProblemModel.
	ProblemModel interface {
		problemModel
		SearchProblemByFields(ctx context.Context, builder squirrel.SelectBuilder, page int64, pageSize int64, author string, oj string, title string, problem_code string, description string, source string, auth int64, level int64, order bool) ([]*Problem, error)
		SelectBuilder() squirrel.SelectBuilder
		PartialUpdate(ctx context.Context, newData *Problem) error
	}

	customProblemModel struct {
		*defaultProblemModel
	}
)

// NewProblemModel returns a model for the database table.
func NewProblemModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProblemModel {
	return &customProblemModel{
		defaultProblemModel: newProblemModel(conn, c, opts...),
	}
}

func (m *customProblemModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *customProblemModel) SearchProblemByFields(ctx context.Context, builder squirrel.SelectBuilder, page int64, pageSize int64, author string, oj string, title string, problemCode string, description string, source string, auth int64, level int64, order bool) ([]*Problem, error) {
	builder = builder.Columns(problemRows)

	if auth != 0 {
		builder = builder.Where("auth = ?", auth)
	}
	if level != 0 {
		builder = builder.Where("level = ?", level)
	}
	if author != "" {
		builder = builder.Where("author like ?", "%"+author+"%")
	}
	if oj != "" {
		builder = builder.Where("oj like ?", "%"+oj+"%")
	}
	if title != "" {
		builder = builder.Where("title like ?", "%"+title+"%")
	}
	if description != "" {
		builder = builder.Where("description like ?", "%"+description+"%")
	}
	if source != "" {
		builder = builder.Where("source like ?", "%"+source+"%")
	}
	if problemCode != "" {
		builder = builder.Where("problem_code like ?", "%"+problemCode+"%")
	}
	if order {
		builder = builder.OrderBy("problem_id asc")
	} else {
		builder = builder.OrderBy("problem_id desc")
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Problem
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customProblemModel) PartialUpdate(ctx context.Context, newData *Problem) error {
	data, err := m.FindOne(ctx, newData.ProblemId)
	if err != nil {
		return err
	}
	if newData.Author == "" {
		newData.Author = data.Author
	}
	newData.Addtime = data.Addtime
	if newData.Oj == "" {
		newData.Oj = data.Oj
	}
	if newData.Title == "" {
		newData.Title = data.Title
	}
	if newData.Des == "" {
		newData.Des = data.Des
	}
	if newData.Time == 0 {
		newData.Time = data.Time
	}
	if newData.Memory == 0 {
		newData.Memory = data.Memory
	}
	if newData.Input == "" {
		newData.Input = data.Input
	}
	if newData.Output == "" {
		newData.Output = data.Output
	}
	if newData.Sinput == "" {
		newData.Sinput = data.Sinput
	}
	if newData.Soutput == "" {
		newData.Soutput = data.Soutput
	}
	if newData.Source == "" {
		newData.Source = data.Source
	}
	if newData.Auth == 0 {
		newData.Auth = data.Auth
	}
	if newData.Level == 0 {
		newData.Level = data.Level
	}
	if newData.Hint == "" {
		newData.Hint = data.Hint
	}
	if newData.ProblemCode == "" {
		newData.ProblemCode = data.ProblemCode
	}

	return m.Update(ctx, newData)
}
