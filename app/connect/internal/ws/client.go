package ws

import (
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"go-chat/pkg/protobuf/ws_pb"
	"google.golang.org/protobuf/proto"
	"time"
)

// Client 客户端连接信息
type Client struct {
	manager       *ClientManager
	userId        int64
	connectId     string
	conn          *websocket.Conn // 连接
	heartbeatTime int64           // 前一次心跳时间
}

// 读取信息，即收到消息,维持心跳消息
func (c *Client) read() {
	defer func() {
		c.manager.unregisterChan <- c
		_ = c.conn.Close()
	}()
	for {
		_, body, err := c.conn.ReadMessage()
		if err != nil {
			logx.Error(err)
			break
		}
		res := &ws_pb.WsRequest{}
		err = proto.Unmarshal(body, res)
		if err != nil {
			logx.Error(err)
			break
		}
		if res.Type == 1 { // 维持心跳消息
			// 刷新连接时间
			c.heartbeatTime = time.Now().Unix()
			res, err := proto.Marshal(&ws_pb.WsResponse{Type: 233})
			if err != nil {
				logx.Error(err)
				continue
			}
			_ = c.Send(res)
		} else if res.Type == 2 { // 下线
			break
		}
	}
}

// Send 发送消息
func (c *Client) Send(body []byte) error {
	err := c.conn.WriteMessage(websocket.BinaryMessage, body)
	if err != nil {
		logx.Error(err)
	}
	return nil
}
