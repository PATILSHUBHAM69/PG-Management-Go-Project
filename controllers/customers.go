package controllers

import (
	"fmt"
	"net/http"

	"github.com/PATILSHUBHAM69/PG-Management-Go-Project/database"
	"github.com/PATILSHUBHAM69/PG-Management-Go-Project/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Get_All_PG(c *gin.Context) {
	database.Connect()
	results, err := Db.Query("SELECT * FROM Property_Details")
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
		output = fmt.Sprintf("%d %s %s %s %s  %s %s %s %s %s \n", propertyid, propertyname, contactno, propertytype, propertyaddress, city_, pincode_, landmark, ammeneties_, price_, advancedeposit)
		c.IndentedJSON(http.StatusOK, output)
	}
}

func Get_PG_ByLocation(c *gin.Context) {
	database.Connect()
	var get_pgLocation models.User
	err := c.BindJSON(&get_pgLocation)
	if err != nil {
		return
	}
	results, err := db.Query("SELECT * FROM Property_Details WHERE LandMark=%s", get_pgLocation.LandMark)
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
		output = fmt.Sprintf("%d %s %s %s %s  %s %s %s %s %s \n", propertyid, propertyname, contactno, propertytype, propertyaddress, city_, pincode_, landmark, ammeneties_, price_, advancedeposit)
		c.IndentedJSON(http.StatusOK, output)
	}
}

func Book_pg(c *gin.Context) {
	database.Connect()
	var add_booking models.Booking
	err := c.BindJSON(&add_booking)
	if err != nil {
		return
	}
	query := fmt.Sprintf(`INSERT INTO Booking_Details VALUES("%d","%s","%s" "%d","%s", "%d", "%s","%s")`, add_booking.Customer_ID, add_booking.Customer_Name, add_booking.Cus_Contact_No, add_booking.Property_ID, add_booking.Booking_time, add_booking.Booking_ID, add_booking.From_Date, add_booking.To_Date)

	insert, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Yes, PG Book Successfully!")

}

func Update_booking(c *gin.Context) {
	database.Connect()
	var edit_booking models.Booking
	err := c.BindJSON(&edit_booking)
	if err != nil {
		return
	}
	query := fmt.Sprintf("UPDATE Booking_details SET Customer_ID='%s', Customer_name='%s',Cus_Contact_No='%s',Property_ID='%s', Booking_time='%s', From_Date='%s', To_Date='%s'  WHERE Booking_ID=%s", edit_booking.Customer_ID, edit_booking.Customer_Name, edit_booking.Cus_Contact_No, edit_booking.Property_ID, edit_booking.Booking_time, edit_booking.From_Date, edit_booking.To_Date)
	results, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer results.Close()
}

func Delete_Booking(c *gin.Context) {
	database.Connect()
	var delete_booking models.Booking
	err := c.BindJSON(&delete_booking)
	if err != nil {
		return
	}
	query := fmt.Sprintf("DELETE FROM Booking_Details WHERE Booking_ID=%d", delete_booking.Booking_ID)
	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()
}
