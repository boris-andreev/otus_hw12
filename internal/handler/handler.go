package handler

import (
	"hw12/internal/model"
	"hw12/internal/utils/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service
}

// @Summary Login endpoint for education purpose
// @Description Urername must be "student", Password must not be empty
// @Tags auth
// @Accept json
// @Produce json
// @Param input body model.Login true "Login"
// @Success 200 {object} string "Login successful"
// @Failure 400 {object} object "Bad request"
// @Failure 401 {object} object "Unauthorized"
// @Failure 500 {object} object "Internal server error"
// @Router /login [post]
func (h *Handler) Login() func(c *gin.Context) {
	return func(c *gin.Context) {
		req := &model.Login{}
		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}

		if "student" != req.Username {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password"})
			return
		}

		token, err := jwt.GenerateToken(req.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
			"token":   token,
		})
	}
}

// @Summary Get homework item
// @Security BearerAuth
// @Tags homework
// @Produce	json
// @Param id path int true  "Item Id"
// @Success 200 {object} model.HomeworkItem
// @Failure 400 {string} string "Invalid request"
// @Failure 404 {string} string "Item not found"
// @Router /api/homework/{id} [get]
func (h Handler) GetHomeworkItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		getItem(ctx, h.service.GetHomeworkItem)
	}
}

// @Summary Get homework items
// @Security BearerAuth
// @Tags homework
// @Produce	json
// @Success 200 {object} []model.HomeworkItem
// @Failure 500 {string} string "Something went wrong"
// @Router /api/homework/ [get]
func (h Handler) GetHomeworkItems() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		getItems(ctx, h.service.GetHomeworkItems)
	}
}

// @Summary Create homework item
// @Security BearerAuth
// @Tags homework
// @Accept json
// @Produce	json
// @Param input body model.HomeworkItem true "Homework item"
// @Success 200 {object} model.HomeworkItem
// @Failure 400 {string} string "Invalid request"
// @Router /api/homework [post]
func (h *Handler) CreateHomeworkItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		createOrModify(ctx, &model.HomeworkItem{}, h.service.CreateItem)
	}
}

// @Summary Delete homework item
// @Security BearerAuth
// @Tags homework
// @Param id path int true  "Item Id"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Something went wrong"
// @Router /api/homework/{id} [delete]
func (h *Handler) DeleteHomeworkItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		processDeleteRequest(ctx, h.service.DeleteHomeworkItem)
	}
}

// @Summary Update homework item
// @Security BearerAuth
// @Tags homework
// @Param id path int true  "Item Id"
// @Param input body model.HomeworkItem true "Homework item"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Something went wrong"
// @Router /api/homework/{id} [post]
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

// @Summary Get study item
// @Security BearerAuth
// @Tags study
// @Produce	json
// @Param id path int true  "Item Id"
// @Success 200 {object} model.StudyItem
// @Failure 400 {string} string "Invalid request"
// @Failure 404 {string} string "Item not found"
// @Router /api/study/{id} [get]
func (h Handler) GetStudyItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		getItem(ctx, h.service.GetStudyItem)
	}
}

// @Summary Get study items
// @Security BearerAuth
// @Tags study
// @Produce	json
// @Success 200 {object} []model.StudyItem
// @Failure 500 {string} string "Something went wrong"
// @Router /api/study/ [get]
func (h Handler) GetStudyItems() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		getItems(ctx, h.service.GetStudyItems)
	}
}

// @Summary Create study items
// @Security BearerAuth
// @Tags study
// @Accept json
// @Produce	json
// @Param input body model.StudyItem true "Study item"
// @Success 200 {object} model.StudyItem
// @Failure 400 {string} string "Invalid request"
// @Router /api/study [post]
func (h *Handler) CreateStudyItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		createOrModify(ctx, &model.StudyItem{}, h.service.CreateItem)
	}
}

// @Summary Delete study item
// @Security BearerAuth
// @Tags study
// @Param id path int true  "Item Id"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Something went wrong"
// @Router /api/study/{id} [delete]
func (h *Handler) DeleteStudyItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		processDeleteRequest(ctx, h.service.DeleteStudyItem)
	}
}

// @Summary Update study item
// @Security BearerAuth
// @Tags study
// @Param id path int true  "Item Id"
// @Param input body model.StudyItem true "Study item"
// @Success 200 {object} model.StudyItem
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Something went wrong"
// @Router /api/study/{id} [post]
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

// @Summary Get workout item
// @Security BearerAuth
// @Tags workout
// @Produce	json
// @Param id path int true  "Item Id"
// @Success 200 {object} model.WorkoutItem
// @Failure 400 {string} string "Invalid request"
// @Failure 404 {string} string "Item not found"
// @Router /api/workout/{id} [get]
func (h Handler) GetWorkoutItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		getItem(ctx, h.service.GetWorkoutItem)
	}
}

// @Summary Get workout items
// @Security BearerAuth
// @Tags workout
// @Produce	json
// @Success 200 {object} []model.WorkoutItem
// @Failure 500 {string} string "Something went wrong"
// @Router /api/workout/ [get]
func (h Handler) GetWorkoutItems() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		getItems(ctx, h.service.GetWorkoutItems)
	}
}

// @Summary Create workout item
// @Tags workout
// @Accept json
// @Produce	json
// @Param input body model.WorkoutItem true "Workout item"
// @Success 200 {object} model.WorkoutItem
// @Failure 400 {string} string "Invalid request"
// @Router /api/workout [post]
func (h *Handler) CreateWorkoutItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		createOrModify(ctx, &model.WorkoutItem{}, h.service.CreateItem)
	}
}

// @Summary Delete workout item
// @Tags workout
// @Param id path int true  "Item Id"
// @Success 200 {object} model.WorkoutItem
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Something went wrong"
// @Router /api/workout/{id} [delete]
func (h *Handler) DeleteWorkoutItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		processDeleteRequest(ctx, h.service.DeleteWorkoutItem)
	}
}

// @Summary Update study item
// @Tags workout
// @Param id path int true  "Item Id"
// @Param input body model.WorkoutItem true "Workout item"
// @Success 200 {object} model.WorkoutItem
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Something went wrong"
// @Router /api/workout/{id} [post]
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
