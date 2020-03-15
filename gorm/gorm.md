## Gorm

### 安装
```golang
go get -u github.com/jinzhu/gorm
```

### 使用
```golang
package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Product struct {
	gorm.Model  //提供ID,TIME默认字段
	Code  string `gorm:"type:varchar"`
	Price uint
}

func main() {
	db, err := gorm.Open("postgres", "host=127.0.0.1  port=5432 user=pguser dbname=testdb sslmode=disable password=pgpass")
	db.SingularTable(true) //去掉复数表名的约定。默认小写，驼峰变下划线，加复数
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
	db.Last(&lastProduct)  //获取最后添加的数据
	lastProduct.Price = 5000
	//保存
	db.Save(&lastProduct) //保存，没有就创建，有就更新

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	db.Delete(&product)
}

```

### 数据库连接
```golang
//Mysql
import _ "github.com/jinzhu/gorm/dialects/mysql"
db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
//PostgreSql
import _ "github.com/jinzhu/gorm/dialects/postgres"
db, err := gorm.Open("postgres", "host=myhost user=gorm dbname=gorm sslmode=disable password=mypassword")
```

### gorm.Model
```golang
type Model struct {
  ID        uint `gorm:"primary_key"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt *time.Time
}
```

### 索引
```golang
// 为`name`列添加索引`idx_user_name`
db.Model(&User{}).AddIndex("idx_user_name", "name")

// 为`name`, `age`列添加索引`idx_user_name_age`
db.Model(&User{}).AddIndex("idx_user_name_age", "name", "age")

// 添加唯一索引
db.Model(&User{}).AddUniqueIndex("idx_user_name", "name")

// 为多列添加唯一索引
db.Model(&User{}).AddUniqueIndex("idx_user_name_age", "name", "age")

// 删除索引
db.Model(&User{}).RemoveIndex("idx_user_name")
```

## GORM 中文文档  
 > [http://gorm.book.jasperxu.com](http://gorm.book.jasperxu.com)\
 > [http://gorm.io/zh_CN/docs](http://gorm.io/zh_CN/docs)