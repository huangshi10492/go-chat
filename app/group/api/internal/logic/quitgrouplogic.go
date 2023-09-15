package logic

import (
	"context"
	"encoding/json"
	"go-chat/app/group/rpc/groupclient"

	"go-chat/app/group/api/internal/svc"
	"go-chat/app/group/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuitGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuitGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuitGroupLogic {
	return &QuitGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuitGroupLogic) QuitGroup(req *types.QuitGroupRequest) error {
	uid, _ := l.ctx.Value("id").(json.Number).Int64()
	_, err := l.svcCtx.GroupRpc.QuitGroup(l.ctx, &groupclient.QuitGroupRequest{
		MemberId: req.MemberId,
		UserId:   uid,
		OwnerId:  req.OwnerId,
		GroupId:  req.GroupId,
	})
	if err != nil {
		return err
	}
	return nil
}
