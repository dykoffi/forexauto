package scheduler

import (
	"fmt"
	"sync"

	"github.com/dykoffi/forexauto/src/logger"
	"github.com/dykoffi/forexauto/src/process"
	"github.com/robfig/cron/v3"
)

type SchedulerInterface interface {
	New() *ScheduleService
	RunCrons()
}

type ScheduleService struct {
	cron    *cron.Cron
	logger  *logger.LoggerService
	process *process.ProcessService
}

var (
	iScheduleService ScheduleService
	once             sync.Once
)

func New(cron *cron.Cron, logger *logger.LoggerService, process *process.ProcessService) *ScheduleService {

	once.Do(func() {
		iScheduleService = ScheduleService{
			cron:    cron,
			logger:  logger,
			process: process,
		}
	})

	return &iScheduleService
}

func (ss *ScheduleService) RunCrons() error {
	ss.logger.Info("Running crons ...")
	// fmt.Println(ss.cron)
	_, err := ss.cron.AddFunc("04 15 * * 1-6", func() {
		fmt.Println("test")
		if err := ss.process.CollectIntraDayForex(); err != nil {
			ss.logger.Error(err.Error())
		} else {
			ss.logger.Info("Loading successfull")
		}
	})

	if err != nil {
		return err
	}

	ss.cron.Run()

	return nil
}
