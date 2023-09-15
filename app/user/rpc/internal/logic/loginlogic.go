package logic

import (
	"context"
	"errors"
	"go-chat/app/user/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go-chat/app/user/rpc/internal/svc"
	"go-chat/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// 检查用户是否存在
	res, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if errors.Is(err, model.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "用户或密码错误")
	}
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}
	// 检查密码是否正确
	if res.Password != in.Password {
		return nil, status.Error(codes.Unauthenticated, "用户或密码错误")
	}
	return &user.LoginResponse{
		Id:     res.Id,
		Mobile: res.Mobile,
	}, nil
}
