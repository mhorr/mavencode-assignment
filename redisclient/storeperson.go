package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/mhorr/mavencode-assignment/shared"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", err, msg)
	}
}

func main() {
	for {
		var l shared.RabbitListener
		for initialized := false; !initialized; {
			log.Println("storeperson.go:main(): calling shared.NewRabbitListener()")
			nl, err := shared.NewRabbitListener(storePerson)
			if err != nil {
				log.Printf("%s: %s", err, "Unable to set up Rabbit listener.")
				time.Sleep(10 * time.Second)
				continue
			}
			initialized = true
			l = nl
		}
		l.Listen() // won't return.
	}
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
