package setup

import (
	"github.com/AminN77/we-connect/api/controller"
	fiberPkg "github.com/AminN77/we-connect/pkg/fiber"
	"github.com/gofiber/fiber/v2"
)

func SetRouter(c *controller.Controller) *fiber.App {
	app := fiberPkg.NewFiberRouter()

	v1 := app.Group("/api/v1")
	{
		v1.Get("/financialData", c.Get)
		v1.Post("/financialData", c.Add)
	}

	return app
}
