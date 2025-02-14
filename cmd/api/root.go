/*
Copyright © 2024 Cranom Technologies Limited info@cranom.tech
*/
package api

import (
	"context"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Send a post request to /{topic}. Replace `topic` with the topic you would like the event to be published. The post request body/data payload should be event payload you want to send")
	})

	redisClient := NewRedisClient()

	app.Post("/topics/:topic", func(c *fiber.Ctx) error {
		// Publish to redis channel for driver to work on it
		topic := c.Params("topic")

		if topic == "master" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "master is a reserved topic",
			})
		}
		errf := redisClient.Publish(context.Background(), topic,
			c.Body(),
		)
		if errf.Err() != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": errf.Err().Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "message published",
		})
	})

	app.Post("/deferred-exec", func(c *fiber.Ctx) error {
		// Publish to redis channel for driver to work on it
		errf := redisClient.Publish(context.Background(), "master",
			c.Body(),
		)
		if errf.Err() != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": errf.Err().Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "message published",
		})
	})

	app.Listen(":3000")

}
