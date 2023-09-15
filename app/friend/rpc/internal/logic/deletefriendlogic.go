package logic

import (
	"context"
	"errors"
	"go-chat/app/friend/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go-chat/app/friend/rpc/friend"
	"go-chat/app/friend/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFriendLogic {
	return &DeleteFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteFriendLogic) DeleteFriend(in *friend.DeleteFriendReq) (*friend.Empty, error) {
	// todo 好友表删除数据(删2个)
	_, err := l.svcCtx.FriendModel.FindOneByUserIdFriendId(l.ctx, in.UserId, in.FriendId)
	if errors.Is(err, model.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "记录不存在")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	_, err = l.svcCtx.FriendModel.FindOneByUserIdFriendId(l.ctx, in.FriendId, in.UserId)
	if errors.Is(err, model.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "记录不存在")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	err = l.svcCtx.FriendModel.DeleteFriend(l.ctx, in.UserId, in.FriendId)
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	return &friend.Empty{}, nil
}
