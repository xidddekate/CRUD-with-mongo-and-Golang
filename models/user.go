package models

import "time"

type User struct {
	ID          interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string      `json:"name" bson:"name"`
	DOB         string      `json:"dob" bson:"dob"`
	Address     string      `json:"address" bson:"address"`
	Description string      `json:"description" bson:"description"`
	CreatedAt   time.Time   `bson:"created_at" json:"created_at,omitempty"`
}

type UserUpdate struct {
	ModifiedCount int64 `json:"modifiedCount"`
	Result        User
}

type UserDelete struct {
	DeletedCount int64 `json:"deletedCount"`
}
