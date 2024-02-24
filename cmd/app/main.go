package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/saboyagustavo/go-monitoring/internal/database"
	"github.com/saboyagustavo/go-monitoring/internal/presentation"
	"github.com/saboyagustavo/go-monitoring/internal/service"
	"github.com/saboyagustavo/go-monitoring/internal/usecase"
)

var dbSource string

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file: ", err.Error())
	}

	dbSource = os.Getenv("DB_SOURCE")
}

func init() {
	loadEnv()

	art :=
		`
	 ▄ •           • ▌ ▄ ·.        ▐ ▄ ▪ ▄▄▄▄▄      ▄▄▄  ▪   ▐ ▄  ▄▄ • 
	▐█ ▀▀█▪         ·██ ▐███▪▪     •█▌▐███•██  ▪     ▀▄ █·██ •█▌▐█▐█ ▀ ▪
	▄█ ▀█▄ ▄█▀▄     ▐█ ▌▐▌▐█· ▄█▀▄ ▐█▐▐▌▐█·▐█.▪ ▄█▀▄ ▐▀▀▄ ▐█·▐█▐▐▌▄█ ▀█▄
	▐█▄▪▐█▐█▌.▐▌    ██ ██▌▐█▌▐█▌.▐▌██▐█▌▐█▌▐█▌·▐█▌.▐▌▐█•█▌▐█▌██▐█▌▐█▄▪▐█
	·▀▀▀▀  ▀█▄▀▪    ▀▀  █▪▀▀▀ ▀█▄▀▪▀▀ █▪▀▀▀▀▀▀  ▀█▄▀▪.▀  ▀▀▀▀▀▀ █▪·▀▀▀▀ 
	`
	fmt.Println("APP && API STATUSES AND HEALTHCHECK")
	fmt.Println(art)
}

func main() {
	db, err := sql.Open("mysql", dbSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	resourceDB := database.NewResourceDB(db)
	monitorService := service.NewMonitorService(resourceDB)
	monitorUsecase := usecase.NewMonitorUsecase(monitorService)
	prompt := presentation.NewPrompt(monitorUsecase)

	prompt.ExecMainMenu()
}
