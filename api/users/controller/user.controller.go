package controller

import (
	"crud-echo-postgres-redis/api/users/model"
	"crud-echo-postgres-redis/api/users/service"
	nr "crud-echo-postgres-redis/newrelic"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type response struct {
	Id      string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func GetAllUsers(c echo.Context) error {
	newRelicApp := nr.GetNewRelicApp()

	newRelicApp.RecordCustomEvent("GetAllUsers", map[string]interface{}{
		"myString": "hello",
		"myFloat":  0.603,
		"myInt":    123,
		"myBool":   true,
	})

	users, err := service.GetAllUsers()
	if err != nil {
		log.Fatalf("Unable to get all users. %v", err)
	}

	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")

	user, err := service.GetUser(id)
	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}

	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	user := new(model.User)

	if err := c.Bind(user); err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	insertId := service.CreateUser(user)

	res := response{
		Id:      insertId,
		Message: "User created successfully",
	}

	return c.JSON(http.StatusCreated, res)
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")

	user := new(model.User)

	if err := c.Bind(user); err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	updatedRows := service.UpdateUser(id, user)

	msg := fmt.Sprintf("User updated successfully. Total rows/records affected %v", updatedRows)

	res := response{
		Id:      id,
		Message: msg,
	}

	return c.JSON(http.StatusCreated, res)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")

	updatedRows := service.DeleteUser(id)

	msg := fmt.Sprintf("User deleted successfully. Total rows/records affected %v", updatedRows)

	res := response{
		Id:      id,
		Message: msg,
	}

	return c.JSON(http.StatusOK, res)
}
