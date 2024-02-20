package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/saboyagustavo/go-monitoring/internal/database"
	"github.com/saboyagustavo/go-monitoring/internal/presentation"
	"github.com/saboyagustavo/go-monitoring/internal/service"
	"github.com/saboyagustavo/go-monitoring/internal/usecase"
)

func init() {
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
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
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
