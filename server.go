package main

import (
	"flag"
	"fmt"
	"github.com/agurha/scheduler/controllers"
	"github.com/agurha/scheduler/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var (
	port = flag.String("port", "3000", "Port to run server on")
)

// var job []models.Job

func main() {

	//Before we set the router ... we should start a goroutine
	// which would be infinite loop and checks job every second
	go runinfinitely()

	r := mux.NewRouter()

	r.HandleFunc("/", controllers.RootHandler)

	r.HandleFunc("/schedule", controllers.ScheduleJob).Methods("POST")

	http.Handle("/", r)

	log.Printf("Scheduler Server running on port %s\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *port), r))

}

func runinfinitely() {

	var job *models.Job

	log.Println("We're going to run this in a go routine and infinitely")
	for {

		now := time.Now()

		for _, j := range job.AllJobs() {

			if j.JobExecutionChecker(now) {

				go j.RunTask(now)
			}

		}

		time.Sleep(time.Second)
	}
}
