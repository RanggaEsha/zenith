package config

import (
	"aidanwoods.dev/go-paseto"
	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	"log"
)

type (
	// Config contains configuration settings loaded from environment variables.
	Config struct {
		DatabaseHost     string `env:"DB_HOST"`
		DatabasePort     string `env:"DB_PORT"`
		DatabaseName     string `env:"DB_NAME"`
		DatabaseUser     string `env:"DB_USER"`
		DatabasePassword string `env:"DB_PASSWORD"`
		SslMode          string `env:"SSL_MODE"`
		Timezone         string `env:"TIMEZONE"`
		PasswordSalt     string `env:"PASSWORD_SALT"`
		SMTPHost         string `env:"SMTP_HOST"`
		SMTPPort         int    `env:"SMTP_PORT"`
		SMTPUsername     string `env:"SMTP_USERNAME"`
		SMTPPassword     string `env:"SMTP_PASSWORD"`
		RedisHost        string `env:"REDIS_HOST"`
		RedisPort        int    `env:"REDIS_PORT"`
		RedisDatabase    int    `env:"REDIS_DB"`
		RedisUsername    string `env:"REDIS_USERNAME"`
		RedisPassword    string `env:"REDIS_PASSWORD"`
	}
)

var (
	SecretKey = paseto.NewV4AsymmetricSecretKey()
	PublicKey = SecretKey.Public()
)

// NewConfig creates and loads a new Config instance from the environment.
func NewConfig(filenames ...string) *Config {
	config := loadEnvFile(filenames...)
	return &config
}

// loadEnvFile loads the configuration from the provided `.env` files and environment variables.
func loadEnvFile(filenames ...string) Config {
	if err := godotenv.Load(filenames...); err != nil {
		log.Fatalf("Error loading `.env` file: %v", err)
	}

	var config Config
	if _, err := env.UnmarshalFromEnviron(&config); err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

	return config
}
