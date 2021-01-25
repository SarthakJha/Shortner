package shortner

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// CheckCache check for the url in cache
func (rc *RedisClient) CheckCache(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	id := c.Params("id")
	fmt.Println(id)
	res := rc.Client.Get(ctx, id)
	if res.Err() != nil {
		cancel()
		fmt.Println("not found in cache")
		return c.Next()
	}
	cancel()
	fmt.Println("found in cache")
	return c.Redirect(res.Val()) // wont redirect to res.String()
}

// WriteCache writes validated request into cache
func (rc *RedisClient) WriteCache(c *fiber.Ctx) error {
	reqBody := URLRequest{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	if err := c.BodyParser(&reqBody); err != nil {
		cancel()
		fmt.Println("error parsing")
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	res := rc.Client.Set(ctx, reqBody.Event, reqBody.URLMain, 10*time.Second)
	if res.Err() != nil {
		fmt.Println("error SET")
		cancel()
		return c.Status(400).JSON(fiber.Map{
			"error": "error in redis",
		})
	}
	cancel()
	fmt.Println("cache Written")
	return c.Next()
}

// ReWriteCache rewrites the cache after key is expired
func (rc *RedisClient) ReWriteCache(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	id := c.Params("id")
	fmt.Printf("rewriting: %v\n", c.Locals("main"))
	res := rc.Client.Set(ctx, id, c.Locals("main"), 10*time.Second)
	if res.Err() != nil {
		fmt.Println("error SET")
		cancel()
		return c.Status(400).JSON(fiber.Map{
			"error": "error in redis",
		})
	}
	cancel()
	fmt.Println("cache ReWritten")
	return c.Redirect(c.Locals("main").(string))
}
