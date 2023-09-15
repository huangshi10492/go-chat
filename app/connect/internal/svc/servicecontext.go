package svc

import (
	"go-chat/app/connect/internal/config"
	"go-chat/app/connect/internal/ws"
)

type ServiceContext struct {
	Config    config.Config
	WebSocket *ws.ClientManager
	TokenMap  map[string]int64
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		WebSocket: ws.NewClientManager(),
		TokenMap:  make(map[string]int64),
	}
}
