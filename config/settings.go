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
	PORT   int16
	USERS  map[string]string
	AGENTS []string
}

func newSettings() *settings {
	portString := os.Getenv("PORT")
	port, err := strconv.Atoi(portString)
	if err != nil {
		log.Fatal(err)
	}
	return &settings{
		PORT: int16(port),
		USERS: map[string]string{
			// Add all users
			os.Getenv("USER1"):  os.Getenv("PASSWORD1"),
			os.Getenv("AGENT1"): string(os.Getenv("PASSWORD2")),
			os.Getenv("AGENT2"): string(os.Getenv("PASSOWRD3")),
		},
		AGENTS: []string{
			os.Getenv("AGENT1"),
		},
	}
}

func init() {
	env := os.Getenv("GO_ENV")
	if env != "prod" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("No .env file")
		}
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
