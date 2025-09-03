package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"hw12/internal/app"
	"hw12/internal/repository"
	"hw12/internal/service"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup
	todoService := service.NewTodoServise(repository.NewTodoRepository(), ctx, &wg)
	app := app.New(ctx, &wg, todoService)

	app.Start()

	wg.Wait()
	fmt.Println("\nGracefull shutdown is ok")
}
