package model

type HomeworkItem struct {
	id          int
	Description string
}

func (h HomeworkItem) GetId() int {
	return h.id
}
