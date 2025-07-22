package service

import (
	"hw12/internal/model"
	"hw12/internal/repository"
)

type TodoService struct {
	todoRepository *repository.TodoRepository
}

func (t *TodoService) BulkSave() {
	t.todoRepository.SaveItem(model.HomeworkItem{Description: "math homework"})
	t.todoRepository.SaveItem(model.StudyItem{Topic: "math homework"})
	t.todoRepository.SaveItem(model.WorkoutItem{Target: "Grow musculs"})
}

func (t *TodoService) PrintItems() {
	t.todoRepository.PrintItems()
}

func NewTodoServise() *TodoService {
	return &TodoService{
		todoRepository: repository.NewTodoRepository(),
	}
}
