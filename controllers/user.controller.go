package controllers

import (
	"crud-echo-postgres-redis/dao"
	"crud-echo-postgres-redis/models"
	"encoding/json"
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
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "GET")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	users, err := dao.GetAllUsers()
	if err != nil {
		log.Fatalf("Unable to get all users. %v", err)
	}

	return c.JSON(http.StatusCreated, users)
}

func GetUser(c echo.Context) error {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "GET")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(req)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	user, err := dao.GetUser(int64(id))
	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}

	return c.JSON(http.StatusCreated, user)
}

func CreateUser(c echo.Context) error {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "POST")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user models.User

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	insertID := dao.InsertUser(user)

	res := response{
		Id:      insertID,
		Message: "User created successfully",
	}

	return c.JSON(http.StatusCreated, res)
}

func UpdateUser(c echo.Context) error {
	res.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "PUT")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(req)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	var user models.User

	err = json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
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
	res.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "DELETE")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(req)

	id, err := strconv.Atoi(params["id"])

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
