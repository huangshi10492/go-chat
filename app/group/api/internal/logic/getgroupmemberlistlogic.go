package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"go-chat/app/group/rpc/groupclient"

	"go-chat/app/group/api/internal/svc"
	"go-chat/app/group/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupMemberListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGroupMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupMemberListLogic {
	return &GetGroupMemberListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupMemberListLogic) GetGroupMemberList(req *types.GetGroupMemberListRequest) (resp *types.GetGroupMemberListResponse, err error) {
	res, err := l.svcCtx.GroupRpc.GetGroupMemberList(l.ctx, &groupclient.GetGroupMemberListRequest{
		GroupId: req.GroupId,
	})
	if err != nil {
		return nil, err
	}
	var t []types.GroupMemberList
	err = copier.Copy(&t, res.MemberList)
	resp = &types.GetGroupMemberListResponse{
		MemberList: t,
	}
	return
}
