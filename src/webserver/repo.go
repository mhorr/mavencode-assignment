package main

import (
	"../shared"
)

// var currentID int
// var todos Todos

// func init() {
// 	RepoCreateTodo(Todo{Name: "Write presentation"})
// 	RepoCreateTodo(Todo{Name: "Host meetup"})
// }

// // RepoFindTodo looks for an item and returns it if found
// func RepoFindTodo(id int) Todo {
// 	for _, t := range todos {
// 		if t.ID == id {
// 			return t
// 		}
// 	}
// 	return Todo{}
// }

// // RepoCreateTodo adds a todo to the DB
// func RepoCreateTodo(t Todo) Todo {
// 	currentID++
// 	t.ID = currentID
// 	todos = append(todos, t)
// 	return t
// }

// // RepoDestroyTodo removes the todo with the passed-in id from the DB
// func RepoDestroyTodo(id int) error {
// 	for i, t := range todos {
// 		if t.ID == id {
// 			todos = append(todos[:i], todos[i+1:]...)
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
// }

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
