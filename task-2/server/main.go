package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var count int

// Book Struct
type Program struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Rating   int    `json:"rating"`
	Year     int    `json:"year"`
	PlusOnly bool   `json:"plus_only"`
}

var programs []Program

// Give all the projects
func getPrograms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(programs)
}

func main() {

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Initializing with some dummy data
	programs = append(programs, Program{ID: count, Name: "Jerks", Rating: 14, Year: 2019, PlusOnly: false})
	programs = append(programs, Program{ID: count, Name: "Derks", Rating: 16, Year: 2020, PlusOnly: true})
	programs = append(programs, Program{ID: count, Name: "Kerks", Rating: 18, Year: 2012, PlusOnly: false})
	programs = append(programs, Program{ID: count, Name: "Merks", Rating: 3, Year: 1995, PlusOnly: true})

	count++

	// Handling the endpoints
	router.HandleFunc("/api/v1/programs", getPrograms).Methods("GET")

	// Running the Server
	log.Fatal(http.ListenAndServe(":8080", router))
}
