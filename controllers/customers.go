package controllers

import (
	"database/sql"
	"fmt"
	"log"
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
			var price_month int
			var advance_deposit_month int
			var price_day int
			var advance_deposit_day int
			err = results.Scan(&propertyid, &propertyname, &contactno, &propertytype, &propertyaddress, &city_, &pincode_, &landmark, &ammeneties_, &price_month, &advance_deposit_month, &price_day, &advance_deposit_day)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf(" Property_ID=%d,  Property_Name='%s'  Contact_Name='%s'  Property_Type='%s'  Property_Address='%s'  City='%s'  Pincode='%s'  Landmark='%s'  Ammeneties='%s'  Price_Month=%d Advance_Deposit_Month=%d Price_Day=%d   Advance_Deposit_Day=%d", propertyid, propertyname, contactno, propertytype, propertyaddress, city_, pincode_, landmark, ammeneties_, price_month, advance_deposit_month, price_day, advance_deposit_day)

			c.IndentedJSON(200, "PG")
			c.JSON(http.StatusOK, gin.H{"": output})
		}
	}
}

func Get_PG_ByLocation() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
		if err != nil {
			log.Fatal(err)
		}
		var pg_bylocation models.User
		err = c.BindJSON(&pg_bylocation)
		if err != nil {
			return
		}
		query_data := fmt.Sprintf("SELECT * FROM PropertyDetails WHERE Landmark='%s' AND City='%s'", *pg_bylocation.City, *pg_bylocation.Landmark)
		fmt.Println(query_data)
		results, err := db.Query(query_data)
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
			var city string
			var pincode string
			var landmark string
			var ammeneties string
			var price_month int
			var advance_deposit_month int
			var price_day int
			var advance_deposit_day int
			err = results.Scan(&propertyid, &propertyname, &contactno, &propertytype, &propertyaddress, &city, &pincode, &landmark, &ammeneties, &price_month, &advance_deposit_month, &price_day, &advance_deposit_day)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf(" Property_ID=%d,  Property_Name='%s'  Contact_Name='%s'  Property_Type='%s'  Property_Address='%s'  City='%s'  Pincode='%s'  Landmark='%s'  Ammeneties='%s'  Price_Month=%d Advance_Deposit_Month=%d Price_Day=%d   Advance_Deposit_Day=%d", propertyid, propertyname, contactno, propertytype, propertyaddress, city, pincode, landmark, ammeneties, price_month, advance_deposit_month, price_day, advance_deposit_day)

			c.IndentedJSON(200, "PG")
			c.JSON(http.StatusOK, gin.H{"": output})
		}
	}
}

func Get_PG_ByType() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
		if err != nil {
			log.Fatal(err)
		}
		var pg_bytype models.User
		err = c.BindJSON(&pg_bytype)
		if err != nil {
			return
		}
		query_data := fmt.Sprintf("SELECT * FROM PropertyDetails WHERE Property_Type='%s'", *pg_bytype.Property_Type)
		fmt.Println(query_data)
		results, err := db.Query(query_data)
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
			var city string
			var pincode string
			var landmark string
			var ammeneties string
			var price_month int
			var advance_deposit_month int
			var price_day int
			var advance_deposit_day int
			err = results.Scan(&propertyid, &propertyname, &contactno, &propertytype, &propertyaddress, &city, &pincode, &landmark, &ammeneties, &price_month, &advance_deposit_month, &price_day, &advance_deposit_day)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf(" Property_ID=%d,  Property_Name='%s'  Contact_Name='%s'  Property_Type='%s'  Property_Address='%s'  City='%s'  Pincode='%s'  Landmark='%s'  Ammeneties='%s'  Price_Month=%d Advance_Deposit_Month=%d Price_Day=%d   Advance_Deposit_Day=%d", propertyid, propertyname, contactno, propertytype, propertyaddress, city, pincode, landmark, ammeneties, price_month, advance_deposit_month, price_day, advance_deposit_day)

			c.IndentedJSON(200, "PG")
			c.JSON(http.StatusOK, gin.H{"": output})
		}
	}
}

func Get_PG_Price_Month() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
		if err != nil {
			log.Fatal(err)
		}
		var pg_byprice_month models.User
		err = c.BindJSON(&pg_byprice_month)
		if err != nil {
			return
		}
		query_data := fmt.Sprintf("SELECT * FROM PropertyDetails WHERE Price_Month <'%d'", *pg_byprice_month.Price_Month)
		fmt.Println(query_data)
		results, err := db.Query(query_data)
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
			var city string
			var pincode string
			var landmark string
			var ammeneties string
			var price_month int
			var advance_deposit_month int
			var price_day int
			var advance_deposit_day int
			err = results.Scan(&propertyid, &propertyname, &contactno, &propertytype, &propertyaddress, &city, &pincode, &landmark, &ammeneties, &price_month, &advance_deposit_month, &price_day, &advance_deposit_day)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf(" Property_ID=%d,  Property_Name='%s'  Contact_Name='%s'  Property_Type='%s'  Property_Address='%s'  City='%s'  Pincode='%s'  Landmark='%s'  Ammeneties='%s'  Price_Month=%d Advance_Deposit_Month=%d Price_Day=%d   Advance_Deposit_Day=%d", propertyid, propertyname, contactno, propertytype, propertyaddress, city, pincode, landmark, ammeneties, price_month, advance_deposit_month, price_day, advance_deposit_day)

			c.IndentedJSON(200, "PG below given amount")
			c.JSON(http.StatusOK, gin.H{"": output})
		}
	}
}

func Get_PG_Price_Day() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
		if err != nil {
			log.Fatal(err)
		}
		var pg_byprice_day models.User
		err = c.BindJSON(&pg_byprice_day)
		if err != nil {
			return
		}
		query_data := fmt.Sprintf("SELECT * FROM PropertyDetails WHERE Price_Day <'%d'", *pg_byprice_day.Price_Day)
		fmt.Println(query_data)
		results, err := db.Query(query_data)
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
			var city string
			var pincode string
			var landmark string
			var ammeneties string
			var price_month int
			var advance_deposit_month int
			var price_day int
			var advance_deposit_day int
			err = results.Scan(&propertyid, &propertyname, &contactno, &propertytype, &propertyaddress, &city, &pincode, &landmark, &ammeneties, &price_month, &advance_deposit_month, &price_day, &advance_deposit_day)
			if err != nil {
				panic(err.Error())
			}
			output = fmt.Sprintf(" Property_ID=%d,  Property_Name='%s'  Contact_Name='%s'  Property_Type='%s'  Property_Address='%s'  City='%s'  Pincode='%s'  Landmark='%s'  Ammeneties='%s'  Price_Month=%d Advance_Deposit_Month=%d Price_Day=%d   Advance_Deposit_Day=%d", propertyid, propertyname, contactno, propertytype, propertyaddress, city, pincode, landmark, ammeneties, price_month, advance_deposit_month, price_day, advance_deposit_day)

			c.IndentedJSON(200, "PG below given amount")
			c.JSON(http.StatusOK, gin.H{"": output})
		}
	}
}

func Book_PG() gin.HandlerFunc {

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
		query := fmt.Sprintf(`INSERT INTO BookingDetails (Customer_Name,Cus_Contact_No,Property_ID,From_Date,To_date) VALUES("%s","%s",%d,"%s","%s")`, add_booking.Customer_Name, add_booking.Cus_Contact_No, add_booking.Property_ID, add_booking.From_Date, add_booking.To_Date)

		insert, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
		c.IndentedJSON(200, "Yes, PG Book Successfully!")

	}
}

func See_Booking_Cus() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(127.0.0.1:3306)/pgmanagement")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		var check_booking_cus models.Booking_pg
		err = c.BindJSON(&check_booking_cus)
		if err != nil {
			return
		}
		query := fmt.Sprintf("SELECT Booking_ID,Customer_Name,Cus_Contact_No,Customer_ID,Property_ID,From_date,To_Date FROM BookingDetails WHERE Customer_ID = '%s'", check_booking_cus.Customer_ID)
		results, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		defer results.Close()
		var output interface{}
		for results.Next() {
			var Booking_id int
			var Customer_name string
			var Cus_contact_no string
			var Customer_id string
			var Property_id int
			var From_date string
			var To_date string
			err = results.Scan(&Booking_id, &Customer_name, &Cus_contact_no, &Customer_id, &Property_id, &From_date, &To_date)
			if err != nil {
				panic(err.Error())
			}
			c.IndentedJSON(200, "See Your Booking")
			output = fmt.Sprintf("Booking_Id=%d,Customer_Name='%s',Cus_Contact_No='%s',Customer_Id='%s',Property_Id=%d,From_Date='%s',To_Date='%s'", Booking_id, Customer_name, Cus_contact_no, Customer_id, Property_id, From_date, To_date)
			c.JSON(http.StatusOK, gin.H{"": output})
		}
	}
}

func Update_Booking() gin.HandlerFunc {

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
		query := fmt.Sprintf("UPDATE BookingDetails SET Customer_Name='%s',Cus_Contact_No='%s',From_Date='%s',To_Date='%s' WHERE Booking_ID=%d", edit_booking.Customer_Name, edit_booking.Cus_Contact_No, edit_booking.From_Date, edit_booking.To_Date, edit_booking.Booking_ID)
		_, err = db.Exec(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, "Yes, Booking Update Successfully!")
		c.IndentedJSON(http.StatusCreated, edit_booking)
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
		c.IndentedJSON(201, "Yes, Booking Delete Successfully!")
		defer results.Close()
	}
}
