package presentation

import (
	"fmt"

	"github.com/saboyagustavo/go-monitoring/internal/usecase"
	"github.com/saboyagustavo/go-monitoring/utils"
)

type Prompt struct {
	Option         PromptOption
	MonitorUsecase *usecase.MonitorUsecase
}

func NewPrompt(monitorUsecase *usecase.MonitorUsecase) *Prompt {
	return &Prompt{
		MonitorUsecase: monitorUsecase,
	}
}

type PromptOption uint8

const (
	_ PromptOption = iota
	MONITOR
	AUDIT
	EXIT
)

var optionOrder = []PromptOption{MONITOR, AUDIT, EXIT}

var optionTexts = map[PromptOption]string{
	MONITOR: "START MONITOR",
	AUDIT:   "AUDIT LOGGING",
	EXIT:    "LEAVE PROGRAM",
}

func (p *Prompt) ExecOptionAction() {
	switch p.Option {
	case MONITOR:
		p.MonitorUsecase.ExecMonitorResources()
	case EXIT:
		p.MonitorUsecase.ExecExitProgram()
	}
}

func (p *Prompt) ExecMainMenu() {
	for {
		p.DisplayOptions()
		p.GetUserInput()
		p.ExecOptionAction()
	}
}

func (p *Prompt) IsValidOption() bool {
	for _, option := range optionOrder {
		if p.Option == option {
			return true
		}
	}
	return false
}

func (p *Prompt) GetUserInput() {
	fmt.Println("\nSELECT ONE OF THE MAIN OPTIONS ABOVE:")

	for {
		_, err := fmt.Scan(&p.Option)
		if err != nil || !p.IsValidOption() {
			fmt.Println("PLEASE ENTER A VALID NUMBER:")
		} else {
			fmt.Printf("GOT %d — %s \n", p.Option, optionTexts[p.Option])
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
