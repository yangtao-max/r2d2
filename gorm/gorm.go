package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Product struct {
	gorm.Model
	Code  string `gorm:"type:varchar"`
	Price uint
}

func main() {
	db, err := gorm.Open("postgres", "host=127.0.0.1  port=5432 user=pguser dbname=testdb sslmode=disable password=pgpass")
	db.SingularTable(true)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// 初始化数据表
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	var product Product
	db.First(&product, 1)                   // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212
	var lastProduct Product
	db.Last(&lastProduct)
	lastProduct.Price = 5000
	//保存
	db.Save(&lastProduct)

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	db.Delete(&product)
}
