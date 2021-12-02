package venv

import (
	"github.com/georgettica/venv/mock"
	"github.com/georgettica/venv/os"
)

type Env interface {
	Environ() []string
	Getenv(key string) string
	LookupEnv(key string) (string, bool)
	Setenv(key, value string) error
	Clearenv()
}

func OS() Env {
	return os.New()
}

func Mock() Env {
	return mock.New()
}
