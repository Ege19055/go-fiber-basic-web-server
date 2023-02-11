package main

import (
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.SendFile("index.html")
	})
	app.Get("/about", func(c*fiber.Ctx){
		c.SendFile("about.html")
	})

	

	app.Listen(":3000")
}