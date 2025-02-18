package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Env          string // dev || prod
	GRPC         GRPCConfig
	RedisAddress string
	RedisDb      int
}

type GRPCConfig struct {
	Port    int
	Timeout time.Duration
}

func MustLoad() *Config {
	loadEnvFile()

	env := getEnv("ENV", "dev")
	grpcPort := getEnvAsInt("GRPC_PORT", 50051)
	grpcTimeout := getEnvAsDuration("GRPC_TIMEOUT", 10*time.Second)
	redisAddress := getEnv("REDIS_ADDRESS", "localhost:5252")
	redisDb := getEnvAsInt("REDIS_DB", 0)

	return &Config{
		Env: env,
		GRPC: GRPCConfig{
			Port:    grpcPort,
			Timeout: grpcTimeout,
		},
		RedisAddress: redisAddress,
		RedisDb:      redisDb,
	}
}

func loadEnvFile() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	valueStr := getEnv(key, "")
	if value, err := time.ParseDuration(valueStr); err == nil {
		return value
	}
	return defaultValue
}
