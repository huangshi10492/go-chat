package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
)

var _ MailboxModel = (*customMailboxModel)(nil)

type (
	// MailboxModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMailboxModel.
	MailboxModel interface {
		mailboxModel

		InsertMany(ctx context.Context, data []any) error
	}

	customMailboxModel struct {
		*defaultMailboxModel
	}
)

// NewMailboxModel returns a model for the mongo.
func NewMailboxModel(url, db, collection string) MailboxModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customMailboxModel{
		defaultMailboxModel: newDefaultMailboxModel(conn),
	}
}

func (m *customMailboxModel) InsertMany(ctx context.Context, data []any) error {
	_, err := m.conn.InsertMany(ctx, data)
	return err
}
