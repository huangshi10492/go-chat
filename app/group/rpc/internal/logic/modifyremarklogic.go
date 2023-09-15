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

type ModifyRemarkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewModifyRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyRemarkLogic {
	return &ModifyRemarkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ModifyRemarkLogic) ModifyRemark(in *group.ModifyRemarkRequest) (*group.Empty, error) {
	_, err := l.svcCtx.UserGroupModel.FindOneByUserIdGroupId(l.ctx, in.UserId, in.GroupId)
	if errors.Is(err, model.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "未加入该群")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	err = l.svcCtx.UserGroupModel.Update(l.ctx, &model.UserGroup{
		UserId:  in.UserId,
		GroupId: in.GroupId,
		Remark:  in.Remark,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	return &group.Empty{}, nil
}
