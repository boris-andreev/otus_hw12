package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"hw12/internal/repository"
	"hw12/internal/service"
)

func main() {
	var wg sync.WaitGroup
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	//defer stop()

	todoService := service.NewTodoServise(repository.NewTodoRepository(), ctx, &wg)
	todoService.Produce()
	todoService.Listen()

	wg.Wait()
	fmt.Println("\nGracefull shutdown is ok")
}
