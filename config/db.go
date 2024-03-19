package config

import (
	"fmt"
	"os"
)

func GetDBPath() string {
	// Build := os.Getenv("BUILD")
	DBDRIVER := os.Getenv("DB_DRIVER")
	PATH := fmt.Sprintf("%s://localhost:%s", DBDRIVER, "27017")
	return PATH
}
