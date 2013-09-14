package models

import (
// "github.com/agurha/scheduler/utils"
// "time"
)

type Job struct {
	// Id            int
	ScheduleTime  string
	CallbackUrl   string
	JobDescriptor string
}

// fun NewJob(scheduleTime time.Time, callbackUrl string, jobDescriptor string) (Job, err) {

// 	job := new(Job)
// 	id, err := utils.GenUUID()

//  	job.Id = id
//  	job.ScheduleTime = scheduleTime
//  	job.CallbackUrl = callbackUrl
//  	job.JobDescriptor = jobDescriptor

//  	return job, nil
// }
