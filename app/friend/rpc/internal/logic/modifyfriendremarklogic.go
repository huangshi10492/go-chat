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

type ModifyFriendRemarkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyFriendRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyFriendRemarkLogic {
	return &ModifyFriendRemarkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifyFriendRemarkLogic) ModifyFriendRemark(in *friend.ModifyFriendRemarkReq) (*friend.Empty, error) {
	// 根据userId和friendId修改对应的数据
	res, err := l.svcCtx.FriendModel.FindOneByUserIdFriendId(l.ctx, in.UserId, in.FriendId)
	if errors.Is(err, model.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "未找到好友")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	res.Remark = in.Remark
	err = l.svcCtx.FriendModel.Update(l.ctx, res)
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	return &friend.Empty{}, nil
}
