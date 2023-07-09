package subsribers

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"mqtt-sub/app/helpers"
	"os"
)

func SubTopicHistory(client mqtt.Client) {
	topic := os.Getenv("TOPIC_NAME")
	token := client.Subscribe(topic, 1, func(client mqtt.Client, message mqtt.Message) {
		log.Printf("Received message from topic: %s\n", message.Topic())
		err := helpers.Base64ToImage(string(message.Payload()))
		if err != nil {
			log.Panic(err)
		}
	})
	token.Wait()
	log.Printf("Subscribed to topic: %s", topic)
}
