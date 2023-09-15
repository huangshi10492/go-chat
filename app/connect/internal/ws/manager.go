package ws

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"go-chat/pkg/protobuf/ws_pb"
	"google.golang.org/protobuf/proto"
	"net/http"
	"sync"
	"time"
)

// 消息类型
const (
	HeartbeatCheckTime   = 29   // 心跳检测几秒检测一次
	HeartbeatTime        = 120  // 心跳距离上一次的最大时间
	ChanBufferRegister   = 1024 // 注册chan缓冲
	ChanBufferUnregister = 1024 // 注销chan大小
)

// ClientManager 客户端管理
type ClientManager struct {
	mu             *sync.RWMutex
	Clients        map[int64][]*Client // 保存连接
	registerChan   chan *Client        // 注册chan
	unregisterChan chan *Client        // 注销chan
}

func NewClientManager() *ClientManager {
	cm := &ClientManager{
		Clients:        make(map[int64][]*Client),
		mu:             new(sync.RWMutex),
		registerChan:   make(chan *Client, ChanBufferRegister),
		unregisterChan: make(chan *Client, ChanBufferUnregister),
	}
	// 注册注销
	go func() {
		defer func() {
			if r := recover(); r != nil {
				logx.Info(r)
			}
		}()
		cm.register()
	}()
	// 检查心跳
	go func() {
		defer func() {
			if r := recover(); r != nil {
				logx.Info(r)
			}
		}()
		cm.heartbeatCheck()
	}()
	return cm
}

// 注册注销
func (m *ClientManager) register() {
	for {
		select {
		case client := <-m.registerChan: // 新注册，新连接
			m.addClient(client)
			res, err := proto.Marshal(&ws_pb.WsResponse{Type: 233})
			if err != nil {
				logx.Error(err)
				continue
			}
			_ = client.conn.WriteMessage(websocket.BinaryMessage, res)

		case client := <-m.unregisterChan: // 注销，或者没有心跳
			// 关闭连接
			_ = client.conn.Close()
			m.removeClient(client)
		}
	}
}

func (m *ClientManager) addClient(client *Client) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Clients[client.userId] = append(m.Clients[client.userId], client)
}

func (m *ClientManager) removeClient(client *Client) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i, c := range m.Clients[client.userId] {
		if c.connectId == client.connectId {
			m.Clients[client.userId] = append(m.Clients[client.userId][:i], m.Clients[client.userId][i+1:]...)
			return
		}
	}
}

// 维持心跳
func (m *ClientManager) heartbeatCheck() {
	for {
		t := time.Now().Unix()
		m.mu.RLock()
		for _, clients := range m.Clients {
			for _, c := range clients {
				if time.Now().Unix()-c.heartbeatTime > HeartbeatTime {
					_ = c.conn.Close()
					m.removeClient(c)
				}
			}
		}
		m.mu.RUnlock()
		t = time.Now().Unix() - t
		if HeartbeatCheckTime > t {
			time.Sleep(time.Second * time.Duration(HeartbeatCheckTime-t))
		}
	}
}

func (m *ClientManager) GetClients(userId int64) []*Client {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.Clients[userId]
}

// NewClient 新建连接
func (m *ClientManager) NewClient(w http.ResponseWriter, r *http.Request, userId int64) {
	conn, err := (&websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(w, r, nil)
	if err != nil {
		logx.Error(err)
		return
	}
	// 创建一个实例连接
	client := &Client{
		manager:       m,
		userId:        userId,
		connectId:     uuid.NewString(),
		heartbeatTime: time.Now().Unix(),
		conn:          conn,
	}

	// 用户注册到用户连接管理
	m.registerChan <- client
	go client.read()
}
