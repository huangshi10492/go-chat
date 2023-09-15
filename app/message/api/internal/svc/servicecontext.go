package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-chat/app/message/api/internal/config"
	"go-chat/app/message/rpc/messageclient"
)

type ServiceContext struct {
	Config     config.Config
	MessageRpc messageclient.Message
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		MessageRpc: messageclient.NewMessage(zrpc.MustNewClient(c.MessageRpc)),
	}
}
