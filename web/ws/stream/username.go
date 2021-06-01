package stream

import (
	"context"

	"dev.sum7.eu/genofire/golang-lib/web/ws"
	"github.com/bdlm/log"
)

// GetUsername of subscriber
func (we *endpoint) getUsername(s *ws.Subscriber) string {
	we.usernameMU.RLock()
	defer we.usernameMU.RUnlock()

	username, ok := we.subscriberUsername[s]
	if !ok {
		return "unknown"
	}
	return username
}

// setUsername to subcriber
func (we *endpoint) setUsername(s *ws.Subscriber, username string) bool {
	we.usernameMU.Lock()
	defer we.usernameMU.Unlock()

	if _, ok := we.usernames[username]; ok {
		return false
	}
	we.usernames[username] = s
	we.subscriberUsername[s] = username
	return true
}

// deleteUsername of sucbriber
func (we *endpoint) deleteUsername(s *ws.Subscriber) {
	we.usernameMU.Lock()
	defer we.usernameMU.Unlock()

	username := we.subscriberUsername[s]
	delete(we.usernames, username)
	delete(we.subscriberUsername, s)
}

// usernameHandler for MessageTypeUsername
func (we *endpoint) usernameHandler(_ context.Context, msg *ws.Message) {
	if u, ok := msg.Body[ws.BodySet]; ok {
		username := u.(string)
		if we.setUsername(msg.Subscriber, username) {
			msg.Body[ws.BodySet] = username
		} else {
			msg.Body[ws.BodyError] = "already in use"
		}
	}
	if _, ok := msg.Body[ws.BodyGet]; ok {
		msg.Body[ws.BodyGet] = we.getUsername(msg.Subscriber)
	}
	log.Warnf("usernameHandler: %v", msg.Body)
	msg.Reply(msg)
}
