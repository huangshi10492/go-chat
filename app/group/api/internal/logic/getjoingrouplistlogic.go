package logic

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"go-chat/app/group/rpc/groupclient"

	"go-chat/app/group/api/internal/svc"
	"go-chat/app/group/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetJoinGroupListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetJoinGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetJoinGroupListLogic {
	return &GetJoinGroupListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetJoinGroupListLogic) GetJoinGroupList() (resp *types.GetJoinListResponse, err error) {
	uid, _ := l.ctx.Value("id").(json.Number).Int64()
	res, err := l.svcCtx.GroupRpc.GetJoinList(l.ctx, &groupclient.GetJoinListRequest{
		UserId: uid,
	})
	var t []types.GroupList
	err = copier.Copy(&t, res.GroupList)
	if err != nil {
		return nil, err
	}
	return
}
