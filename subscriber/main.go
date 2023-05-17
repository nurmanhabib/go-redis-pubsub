package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func main() {
	subscriber := redisClient.Subscribe(context.Background(), "my-channel")

	for {
		msg, err := subscriber.ReceiveMessage(context.Background())
		if err != nil {
			panic(err)
		}

		log.Printf("Received message: %v", msg)
	}
}
