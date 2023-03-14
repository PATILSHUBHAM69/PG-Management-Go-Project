package routes

import (
	controller "github.com/PATILSHUBHAM69/PG-Management-Go-Project/controllers"
	"github.com/PATILSHUBHAM69/PG-Management-Go-Project/middleware"
	"github.com/gin-gonic/gin"
)

func CustomerRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/allpg", controller.Get_All_PG())
	incomingRoutes.POST("/bookpg", controller.Book_pg())
	incomingRoutes.PATCH("/updatebooking", controller.Update_booking())
	incomingRoutes.DELETE("/canclebooking", controller.Delete_Booking())
}
