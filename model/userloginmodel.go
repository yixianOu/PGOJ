package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserLoginModel = (*customUserLoginModel)(nil)

type (
	// UserLoginModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserLoginModel.
	UserLoginModel interface {
		userLoginModel
		SearchUserByFields(ctx context.Context, builder squirrel.SelectBuilder, page int64, pageSize int64, roleLevel int64, username string, order bool) ([]*UserLogin, error)
		SelectBuilder() squirrel.SelectBuilder
		PartialUpdate(ctx context.Context, newData *UserLogin) error
	}

	customUserLoginModel struct {
		*defaultUserLoginModel
	}
)

// NewUserLoginModel returns a model for the database table.
func NewUserLoginModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserLoginModel {
	return &customUserLoginModel{
		defaultUserLoginModel: newUserLoginModel(conn, c, opts...),
	}
}

func (m *customUserLoginModel) SearchUserByFields(ctx context.Context, builder squirrel.SelectBuilder, page int64, pageSize int64, roleLevel int64, username string, order bool) ([]*UserLogin, error) {
	builder = builder.Columns(userLoginRows)

	if roleLevel != 0 {
		builder = builder.Where("role_level = ?", roleLevel)
	}
	if username != "" {
		builder = builder.Where("username like ?", "%"+username+"%")
	}
	if order {
		builder = builder.OrderBy("id desc")
	} else {
		builder = builder.OrderBy("id asc")
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserLogin
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customUserLoginModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *customUserLoginModel) PartialUpdate(ctx context.Context, newData *UserLogin) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	if newData.Password == "" {
		newData.Password = data.Password
	}
	if newData.RoleLevel == 0 {
		newData.RoleLevel = data.RoleLevel
	}
	if newData.Username == "" {
		newData.Username = data.Username
	}
	if newData.Email == "" {
		newData.Email = data.Email
	}
	if newData.CoverImageUrl == "" {
		newData.CoverImageUrl = data.CoverImageUrl
	}

	return m.Update(ctx, newData)
}
