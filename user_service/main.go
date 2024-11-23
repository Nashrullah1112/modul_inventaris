package main

import (
	"user/Config"
	"user/Route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,User-Agent,Content-Length",
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	Config.InitDB()

	Route.NewRouter(app)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}

}
