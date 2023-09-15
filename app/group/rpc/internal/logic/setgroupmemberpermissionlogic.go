package logic

import (
	"context"
	"errors"
	"go-chat/app/group/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go-chat/app/group/rpc/group"
	"go-chat/app/group/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetGroupMemberPermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetGroupMemberPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetGroupMemberPermissionLogic {
	return &SetGroupMemberPermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetGroupMemberPermissionLogic) SetGroupMemberPermission(in *group.SetGroupMemberPermissionRequest) (*group.Empty, error) {
	if in.Permission > 3 {
		return nil, status.Error(codes.InvalidArgument, "参数错误")
	}
	res1, err := l.svcCtx.UserGroupModel.FindOneByUserIdGroupId(l.ctx, in.MemberId, in.GroupId)
	if errors.Is(err, model.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "用户不在该群中")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	res2, err := l.svcCtx.UserGroupModel.FindOneByUserIdGroupId(l.ctx, in.UserId, in.GroupId)
	if errors.Is(err, model.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "用户不在该群中")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	if res1.Permission <= res2.Permission || res2.Permission > 1 {
		return nil, status.Error(codes.PermissionDenied, "权限不足")
	}
	res1.Permission = in.Permission
	err = l.svcCtx.UserGroupModel.Update(l.ctx, res1)
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	return &group.Empty{}, nil
}
