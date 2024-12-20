package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserProfileModel = (*customUserProfileModel)(nil)

type (
	// UserProfileModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserProfileModel.
	UserProfileModel interface {
		userProfileModel
		SelectBuilder() squirrel.SelectBuilder
		SearchUserProfileByFields(ctx context.Context, builder squirrel.SelectBuilder, page int64, pageSize int64, description string, school string, orderByScore bool) ([]*UserProfile, error)
		SortUserByScoreAndReturnRank(ctx context.Context, userId int64) (int64, error)
		PartialUpdateProfile(ctx context.Context, newData *UserProfile) error
	}

	customUserProfileModel struct {
		*defaultUserProfileModel
	}
)

// NewUserProfileModel returns a model for the database table.
func NewUserProfileModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserProfileModel {
	return &customUserProfileModel{
		defaultUserProfileModel: newUserProfileModel(conn, c, opts...),
	}
}

func (m *customUserProfileModel) PartialUpdateProfile(ctx context.Context, newData *UserProfile) error {
	var oldProfile *UserProfile
	var err error
	if newData.Id != 0 {
		oldProfile, err = m.FindOne(ctx, newData.Id)
		if err != nil {
			return err
		}
		newData.UserId = oldProfile.UserId
	} else {
		oldProfile, err = m.FindOneByUserId(ctx, newData.UserId)
		if err != nil {
			return err
		}
		newData.Id = oldProfile.Id
	}

	if newData.Score == 0 {
		newData.Score = oldProfile.Score
	}
	if newData.Rating == 0 {
		newData.Rating = oldProfile.Rating
	}
	if newData.SubmitCount == 0 {
		newData.SubmitCount = oldProfile.SubmitCount
	}
	if newData.ACCount == 0 {
		newData.ACCount = oldProfile.ACCount
	}
	if !newData.Description.Valid {
		newData.Description = oldProfile.Description
	}
	if !newData.Name.Valid {
		newData.Name = oldProfile.Name
	}
	if !newData.School.Valid {
		newData.School = oldProfile.School
	}
	if !newData.Phone.Valid {
		newData.Phone = oldProfile.Phone
	}

	return m.Update(ctx, newData)
}

func (m *customUserProfileModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

// SearchUserProfileByFields 根据查询条件列出所有profile并且按照指定的字段排序
func (m *customUserProfileModel) SearchUserProfileByFields(ctx context.Context, builder squirrel.SelectBuilder, page int64, pageSize int64, description string, school string, orderByScore bool) ([]*UserProfile, error) {
	builder = builder.Columns(userProfileRows)

	if description != "" {
		builder = builder.Where("description like ?", "%"+description+"%")
	}
	if school != "" {
		builder = builder.Where("school like ?", "%"+school+"%")
	}
	if orderByScore {
		//score越高越靠前
		builder = builder.OrderBy("score desc")
	} else {
		builder = builder.OrderBy("id desc")
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserProfile
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// SortUserByScoreAndReturnRank 根据分数排序并返回排名
func (m *customUserProfileModel) SortUserByScoreAndReturnRank(ctx context.Context, userId int64) (int64, error) {
	builder := squirrel.Select("user_id").From(m.table).OrderBy("score desc")
	query, values, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	var resp []int64
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	if err != nil {
		return 0, err
	}

	for i, v := range resp {
		if v == userId {
			return int64(i + 1), nil
		}
	}
	//如果没有找到，返回0
	return 0, nil
}
