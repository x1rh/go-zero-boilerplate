package postservicelogic

import (
	"context"

	"go-zero-boilerplate/app/post-service/api/pb"
	"go-zero-boilerplate/app/post-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostDeleteLogic {
	return &PostDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostDeleteLogic) PostDelete(in *pb.PostIdReq) (*pb.CodeMessage, error) {
	// todo: add your logic here and delete this line

	return &pb.CodeMessage{}, nil
}
