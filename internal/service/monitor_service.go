package service

import (
	"log"

	"github.com/saboyagustavo/go-monitoring/internal/database"
	"github.com/saboyagustavo/go-monitoring/internal/entity"
)

type MonitorService struct {
	ResourceDB *database.ResourceDB
}

func NewMonitorService(resourceDB *database.ResourceDB) *MonitorService {
	return &MonitorService{ResourceDB: resourceDB}
}

func (ms *MonitorService) MonitorResources() {
	resources, err := ms.ResourceDB.GetResources()
	if err != nil {
		log.Printf("an error occurred when trying to get the resources: %v", err.Error())
	}

	for _, r := range resources {
		m := entity.NewMonitor(r)
		m.Exec()
	}
}
