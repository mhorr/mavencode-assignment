package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"../shared"
	"github.com/gorilla/mux"
)

const ONE_MB = 1048576

// Index is the handler for '/'
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

// // TodoIndex is the handler for '/todo'
// func TodoIndex(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json; charset=UTEF-8")
// 	w.WriteHeader(http.StatusOK)
// 	if err := json.NewEncoder(w).Encode(todos); err != nil {
// 		panic(err)
// 	}
// }

// // TodoShow is the handler for /todo/<ID>
// func TodoShow(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	todoID := vars["todoId"]
// 	fmt.Fprintln(w, "Todo show:", todoID)
// }

// // TodoCreate creates a new todo entry
// func TodoCreate(w http.ResponseWriter, r *http.Request) {
// 	var todo Todo
// 	body, err := ioutil.ReadAll(io.LimitReader(r.Body, ONE_MB))
// 	if err != nil {
// 		panic(err)
// 	}
// 	if err := r.Body.Close(); err != nil {
// 		panic(err)
// 	}
// 	if err := json.Unmarshal(body, &todo); err != nil {
// 		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 		w.WriteHeader(422) // unprocessable entity
// 		if err := json.NewEncoder(w).Encode(err); err != nil {
// 			panic(err)
// 		}
// 	}

// 	t := RepoCreateTodo(todo)
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.WriteHeader(http.StatusCreated)
// 	if err := json.NewEncoder(w).Encode(t); err != nil {
// 		panic(err)
// 	}
// }

// PersonCreate creates a new person object
func PersonCreate(w http.ResponseWriter, r *http.Request) {
	var person shared.Person

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, ONE_MB))
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
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
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
