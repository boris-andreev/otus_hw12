package repository

import (
	"fmt"
	"sync"
	"time"

	"hw12/internal/model"
)

type Identifier interface {
	GetId() int
}

type TodoRepository struct {
	homeworks []model.HomeworkItem
	studies   []model.StudyItem
	workouts  []model.WorkoutItem

	homeworksMutex *sync.Mutex
	studiesMutex   *sync.Mutex
	workoutsMutex  *sync.Mutex

	items chan Identifier
}

func (t *TodoRepository) Listen() {
	var once sync.Once

	once.Do(func() {
		go t.log()
	})

	for {
		item, ok := <-t.items

		if !ok {
			break
		}

		switch item.(type) {
		case model.HomeworkItem:
			go appendItem[model.HomeworkItem](&t.homeworks, item.(model.HomeworkItem), t.homeworksMutex)
		case model.StudyItem:
			go appendItem[model.StudyItem](&t.studies, item.(model.StudyItem), t.studiesMutex)
		case model.WorkoutItem:
			go appendItem[model.WorkoutItem](&t.workouts, item.(model.WorkoutItem), t.workoutsMutex)
		}
	}
}

func (t *TodoRepository) log() {
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	homeworkItemsAdded := 0
	studyItemsAdded := 0
	workoutItemsAdded := 0

	for {
		select {
		case <-ticker.C:
			func() {
				go logAdded(&t.homeworks, t.homeworksMutex, &homeworkItemsAdded, "Homeworks were added:")
				go logAdded(&t.studies, t.studiesMutex, &studyItemsAdded, "Studies were added:")
				go logAdded(&t.workouts, t.workoutsMutex, &workoutItemsAdded, "Workouts were added:")
			}()
		}
	}
}

func appendItem[T model.HomeworkItem | model.StudyItem | model.WorkoutItem](slice *[]T, item T, mutex *sync.Mutex) {
	defer mutex.Unlock()

	mutex.Lock()
	*slice = append(*slice, item)
}

func logAdded[T any](slice *[]T, mutex *sync.Mutex, counter *int, message string) {
	defer mutex.Unlock()

	mutex.Lock()
	itemsWereAdded := *counter < len(*slice)

	if itemsWereAdded {
		fmt.Println(message, (*slice)[*counter:])
		*counter = len(*slice)
	}

	return
}

func NewTodoRepository(items chan Identifier) *TodoRepository {
	return &TodoRepository{
		items:          items,
		homeworksMutex: &sync.Mutex{},
		studiesMutex:   &sync.Mutex{},
		workoutsMutex:  &sync.Mutex{},
	}
}
