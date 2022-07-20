package k6setenv

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	Env string
}

func (e *Environment) LoadEnvironment() {
	e.Env = os.Getenv("ENV")
	_ = godotenv.Load(fmt.Sprintf("%s.env", e.Env))
}

func (e *Environment) GetEnvironmentVariable(varName string) string {
	return os.Getenv(varName)
}
