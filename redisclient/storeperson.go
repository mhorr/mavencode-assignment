package main

import (
	"encoding/json"
	"fmt"
	"log"

	"../shared"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", err, msg)
	}
}

func main() {
	l, err := shared.NewRabbitListener(storePerson)
	handleError(err, "Unable to set up Rabbit listener.")
	l.Listen() // won't return.
}

func storePerson(b []byte) {
	var p shared.Person
	err := json.Unmarshal(b, &p)
	handleError(err, "Unable to unmarshal Person from bytes.")
	fmt.Printf("Unmarshaled person: %v\n", p)
	s, err := shared.GetRedisPersonStore()
	handleError(err, "Unable to get Redis Person store object")
	err = s.Store(p)
	handleError(err, "Unable to store person.")
}
