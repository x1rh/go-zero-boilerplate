package identitylogic

import (
	"context"

	"go-zero-boilerplate/model/user"
	"go-zero-boilerplate/pkg/zero-contrib/appctx"
	"go-zero-boilerplate/pkg/zero-contrib/errx"
	"go-zero-boilerplate/app/user-service/api/pb"
	"go-zero-boilerplate/app/user-service/internal/svc"
	
	"github.com/jinzhu/copier"
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
	var u user.User
	if err := l.svcCtx.DB.Model(user.User{}).Where("id=?", uid).First(&u).Error; err != nil {
		return nil, errx.Error(errx.InvalidArgument, err, errx.MsgInvalidArgument)
	}
	
	userInfo := &pb.UserInfo{}
	if err := copier.Copy(userInfo, u); err != nil { 
		return nil, errx.Error(errx.CodeInternalServerErr, err, errx.MsgInternalServerErr)
	}
	userInfo.Uid = uint64(u.ID)

	return &pb.UserinfoResp{Data: userInfo}, nil
}
