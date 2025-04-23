package controller

import (
	"agnos-assignment/app/service"

	"github.com/gin-gonic/gin"
)

type WebsocketControllerInterface interface {
	ConnectChat(c *gin.Context)
}

type WebsocketController struct {
	Svc *service.WebSocketService
}

func WebsocketControllerInit(s *service.WebSocketService) *WebsocketController {
	return &WebsocketController{Svc: s}
}

func (w *WebsocketController) ConnectChat(c *gin.Context) {
	room := c.Param("room")
	w.Svc.HandleConnection(c.Writer, c.Request, room)
}
