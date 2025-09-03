package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"hw12/internal/httperrors"
	"hw12/internal/model"
)

type service interface {
	CreateItem(item model.Identifier)
	UpdateItem(item model.Identifier)
	DeleteHomeworkItem(id int) error
	DeleteStudyItem(id int) error
	DeleteWorkoutItem(id int) error
	GetHomeworkItem(id int) (*model.HomeworkItem, error)
	GetStudyItem(id int) (*model.StudyItem, error)
	GetWorkoutItem(id int) (*model.WorkoutItem, error)
	GetHomeworkItems() ([]*model.HomeworkItem, error)
	GetStudyItems() ([]*model.StudyItem, error)
	GetWorkoutItems() ([]*model.WorkoutItem, error)
}

func getItem[T model.ItemWithId](ctx *gin.Context, getter func(id int) (T, error)) {
	id, success := getIdFromRoute(ctx)

	if !success {
		ctx.JSON(http.StatusBadRequest, httperrors.ErrorMessage{
			Message: fmt.Sprintf("invalid id: %d", id),
		})

		return
	}

	res, err := getter(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httperrors.ErrorMessage{
			Message: "Something went wrong",
		})

		return
	}

	if res == nil {
		ctx.JSON(http.StatusNotFound, httperrors.ErrorMessage{
			Message: fmt.Sprintf("Item with id: %d not found", id),
		})

		return
	}

	ctx.JSON(http.StatusOK, res)
}

func getItems[T model.ItemWithId](ctx *gin.Context, getter func() ([]T, error)) {
	res, err := getter()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httperrors.ErrorMessage{
			Message: "Something went wrong",
		})

		return
	}

	ctx.JSON(http.StatusOK, res)
}

func createOrModify[T model.ItemWithId](ctx *gin.Context, item T, mutator func(item model.Identifier)) {
	err := ctx.ShouldBindJSON(item)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, httperrors.ErrorMessage{
			Message: err.Error(),
		})

		return
	}

	mutator(item)
}

func processDeleteRequest(ctx *gin.Context, deleter func(int) error) {
	id, success := getIdFromRoute(ctx)
	if !success {
		return
	}

	err := deleter(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, httperrors.ErrorMessage{
			Message: "Sommething went wrong",
		})
	}
}

func getIdFromRoute(ctx *gin.Context) (int, bool) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, httperrors.ErrorMessage{
			Message: err.Error(),
		})

		return 0, false
	}

	return id, true
}
