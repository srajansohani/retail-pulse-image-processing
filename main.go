package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/srajansohani/image-process-service/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/submit/", handlers.SubmitJob).Methods("POST")
	r.HandleFunc("/api/status", handlers.GetJobInfo).Methods("GET")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
