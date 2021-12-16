package main

import (
	"encoding/json"
	"fmt"
	"github.com/abhishekambastha/personal-diary/pkg/domain"
	"github.com/abhishekambastha/personal-diary/pkg/repository"
	mux "github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Println("Welcome to Personal Diary Project!")

	// Gorilla Routing
	r := mux.NewRouter()
	r.HandleFunc("/v1/entry/{id}", getEntryById).Methods("GET")
	r.HandleFunc("/v1/entry/", addEntry).Methods("POST")
	r.HandleFunc("/v1/entry/", getAllEntries).Methods("GET")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
	}
}

func getAllEntries(writer http.ResponseWriter, request *http.Request) {
	err := json.NewEncoder(writer).Encode(repository.GetAllEntries())
	if err != nil {
		return
	}
}

func addEntry(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var entry domain.Entry
	err := decoder.Decode(&entry)
	if err != nil {
		panic(err)
	}

	createdEntry := repository.AddEntry(entry)

	eErr := json.NewEncoder(writer).Encode(createdEntry)
	if eErr != nil {
		return
	}
}

func getEntryById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	entry := repository.GetEntryById(vars["id"])

	err := json.NewEncoder(writer).Encode(entry)
	if err != nil {
		return
	}
}
