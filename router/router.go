package router

import (
	"crud-echo-postgres-redis/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.POST},
	}))

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

	return e
}
