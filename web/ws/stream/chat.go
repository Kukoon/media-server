package stream

import (
	"context"

	"dev.sum7.eu/genofire/golang-lib/web/ws"
	"github.com/bdlm/log"
)

// chatHandler for MessageTypeChat
func (this *endpoint) chatHandler(_ context.Context, msg *ws.Message) {
	msg.Body["username"] = this.getUsername(msg.Subscriber)
	this.chatMessages = append(this.chatMessages, msg)
	log.Warnf("chatHandler: %v", msg.Body)
	this.Broadcast(msg)
}
