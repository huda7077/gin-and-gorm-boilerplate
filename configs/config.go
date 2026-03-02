package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type configImpl struct{}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Config {
	if err := godotenv.Load(filenames...); err != nil {
		log.Println("[config] .env file not found, using system environment variables")
	}
	return &configImpl{}
}
