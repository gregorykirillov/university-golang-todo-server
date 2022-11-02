package initializers

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables(key string) string {
	godotenv.Load(".env")

	return os.Getenv(key)
}
