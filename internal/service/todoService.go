package service

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"hw12/internal/model"
	"hw12/internal/repository"
)

type TodoService struct {
	items      chan repository.Identifier
	identity   atomic.Int64
	repository *repository.TodoRepository
	ctx        context.Context
	wg         *sync.WaitGroup
}

func (t *TodoService) BulkSave() {
	t.identity.Add(1)
	id := int(t.identity.Load())
	t.items <- model.HomeworkItem{Id: id, Description: "Math homework"}
	t.items <- model.StudyItem{Id: id, Topic: "Math lesson"}
	t.items <- model.WorkoutItem{Id: id, Target: "Grow musculs"}
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
					homeworkItemsAdded = logAddedItems(t.repository.GetHomeworskCount(), homeworkItemsAdded, "Homeworks were added:", t.repository.GetHomewors)
					studyItemsAdded = logAddedItems(t.repository.GetStudiesCount(), studyItemsAdded, "Studies were added:", t.repository.GetStudies)
					workoutItemsAdded = logAddedItems(t.repository.GetWorkoutCount(), workoutItemsAdded, "Workouts were added:", t.repository.GetWorkouts)
				}()
			case <-t.ctx.Done():
				return
			}
		}
	}()
}

func logAddedItems[T any](itemsCount int, counter int, message string, getItems func(int) []T) int {
	itemsWereAdded := counter < itemsCount

	if itemsWereAdded {
		fmt.Println(message, getItems(counter))
	}

	return itemsCount
}

func NewTodoServise(repo *repository.TodoRepository, ctx context.Context, wg *sync.WaitGroup) *TodoService {
	return &TodoService{
		items:      make(chan repository.Identifier),
		repository: repo,
		ctx:        ctx,
		wg:         wg,
	}
}
