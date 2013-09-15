package main

import (
	"flag"
	"fmt"
	"github.com/agurha/scheduler/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	port = flag.String("port", "3000", "Port to run server on")
)

func main() {

	//Before we set the router ... we should start a channel
	// which would be infinite loop

	r := mux.NewRouter()

	r.HandleFunc("/", controllers.RootHandler)

	r.HandleFunc("/schedule", controllers.ScheduleJob).Methods("POST")

	http.Handle("/", r)

	log.Printf("Scheduler Server running on port %s\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *port), r))

}
