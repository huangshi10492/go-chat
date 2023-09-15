package mqs

import (
	"context"
	"go-chat/app/connect/internal/svc"
	"go-chat/pkg/protobuf/mq_pb"
	"go-chat/pkg/protobuf/ws_pb"
	"google.golang.org/protobuf/proto"
)

type MessageQueue struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageQueue(ctx context.Context, svcCtx *svc.ServiceContext) *MessageQueue {
	return &MessageQueue{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageQueue) Consume(_, val string) error {
	// 消息处理
	res := &mq_pb.Notify{}
	err := proto.Unmarshal([]byte(val), res)
	if err != nil {
		return err
	}
	mailbox := &ws_pb.Mailbox{
		SenderId:    res.Common.SenderId,
		ContentType: res.Common.ContentType,
		Content:     res.Common.Content,
		CreateAt:    res.Common.CreateAt,
		ClientId:    res.Common.ClientId,
	}
	switch res.NotifyType {
	case 1:
		// 发送消息
		resp := &ws_pb.MessageResponse{
			Message: &ws_pb.Message{
				ID:          res.Message.ID,
				GroupId:     res.Message.GroupId,
				MessageType: res.Message.MessageType,
				MessageData: res.Message.MessageData,
			},
			Mailbox: mailbox,
		}
		for _, box := range res.MailBox {
			clients := l.svcCtx.WebSocket.GetClients(box.ReceiverId)
			if len(clients) == 0 {
				continue
			}
			resp.Mailbox.Seq = box.Seq
			resp.Mailbox.ReceiverId = box.ReceiverId
			body, err := proto.Marshal(&ws_pb.WsResponse{
				Type: 1,
				Data: &ws_pb.WsResponse_MessageResponse{MessageResponse: resp},
			})
			if err != nil {
				continue
			}
			for _, client := range clients {
				err = client.Send(body)
			}
		}

	case 2:
		// 已读
		resp := &ws_pb.MailboxResponse{
			Mailbox: mailbox,
		}
		for _, box := range res.MailBox {
			clients := l.svcCtx.WebSocket.GetClients(box.ReceiverId)
			if len(clients) == 0 {
				continue
			}
			resp.Mailbox.Seq = box.Seq
			resp.Mailbox.ReceiverId = box.ReceiverId
			body, err := proto.Marshal(&ws_pb.WsResponse{
				Type: 2,
				Data: &ws_pb.WsResponse_MailboxResponse{MailboxResponse: resp},
			})
			if err != nil {
				continue
			}
			for _, client := range clients {
				err = client.Send(body)
			}
		}

	case 3:
		// 撤回
		resp := &ws_pb.MailboxResponse{
			Mailbox: mailbox,
		}
		for _, box := range res.MailBox {
			clients := l.svcCtx.WebSocket.GetClients(box.ReceiverId)
			if len(clients) == 0 {
				continue
			}
			resp.Mailbox.Seq = box.Seq
			resp.Mailbox.ReceiverId = box.ReceiverId
			body, err := proto.Marshal(&ws_pb.WsResponse{
				Type: 2,
				Data: &ws_pb.WsResponse_MailboxResponse{MailboxResponse: resp},
			})
			if err != nil {
				continue
			}
			for _, client := range clients {
				err = client.Send(body)
			}
		}
	}
	return nil
}
