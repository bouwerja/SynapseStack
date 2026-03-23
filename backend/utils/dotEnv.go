package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

type DatabaseDetails struct {
	ServerName string
	Database   string
	Username   string
	Password   string
	Port       string
}

func GetDatabaseDetails() (DatabaseDetails, error) {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)

	err := godotenv.Load(filepath.Join(basePath, "..", ".env"))
	if err != nil {
		errString := fmt.Sprintf("error loading .env file: %s", err)
		return DatabaseDetails{}, errors.New(errString)
	}

	serverName := os.Getenv("MYSQL_SERVER_NAME")
	database := os.Getenv("MYSQL_DATABASE_NAME")
	user := os.Getenv("MYSQL_USERNAME")
	pwd := os.Getenv("MYSQL_PASSWORD")
	port := os.Getenv("MYSQL_PORT")

	if serverName == "" || database == "" || user == "" || pwd == "" || port == "" {
		errString := fmt.Sprintf("Environemnt Variables not found in environment:\nServer: %s\nDatabase: %s\nUser: %s\nPassword: %s\nPort: %s", serverName, database, user, pwd, port)
		return DatabaseDetails{}, errors.New(errString)
	}

	return DatabaseDetails{ServerName: serverName, Database: database, Username: user, Password: pwd, Port: port}, nil
}
