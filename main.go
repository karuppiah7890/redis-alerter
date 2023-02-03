package main

import "fmt"

func main() {
	fmt.Printf("redis-alerter")

	// Get optional custom emoji shortcut in text. Default is "", ie
	// no shortcut. This need not have ':' prefix and suffix, if it's there
	// also it's fine

	// Get optional name for Redis. Default is "Redis".
	// This will be used in the alert messages

	// Get optional environment name for the environment where
	// Redis is running. Default is "production". This name will
	// be used in the alert messages

	// Get redis host name / DNS / IP
	// Get redis port number

	// TODO: Connect to Redis

	// If error like connection refused - send alert that Redis is down
	// "[${redis-name} (Redis) | Redis ][${:redis:}] is down in ${envirnment-name} environment"

	// Provide sample Redis image for making a custom emoji
	// to show in slack alert message using emoji shortcut

	// Provide sample Bot image

	// Log all things
}
