package project

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type Client struct {
	Connection *websocket.Conn
}

type Manager struct {
	Clients map[*Client]bool
	Lock    sync.Mutex
}

var manager = Manager{
	Clients: make(map[*Client]bool),
}

func (m *Manager) AddClient(c *websocket.Conn) {
	client := &Client{Connection: c}
	m.Lock.Lock()
	defer m.Lock.Unlock()
	m.Clients[client] = true
}

func (m *Manager) RemoveClient(c *websocket.Conn) {
	m.Lock.Lock()
	defer m.Lock.Unlock()
	for client := range m.Clients {
		if client.Connection == c {
			delete(m.Clients, client)
			return
		}
	}
}

func (m *Manager) Broadcast(message []byte) {
	m.Lock.Lock()
	defer m.Lock.Unlock()
	for client := range m.Clients {
		client.Connection.WriteMessage(websocket.TextMessage, message)
	}
}
