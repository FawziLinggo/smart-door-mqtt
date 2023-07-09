package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
	"log"
	"mqtt-sub/app/constants"
	"mqtt-sub/app/helpers"
	"os"
)

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Connected")
}
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Printf("Connect lost: %v", err)
}

func main() {
	client := mqtt.NewClient(optsConfig())
	topic := os.Getenv("TOPIC_NAME")
	pathImage := os.Getenv("PATH_IMAGE") + "testing.jpeg"
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error connecting to broker: %v", token.Error())
	}

	base64Image, err := helpers.ImageToBase64(pathImage)
	if err != nil {
		log.Fatal(err)
	}
	publish(client, topic, base64Image)
	client.Disconnect(250)
}

func optsConfig() *mqtt.ClientOptions {
	err := godotenv.Load(constants.EnvironmentDirectory)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		broker   = os.Getenv("BROKER")
		username = os.Getenv("USERNAME_MQTT")
		password = os.Getenv("PASSWORD_MQTT")
		port     = os.Getenv("PORT")
		//clientID = os.Getenv("CLIENT_ID_MQTT")
	)

	mqtt.ERROR = log.New(os.Stdout, "[ERROR] ", 0)
	mqtt.CRITICAL = log.New(os.Stdout, "[CRIT] ", 0)
	mqtt.WARN = log.New(os.Stdout, "[WARN]  ", 0)
	//mqtt.DEBUG = log.New(os.Stdout, "[DEBUG]  ", 0)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", broker, port))
	opts.SetClientID("pub")
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	return opts
}

func publish(client mqtt.Client, topic string, payload string) {
	//token := client.Publish(topic, 2, false, payload)
	token := client.Publish(topic, 2, false, payload)
	token.Wait()
	if token.Error() != nil {
		log.Printf("Error publishing message: %v", token.Error())
	} else {
		log.Printf("Published message: %s to topic: %s", "payload", topic)
	}
}
