package scheduler

import (
	"github.com/dykoffi/forexauto/src/logger"
	"github.com/dykoffi/forexauto/src/process"
	"github.com/robfig/cron/v3"
)

type SchedulerInterface interface {
	New() *ScheduleService
	Cron()
}

type ScheduleService struct {
	cron   *cron.Cron
	logger *logger.LoggerService
}

var IScheduleService ScheduleService

func New() *ScheduleService {
	if (IScheduleService != ScheduleService{}) {
		return &IScheduleService
	}

	IScheduleService := ScheduleService{
		cron:   cron.New(),
		logger: logger.New(),
	}

	return &IScheduleService
}

func (ss *ScheduleService) RunCrons() {
	ss.logger.Info("Running crons ...")
	ss.cron.AddFunc("* 23 * * 2-7", process.New().CollectIntraDayForex)
	ss.cron.Start()
}
