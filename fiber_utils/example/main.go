package main

import (
	"github.com/pkg/errors"

	"github.com/gofiber/fiber/v2"
	"github.com/teadove/teasutils/fiber_utils"
)

func main() {
	app := fiber.New(fiber.Config{Immutable: true, ErrorHandler: fiber_utils.ErrHandler()})
	app.Use(fiber_utils.MiddlewareLogger())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/err", func(c *fiber.Ctx) error {
		return errors.New("error occured")
	})
	app.Get("/parse-err", func(c *fiber.Ctx) error {
		return c.JSON(func() {})
	})

	err := app.Listen(":8000")
	if err != nil {
		panic(err)
	}
}
