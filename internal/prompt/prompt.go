package prompt

import (
	"fmt"
	"os"

	"github.com/saboyagustavo/go-monitoring/utils"
)

type PromptOptions uint8

const (
	_ PromptOptions = iota
	START
	AUDIT
	EXIT
)

type PromptAction func(p *Prompt)

type Prompt struct {
	Options PromptOptions
}

var optionOrder = []PromptOptions{START, AUDIT, EXIT}

var optionTexts = map[PromptOptions]string{
	START: "START MONITOR",
	AUDIT: "AUDIT LOGGING",
	EXIT:  "LEAVE PROGRAM",
}

func (p *Prompt) ExecOptionAction() {
	switch p.Options {
	case EXIT:
		fmt.Println("EXITING THE PROGRAM... BYE!")
		os.Exit(0)
	}
}

func (p *Prompt) ExecMainMenu() {
	p.DisplayOptions()
	p.GetUserInput()
	p.ExecOptionAction()
}

func (p *Prompt) IsValidOption() bool {
	for _, option := range optionOrder {
		if p.Options == option {
			return true
		}
	}
	return false
}

func (p *Prompt) GetUserInput() {
	fmt.Println("\nSELECT ONE OF THE MAIN OPTIONS ABOVE:")

	for {
		_, err := fmt.Scan(&p.Options)
		if err != nil || !p.IsValidOption() {
			fmt.Println("PLEASE ENTER A VALID NUMBER:")
		} else {
			fmt.Printf("GOT %d — %s \n", p.Options, optionTexts[p.Options])
			break
		}
		utils.ClearInputBuffer()
	}
}

func (p *Prompt) DisplayOptions() {
	for _, option := range optionOrder {
		fmt.Printf("%d — %s\n", option, optionTexts[option])
	}
}
