package scheduler

import (
	"fmt"
	"sync"

	"github.com/dykoffi/forexauto/src/logger"
	"github.com/dykoffi/forexauto/src/process"
	"github.com/robfig/cron/v3"
)

type Interface interface {
	RunCrons(cronExpr string) error
}

type SchedulerService struct {
	cron    *cron.Cron
	logger  logger.LoggerInterface
	process process.ProcessInterface
}

var (
	iScheduleService SchedulerService
	once             sync.Once
)

func New(cron *cron.Cron, logger logger.LoggerInterface, process process.ProcessInterface) *SchedulerService {

	once.Do(func() {
		iScheduleService = SchedulerService{
			cron:    cron,
			logger:  logger,
			process: process,
		}
	})

	return &iScheduleService
}

func (ss *SchedulerService) RunCrons(cronExpr string) error {
	ss.logger.Info("Running crons ...")
	// fmt.Println(ss.cron)
	_, err := ss.cron.AddFunc(cronExpr, func() {
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
