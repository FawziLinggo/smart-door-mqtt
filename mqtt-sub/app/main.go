package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
	"log"
	"mqtt-sub/app/config"
	"mqtt-sub/app/constants"
	"mqtt-sub/app/subsribers"
	"time"
)

func main() {
	// Load .env file
	err := godotenv.Load(constants.EnvironmentDirectory)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := mqtt.NewClient(config.OptsConfig())
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error connecting to broker: %v", token.Error())
	}

	go subsribers.SubTopicHistory(client)

	for {
		time.Sleep(time.Second)
	}

}
