package dotenv

import (
	"os"

	"github.com/joho/godotenv"
)

// LoadEnvironmentVariables loads the environment variables based
// on the current environment.
func LoadEnvironmentVariables(dir string) {
	appENV := os.Getenv("APP_ENV")
	if appENV == "" {
		appENV = "dev"
	}
	godotenv.Load(dir+"env/.env", dir+"env/.env."+appENV)
}
