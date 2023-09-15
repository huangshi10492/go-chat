package logic

import (
	"context"
	"encoding/json"
	"go-chat/app/group/rpc/groupclient"

	"go-chat/app/group/api/internal/svc"
	"go-chat/app/group/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateGroupLogic) CreateGroup(req *types.CreateGroupRequest) (resp *types.CreateGroupResponse, err error) {
	uid, _ := l.ctx.Value("id").(json.Number).Int64()
	res, err := l.svcCtx.GroupRpc.CreateGroup(l.ctx, &groupclient.CreateGroupRequest{
		UserId: uid,
		Name:   req.Name,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.CreateGroupResponse{
		Id:   res.Id,
		Name: res.Name,
	}
	return
}
