package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProblemTagModel = (*customProblemTagModel)(nil)

type (
	// ProblemTagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProblemTagModel.
	ProblemTagModel interface {
		problemTagModel
		SelectBuilder() squirrel.SelectBuilder
		FindTagsByProblemId(ctx context.Context, builder squirrel.SelectBuilder, problemId int64) ([]*ProblemTag, error)
		FindProblemsByTagId(ctx context.Context, builder squirrel.SelectBuilder, tagId int64, page int64, limit int64) ([]*ProblemTag, error)
	}

	customProblemTagModel struct {
		*defaultProblemTagModel
	}
)

// NewProblemTagModel returns a model for the database table.
func NewProblemTagModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProblemTagModel {
	return &customProblemTagModel{
		defaultProblemTagModel: newProblemTagModel(conn, c, opts...),
	}
}

func (m *customProblemTagModel) FindTagsByProblemId(ctx context.Context, builder squirrel.SelectBuilder, problemId int64) (ProblemTags []*ProblemTag, err error) {
	builder = builder.Columns(problemTagRows)

	query, values, err := builder.Where(squirrel.Eq{"problem_id": problemId}).ToSql()
	if err != nil {
		return nil, err
	}

	err = m.QueryRowsNoCacheCtx(ctx, &ProblemTags, query, values...)
	return
}

func (m *customProblemTagModel) FindProblemsByTagId(ctx context.Context, builder squirrel.SelectBuilder, tagId int64, page int64, limit int64) (ProblemTags []*ProblemTag, err error) {
	builder = builder.Columns(problemTagRows)

	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * limit
	builder = builder.Offset(uint64(offset)).Limit(uint64(limit))

	query, values, err := builder.Where(squirrel.Eq{"tag_id": tagId}).ToSql()
	if err != nil {
		return nil, err
	}

	err = m.QueryRowsNoCacheCtx(ctx, &ProblemTags, query, values...)
	return
}

func (m *customProblemTagModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}
