package models

type User struct {
	UserId    int64  `json:"user_id" bson:"user_id"`
	UserName  string `json:"username" bson:"username"`
	FirstName string `json:"first_name" bson:"first_name"`
}
