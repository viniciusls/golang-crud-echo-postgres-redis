package dao

import (
	"crud-echo-postgres-redis/api/users/model"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"log"
	"strconv"
)

type Item struct {
	Id       string
	Name     string
	Location string
	Age      int64
}

const tableName = "users"

func createConnection() *dynamodb.DynamoDB {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	fmt.Println("Successfully connected to database!")

	return svc
}

func GetAllUsers() ([]model.User, error) {
	_ = createConnection()

	var users []model.User

	// TBD

	return users, nil
}

func GetUser(id string) (model.User, error) {
	db := createConnection()

	var user model.User

	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}

	if result.Item == nil {
		msg := "Could not find '" + id + "'"
		return user, errors.New(msg)
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		log.Fatalf("Failed to unmarshal Record, %v", err)
	}

	return user, err
}

func CreateUser(user *model.User) string {
	db := createConnection()

	id := uuid.New().String()
	item := Item{
		Id:       id,
		Name:     user.Name,
		Location: user.Location,
		Age:      user.Age,
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Fatalf("Got error marshalling new user item: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = db.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}

	return id
}

func UpdateUser(id string, user *model.User) int64 {
	db := createConnection()

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":name": {
				S: aws.String(user.Name),
			},
			":location": {
				S: aws.String(user.Location),
			},
			":age": {
				N: aws.String(strconv.FormatInt(user.Age, 10)),
			},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(id),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set Name = :name, Location = :location, Age = :age"),
	}

	_, err := db.UpdateItem(input)
	if err != nil {
		log.Fatalf("Got error calling UpdateItem: %s", err)
	}

	return 1
}

func DeleteUser(id string) int64 {
	db := createConnection()

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(tableName),
	}

	_, err := db.DeleteItem(input)
	if err != nil {
		log.Fatalf("Got error calling DeleteItem: %s", err)
	}

	return 1
}
