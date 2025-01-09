package consumer

import (
	"context"
	"fmt"

	api "github.com/jim-junior/eda/cmd/api"
)

func RunConsumer(topic string) error {
	redisClient := api.NewRedisClient()
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
