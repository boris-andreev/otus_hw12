package main

import (
	"time"

	"hw12/internal/repository"
	"hw12/internal/service"
)

func main() {
	items := make(chan repository.Identifier)
	todoService := service.NewTodoServise(items)
	todoRepository := repository.NewTodoRepository(items)
	go todoRepository.Listen()

	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			go todoService.BulkSave()
		}
	}
}
