package dtos

import "time"

type User struct {
	UserName    string `json:"user_name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phone_number"`
}

type UserResponse struct {
	UserName    string `json:"user_name,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TigerDetails struct {
	Name     string `json:"name" binding:"required"`
	DOB      string `json:"dob" binding:"required"`
	LastSeen string `json:"last_seen" binding:"required"`
	Lat      string `json:"lat" binding:"required"`
	Long     string `json:"long" binding:"required"`
}

type AccessToken struct {
	Token      string    `json:"token"`
	UserId     int       `json:"user_id"`
	ExpiryTime time.Time `json:"expiry_time"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type TigerResponse struct {
	Id   int    `json:"int"`
	Name string `json:"name"`
}

type TigerSightDetails struct {
	TigerId  int    `json:"tiger_id" binding:"required"`
	LastSeen string `json:"last_seen" binding:"required"`
	Lat      string `json:"lat" binding:"required"`
	Long     string `json:"long" binding:"required"`
}
