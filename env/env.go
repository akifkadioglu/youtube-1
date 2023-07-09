package env

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

const (
	PASSWORD string = "DATABASE_PASSWORD"
)

const (
	LOCAL int = iota
	PROD
	TEST
)

func InitEnv(t int) {
	var err error
	switch t {
	case LOCAL:
		err = godotenv.Load(filepath.Join("./env", ".env"))
	case PROD:
		err = godotenv.Load(filepath.Join("/etc/secrets", ".env"))
	case TEST:

	}
	if err != nil {
		log.Fatal("env file doesnt exist", err)
	}
}

func Getenv(key string) string {
	return os.Getenv(key)
}
