package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID               primitive.ObjectID `bson:"_id"`
	First_name       *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_name        *string            `json:"last_name" validate:"required,min=2,max=100"`
	Password         *string            `json:"Password" validate:"required,min=6"`
	Email            *string            `json:"email" validate:"email,required"`
	Phone            *string            `json:"phone" validate:"required"`
	Token            *string            `json:"token"`
	User_type        *string            `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Refresh_token    *string            `json:"refresh_token"`
	Created_at       time.Time          `json:"created_at"`
	Updated_at       time.Time          `json:"updated_at"`
	User_id          *string            `json:"user_id"`
	Property_ID      *int               `json:"property_id"`
	Property_Name    *string            `json:"property_name"`
	Contact_No       *string            `json:"contact_no"`
	Property_Type    *string            `json:"property_type"`
	Property_Address *string            `json:"property_address"`
	City             *string            `json:"city"`
	Pin_Code         *string            `json:"pin_code"`
	LandMark         *string            `json:"landmark"`
	Property_Image   *string            `json:"property_image"`
	Ammeneties       *string            `json:"ammeneties"`
	Price            *string            `json:"price"`
	Advance_Deposit  *string            `json:"advance_deposit"`
}
