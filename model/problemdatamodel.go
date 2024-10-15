package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProblemdataModel = (*customProblemdataModel)(nil)

type (
	// ProblemdataModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProblemdataModel.
	ProblemdataModel interface {
		problemdataModel
		SearchProblemdataByFields(ctx context.Context, builder squirrel.SelectBuilder, page int64, pageSize int64, scoreFloor int64, scoreCeil int64, Auth int64, order bool) ([]*Problemdata, error)
		SelectBuilder() squirrel.SelectBuilder
		PartialUpdate(ctx context.Context, newData *Problemdata) error
	}

	customProblemdataModel struct {
		*defaultProblemdataModel
	}
)

// NewProblemdataModel returns a model for the database table.
func NewProblemdataModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProblemdataModel {
	return &customProblemdataModel{
		defaultProblemdataModel: newProblemdataModel(conn, c, opts...),
	}
}

func (m *customProblemdataModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *customProblemdataModel) SearchProblemdataByFields(ctx context.Context, builder squirrel.SelectBuilder, page int64, pageSize int64, scoreFloor int64, scoreCeil int64, Auth int64, order bool) ([]*Problemdata, error) {
	builder = builder.Columns(problemdataRows)

	if Auth != 0 {
		builder = builder.Where("auth = ?", Auth)
	}
	if scoreCeil != 0 {
		builder = builder.Where("score <= ?", scoreCeil)
	}
	if scoreFloor != 0 {
		builder = builder.Where("score >= ?", scoreFloor)
	}
	if order {
		builder = builder.OrderBy("problemdata_id DESC")
	} else {
		builder = builder.OrderBy("problemdata_id ASC")
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Problemdata
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customProblemdataModel) PartialUpdate(ctx context.Context, newData *Problemdata) error {
	sq := squirrel.Update(m.table).Where(squirrel.Eq{"problemdata_id": newData.ProblemdataId}).Set("submission", squirrel.Expr("submission + ?", 1))
	if newData.Ac != 0 {
		sq = sq.Set("ac", squirrel.Expr("ac + ?", newData.Ac))
	}
	if newData.Mle != 0 {
		sq = sq.Set("mle", squirrel.Expr("mle + ?", newData.Mle))
	}
	if newData.Tle != 0 {
		sq = sq.Set("tle", squirrel.Expr("tle + ?", newData.Tle))
	}
	if newData.Rte != 0 {
		sq = sq.Set("rte", squirrel.Expr("rte + ?", newData.Rte))
	}
	if newData.Ole != 0 {
		sq = sq.Set("ole", squirrel.Expr("ole + ?", newData.Ole))
	}
	if newData.Ce != 0 {
		sq = sq.Set("ce", squirrel.Expr("ce + ?", newData.Ce))
	}
	if newData.Wa != 0 {
		sq = sq.Set("wa", squirrel.Expr("wa + ?", newData.Wa))
	}
	if newData.Ue != 0 {
		sq = sq.Set("ue", squirrel.Expr("ue + ?", newData.Ue))
	}
	if newData.Sf != 0 {
		sq = sq.Set("sf", squirrel.Expr("sf + ?", newData.Sf))
	}
	if newData.Fe != 0 {
		sq = sq.Set("fe", squirrel.Expr("fe + ?", newData.Fe))
	}
	if newData.Score != 0 {
		sq = sq.Set("score", newData.Score)
	}
	if newData.Auth != 0 {
		sq = sq.Set("auth", newData.Auth)
	}

	query, values, err := sq.ToSql()
	if err != nil {
		return err
	}

	_, err = m.ExecNoCacheCtx(ctx, query, values...)
	return nil
}

//data, err := m.FindOne(ctx, newData.ProblemdataId)
//if err != nil {
//	return err
//}
//if newData.ProblemId == 0 {
//	newData.ProblemId = data.ProblemId
//}
//if newData.Submission == 0 {
//	newData.Submission = data.Submission
//}
//if newData.Ac == 0 {
//	newData.Ac = data.Ac
//}
//if newData.Mle == 0 {
//	newData.Mle = data.Mle
//}
//if newData.Tle == 0 {
//	newData.Tle = data.Tle
//}
//if newData.Rte == 0 {
//	newData.Rte = data.Rte
//}
//if newData.Wa == 0 {
//	newData.Wa = data.Wa
//}
//if newData.Ole == 0 {
//	newData.Ole = data.Ole
//}
//if newData.Ce == 0 {
//	newData.Ce = data.Ce
//}
//if newData.Ue == 0 {
//	newData.Ue = data.Ue
//}
//if newData.Score == 0 {
//	newData.Score = data.Score
//}
//if newData.Auth == 0 {
//	newData.Auth = data.Auth
//}
//if newData.Sf == 0 {
//	newData.Sf = data.Sf
//}
//if newData.Fe == 0 {
//	newData.Fe = data.Fe
//}
//
//return m.Update(ctx, newData)
