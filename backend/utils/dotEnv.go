package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

type DatabaseDetails struct {
	ServerName string
}

func GetDatabaseDetails() (DatabaseDetails, error) {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)

	err := godotenv.Load(filepath.Join(basePath, "..", ".env"))
	if err != nil {
		return DatabaseDetails{}, fmt.Errorf("error loading .env file: %w", err)
	}

	serverName := os.Getenv("MYSQL_SERVER_NAME")

	if serverName == "" {
		return DatabaseDetails{}, fmt.Errorf("MYSQL_SERVER_NAME not found in environment")
	}

	return DatabaseDetails{ServerName: serverName}, nil
}
