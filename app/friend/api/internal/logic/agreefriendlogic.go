package logic

import (
	"context"
	"encoding/json"
	"go-chat/app/friend/rpc/friendclient"

	"go-chat/app/friend/api/internal/svc"
	"go-chat/app/friend/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgreeFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgreeFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AgreeFriendLogic {
	return &AgreeFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgreeFriendLogic) AgreeFriend(req *types.AgreeFriendReq) error {
	// 获取userId
	uid, _ := l.ctx.Value("id").(json.Number).Int64()
	// rpc请求
	_, err := l.svcCtx.FriendRpc.AgreeFriend(l.ctx, &friendclient.AgreeFriendReq{
		RequestId: req.RequestId,
		UserId:    uid,
		Remark:    req.Remark,
		Agree:     req.Agree,
	})
	return err
}
