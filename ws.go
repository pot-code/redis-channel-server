package main

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
)

type WebsocketHandler struct {
	cm *RedisChannelManger
}

func NewWebsocketHandler(cm *RedisChannelManger) *WebsocketHandler {
	return &WebsocketHandler{
		cm: cm,
	}
}

func (h *WebsocketHandler) ServeHTTP(e echo.Context) error {
	ws, err := upgrader.Upgrade(e.Response(), e.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	cid := e.QueryParam("channel")
	for msg := range h.cm.Subscribe(e.Request().Context(), cid) {
		err := ws.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return err
		}
	}
	return nil
}
