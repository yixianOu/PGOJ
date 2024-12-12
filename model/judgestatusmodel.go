package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ JudgestatusModel = (*customJudgestatusModel)(nil)

type (
	// JudgestatusModel is an interface to be customized, add more methods here,
	// and implement the added methods in customJudgestatusModel.
	JudgestatusModel interface {
		judgestatusModel
		SelectBuilder() squirrel.SelectBuilder
		PartialUpdate(ctx context.Context, newData *Judgestatus) error
		SearchJudgestatusByFields(ctx context.Context, builder squirrel.SelectBuilder, page int64, limit int64, userId int64, problemId int64, result string, language string, submittime int64, contest int64, problemtitle string, order bool) ([]*Judgestatus, error)
	}

	customJudgestatusModel struct {
		*defaultJudgestatusModel
	}
)

// NewJudgestatusModel returns a model for the database table.
func NewJudgestatusModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) JudgestatusModel {
	return &customJudgestatusModel{
		defaultJudgestatusModel: newJudgestatusModel(conn, c, opts...),
	}
}

func (m *customJudgestatusModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *customJudgestatusModel) SearchJudgestatusByFields(ctx context.Context, builder squirrel.SelectBuilder, page int64, limit int64, userId int64, problemId int64, result string, language string, submittime int64, contest int64, problemtitle string, order bool) ([]*Judgestatus, error) {
	builder = builder.Columns(judgestatusRows)

	if problemId != 0 {
		builder = builder.Where("problem_id = ?", problemId)
	}
	if userId != 0 {
		builder = builder.Where("user_id = ?", userId)
	}
	if result != "" {
		builder = builder.Where("result = ?", result)
	}
	if language != "" {
		builder = builder.Where("language like ?", "%"+language+"%")
	}
	if problemtitle != "" {
		builder = builder.Where("problemtitle like ?", "%"+problemtitle+"%")
	}
	if contest != 0 {
		builder = builder.Where("contest = ?", contest)
	}
	if submittime != 0 {
		t := time.Unix(submittime, 0)
		builder = builder.Where(squirrel.GtOrEq{"create_time": t})
	}

	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * limit
	if limit != 0 {
		builder = builder.Limit(uint64(limit)).Offset(uint64(offset))
	}
	if order {
		builder = builder.OrderBy("judge_id asc")
	} else {
		builder = builder.OrderBy("judge_id desc")
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	var resp []*Judgestatus
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, args...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customJudgestatusModel) PartialUpdate(ctx context.Context, newData *Judgestatus) error {
	data, err := m.FindOne(ctx, newData.JudgeId)
	if err != nil {
		return err
	}
	if newData.UserId == 0 {
		newData.UserId = data.UserId
	}
	if newData.ProblemId == 0 {
		newData.ProblemId = data.ProblemId
	}
	if newData.Oj == "" {
		newData.Oj = data.Oj
	}
	if newData.Result == "" {
		newData.Result = data.Result
	}
	if newData.Time == 0 {
		newData.Time = data.Time
	}
	if newData.Memory == 0 {
		newData.Memory = data.Memory
	}
	if newData.Language == "" {
		newData.Language = data.Language
	}
	if newData.Code == "" {
		newData.Code = data.Code
	}
	if newData.CreateTime.IsZero() {
		newData.CreateTime = data.CreateTime
	}
	if newData.Judger == "" {
		newData.Judger = data.Judger
	}
	if newData.Contest == 0 {
		newData.Contest = data.Contest
	}
	if newData.Contestproblem == 0 {
		newData.Contestproblem = data.Contestproblem
	}
	if newData.Code == "" {
		newData.Code = data.Code
	}
	if newData.Length == 0 {
		newData.Length = data.Length
	}
	if newData.Testcase == "" {
		newData.Testcase = data.Testcase
	}
	if newData.Ip == "" {
		newData.Ip = data.Ip
	}
	if newData.Message == "" {
		newData.Message = data.Message
	}
	if newData.Problemtitle == "" {
		newData.Problemtitle = data.Problemtitle
	}
	if newData.Rating == 0 {
		newData.Rating = data.Rating
	}
	if newData.InputData.Valid == false {
		newData.InputData = data.InputData
	}
	if newData.SampleOutPut.Valid == false {
		newData.SampleOutPut = data.SampleOutPut
	}
	if newData.UserOutPut.Valid == false {
		newData.UserOutPut = data.UserOutPut
	}

	return m.Update(ctx, newData)
}
