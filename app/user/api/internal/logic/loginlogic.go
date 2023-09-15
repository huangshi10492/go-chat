package logic

import (
	"context"
	"github.com/zeromicro/x/errors"
	"go-chat/app/user/rpc/userclient"
	"go-chat/pkg/jwt"
	"time"

	"go-chat/app/user/api/internal/svc"
	"go-chat/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	res, err := l.svcCtx.UserRpc.Login(l.ctx,
		&userclient.LoginRequest{
			Mobile:   req.Mobile,
			Password: req.Password,
		})
	if err != nil {
		return nil, errors.New(500, "登录失败")
	}
	token, err := jwt.GetJwtToken(l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(), l.svcCtx.Config.Auth.AccessExpire, res.Id)
	if err != nil {
		return nil, errors.New(500, "登录失败")
	}
	return &types.LoginResponse{
		AccessToken: token,
	}, nil
}
