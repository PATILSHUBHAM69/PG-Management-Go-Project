package routes

import (
	controller "github.com/PATILSHUBHAM69/PG-Management-Go-Project/controllers"
	"github.com/PATILSHUBHAM69/PG-Management-Go-Project/middleware"
	"github.com/gin-gonic/gin"
)

func CustomerRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/allpg", controller.Get_All_PG())
	incomingRoutes.GET("/all_pg_bylocation", controller.Get_PG_ByLocation())
	incomingRoutes.GET("/all_pg_bytype", controller.Get_PG_ByType())
	incomingRoutes.GET("/all_pg_byprice_month", controller.Get_PG_Price_Month())
	incomingRoutes.GET("/all_pg_byprice_day", controller.Get_PG_Price_Day())
	incomingRoutes.POST("/bookpg", controller.Book_pg())
	incomingRoutes.PATCH("/updatebooking", controller.Update_booking())
	incomingRoutes.DELETE("/canclebooking", controller.Delete_Booking())
}
