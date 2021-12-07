package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "test:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	// 链接方法一：gorm 数据库链接 mysql
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 链接方法二：gorm 通过现在sql初始化 *gorm.DB
	//sqlDB, err := sql.Open("mysql", dsn)
	//gormDB, err := gorm.Open(mysql.New(mysql.Config{
	//	Conn: sqlDB,
	//}), &gorm.Config{})

	db.Migrator().DropTable(&Product{})

	// 迁移 schema 自动迁移只增不减
	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;").AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})
	db.Create(&Product{Code: "D42中文", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // 根据整形主键查找
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	db.Delete(&product)

}
