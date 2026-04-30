package fiber_utils

import (
	_ "embed"

	"github.com/gofiber/fiber/v3"
)

////go:embed openapi.yaml
//var openapi []byte

//go:embed swagger.html
var swagger []byte

func WithSwagger(app *fiber.App, openapiSpec []byte) {
	app.Get("/openapi.yaml", func(c fiber.Ctx) error { return c.Send(openapiSpec) })
	app.Get("/docs", func(c fiber.Ctx) error { return c.Type("html").Send(swagger) })
}
