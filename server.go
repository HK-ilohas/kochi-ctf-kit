package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"problemServer/api"
	"problemServer/handler"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/login", handler.Login())

	users := e.Group("/users")
	users.Use(middleware.JWT([]byte("secret")))
	users.GET("/welcome", handler.Restricted()) // users/welcome

	users.GET("/team", api.GetTeamData())
	e.Logger.Fatal(e.Start(":1323"))
}
