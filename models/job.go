package models

type Job struct {
	ScheduleTime  CronTime
	CallbackUrl   string
	JobDescriptor string
}
