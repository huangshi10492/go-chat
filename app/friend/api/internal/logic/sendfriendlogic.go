package logic

import (
	"context"
	"encoding/json"
	"go-chat/app/friend/rpc/friendclient"

	"go-chat/app/friend/api/internal/svc"
	"go-chat/app/friend/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendFriendLogic {
	return &SendFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendFriendLogic) SendFriend(req *types.SendFriendReq) error {
	// 获取userId
	uid, _ := l.ctx.Value("id").(json.Number).Int64()
	// rpc请求
	_, err := l.svcCtx.FriendRpc.SendFriend(l.ctx, &friendclient.SendFriendReq{
		UserId:   uid,
		FriendId: req.FriendId,
		Reason:   req.Reason,
		Remark:   req.Remark,
	})
	return err
}
