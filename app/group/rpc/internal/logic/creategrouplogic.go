package logic

import (
	"context"
	"go-chat/app/group/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go-chat/app/group/rpc/group"
	"go-chat/app/group/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateGroupLogic) CreateGroup(in *group.CreateGroupRequest) (*group.CreateGroupResponse, error) {
	res, err := l.svcCtx.GroupModel.Insert(l.ctx, &model.Group{
		Name: in.Name,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	groupId, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	res, err = l.svcCtx.UserGroupModel.Insert(l.ctx, &model.UserGroup{
		UserId:     in.UserId,
		GroupId:    groupId,
		Permission: 3,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	return &group.CreateGroupResponse{
		Id:   groupId,
		Name: in.Name,
	}, nil
}
