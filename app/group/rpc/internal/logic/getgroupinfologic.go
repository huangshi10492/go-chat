package logic

import (
	"context"

	"go-chat/app/group/rpc/group"
	"go-chat/app/group/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupInfoLogic {
	return &GetGroupInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupInfoLogic) GetGroupInfo(in *group.GetGroupInfoRequest) (*group.GetGroupInfoResponse, error) {
	res, err := l.svcCtx.GroupModel.FindOne(l.ctx, in.GroupId)
	if err != nil {
		return nil, err
	}
	return &group.GetGroupInfoResponse{
		Id:   res.Id,
		Name: res.Name,
	}, nil
}
