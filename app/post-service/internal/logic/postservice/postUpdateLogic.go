package postservicelogic

import (
	"context"

	"go-zero-boilerplate/app/post-service/api/pb"
	"go-zero-boilerplate/app/post-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostUpdateLogic {
	return &PostUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostUpdateLogic) PostUpdate(in *pb.PostUpdateReq) (*pb.CodeMessage, error) {
	// todo: add your logic here and delete this line

	return &pb.CodeMessage{}, nil
}
