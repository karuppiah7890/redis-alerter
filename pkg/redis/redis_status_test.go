package redis_test

import (
	"context"
	"testing"

	"github.com/karuppiah7890/redis-alerter/pkg/redis"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestGetRedisStatus(t *testing.T) {
	t.Run("Redis is up", func(t *testing.T) {
		ctx := context.Background()
		req := testcontainers.ContainerRequest{
			Image:        "redis:7",
			ExposedPorts: []string{"6379/tcp"},
			WaitingFor:   wait.ForLog("Ready to accept connections"),
		}
		redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		})
		if err != nil {
			t.Error(err)
		}
		defer func() {
			if err := redisC.Terminate(ctx); err != nil {
				t.Fatalf("failed to terminate container: %s", err.Error())
			}
		}()

		host, err := redisC.Host(ctx)
		if err != nil {
			t.Errorf("error occurred: expected no errors while getting test redis host but got one: %v", err)
		}

		port, err := redisC.MappedPort(ctx, "6379/tcp")
		if err != nil {
			t.Errorf("error occurred: expected no errors while getting test redis port but got one: %v", err)
		}

		redisStatus := redis.GetRedisStatus(host, port.Int())

		if redisStatus.IsUp == false {
			t.Error("error occurred: expected Redis to be up and running, but the Redis is not up")
		}
	})

	t.Run("Redis is down", func(t *testing.T) {
		redisStatus := redis.GetRedisStatus("localhost", 6379)

		if redisStatus.IsUp == true {
			t.Error("error occurred: expected Redis to be down, but the Redis is up and running")
		}

		if redisStatus.Errors == nil {
			t.Error("error occurred: expected atleast one error but got none")
		}
	})
}
