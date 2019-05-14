package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ServiceConfig is set of common configs that every service needs to work
type ServiceConfig struct {
	NatsHost   string
	NatsPort   int
	TracerHost string
	TracerPort int
	Port       int
}

func (c *ServiceConfig) GetNatsURL() string {
	return fmt.Sprintf("nats://%v:%v", c.NatsHost, c.NatsPort)
}

func (c *ServiceConfig) GetTracerURL() string {
	return fmt.Sprintf("%v:%v", c.TracerHost, c.TracerPort)
}

// New returns a new Config struct
func New() *ServiceConfig {
	return &ServiceConfig{
		NatsHost:   GetEnv("NATS_HOST", "nats"),
		NatsPort:   GetEnvAsInt("NATS_PORT", 4222),
		TracerHost: GetEnv("TRACER_HOST", "tracer"),
		TracerPort: GetEnvAsInt("TRACER_PORT", 6831),
		Port:       GetEnvAsInt("PORT", 5000),
	}
}

// GetEnv is a function to read an environment or return a default value
func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// GetEnvAsInt reads an environment variable into integer or return a default value
func GetEnvAsInt(name string, defaultVal int) int {
	valueStr := GetEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// GetEnvAsBool reads an environment variable into a bool or return default value
func GetEnvAsBool(name string, defaultVal bool) bool {
	valStr := GetEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

// GetEnvAsSlice reads an environment variable into a string slice or return default value
func GetEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := GetEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
