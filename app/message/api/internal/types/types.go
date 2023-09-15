// Code generated by goctl. DO NOT EDIT.
package types

type SendMessageReq struct {
	SenderType  int32  `json:"senderType"`
	GroupId     int64  `json:"groupId"`
	ReceiverId  int64  `json:"receiverId"`
	ContentType int32  `json:"contentType"`
	Content     string `json:"content"`
	ClientId    int64  `json:"clientId"`
}

type ReadReq struct {
	MessageId string `json:"messageId"`
	ClientId  int64  `json:"clientId"`
}

type WithdrawReq struct {
	SenderType int32  `json:"senderType"`
	GroupId    int64  `json:"groupId"`
	ReceiverId int64  `json:"receiverId"`
	MessageId  string `json:"messageId"`
	ClientId   int64  `json:"clientId"`
}