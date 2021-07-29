package models

const (
	CronJobScheduleEveryMinute = "0 * * * * *"
	CronJobScheduleEveryHour   = "0 0 * * * *"
	CronJobScheduleEveryDay    = "0 0 0 * * *"
)

type CronJob struct {
	Name     string
	Schedule string
	Fn       func()
}
