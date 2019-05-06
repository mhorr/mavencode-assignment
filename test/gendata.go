package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/mhorr/mavencode-assignment/shared"
)

// FakePerson creates a fake person.
func FakePerson() shared.Person {
	var p shared.Person
	p.Firstname = randomdata.FirstName(randomdata.RandomGender)
	p.Lastname = randomdata.LastName()
	p.Address = randomdata.Address()
	t := time.Now().Add(time.Duration(-rand.Intn(3600)) * time.Second)
	p.Timestamp = &t
	return p
}

// CheckErrWithPanic will long and panic if err is non-nil
func CheckErrWithPanic(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s\n", err, msg)
	}
}

// MakePeople creates the passed-in number of Person
// objects, adds them to an array, and returns the
// array serialized as a JSON byte array.
func MakePeople(numtogenerate int) []shared.Person {
	persons := []shared.Person{}
	for i := 1; i <= numtogenerate; i++ {
		persons = append(persons, FakePerson())
	}
	return persons
}

// SendPerson posts the person to our person endpoint
func SendPerson(person []byte) {
	url := "http://localhost:8080/person"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(person))
	req.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	CheckErrWithPanic(err, "Failed doing POST request.")
	defer resp.Body.Close()
	log.Printf("response Status: %s\n", resp.Status)
	log.Printf("response Headers: %#v\n", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Printf("response Body: %s", string(body))
}

func main() {
	var numtogenerate int
	if len(os.Args) > 1 {
		if parsedNum, err := strconv.Atoi(os.Args[1]); err != nil {
			log.Printf("Argument %s not converted to int. Defaulting\n", os.Args[1])
			numtogenerate = 10
		} else {
			numtogenerate = parsedNum
		}
	} else {
		numtogenerate = 10
	}
	people := MakePeople(numtogenerate)
	for _, person := range people {
		pb, _ := json.Marshal(person)
		SendPerson(pb)
	}
}
