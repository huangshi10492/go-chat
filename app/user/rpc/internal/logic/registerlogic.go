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

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// 判断用户是否已经注册
	_, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if !errors.Is(err, model.ErrNotFound) {
		return nil, status.Error(codes.AlreadyExists, "手机号码已经注册")
	}
	newUser := model.User{
		Name:     in.Name,
		Mobile:   in.Mobile,
		Password: in.Password,
	}
	res, err := l.svcCtx.UserModel.Insert(l.ctx, &newUser)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "内部错误: %v", err)
	}
	newUser.Id, err = res.LastInsertId()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "内部错误: %v", err)
	}
	return &user.RegisterResponse{
		Id:     newUser.Id,
		Name:   newUser.Name,
		Gender: newUser.Gender,
		Avatar: newUser.Avatar,
		Mobile: newUser.Mobile,
		Email:  newUser.Email,
	}, nil
}
