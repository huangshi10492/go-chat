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

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	// 检查用户是否存在
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if errors.Is(err, model.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "用户或密码错误")
	}
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}
	return &user.UserInfoResponse{
		Id:     res.Id,
		Name:   res.Name,
		Gender: res.Gender,
		Avatar: res.Avatar,
		Mobile: res.Mobile,
		Email:  res.Email,
	}, nil
}
