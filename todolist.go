package activetask

import (
	"log"

	"github.com/gammons/todolist/todolist"
)

func GetTaskById(taskId int) *todolist.Todo {
	store := todolist.FileStore{}
	todos, err := store.Load()
	if err != nil {
		log.Printf("Error loading ToDos", err)
		return nil
	}
	log.Printf("Checking for task ID %d in %d todos", taskId, len(todos))

	for _, todo := range todos {
		if todo.Id == taskId {
			return todo
		}
	}

	return nil

}
