package postservicelogic

import (
	"context"

	"go-zero-boilerplate/app/post-service/api/pb"
	"go-zero-boilerplate/app/post-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostCreateLogic {
	return &PostCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostCreateLogic) PostCreate(in *pb.PostCreateReq) (*pb.PostCreateResp, error) {
	// todo: add your logic here and delete this line

	return &pb.PostCreateResp{}, nil
}
