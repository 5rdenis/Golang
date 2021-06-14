package main

import (

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", getAllUsers)
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PATCH("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
	e.PATCH("/users/:idFrom/:idTo", transfer)

	e.Logger.Fatal(e.Start(":1323"))
}