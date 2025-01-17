package consumer

import (
	"context"
	"fmt"
	"time"

	api "github.com/jim-junior/eda/cmd/api"
)

func RunMasterConsumer() error {
	ctx := context.Background()
	redisClient := api.NewRedisClient()
	fmt.Println("Master Node running")
	for {
		// Get the resource from the message queue
		pubsub := redisClient.Subscribe(ctx, "master")

		ch := pubsub.Channel()

		for msg := range ch {

			fmt.Println("Recieved Message")
			_, err := redisClient.LPush(ctx, "message:queue", msg.String()).Result()

			if err != nil {
				panic(err)
			}
		}
	}
}

func RunDFEConsumer() error {
	ctx := context.Background()
	redisClient := api.NewRedisClient()
	fmt.Println("This consumer demonstrates Deffered Execution and Eventual Consistency")

	for {
		// Get the resource from the message queue

		fmt.Println("Fetching Events from Event Queue")

		len, err := redisClient.LLen(ctx, "message:queue").Result()

		if err != nil {
			panic(err)
		}

		if len == 0 {
			time.Sleep(10 * time.Second)
			continue
		}

		msg, err := redisClient.RPop(ctx, "message:queue").Result()

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(msg)

		time.Sleep(10 * time.Second)

	}
}
