package healthlogic

import (
	"context"

	"go-zero-boilerplate/app/post-service/api/pb"
	"go-zero-boilerplate/app/post-service/internal/svc"

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
	// todo: add your logic here and delete this line

	return &pb.Response{}, nil
}
