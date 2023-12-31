// Code generated by goctl. DO NOT EDIT.
package types

type SendFriendReq struct {
	FriendId int64  `json:"friendId"`
	Reason   string `json:"reason"`
	Remark   string `json:"remark"`
}

type AgreeFriendReq struct {
	RequestId int64  `json:"requestId"`
	Remark    string `json:"remark"`
	Agree     bool   `json:"agree"`
}

type ModifyFriendRemarkReq struct {
	FriendId int64  `json:"friendId"`
	Remark   string `json:"remark"`
}

type DeleteFriendReq struct {
	FriendId int64 `json:"friendId"`
}

type GetFriendListResp struct {
	FriendList []FriendInfo `json:"friendList"`
}

type FriendInfo struct {
	FriendId int64  `json:"friendId"`
	Name     string `json:"name"`
	Remark   string `json:"remark"`
}

type GetFriendRequestListResp struct {
	FriendRequestList []FriendRequestInfo `json:"friendRequestList"`
}

type FriendRequestInfo struct {
	RequestId int64  `json:"requestId"`
	FriendId  int64  `json:"friendId"`
	Name      string `json:"name"`
	Reason    string `json:"reason"`
	Status    int64  `json:"status"`
}
