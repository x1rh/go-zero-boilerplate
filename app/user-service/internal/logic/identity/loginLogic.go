package identitylogic

import (
	"context"

	"go-zero-boilerplate/app/user-service/api/pb"
	"go-zero-boilerplate/app/user-service/internal/svc"
	"go-zero-boilerplate/model/user"
	"go-zero-boilerplate/pkg/zero-contrib/errx"
	"go-zero-boilerplate/pkg/zero-contrib/jwtx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, errx.Error(errx.InvalidArgument, err, errx.MsgInvalidArgument)
	}
	var u user.User
	err := l.svcCtx.DB.Model(user.User{}).Where("email = ?", in.Email).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg := "invalid user or password"
			return nil, errx.Error(errx.InvalidArgument, msg, msg)
		}
		return nil, errx.Error(errx.Internal, err, errx.MsgInternalServerErr)
	}

	// check password

	token, err := l.svcCtx.JwtManager.Gen(jwtx.User{
		Uid: int64(u.ID),
	})
	if err != nil {
		return nil, errx.Error(errx.Internal, err, errx.MsgInternalServerErr)
	}

	return &pb.LoginResp{Jwt: token}, nil
}
