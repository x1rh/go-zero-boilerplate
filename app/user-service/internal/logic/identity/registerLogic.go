package identitylogic

import (
	"context"

	"go-zero-boilerplate/app/user-service/api/pb"
	"go-zero-boilerplate/app/user-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	// todo: add your logic here and delete this line

	return &pb.RegisterResp{}, nil
}
