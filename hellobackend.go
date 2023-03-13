package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HelloRequest struct {
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/api/hello", handleHello)
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req HelloRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello, %s!", req.Name)
}
