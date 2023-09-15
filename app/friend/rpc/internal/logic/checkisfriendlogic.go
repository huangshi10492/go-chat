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

type CheckIsFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckIsFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckIsFriendLogic {
	return &CheckIsFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckIsFriendLogic) CheckIsFriend(in *friend.CheckIsFriendRequest) (*friend.CheckIsFriendResponse, error) {
	_, err := l.svcCtx.FriendModel.FindOneByUserIdFriendId(l.ctx, in.UserId, in.FriendId)
	if errors.Is(err, model.ErrNotFound) {
		return &friend.CheckIsFriendResponse{
			IsFriend: false,
		}, nil
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "内部错误")
	}
	return &friend.CheckIsFriendResponse{
		IsFriend: true,
	}, nil
}
