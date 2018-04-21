package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/handlers"
	_ "github.com/lib/pq"

	config "./config"
	model "./models"
	mqtt "./mqtt"
	utils "./utils"
)

func main() {
	c := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	config.LoadConfig()

	model.InitConn()

	go func() {
		mqtt.MqttRun()
	}()

	go func() {
		<-c
		fmt.Println("\n[CDA] Encerrando aplicação...\n")
		mqtt.CloseConnection()
		model.CloseConnection()
		fmt.Println("\n[CDA] Aplicação encerrada!\n")
		os.Exit(0)
		done <- true
	}()

	port := fmt.Sprintf(":%s", config.APIPort)
	fmt.Println("[API] API iniciada", port)

	AllowedHeaders := []string{"X-Requested-With", "Content-Type", "Authorization"}
	AllowedMethods := []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}
	AllowedOrigins := []string{"*"}

	err := http.ListenAndServe(port, handlers.CORS(
		handlers.AllowedHeaders(AllowedHeaders),
		handlers.AllowedMethods(AllowedMethods),
		handlers.AllowedOrigins(AllowedOrigins))(routes()))

	utils.CheckErr(err)

	<-done
	fmt.Println("Aplicação encerrada!")

}
