package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FriendRequestModel = (*customFriendRequestModel)(nil)

type (
	// FriendRequestModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFriendRequestModel.
	FriendRequestModel interface {
		friendRequestModel
		FindAllByUserId(ctx context.Context, userId int64) ([]*FriendRequest, error)
		GetFriendRequestList(ctx context.Context, id int64) ([]*FriendRequest, error)
	}

	customFriendRequestModel struct {
		*defaultFriendRequestModel
	}
)

// NewFriendRequestModel returns a model for the database table.
func NewFriendRequestModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) FriendRequestModel {
	return &customFriendRequestModel{
		defaultFriendRequestModel: newFriendRequestModel(conn, c, opts...),
	}
}

// FindAllByUserId 获取收到的申请列表
func (m *customFriendRequestModel) FindAllByUserId(ctx context.Context, userId int64) ([]*FriendRequest, error) {
	var resp []*FriendRequest

	query := fmt.Sprintf("select %s from %s where `friend_id` = ?", friendRequestRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userId)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customFriendRequestModel) GetFriendRequestList(ctx context.Context, id int64) ([]*FriendRequest, error) {
	var resp []*FriendRequest
	query := fmt.Sprintf("select * from %s where `friend_id` = ?", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, id)
	return resp, err
}
