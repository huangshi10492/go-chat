package logic

import (
	"context"
	"encoding/json"
	"go-chat/app/message/rpc/messageclient"

	"go-chat/app/message/api/internal/svc"
	"go-chat/app/message/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendMessageLogic) SendMessage(req *types.SendMessageReq) error {
	uid, _ := l.ctx.Value("id").(json.Number).Int64()
	_, err := l.svcCtx.MessageRpc.MessageHandler(l.ctx, &messageclient.MessageHandlerRequest{
		SenderId:    uid,
		SenderType:  req.SenderType,
		GroupId:     req.GroupId,
		ReceiverId:  req.ReceiverId,
		ContentType: req.ContentType,
		Content:     req.Content,
		ClientId:    req.ClientId,
	})
	return err
}
