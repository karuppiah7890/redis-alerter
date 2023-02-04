package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisStatus struct {
	IsUp   bool
	Errors []error
}

func GetRedisStatus(host string, port int) RedisStatus {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", host, port),
	})

	val, err := rdb.Ping(ctx).Result()

	if err != nil {
		return RedisStatus{
			IsUp:   false,
			Errors: []error{fmt.Errorf("error occurred while pinging redis: %v", err)},
		}
	}

	if val == "PONG" {
		return RedisStatus{
			IsUp: true,
		}
	}

	return RedisStatus{
		IsUp: false,
	}
}
