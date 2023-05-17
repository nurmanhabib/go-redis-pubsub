package main

import (
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := r.URL.Query().Get("msg")

		redisClient.Publish(r.Context(), "my-channel", message)

		fmt.Fprintf(w, "Published: %s", message)
	})

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
