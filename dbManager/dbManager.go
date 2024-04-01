package dbManager

import (
	"gorm.io/gorm"
	"tigerhall/middleware"
)

func GetDbInstance() (*gorm.DB, error) {
	var dbc = middleware.DBC
	return dbc, nil
}
