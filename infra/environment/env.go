package environment

import (
	"github.com/joho/godotenv"
	"os"
)

type Env struct{}

func (e *Env) load() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}
	return nil
}

func (e *Env) GetVariable(key string) (string, error) {
	if err := e.load(); err != nil {
		return "", err
	}
	return os.Getenv(key), nil
}
