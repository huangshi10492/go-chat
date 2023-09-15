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

type JoinGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJoinGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinGroupLogic {
	return &JoinGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JoinGroupLogic) JoinGroup(in *group.JoinGroupRequest) (*group.Empty, error) {
	_, err := l.svcCtx.UserGroupModel.FindOneByUserIdGroupId(l.ctx, in.UserId, in.GroupId)
	if !errors.Is(err, model.ErrNotFound) {
		return nil, status.Error(codes.AlreadyExists, "已经加入了该群")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	_, err = l.svcCtx.GroupModel.FindOne(l.ctx, in.GroupId)
	if errors.Is(err, model.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "该群不存在")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	_, err = l.svcCtx.UserGroupModel.Insert(l.ctx, &model.UserGroup{
		UserId:  in.UserId,
		GroupId: in.GroupId,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	return &group.Empty{}, nil
}
