package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go-chat/app/group/rpc/group"
	"go-chat/app/group/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetJoinListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetJoinListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetJoinListLogic {
	return &GetJoinListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetJoinListLogic) GetJoinList(in *group.GetJoinListRequest) (*group.GetJoinListResponse, error) {
	res, err := l.svcCtx.UserGroupModel.GetJoinList(l.ctx, in.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	var resp []*group.GroupList
	err = copier.Copy(&resp, res)
	return &group.GetJoinListResponse{
		GroupList: resp,
	}, nil
}
