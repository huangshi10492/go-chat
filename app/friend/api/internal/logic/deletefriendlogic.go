package logic

import (
	"context"
	"encoding/json"
	"go-chat/app/friend/rpc/friendclient"

	"go-chat/app/friend/api/internal/svc"
	"go-chat/app/friend/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFriendLogic {
	return &DeleteFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFriendLogic) DeleteFriend(req *types.DeleteFriendReq) error {
	// 获取userId
	uid, _ := l.ctx.Value("id").(json.Number).Int64()
	// rpc请求
	_, err := l.svcCtx.FriendRpc.DeleteFriend(l.ctx, &friendclient.DeleteFriendReq{
		UserId:   uid,
		FriendId: req.FriendId,
	})
	return err
}
