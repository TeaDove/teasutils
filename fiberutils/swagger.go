package fiberutils

import (
	_ "embed"
	"encoding/json"

	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v3"
	"gopkg.in/yaml.v3"
)

//go:embed swagger.html
var swagger []byte

// MustWithSwagger is like WithSwagger but panics if specYAML cannot be parsed.
// Intended for wiring routes at startup with a spec embedded at compile time.
func MustWithSwagger(app *fiber.App, specYAML []byte) {
	err := WithSwagger(app, specYAML)
	if err != nil {
		panic(errors.Wrap(err, "with swagger"))
	}
}

// WithSwagger registers documentation routes on app from an OpenAPI spec
// given as YAML: /openapi.yaml (the spec verbatim), /openapi.json (the spec
// converted to JSON) and /docs (the embedded Swagger UI page).
// It returns an error if specYAML is not valid YAML.
func WithSwagger(app *fiber.App, specYAML []byte) error {
	var spec map[string]any

	err := yaml.Unmarshal(specYAML, &spec)
	if err != nil {
		return errors.Wrap(err, "failed to decode spec")
	}

	specJSON, err := json.Marshal(spec)
	if err != nil {
		return errors.Wrap(err, "failed to marshal spec")
	}

	app.Get("/openapi.yaml", func(c fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, "application/yaml")

		return c.Send(specYAML)
	})
	app.Get("/openapi.json", func(c fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

		return c.Send(specJSON)
	})
	app.Get("/docs", func(c fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

		return c.Send(swagger)
	})

	return nil
}
