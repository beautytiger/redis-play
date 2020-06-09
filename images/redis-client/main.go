package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()
	ctx := context.Background()
	// pong, err := client.Ping(ctx).Result()
	// fmt.Println(pong, err)
	// Output: PONG <nil>
	limit := 1 << 16
	fmt.Println(limit)
	for i := 0; i < limit; i++ {
		err := client.Set(ctx, strconv.Itoa(i), uuid.New().String(), 0).Err()
		if err != nil {
			panic(err)
		}
	}
	// for j := 0; j < limit; j++ {
	// 	v, err := client.Get(ctx, strconv.Itoa(j)).Result()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(v)
	// }
}
