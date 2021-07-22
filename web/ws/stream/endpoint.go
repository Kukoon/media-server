package stream

import (
	"context"
	"sync"
	"time"

	"dev.sum7.eu/genofire/golang-lib/web"
	"dev.sum7.eu/genofire/golang-lib/web/ws"
	"dev.sum7.eu/genofire/golang-lib/worker"
	"github.com/bdlm/log"
	"github.com/google/uuid"
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
	web                *web.Service
	Worker             *worker.Worker
	channelID          uuid.UUID
	Running            bool
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
func NewEndpoint(web *web.Service, channelID uuid.UUID) *endpoint {
	we := endpoint{
		WebsocketEndpoint:  ws.NewEndpoint(),
		web:                web,
		channelID:          channelID,
		usernames:          make(map[string]*ws.Subscriber),
		subscriberUsername: make(map[*ws.Subscriber]string),
	}

	we.OnOpen = we.onOpen
	we.OnClose = we.onClose
	we.Worker = worker.NewWorker(5*time.Second, func() {
		we.SendStatus(nil)
	})
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

	we.Worker.Start()

	return &we
}
