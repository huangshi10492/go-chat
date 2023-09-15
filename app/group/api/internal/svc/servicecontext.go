package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-chat/app/group/api/internal/config"
	"go-chat/app/group/rpc/groupclient"
)

type ServiceContext struct {
	Config   config.Config
	GroupRpc groupclient.Group
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		GroupRpc: groupclient.NewGroup(zrpc.MustNewClient(c.GroupRpc)),
	}
}
