package handler

import (
	"hw12/internal/model"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service
}

// @Summary Get homework item
// @Tags homework
// @Produce	json
// @Param id path int true  "Item Id"
// @Success 200 {object} model.HomeworkItem
// @Failure 400 {string} string "Invalid request"
// @Router /homework/:id [post]
func (h Handler) GetHomeworkItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		getItem(ctx, h.service.GetHomeworkItem)
	}
}

// @Summary Get homework item
// @Tags homework
// @Produce	json
// @Success 200 {object} []model.HomeworkItem
// @Failure 500 {string} string "Something went wrong"
// @Router /homework/:id [post]
func (h Handler) GetHomeworkItems() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		getItems(ctx, h.service.GetHomeworkItems)
	}
}

func (h *Handler) CreateHomeworkItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		createOrModify(ctx, &model.HomeworkItem{}, h.service.CreateItem)
	}
}

func (h *Handler) DeleteHomeworkItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		processDeleteRequest(ctx, h.service.DeleteHomeworkItem)
	}
}

func (h *Handler) UpdateHomeworkItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, success := getIdFromRoute(ctx)
		if !success {
			return
		}

		item := &model.HomeworkItem{}
		createOrModify(ctx, item, func(item model.Identifier) {
			item.SetId(id)
			h.service.UpdateItem(item)
		})
	}
}

func (h Handler) GetStudyItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		getItem(ctx, h.service.GetStudyItem)
	}
}

func (h Handler) GetStudyItems() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		getItems(ctx, h.service.GetStudyItems)
	}
}

func (h *Handler) CreateStudyItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		createOrModify(ctx, &model.StudyItem{}, h.service.CreateItem)
	}
}

func (h *Handler) DeleteStudyItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		processDeleteRequest(ctx, h.service.DeleteStudyItem)
	}
}

func (h *Handler) UpdateStudyItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, success := getIdFromRoute(ctx)
		if !success {
			return
		}

		item := &model.StudyItem{}

		createOrModify(ctx, item, func(item model.Identifier) {
			item.SetId(id)
			h.service.UpdateItem(item)
		})
	}
}

func (h Handler) GetWorkoutItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		getItem(ctx, h.service.GetWorkoutItem)
	}
}

func (h Handler) GetWorkoutItems() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		getItems(ctx, h.service.GetWorkoutItems)
	}
}

func (h *Handler) CreateWorkoutItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		createOrModify(ctx, &model.WorkoutItem{}, h.service.CreateItem)
	}
}

func (h *Handler) DeleteWorkoutItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		processDeleteRequest(ctx, h.service.DeleteWorkoutItem)
	}
}

func (h *Handler) UpdateWorkoutItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, success := getIdFromRoute(ctx)

		if !success {
			return
		}

		item := &model.WorkoutItem{}

		createOrModify(ctx, item, func(item model.Identifier) {
			item.SetId(id)
			h.service.UpdateItem(item)
		})
	}
}

func New(service service) *Handler {
	return &Handler{
		service: service,
	}
}
