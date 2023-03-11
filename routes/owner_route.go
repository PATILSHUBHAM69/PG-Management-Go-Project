package routes

import (
	controller "github.com/PATILSHUBHAM69/PG-Management-Go-Project/controllers"
	"github.com/PATILSHUBHAM69/PG-Management-Go-Project/middleware"
	"github.com/gin-gonic/gin"
)

func PgOwnerRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.Add_Property())
	incomingRoutes.GET("/users/:user_id", controller.Update_Property())
	incomingRoutes.GET("/getUser", controller.Delete_property())
	incomingRoutes.GET("/getUser", controller.See_bookings())
}
