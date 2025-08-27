package model

type HomeworkItem struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

func (h *HomeworkItem) GetId() int {
	return h.Id
}

func (h *HomeworkItem) SetId(id int) {
	h.Id = id
}
