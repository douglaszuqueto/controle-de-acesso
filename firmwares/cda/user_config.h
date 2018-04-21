#ifndef _USER_CONFIG_H_
#define _USER_CONFIG_H_

#define DEBUG

/* WIFI */

#define         WIFI_SSID     "YOUR_SSID"
#define         WIFI_PASSWORD "YOUR_PASSWORD"

IPAddress       ip(192, 168, 0, 120);
IPAddress       gw(192, 168, 0, 1);
IPAddress       subnet(255, 255, 255, 0);

/* MQTT */

const char*     BROKER_MQTT = "192.168.0.150";
int             BROKER_PORT = 1883;

const char*     willTopic   = "/device/state/9640801";
int             willQoS     = 0;
bool            willRetain  = false;
const char*     willMessage = "offline";

const char*     TOPIC_PING  = "/device/ping/9640801";
const char*     TOPIC_PONG  = "/device/pong/9640801";

#endif
