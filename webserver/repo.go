package main

import (
	"../shared"
)

// RepoSendPerson places the person on the queue for storage in Redis
func RepoSendPerson(p shared.Person) (shared.Person, error) {
	w, err := shared.GetRabbitPersonWriter()
	if err != nil {
		return p, err
	}
	defer w.Cleanup()

	err = w.Write(p)

	if err != nil {
		return p, err
	}

	return p, nil
}
