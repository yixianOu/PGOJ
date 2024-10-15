package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TagModel = (*customTagModel)(nil)

type (
	// TagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTagModel.
	TagModel interface {
		tagModel
		SelectBuilder() squirrel.SelectBuilder
		SearchTagsByFields(ctx context.Context, builder squirrel.SelectBuilder, page int64, pageSize int64, tagName string, order bool) ([]*Tag, error)
	}

	customTagModel struct {
		*defaultTagModel
	}
)

// NewTagModel returns a model for the database table.
func NewTagModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TagModel {
	return &customTagModel{
		defaultTagModel: newTagModel(conn, c, opts...),
	}
}

func (m *customTagModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *customTagModel) SearchTagsByFields(ctx context.Context, builder squirrel.SelectBuilder, page int64, pageSize int64, tagName string, order bool) ([]*Tag, error) {
	builder = builder.Columns(tagRows)

	if tagName != "" {
		builder = builder.Where("tag_name like ?", "%"+tagName+"%")
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	builder = builder.Offset(uint64(offset)).Limit(uint64(pageSize))

	if order {
		builder = builder.OrderBy("tag_id DESC")
	} else {
		builder = builder.OrderBy("tag_id ASC")
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var list []*Tag
	err = m.QueryRowsNoCacheCtx(ctx, &list, query, args...)
	if err != nil {
		return nil, err
	}

	return list, nil
}
