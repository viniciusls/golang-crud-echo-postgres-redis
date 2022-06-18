package service

import (
	"crud-echo-postgres-redis/api/users/dao_dynamodb"
	"crud-echo-postgres-redis/api/users/model"
	"crud-echo-postgres-redis/redis"
	"encoding/json"
	"log"
)

func GetAllUsers() ([]model.User, error) {
	cacheKey := "all_users"
	cachedContent, err := redis.Get(cacheKey)
	if err == nil {
		var users []model.User
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

	redis.Set(cacheKey, string(serialized), 0)

	return users, err
}

func GetUser(id string) (model.User, error) {
	cacheKey := "user_" + id
	cachedContent, err := redis.Get(cacheKey)
	if err == nil {
		var user model.User
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

	redis.Set(cacheKey, string(serialized), 0)

	return user, err
}

func CreateUser(user *model.User) string {
	insertId := dao.CreateUser(user)

	_, err := redis.Del("all_users")
	if err != nil {
		log.Printf("Unable to cleanup all_users cache. %v", err)
	}

	return insertId
}

func UpdateUser(id string, user *model.User) int64 {
	rowsAffected := dao.UpdateUser(id, user)

	_, err := redis.Del("all_users", "user_"+id)
	if err != nil {
		log.Printf("Unable to cleanup all_users cache. %v", err)
	}

	return rowsAffected
}

func DeleteUser(id string) int64 {
	rowsAffected := dao.DeleteUser(id)

	_, err := redis.Del("all_users", "user_"+id)
	if err != nil {
		log.Printf("Unable to cleanup all_users cache. %v", err)
	}

	return rowsAffected
}
