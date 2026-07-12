package main

import (
	"time"

	"github.com/cockroachdb/errors"

	"github.com/gofiber/fiber/v3"
	"github.com/teadove/teasutils/fiberutils"
)

func main() {
	app := fiber.New(fiber.Config{Immutable: true, ErrorHandler: fiberutils.ErrHandler()})
	app.Use(fiberutils.MiddlewareLogger())
	app.Use(fiberutils.MiddlewareCtxTimeout(time.Second))

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/err", func(_ fiber.Ctx) error {
		return errors.New("error occurred")
	})
	app.Get("/parse-err", func(c fiber.Ctx) error {
		return c.JSON(func() {})
	})

	err := app.Listen(":8000")
	if err != nil {
		panic(err)
	}
}
