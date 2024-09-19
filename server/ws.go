package server

import (
	"fmt"
	"github.com/lxzan/gws"
)

const SessionUserID = "userID"

func (s *Server) OnOpen(socket *gws.Conn) {}
func (s *Server) OnPing(socket *gws.Conn, payload []byte) {
	_ = socket.WritePong(nil)
}
func (s *Server) OnPong(socket *gws.Conn, payload []byte) {}

func (s *Server) OnClose(socket *gws.Conn, err error) {
	userID, _ := socket.Session().Load(SessionUserID)
	s.RemoveUser(userID.(string))
}

func (s *Server) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()

	msg, err := parseMessage(message.Bytes())
	if err != nil {
		fmt.Printf("parse message error: %v\n", err)
		return
	}
	switch msg.Type {
	case MessageTypeIdentify:
		s.OnIdentify(socket, msg)
	case MessageTypeText, MessageTypeFile:
		s.Broadcast(true, msg)
	}
}

func (s *Server) OnIdentify(socket *gws.Conn, msg *Message) {
	socket.Session().Store(SessionUserID, msg.UserID)
	s.AddUser(msg.UserID, socket)
	s.Broadcast(true, msg)
}
