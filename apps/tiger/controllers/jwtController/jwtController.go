package jwtController

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"tigerhall/dtos"
	"tigerhall/models"
	"tigerhall/repositories/tigerRepository"
	"time"
)

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
