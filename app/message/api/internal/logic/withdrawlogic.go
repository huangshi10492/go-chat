package logic

import (
	"context"
	"encoding/json"
	"go-chat/app/message/rpc/messageclient"

	"go-chat/app/message/api/internal/svc"
	"go-chat/app/message/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WithdrawLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWithdrawLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WithdrawLogic {
	return &WithdrawLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WithdrawLogic) Withdraw(req *types.WithdrawReq) error {
	uid, _ := l.ctx.Value("id").(json.Number).Int64()
	_, err := l.svcCtx.MessageRpc.WithdrawHandler(l.ctx, &messageclient.WithdrawHandlerRequest{
		SenderId:   uid,
		SenderType: req.SenderType,
		GroupId:    req.GroupId,
		ReceiverId: req.ReceiverId,
		MessageId:  req.MessageId,
		ClientId:   req.ClientId,
	})
	return err
}
