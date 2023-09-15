package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserGroupModel = (*customUserGroupModel)(nil)

type (
	// UserGroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserGroupModel.
	UserGroupModel interface {
		userGroupModel

		DissolveGroup(ctx context.Context, ids []int64, groupId int64) error
		FindUserIdByGroupId(ctx context.Context, groupId int64) ([]*UserGroup, error)
		GetJoinList(ctx context.Context, userId int64) ([]*UserGroup, error)
		SwitchOwner(ctx context.Context, fromId int64, toId int64) error
	}

	customUserGroupModel struct {
		*defaultUserGroupModel
	}
)

// NewUserGroupModel returns a model for the database table.
func NewUserGroupModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserGroupModel {
	return &customUserGroupModel{
		defaultUserGroupModel: newUserGroupModel(conn, c, opts...),
	}
}

func (m *customUserGroupModel) DissolveGroup(ctx context.Context, userIds []int64, groupId int64) error {
	query := fmt.Sprintf("delete from %s where `group_id` = ? and where `user_id` = ?", m.table)
	err := m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		for _, id := range userIds {
			_, err := session.ExecCtx(ctx, query, groupId, id)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
func (m *customUserGroupModel) FindUserIdByGroupId(ctx context.Context, groupId int64) ([]*UserGroup, error) {
	var resp []*UserGroup
	query := fmt.Sprintf("select %s from %s where `group_id` = ?", userGroupRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, groupId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *customUserGroupModel) GetJoinList(ctx context.Context, userId int64) ([]*UserGroup, error) {
	var resp []*UserGroup
	query := fmt.Sprintf("select %s from %s where `user_id` = ?", userGroupRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (m *customUserGroupModel) SwitchOwner(ctx context.Context, fromId int64, toId int64) error {
	query1 := fmt.Sprintf("update %s set `permission` where `id` = ?", m.table)
	query2 := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		_, err := session.ExecCtx(ctx, query1, 3, toId)
		if err != nil {
			return err
		}
		_, err = session.ExecCtx(ctx, query2, fromId)
		return err
	})
}
