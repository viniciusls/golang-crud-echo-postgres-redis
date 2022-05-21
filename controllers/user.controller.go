package controllers

import (
	"crud-echo-postgres-redis/models"
	"crud-echo-postgres-redis/services"
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
	users, err := services.GetAllUsers()
	if err != nil {
		log.Fatalf("Unable to get all users. %v", err)
	}

	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	user, err := services.GetUser(int64(id))
	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}

	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	insertId := services.CreateUser(user)

	res := response{
		Id:      insertId,
		Message: "User created successfully",
	}

	return c.JSON(http.StatusCreated, res)
}

func UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.FormValue("id"))

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	user := new(models.User)

	if err := c.Bind(user); err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	updatedRows := services.UpdateUser(int64(id), user)

	msg := fmt.Sprintf("User updated successfully. Total rows/records affected %v", updatedRows)

	res := response{
		Id:      int64(id),
		Message: msg,
	}

	return c.JSON(http.StatusCreated, res)
}

func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.FormValue("id"))

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	updatedRows := services.DeleteUser(int64(id))

	msg := fmt.Sprintf("User updated successfully. Total rows/records affected %v", updatedRows)

	res := response{
		Id:      int64(id),
		Message: msg,
	}

	return c.JSON(http.StatusOK, res)
}
