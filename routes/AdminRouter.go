package routes

import (
	controller "github.com/PATILSHUBHAM69/PG-Management-Go-Project/controllers"
	"github.com/PATILSHUBHAM69/PG-Management-Go-Project/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())

}
