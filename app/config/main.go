package config

import "os"

var (
	// DbHost dbHost
	DbHost = "127.0.0.1"

	// APIPort apiPort
	APIPort = "3000"

	// APPPort appPort
	APPPort = "8080"

	// MqttHost mqttHost
	MqttHost = "127.0.0.1"
	//MqttPort mqttPort
	MqttPort = "1883"
)

// LoadConfig loadConfig
func LoadConfig() {
	loadBDConfig()
	loadAPIConfig()
	loadAppConfig()
	loadMQTTConfig()
}

func loadBDConfig() {
	host := os.Getenv("CDA_BD_HOST")

	if len(host) > 0 {
		DbHost = host
	}
}

func loadAPIConfig() {
	port := os.Getenv("CDA_API_PORT")

	if len(port) > 0 {
		APIPort = port
	}
}

func loadAppConfig() {
	port := os.Getenv("CDA_APP_PORT")

	if len(port) > 0 {
		APPPort = port
	}
}

func loadMQTTConfig() {
	host := os.Getenv("CDA_MQTT_HOST")
	port := os.Getenv("CDA_MQTT_PORT")

	if len(host) > 0 {
		MqttHost = host
	}

	if len(port) > 0 {
		MqttPort = port
	}
}
