package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	SenderId    int64              `bson:"senderId"`
	GroupId     int64              `bson:"groupId,omitempty"`
	MessageType int32              `bson:"messageType"` //1:文本 2:图片 3:文件
	MessageData string             `bson:"messageData"`
	UpdateAt    time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt    time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
