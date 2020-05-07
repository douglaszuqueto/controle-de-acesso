package mqtt

import (
	"fmt"
	"io/ioutil"
	"net/http"

	models "github.com/douglaszuqueto/controle-de-acesso/app/models"
	utils "github.com/douglaszuqueto/controle-de-acesso/app/utils"
)

func sendPong(payload mqttPayload) {
	tag := models.VerifyTag(payload.Tag, payload.ChipID)

	// sendPongToDevice(payload)
	sendPongToDevice2(tag, payload)
	sendPongToFront(tag, payload)
}

func sendPongToDevice2(tag *models.Ping, payload mqttPayload) {
	topic := fmt.Sprintf("/device/pong/%s", payload.ChipID)
	msg := fmt.Sprintf("%s|%v", tag.User, tag.State)

	client.Publish(topic, 1, false, msg)
}

func sendPongToDevice(payload mqttPayload) {
	uri := fmt.Sprintf("http://127.0.0.1:3000/api/verify/%s/%s", payload.Tag, payload.ChipID)
	response, err := http.Get(uri)
	utils.CheckErr(err)

	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	utils.CheckErr(err)

	topic := fmt.Sprintf("/device/pong/%s", payload.ChipID)
	client.Publish(topic, 1, false, contents)
}

func sendPongToFront(tag *models.Ping, payload mqttPayload) {
	in := fmt.Sprintf(`{"chip_id":"%s","state":"%v","gravatar":"%s","user":"%s"}`, payload.ChipID, tag.State, tag.Gravatar, tag.User)
	client.Publish("/device/verify", 1, false, in)

	createLog(tag)
}

func createLog(p *models.Ping) {
	err := models.CreateLog(p)
	utils.CheckErr(err)
}
