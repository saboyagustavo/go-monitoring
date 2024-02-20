package entity

import (
	"fmt"
	"time"

	webserver "github.com/saboyagustavo/go-monitoring/internal/webserver"
)

type Monitor struct {
	Handler  *webserver.WebMonitorHandler
	Resource *Resource
}

func NewMonitor(resource *Resource) *Monitor {
	return &Monitor{
		Handler:  webserver.NewWebMonitorHandler(),
		Resource: resource,
	}
}

func (m *Monitor) Exec() {
	switch m.Resource.HTTPMethod {
	case "GET":
		m.Handler.GetResource(m.Resource.URL, m.Resource.StatusCode)
	}
}

func (m *Monitor) Worker(ID string, data chan int) {
	for x := range data {
		fmt.Printf("Worker %s received %d\n", ID, x)
		time.Sleep(time.Second)
	}
}
