package logic

import (
	"context"
	"encoding/json"
	"go-chat/app/group/rpc/groupclient"

	"go-chat/app/group/api/internal/svc"
	"go-chat/app/group/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetGroupMemberPermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetGroupMemberPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetGroupMemberPermissionLogic {
	return &SetGroupMemberPermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetGroupMemberPermissionLogic) SetGroupMemberPermission(req *types.SetGroupMemberPermissionRequest) error {
	uid, _ := l.ctx.Value("id").(json.Number).Int64()
	_, err := l.svcCtx.GroupRpc.SetGroupMemberPermission(l.ctx, &groupclient.SetGroupMemberPermissionRequest{
		GroupId:    req.GroupId,
		UserId:     uid,
		Permission: req.Permission,
		MemberId:   req.MemberId,
	})
	if err != nil {
		return err
	}
	return nil
}
