package router

import (
	"crud-echo-postgres-redis/api/users/controller"
	_ "crud-echo-postgres-redis/docs"
	nr "crud-echo-postgres-redis/newrelic"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

func New() *echo.Echo {
	fmt.Println("Starting Echo app")

	newRelicApp, err := nr.InitNewRelicApp()

	// logs an error if New Relic failed to start
	if err != nil {
		log.Error(err)
	}

	newRelicApp.RecordCustomEvent("InitRouter", map[string]interface{}{
		"myString": "hello",
		"myFloat":  0.603,
		"myInt":    123,
		"myBool":   true,
	})

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
	e.GET("/health", HealthCheck)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/users", func(c echo.Context) error {
		return controller.GetAllUsers(c)
	})
	e.POST("/users", func(c echo.Context) error {
		return controller.CreateUser(c)
	})
	e.GET("/users/:id", func(c echo.Context) error {
		return controller.GetUser(c)
	})
	e.PUT("/users/:id", func(c echo.Context) error {
		return controller.UpdateUser(c)
	})
	e.DELETE("/users/:id", func(c echo.Context) error {
		return controller.DeleteUser(c)
	})

	return e
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags HealthCheck
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}
