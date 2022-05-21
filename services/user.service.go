package services

import (
	"crud-echo-postgres-redis/dao"
	"crud-echo-postgres-redis/helper"
	"crud-echo-postgres-redis/models"
	"encoding/json"
	"log"
	"strconv"
)

func GetAllUsers() ([]models.User, error) {
	cacheKey := "all_users"
	cachedContent, err := helper.Get(cacheKey)
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

	helper.Set(cacheKey, string(serialized), 0)

	return users, err
}

func GetUser(id int64) (models.User, error) {
	cacheKey := "user_" + strconv.Itoa(int(id))
	cachedContent, err := helper.Get(cacheKey)
	if err == nil {
		var user models.User
		if err := json.Unmarshal([]byte(cachedContent), &user); err != nil {
			log.Fatalf("Unable to convert cached content to user. %v", err)
		}

		return user, nil
	}

	user, err := dao.GetUser(id)

	serialized, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Unable to convert obj to string. %v", err)
	}

	helper.Set(cacheKey, string(serialized), 0)

	return user, err
}

func CreateUser(user *models.User) int64 {
	insertId := dao.CreateUser(user)

	_, err := helper.Del("all_users")
	if err != nil {
		log.Printf("Unable to cleanup all_users cache. %v", err)
	}

	return insertId
}

func UpdateUser(id int64, user *models.User) int64 {
	return dao.UpdateUser(id, user)
}

func DeleteUser(id int64) int64 {
	return dao.DeleteUser(id)
}
