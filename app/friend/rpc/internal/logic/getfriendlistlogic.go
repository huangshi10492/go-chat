package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go-chat/app/friend/rpc/friend"
	"go-chat/app/friend/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendListLogic {
	return &GetFriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendListLogic) GetFriendList(in *friend.GetFriendListReq) (*friend.GetFriendListResp, error) {
	// 好友表读取数据
	res, err := l.svcCtx.FriendModel.GetFriendList(l.ctx, in.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, "获取好友列表失败")
	}
	var list []*friend.FriendInfo
	err = copier.Copy(list, res)
	if err != nil {
		return nil, status.Error(codes.Internal, "获取好友列表失败")
	}
	return &friend.GetFriendListResp{
		FriendList: list,
	}, nil
}
