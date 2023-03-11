package routes

import (
	controller "github.com/PATILSHUBHAM69/PG-Management-Go-Project/controllers"
	"github.com/PATILSHUBHAM69/PG-Management-Go-Project/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.Get_All_PG())
	incomingRoutes.GET("/users/:user_id", controller.Get_PG_ByLocation())
	incomingRoutes.GET("/getUser", controller.Book_pg())
	incomingRoutes.GET("/getUser", controller.Update_booking())
	incomingRoutes.GET("/getUser", controller.Delete_Booking())
}
