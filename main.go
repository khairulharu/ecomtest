package main

import (
	"ecomtest/internal/api"
	"ecomtest/internal/component"
	"ecomtest/internal/config"
	"ecomtest/internal/repository"
	"ecomtest/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config := config.Get()
	dbConnection := component.GetDatabaseConnection(config)

	productRepository := repository.NewProduct(dbConnection)

	productService := service.NewProduct(productRepository)

	app := fiber.New()
	app.Use(logger.New())

	api.NewProduct(productService, app)

	app.Listen(config.SRV.Host + ":" + config.SRV.Port)
}
