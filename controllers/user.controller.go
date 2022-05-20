package controllers

import (
	"crud-echo-postgres-redis/dao"
	"crud-echo-postgres-redis/models"
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
	users, err := dao.GetAllUsers()
	if err != nil {
		log.Fatalf("Unable to get all users. %v", err)
	}

	return c.JSON(http.StatusCreated, users)
}

func GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	user, err := dao.GetUser(int64(id))
	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}

	return c.JSON(http.StatusCreated, user)
}

func CreateUser(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	insertId := dao.InsertUser(user)

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

	var user models.User

	if err := c.Bind(user); err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	updatedRows := dao.UpdateUser(int64(id), user)

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

	updatedRows := dao.DeleteUser(int64(id))

	msg := fmt.Sprintf("User updated successfully. Total rows/records affected %v", updatedRows)

	res := response{
		Id:      int64(id),
		Message: msg,
	}

	return c.JSON(http.StatusCreated, res)
}
