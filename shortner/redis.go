package shortner

import (
	"github.com/go-redis/redis/v8"
)

// RedisClient contains redis client
type RedisClient struct {
	Client *redis.Client
}

// InitRedis initialises redis client
func InitRedis() (RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	redisClient := RedisClient{
		Client: client,
	}
	return redisClient, nil
}
