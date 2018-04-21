package mqtt

import (
	"fmt"
	"log"
	"net/url"
	"time"

	config "../config"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var client mqtt.Client
var token mqtt.Token

func connectionLostHandler(client mqtt.Client, err error) {
	fmt.Println("[MQTT] Conexão perdida")
}

func onConnectHandler(client mqtt.Client) {
	fmt.Println("[MQTT] Conectado")
	sub("/device/ping/#")
}

func connect(clientID string, uri *url.URL) mqtt.Client {
	opts := mqtt.NewClientOptions()

	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.ClientID = clientID
	opts.AutoReconnect = true
	//opts.Username = "api"
	//opts.Password = "api"
	opts.SetKeepAlive(5 * time.Second)
	opts.SetPingTimeout(5 * time.Second)
	opts.SetConnectionLostHandler(connectionLostHandler)
	opts.SetOnConnectHandler(onConnectHandler)

	client = mqtt.NewClient(opts)

	for {
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			fmt.Println("[MQTT] tentando reconectar-se...")
			time.Sleep(3 * time.Second)
			continue
		}
		fmt.Println("[MQTT] Conectado com sucesso", uri)
		break
	}

	return client
}

// MqttRun mqttRun
func MqttRun() {

	host := fmt.Sprintf("tcp://%s:%s", config.MqttHost, config.MqttPort)
	uri, err := url.Parse(host)
	if err != nil {
		log.Fatal(err)
	}

	client = connect("sub", uri)
}

// CloseConnection closeConnection
func CloseConnection() {
	fmt.Println("[MQTT] Fechando conexão...")
	client.Disconnect(0)
	fmt.Println("[MQTT] Conexão fechada!")
}
