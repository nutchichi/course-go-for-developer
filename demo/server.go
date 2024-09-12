package main

import (
	"fmt"
	"net/http"
)

func flightHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("flightHandler Method:", r.Method)

	if r.Method == "GET" {
		raw := `{
			"flightNumber": "ABC-123",
			"destination": "Japan"
		}`
		w.Write([]byte(raw))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/flights", flightHandler)

	fmt.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}
