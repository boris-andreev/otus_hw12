package main

import (
	"time"

	"hw12/internal/repository"
	"hw12/internal/service"
)

func main() {
	todoService := service.NewTodoServise(repository.NewTodoRepository())
	go todoService.Listen()

	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			go todoService.BulkSave()
		}
	}
}
