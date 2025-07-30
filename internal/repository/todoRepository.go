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
	defer mutex.Unlock()

	mutex.Lock()
	*slice = append(*slice, item)
}

func (t *TodoRepository) GetHomeworskCount() int {
	defer t.homeworksMutex.RUnlock()

	t.homeworksMutex.RLock()

	return len(t.homeworks)
}

func (t *TodoRepository) GetStudiesCount() int {
	defer t.studiesMutex.RUnlock()

	t.studiesMutex.RLock()

	return len(t.studies)
}

func (t *TodoRepository) GetWorkoutCount() int {
	defer t.workoutsMutex.RUnlock()

	t.workoutsMutex.RLock()

	return len(t.workouts)
}

func (t *TodoRepository) GetHomewors(startIndex int) []model.HomeworkItem {
	defer t.homeworksMutex.RUnlock()

	t.homeworksMutex.RLock()

	return t.homeworks[startIndex:]
}

func (t *TodoRepository) GetStudies(startIndex int) []model.StudyItem {
	defer t.studiesMutex.RUnlock()

	t.studiesMutex.RLock()

	return t.studies[startIndex:]
}

func (t *TodoRepository) GetWorkouts(startIndex int) []model.WorkoutItem {
	defer t.workoutsMutex.RUnlock()

	t.workoutsMutex.RLock()

	return t.workouts[startIndex:]
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{
		homeworksMutex: &sync.RWMutex{},
		studiesMutex:   &sync.RWMutex{},
		workoutsMutex:  &sync.RWMutex{},
	}
}
