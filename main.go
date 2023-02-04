package main

import (
	"fmt"
	"log"

	"github.com/karuppiah7890/redis-alerter/pkg/config"
	"github.com/karuppiah7890/redis-alerter/pkg/redis"
	"github.com/karuppiah7890/redis-alerter/pkg/slack"
)

func main() {
	c, err := config.NewConfigFromEnvVars()
	if err != nil {
		log.Fatalf("error occurred while getting configuration from environment variables: %v", err)
	}

	// TODO: Use Mocks to test the integration with ease for different cases with unit tests
	redisStatus := redis.GetRedisStatus(c.GetRedisHost(), c.GetRedisPort())

	if !redisStatus.IsUp {
		message := fmt.Sprintf("Critical alert :rotating_light:! %s is down in %s environment :rotating_light:", c.GetRedisName(), c.GetEnvironmentName())
		// TODO: Use Mocks to test the integration with ease for different cases with unit tests
		err := slack.SendMessage(c.GetSlackToken(), c.GetSlackChanel(), message)
		if err != nil {
			log.Fatalf("error occurred while sending slack alert message: %v", err)
		}
	}

	// Get optional custom emoji shortcut in text. Default is "", ie
	// no shortcut. This need not have ':' prefix and suffix, if it's there
	// also it's fine

	// Get optional environment name for the environment where
	// Redis is running. Default is "production". This name will
	// be used in the alert messages

	// Get redis host name / DNS / IP
	// Get redis port number

	// TODO: Connect to Redis

	// If error like connection refused - send alert that Redis is down
	// "Critical alert :rotating_light: ! [${redis-name} (Redis) | Redis ][${:redis:}] is down in ${envirnment-name} environment"

	// Provide sample Redis image for making a custom emoji
	// to show in slack alert message using emoji shortcut

	// Provide sample Bot image

	// Log all things
}
