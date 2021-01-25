package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/SarthakJha/shawty/shortner"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err.Error())
	}
	db, err := shortner.SeedDb()
	if err != nil {
		log.Fatalln(err)
	}
	redisClient, err := shortner.InitRedis()
	defer func() {
		err := redisClient.Client.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		// fmt.Println(c.IP()) // 127.0.0.1 (ip of the client)
		return c.Next()
	})

	app.Post("/api/create", db.ValidateRequest, redisClient.WriteCache, db.CreateShortURL)

	app.Get("/:id", redisClient.CheckCache, db.RedirectRequest, redisClient.ReWriteCache)

	app.All("*", func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"error": "Route not found!",
		})
	})
	defer func() {
		err := db.Client.Disconnect(ctx)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Disconnected from mongo...")
	}()
	defer cancel()
	log.Fatalln(app.Listen(":" + os.Getenv("PORT")))
}
