package tigerController

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"tigerhall/apps/tiger/serializer/tigerSerializer"
	"tigerhall/config"
	"tigerhall/config/constants"
	"tigerhall/dbManager"
	"tigerhall/dtos"
	"tigerhall/models"
	"tigerhall/repositories/tigerRepository"
	"time"
)

func CreateUser(user dtos.User) (dtos.UserResponse, error) {
	db, _ := dbManager.GetDbInstance()
	userData, err := tigerRepository.GetUser(db, user)
	if err != nil {
		return dtos.UserResponse{}, err
	}
	if userData.UserName == "" {
		return dtos.UserResponse{}, config.ValidationError(constants.UserAlreadyExists)
	}
	userData, err = tigerRepository.CreateUser(db, user)
	if err != nil {
		return dtos.UserResponse{}, err
	}
	return tigerSerializer.UserSerializer(userData)
}

func LoginUser(login dtos.Login) (dtos.LoginResponse, error) {
	db, _ := dbManager.GetDbInstance()
	userData, err := tigerRepository.GetUserByEmail(db, login)
	if err != nil {
		return dtos.LoginResponse{}, err
	}
	if userData.UserName == "" {
		return dtos.LoginResponse{}, config.ValidationError(constants.UserAlreadyExists)
	}

	if !validatePassword(login.Password, userData.Password) {
		return dtos.LoginResponse{}, config.ValidationError(constants.PasswordNotValid)
	}
	accessTokenDtos := CreateJwtToken(db, userData)
	return tigerSerializer.LoginSerializer(accessTokenDtos)
}

func validatePassword(passwordHash string, password string) bool {
	// todo: add the password matching algorithm here
	return true
}
func CreateTiger(tiger dtos.TigerDetails) (dtos.TigerResponse, error) {
	db, _ := dbManager.GetDbInstance()
	tigerInfo, err := tigerRepository.GetTiger(db, tiger)
	if err != nil {
		return dtos.TigerResponse{}, nil
	}
	if tigerInfo.Name != "" {
		return tigerSerializer.TigerSerializer(tigerInfo)
	}
	tigerData, err := tigerRepository.CreateTiger(db, tiger)
	if err != nil {
		return dtos.TigerResponse{}, nil
	}
	return tigerSerializer.TigerSerializer(tigerData)
}

func CreateJwtToken(db *gorm.DB, userData models.User) dtos.AccessToken {
	var secretKey = []byte("secret-key")
	expiryTime := time.Now().Add(time.Hour * 24)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username":     userData.UserName,
			"phone_number": userData.PhoneNumber,
			"email":        userData.Email,
			"exp":          expiryTime,
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println(err)
	}
	accessToken := dtos.AccessToken{
		Token:      tokenString,
		UserId:     userData.Id,
		ExpiryTime: expiryTime,
	}
	data, err := tigerRepository.CreateToken(db, accessToken)
	if err != nil {
		fmt.Println(err)
	}
	accessTokenDtos := dtos.AccessToken{
		Token: data.AccessToken,
	}
	return accessTokenDtos
}

func TigerDetails() (map[string]interface{}, error) {
	db, _ := dbManager.GetDbInstance()
	tigerRepository.GetTigerByLastSeen(db)
	//if err != nil {
	//	fmt.Println(err, tigerInfo)
	//}
	return map[string]interface{}{}, nil
}
