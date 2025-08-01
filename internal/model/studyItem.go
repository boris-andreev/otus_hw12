package model

type StudyItem struct {
	Id    int
	Topic string
}

func (s StudyItem) GetId() int {
	return s.Id
}
