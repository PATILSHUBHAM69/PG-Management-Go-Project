package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type info struct {
	First_Name string
	Last_Name  string
	City       string
	Contact_No string
}

func GetUser2() gin.HandlerFunc {

	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/go_db8")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		results, err := db.Query("SELECT * FROM Student_Info")
		if err != nil {
			panic(err.Error())
		}
		defer results.Close()
		var output interface{}
		for results.Next() {
			var F_Name string
			var L_Name string
			var City_Name string
			var Con_No string
			err = results.Scan(&F_Name, &L_Name, &City_Name, &Con_No)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%s %s %s %s \n", F_Name, L_Name, City_Name, Con_No)

			c.IndentedJSON(http.StatusOK, output)
		}

	}
}
