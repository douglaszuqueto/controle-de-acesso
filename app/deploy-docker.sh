# Build the front-end application

#CDA_API_HOST=127.0.0.1 CDA_MQTT_HOST=127.0.0.1 yarn --cwd ../front run build

# Copy files from front/dist to public

#rm -rf ./public
#mkdir public
#cp -r ../front/dist/* ./public/

# Generate assets from front-end build

#./bin/statik -f

docker build -t cda:latest .

#docker run -it --rm --name cda -p 3000:3000 -p 8080:8080 --link mosquitto:mosquitto --link postgres:postgres -e CDA_BD_HOST=postgres -e CDA_MQTT_HOST=mosquitto cda

docker run -it --rm --name cda -p 3000:3000 -p 8080:8080 --link mosquitto:mosquitto --link postgres:postgres -e CDA_BD_HOST=postgres -e CDA_MQTT_HOST=mosquitto -e WITH_FRONTEND=true -m 1GB --cpus 1 cda
