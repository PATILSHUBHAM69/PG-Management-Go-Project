package main

import (
	database "github.com/PATILSHUBHAM69/PG-Management-Go-Project/database"

	routes "github.com/PATILSHUBHAM69/PG-Management-Go-Project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	port := "8000"
	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.AdminRoutes(router)
	routes.CustomerRoutes(router)
	routes.PgOwnerRoutes(router)

	router.Run(":" + port)

}
