package logic

import (
	"context"
	"go-chat/app/user/rpc/userclient"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go-chat/app/friend/rpc/friend"
	"go-chat/app/friend/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendRequestListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendRequestListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendRequestListLogic {
	return &GetFriendRequestListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendRequestListLogic) GetFriendRequestList(in *friend.GetFriendRequestListReq) (*friend.GetFriendRequestListResp, error) {
	// 读取申请列表数据
	res, err := l.svcCtx.FriendRequestModel.GetFriendRequestList(l.ctx, in.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, "获取申请列表失败")
	}
	var list []*friend.FriendRequestInfo
	for _, r := range res {
		res2, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &userclient.UserInfoRequest{
			Id: r.UserId,
		})
		list = append(list, &friend.FriendRequestInfo{
			RequestId: r.Id,
			FriendId:  r.UserId,
			Name:      res2.Name,
			Reason:    r.Reason,
			Status:    r.Status,
		})
	}
	return &friend.GetFriendRequestListResp{FriendRequestList: list}, nil
}
