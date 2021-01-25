package shortner

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// ValidateRequest checks req is unique
func (mg *MongoConn) ValidateRequest(c *fiber.Ctx) error {
	reqBody := URLRequest{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	if err := c.BodyParser(&reqBody); err != nil {
		cancel()
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res := mg.Db.Collection(os.Getenv("COLLECTION_NAME")).FindOne(ctx, bson.D{
		{Key: "short_id", Value: reqBody.Event},
	})
	if res.Err() == nil {
		cancel()
		return c.Status(400).JSON(fiber.Map{
			"error": "Event already exists! Try with different event",
		})
	}
	fmt.Println("request validated")
	cancel()
	return c.Next()
}
