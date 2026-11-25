package config

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string, defaultValue string) string {
	var val = os.Getenv(key)
	if val == "" {
		return defaultValue
	}

	return val
}

func SetEnv(key, value string) error {
	return os.Setenv(key, value)
}

func LoadEnvFile() error {
	appEnv := os.Getenv("WEB_LABEL_PRINTER_APP_ENV")
	fileName := ".env"
	if appEnv != "" {
		fileName = ".env." + appEnv
	}

	err := godotenv.Load(fileName)
	if err != nil {
		//  return fmt.Errorf("load env file(%s) failed, err=%w", fileName, err) -- do not return an error here, the .env file is allowed to be missing
	}

	return nil
}
