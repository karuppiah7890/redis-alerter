package config

import (
	"fmt"
	"os"
	"strconv"
)

// All configuration is through environment variables

const REDIS_NAME_ENV_VAR = "REDIS_NAME"
const DEFAULT_REDIS_NAME = "Redis"
const REDIS_HOST_ENV_VAR = "REDIS_HOST"
const DEFAULT_REDIS_HOST = "localhost"
const REDIS_PORT_ENV_VAR = "REDIS_PORT"
const DEFAULT_REDIS_PORT = 6379
const ENVIRONMENT_NAME_ENV_VAR = "ENVIRONMENT_NAME"
const DEFAULT_ENVIRONMENT_NAME = "Production"
const SLACK_TOKEN_ENV_VAR = "SLACK_TOKEN"
const SLACK_CHANNEL_ENV_VAR = "SLACK_CHANNEL"

type Config struct {
	redisName       string
	redisHost       string
	redisPort       int
	environmentName string
	slackToken      string
	slackChannel    string
}

func NewConfigFromEnvVars() (*Config, error) {
	redisName := getRedisName()

	redisHost := getRedisHost()

	redisPort, err := getRedisPort()
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting redis port: %v", err)
	}

	environmentName := getEnvironmentName()

	slackToken, err := getSlackToken()
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting slack token: %v", err)
	}

	slackChannel, err := getSlackChannel()
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting slack channel: %v", err)
	}

	return &Config{
		redisName:       redisName,
		redisHost:       redisHost,
		redisPort:       redisPort,
		environmentName: environmentName,
		slackToken:      slackToken,
		slackChannel:    slackChannel,
	}, nil
}

// Get optional name for the Redis instance. Default is "Redis".
// This will be used in the alert messages
func getRedisName() string {
	redisName, ok := os.LookupEnv(REDIS_NAME_ENV_VAR)
	if !ok {
		return DEFAULT_REDIS_NAME
	}

	return fmt.Sprintf("%s (Redis)", redisName)
}

// Get redis host name / DNS / IP. Default is "localhost"
func getRedisHost() string {
	redisHost, ok := os.LookupEnv(REDIS_HOST_ENV_VAR)
	if !ok {
		return DEFAULT_REDIS_HOST
	}

	return redisHost
}

// Get redis port number. Default is 6379
func getRedisPort() (int, error) {
	redisPortEnvStr, ok := os.LookupEnv(REDIS_PORT_ENV_VAR)
	if !ok {
		return DEFAULT_REDIS_PORT, nil
	}

	redisPort, err := strconv.Atoi(redisPortEnvStr)
	if err != nil {
		return 0, fmt.Errorf("%s environment variable value (%s) cannot be converted to integer", REDIS_PORT_ENV_VAR, redisPortEnvStr)
	}

	return redisPort, nil
}

// Get optional environment name for the environment where
// the services are running. Default is "Production". This name will
// be used in the alert messages
func getEnvironmentName() string {
	environmentName, ok := os.LookupEnv(ENVIRONMENT_NAME_ENV_VAR)
	if !ok {
		return DEFAULT_ENVIRONMENT_NAME
	}

	return environmentName
}

func getSlackToken() (string, error) {
	slackToken, ok := os.LookupEnv(SLACK_TOKEN_ENV_VAR)
	if !ok {
		return "", fmt.Errorf("%s environment variable is not defined and is required. Please define it", SLACK_TOKEN_ENV_VAR)
	}
	return slackToken, nil
}

func getSlackChannel() (string, error) {
	slackChannel, ok := os.LookupEnv(SLACK_CHANNEL_ENV_VAR)
	if !ok {
		return "", fmt.Errorf("%s environment variable is not defined and is required. Please define it", SLACK_CHANNEL_ENV_VAR)
	}
	return slackChannel, nil
}

func (c *Config) GetRedisName() string {
	return c.redisName
}

func (c *Config) GetRedisHost() string {
	return c.redisHost
}

func (c *Config) GetRedisPort() int {
	return c.redisPort
}

func (c *Config) GetEnvironmentName() string {
	return c.environmentName
}

func (c *Config) GetSlackToken() string {
	return c.slackToken
}

func (c *Config) GetSlackChanel() string {
	return c.slackChannel
}
