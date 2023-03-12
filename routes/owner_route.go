package routes

import (
	controller "github.com/PATILSHUBHAM69/PG-Management-Go-Project/controllers"
	"github.com/PATILSHUBHAM69/PG-Management-Go-Project/middleware"
	"github.com/gin-gonic/gin"
)

func PgOwnerRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.POST("/addproperty", controller.Add_Property())
	incomingRoutes.PATCH("/updateproperty", controller.Update_Property())
	incomingRoutes.DELETE("/deleteproperty", controller.Delete_property())
	incomingRoutes.GET("/seebooking", controller.See_bookings())
}
