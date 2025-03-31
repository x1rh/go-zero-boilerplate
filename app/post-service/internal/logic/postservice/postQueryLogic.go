package postservicelogic

import (
	"context"

	"go-zero-boilerplate/app/post-service/api/pb"
	"go-zero-boilerplate/app/post-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostQueryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostQueryLogic {
	return &PostQueryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostQueryLogic) PostQuery(in *pb.PostIdReq) (*pb.PostQueryResp, error) {
	// todo: add your logic here and delete this line

	return &pb.PostQueryResp{}, nil
}
