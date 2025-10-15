package logic

import (
	"context"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"oj-micro/common/dataType"
	"oj-micro/common/xcode"
	"oj-micro/judgeStatus/cmd/rpc/internal/svc"
	"oj-micro/judgeStatus/cmd/rpc/pb"
	"oj-micro/judgeStatus/model"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddJudgestatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddJudgestatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddJudgestatusLogic {
	return &AddJudgestatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type judgeTest struct {
	JudgeId int64 `json:"judge_id"`
	//UserId int64  `json:"user_id"`
	ProblemId   int64  `json:"problem_id"`
	ProblemCode string `json:"problem_code"`
	Language    string `json:"language"`
	Code        string `json:"code"`
	TimeLimit   int64  `json:"time_limit"`
	MemoryLimit int64  `json:"memory_limit"`
}

type testCaseResult struct {
	JudgeId      int64  `json:"judge_id"`
	CaseId       int64  `json:"case_id"`
	Result       string `json:"result"`
	TimeCost     int64  `json:"time_cost"`
	MemoryCost   int64  `json:"memory_cost"`
	Message      string `json:"message"`
	InputData    string `json:"input_data"`
	SampleOutPut string `json:"sample_output"`
	UserOutPut   string `json:"user_output"`
}

func (l *AddJudgestatusLogic) AddJudgestatus(in *pb.AddJudgestatusReq, stream pb.JudgeService_AddJudgestatusServer) error {
	// Create a new record in the database
	result, err := l.svcCtx.JudgeStatusModel.Insert(l.ctx, &model.Judgestatus{
		UserId:       in.UserId,
		ProblemId:    in.ProblemId,
		Problemtitle: in.ProblemTitle,
		Oj:           in.Oj,
		Length:       in.Length,
		Language:     in.Language,
		//Submittime:     time.Now(),
		Judger:         in.Judger,
		Contest:        in.Contest,
		Contestproblem: in.ContestProblem,
		Code:           in.Code,
		Rating:         in.Rating,
		Ip:             in.Ip,
	})
	if err != nil {
		l.Logger.Errorf("JudgeStatusModel Insert error: %v", err)
		return xcode.ServerErr
	}

	// Get the last inserted ID
	judgeId, err := result.LastInsertId()
	if err != nil {
		l.Logger.Errorf("JudgeStatusModel LastInsertId error: %v", err)
		return xcode.ServerErr
	}

	// Create a message to send to the JetStream
	test, err := json.Marshal(judgeTest{
		JudgeId: judgeId,
		//UserId:      in.UserId,
		ProblemId:   in.ProblemId,
		ProblemCode: in.ProblemCode,
		Language:    in.Language,
		Code:        in.Code,
		TimeLimit:   in.TimeLimit,
		MemoryLimit: in.MemoryLimit,
	})
	if err != nil {
		l.Logger.Errorf("json.Marshal error: %v", err)
		return xcode.ServerErr
	}

	// Send the message to the JudgeSide
	PubAck, err := l.svcCtx.JetStream.PublishAsync(l.svcCtx.Config.NatsConf.StreamSubject, test)
	if err != nil {
		l.Logger.Errorf("JetStream PublishAsync error: %v, ack: %v", err, PubAck)
		return xcode.ServerErr
	}

	var wg sync.WaitGroup
	wg.Add(int(in.CaseNum))
	var errChannel chan error
	skipWG := func() {
		for in.CaseNum > 0 {
			wg.Done()
			atomic.AddInt64(&in.CaseNum, -1)
		}
	}

	// Subscribe to the judge result
	subscription, err := l.svcCtx.NatsClient.Subscribe(strconv.FormatInt(judgeId, 10), func(msg *nats.Msg) {
		var result testCaseResult
		err := json.Unmarshal(msg.Data, &result)
		if err != nil {
			l.Logger.Errorf("json.Unmarshal error: %v", err)
			errChannel <- err
			return
		}

		err = stream.Send(&pb.AddJudgestatusResp{
			JudgeId:      judgeId,
			TestCase:     result.CaseId,
			Result:       result.Result,
			TimeCost:     result.TimeCost,
			MemoryCost:   result.MemoryCost,
			Message:      result.Message,
			InputData:    result.InputData,
			SampleOutput: result.SampleOutPut,
			UserOutput:   result.UserOutPut,
		})
		if err != nil {
			l.Logger.Errorf("stream.Send error: %v", err)
			errChannel <- err
			return
		}

		if result.Result == string(dataType.CompileError) {
			skipWG()
			return
		}

		atomic.AddInt64(&in.CaseNum, -1)
		wg.Done()
	})
	if err != nil {
		l.Logger.Errorf("NatsClient Subscribe error: %v", err)
		return xcode.ServerErr
	}

	defer func(subscription *nats.Subscription) {
		err := subscription.Unsubscribe()
		if err != nil {
			l.Logger.Errorf("NatsClient Unsubscribe error: %v", err)
			return
		}
	}(subscription)

	var codeChannel chan xcode.Code
	go func(codeChannel chan<- xcode.Code, errChannel <-chan error) {
		select {
		case err = <-PubAck.Err():
			l.Logger.Errorf("JetStream PubAck error: %v", err)
			skipWG()
			codeChannel <- xcode.ServerErr
		case <-l.ctx.Done():
			l.Logger.Errorf("context Done: %v", l.ctx.Err())
			skipWG()
			codeChannel <- xcode.RequestTimeout
		case err = <-errChannel:
			l.Logger.Errorf("errChannel error: %v", err)
			skipWG()
			codeChannel <- xcode.ServerErr
		}
	}(codeChannel, errChannel)

	wg.Wait()
	select {
	case code := <-codeChannel:
		return code
	default:
		return nil
	}
}
