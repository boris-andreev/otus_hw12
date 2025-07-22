package repository

import (
	"fmt"

	"hw12/internal/model"
)

type Identifier interface {
	GetId() int
}

type TodoRepository struct {
	homeworks []model.HomeworkItem
	studies   []model.StudyItem
	workouts  []model.WorkoutItem
}

func (t *TodoRepository) SaveItem(item Identifier) {
	switch item.(type) {
	case model.HomeworkItem:
		t.homeworks = append(t.homeworks, item.(model.HomeworkItem))
	case model.StudyItem:
		t.studies = append(t.studies, item.(model.StudyItem))
	case model.WorkoutItem:
		t.workouts = append(t.workouts, item.(model.WorkoutItem))
	}
}

func (t *TodoRepository) PrintItems() {
	fmt.Printf("Homework items: %d; Study items: %d; Workout items: %d\n", len(t.homeworks), len(t.studies), len(t.workouts))
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{}
}
