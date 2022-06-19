package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// db connection
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate : creating schema?
	db.AutoMigrate((&Product{}))
	
	// create
	db.Create(&Product{Code: "D42", Price: 100})
	
	// Read
	var product Product // to store result
	db.First(&product, 1) // find product with primary key 1
	fmt.Println(product) // result stored in product
}
