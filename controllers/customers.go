package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/PATILSHUBHAM69/PG-Management-Go-Project/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Get_All_PG() gin.HandlerFunc {

	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		results, err := db.Query("SELECT * FROM PropertyDetails")
		if err != nil {
			panic(err.Error())
		}
		defer results.Close()
		var output interface{}
		for results.Next() {
			var propertyid int
			var propertyname string
			var contactno string
			var propertytype string
			var propertyaddress string
			var city_ string
			var pincode_ string
			var landmark string
			var ammeneties_ string
			var price_ string
			var advancedeposit string
			err = results.Scan(&propertyid, &propertyname, &contactno, &propertytype, &propertyaddress, &city_, &pincode_, &landmark, &ammeneties_, &price_, &advancedeposit)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf(" Property_ID=%d,  Property_Name='%s'  Contact_Name='%s'  Property_Type='%s'  Property_Address='%s'  City='%s'  Pincode='%s'  Landmark='%s'  Ammeneties='%s'  Price='%s'  Advance_Deposit='%s'", propertyid, propertyname, contactno, propertytype, propertyaddress, city_, pincode_, landmark, ammeneties_, price_, advancedeposit)

			c.JSON(http.StatusOK, gin.H{"All PG": output})
		}
	}
}

func Get_PG_ByLocation() gin.HandlerFunc {

	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		var get_pgLocation models.User
		err = c.BindJSON(&get_pgLocation)
		if err != nil {
			return
		}
		fmt.Println(*get_pgLocation.Landmark)

		results, err := db.Query("SELECT * FROM PropertyDetails WHERE Landmark = '%s' ", *get_pgLocation.Landmark)
		fmt.Println("111")
		if err != nil {
			fmt.Println("222")
			panic(err.Error())
		}
		defer results.Close()
		var output interface{}
		for results.Next() {
			var propertyid int
			var propertyname string
			var contactno string
			var propertytype string
			var propertyaddress string
			var city_ string
			var pincode_ string
			var landmark string
			var ammeneties_ string
			var price_ string
			var advancedeposit string
			err = results.Scan(&propertyid, &propertyname, &contactno, &propertytype, &propertyaddress, &city_, &pincode_, &landmark, &ammeneties_, &price_, &advancedeposit)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf(" Property_ID=%d,  Property_Name='%s'  Contact_Name='%s'  Property_Type='%s'  Property_Address='%s'  City='%s'  Pincode='%s'  Landmark='%s'  Ammeneties='%s'  Price='%s'  Advance_Deposit='%s'", propertyid, propertyname, contactno, propertytype, propertyaddress, city_, pincode_, landmark, ammeneties_, price_, advancedeposit)

			c.JSON(http.StatusOK, gin.H{"All PG By Location": output})
		}
	}
}

func Book_pg() gin.HandlerFunc {

	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		var add_booking models.Booking_pg
		err = c.BindJSON(&add_booking)
		if err != nil {
			return
		}
		query := fmt.Sprintf(`INSERT INTO BookingDetails (Customer_ID,Customer_Name,Cus_Contact_No,Property_ID,Booking_time,From_Date,To_date) VALUES(%d,"%s","%s",%d,"%s","%s","%s")`, add_booking.Customer_ID, add_booking.Customer_Name, add_booking.Cus_Contact_No, add_booking.Property_ID, add_booking.Booking_time, add_booking.From_Date, add_booking.To_Date)

		insert, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
		fmt.Println("Yes, PG Book Successfully!")

	}
}

func Update_booking() gin.HandlerFunc {

	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		var edit_booking models.Booking_pg
		err = c.BindJSON(&edit_booking)
		if err != nil {
			return
		}
		query := fmt.Sprintf("UPDATE BookingDetails SET Customer_ID=%d,Customer_Name='%s',Cus_Contact_No='%s',From_Date='%s',To_Date='%s' WHERE Booking_ID=%d", edit_booking.Customer_ID, edit_booking.Customer_Name, edit_booking.Cus_Contact_No, edit_booking.From_Date, edit_booking.To_Date, edit_booking.Booking_ID)
		_, err = db.Exec(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
}

func Delete_Booking() gin.HandlerFunc {

	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		var delete_booking models.Booking_pg
		err = c.BindJSON(&delete_booking)
		if err != nil {
			return
		}
		query := fmt.Sprintf("DELETE FROM BookingDetails WHERE Booking_ID=%d", delete_booking.Booking_ID)
		results, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		defer results.Close()
	}
}
