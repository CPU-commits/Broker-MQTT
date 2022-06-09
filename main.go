package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/CPU-commits/broker_mqtt/auth"
	"github.com/CPU-commits/broker_mqtt/config"
	mqtt "github.com/mochi-co/mqtt/server"
	"github.com/mochi-co/mqtt/server/listeners"
)

var settings = config.GetSettings()

func main() {
	// Init signal and done
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()
	log.Printf("Mochi MQTT Server initializing...")
	// Create server
	server := mqtt.NewServer(nil)
	// Create TCP Listener
	tcp := listeners.NewTCP(
		"t1",
		fmt.Sprintf(":%s", strconv.Itoa(int(settings.PORT))),
	)
	// Add Listener to server
	err := server.AddListener(tcp, &listeners.Config{
		Auth: &auth.Auth{
			// Users and password allowed
			Users: map[string]string{},
			// AllowedTopics
			AllowedTopics: map[string][]string{},
		},
	}) // Need to add TLS
	if err != nil {
		log.Fatal(err)
	}
	// Start
	go func() {
		err = server.Serve()
		if err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("Server started")
	<-done

	// Close server
	server.Close()
	log.Printf("Server finished")
}
