package prompt

import (
	"fmt"
	"os"

	entity "github.com/saboyagustavo/go-monitoring/internal/entity"
	"github.com/saboyagustavo/go-monitoring/utils"
)

type PromptOptions uint8

const (
	_ PromptOptions = iota
	MONITOR
	AUDIT
	EXIT
)

/* TODO:
resource CRUD
 what URL?
 is it a valid http/https?
which method?
	GET
what status expected?
 is it a valid status?
Want to notify by email when error?
 boolean
*/

type PromptAction func(p *Prompt)

type Prompt struct {
	Options PromptOptions
}

var optionOrder = []PromptOptions{MONITOR, AUDIT, EXIT}

var optionTexts = map[PromptOptions]string{
	MONITOR: "START MONITOR",
	AUDIT:   "AUDIT LOGGING",
	EXIT:    "LEAVE PROGRAM",
}

func (p *Prompt) ExecOptionAction() {
	switch p.Options {
	case MONITOR:
		/*
			TODO:
				- Retrieve from an in-memory database (SQLite?) a Map with all urls to monitor
				- plus its corresponding method and expected payload
					-- do this in a routine for each monitored app
					-- loop on a range and apply a time sleep for doing the monitoring from times to times
				- CRUD operation for a new path to be monitored
				- Send a mail if one of the monitored services is down
				- Unit testing for the monitor
		*/
		m := entity.NewMonitor("https://httpstat.us/random/200", "GET")
		m2 := entity.NewMonitor("https://random-status-code.herokuapp.com/", "GET")
		m.Exec()
		m2.Exec()
	case EXIT:
		fmt.Println("EXITING THE PROGRAM... BYE!")
		os.Exit(0)
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
