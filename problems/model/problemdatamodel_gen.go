// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	problemdataFieldNames          = builder.RawFieldNames(&Problemdata{})
	problemdataRows                = strings.Join(problemdataFieldNames, ",")
	problemdataRowsExpectAutoSet   = strings.Join(stringx.Remove(problemdataFieldNames, "`problemdata_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	problemdataRowsWithPlaceHolder = strings.Join(stringx.Remove(problemdataFieldNames, "`problemdata_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheOjMicroProblemdataProblemdataIdPrefix = "cache:ojMicro:problemdata:problemdataId:"
	cacheOjMicroProblemdataProblemIdPrefix     = "cache:ojMicro:problemdata:problemId:"
)

type (
	problemdataModel interface {
		Insert(ctx context.Context, data *Problemdata) (sql.Result, error)
		FindOne(ctx context.Context, problemdataId int64) (*Problemdata, error)
		FindOneByProblemId(ctx context.Context, problemId int64) (*Problemdata, error)
		Update(ctx context.Context, data *Problemdata) error
		Delete(ctx context.Context, problemdataId int64) error
	}

	defaultProblemdataModel struct {
		sqlc.CachedConn
		table string
	}

	Problemdata struct {
		ProblemdataId int64 `db:"problemdata_id"`
		ProblemId     int64 `db:"problem_id"`
		Submission    int64 `db:"submission"`
		Ac            int64 `db:"ac"`
		Mle           int64 `db:"mle"`
		Tle           int64 `db:"tle"`
		Rte           int64 `db:"rte"`
		Ole           int64 `db:"ole"`
		Ce            int64 `db:"ce"`
		Wa            int64 `db:"wa"`
		Ue            int64 `db:"ue"`
		Score         int64 `db:"score"`
		Auth          int64 `db:"auth"`
		Sf            int64 `db:"sf"`
		Fe            int64 `db:"fe"`
	}
)

func newProblemdataModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultProblemdataModel {
	return &defaultProblemdataModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`problemdata`",
	}
}

func (m *defaultProblemdataModel) Delete(ctx context.Context, problemdataId int64) error {
	data, err := m.FindOne(ctx, problemdataId)
	if err != nil {
		return err
	}

	ojMicroProblemdataProblemIdKey := fmt.Sprintf("%s%v", cacheOjMicroProblemdataProblemIdPrefix, data.ProblemId)
	ojMicroProblemdataProblemdataIdKey := fmt.Sprintf("%s%v", cacheOjMicroProblemdataProblemdataIdPrefix, problemdataId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `problemdata_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, problemdataId)
	}, ojMicroProblemdataProblemIdKey, ojMicroProblemdataProblemdataIdKey)
	return err
}

func (m *defaultProblemdataModel) FindOne(ctx context.Context, problemdataId int64) (*Problemdata, error) {
	ojMicroProblemdataProblemdataIdKey := fmt.Sprintf("%s%v", cacheOjMicroProblemdataProblemdataIdPrefix, problemdataId)
	var resp Problemdata
	err := m.QueryRowCtx(ctx, &resp, ojMicroProblemdataProblemdataIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `problemdata_id` = ? limit 1", problemdataRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, problemdataId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultProblemdataModel) FindOneByProblemId(ctx context.Context, problemId int64) (*Problemdata, error) {
	ojMicroProblemdataProblemIdKey := fmt.Sprintf("%s%v", cacheOjMicroProblemdataProblemIdPrefix, problemId)
	var resp Problemdata
	err := m.QueryRowIndexCtx(ctx, &resp, ojMicroProblemdataProblemIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `problem_id` = ? limit 1", problemdataRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, problemId); err != nil {
			return nil, err
		}
		return resp.ProblemdataId, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultProblemdataModel) Insert(ctx context.Context, data *Problemdata) (sql.Result, error) {
	ojMicroProblemdataProblemIdKey := fmt.Sprintf("%s%v", cacheOjMicroProblemdataProblemIdPrefix, data.ProblemId)
	ojMicroProblemdataProblemdataIdKey := fmt.Sprintf("%s%v", cacheOjMicroProblemdataProblemdataIdPrefix, data.ProblemdataId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, problemdataRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ProblemId, data.Submission, data.Ac, data.Mle, data.Tle, data.Rte, data.Ole, data.Ce, data.Wa, data.Ue, data.Score, data.Auth, data.Sf, data.Fe)
	}, ojMicroProblemdataProblemIdKey, ojMicroProblemdataProblemdataIdKey)
	return ret, err
}

func (m *defaultProblemdataModel) Update(ctx context.Context, newData *Problemdata) error {
	data, err := m.FindOne(ctx, newData.ProblemdataId)
	if err != nil {
		return err
	}

	ojMicroProblemdataProblemIdKey := fmt.Sprintf("%s%v", cacheOjMicroProblemdataProblemIdPrefix, data.ProblemId)
	ojMicroProblemdataProblemdataIdKey := fmt.Sprintf("%s%v", cacheOjMicroProblemdataProblemdataIdPrefix, data.ProblemdataId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `problemdata_id` = ?", m.table, problemdataRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.ProblemId, newData.Submission, newData.Ac, newData.Mle, newData.Tle, newData.Rte, newData.Ole, newData.Ce, newData.Wa, newData.Ue, newData.Score, newData.Auth, newData.Sf, newData.Fe, newData.ProblemdataId)
	}, ojMicroProblemdataProblemIdKey, ojMicroProblemdataProblemdataIdKey)
	return err
}

func (m *defaultProblemdataModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheOjMicroProblemdataProblemdataIdPrefix, primary)
}

func (m *defaultProblemdataModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `problemdata_id` = ? limit 1", problemdataRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultProblemdataModel) tableName() string {
	return m.table
}
