package db

import (
	"api/config"
	"api/domain/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
  
func NewDB(config *config.Config) *gorm.DB {
	dsn := config.DB_USER + ":" + config.DB_PASS + "@tcp(" + config.DB_HOST + ":3306)/" + config.DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Post{}, &model.Tag{})

	return db
}
