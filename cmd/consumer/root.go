package consumer

import (
	"context"
	"fmt"

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

			fmt.Println("Received event on topic: " + topic)
			fmt.Println("Event data: " + msg.String())
		}
	}
}
