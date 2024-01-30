package main

import (
	"fmt"

	"github.com/saboyagustavo/go-monitoring/internal/prompt"
)

func main() {
	art :=
		`
	▄▄ •           • ▌ ▄ ·.        ▐ ▄ ▪ ▄▄▄▄▄      ▄▄▄  ▪   ▐ ▄  ▄▄ • 
	▐█ ▀ ▪▪         ·██ ▐███▪▪     •█▌▐███•██  ▪     ▀▄ █·██ •█▌▐█▐█ ▀ ▪
	▄█ ▀█▄ ▄█▀▄     ▐█ ▌▐▌▐█· ▄█▀▄ ▐█▐▐▌▐█·▐█.▪ ▄█▀▄ ▐▀▀▄ ▐█·▐█▐▐▌▄█ ▀█▄
	▐█▄▪▐█▐█▌.▐▌    ██ ██▌▐█▌▐█▌.▐▌██▐█▌▐█▌▐█▌·▐█▌.▐▌▐█•█▌▐█▌██▐█▌▐█▄▪▐█
	·▀▀▀▀  ▀█▄▀▪    ▀▀  █▪▀▀▀ ▀█▄▀▪▀▀ █▪▀▀▀▀▀▀  ▀█▄▀▪.▀  ▀▀▀▀▀▀ █▪·▀▀▀▀ 
	`

	fmt.Println("APP && API STATUSES AND HEALTHCHECK")
	fmt.Println(art)
	fmt.Println("Please select one of the main options below")
	prompt.MainMenu()
}
