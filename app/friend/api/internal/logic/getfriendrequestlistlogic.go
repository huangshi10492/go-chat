package logic

import (
	"context"
	"encoding/json"
	"go-chat/app/friend/rpc/friendclient"

	"go-chat/app/friend/api/internal/svc"
	"go-chat/app/friend/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendRequestListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFriendRequestListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendRequestListLogic {
	return &GetFriendRequestListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFriendRequestListLogic) GetFriendRequestList() (*types.GetFriendRequestListResp, error) {
	// 获取userId
	uid, _ := l.ctx.Value("id").(json.Number).Int64()
	// rpc请求
	res, err := l.svcCtx.FriendRpc.GetFriendRequestList(l.ctx, &friendclient.GetFriendRequestListReq{
		UserId: uid,
	})
	if err != nil {
		return nil, err
	}
	var t []types.FriendRequestInfo
	for _, r := range res.FriendRequestList {
		t = append(t, types.FriendRequestInfo{
			RequestId: r.RequestId,
			FriendId:  r.FriendId,
			Name:      r.Name,
			Reason:    r.Reason,
			Status:    r.Status,
		})
	}
	return &types.GetFriendRequestListResp{FriendRequestList: t}, nil
}
