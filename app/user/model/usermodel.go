package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel

		UsersInfo(ctx context.Context, userIds []int64) ([]*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c, opts...),
	}
}

func (m *customUserModel) UsersInfo(ctx context.Context, userIds []int64) ([]*User, error) {
	var resp []*User

	query := fmt.Sprintf("select %s from %s where `id` in (?)", userRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userIds)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
