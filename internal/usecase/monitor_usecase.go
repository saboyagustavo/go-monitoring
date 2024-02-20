package usecase

import (
	"fmt"
	"os"

	"github.com/saboyagustavo/go-monitoring/internal/service"
)

type MonitorUsecase struct {
	MonitorService *service.MonitorService
}

func NewMonitorUsecase(monitorService *service.MonitorService) *MonitorUsecase {
	return &MonitorUsecase{MonitorService: monitorService}
}

func (muc *MonitorUsecase) ExecAuditLogs() {

}

func (muc *MonitorUsecase) ExecExitProgram() {
	fmt.Println("EXITING THE PROGRAM... BYE!")
	os.Exit(0)
}

func (muc *MonitorUsecase) ExecMonitorResources() {
	muc.MonitorService.MonitorResources()
}
