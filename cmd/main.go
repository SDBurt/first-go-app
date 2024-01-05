package main

import (
	"context"
	"first-go/handler"

	"github.com/labstack/echo/v4"
)

type DB struct{}

func main() {
	app := echo.New()

	userHandler := handler.UserHandler{}
	app.Use(withUser)
	app.GET("/user", userHandler.HandleUserShow)
	app.Static("/static", "static")
	app.Start(":8080")
}

func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "user", "a@gg.com")
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
