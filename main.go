package main

import (
	"log"

	"github.com/aaronmurniadi/panacea/controller"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		ServerHeader:      "Fiber",
		AppName:           "Test App v1.0.1",
		Prefork:           false,
		CaseSensitive:     true,
		StrictRouting:     true,
		EnablePrintRoutes: true,
	})

	app.Get("/", controller.HomePageController)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
		return
	}
}
