package main

import (
	"time"

	"hw12/internal/service"
)

func main() {
	todoService := service.NewTodoServise()

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			func() {
				todoService.BulkSave()
				todoService.PrintItems()
			}()
		}
	}

}
