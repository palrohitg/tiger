package tigerSerializer

import (
	"encoding/json"
	"tigerhall/dtos"
	"tigerhall/models"
)

func UserSerializer(user models.User) (dtos.UserResponse, error) {
	var userResponse dtos.UserResponse
	cardString, _ := json.Marshal(user)
	_ = json.Unmarshal(cardString, &userResponse)
	return userResponse, nil
}

func LoginSerializer(accessToken dtos.AccessToken) (dtos.LoginResponse, error) {
	var loginResponse dtos.LoginResponse
	cardString, _ := json.Marshal(accessToken)
	_ = json.Unmarshal(cardString, &loginResponse)
	return loginResponse, nil
}

func TigerSerializer(tiger models.Tiger) (dtos.TigerResponse, error) {
	var tigerResponse dtos.TigerResponse
	cardString, _ := json.Marshal(tiger)
	_ = json.Unmarshal(cardString, &tigerResponse)
	return tigerResponse, nil
}

func TigerSightSerializer(tigerSight models.TigerSight) (dtos.TigerSightResponse, error) {
	var tigerSightResponse dtos.TigerSightResponse
	cardString, _ := json.Marshal(tigerSight)
	_ = json.Unmarshal(cardString, &tigerSightResponse)
	return tigerSightResponse, nil
}

func TigerWithSightDetailsSerializer(tigerWithSight []dtos.TigerWithSightDetails) ([]dtos.TigerWithSightDetailsResponse, error) {
	var tigerWithSightResponse []dtos.TigerWithSightDetailsResponse
	cardString, _ := json.Marshal(tigerWithSight)
	_ = json.Unmarshal(cardString, &tigerWithSightResponse)
	return tigerWithSightResponse, nil
}
