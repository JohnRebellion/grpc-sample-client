package main

import (
	"grpc-sample-client/envRouting"
	"grpc-sample-client/person"

	"github.com/gofiber/fiber/v2"
)

func main() {
	envRouting.LoadEnv()
	app := fiber.New()
	app.Get("/person", func(c *fiber.Ctx) error {
		people, err := person.GetAll()

		if err == nil {
			return c.JSON(people)
		}

		return err
	})

	app.Listen(":8111")
}
