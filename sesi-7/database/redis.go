package database

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func ConnectRedis(ctx context.Context, host, password string) (client *redis.Client, err error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
	})

	err = rdb.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}
