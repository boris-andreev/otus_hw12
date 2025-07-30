package service

import (
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

	once.Do(func() {
		go t.log()
	})

	for {
		item, ok := <-t.items

		if !ok {
			break
		}

		t.repository.SaveItem(item)
	}
}

func (t *TodoService) log() {
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	var homeworkItemsAdded atomic.Int64
	var studyItemsAdded atomic.Int64
	var workoutItemsAdded atomic.Int64

	for {
		select {
		case <-ticker.C:
			func() {
				go logAddedItems(t.repository.GetHomeworskCount(), &homeworkItemsAdded, "Homeworks were added:", t.repository.GetHomewors)
				go logAddedItems(t.repository.GetStudiesCount(), &studyItemsAdded, "Studies were added:", t.repository.GetStudies)
				go logAddedItems(t.repository.GetWorkoutCount(), &workoutItemsAdded, "Workouts were added:", t.repository.GetWorkouts)
			}()
		}
	}
}

func logAddedItems[T any](itemsCount int, counter *atomic.Int64, message string, getItems func(int) []T) {
	cnt := int((*counter).Load())
	itemsWereAdded := cnt < itemsCount

	if itemsWereAdded {
		fmt.Println(message, getItems(cnt))
		(*counter).Swap(int64(itemsCount))
	}

	return
}

func NewTodoServise(repo *repository.TodoRepository) *TodoService {
	return &TodoService{
		items:      make(chan repository.Identifier),
		repository: repo,
	}
}
