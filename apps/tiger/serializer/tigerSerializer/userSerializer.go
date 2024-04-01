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
