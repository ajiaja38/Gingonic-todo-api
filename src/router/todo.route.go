package router

import (
	"learning-gin/src/controller"
	"learning-gin/src/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func SetupTodoRoutes(r *gin.Engine, dbConn *gorm.DB, log *logrus.Logger) {
	todoService := service.NewTodoService(dbConn, log)
	todoController := controller.NewTodoController(todoService, log)

	api := r.Group("/api/v1")
	{
		api.POST("/todo", todoController.Create)
		api.GET("/todos", todoController.GetAll)
		api.GET("/todo/:id", todoController.GetById)
	}
}
