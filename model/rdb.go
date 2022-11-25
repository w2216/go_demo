package model

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client
var RdbErr error
var RdbCtx = context.Background()

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "sddphp.cn:19979",
		Password: "20210101", // no password set
		DB:       0,          // use default DB
	})

	err := Rdb.Ping(RdbCtx).Err()
	if err != nil {
		panic(err)
	}
}
