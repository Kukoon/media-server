package stream

import (
	"time"

	"dev.sum7.eu/genofire/golang-lib/web/ws"
	"github.com/Kukoon/media-server/models"
	"go.uber.org/zap"
)

// SendStatus
func (we *endpoint) SendStatus(origin *ws.Message) {
	now := time.Now()
	m := &models.Stream{}
	if err := we.web.DB.
		Where("channel_id=?", we.channelID).
		Where("listen_at < ?", now).
		Where("start_at < ?", now).
		Order("start_at DESC").
		First(m).Error; err != nil {
		we.log.Warn("websocket: send status could not fetch current stream",
			zap.String("channel_id", we.channelID.String()),
			zap.Error(err),
		)
	}
	we.usernameMU.RLock()
	msg := ws.Message{
		Type: MessageTypeStatus,
		Body: map[string]interface{}{
			"viewers":   len(we.Subscribers),
			"chatusers": len(we.usernames),
			"running":   m.Running,
			"stream":    m.ID,
		},
	}
	we.usernameMU.RUnlock()
	if origin != nil {
		origin.Reply(&msg)
	} else {
		we.Broadcast(&msg)
	}
}
