package service

import (
	"agnos-assignment/app/pkg/chat"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebsocketServiceInterface interface {
	HandleConnection(w http.ResponseWriter, r *http.Request, roomName string)
}

type WebSocketService struct {
	Hub *chat.Hub
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func WebsocketServiceInit(hub *chat.Hub) *WebSocketService {
	return &WebSocketService{
		Hub: hub,
	}
}

func (s *WebSocketService) HandleConnection(w http.ResponseWriter, r *http.Request, roomName string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	room := s.Hub.GetRoom(roomName)
	client := chat.NewClient(conn, room)

	room.Register <- client

	go func() {
		defer func() {
			room.Unregister <- client
			client.Conn.Close()
		}()
		for {
			_, message, err := client.Conn.ReadMessage()
			if err != nil {
				break
			}
			room.Broadcast <- message
		}
	}()

	go func() {
		for msg := range client.Send {
			client.Conn.WriteMessage(websocket.TextMessage, msg)
		}
	}()
}
