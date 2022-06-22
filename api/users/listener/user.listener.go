package listener

import (
	"crud-echo-postgres-redis/api/users/model"
	"crud-echo-postgres-redis/api/users/service"
	"crud-echo-postgres-redis/config"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
	"time"
)

func createConsumer() *kafka.Consumer {
	env, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading app.env file")
	}

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": env.KafkaSource,
		"group.id":          "crud-golang",
		"auto.offset.reset": "smallest"})

	if err != nil {
		fmt.Printf("Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	return consumer
}

func Run() {
	fmt.Println("Starting Kafka consumer")
	consumer := createConsumer()
	err := consumer.SubscribeTopics([]string{"users"}, nil)

	if err != nil {
		fmt.Printf("Failed to subscribe topic: %s\n", err)
		os.Exit(1)
	}

	run := true
	for run == true {
		ev := consumer.Poll(0)
		switch e := ev.(type) {
		case *kafka.Message:
			fmt.Println(e)
			fmt.Println(string(e.Key))
			fmt.Println(string(e.Value))

			var user = new(model.User)
			err := json.Unmarshal(e.Value, &user)
			if err != nil {
				log.Printf("fail to unmarshal user. %v", err)
				return
			}

			service.CreateUser(user)
		case kafka.PartitionEOF:
			fmt.Printf("%% Reached %v\n", e)
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
		}

		time.Sleep(5 * time.Second)
	}

	err = consumer.Close()
	if err != nil {
		panic(err)
	}
}
