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

type QuitGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuitGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuitGroupLogic {
	return &QuitGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QuitGroupLogic) QuitGroup(in *group.QuitGroupRequest) (*group.Empty, error) {
	res1, err := l.svcCtx.UserGroupModel.FindOneByUserIdGroupId(l.ctx, in.GroupId, in.MemberId)
	if errors.Is(err, model.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "指定退出的用户不在群中")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	if in.UserId == in.MemberId {
		if res1.Permission == 0 {
			res2, err := l.svcCtx.UserGroupModel.FindOneByUserIdGroupId(l.ctx, in.GroupId, in.OwnerId)
			if errors.Is(err, model.ErrNotFound) {
				return nil, status.Error(codes.NotFound, "指定群主的用户不在群中")
			}
			if err != nil {
				return nil, status.Error(codes.Internal, "数据库错误")
			}
			err = l.svcCtx.UserGroupModel.SwitchOwner(l.ctx, res1.Id, res2.Id)
			if err != nil {
				return nil, status.Error(codes.Internal, "数据库错误")
			}
		} else {
			err = l.svcCtx.UserGroupModel.Delete(l.ctx, res1.Id)
			if err != nil {
				return nil, status.Error(codes.Internal, "数据库错误")
			}
		}
	} else {
		res3, err := l.svcCtx.UserGroupModel.FindOneByUserIdGroupId(l.ctx, in.GroupId, in.UserId)
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "操作用户不在群中")
		}
		if err != nil {
			return nil, status.Error(codes.Internal, "数据库错误")
		}
		if res3.Permission >= res1.Permission {
			return nil, status.Error(codes.PermissionDenied, "权限不足")
		}
		err = l.svcCtx.UserGroupModel.Delete(l.ctx, res1.Id)
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "指定退出的用户不在群中")
		}
		if err != nil {
			return nil, status.Error(codes.Internal, "数据库错误")
		}
	}
	return &group.Empty{}, nil
}
