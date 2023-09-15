// Code generated by goctl. DO NOT EDIT.
package types

type CreateGroupRequest struct {
	Name string `json:"name"`
}

type CreateGroupResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type JoinGroupRequest struct {
	GroupId int64 `json:"groupId"`
}

type GetJoinListResponse struct {
	GroupList []GroupList `json:"groupList"`
}

type GroupList struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Remark string `json:"remark"`
}

type GetGroupMemberListRequest struct {
	GroupId int64 `json:"groupId"`
}

type GetGroupMemberListResponse struct {
	MemberList []GroupMemberList `json:"groupMemberList"`
}

type GroupMemberList struct {
	UserId     int64  `json:"userId"`
	Name       string `json:"name"`
	Permission int64  `json:"permission"`
}

type QuitGroupRequest struct {
	MemberId int64 `json:"memberId"`
	GroupId  int64 `json:"groupId"`
	OwnerId  int64 `json:"ownerId"`
}

type SetGroupMemberPermissionRequest struct {
	GroupId    int64 `json:"groupId"`
	MemberId   int64 `json:"memberId"`
	Permission int64 `json:"permission"`
}

type ModifyRemarkRequest struct {
	GroupId int64  `json:"groupId"`
	Remark  string `json:"remark"`
}

type DissolveGroupRequest struct {
	GroupId int64 `json:"groupId"`
}

type GetGroupInfoRequest struct {
	GroupId int64 `json:"groupId"`
}

type GetGroupInfoResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}