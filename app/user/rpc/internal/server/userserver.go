// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"go-chat/app/user/rpc/internal/logic"
	"go-chat/app/user/rpc/internal/svc"
	"go-chat/app/user/rpc/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Login(ctx context.Context, in *user.LoginRequest) (*user.LoginResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) Register(ctx context.Context, in *user.RegisterRequest) (*user.RegisterResponse, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServer) UserInfo(ctx context.Context, in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	l := logic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}

func (s *UserServer) UsersInfo(ctx context.Context, in *user.UsersInfoRequest) (*user.UsersInfoResponse, error) {
	l := logic.NewUsersInfoLogic(ctx, s.svcCtx)
	return l.UsersInfo(in)
}
