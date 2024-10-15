package svc

import (
	"context"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"oj-micro/judgeStatus/cmd/rpc/internal/config"
	"oj-micro/judgeStatus/model"
)

type ServiceContext struct {
	Config           config.Config
	JudgeStatusModel model.JudgestatusModel
	NatsClient       *nats.Conn
	JetStream        jetstream.JetStream
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbConn := sqlx.NewMysql(c.DB.DataSource)
	natsConn, err := nats.Connect(c.NatsConf.Host)
	if err != nil {
		panic(err)
	}
	js, err := jetstream.New(natsConn)
	if err != nil {
		panic(err)
	}
	err = NewJudgeStream(js, c)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:           c,
		JudgeStatusModel: model.NewJudgestatusModel(dbConn, c.Cache),
		NatsClient:       natsConn,
		JetStream:        js,
	}
}

func NewJudgeStream(js jetstream.JetStream, c config.Config) error {
	cfg := jetstream.StreamConfig{
		Name:        c.NatsConf.StreamName,
		Subjects:    []string{c.NatsConf.StreamSubject},
		Retention:   jetstream.WorkQueuePolicy,
		Storage:     jetstream.MemoryStorage,
		NoAck:       false,
		DenyDelete:  true,
		DenyPurge:   true,
		AllowRollup: false,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := js.CreateOrUpdateStream(ctx, cfg)
	if err != nil {
		return err
	}

	return err
}
