package main

import (
	"context"
	"fmt"

	"encoding/json"

	"github.com/go-redis/redis/v8"
)

type TimeStamp struct {
	SecsSinceEpoch int64

	NanosSinceEpoch int64
}

type CurrencyPrice struct {
	Currency string

	Id string

	Price string

	Name string

	LastUpdated TimeStamp
}

func main() {

	fmt.Println(("what's this go?? Just hello"))

	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	val, err := rdb.Get(ctx, "nomics_price_SOLUSD").Result()
	if err != nil {
		panic(err)
	}

	b := []byte(val) // convert string into byte slice

	var curr CurrencyPrice

	err = json.Unmarshal(b, &curr)

	if err != nil {
		panic(err)
	}

	fmt.Println("Price of SOL is :", curr.Price)

}
