package controller

import (
	"crud-echo-postgres-redis/api/users/model"
	"crud-echo-postgres-redis/api/users/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type response struct {
	Id      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func GetAllUsers(c echo.Context) error {
	users, err := service.GetAllUsers()
	if err != nil {
		log.Fatalf("Unable to get all users. %v", err)
	}

	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	user, err := service.GetUser(int64(id))
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
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	user := new(model.User)

	if err := c.Bind(user); err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	updatedRows := service.UpdateUser(int64(id), user)

	msg := fmt.Sprintf("User updated successfully. Total rows/records affected %v", updatedRows)

	res := response{
		Id:      int64(id),
		Message: msg,
	}

	return c.JSON(http.StatusCreated, res)
}

func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	updatedRows := service.DeleteUser(int64(id))

	msg := fmt.Sprintf("User deleted successfully. Total rows/records affected %v", updatedRows)

	res := response{
		Id:      int64(id),
		Message: msg,
	}

	return c.JSON(http.StatusOK, res)
}
