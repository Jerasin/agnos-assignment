// go:build wireinject
// go:build wireinject
//go:build wireinject
// +build wireinject

package module

import (
	"agnos-assignment/app/controller"
	"agnos-assignment/app/pkg/chat"
	"agnos-assignment/app/service"

	"github.com/google/wire"
)

var hub = wire.NewSet(chat.NewHub)

var websocketCtrlSet = wire.NewSet(controller.WebsocketControllerInit,
	wire.Bind(new(controller.WebsocketControllerInterface), new(*controller.WebsocketController)),
)

var websocketSvcSet = wire.NewSet(service.WebsocketServiceInit,
	wire.Bind(new(service.WebsocketServiceInterface), new(*service.WebSocketService)),
)

type WebsocketModule struct {
	WebsocketCtrl controller.WebsocketControllerInterface
	WebsocketSvc  service.WebsocketServiceInterface
}

func NewWebsocketModule(
	websocketCtrl controller.WebsocketControllerInterface,
	websocketSvc service.WebsocketServiceInterface,
) *WebsocketModule {
	return &WebsocketModule{
		WebsocketCtrl: websocketCtrl,
		WebsocketSvc:  websocketSvc,
	}
}

func WebsocketModuleInit() *WebsocketModule {
	wire.Build(NewWebsocketModule, websocketCtrlSet, websocketSvcSet, hub)
	return nil
}
