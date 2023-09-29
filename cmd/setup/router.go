// Package setup gathers the different components setup procedure which will be used in the entrypoint
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

	// middlewares
	app.Use(recover.New())

	// routes
	v1 := app.Group("/api/v1")
	{
		v1.Get("/financialData", c.Get)
	}

	// swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	app.Get("/swagger/*", swagger.HandlerDefault)

	return app
}
