package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-chat/app/message/model"
	"go-chat/app/message/rpc/internal/svc"
	"go-chat/app/message/rpc/message"
	"go-chat/pkg/protobuf/mq_pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"time"
)

type ReadHandlerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReadHandlerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadHandlerLogic {
	return &ReadHandlerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReadHandlerLogic) ReadHandler(in *message.ReadHandlerRequest) (*message.Empty, error) {
	res, err := l.svcCtx.MessageModel.FindOne(l.ctx, in.MessageId)
	if err != nil {
		return nil, err
	}
	ids := []int64{res.SenderId, in.SenderId}
	var notifyData []*mq_pb.Mailbox
	for _, id := range ids {
		seq, err := getSeq(l.ctx, l.svcCtx.Redis, id)
		if err != nil {
			return nil, err
		}
		err = l.svcCtx.MailboxModel.Insert(l.ctx, &model.Mailbox{
			SenderId:    in.SenderId,
			ReceiverId:  id,
			ContentType: 2,
			Content:     in.MessageId,
			Seq:         seq,
		})
		if err != nil {
			logx.Error(err)
		}
		notifyData = append(notifyData, &mq_pb.Mailbox{
			ReceiverId: id,
			Seq:        seq,
		})
	}
	// 消息推送
	notify := &mq_pb.Notify{
		NotifyType: 2,
		MailBox:    notifyData,
		Common: &mq_pb.NotifyCommon{
			SenderId:    in.SenderId,
			ContentType: 1,
			Content:     in.MessageId,
			CreateAt:    time.Now().Unix(),
			ClientId:    in.ClientId,
		},
	}
	data, err := proto.Marshal(notify)
	if err != nil {
		return nil, status.Error(codes.Internal, "序列化错误")
	}
	err = l.svcCtx.KqPusherClient.Push(string(data))
	if err != nil {
		logx.Error(err)
	}
	return &message.Empty{}, nil
}
