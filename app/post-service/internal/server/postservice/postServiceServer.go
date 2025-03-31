// Code generated by goctl. DO NOT EDIT.
// Source: post_service.proto

package server

import (
	"context"

	"go-zero-boilerplate/app/post-service/api/pb"
	"go-zero-boilerplate/app/post-service/internal/logic/postservice"
	"go-zero-boilerplate/app/post-service/internal/svc"
)

type PostServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedPostServiceServer
}

func NewPostServiceServer(svcCtx *svc.ServiceContext) *PostServiceServer {
	return &PostServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *PostServiceServer) PostCreate(ctx context.Context, in *pb.PostCreateReq) (*pb.PostCreateResp, error) {
	l := postservicelogic.NewPostCreateLogic(ctx, s.svcCtx)
	return l.PostCreate(in)
}

func (s *PostServiceServer) PostQuery(ctx context.Context, in *pb.PostIdReq) (*pb.PostQueryResp, error) {
	l := postservicelogic.NewPostQueryLogic(ctx, s.svcCtx)
	return l.PostQuery(in)
}

func (s *PostServiceServer) PostUpdate(ctx context.Context, in *pb.PostUpdateReq) (*pb.CodeMessage, error) {
	l := postservicelogic.NewPostUpdateLogic(ctx, s.svcCtx)
	return l.PostUpdate(in)
}

func (s *PostServiceServer) PostDelete(ctx context.Context, in *pb.PostIdReq) (*pb.CodeMessage, error) {
	l := postservicelogic.NewPostDeleteLogic(ctx, s.svcCtx)
	return l.PostDelete(in)
}
