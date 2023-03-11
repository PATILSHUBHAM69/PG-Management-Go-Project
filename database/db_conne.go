package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() {
	var err error

	// Make Database Connection
	db, err = sql.Open("mysql", "root:india@123@tcp(localhost:3306)/pgmanagement")

	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MySQL database!")

	//Create Tabel
	Create_Property, err := db.Query("CREATE TABLE IF NOT EXISTS Property_Details (Property_ID INT NOT NULL AUTO_INCREMEENT=1001, Property_Name varchar(25), Contact_No varchar(15), Property_Type varchar(255), Property_Address varchar(200), City varchar(20), Pin_Code varchar(6), Landmark varchar(30), Ammeneties varchar(200), Price varchar(15), Advance_Deposit varchar(15);")
	if err != nil {
		panic(err.Error())
	}
	defer Create_Property.Close()
	fmt.Println("Property Details Successfully Created")

	Create_Booking, err := db.Query("CREATE TABLE IF NOT EXISTS Booking_Details (Customer_ID INT NOT NULL AUTO_INCREMEENT=101,Customer_Name varchar(25), Cus_Contact_No varchar(15), Property_ID INT, Booking_time varchar(15), Booking_ID INT, From_Date varchar(20), To_Date varchar(20);")
	if err != nil {
		panic(err.Error())
	}
	defer Create_Booking.Close()
	fmt.Println("Property Details Successfully Created")

}
