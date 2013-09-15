package controllers

import (
	"encoding/json"
	// "fmt"
	// "github.com/crowdmob/goamz/aws"
	// "github.com/crowdmob/goamz/sqs"
	"io/ioutil"
	"log"
	"net/http"
)

type Response map[string]interface{}

func (r Response) String() (s string) {

	b, err := json.Marshal(r)

	if err != nil {

		s = ""
		return
	}

	s = string(b)
	log.Println("String we have is :", s)
	return
}

func addResponseHeaders(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

}

func httpErrorResponse(w http.ResponseWriter, err error) {

	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func getRequestBody(r *http.Request) ([]byte, error) {

	jsonBody, err := ioutil.ReadAll(r.Body)

	return jsonBody, err

}
