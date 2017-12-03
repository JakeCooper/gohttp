package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Types

// reqBody : Request Template
type reqBody struct {
	Data string `json:"data"`
}

// resBody : Response Template
type resBody struct {
	Data string `json:"data"`
}

// End Types

var (
	done = make(chan struct{})
)

// Very minimal error logging function, but very easy to integrate any other type of logging
// by modifying the contents of this function with whatever logging platform Coinbase uses.
func log(err error) {
	fmt.Printf("Log: %+v\n", err)
}

func validateRequest(req reqBody) bool {
	if req.Data == "" {
		return false
	}
	return true
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	req := reqBody{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || !validateRequest(req) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Malformed/Invalid Request. See README for more information."))
		return
	}

	err = json.NewEncoder(w).Encode(&resBody{})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Failed to Encode Response: %s", err.Error())))
	}
}

func main() {
	// HTTP Server Setup
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
	<-done // Use channel to hang server.
}
