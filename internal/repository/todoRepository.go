package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"hw12/internal/model"
)

type Identifier interface {
	GetId() int
	SetId(int)
}

const (
	homeworksJson = "./homeworks.json"
	studiesJson   = "./studies.json"
	workoutsJson  = "./workouts.json"
)

type TodoRepository struct {
	homeworks []*model.HomeworkItem
	studies   []*model.StudyItem
	workouts  []*model.WorkoutItem

	homeworksMutex *sync.RWMutex
	studiesMutex   *sync.RWMutex
	workoutsMutex  *sync.RWMutex

	items chan Identifier
}

func (t *TodoRepository) SaveItem(item Identifier) {
	setId := func(id int) { item.SetId(id) }
	switch item.(type) {
	case *model.HomeworkItem:
		appendItem[model.HomeworkItem](&t.homeworks, item.(*model.HomeworkItem), t.homeworksMutex, homeworksJson, setId)
	case *model.StudyItem:
		appendItem[model.StudyItem](&t.studies, item.(*model.StudyItem), t.studiesMutex, studiesJson, setId)
	case *model.WorkoutItem:
		appendItem[model.WorkoutItem](&t.workouts, item.(*model.WorkoutItem), t.workoutsMutex, workoutsJson, setId)
	}
}

func appendItem[T model.HomeworkItem | model.StudyItem | model.WorkoutItem](
	slice *[]*T,
	item *T,
	mutex *sync.RWMutex,
	fileName string,
	setId func(id int)) {

	mutex.Lock()
	defer mutex.Unlock()

	setId(len(*slice))

	*slice = append(*slice, item)
	err := appendToFile(fileName, item)
	if err != nil {
		panic(err)
	}
}

func (t *TodoRepository) GetHomeworksCount() int {
	t.homeworksMutex.RLock()
	defer t.homeworksMutex.RUnlock()

	return len(t.homeworks)
}

func (t *TodoRepository) GetStudiesCount() int {
	t.studiesMutex.RLock()
	defer t.studiesMutex.RUnlock()

	return len(t.studies)
}

func (t *TodoRepository) GetWorkoutCount() int {
	t.workoutsMutex.RLock()
	defer t.workoutsMutex.RUnlock()

	return len(t.workouts)
}

func (t *TodoRepository) GetHomewors(startIndex int) (int, []*model.HomeworkItem) {
	return getItems(startIndex, t.homeworksMutex, t.homeworks)
}

func (t *TodoRepository) GetStudies(startIndex int) (int, []*model.StudyItem) {
	return getItems(startIndex, t.studiesMutex, t.studies)
}

func (t *TodoRepository) GetWorkouts(startIndex int) (int, []*model.WorkoutItem) {
	return getItems(startIndex, t.workoutsMutex, t.workouts)
}

func getItems[T model.HomeworkItem | model.StudyItem | model.WorkoutItem](
	startIndex int,
	mu *sync.RWMutex,
	slice []*T) (int, []*T) {

	mu.RLock()
	defer mu.RUnlock()

	length := len(slice)
	if length < startIndex {
		return length, []*T{}
	}

	sliceCopy := make([]*T, length-startIndex)
	copy(sliceCopy, slice[startIndex:length])

	return length, sliceCopy
}

func NewTodoRepository() *TodoRepository {
	result := &TodoRepository{
		homeworksMutex: &sync.RWMutex{},
		studiesMutex:   &sync.RWMutex{},
		workoutsMutex:  &sync.RWMutex{},
	}
	var err error

	result.homeworks, err = readFromFile[model.HomeworkItem](homeworksJson)
	if err != nil {
		panic(err)
	}

	result.studies, err = readFromFile[model.StudyItem](studiesJson)
	if err != nil {
		panic(err)
	}

	result.workouts, err = readFromFile[model.WorkoutItem](workoutsJson)
	if err != nil {
		panic(err)
	}

	return result
}

func readFromFile[T model.HomeworkItem | model.StudyItem | model.WorkoutItem](fileName string) ([]*T, error) {
	var result []*T
	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return result, nil
		}

		return nil, errors.New(fmt.Sprintf("Cannot open file %s, err: %s", fileName, err.Error()))
	}
	defer file.Close()

	writer := new(strings.Builder)
	io.Copy(writer, file)
	reader := strings.NewReader(writer.String())

	decoder := json.NewDecoder(reader)

	for {
		var item T
		err := decoder.Decode(&item)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Cannot decode item from file %s, err: %s", fileName, err.Error()))
		}
		result = append(result, &item)
	}

	return result, nil
}

func appendToFile[T model.HomeworkItem | model.StudyItem | model.WorkoutItem](fileName string, dataToAppend *T) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return errors.New(fmt.Sprintf("Cannot open file %s, err: %s", fileName, err.Error()))
	}
	defer file.Close()

	buf, err := json.Marshal(dataToAppend)
	if err != nil {
		return errors.New(fmt.Sprintf("Cannot marshal data in order to write to file %s, err: %s", fileName, err.Error()))
	}

	_, err = file.Write(buf)
	if err != nil {
		return errors.New(fmt.Sprintf("Cannot write data to file %s, err: %s", fileName, err.Error()))
	}

	return nil
}
