package logic

import (
	"context"
	"encoding/json"
	"go-chat/app/friend/rpc/friendclient"

	"go-chat/app/friend/api/internal/svc"
	"go-chat/app/friend/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyFriendRemarkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyFriendRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyFriendRemarkLogic {
	return &ModifyFriendRemarkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyFriendRemarkLogic) ModifyFriendRemark(req *types.ModifyFriendRemarkReq) error {
	// 获取userId
	uid, _ := l.ctx.Value("id").(json.Number).Int64()
	// rpc请求
	_, err := l.svcCtx.FriendRpc.ModifyFriendRemark(l.ctx, &friendclient.ModifyFriendRemarkReq{
		UserId:   uid,
		FriendId: req.FriendId,
		Remark:   req.Remark,
	})
	return err
}
