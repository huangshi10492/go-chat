package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FriendModel = (*customFriendModel)(nil)

type (
	// FriendModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFriendModel.
	FriendModel interface {
		friendModel
		DeleteFriend(ctx context.Context, id1 int64, id2 int64) error
		GetFriendList(ctx context.Context, id int64) ([]*Friend, error)
	}

	customFriendModel struct {
		*defaultFriendModel
	}
)

// NewFriendModel returns a model for the database table.
func NewFriendModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) FriendModel {
	return &customFriendModel{
		defaultFriendModel: newFriendModel(conn, c, opts...),
	}
}

func (m *customFriendModel) DeleteFriend(ctx context.Context, id1 int64, id2 int64) error {
	err := m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		query := fmt.Sprintf("delete from %s where `user_id` = ? and `friend_id` = ?", m.table)
		_, err := session.ExecCtx(ctx, query, id1, id2)
		if err != nil {
			return err
		}
		_, err = session.ExecCtx(ctx, query, id2, id1)
		return err
	})
	return err
}
func (m *customFriendModel) GetFriendList(ctx context.Context, id int64) ([]*Friend, error) {
	var friends []*Friend
	query := fmt.Sprintf("select * from %s where `user_id` = ?", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &friends, query, id)
	return friends, err
}
