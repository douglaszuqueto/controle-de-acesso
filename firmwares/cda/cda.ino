/************************** Inclusão das Bibliotecas **************************/

#include <ESP8266WiFi.h>
#include <PubSubClient.h>
#include <SPI.h>
#include <MFRC522.h>

/****************************** User Config ***********************************/

#include "user_config_override.h"

/**************************** DEBUG *******************************/

#ifdef DEBUG
#define DEBUG_PRINTLN(m) Serial.println(m)
#define DEBUG_PRINT(m) Serial.print(m)

#else
#define DEBUG_PRINTLN(m)
#define DEBUG_PRINT(m)

#endif

/**************************** Variaveis globais *******************************/

#define RST_PIN         5
#define SS_PIN          4

int rele = D3;
int button = D1;

volatile bool isReady = true;
unsigned long previousMillis      = 0;
int retries = 5;

/************************* Instanciação dos objetos  **************************/

WiFiClient client;
PubSubClient mqtt(client);
MFRC522 mfrc522(SS_PIN, RST_PIN);

/*********************** Implementação dos Prototypes *************************/

void initSerial() {
#ifdef DEBUG
  Serial.begin(115200);
#endif
}

void initWiFi() {
  delay(10);
  DEBUG_PRINTLN("");
  DEBUG_PRINT("[WIFI] Conectando-se em " + String(WIFI_SSID));

  WiFi.config(ip, gw, subnet);
  WiFi.begin(WIFI_SSID, WIFI_PASSWORD);

  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    DEBUG_PRINT(".");
  }

  DEBUG_PRINTLN("");
  DEBUG_PRINT(F("[WIFI] SSID: "));
  DEBUG_PRINTLN(WIFI_SSID);
  DEBUG_PRINT(F("[WIFI] IP: "));
  DEBUG_PRINTLN(WiFi.localIP());
  DEBUG_PRINT(F("[WIFI] Mac: "));
  DEBUG_PRINTLN(WiFi.macAddress());
  DEBUG_PRINTLN("");
}

void initMQTT() {
  mqtt.setServer(BROKER_MQTT, BROKER_PORT);
  mqtt.setCallback(mqtt_callback);
}

void recconectWiFi() {
  while (WiFi.status() != WL_CONNECTED) {
    delay(100);
    Serial.print(".");
  }
}

void reconnectMQTT() {
  if (mqtt.connected()) {
    return;
  }

  while (!mqtt.connected()) {
    Serial.println("[BROKER] Tentando se conectar ao Broker MQTT: " + String(BROKER_MQTT));

    if (retries == 0) {
      ESP.restart();
    }

    if (mqtt.connect("9640801", willTopic, willQoS, willRetain, willMessage)) {
      mqtt.subscribe(TOPIC_PONG);
      mqtt.publish(willTopic, "online");

      DEBUG_PRINTLN("[BROKER] Conectado");
      DEBUG_PRINTLN("[RFID] Aproxime o seu cartao do leitor...");

    } else {
      DEBUG_PRINTLN("[BROKER] Falha ao Reconectar");
      DEBUG_PRINTLN("[BROKER] Tentando se reconectar em 2 segundos");
      retries--;
      delay(2000);
    }
  }
}

void initRFID() {
  SPI.begin();
  mfrc522.PCD_Init();

  DEBUG_PRINTLN("[RFID] Iniciando...");
  DEBUG_PRINTLN();
}

/*********************** Lógica principal do projeto *************************/

void rfidProcess()
{
  if ( ! mfrc522.PICC_IsNewCardPresent()) {
    delay(500);
    return;
  }

  if ( ! mfrc522.PICC_ReadCardSerial()) {
    delay(500);
    return;
  }

//  if (!isReady) {
//    return;
//  }

  DEBUG_PRINT("[RFID] UID da tag : ");
  String conteudo = "";
  byte letra;
  for (byte i = 0; i < mfrc522.uid.size; i++)
  {
    conteudo.concat(String(mfrc522.uid.uidByte[i] < 0x10 ? "0" : ""));
    conteudo.concat(String(mfrc522.uid.uidByte[i], HEX));
  }
  conteudo.toUpperCase();
  DEBUG_PRINTLN(conteudo);

  yield();
  sendMqtt(conteudo);
}

void processar(String user, String state)
{
  int s = state.toInt();

  if (s != 0) {
    DEBUG_PRINTLN("[Controle] Bem vindo(a) " + String(user));
  }

  switch (s) {
    case 0:
      DEBUG_PRINTLN("[Controle] Acesso negado");
      DEBUG_PRINTLN("[Controle] Catraca bloqueada...");
      catraca(0);
      break;
    case 1:
      DEBUG_PRINTLN("[Controle] Acesso Liberado");
      DEBUG_PRINTLN("[Controle] Liberando catraca...");
      catraca(1);
      delay(500);
      catraca(0);
      delay(1000);
      break;
    case 2:
      DEBUG_PRINTLN("[Controle] Acesso não autorizado");
      DEBUG_PRINTLN("[Controle] Catraca bloqueada...");
      catraca(0);
      break;
    case 3:
      DEBUG_PRINTLN("[Controle] Acesso inativo");
      catraca(0);
      break;
    default:
      catraca(0);
      break;
  }
  DEBUG_PRINTLN("");
  DEBUG_PRINTLN("[RFID] Aproxime o seu cartao do leitor...");

  isReady = true;
}

void mqtt_callback(char* topic, byte* payload, unsigned int length) {

  String user;
  String state;
  for (int i = 0; i < length; i++) {
    char c = (char)payload[i];
    user += c;

    if ((char)payload[i] == '|') {
      state = (char)payload[i + 1];
      break;
    }
  }
  DEBUG_PRINTLN("[MQTT] Tópico => " + String(topic) + " | Valor => " + String(state));
  yield();
  processar(user, state);
}

void sendMqtt(String uuid) {

  String payload;

  payload += "{\"tag\":\"";
  payload += uuid;
  payload += "\",\"chip_id\":\"";
  payload += ESP.getChipId();
  payload += "\"}";
  char data[50];
  payload.toCharArray(data, 50);
  mqtt.publish(TOPIC_PING, data);

  isReady = false;
}

void sendChipID() {
  char data[10];
  String message = String(ESP.getChipId());
  message.toCharArray(data, 10);
  mqtt.publish("/help/chip_id", data);
  DEBUG_PRINTLN("[HELP] Enviando ChipID: " + message);
}

void catraca(int state) {
  digitalWrite(rele, state == 0 ? HIGH : LOW);
}

void setup() {
  initSerial();
  initWiFi();
  initMQTT();
  initRFID();

  pinMode(rele, OUTPUT);
  pinMode(button, INPUT_PULLUP);

  catraca(LOW);
}

void loop() {
  recconectWiFi();
  reconnectMQTT();

  mqtt.loop();

  rfidProcess();

  if (mqtt.connected() && digitalRead(button) == LOW) {
    sendChipID();
    delay(250);
  }

  yield();
}
