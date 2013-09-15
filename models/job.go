package models

import (
	"log"
	"time"
)

type Job struct {
	ScheduleTime CronTime
	CallbackUrl  string
	Task         string
}

func (j Job) JobExecutionChecker(t time.Time) (ok bool) {

	log.Println("We already called Job execution checker")

	ok = (j.ScheduleTime.Second == int(t.Second())) &&
		(j.ScheduleTime.Minute == int(t.Minute())) &&
		(j.ScheduleTime.Hour == int(t.Hour())) &&
		(j.ScheduleTime.DayOfMonth == int(t.Day())) &&
		(j.ScheduleTime.Month == int(t.Month())) &&
		(j.ScheduleTime.DayOfWeek == int(t.Weekday()))

	return ok
}

func (j Job) RunTask(t time.Time) (ok bool) {

	// Here we should call the callback and run the Task

	ok = true
	return ok
}

// func CreateNewJob(schedule CronTime, callback string, task string) {

// 	j := Job{schedule, callback, task}
// 	return j
// }

func (j Job) AllJobs() []*Job {

	return j.JobsSnapshot()
}

func (j Job) JobsSnapshot() []*Job {

	allJobs := []*Job{}

	for _, e := range j.AllJobs() {

		allJobs = append(allJobs, &Job{

			ScheduleTime: e.ScheduleTime,
			CallbackUrl:  e.CallbackUrl,
			Task:         e.Task,
		})
	}

	return allJobs
}
