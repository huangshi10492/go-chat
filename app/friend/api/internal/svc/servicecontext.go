package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-chat/app/friend/api/internal/config"
	"go-chat/app/friend/rpc/friendclient"
)

type ServiceContext struct {
	Config    config.Config
	FriendRpc friendclient.Friend
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		FriendRpc: friendclient.NewFriend(zrpc.MustNewClient(c.FriendRpc)),
	}
}
