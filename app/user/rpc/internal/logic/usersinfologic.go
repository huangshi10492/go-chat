package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go-chat/app/user/rpc/internal/svc"
	"go-chat/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UsersInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUsersInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsersInfoLogic {
	return &UsersInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UsersInfoLogic) UsersInfo(in *user.UsersInfoRequest) (*user.UsersInfoResponse, error) {
	res, err := l.svcCtx.UserModel.UsersInfo(l.ctx, in.Ids)
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	var resp []*user.UserInfoResponse
	err = copier.Copy(resp, res)
	if err != nil {
		return nil, status.Error(codes.Internal, "内部错误")
	}
	return &user.UsersInfoResponse{
		UsersInfo: resp,
	}, nil
}
