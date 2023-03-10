// package routes

// import (
// 	controller "pg-management/controllers"

// 	"github.com/gin-gonic/gin"
// )

// func pg_owner_routes(incomingRoutes *gin.Engine) {
// 	incomingRoutes.GET("/get_ownerID", controller.Getowner_byID())
// 	incomingRoutes.GET("/get", controller.Get_Alluser())
// 	incomingRoutes.POST("/sign_up", controller.Sign_up())
// 	incomingRoutes.POST("/owner_login", controller.Login())
// 	incomingRoutes.PATCH("/owner_update", controller.Update_owner())
// 	incomingRoutes.DELETE("/owner_delete", controller.Delete_owner())
// 	incomingRoutes.GET("/see_booking", controller.See_booking())
// }
