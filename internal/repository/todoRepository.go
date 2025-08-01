package repository

import (
	"sync"

	"hw12/internal/model"
)

type Identifier interface {
	GetId() int
}

type TodoRepository struct {
	homeworks []model.HomeworkItem
	studies   []model.StudyItem
	workouts  []model.WorkoutItem

	homeworksMutex *sync.RWMutex
	studiesMutex   *sync.RWMutex
	workoutsMutex  *sync.RWMutex

	items chan Identifier
}

func (t *TodoRepository) SaveItem(item Identifier) {
	switch item.(type) {
	case model.HomeworkItem:
		appendItem[model.HomeworkItem](&t.homeworks, item.(model.HomeworkItem), t.homeworksMutex)
	case model.StudyItem:
		appendItem[model.StudyItem](&t.studies, item.(model.StudyItem), t.studiesMutex)
	case model.WorkoutItem:
		appendItem[model.WorkoutItem](&t.workouts, item.(model.WorkoutItem), t.workoutsMutex)
	}
}

func appendItem[T model.HomeworkItem | model.StudyItem | model.WorkoutItem](slice *[]T, item T, mutex *sync.RWMutex) {
	mutex.Lock()
	defer mutex.Unlock()

	*slice = append(*slice, item)
}

func (t *TodoRepository) GetHomeworskCount() int {
	t.homeworksMutex.RLock()
	defer t.homeworksMutex.RUnlock()

	return len(t.homeworks)
}

func (t *TodoRepository) GetStudiesCount() int {
	t.studiesMutex.RLock()
	defer t.studiesMutex.RUnlock()

	return len(t.studies)
}

func (t *TodoRepository) GetWorkoutCount() int {
	t.workoutsMutex.RLock()
	defer t.workoutsMutex.RUnlock()

	return len(t.workouts)
}

func (t *TodoRepository) GetHomewors(startIndex int) []model.HomeworkItem {
	t.homeworksMutex.RLock()
	defer t.homeworksMutex.RUnlock()

	return t.homeworks[startIndex:]
}

func (t *TodoRepository) GetStudies(startIndex int) []model.StudyItem {
	t.studiesMutex.RLock()
	defer t.studiesMutex.RUnlock()

	return t.studies[startIndex:]
}

func (t *TodoRepository) GetWorkouts(startIndex int) []model.WorkoutItem {
	t.workoutsMutex.RLock()
	defer t.workoutsMutex.RUnlock()

	return t.workouts[startIndex:]
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{
		homeworksMutex: &sync.RWMutex{},
		studiesMutex:   &sync.RWMutex{},
		workoutsMutex:  &sync.RWMutex{},
	}
}
