package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
)

func main() {
	rd := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	cm := NewRedisChannelManger(rd)
	a := NewApi(cm)
	ws := NewWebsocketHandler(cm)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/ws", ws.ServeHTTP)
	e.POST("/pub", a.publish)
	e.Logger.Fatal(e.Start(":1323"))
}
