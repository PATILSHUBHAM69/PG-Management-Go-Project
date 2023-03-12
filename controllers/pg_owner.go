package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/PATILSHUBHAM69/PG-Management-Go-Project/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Add_Property() gin.HandlerFunc {

	return func(c *gin.Context) {
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

		query := fmt.Sprintf(`INSERT INTO PropertyDetails (Property_Name,Contact_No,Property_Type,Property_Address,City,Pin_Code,Landmark,Ammeneties,Price,Advance_Deposit) VALUES("%s", "%s","%s", "%s", "%s", "%s", "%s","%s", "%s", "%s")`, *add_property.Property_Name, *add_property.Contact_No, *add_property.Property_Type, *add_property.Property_Address, *add_property.City, *add_property.Pin_Code, *add_property.Landmark, *add_property.Ammeneties, *add_property.Price, *add_property.Advance_Deposit)

		insert, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
		fmt.Println("Yes, Property added Successfully!")
	}
}

func Update_Property() gin.HandlerFunc {

	return func(c *gin.Context) {
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
		query := fmt.Sprintf("UPDATE PropertyDetails SET Property_Name='%s',Contact_No='%s',Property_Type='%s',Property_Address='%s',City='%s',Pin_Code='%s',LandMark='%s',Ammeneties='%s',Price='%s',Advance_Deposit='%s' WHERE Property_ID ='%d' ", *edit_property.Property_Name, *edit_property.Contact_No, *edit_property.Property_Type, *edit_property.Property_Address, *edit_property.City, *edit_property.Pin_Code, *edit_property.Landmark, *edit_property.Ammeneties, *edit_property.Price, *edit_property.Advance_Deposit, *edit_property.Property_ID)
		_, err = db.Exec(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	}
}

func Delete_property() gin.HandlerFunc {

	return func(c *gin.Context) {
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
		query := fmt.Sprintf("DELETE FROM PropertyDetails WHERE Property_ID=%d", *delete_property.Property_ID)
		results, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		defer results.Close()
	}
}

func See_bookings() gin.HandlerFunc {

	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		var check_bookings models.Booking_pg
		err = c.BindJSON(&check_bookings)
		if err != nil {
			return
		}
		query := fmt.Sprintf("SELECT * FROM BookingDetails WHERE Property_ID=%d", check_bookings.Property_ID)
		results, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		defer results.Close()
		var output interface{}
		for results.Next() {
			var Booking_id int
			var Customer_id int
			var Customer_name string
			var Cus_Contact_no string
			var Property_id int
			var Booking_time string
			var From_date string
			var To_date string
			err = results.Scan(&Booking_id, &Customer_id, &Customer_name, &Cus_Contact_no, &Property_id, &Booking_time, &From_date, &To_date)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf("Booking_Id=%d, Customer_ID=%d, Customer_Name='%s', Cus_Contact_No='%s', Property_ID=%d, Booking_Time='%s', From_Date='%s', To_date='%s' ", Booking_id, Customer_id, Customer_name, Cus_Contact_no, Property_id, Booking_time, From_date, To_date)
			c.JSON(http.StatusOK, gin.H{" PG booking ": output})
		}
	}
}
