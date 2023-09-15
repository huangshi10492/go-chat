package logic

import (
	"context"
	"encoding/json"
	"go-chat/app/group/rpc/groupclient"

	"go-chat/app/group/api/internal/svc"
	"go-chat/app/group/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyRemarkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyRemarkLogic {
	return &ModifyRemarkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyRemarkLogic) ModifyRemark(req *types.ModifyRemarkRequest) error {
	uid, _ := l.ctx.Value("id").(json.Number).Int64()
	_, err := l.svcCtx.GroupRpc.ModifyRemark(l.ctx, &groupclient.ModifyRemarkRequest{
		UserId:  uid,
		GroupId: req.GroupId,
		Remark:  req.Remark,
	})
	if err != nil {
		return err
	}
	return nil
}
