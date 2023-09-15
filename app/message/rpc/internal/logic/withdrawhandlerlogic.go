package logic

import (
	"context"
	"go-chat/app/friend/rpc/friendclient"
	"go-chat/app/group/rpc/groupclient"
	"go-chat/app/message/model"
	"go-chat/pkg/protobuf/mq_pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"time"

	"go-chat/app/message/rpc/internal/svc"
	"go-chat/app/message/rpc/message"

	"github.com/zeromicro/go-zero/core/logx"
)

type WithdrawHandlerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWithdrawHandlerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WithdrawHandlerLogic {
	return &WithdrawHandlerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WithdrawHandlerLogic) WithdrawHandler(in *message.WithdrawHandlerRequest) (*message.Empty, error) {
	res, err := l.svcCtx.MessageModel.FindOne(l.ctx, in.MessageId)
	if err != nil {
		return nil, err
	}
	if res.SenderId != in.SenderId {
		return nil, status.Error(codes.PermissionDenied, "只有发送者才能撤回消息")
	}
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
			ContentType: 3,
			Content:     in.MessageId,
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
		NotifyType: 3,
		MailBox:    notifyData,
		Common: &mq_pb.NotifyCommon{
			SenderId:    in.SenderId,
			ContentType: 3,
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
