package router

import (
	"learning-gin/src/controller"
	"learning-gin/src/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func SetupTodoRoutes(api *gin.RouterGroup, dbConn *gorm.DB, log *logrus.Logger) {
	todoService := service.NewTodoService(dbConn, log)
	todoController := controller.NewTodoController(todoService, log)

	api.POST("/todo", todoController.Create)
	api.GET("/todos", todoController.GetAll)
	api.GET("/todo/:id", todoController.GetById)
	api.PUT("/todo/:id", todoController.UpdateHandler)
	api.DELETE("/todo/:id", todoController.Delete)
}
