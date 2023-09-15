// Code generated by goctl. DO NOT EDIT.
// Source: group.proto

package groupclient

import (
	"context"

	"go-chat/app/group/rpc/group"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateGroupRequest              = group.CreateGroupRequest
	CreateGroupResponse             = group.CreateGroupResponse
	DissolveGroupRequest            = group.DissolveGroupRequest
	Empty                           = group.Empty
	GetGroupInfoRequest             = group.GetGroupInfoRequest
	GetGroupInfoResponse            = group.GetGroupInfoResponse
	GetGroupMemberListRequest       = group.GetGroupMemberListRequest
	GetGroupMemberListResponse      = group.GetGroupMemberListResponse
	GetJoinListRequest              = group.GetJoinListRequest
	GetJoinListResponse             = group.GetJoinListResponse
	GroupList                       = group.GroupList
	GroupMemberList                 = group.GroupMemberList
	JoinGroupRequest                = group.JoinGroupRequest
	ModifyRemarkRequest             = group.ModifyRemarkRequest
	QuitGroupRequest                = group.QuitGroupRequest
	SetGroupMemberPermissionRequest = group.SetGroupMemberPermissionRequest

	Group interface {
		CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error)
		JoinGroup(ctx context.Context, in *JoinGroupRequest, opts ...grpc.CallOption) (*Empty, error)
		GetJoinList(ctx context.Context, in *GetJoinListRequest, opts ...grpc.CallOption) (*GetJoinListResponse, error)
		GetGroupMemberList(ctx context.Context, in *GetGroupMemberListRequest, opts ...grpc.CallOption) (*GetGroupMemberListResponse, error)
		QuitGroup(ctx context.Context, in *QuitGroupRequest, opts ...grpc.CallOption) (*Empty, error)
		SetGroupMemberPermission(ctx context.Context, in *SetGroupMemberPermissionRequest, opts ...grpc.CallOption) (*Empty, error)
		ModifyRemark(ctx context.Context, in *ModifyRemarkRequest, opts ...grpc.CallOption) (*Empty, error)
		DissolveGroup(ctx context.Context, in *DissolveGroupRequest, opts ...grpc.CallOption) (*Empty, error)
		GetGroupInfo(ctx context.Context, in *GetGroupInfoRequest, opts ...grpc.CallOption) (*GetGroupInfoResponse, error)
	}

	defaultGroup struct {
		cli zrpc.Client
	}
)

func NewGroup(cli zrpc.Client) Group {
	return &defaultGroup{
		cli: cli,
	}
}

func (m *defaultGroup) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.CreateGroup(ctx, in, opts...)
}

func (m *defaultGroup) JoinGroup(ctx context.Context, in *JoinGroupRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.JoinGroup(ctx, in, opts...)
}

func (m *defaultGroup) GetJoinList(ctx context.Context, in *GetJoinListRequest, opts ...grpc.CallOption) (*GetJoinListResponse, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.GetJoinList(ctx, in, opts...)
}

func (m *defaultGroup) GetGroupMemberList(ctx context.Context, in *GetGroupMemberListRequest, opts ...grpc.CallOption) (*GetGroupMemberListResponse, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.GetGroupMemberList(ctx, in, opts...)
}

func (m *defaultGroup) QuitGroup(ctx context.Context, in *QuitGroupRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.QuitGroup(ctx, in, opts...)
}

func (m *defaultGroup) SetGroupMemberPermission(ctx context.Context, in *SetGroupMemberPermissionRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.SetGroupMemberPermission(ctx, in, opts...)
}

func (m *defaultGroup) ModifyRemark(ctx context.Context, in *ModifyRemarkRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.ModifyRemark(ctx, in, opts...)
}

func (m *defaultGroup) DissolveGroup(ctx context.Context, in *DissolveGroupRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.DissolveGroup(ctx, in, opts...)
}

func (m *defaultGroup) GetGroupInfo(ctx context.Context, in *GetGroupInfoRequest, opts ...grpc.CallOption) (*GetGroupInfoResponse, error) {
	client := group.NewGroupClient(m.cli.Conn())
	return client.GetGroupInfo(ctx, in, opts...)
}
