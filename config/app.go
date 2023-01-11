package config

import (
	"os"

	"govel/app/exception"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
	LoadEnv(filenames ...string)
}

type configImpl struct{}

func New() Config {
	return &configImpl{}
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func (config *configImpl) LoadEnv(filenames ...string) {
	err := godotenv.Load(filenames...)
	exception.PanicIfNeeded(err)
}
