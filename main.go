package main

import (
	"github.com/dykoffi/forexauto/src/config"
	"github.com/dykoffi/forexauto/src/data"
	"github.com/dykoffi/forexauto/src/db"
	"github.com/dykoffi/forexauto/src/logger"
	"github.com/dykoffi/forexauto/src/process"
	"github.com/dykoffi/forexauto/src/scheduler"
	"github.com/robfig/cron/v3"
)

func main() {
	cronService := cron.New()
	configService := config.New()
	dbService := db.New(configService)
	dataService := data.New(configService)
	loggerService := logger.New(configService)
	processService := process.New(configService, dataService, dbService)

	if err := scheduler.New(cronService, loggerService, processService).RunCrons(); err != nil {
		loggerService.Error(err.Error())
	}

}
