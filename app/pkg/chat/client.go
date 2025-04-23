package chat

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
	Room *Room
}

func NewClient(conn *websocket.Conn, room *Room) *Client {
	return &Client{
		Conn: conn,
		Send: make(chan []byte, 256),
		Room: room,
	}
}
