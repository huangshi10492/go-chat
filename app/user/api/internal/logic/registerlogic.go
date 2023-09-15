package logic

import (
	"context"
	"github.com/zeromicro/x/errors"
	"go-chat/app/user/rpc/userclient"

	"go-chat/app/user/api/internal/svc"
	"go-chat/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	res, err := l.svcCtx.UserRpc.Register(l.ctx,
		&userclient.RegisterRequest{
			Name:     req.Name,
			Mobile:   req.Mobile,
			Password: req.Password})
	if err != nil {
		return nil, errors.New(500, "注册失败")
	}
	return &types.RegisterResponse{
		Id:     res.Id,
		Name:   res.Name,
		Gender: res.Gender,
		Avatar: res.Avatar,
		Mobile: res.Mobile,
		Email:  res.Email,
	}, nil
}
