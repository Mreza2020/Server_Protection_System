package Server_Protection_System

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// SaveOtpToRedis
//
// It is for storing temporary code in the Redis system.
//
// Parameters:
//
// email string , otp string , Storage int
//
// NOTE:
//
// - Enter the storage duration based on your needs.
//
// - Duration is based on minutes.
func SaveOtpToRedis(email string, otp string, Storage int) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	RedisClient.Set(ctx, "otp:"+email, otp, time.Duration(Storage)*time.Minute)

}

// # GetOtpFromRedis
//
// # Redis code recipient based on email
//
// # Parameters:
//
// email string , ch chan string
//
// # Returns:
//
// OTP not found And Error retrieving OTP And otpcod
func GetOtpFromRedis(email string, ch chan string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	otp, err := RedisClient.Get(ctx, "otp:"+email).Result()
	if err == redis.Nil {
		ch <- "OTP not found"
	} else if err != nil {
		ch <- "Error retrieving OTP"
	} else {
		ch <- otp
	}

}
