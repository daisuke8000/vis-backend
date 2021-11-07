package main

import (
	"github.com/daisuke8000/backend/src/database"
	"github.com/daisuke8000/backend/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()
	database.AutoMigrate()

	app := gin.Default()

	routes.Setup(app)

	//app.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	app.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
