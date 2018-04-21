export GOPATH=$(pwd)

# Build the front-end application

#CDA_API_HOST='controle-de-acesso.douglaszuqueto.com:8080' CDA_MQTT_HOST='controle-de-acesso.douglaszuqueto.com:8083' yarn --cwd ../front run build

# CDA_API_HOST="192.168.0.160:3000" CDA_MQTT_HOST="192.168.0.160:8083" yarn --cwd ../front run build
CDA_API_HOST="192.168.0.100:3000" CDA_MQTT_HOST="192.168.0.100:8083" yarn --cwd ../front run build

# Copy files from front/dist to public

rm -rf ./public
mkdir public
cp -r ../front/dist/* ./public

# Generate assets from front-end build

./bin/statik -f

GOARCH=arm go build -o rpi-app

scp rpi-app pi@192.168.0.100:/home/pi/

# scp rpi-app pi@192.168.0.160:/home/pi/tmp/
#ssh pi@192.168.0.150 "bash cda-pre-install.sh"
