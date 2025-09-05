package app

import (
	"context"
	"errors"
	"fmt"
	"hw12/internal/handler"
	"hw12/internal/handler/authmiddleware"
	"hw12/internal/service"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	ginSwaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "hw12/docs"
)

type App struct {
	router *gin.Engine
	server *http.Server
	ctx    context.Context
	wg     *sync.WaitGroup
}

func New(ctx context.Context, wg *sync.WaitGroup, todoService *service.TodoService) *App {
	router := gin.Default()
	handler := handler.New(todoService)
	configureRouting(router, handler)

	return &App{
		router: router,
		server: &http.Server{
			Addr:    ":8080",
			Handler: router,
		},
		ctx: ctx,
		wg:  wg,
	}
}

func configureRouting(router *gin.Engine, handler *handler.Handler) {

	router.GET("/swagger/*any", ginSwagger.WrapHandler(ginSwaggerFiles.Handler))
	router.POST("/login", handler.Login())

	router.Use(authmiddleware.Handle())

	api := router.Group("/api")

	homeworkGroup := api.Group("/homework")
	homeworkGroup.POST("/", handler.CreateHomeworkItem())
	homeworkGroup.DELETE("/:id", handler.DeleteHomeworkItem())
	homeworkGroup.PUT("/:id", handler.UpdateHomeworkItem())
	homeworkGroup.GET("/:id", handler.GetHomeworkItem())
	homeworkGroup.GET("/", handler.GetHomeworkItems())

	studyGroup := api.Group("/study")
	studyGroup.POST("/", handler.CreateStudyItem())
	studyGroup.DELETE("/:id", handler.DeleteStudyItem())
	studyGroup.PUT("/:id", handler.UpdateStudyItem())
	studyGroup.GET("/:id", handler.GetStudyItem())
	studyGroup.GET("/", handler.GetStudyItems())

	workoutGroup := api.Group("/workout")
	workoutGroup.POST("/", handler.CreateWorkoutItem())
	workoutGroup.DELETE("/:id", handler.DeleteWorkoutItem())
	workoutGroup.PUT("/:id", handler.UpdateWorkoutItem())
	workoutGroup.GET("/:id", handler.GetWorkoutItem())
	workoutGroup.GET("/", handler.GetWorkoutItems())
}

func (a *App) Start() {
	go func() {
		if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(fmt.Sprintf("Server failed to start: %v", err))
		}
	}()

	a.listenForFinish()
}

func (a *App) listenForFinish() {
	a.wg.Add(1)

	go func() {
		defer a.wg.Done()

		for {
			select {
			case <-a.ctx.Done():
				fmt.Println("Shutting down server...")

				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				if err := a.server.Shutdown(ctx); err != nil {
					fmt.Printf("Server forced to shutdown: %v", err)
				}

				log.Println("Server exited")
				return
			}
		}
	}()
}
