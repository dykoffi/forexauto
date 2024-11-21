package scheduler

import (
	"fmt"
	"sync"

	"github.com/dykoffi/forexauto/src/logger"
	"github.com/dykoffi/forexauto/src/process"
	"github.com/robfig/cron/v3"
)

type Interface interface {
	RunCrons()
}

type Service struct {
	cron    *cron.Cron
	logger  logger.Interface
	process process.Interface
}

var (
	iScheduleService Service
	once             sync.Once
)

func New(cron *cron.Cron, logger logger.Interface, process process.Interface) *Service {

	once.Do(func() {
		iScheduleService = Service{
			cron:    cron,
			logger:  logger,
			process: process,
		}
	})

	return &iScheduleService
}

func (ss *Service) RunCrons() error {
	ss.logger.Info("Running crons ...")
	// fmt.Println(ss.cron)
	_, err := ss.cron.AddFunc("09 18 * * 1-6", func() {
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
