package identitylogic

import (
	"context"

	"go-zero-boilerplate/app/zero-service/api/pb"
	"go-zero-boilerplate/app/zero-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type WalletLoginNonceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWalletLoginNonceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WalletLoginNonceLogic {
	return &WalletLoginNonceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WalletLoginNonceLogic) WalletLoginNonce(in *pb.WalletLoginNonceReq) (*pb.WalletLoginNonceResp, error) {
	// todo: add your logic here and delete this line

	return &pb.WalletLoginNonceResp{}, nil
}
