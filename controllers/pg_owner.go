package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/PATILSHUBHAM69/PG-Management-Go-Project/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Add_Property(c *gin.Context) {
	db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var add_property models.User
	err = c.BindJSON(&add_property)
	if err != nil {
		return
	}

	query_edit_property := fmt.Sprintf(`INSERT INTO Property_Details VALUES("%d", "%s","%s", "%s", "%s", "%s", "%s","%s", "%s", "%s", "%s")`, add_property.Property_ID, add_property.Property_Name, &add_property.Contact_No, add_property.Property_Type, add_property.Property_Address, add_property.City, add_property.Pin_Code, &add_property.LandMark, add_property.Ammeneties, add_property.Price, add_property.Advance_Deposit)

	insert, err := db.Query(query_edit_property)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Yes, Property added Successfully!")
}

func Update_Property(c *gin.Context) {
	db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var edit_property models.User
	err = c.BindJSON(&edit_property)
	if err != nil {
		return
	}
	query := fmt.Sprintf("UPDATE Property_Details SET Property_Name='%s',Contact_No='%s', Property_Type='%s', Property_Address='%s', City='%s', Pin_Code='%s', LandMark='%s' Ammeneties='%s', Price='%s', Advance_Deposit='%s' WHERE Property_ID=%d", edit_property.Property_Name, edit_property.Contact_No, edit_property.Property_Type, edit_property.Property_Address, edit_property.City, edit_property.Pin_Code, &edit_property.LandMark, edit_property.Ammeneties, edit_property.Price, edit_property.Advance_Deposit, edit_property.Property_ID)
	results, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer results.Close()
}

func Delete_property(c *gin.Context) {
	db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var delete_property models.User
	err = c.BindJSON(&delete_property)
	if err != nil {
		return
	}
	query := fmt.Sprintf("DELETE FROM Property_Details WHERE Property_ID=%d", delete_property.ID)
	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()
}

func See_bookings(c *gin.Context) {
	db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var check_bookings models.Booking
	err = c.BindJSON(&check_bookings)
	if err != nil {
		return
	}
	query := fmt.Sprintf("SELECT * FROM Booking_Details WHERE Property_ID=%d", check_bookings.Property_ID)
	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()
	var output interface{}
	for results.Next() {
		var Customer_id int
		var Property_id int
		var Booking_time string
		var Booking_id int
		err = results.Scan(&Customer_id, &Property_id, &Booking_time, &Booking_id)
		if err != nil {
			panic(err.Error())
		}
		output = fmt.Sprintf("%d %d %s %d \n", Customer_id, Property_id, Booking_time, Booking_id)
		c.IndentedJSON(http.StatusOK, output)
	}
}
