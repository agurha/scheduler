package main

import (
	"github.com/agurha/scheduler/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", controllers.RootHandler)

	r.HandleFunc("/schedule", controllers.ScheduleJob).Methods("POST")

	http.Handle("/", r)

	log.Println("Scheduler Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
