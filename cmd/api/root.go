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
		return c.SendString("Send a post request to /send-sms")
	})

	redisClient := NewRedisClient()

	app.Post("/send-sms", func(c *fiber.Ctx) error {
		// Publish to redis channel for driver to work on it
		errf := redisClient.Publish(context.Background(), "messages",
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
