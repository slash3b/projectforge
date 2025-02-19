// Content managed by Project Forge, see [projectforge.md] for details.
package websocket

import (
	"fmt"
	"sync/atomic"

	"github.com/fasthttp/websocket"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"

	"projectforge.dev/projectforge/app/util"
)

var count int64

func (s *Service) write(connID uuid.UUID, message string, logger util.Logger) error {
	if connID == systemID {
		logger.Warnf("admin message sent: %s", message)
		return nil
	}

	s.connectionsMu.Lock()
	c, ok := s.connections[connID]
	s.connectionsMu.Unlock()

	atomic.AddInt64(&count, 1)

	if !ok {
		return errors.New("cannot load connection [" + connID.String() + "]")
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	err := c.socket.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		return errors.Wrap(err, "unable to write to websocket")
	}
	return nil
}

func (s *Service) WriteMessage(connID uuid.UUID, message *Message, logger util.Logger) error {
	s.WriteTap(message, logger)
	return s.write(connID, util.ToJSON(message), logger)
}

func (s *Service) WriteChannel(message *Message, logger util.Logger, except ...uuid.UUID) error {
	if message == nil {
		return errors.New("no message provided")
	}
	if message.Channel == "" {
		return errors.New("no channel provided")
	}
	conns, ok := s.channels[message.Channel]
	if !ok {
		return nil
	}
	json := util.ToJSON(message)
	logger.Debugf("sending message [%v::%v] to [%v] connections", message.Channel, message.Cmd, len(conns.MemberIDs))
	for _, conn := range conns.MemberIDs {
		if !slices.Contains(except, conn) {
			connID := conn

			go func() {
				defer func() { _ = recover() }()
				_ = s.write(connID, json, logger)
			}()
		}
	}
	s.WriteTap(message, logger)
	return nil
}

func (s *Service) ReadLoop(connID uuid.UUID, logger util.Logger) error {
	c, ok := s.connections[connID]
	if !ok {
		return errors.New("cannot load connection [" + connID.String() + "]")
	}
	d := func() error {
		if s.onClose != nil {
			return s.onClose(s, c, logger)
		}
		return nil
	}
	m := func(m *Message) error {
		return OnMessage(s, connID, m, logger)
	}
	return ReadSocketLoop(connID, c.socket, m, d, logger)
}

func ReadSocketLoop(connID uuid.UUID, sock *websocket.Conn, onMessage func(m *Message) error, onDisconnect func() error, logger util.Logger) error {
	defer func() {
		_ = sock.Close()
		if onDisconnect != nil {
			err := onDisconnect()
			if err != nil {
				logger.Warn(fmt.Sprintf("error running onDisconnect for [%v]: %+v", connID, err))
			}
		}
		if onDisconnect == nil {
			err := onDisconnect()
			if err != nil {
				logger.Warn(fmt.Sprintf("error running disconnect for [%v]: %+v", connID, err))
			}
		}
		logger.Debug(fmt.Sprintf("closed websocket [%v]", connID.String()))
	}()

	for {
		_, message, err := sock.ReadMessage()
		if err != nil {
			return err
		}
		m := &Message{}
		err = util.FromJSON(message, m)
		if err != nil {
			return errors.Wrap(err, "error decoding websocket message")
		}
		err = onMessage(m)
		if err != nil {
			return errors.Wrap(err, "error handling websocket message")
		}
	}
}
