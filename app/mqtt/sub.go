package mqtt

import (
	"encoding/json"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type mqttPayload struct {
	Tag    string `json:"tag"`
	ChipID string `json:"chip_id"`
}

type topicStateTagPayload struct {
	chipID   string
	state    int
	gravatar string
	user     string
}

func sub(topic string) {
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		payload := mqttPayload{}
		err := json.Unmarshal(msg.Payload(), &payload)
		if err != nil {
			fmt.Println(err)
			return
		}

		sendPong(payload)
	})
}
