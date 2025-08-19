package model

type HomeworkItem struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

func (h HomeworkItem) GetId() int {
	return h.Id
}
