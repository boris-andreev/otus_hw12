package model

type StudyItem struct {
	id    int
	Topic string
}

func (s StudyItem) GetId() int {
	return s.id
}
