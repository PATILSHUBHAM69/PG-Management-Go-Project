package controllers

import (
	"pg-managment/models"
	"database/sql"
	"fmt"
	"net/http"
	"pg-managment/database"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)


func Getuser_byID() gin.HandlerFunc{
	retrun func(c *gin.Context) {
		database.Connect()
		var get_pg models.Owner
		err = c.BindJSON(&get_pro)
		if err != nil {
			return
		}
		query := fmt.Sprintf("SELECT * FROM Owner_Details WHERE ID=%d", get_pg.ID)
		results, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		defer results.Close()
		var output interface{}
		for results.Next() {
			var Id_ uint
			var F_Name string
			var L_Name string
			var Email string
			var Phonenumber string
			var City string
			var PinCode string
			var State string
			var Location string
			err = results.Scan(&Id_, &F_Name, &L_Name, &Email, &Phonenumber,&City, &PinCode, &State, &Location)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("%d %s %s %s %s %s %d %s %s %s %s \n", Id_, F_Name, L_Name, Email, Phonenumber,City, PinCode, State, Location)
			c.IndentedJSON(http.StatusOK, output)	
		}
	}
}

func Get_Allowner(c *gin.Context) {
	database.Connect()
	results, err := db.Query("SELECT * FROM Owner_Details")
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()
	var output interface{}
	for results.Next() {
		var Id_ uint
		var F_Name string
		var L_Name string
		var Email string
		var Phonenumber string
		var City string
		var PinCode string
		var State string
		var Location string
		err = results.Scan(&Id_, &F_Name, &L_Name, &Email, &Phonenumber, &City, &PinCode, &State, &Location)
		if err != nil {
			panic(err.Error())
		}
		output = fmt.Sprintf("%d %s %s %s %s  %s %s %s %s \n", Id_, F_Name, L_Name, Email, Phonenumber, City, PinCode, State, Location)
		c.IndentedJSON(http.StatusOK, output)
	}
}

func Signup_owner(c *gin.Context) {
	database.Connect()
	var sign_up models.Owner
	err = c.BindJSON(&sign_up)
	if err != nil {
		return
	}

	query_data := fmt.Sprintf(`INSERT INTO Owner_Details VALUES("%d", "%s", "%s", "%s", "%s", "%s", "%s", "%s", "%s")`, sign_up.ID, sign_up.First_Name, sign_up.Last_Name, sign_up.Email, sign_up.PhoneNumber,sign_up.City, sign_up.Pin_Code, sign_up.State, sign_up.Location)

	insert, err := db.Query(query_data)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Yes, values added!")
}

func Owner_Login(c *gin.Context) {
	database.Connect()
	var login_owner models.Owner
	err = c.BindJSON(&login_owner)
	if err != nil {
		return
	}

}


func Update_user(c *gin.Context) {
	database.Connect()
	var data models.owner
	err = c.BindJSON(&data)
	if err != nil {
		return
	}
	query := fmt.Sprintf("UPDATE Owner_details SET First_Name='%s', Last_Name='%s', Email='%s', PhoneNumber='%s', PropertyID='%d', PG_Type='%s', City='%s',Pin_Code='%d', State='%s', Location='%s' WHERE ID=%d", data.First_Name, data.Last_Name, data.Email, data.PhoneNumber, data.PropertyID, data.PG_Type, data.City, data.Pin_Code, data.State, data.Location, data.ID)
	results, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer results.Close()
}

func Delete_user(c *gin.Context) {
	database.Connect()
	var get_pro models.owner
	err = c.BindJSON(&get_pro)
	if err != nil {
		return
	}
	query := fmt.Sprintf("DELETE FROM Owner_Details WHERE ID=%d", get_pro.ID)
	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()
}

func Get_allpg(c *gin.Context) {
	database.Connect()
	var get_pro models.owner
	err = c.BindJSON(&get_pro)
	if err != nil {
		return
	}
}

func Book_pg(c *gin.Context) {
	database.Connect()
	var get_pro models.owner
	err = c.BindJSON(&get_pro)
	if err != nil {
		return
	}
}
