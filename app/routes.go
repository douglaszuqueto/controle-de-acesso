package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	config "github.com/douglaszuqueto/controle-de-acesso/app/config"
	handlers "github.com/douglaszuqueto/controle-de-acesso/app/handlers"
	utils "github.com/douglaszuqueto/controle-de-acesso/app/utils"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"
)

// Error error
type Error struct {
	Error   bool
	Message string
}

// Response response
type Response struct {
	Error Error
	Data  interface{}
}

func loadUserRoutes(api *mux.Router) {
	user := api.PathPrefix("/user").Subrouter()

	userHandler := handlers.UserHandler{}

	user.HandleFunc("", userHandler.Index).Methods("GET")
	user.HandleFunc("", userHandler.Create).Methods("POST")
	user.HandleFunc("/{id}", userHandler.Get).Methods("GET")
	user.HandleFunc("/{id}/tags", userHandler.GetUserTags).Methods("GET")
	user.HandleFunc("/{id}", userHandler.Update).Methods("PUT")
	user.HandleFunc("/{id}", userHandler.Remove).Methods("DELETE")

}

func loadTagRoutes(api *mux.Router) {
	tag := api.PathPrefix("/tag").Subrouter()

	tagHandler := handlers.TagHandler{}

	tag.HandleFunc("", tagHandler.Index).Methods("GET")
	tag.HandleFunc("", tagHandler.Create).Methods("POST")
	tag.HandleFunc("/{id}", tagHandler.Get).Methods("GET")
	tag.HandleFunc("/{id}", tagHandler.Update).Methods("PUT")
	tag.HandleFunc("/{id}", tagHandler.Remove).Methods("DELETE")
	tag.HandleFunc("/attach", tagHandler.AttachTagDevice).Methods("POST")
	tag.HandleFunc("/dettach", tagHandler.DettachTagDevice).Methods("POST")
}

func loadDeviceRoutes(api *mux.Router) {
	device := api.PathPrefix("/device").Subrouter()

	deviceHandler := handlers.DeviceHandler{}

	device.HandleFunc("", deviceHandler.Index).Methods("GET")
	device.HandleFunc("", deviceHandler.Create).Methods("POST")
	device.HandleFunc("/{id}", deviceHandler.Get).Methods("GET")
	device.HandleFunc("/{id}/tags", deviceHandler.Tags).Methods("GET")
	device.HandleFunc("/{id}/tags/{id_tag}", deviceHandler.TagActivateDeactivate).Methods("PUT")
	device.HandleFunc("/{id}", deviceHandler.Update).Methods("PUT")
	device.HandleFunc("/{id}", deviceHandler.Remove).Methods("DELETE")
}

func loadLogRoutes(api *mux.Router) {
	log := api.PathPrefix("/log").Subrouter()

	logHandler := handlers.LogHandler{}

	log.HandleFunc("", logHandler.Index).Methods("GET")
}

func jwtHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header["Authorization"]

		if len(authorization) == 0 {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Response{Error{true, "token not found"}, nil})
			return
		}

		_, err := jwt.Parse(authorization[0], func(token *jwt.Token) (interface{}, error) {
			return []byte("dz"), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Response{Error{true, "forbidden"}, nil})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func routes() *mux.Router {

	router := mux.NewRouter()
	authHandler := handlers.AuthHandler{}
	router.HandleFunc("/api/auth", authHandler.Auth).Methods("POST")
	router.HandleFunc("/api/verify/{tag}/{id_chip}", authHandler.Ping).Methods("GET")

	api := router.PathPrefix("/api").Subrouter()
	api.Use(jwtHandler)

	loadUserRoutes(api)
	loadTagRoutes(api)
	loadDeviceRoutes(api)
	loadLogRoutes(api)

	loadFrontEnd()

	return router
}

func loadFrontEnd() {
	withFront := os.Getenv("WITH_FRONTEND")

	if withFront == "true" {
		statikFS, err := fs.New()
		utils.CheckErr(err)

		http.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))
		go func() {
			port := fmt.Sprintf(":%s", config.APPPort)
			fmt.Println("[APP] Iniciando...", port)
			err := http.ListenAndServe(port, nil)
			utils.CheckErr(err)
		}()
	}
}
