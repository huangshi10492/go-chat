package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mailbox struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	SenderId    int64              `bson:"senderId"`
	ReceiverId  int64              `bson:"receiverId"`
	ContentType int32              `bson:"content_type"` //1:指向消息id 2:已读的消息id 3:撤销的消息id
	Content     string             `bson:"content"`
	Seq         int64              `bson:"seq"`
	UpdateAt    time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt    time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
