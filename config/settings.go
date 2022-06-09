package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var lock = &sync.Mutex{}
var singleSettings *settings

type settings struct {
	PORT int8
}

func newSettings() *settings {
	portString := os.Getenv("PORT")
	port, err := strconv.Atoi(portString)
	if err != nil {
		log.Fatal(err)
	}
	return &settings{
		PORT: int8(port),
	}
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No .env file")
	}
}

func GetSettings() *settings {
	if singleSettings == nil {
		lock.Lock()
		defer lock.Unlock()
		singleSettings = newSettings()
	}
	return singleSettings
}
