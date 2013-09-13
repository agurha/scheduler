package controllers

import (
	"log"
	"net/http"
)

func ScheduleJob(w http.ResponseWriter, r *http.Request) {

	log.Println("Schedule a Job")

	addResponseHeaders(w)

	jsonBody, err := getRequestBody(r)
	if err != nil {

		http.Error(w, "Failed to get request Body", http.StatusBadRequest)
		return
	}

	log.Println(jsonBody)

}
