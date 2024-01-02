package main

import (
	"first-go/handler"

	"github.com/labstack/echo/v4"
)

type DB struct{}

func main() {
	app := echo.New()

	// var db DB

	userHandler := handler.UserHandler{}
	// userHandler := hander.UserHandler{}
	app.GET("/user", userHandler.HandleUserShow)

	app.Start(":8080")
}
