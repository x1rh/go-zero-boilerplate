package healthlogic

import (
	"context"
	"fmt"
	"time"

	"go-zero-boilerplate/app/zero-service/api/pb"
	"go-zero-boilerplate/app/zero-service/internal/svc"
	"go-zero-boilerplate/pkg/zero-contrib/errx"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *pb.Request) (*pb.Response, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, errx.Error(errx.InvalidArgument, err, errx.MsgInvalidArgument)
	}
	return &pb.Response{Pong: fmt.Sprintf("recv ping=%s, ts=%d", in.Ping, time.Now().UTC())}, nil
}
