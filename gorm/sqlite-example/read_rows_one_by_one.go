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
	var product Product
	rows, err := db.Model(&Product{}).Rows() // read Product model , row one by one
	if err != nil { // check err
		panic("failed to connect db")
	}
	defer rows.Close() // close when done
	for rows.Next() { // iterator
		db.ScanRows(rows, &product) // copy to product
		fmt.Println(product)
	}
}
