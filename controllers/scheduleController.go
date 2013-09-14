package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/agurha/scheduler/models"
	"log"
	"net/http"
)

/*
	Schedule Job request handler accepts a Json Object of structure

	{
  		"ScheduleTime" : "Hello",
  		"CallbackUrl" : "https://somecallback.com",
  		"JobDescriptor" : "SendPush"
	}
*/
func ScheduleJob(w http.ResponseWriter, r *http.Request) {

	log.Println("Schedule a Job")

	addResponseHeaders(w)

	decoder := json.NewDecoder(r.Body)

	var job *models.Job

	err := decoder.Decode(&job)

	if err != nil {

		http.Error(w, "Failed to get request Body", http.StatusBadRequest)
		return
	}

	log.Println(job.ScheduleTime)

	fmt.Fprintf(w, "Job Posted Successfully to %s", r.URL.Path[1:])
}
