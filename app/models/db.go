package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	config "../config"
	validator "gopkg.in/go-playground/validator.v9"
)

var db *sql.DB
var validate *validator.Validate

// InitConn init connection
func InitConn() {
	connStr := fmt.Sprintf("postgres://postgres:root@%s/controle-de-acesso?sslmode=disable", config.DbHost)

	var err error
	db, err = sql.Open("postgres", connStr)

	db.SetMaxIdleConns(500)
	db.SetMaxOpenConns(1000)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[BD] Iniciado", config.DbHost)

	validate = validator.New()
}

// Query query
func Query(sql string) (*sql.Rows, error) {
	rows, err := db.Query(sql)

	return rows, err
}

// CloseConnection closeConnection
func CloseConnection() {
	fmt.Println("[BD] Fechando conexão...")
	db.Close()
	fmt.Println("[BD] Conexão fechada!")
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
