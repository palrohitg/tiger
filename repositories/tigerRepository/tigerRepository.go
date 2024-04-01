package tigerRepository

import (
	"fmt"
	"gorm.io/gorm"
	"tigerhall/config/constants"
	"tigerhall/dtos"
	"tigerhall/models"
	"time"
)

func CreateUser(db *gorm.DB, user dtos.User) (models.User, error) {
	userDtos := models.User{
		UserName:    user.UserName,
		Password:    user.Password,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}
	userDtos.CreatedAt = time.Now()
	userDtos.UpdatedAt = time.Now()
	userDtos.IsActive = true
	if err := db.Table(constants.User).Create(&userDtos).Error; err != nil {
		return userDtos, err
	}
	return userDtos, nil
}

func GetUser(db *gorm.DB, user dtos.User) (models.User, error) {
	userDtos := models.User{}
	res := db.Table(constants.User).
		Where("email = ?  OR  phone_number = ?", user.Email, user.PhoneNumber).
		Find(&userDtos)
	if res.Error != nil {
		return userDtos, res.Error
	}
	if res.RowsAffected != 0 {
		return userDtos, nil
	}
	return userDtos, nil
}

func GetUserByEmail(db *gorm.DB, login dtos.Login) (models.User, error) {
	userDtos := models.User{}
	res := db.Table(constants.User).
		Where("email = ? AND is_active = ?", login.Email, true).
		Find(&userDtos)
	if res.Error != nil {
		return userDtos, res.Error
	}
	if res.RowsAffected != 0 {
		return userDtos, nil
	}
	return userDtos, nil
}

func CreateToken(db *gorm.DB, accessToken dtos.AccessToken) (models.AccessToken, error) {
	accessTokenDtos := models.AccessToken{
		AccessToken: accessToken.Token,
		UserId:      accessToken.UserId,
		ExpiryTime:  accessToken.ExpiryTime,
	}
	accessTokenDtos.CreatedAt = time.Now()
	accessTokenDtos.UpdatedAt = time.Now()
	accessTokenDtos.IsActive = true
	if err := db.Table(constants.AccessToken).Create(&accessTokenDtos).Error; err != nil {
		return accessTokenDtos, err
	}
	return accessTokenDtos, nil
}

func CreateTiger(db *gorm.DB, tiger dtos.TigerDetails) (models.Tiger, error) {
	tigerDtos := models.Tiger{
		Name: tiger.Name,
		DOB:  tiger.DOB,
	}
	tigerDtos.CreatedAt = time.Now()
	tigerDtos.UpdatedAt = time.Now()
	tigerDtos.IsActive = true
	if err := db.Table(constants.Tiger).Create(&tigerDtos).Error; err != nil {
		return tigerDtos, err
	}
	tigerSight := models.TigerSight{
		TigerId:  tigerDtos.Id,
		Lat:      tiger.Lat,
		Long:     tiger.Long,
		LastSeen: tiger.LastSeen,
	}
	CreateTigerDetail(db, tigerSight)
	return tigerDtos, nil
}

func CreateTigerDetail(db *gorm.DB, tigerSight models.TigerSight) (models.TigerSight, error) {
	tigerSight.CreatedAt = time.Now()
	tigerSight.UpdatedAt = time.Now()
	tigerSight.IsActive = true
	if err := db.Table(constants.TigerSight).Create(&tigerSight).Error; err != nil {
		return tigerSight, nil
	}
	return tigerSight, nil
}

func GetTiger(db *gorm.DB, tiger dtos.TigerDetails) (models.Tiger, error) {
	tigerDtos := models.Tiger{}
	res := db.Table(constants.Tiger).
		Where("name = ? and is_active = ?", tiger.Name, true).
		Find(&tigerDtos)
	if res.Error != nil {
		return tigerDtos, nil
	}
	if res.RowsAffected != 0 {
		return tigerDtos, nil
	}
	return tigerDtos, nil
}

/*
res := db.Table(constants.Tiger).

	Joins("JOIN zoo ON tiger.zoo_id = zoo.id").
	Where("tiger.name = ? AND tiger.is_active = ?", tiger.Name, true).
	Select("tiger.*, zoo.name as zoo_name"). // Select the columns you want
	Find(&tigerDtos)
*/
func GetTigerByLastSeen(db *gorm.DB) {
	// Make the Join query to filter out the records once.
	// Find the elements
	tigerDtos := models.Tiger{}
	res := db.Table(constants.Tiger).
		Joins("JOIN tiger_sight ON tiger.id = tiger_sight.tiger_id").
		Find(&tigerDtos)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	fmt.Println(tigerDtos)
	return
}
