package main

import (
	"learning-gin/src/config"
	"learning-gin/src/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	log.SetLevel(logrus.InfoLevel)
}

func welcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    200,
		"status":  true,
		"message": "Hello Gin Server! 🚀",
	})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConn, err := config.SetupDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	port := ":5600"
	globalPrefix := "api/v1"

	r := gin.Default()

	api := r.Group(globalPrefix)
	{
		api.GET("/", welcome)
	}

	router.SetupTodoRoutes(r, dbConn, log)

	log.Infof("🚀 Application listening on http://0.0.0.0%s/%s", port, globalPrefix)

	r.Run(port)
}
