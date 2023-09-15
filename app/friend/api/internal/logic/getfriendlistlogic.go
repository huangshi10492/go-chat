package logic

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/x/errors"
	"go-chat/app/friend/rpc/friendclient"

	"go-chat/app/friend/api/internal/svc"
	"go-chat/app/friend/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendListLogic {
	return &GetFriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFriendListLogic) GetFriendList() (resp *types.GetFriendListResp, err error) {
	// 获取userId
	uid, _ := l.ctx.Value("id").(json.Number).Int64()
	// rpc请求
	res, err := l.svcCtx.FriendRpc.GetFriendList(l.ctx, &friendclient.GetFriendListReq{
		UserId: uid,
	})
	if err != nil {
		return nil, errors.New(500, "获取失败")
	}
	err = copier.Copy(resp, res)
	if err != nil {
		return nil, errors.New(500, "获取失败")
	}
	return
}
