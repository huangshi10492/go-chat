package logic

import (
	"context"
	"encoding/json"
	"go-chat/app/message/rpc/messageclient"

	"go-chat/app/message/api/internal/svc"
	"go-chat/app/message/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadLogic {
	return &ReadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReadLogic) Read(req *types.ReadReq) error {
	uid, _ := l.ctx.Value("id").(json.Number).Int64()
	_, err := l.svcCtx.MessageRpc.ReadHandler(l.ctx, &messageclient.ReadHandlerRequest{
		SenderId:  uid,
		MessageId: req.MessageId,
		ClientId:  req.ClientId,
	})
	return err
}
