package consumer

import (
	"context"
	"fmt"

	api "github.com/jim-junior/eda/cmd/api"
)

func RunConsumer() error {
	redisClient := api.NewRedisClient()
	for {
		// Get the resource from the message queue
		pubsub := redisClient.Subscribe(context.Background(), "messages")

		ch := pubsub.Channel()

		for msg := range ch {

			fmt.Println("Recieved Message")
			fmt.Println(msg)
		}
	}
}
