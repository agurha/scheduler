package models

import (
	"time"
)

type Job struct {
	ScheduleTime CronTime
	CallbackUrl  string
	Task         string
}

func (j Job) JobExecutionChecker(t time.Time) (ok bool) {

	ok = (j.ScheduleTime.Second == int(t.Second())) &&
		(j.ScheduleTime.Minute == int(t.Minute())) &&
		(j.ScheduleTime.Hour == int(t.Hour())) &&
		(j.ScheduleTime.DayOfMonth == int(t.Day())) &&
		(j.ScheduleTime.Month == int(t.Month())) &&
		(j.ScheduleTime.DayOfWeek == int(t.Weekday()))

	return ok
}
