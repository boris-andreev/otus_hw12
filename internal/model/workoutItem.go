package model

type WorkoutItem struct {
	id     int
	Target string
}

func (w WorkoutItem) GetId() int {
	return w.id
}
