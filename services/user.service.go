package services

import (
	"crud-echo-postgres-redis/dao"
	"crud-echo-postgres-redis/helper"
	"crud-echo-postgres-redis/models"
	"encoding/json"
	"log"
)

func GetAllUsers() ([]models.User, error) {
	cachedContent, err := helper.Get("all_users")
	if err == nil {
		var users []models.User
		if err := json.Unmarshal([]byte(cachedContent), &users); err != nil {
			log.Fatalf("Unable to convert cached content to users. %v", err)
		}

		return users, nil
	}

	users, err := dao.GetAllUsers()

	serialized, err := json.Marshal(users)
	if err != nil {
		log.Fatalf("Unable to convert array to string. %v", err)
	}

	helper.Set("all_users", string(serialized), 0)

	return users, err
}

func GetUser(id int64) (models.User, error) {
	return dao.GetUser(id)
}

func CreateUser(user *models.User) int64 {
	return dao.CreateUser(user)
}

func UpdateUser(id int64, user *models.User) int64 {
	return dao.UpdateUser(id, user)
}

func DeleteUser(id int64) int64 {
	return dao.DeleteUser(id)
}
