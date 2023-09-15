package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"go-chat/app/group/rpc/groupclient"

	"go-chat/app/group/api/internal/svc"
	"go-chat/app/group/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupInfoLogic {
	return &GetGroupInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupInfoLogic) GetGroupInfo(req *types.GetGroupInfoRequest) (resp *types.GetGroupInfoResponse, err error) {
	res, err := l.svcCtx.GroupRpc.GetGroupInfo(l.ctx, &groupclient.GetGroupInfoRequest{
		GroupId: req.GroupId,
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&resp, res)
	if err != nil {
		return nil, err
	}
	return
}
