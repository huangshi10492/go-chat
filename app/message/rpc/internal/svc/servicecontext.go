package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"go-chat/app/friend/rpc/friendclient"
	"go-chat/app/group/rpc/groupclient"
	"go-chat/app/message/model"
	"go-chat/app/message/rpc/internal/config"
)

type ServiceContext struct {
	Config         config.Config
	Redis          *redis.Redis
	FriendRpc      friendclient.Friend
	GroupRpc       groupclient.Group
	MailboxModel   model.MailboxModel
	MessageModel   model.MessageModel
	KqPusherClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		Redis:          redis.MustNewRedis(c.RedisClient),
		FriendRpc:      friendclient.NewFriend(zrpc.MustNewClient(c.FriendRpc)),
		GroupRpc:       groupclient.NewGroup(zrpc.MustNewClient(c.GroupRpc)),
		MailboxModel:   model.NewMailboxModel(c.MailBoxModel.Url, c.MailBoxModel.DB, c.MailBoxModel.Collection),
		MessageModel:   model.NewMessageModel(c.MessageModel.Url, c.MessageModel.DB, c.MessageModel.Collection),
		KqPusherClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
	}
}
