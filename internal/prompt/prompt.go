package prompt

import (
	"fmt"

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
	Actions []PromptAction
}

var optionOrder = []PromptOptions{START, AUDIT, EXIT}

var optionTexts = map[PromptOptions]string{
	START: "START MONITOR",
	AUDIT: "AUDIT LOGGING",
	EXIT:  "LEAVE PROGRAM",
}

func (p *Prompt) PickMainMenuOption() {
	fmt.Printf("GOT %d — %s ", p.Options, optionTexts[p.Options])
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
			utils.ClearInputBuffer()
		} else {
			p.PickMainMenuOption()
			break
		}
	}
	utils.ClearInputBuffer()
}

func DisplayOptions() {
	for _, option := range optionOrder {
		fmt.Printf("%d — %s\n", option, optionTexts[option])
	}
}

func MainMenu() {
	DisplayOptions()
	prompt := Prompt{}
	prompt.GetUserInput()
}
