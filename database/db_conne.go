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
	Create_Property, err := db.Query("CREATE TABLE IF NOT EXISTS PropertyDetails (Property_ID INT NOT NULL AUTO_INCREMENT, Property_Name VARCHAR(25), Contact_No VARCHAR(255), Property_Type VARCHAR(255), Property_Address VARCHAR(200), City VARCHAR(20), Pin_Code VARCHAR(6), Landmark VARCHAR(30), Ammeneties VARCHAR(200), Price VARCHAR(15), Advance_Deposit VARCHAR(15), PRIMARY KEY (Property_ID) );")
	if err != nil {
		panic(err.Error())
	}
	defer Create_Property.Close()
	fmt.Println("Property Details Successfully Created")

	Create_Booking, err := db.Query("CREATE TABLE IF NOT EXISTS BookingDetails (Customer_ID INT,Customer_Name VARCHAR(25), Cus_Contact_No VARCHAR(15), Property_ID INT, Booking_time VARCHAR(15), Booking_ID INT NOT NULL AUTO_INCREMENT, From_Date VARCHAR(20), To_Date VARCHAR(20), PRIMARY KEY (Booking_ID) );")
	if err != nil {
		panic(err.Error())
	}
	defer Create_Booking.Close()
	fmt.Println("Property Details Successfully Created")

}
