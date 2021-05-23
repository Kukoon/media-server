package ws

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/bdlm/log"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type WebsocketEndpoint struct {
	// publishLimiter controls the rate limit applied to the publish endpoint.
	//
	// Defaults to one publish every 100ms with a burst of 8.
	publishLimiter *rate.Limiter

	subscribersMu sync.Mutex
	Subscribers   map[*Subscriber]struct{}

	// Message Handler
	handlers              map[string]MessageHandleFunc
	DefaultMessageHandler MessageHandleFunc
	OnOpen                SubscriberEventFunc
	OnClose               SubscriberEventFunc
}

// MessageHandleFunc for handling messages
type MessageHandleFunc func(ctx context.Context, msg *Message)

type SubscriberEventFunc func(s *Subscriber)

// Message on websocket
type Message struct {
	Type       string                 `json:"type"`
	Body       map[string]interface{} `json:"body"`
	Reply      chan<- *Message        `json:"-"`
	Subscriber *Subscriber            `json:"-"`
}

func NewEndpoint() *WebsocketEndpoint {
	return &WebsocketEndpoint{
		publishLimiter: rate.NewLimiter(rate.Every(time.Millisecond*100), 8),
		Subscribers:    make(map[*Subscriber]struct{}),
		handlers:       make(map[string]MessageHandleFunc),
	}
}
func (this *WebsocketEndpoint) Handler(ctx *gin.Context) {
	c, err := websocket.Accept(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, false)
		return
	}
	defer c.Close(websocket.StatusInternalError, "")

	err = this.addSubscriber(ctx, c)

	if websocket.CloseStatus(err) == websocket.StatusNormalClosure ||
		websocket.CloseStatus(err) == websocket.StatusGoingAway {
		return
	}
	log.Errorf("subscriber stopped: %s", err)
}

func (this *WebsocketEndpoint) AddMessageHandler(typ string, f MessageHandleFunc) {
	this.handlers[typ] = f
}

type Subscriber struct {
	out       chan *Message
	closeSlow func()
}

func (this *WebsocketEndpoint) readWorker(ctx context.Context, c *websocket.Conn, s *Subscriber) error {
	for {
		var msg Message
		err := wsjson.Read(ctx, c, &msg)
		if err != nil {
			return err
		}
		log.WithField("type", msg.Type).Debug("recieve")
		msg.Subscriber = s
		if handler, ok := this.handlers[msg.Type]; ok {
			handler(ctx, &msg)
		} else if this.DefaultMessageHandler != nil {
			this.DefaultMessageHandler(ctx, &msg)
		}
	}
}

func (this *WebsocketEndpoint) addSubscriber(ctxGin *gin.Context, c *websocket.Conn) error {
	s := &Subscriber{
		out: make(chan *Message, 10),
		closeSlow: func() {
			c.Close(websocket.StatusPolicyViolation, "connection too slow to keep up with messages")
		},
	}

	this.subscribersMu.Lock()
	this.Subscribers[s] = struct{}{}
	this.subscribersMu.Unlock()
	defer func() {
		this.subscribersMu.Lock()
		delete(this.Subscribers, s)
		this.subscribersMu.Unlock()
		if this.OnClose != nil {
			this.OnClose(s)
		}
		log.Debug("websocket closed")
	}()

	if this.OnOpen != nil {
		this.OnOpen(s)
	}

	ctx := ctxGin.Request.Context()

	go func() {
		err := this.readWorker(ctx, c, s)
		if websocket.CloseStatus(err) == websocket.StatusNormalClosure ||
			websocket.CloseStatus(err) == websocket.StatusGoingAway {
			return
		}
		log.Errorf("websocket reading error: %s", err)
	}()
	log.Debug("websocket started")

	for {
		select {
		case msg := <-s.out:
			err := writeTimeout(ctx, time.Second*5, c, msg)
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func writeTimeout(ctx context.Context, timeout time.Duration, c *websocket.Conn, msg *Message) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return wsjson.Write(ctx, c, msg)
}

func (this *WebsocketEndpoint) Broadcast(msg *Message) {
	this.subscribersMu.Lock()
	defer this.subscribersMu.Unlock()

	this.publishLimiter.Wait(context.Background())

	for s := range this.Subscribers {
		if s == msg.Subscriber {
			continue
		}
		select {
		case s.out <- msg:
		default:
			go s.closeSlow()
		}
	}
}
