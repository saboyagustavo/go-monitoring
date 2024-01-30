package main

import (
	"fmt"

	"github.com/saboyagustavo/go-monitoring/internal/prompt"
)

func mainMenu() {
	p := prompt.Prompt{}
	p.ExecMainMenu()
}

func main() {
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
	mainMenu()
}
