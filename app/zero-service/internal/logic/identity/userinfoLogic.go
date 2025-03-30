package identitylogic

import (
	"context"
	"go-zero-boilerplate/pkg/zero-contrib/appctx"
	"go-zero-boilerplate/pkg/zero-contrib/errx"

	"github.com/jinzhu/copier"

	"go-zero-boilerplate/app/zero-service/api/pb"
	"go-zero-boilerplate/app/zero-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserinfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserinfoLogic {
	return &UserinfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserinfoLogic) Userinfo(in *pb.UserinfoReq) (*pb.UserinfoResp, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, errx.Error(errx.InvalidArgument, err, "invalid argument")
	}

	uid := appctx.GetUid(l.ctx)
	u, err := l.svcCtx.UserModel.FindOne(l.ctx, uid)
	if err != nil {
		return nil, errx.Error(errx.Internal, err, errx.MsgInternalServerErr)
	}
	userInfo := &pb.UserInfo{}
	err = copier.Copy(userInfo, u)
	if err != nil {
		return nil, errx.Error(errx.CodeInternalServerErr, err, errx.MsgInternalServerErr)
	}
	userInfo.Uid = uint64(u.Id)

	return &pb.UserinfoResp{Data: userInfo}, nil
}
