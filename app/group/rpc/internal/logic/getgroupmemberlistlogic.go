package logic

import (
	"context"
	"go-chat/app/user/rpc/userclient"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go-chat/app/group/rpc/group"
	"go-chat/app/group/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupMemberListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupMemberListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupMemberListLogic {
	return &GetGroupMemberListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupMemberListLogic) GetGroupMemberList(in *group.GetGroupMemberListRequest) (*group.GetGroupMemberListResponse, error) {
	res1, err := l.svcCtx.UserGroupModel.FindUserIdByGroupId(l.ctx, in.GroupId)
	if err != nil {
		return nil, status.Error(codes.Internal, "数据库错误")
	}
	userIds := make([]int64, len(res1))
	for i, t := range res1 {
		userIds[i] = t.UserId
	}
	res2, err := l.svcCtx.UserRpc.UsersInfo(l.ctx, &userclient.UsersInfoRequest{Ids: userIds})
	if err != nil {
		return nil, status.Error(codes.Internal, "服务错误")
	}
	mergedMap := make(map[int64]group.GroupMemberList)
	var resp []*group.GroupMemberList
	for _, p := range res1 {
		mergedMap[p.UserId] = group.GroupMemberList{
			UserId:     p.UserId,
			Permission: p.Permission,
		}
	}
	for _, s := range res2.UsersInfo {
		t := group.GroupMemberList{
			UserId:     s.Id,
			Permission: mergedMap[s.Id].Permission,
			Name:       s.Name,
		}
		resp = append(resp, &t)
	}
	return &group.GetGroupMemberListResponse{
		MemberList: resp,
	}, nil
}
