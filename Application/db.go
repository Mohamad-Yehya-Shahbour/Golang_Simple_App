package Application

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func makeConnection()*gorm.DB{

	dsn := "root:@tcp(127.0.0.1:3306)/golangdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error while connecting the database", err.Error())
	}

	return db
}
