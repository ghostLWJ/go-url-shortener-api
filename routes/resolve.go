package routes

import (
	"main/cache"

	"github.com/gofiber/fiber/v2"
)

func ResolveURL(c *fiber.Ctx, cache *cache.Cache) error {
	url := c.Params("url")

	value := cache.Get(url)

	if value == nil {
		return c.Status(fiber.StatusNotFound).SendString("URL not found")
	}

	return c.Redirect(value.(string))
}
