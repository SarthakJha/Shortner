package shortner

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateShortURL creates the short url
func (mg *MongoConn) CreateShortURL(c *fiber.Ctx) error {
	fmt.Println("entering db (redir)")
	reqBody := URLRequest{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	if err := c.BodyParser(&reqBody); err != nil {
		cancel()
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	insDoc := CreateURL{
		URLMain: reqBody.URLMain,
		ShortID: reqBody.Event,
	}
	_, err := mg.Db.Collection(os.Getenv("COLLECTION_NAME")).InsertOne(ctx, insDoc)
	if err != nil {
		cancel()
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	baseString := "http://localhost:3000/"
	cancel()
	fmt.Println("short url created")
	return c.Status(201).JSON(fiber.Map{
		"short_url": baseString + insDoc.ShortID,
	})
}

// RedirectRequest redirects the request
func (mg *MongoConn) RedirectRequest(c *fiber.Ctx) error {
	fmt.Println("entering db(redirect)")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	id := c.Params("id")

	res := mg.Db.Collection(os.Getenv("COLLECTION_NAME")).FindOne(ctx, bson.D{
		{Key: "short_id", Value: id},
	})
	if res.Err() != nil {
		cancel()
		return c.Status(404).JSON(fiber.Map{
			"error": "Invalid Route",
		})
	}

	model := DecodeRequest{}
	err := res.Decode(&model)
	if err != nil {
		cancel()
		return c.Status(400).JSON(fiber.Map{
			"error": "error parsing body. recheck request body",
		})
	}
	cancel()
	fmt.Println("going next")
	c.Locals("main", model.URLMain)
	return c.Next()
}
