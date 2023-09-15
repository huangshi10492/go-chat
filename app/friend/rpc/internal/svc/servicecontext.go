package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"go-chat/app/friend/model"
	"go-chat/app/friend/rpc/internal/config"
	"go-chat/app/user/rpc/userclient"
)

type ServiceContext struct {
	Config             config.Config
	FriendModel        model.FriendModel
	FriendRequestModel model.FriendRequestModel
	UserRpc            userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:             c,
		FriendModel:        model.NewFriendModel(conn, c.CacheRedis),
		FriendRequestModel: model.NewFriendRequestModel(conn, c.CacheRedis),
		UserRpc:            userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
