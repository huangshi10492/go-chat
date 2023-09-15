package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go-chat/app/group/rpc/group"
	"go-chat/app/group/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DissolveGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDissolveGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DissolveGroupLogic {
	return &DissolveGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DissolveGroupLogic) DissolveGroup(in *group.DissolveGroupRequest) (*group.Empty, error) {
	res1, err := l.svcCtx.UserGroupModel.FindOneByUserIdGroupId(l.ctx, in.UserId, in.GroupId)
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	if res1.Permission != 3 {
		return nil, status.Error(codes.PermissionDenied, "权限不足")
	}
	res2, err := l.svcCtx.UserGroupModel.FindUserIdByGroupId(l.ctx, in.GroupId)
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	userIds := make([]int64, len(res2))
	for i, t := range res2 {
		userIds[i] = t.UserId
	}
	err = l.svcCtx.UserGroupModel.DissolveGroup(l.ctx, userIds, in.GroupId)
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	err = l.svcCtx.GroupModel.Delete(l.ctx, in.GroupId)
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	return &group.Empty{}, nil
}
