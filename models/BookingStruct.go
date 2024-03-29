package models

type Booking_pg struct {
	Customer_Name  string `json:"customer_name" validate:"required"`
	Cus_Contact_No string `json:"cus_contact_no" validate:"required"`
	Property_ID    int    `json:"property_id" validate:"required"`
	Booking_ID     int    `json:"booking_id"`
	Customer_ID    string `json:"customer_id"`
	From_Date      string `json:"from_date"`
	To_Date        string `json:"to_date"`
}
