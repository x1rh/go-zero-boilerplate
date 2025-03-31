// Code generated by goctl. DO NOT EDIT.
// Source: user_service.proto

package server

import (
	"context"

	"go-zero-boilerplate/app/user-service/api/pb"
	"go-zero-boilerplate/app/user-service/internal/logic/identity"
	"go-zero-boilerplate/app/user-service/internal/svc"
)

type IdentityServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedIdentityServer
}

func NewIdentityServer(svcCtx *svc.ServiceContext) *IdentityServer {
	return &IdentityServer{
		svcCtx: svcCtx,
	}
}

func (s *IdentityServer) Register(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterResp, error) {
	l := identitylogic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *IdentityServer) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginResp, error) {
	l := identitylogic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *IdentityServer) Userinfo(ctx context.Context, in *pb.UserinfoReq) (*pb.UserinfoResp, error) {
	l := identitylogic.NewUserinfoLogic(ctx, s.svcCtx)
	return l.Userinfo(in)
}
