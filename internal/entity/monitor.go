package entity

import (
	"fmt"
	"time"

	webserver "github.com/saboyagustavo/go-monitoring/internal/webserver"
)

type Monitor struct {
	Handler *webserver.WebMonitorHandler
	Path    string
	Method  string
}

func NewMonitor(path, method string) *Monitor {
	//TODO: pega aqui o path e e o método e abastece
	return &Monitor{
		Handler: webserver.NewWebMonitorHandler(),
		Path:    path,
		Method:  method,
	}
}

func (m *Monitor) Exec() {
	switch m.Method {
	case "GET":
		m.Handler.GetResource(m.Path)
	}
}

func (m *Monitor) Worker(ID string, data chan int) {
	for x := range data {
		fmt.Printf("Worker %s received %d\n", ID, x)
		time.Sleep(time.Second)
	}
}
