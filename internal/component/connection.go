package component

import (
	"ecomtest/domain"
	"ecomtest/internal/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDatabaseConnection(cnf *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True",
		cnf.DB.User, cnf.DB.Pass, cnf.DB.Host, cnf.DB.Port, cnf.DB.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error when connect to database:%s", err.Error())
	}
	if err := db.AutoMigrate(&domain.Product{}); err != nil {
		log.Fatalf("migrate error:%s", err.Error())
	}
	fmt.Println("database connect")

	return db
}
