package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID				primitive.ObjectID 		`bson:"_id"`
	First_name		*string 				`json:"first_name" validate:"required, min=2, max=100"`
	Last_name		*string 				`json:"last_name"`
	Email			*string 				`json:"email" validate:"email, required"`
	Password		*string 				`json:"password" validate:"required, min=8"`
	Phone			*int64 					`json:"phone" validate:"required, max=10"`
	Aadhar			*int64 					`json:"aadhar" validate:"required, min=12, max=12"`
	Token			*string 				`json:"token"`
	Refresh_token	*string 				`json:"refresh_token"`
	User_type		*string					`json:"user_type" validate:"required, eq=ADMIN|eq=USER"`
	Created_at		time.Time				`json:"created_at"`
	Updated_at		time.Time				`json:"updated_at"`
	User_id			*string 				`json:"user_id"`
}
