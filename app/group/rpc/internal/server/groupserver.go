// Code generated by goctl. DO NOT EDIT.
// Source: group.proto

package server

import (
	"context"

	"go-chat/app/group/rpc/group"
	"go-chat/app/group/rpc/internal/logic"
	"go-chat/app/group/rpc/internal/svc"
)

type GroupServer struct {
	svcCtx *svc.ServiceContext
	group.UnimplementedGroupServer
}

func NewGroupServer(svcCtx *svc.ServiceContext) *GroupServer {
	return &GroupServer{
		svcCtx: svcCtx,
	}
}

func (s *GroupServer) CreateGroup(ctx context.Context, in *group.CreateGroupRequest) (*group.CreateGroupResponse, error) {
	l := logic.NewCreateGroupLogic(ctx, s.svcCtx)
	return l.CreateGroup(in)
}

func (s *GroupServer) JoinGroup(ctx context.Context, in *group.JoinGroupRequest) (*group.Empty, error) {
	l := logic.NewJoinGroupLogic(ctx, s.svcCtx)
	return l.JoinGroup(in)
}

func (s *GroupServer) GetJoinList(ctx context.Context, in *group.GetJoinListRequest) (*group.GetJoinListResponse, error) {
	l := logic.NewGetJoinListLogic(ctx, s.svcCtx)
	return l.GetJoinList(in)
}

func (s *GroupServer) GetGroupMemberList(ctx context.Context, in *group.GetGroupMemberListRequest) (*group.GetGroupMemberListResponse, error) {
	l := logic.NewGetGroupMemberListLogic(ctx, s.svcCtx)
	return l.GetGroupMemberList(in)
}

func (s *GroupServer) QuitGroup(ctx context.Context, in *group.QuitGroupRequest) (*group.Empty, error) {
	l := logic.NewQuitGroupLogic(ctx, s.svcCtx)
	return l.QuitGroup(in)
}

func (s *GroupServer) SetGroupMemberPermission(ctx context.Context, in *group.SetGroupMemberPermissionRequest) (*group.Empty, error) {
	l := logic.NewSetGroupMemberPermissionLogic(ctx, s.svcCtx)
	return l.SetGroupMemberPermission(in)
}

func (s *GroupServer) ModifyRemark(ctx context.Context, in *group.ModifyRemarkRequest) (*group.Empty, error) {
	l := logic.NewModifyRemarkLogic(ctx, s.svcCtx)
	return l.ModifyRemark(in)
}

func (s *GroupServer) DissolveGroup(ctx context.Context, in *group.DissolveGroupRequest) (*group.Empty, error) {
	l := logic.NewDissolveGroupLogic(ctx, s.svcCtx)
	return l.DissolveGroup(in)
}

func (s *GroupServer) GetGroupInfo(ctx context.Context, in *group.GetGroupInfoRequest) (*group.GetGroupInfoResponse, error) {
	l := logic.NewGetGroupInfoLogic(ctx, s.svcCtx)
	return l.GetGroupInfo(in)
}
