package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DBConnectionURI = ""
	BackendPort     = 0
	LocalStackHost  = "localhost"
)

func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	BackendPort, err = strconv.Atoi(os.Getenv("BACKEND_PORT"))
	if err != nil {
		BackendPort = 9000
	}

	DBConnectionURI = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	LocalStackHost = os.Getenv("LOCALSTACK_HOST")
}
