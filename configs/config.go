package configs

import (
	"os"

	"github.com/huda7077/gin-and-gorm-boilerplate/pkg/exceptions"
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
	err := godotenv.Load(filenames...)
	exceptions.PanicLogging(err)
	return &configImpl{}
}
