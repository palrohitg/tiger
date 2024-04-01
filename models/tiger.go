package models

import "time"

/*
 Use the BaseClass Concepts for common struct here
*/

type BaseModel struct {
	IsActive  bool      `json:"bool"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Tiger struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	BaseModel
}

func (b *Tiger) TableName() string {
	return "tiger"
}

type TigerSight struct {
	Id       int    `json:"id"`
	TigerId  int    `json:"tiger_id"`
	Lat      string `json:"lat"`
	Long     string `json:"long"`
	ImageURL string `json:"image_url"`
	LastSeen string `json:"last_seen"`
	BaseModel
}

func (b *TigerSight) TableName() string {
	return "tiger_sight"
}

type User struct {
	Id          int    `json:"id"`
	UserName    string `json:"user_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	BaseModel
}

func (b *User) TableName() string {
	return "user"
}

type AccessToken struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id"`
	AccessToken string    `json:"access_token"`
	ExpiryTime  time.Time `json:"expiry_time"`
	BaseModel
}

func (b *AccessToken) TableName() string {
	return "access_token"
}
