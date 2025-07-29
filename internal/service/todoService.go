package service

import (
	"hw12/internal/model"
	"hw12/internal/repository"
	"sync/atomic"
)

type TodoService struct {
	items    chan repository.Identifier
	identity atomic.Int64
}

func (t *TodoService) BulkSave() {
	t.identity.Add(1)
	id := int(t.identity.Load())
	t.items <- model.HomeworkItem{Id: id, Description: "Math homework"}
	t.items <- model.StudyItem{Id: id, Topic: "Math lesson"}
	t.items <- model.WorkoutItem{Id: id, Target: "Grow musculs"}
}

func NewTodoServise(items chan repository.Identifier) *TodoService {
	return &TodoService{
		items: items,
	}
}
