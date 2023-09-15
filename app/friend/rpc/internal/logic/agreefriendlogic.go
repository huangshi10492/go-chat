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

type AgreeFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAgreeFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AgreeFriendLogic {
	return &AgreeFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AgreeFriendLogic) AgreeFriend(in *friend.AgreeFriendReq) (*friend.Empty, error) {
	// 判断权限
	res, err := l.svcCtx.FriendRequestModel.FindOne(l.ctx, in.RequestId)
	if errors.Is(err, model.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "请求不存在")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	if res.FriendId != in.UserId {
		return nil, status.Error(codes.PermissionDenied, "权限不足")
	}
	// 设置申请状态
	if res.Status != 0 {
		return nil, status.Error(codes.FailedPrecondition, "重复操作")
	}
	if in.Agree {
		res.Status = 1
	} else {
		res.Status = 2
	}
	err = l.svcCtx.FriendRequestModel.Update(l.ctx, res)
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	// 好友表添加数据
	_, err = l.svcCtx.FriendModel.Insert(l.ctx, &model.Friend{
		UserId:   res.UserId,
		FriendId: res.FriendId,
		Remark:   res.Remark,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	_, err = l.svcCtx.FriendModel.Insert(l.ctx, &model.Friend{
		UserId:   res.FriendId,
		FriendId: res.UserId,
		Remark:   in.Remark,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	return &friend.Empty{}, nil
}
