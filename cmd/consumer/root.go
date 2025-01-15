package consumer

import (
	"context"
	"fmt"
	"time"

	api "github.com/jim-junior/eda/cmd/api"
)

func RunConsumer(topic string) error {
	redisClient := api.NewRedisClient()
	fmt.Println("Consumer subscribed to `" + topic + "` Topic")
	for {
		// Get the resource from the message queue
		pubsub := redisClient.Subscribe(context.Background(), topic)

		ch := pubsub.Channel()

		for msg := range ch {

			fmt.Println("Recieved Message")
			fmt.Println(msg)
		}
	}
}

func RunDFEConsumer(topic string) error {
	redisClient := api.NewRedisClient()
	fmt.Println("This consumer demonstrates Deffered Execution and Eventual Consistency")
	fmt.Println("Consumer subscribed to `" + topic + "` Topic")
	for {
		// Get the resource from the message queue
		pubsub := redisClient.Subscribe(context.Background(), topic)

		ch := pubsub.Channel()

		for msg := range ch {

			time.Sleep(3 * time.Second)

			fmt.Println("Recieved Message")
			fmt.Println(msg)
		}
	}
}
