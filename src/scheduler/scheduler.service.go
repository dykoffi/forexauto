package scheduler

import "time"

type SchedulerInterface interface {
	New() *ScheduleService
	Cron()
}

type ScheduleService struct {
	openDays  string
	openHours string
	beginTime string
}

var IScheduleService ScheduleService

func New() *ScheduleService {
	if (IScheduleService != ScheduleService{}) {
		return &IScheduleService
	}

	return &ScheduleService{}
}

func (ss *ScheduleService) Cron(task func() error, duration time.Duration) {
	for {
		task()
		time.Sleep(duration)
	}
}
