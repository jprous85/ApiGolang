package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

func connect() *gorm.DB {
	db, err := gorm.Open("mysql", "ubuntu:ubuntu@tcp(127.0.0.1:3306)/apicrudgolang?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connect to bbdd successfully!")
	return db
}
