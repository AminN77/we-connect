package setup

import (
	"github.com/AminN77/we-connect/api/controller"
	"github.com/AminN77/we-connect/docs"
	fiberPkg "github.com/AminN77/we-connect/pkg/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

func SetRouter(c *controller.Controller) *fiber.App {
	app := fiberPkg.NewFiberRouter()
	docs.SwaggerInfo.BasePath = "/api/v1"

	app.Use(recover.New())

	v1 := app.Group("/api/v1")
	{
		v1.Get("/financialData", c.Get)
	}

	app.Get("/swagger/*", swagger.HandlerDefault)
	return app
}
