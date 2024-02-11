package config

import (
	"os"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
	"github.com/underbek/examples-go/config"
)

type Config struct {
	config.App
	HTTPServer
	Telegram
}

type HTTPServer struct {
	Port         string        `env:"HTTP_SERVER_PORT" envDefault:":8080"`
	WriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT" envDefault:"15s"`
	ReadTimeout  time.Duration `env:"HTTP_READ_TIMEOUT" envDefault:"15s"`
}

type Telegram struct {
	Token string `env:"TELEGRAM_TOKEN" valid:"required"`
	URL   string `env:"TELEGRAM_URL" valid:"required"`
}

func New[T any]() (T, error) {
	var cfg T

	if envFilePath, ok := os.LookupEnv("ENV_FILE_PATH"); ok {
		if err := godotenv.Load(envFilePath); err != nil {
			return cfg, err
		}
	}

	if err := env.Parse(&cfg); err != nil {
		return cfg, err
	}

	if ok, err := govalidator.ValidateStruct(cfg); !ok {
		return cfg, err
	}

	return cfg, nil
}
