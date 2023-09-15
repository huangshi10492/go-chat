package logic

import (
	"context"
	"go-chat/app/friend/rpc/friendclient"
	"go-chat/app/group/rpc/groupclient"
	"go-chat/app/message/model"
	"go-chat/pkg/protobuf/mq_pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"time"

	"go-chat/app/message/rpc/internal/svc"
	"go-chat/app/message/rpc/message"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageHandlerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMessageHandlerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageHandlerLogic {
	return &MessageHandlerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MessageHandlerLogic) MessageHandler(in *message.MessageHandlerRequest) (*message.Empty, error) {
	// 获取需要写入消息的用户的列表
	var ids []int64
	switch in.SenderType {
	case 1:
		res, err := l.svcCtx.FriendRpc.CheckIsFriend(l.ctx, &friendclient.CheckIsFriendRequest{
			UserId:   in.SenderId,
			FriendId: in.ReceiverId,
		})
		if err != nil {
			return nil, err
		}
		if !res.IsFriend {
			return nil, status.Error(codes.NotFound, "好友不存在")
		}
		ids = []int64{in.SenderId, in.ReceiverId}
		break
	case 2:
		res, err := l.svcCtx.GroupRpc.GetGroupMemberList(l.ctx, &groupclient.GetGroupMemberListRequest{
			GroupId: in.GroupId,
		})
		if err != nil {
			return nil, err
		}
		if !isInArray(in.SenderId, res.MemberList) {
			return nil, status.Error(codes.NotFound, "未加入群")
		}
		for _, t := range res.MemberList {
			ids = append(ids, t.UserId)
		}
	}
	// 消息持久化
	messageId := primitive.NewObjectID()
	err := l.svcCtx.MessageModel.Insert(l.ctx, &model.Message{
		ID:          messageId,
		MessageType: in.ContentType,
		MessageData: in.Content,
		GroupId:     in.GroupId,
		SenderId:    in.SenderId,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	var modelData []any
	var notifyData []*mq_pb.Mailbox
	for _, id := range ids {
		seq, err := getSeq(l.ctx, l.svcCtx.Redis, id)
		if err != nil {
			logx.Error(err)
		}
		modelData = append(modelData, &model.Mailbox{
			SenderId:    in.SenderId,
			ReceiverId:  id,
			ContentType: 1,
			Content:     messageId.Hex(),
			Seq:         seq,
		})
		notifyData = append(notifyData, &mq_pb.Mailbox{
			ReceiverId: id,
			Seq:        seq,
		})
	}
	err = l.svcCtx.MailboxModel.InsertMany(l.ctx, modelData)
	if err != nil {
		logx.Error(err)
	}
	// 消息推送
	notify := &mq_pb.Notify{
		NotifyType: 1,
		MailBox:    notifyData,
		Message: &mq_pb.Message{
			ID:          messageId.Hex(),
			GroupId:     in.GroupId,
			MessageType: in.ContentType,
			MessageData: in.Content,
		},
		Common: &mq_pb.NotifyCommon{
			SenderId:    in.SenderId,
			ContentType: 1,
			Content:     messageId.Hex(),
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

func isInArray(target int64, array []*groupclient.GroupMemberList) bool {
	for _, t := range array {
		if t.UserId == target {
			return true
		}
	}
	return false
}
