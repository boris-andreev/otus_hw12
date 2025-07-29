package model

type HomeworkItem struct {
	Id          int
	Description string
}

func (h HomeworkItem) GetId() int {
	return h.Id
}
