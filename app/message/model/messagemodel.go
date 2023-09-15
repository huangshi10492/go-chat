package model

import (
	"github.com/zeromicro/go-zero/core/stores/mon"
)

var _ MessageModel = (*customMessageModel)(nil)

type (
	// MessageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageModel.
	MessageModel interface {
		messageModel
	}

	customMessageModel struct {
		*defaultMessageModel
	}
)

// NewMessageModel returns a model for the mongo.
func NewMessageModel(url, db, collection string) MessageModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customMessageModel{
		defaultMessageModel: newDefaultMessageModel(conn),
	}
}
