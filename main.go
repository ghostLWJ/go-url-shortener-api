package main

import (
	"fmt"
	"log"
	"main/cache"
	"main/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	cacheInstance := cache.NewCache()

	app := fiber.New()

	app.Use(logger.New())

	setupRoutes(app, cacheInstance)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}

func setupRoutes(app *fiber.App, cache *cache.Cache) {
	app.Get("/:url", func(c *fiber.Ctx) error {
		return routes.ResolveURL(c, cache)
	})

	app.Post("/shorten", func(c *fiber.Ctx) error {
		return routes.ShortenURL(c, cache)
	})
}
