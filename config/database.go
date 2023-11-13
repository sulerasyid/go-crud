package config

import (
	"os"

	"github.com/sulerasyid/go-crud/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDB() (*gorm.DB, error) {

	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")

	db, err := gorm.Open("mysql", dbUsername+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?parseTime=true")
	if err != nil {
		return nil, err
	}

	migrateDDL(db)

	return db, nil

}

func migrateDDL(db *gorm.DB) {
	db.AutoMigrate(&model.Tags{})
}
