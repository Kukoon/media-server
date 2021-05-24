package stream

import (
	"context"
	"sync"
	"time"

	"github.com/bdlm/log"

	"github.com/Kukoon/media-server/web/ws"
)

const (
	MessageTypeChat     = "chat"
	MessageTypePing     = "ping"
	MessageTypeStatus   = "status"
	MessageTypeUsername = "username"
)

type endpoint struct {
	*ws.WebsocketEndpoint
	usernameMU         sync.RWMutex
	usernames          map[string]*ws.Subscriber
	subscriberUsername map[*ws.Subscriber]string
	chatMessages       []*ws.Message
}

func (this *endpoint) onOpen(s *ws.Subscriber, out chan<- *ws.Message) {
	for _, msg := range this.chatMessages {
		out <- msg
	}
}
func (this *endpoint) onClose(s *ws.Subscriber, out chan<- *ws.Message) {
	this.deleteUsername(s)
}

func (this *endpoint) SendStatus(origin *ws.Message) {
	this.usernameMU.RLock()
	msg := ws.Message{
		Type: MessageTypeStatus,
		Body: map[string]interface{}{
			"viewers":   len(this.Subscribers),
			"chatusers": len(this.usernames),
		},
	}
	this.usernameMU.RUnlock()
	if origin != nil {
		origin.Reply <- &msg
	} else {
		this.Broadcast(&msg)
	}
}
func (this *endpoint) chatHandler(_ context.Context, msg *ws.Message) {
	msg.Body["username"] = this.getUsername(msg.Subscriber)
	this.chatMessages = append(this.chatMessages, msg)
	log.Warnf("chatHandler: %v", msg.Body)
	this.Broadcast(msg)
}

func (this *endpoint) getUsername(s *ws.Subscriber) string {
	this.usernameMU.RLock()
	defer this.usernameMU.RUnlock()

	username, ok := this.subscriberUsername[s]
	if !ok {
		return "unknown"
	}
	return username
}
func (this *endpoint) setUsername(s *ws.Subscriber, username string) bool {
	this.usernameMU.Lock()
	defer this.usernameMU.Unlock()

	if _, ok := this.usernames[username]; ok {
		return false
	}
	this.usernames[username] = s
	this.subscriberUsername[s] = username
	return true
}
func (this *endpoint) deleteUsername(s *ws.Subscriber) {
	this.usernameMU.Lock()
	defer this.usernameMU.Unlock()

	username := this.subscriberUsername[s]
	delete(this.usernames, username)
	delete(this.subscriberUsername, s)
}
func (this *endpoint) usernameHandler(_ context.Context, msg *ws.Message) {
	if u, ok := msg.Body[ws.BodySet]; ok {
		username := u.(string)
		if this.setUsername(msg.Subscriber, username) {
			msg.Body[ws.BodySet] = username
		} else {
			msg.Body[ws.BodyError] = "already in use"
		}
	}
	if _, ok := msg.Body[ws.BodyGet]; ok {
		msg.Body[ws.BodyGet] = this.getUsername(msg.Subscriber)
	}
	log.Warnf("usernameHandler: %v", msg.Body)
	msg.Reply <- msg
}

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
		msg.Reply <- msg
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
