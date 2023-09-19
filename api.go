package main

import (
	"github.com/labstack/echo/v4"
)

type Api struct {
	cm *RedisChannelManger
}

func NewApi(cm *RedisChannelManger) *Api {
	return &Api{
		cm: cm,
	}
}

func (a *Api) publish(e echo.Context) error {
	var m Message
	if err := e.Bind(&m); err != nil {
		return err
	}

	err := a.cm.Publish(e.Request().Context(), &m)
	if err != nil {
		return err
	}
	return nil
}
