package Server_Protection_System

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

// # Run_Redis_Server
//
// # This is for starting redis
//
// To start:
//
// Run_Redis_Server(ch chan string)
func Run_Redis_Server(ch chan string) {
	if Env_password_Loaded_string("DB_Redis") != "" {
		if Env_password_Loaded_string("DB_Redis_password") != "" {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			RedisClient = redis.NewClient(&redis.Options{
				Addr:     Env_password_Loaded_string("DB_Redis"),
				Password: Env_password_Loaded_string("DB_Redis_password"),
				DB:       0,
			})
			_, err := RedisClient.Ping(ctx).Result()
			if err != nil {
				ch <- "Redis could not be connected"
			}

		}
	}
}
