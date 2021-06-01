package stream

import (
	"context"
	"sync"
	"time"

	"dev.sum7.eu/genofire/golang-lib/web/ws"
	"github.com/bdlm/log"
)

const (
	// MessageTypeChat of chat messages
	MessageTypeChat = "chat"
	// MessageTypePing of ping messages
	MessageTypePing = "ping"
	// MessageTypeStatus of status messages
	MessageTypeStatus = "status"
	// MessageTypeUsername of username messages
	MessageTypeUsername = "username"
)

type endpoint struct {
	*ws.WebsocketEndpoint
	usernameMU         sync.RWMutex
	usernames          map[string]*ws.Subscriber
	subscriberUsername map[*ws.Subscriber]string
	chatMessages       []*ws.Message
}

// onOpen of subscriber connections - send preaves messages
func (we *endpoint) onOpen(s *ws.Subscriber, out chan<- *ws.Message) {
	for _, msg := range we.chatMessages {
		out <- msg
	}
}

// onClose of subscriber connections - delete username
func (we *endpoint) onClose(s *ws.Subscriber, out chan<- *ws.Message) {
	we.deleteUsername(s)
}

// NewEndpoint of Websocket for stream
func NewEndpoint() *endpoint {
	we := endpoint{
		WebsocketEndpoint:  ws.NewEndpoint(),
		usernames:          make(map[string]*ws.Subscriber),
		subscriberUsername: make(map[*ws.Subscriber]string),
	}

	we.OnOpen = we.onOpen
	we.OnClose = we.onClose
	we.DefaultMessageHandler = func(ctx context.Context, msg *ws.Message) {
		log.WithField("type", msg.Type).Warn("unsupported websocket message")
	}

	we.AddMessageHandler(MessageTypePing, func(ctx context.Context, msg *ws.Message) {
		msg.Reply(msg)
	})
	we.AddMessageHandler(MessageTypeStatus, func(ctx context.Context, msg *ws.Message) {
		we.SendStatus(msg)
	})
	we.AddMessageHandler(MessageTypeUsername, we.usernameHandler)
	we.AddMessageHandler(MessageTypeChat, we.chatHandler)

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				we.SendStatus(nil)
			}
		}
	}()
	return &we
}
