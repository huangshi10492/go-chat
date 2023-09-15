package logic

import (
	"context"
	"go-chat/app/friend/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go-chat/app/friend/rpc/friend"
	"go-chat/app/friend/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendFriendLogic {
	return &SendFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendFriendLogic) SendFriend(in *friend.SendFriendReq) (*friend.Empty, error) {
	// 将申请持久化
	_, err := l.svcCtx.FriendRequestModel.Insert(l.ctx, &model.FriendRequest{
		UserId:   in.UserId,
		FriendId: in.FriendId,
		Remark:   in.Remark,
		Reason:   in.Reason,
		Status:   0,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "保存好友申请失败")
	}
	// todo 发送更新好友申请列表的消息
	return &friend.Empty{}, nil
}
