package config

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
	"log"
	"mqtt-sub/app/constants"
	"os"
)

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Connected")
}
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Printf("Connect lost: %v", err)
}

func OptsConfig() *mqtt.ClientOptions {
	err := godotenv.Load(constants.EnvironmentDirectory)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		broker   = os.Getenv("BROKER")
		username = os.Getenv("USERNAME_MQTT")
		password = os.Getenv("PASSWORD_MQTT")
		port     = os.Getenv("PORT")
		clientID = os.Getenv("CLIENT_ID_MQTT")
	)

	mqtt.ERROR = log.New(os.Stdout, "[ERROR] ", 0)
	mqtt.CRITICAL = log.New(os.Stdout, "[CRIT] ", 0)
	//mqtt.WARN = log.New(os.Stdout, "[WARN]  ", 0)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", broker, port))
	opts.SetClientID(clientID)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	return opts
}
