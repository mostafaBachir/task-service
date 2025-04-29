package project

import (
	"github.com/mostafaBachir/task-service/security"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/contrib/websocket"
)

func WebSocketAuthMiddleware(c *fiber.Ctx) error {
	if err := security.CheckJWTWebSocket(c); err != nil {
		return err
	}
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}

func WebSocketConnect(c *websocket.Conn) {
	manager.AddClient(c)
	defer func() {
		manager.RemoveClient(c)
		c.Close()
	}()

	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			break
		}
	}
}
