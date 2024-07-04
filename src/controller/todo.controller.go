package controller

import (
	"learning-gin/src/model"
	"learning-gin/src/service"
	"net/http"
	"strconv"

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

func (ctrl *TodoController) GetById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)

	if err != nil {
		ctrl.Log.Errorf("Invalid ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid ID",
		}})
		return
	}

	todo, err := ctrl.Service.GetById(uint(id))

	if err != nil {
		ctrl.Log.Errorf("Failed to fetch todo with ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": gin.H{
				"code":    http.StatusInternalServerError,
				"message": "Failed to fetch todo",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success fetching todo",
		"data":    todo,
	})
}

func (ctrl *TodoController) UpdateHandler(c *gin.Context) {
	var todo model.Todo
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctrl.Log.Errorf("Invalid ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid ID",
		}})
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.ID = uint(id)

	if err := ctrl.Service.Update(&todo); err != nil {
		ctrl.Log.Errorf("Failed to update todo with ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": gin.H{
				"code":    http.StatusInternalServerError,
				"message": "Failed to update todo",
			},
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success updating todo",
		"data":    todo,
	})
}

func (ctrl *TodoController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctrl.Log.Errorf("Invalid ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid ID",
		}})
		return
	}

	if err := ctrl.Service.Delete(uint(id)); err != nil {
		ctrl.Log.Errorf("Failed to delete todo with ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to delete todo",
		}})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Success deleting todo",
	})
}
