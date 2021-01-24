package main

import (
	"context"
	"log"
	"time"

	"github.com/SarthakJha/shawty/shortner"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err.Error())
	}
	db, err := shortner.SeedDb()
	if err != nil {
		panic(err.Error())
	}
	app := fiber.New()
	app.Post("/", db.CreateShortURL)
	app.Get("/:id", db.RedirectRequest)
	defer db.Client.Disconnect(ctx)
	defer cancel()
	log.Fatalln(app.Listen(":3000"))
}
