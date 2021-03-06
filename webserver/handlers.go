package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mhorr/mavencode-assignment/shared"
)

const oneMegabyte = 1048576

// Index is the handler for '/'
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

// PersonCreate creates a new person object
func PersonCreate(w http.ResponseWriter, r *http.Request) {
	var person shared.Person

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, oneMegabyte))
	handleWebError(err, "Failed to read message body.")

	err = r.Body.Close()
	handleWebError(err, "Failed to close response body")

	if err := json.Unmarshal(body, &person); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		err := json.NewEncoder(w).Encode(err)
		handleWebError(err, "Failed to encode error as JSON")
	}

	person.EnsureTimeStampIsSet()

	p, err := RepoSendPerson(person)
	handleWebError(err, "Failed to send person for storage.")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(p)
	handleWebError(err, "Failed to encode person as JSON")
}

// PersonList lists all person objects in database
func PersonList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personID := vars["fullname"]
	s, err := shared.GetRedisPersonStore()
	handleWebError(err, "Failed to get RedisPersonStore")
	person, err := s.QueryPersonByFullName(personID)
	handleWebError(err, "Failed to query person "+personID)
	pjs, err := json.Marshal(person)
	handleWebError(err, "Failed to marshal Persons")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(pjs)
}

func handleWebError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", err, msg)
	}
}

// PersonsQuery is the handler for /person/<range>
func PersonsQuery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queryRange := vars["range"]

	s, err := shared.GetRedisPersonStore()
	handleWebError(err, "Failed to get RedisPersonStore")
	persons, err := s.Query(queryRange)
	handleWebError(err, "Failed to query Persons")
	pjs, err := json.Marshal(persons)
	handleWebError(err, "Failed to marshal Persons")
	w.Header().Set("Content-Type", "application/json; charset=UTEF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(pjs)
}
