package prompt

import "fmt"

type PromptOptions uint8

const (
	DEFAULT PromptOptions = iota
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

func PickMainMenuOption(p *Prompt) {
	fmt.Printf("GOT %s — ", optionTexts[p.Options])
}

func CheckValidOption() {
	// TODO: validate input prompt entries
}

func InputPrompt() {
	// TODO: scan input prompt entries
}

func DisplayOptions() {
	for _, option := range optionOrder {
		fmt.Printf("%d — %s\n", option, optionTexts[option])
	}
}

func MainMenu() {
	DisplayOptions()
}
