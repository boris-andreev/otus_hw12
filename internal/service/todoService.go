package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"hw12/internal/model"
	"hw12/internal/repository"
)

type TodoService struct {
	items      chan repository.Identifier
	repository *repository.TodoRepository
	ctx        context.Context
	wg         *sync.WaitGroup
}

func (t *TodoService) BulkSave() {
	t.items <- &model.HomeworkItem{Description: "Math homework"}
	t.items <- &model.StudyItem{Topic: "Math lesson"}
	t.items <- &model.WorkoutItem{Target: "Grow musculs"}
}

func (t *TodoService) Listen() {
	var once sync.Once

	t.wg.Add(1)

	go func() {
		defer t.wg.Done()

		once.Do(func() {
			t.log()
		})

		for item := range t.items {
			t.repository.SaveItem(item)
		}
	}()
}

func (t *TodoService) Produce() {
	t.wg.Add(1)

	go func() {
		defer t.wg.Done()

		ticker := time.NewTicker(50 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				t.BulkSave()
			case <-t.ctx.Done():
				close(t.items)
				return
			}
		}
	}()
}

func (t *TodoService) log() {
	t.wg.Add(1)

	go func() {
		defer t.wg.Done()

		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()

		homeworkItemsAdded := t.repository.GetHomeworskCount()
		studyItemsAdded := t.repository.GetStudiesCount()
		workoutItemsAdded := t.repository.GetWorkoutCount()

		for {
			select {
			case <-ticker.C:
				func() {
					homeworkItemsAdded = logAddedItems(homeworkItemsAdded, "Homeworks were added:", t.repository.GetHomewors)
					studyItemsAdded = logAddedItems(studyItemsAdded, "Studies were added:", t.repository.GetStudies)
					workoutItemsAdded = logAddedItems(workoutItemsAdded, "Workouts were added:", t.repository.GetWorkouts)
				}()
			case <-t.ctx.Done():
				return
			}
		}
	}()
}

func logAddedItems[T any](counter int, message string, getItems func(int) (int, []*T)) int {
	totalCount, items := getItems(counter)

	if totalCount > counter {
		fmt.Print(message, "[")

		for i, item := range items {
			fmt.Print(*item)
			if i < len(items)-1 {
				fmt.Print(", ")
			}
		}

		fmt.Print("]\n")
	}

	return totalCount
}

func NewTodoServise(repo *repository.TodoRepository, ctx context.Context, wg *sync.WaitGroup) *TodoService {
	return &TodoService{
		items:      make(chan repository.Identifier),
		repository: repo,
		ctx:        ctx,
		wg:         wg,
	}
}
