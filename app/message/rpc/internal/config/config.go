package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RedisClient  redis.RedisConf
	GroupRpc     zrpc.RpcClientConf
	FriendRpc    zrpc.RpcClientConf
	MailBoxModel struct {
		Url        string
		DB         string
		Collection string
	}
	MessageModel struct {
		Url        string
		DB         string
		Collection string
	}
	KqPusherConf struct {
		Brokers []string
		Topic   string
	}
}
