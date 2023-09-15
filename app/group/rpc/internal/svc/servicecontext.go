package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"go-chat/app/group/model"
	"go-chat/app/group/rpc/internal/config"
	"go-chat/app/user/rpc/userclient"
)

type ServiceContext struct {
	Config         config.Config
	GroupModel     model.GroupModel
	UserGroupModel model.UserGroupModel
	UserRpc        userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:         c,
		GroupModel:     model.NewGroupModel(conn, c.CacheRedis),
		UserGroupModel: model.NewUserGroupModel(conn, c.CacheRedis),
		UserRpc:        userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
