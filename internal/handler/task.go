package handler

import (
	"net/http"
	"redigo/internal/constant"
	"redigo/internal/domain/taskDomain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type taskHandler struct {
	taskService taskDomain.TaskService
}

func NewTaskHandler(taskService taskDomain.TaskService) *taskHandler {
	return &taskHandler{taskService}
}

func (h taskHandler) Get(c *gin.Context) {
	var tasks []taskDomain.ResponseTask
	err := h.taskService.Get(&tasks)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}

func (h taskHandler) Show(c *gin.Context) {
	taskID, err := uuid.Parse(c.Param("taskID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": http.StatusText(http.StatusBadRequest),
		})
		return
	}

	var task taskDomain.ResponseTask
	if err := h.taskService.Show(taskID, &task); err != nil {
		switch err.Error() {
		case constant.RecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": constant.RecordNotFound,
			})
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": http.StatusText(http.StatusInternalServerError),
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func (h taskHandler) Store(c *gin.Context) {
	var req taskDomain.RequestTask
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	if err := h.taskService.Store(req); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h taskHandler) Update(c *gin.Context) {
	taskID, err := uuid.Parse(c.Param("taskID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": http.StatusText(http.StatusBadRequest),
		})
	}

	var req taskDomain.RequestTask
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": http.StatusText(http.StatusBadRequest),
		})
		return
	}

	if err := h.taskService.Update(taskID, req); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h taskHandler) ToggleDone(c *gin.Context) {
	taskID, err := uuid.Parse(c.Param("taskID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": http.StatusText(http.StatusBadRequest),
		})
		return
	}

	if err := h.taskService.ToggleDone(taskID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
}

func (h taskHandler) Delete(c *gin.Context) {
	taskID, err := uuid.Parse(c.Param("taskID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": http.StatusText(http.StatusBadRequest),
		})
		return
	}

	if err := h.taskService.Delete(taskID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
}
