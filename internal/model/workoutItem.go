package model

type WorkoutItem struct {
	Id     int
	Target string
}

func (w WorkoutItem) GetId() int {
	return w.Id
}
