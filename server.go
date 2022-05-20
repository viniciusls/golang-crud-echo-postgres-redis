package main

import (
	"crud-echo-postgres-redis/controllers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/users", func(c echo.Context) error {
		return controllers.GetAllUsers(c)
	})
	e.POST("/users", func(c echo.Context) error {
		return controllers.CreateUser(c)
	})
	e.GET("/users/:id", func(c echo.Context) error {
		return controllers.GetUser(c)
	})
	e.PUT("/users/:id", func(c echo.Context) error {
		return controllers.UpdateUser(c)
	})
	e.DELETE("/users/:id", func(c echo.Context) error {
		return controllers.DeleteUser(c)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
