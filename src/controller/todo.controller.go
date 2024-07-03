package controller

import (
	"learning-gin/src/model"
	"learning-gin/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TodoController struct {
	Service *service.TodoService
	Log     *logrus.Logger
}

func NewTodoController(service *service.TodoService, log *logrus.Logger) *TodoController {
	return &TodoController{
		Service: service,
		Log:     log,
	}
}

func (ctrl *TodoController) Create(c *gin.Context) {
	var todo model.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.Service.Create(&todo); err != nil {
		ctrl.Log.Errorf("Failed to create todo: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to create todo",
		}})
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "Success creating todo",
		"data":    todo,
	})
}

func (ctrl *TodoController) GetAll(c *gin.Context) {
	todos, err := ctrl.Service.GetAll()
	if err != nil {
		ctrl.Log.Errorf("Failed to fetch todos: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to fetch todos",
		}})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success fetching todos",
		"data":    todos,
	})
}