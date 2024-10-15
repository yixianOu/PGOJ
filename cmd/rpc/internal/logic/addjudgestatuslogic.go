package logic

import (
	"context"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"oj-micro/common/xcode"
	"oj-micro/judgeStatus/model"
	"strconv"
	"sync"
	"time"

	"oj-micro/judgeStatus/cmd/rpc/internal/svc"
	"oj-micro/judgeStatus/cmd/rpc/pb"

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
		UserId:         in.UserId,
		ProblemId:      in.ProblemId,
		Problemtitle:   in.ProblemTitle,
		Oj:             in.Oj,
		Length:         in.Length,
		Language:       in.Language,
		Submittime:     time.Now(),
		Judger:         in.Judger,
		Contest:        in.Contest,
		Contestproblem: in.ContestProblem,
		Code:           in.Code,
		Rating:         in.Rating,
		Ip:             in.Ip,
	})
	if err != nil {
		logx.Errorf("JudgeStatusModel Insert error: %v", err)
		return xcode.ServerErr
	}

	// Get the last inserted ID
	judgeId, err := result.LastInsertId()
	if err != nil {
		logx.Errorf("JudgeStatusModel LastInsertId error: %v", err)
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
		logx.Errorf("json.Marshal error: %v", err)
		return xcode.ServerErr
	}

	// Send the message to the JudgeSide
	PubAck, err := l.svcCtx.JetStream.PublishAsync(l.svcCtx.Config.NatsConf.StreamSubject, test)
	if err != nil {
		logx.Errorf("JetStream PublishAsync error: %v, ack: %v", err, PubAck)
		return xcode.ServerErr
	}

	var wg sync.WaitGroup
	wg.Add(int(in.CaseNum))
	var errChannel chan error

	// Subscribe to the judge result
	subscription, err := l.svcCtx.NatsClient.Subscribe(strconv.FormatInt(judgeId, 10), func(msg *nats.Msg) {
		var result testCaseResult
		err := json.Unmarshal(msg.Data, &result)
		if err != nil {
			logx.Errorf("json.Unmarshal error: %v", err)
			errChannel <- err
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
			logx.Errorf("stream.Send error: %v", err)
			errChannel <- err
		}

		in.CaseNum--
		wg.Done()
	})
	if err != nil {
		logx.Errorf("NatsClient Subscribe error: %v", err)
		return xcode.ServerErr
	}
	defer func(subscription *nats.Subscription) {
		err := subscription.Unsubscribe()
		if err != nil {
			logx.Errorf("NatsClient Unsubscribe error: %v", err)
			return
		}
	}(subscription)

	var codeChannel chan xcode.Code
	go func(codeChannel chan<- xcode.Code, errChannel <-chan error) {
		select {
		case err = <-PubAck.Err():
			logx.Errorf("JetStream PubAck error: %v", err)
			for in.CaseNum > 0 {
				wg.Done()
				in.CaseNum--
			}
			codeChannel <- xcode.ServerErr
		case <-l.ctx.Done():
			logx.Errorf("context Done: %v", l.ctx.Err())
			for in.CaseNum > 0 {
				wg.Done()
				in.CaseNum--
			}
			codeChannel <- xcode.RequestTimeout
		case err = <-errChannel:
			logx.Errorf("errChannel error: %v", err)
			for in.CaseNum > 0 {
				wg.Done()
				in.CaseNum--
			}
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
