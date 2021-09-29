package stream

import (
	"context"

	"dev.sum7.eu/genofire/golang-lib/web/ws"
	"go.uber.org/zap"
)

// chatHandler for MessageTypeChat
func (e *endpoint) chatHandler(_ context.Context, msg *ws.Message) {
	username := e.getUsername(msg.Subscriber)
	msg.Body["username"] = username
	e.chatMessages = append(e.chatMessages, msg)
	e.log.Info("chatHandler", zap.String("username", username))
	e.Broadcast(msg)
}
