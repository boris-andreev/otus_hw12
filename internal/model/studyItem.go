package model

type StudyItem struct {
	Id    int    `json:"id"`
	Topic string `json:"topic"`
}

func (s *StudyItem) GetId() int {
	return s.Id
}

func (s *StudyItem) SetId(id int) {
	s.Id = id
}
