package routes

import (
	controller "github.com/Viru9029/E-Commerce-Project/controllers"
	"github.com/Viru9029/E-Commerce-Project/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
