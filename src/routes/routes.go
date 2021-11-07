package routes

import (
	"github.com/daisuke8000/backend/src/controllers"
	"github.com/gin-gonic/gin"
)

func Setup (app *gin.Engine) {
	api := app.Group("/api")

	admin := api.Group("/admin")
	{
		admin.POST("/register", controllers.Register)
	}

}