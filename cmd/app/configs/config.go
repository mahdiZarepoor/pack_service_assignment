package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"sync"
)

var (
	once   sync.Once
	config Config
)

type App struct {
	Name     string
	Env      string
	Version  string
	URL      string
	Port     string
	LogLevel string
	Debug    bool
}

type Swagger struct {
	Info struct {
		Title       string
		Description string
		Version     string
	}
	Host     string
	Schemes  string
	Username string
	Password string
	Enable   bool
}

// Config represents the application configuration.
type Config struct {
	Swagger Swagger
	App     App
}

// LoadConfig loads configuration from .env file and populates the Config struct.
func LoadConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, fmt.Errorf("error loading .env file: %v", err)
	}

	var app App
	app.Name = os.Getenv("APP_NAME")
	app.Env = os.Getenv("APP_ENV")
	app.Version = os.Getenv("APP_VERSION")
	app.URL = os.Getenv("APP_URL")
	app.Port = os.Getenv("APP_PORT")
	app.Debug = getBoolEnv("APP_DEBUG")
	app.LogLevel = os.Getenv("APP_LOG_LEVEL")

	var swagger Swagger
	swagger.Host = os.Getenv("SWAGGER_HOST")
	swagger.Schemes = os.Getenv("SWAGGER_SCHEMES")
	swagger.Info.Title = os.Getenv("SWAGGER_INFO_TITLE")
	swagger.Info.Description = os.Getenv("SWAGGER_INFO_DESCRIPTION")
	swagger.Info.Version = os.Getenv("SWAGGER_INFO_VERSION")
	swagger.Enable = getBoolEnv("SWAGGER_ENABLE")
	swagger.Username = os.Getenv("SWAGGER_USERNAME")
	swagger.Password = os.Getenv("SWAGGER_PASSWORD")

	return Config{
		App:     app,
		Swagger: swagger,
	}, nil
}

// Helper function to convert string environment variable to bool
func getBoolEnv(key string, defaults ...bool) bool {

	value := os.Getenv(key)
	if value == "" {
		if len(defaults) > 0 && defaults[0] {
			return defaults[0]
		}
	}

	val, _ := strconv.ParseBool(value)
	return val
}

// Helper function to convert string environment variable to int
func getIntEnv(key string, defaultValue int) int {
	val, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return defaultValue
	}
	return val
}

func getStringEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetConfig() Config {
	once.Do(func() {
		var err error
		config, err = LoadConfig()
		if err != nil {
			panic(err)
		}
	})
	return config
}
