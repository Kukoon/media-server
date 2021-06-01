package stream

import (
	"dev.sum7.eu/genofire/golang-lib/web/ws"
)

// SendStatus
func (we *endpoint) SendStatus(origin *ws.Message) {
	we.usernameMU.RLock()
	msg := ws.Message{
		Type: MessageTypeStatus,
		Body: map[string]interface{}{
			"viewers":   len(we.Subscribers),
			"chatusers": len(we.usernames),
		},
	}
	we.usernameMU.RUnlock()
	if origin != nil {
		origin.Reply(&msg)
	} else {
		we.Broadcast(&msg)
	}
}
