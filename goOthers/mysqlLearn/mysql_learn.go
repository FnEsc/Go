package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // sql open driver 会引用
	"time"
)

var MysqlDb *sql.DB
var Err error

const (
	USERNAME = "test"
	PASSWORD = "123456"
	HOST     = "127.0.0.1"
	PORT     = "3306"
	DATABASE = "test"
	CHARSET  = "utf8"
)

type Demo struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Status int    `json:"status"`
	Sex    int    `json:"sex"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
	Qq     string `json:"qq"`
}

func init() {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USERNAME, PASSWORD, HOST, PORT, DATABASE, CHARSET)
	// 打开连接失败
	MysqlDb, Err = sql.Open("mysql", dbDSN)
	if Err != nil {
		panic("数据源配置错误: " + Err.Error())
	}
	// 最大连接数
	MysqlDb.SetMaxOpenConns(100)
	// 闲置连接数
	MysqlDb.SetMaxIdleConns(20)
	// 最大连接周期
	MysqlDb.SetConnMaxLifetime(100 * time.Second)
	if Err = MysqlDb.Ping(); nil != Err {
		panic("数据库链接失败: " + Err.Error())
	}
	// defer MysqlDb.Close();
}

func MysqlQueryRow() {
	var demo Demo
	// 查询 QueryRow 返回一条
	row := MysqlDb.QueryRow("SELECT `id`, `name` FROM  demo WHERE `status` = ?", 1)
	Err = row.Scan(&demo.Id, &demo.Name)
	if Err == nil {
		fmt.Println(demo.Name, demo.Id)
	} else {
		fmt.Println(Err.Error())
	}
}

func MysqlQuery() {
	// 查询 Query
	var demos []Demo
	var demo Demo
	rows, Err := MysqlDb.Query("SELECT `id`, `name` FROM  demo WHERE `status` = ?", 1)
	if Err == nil {
		// 依次循环取
		for rows.Next() {
			rows.Scan(&demo.Id, &demo.Name)
			demos = append(demos, demo)
		}
		fmt.Println(demos)
	} else {
		fmt.Println(Err.Error())
	}
}

func MysqlInsert() {
	res, Err := MysqlDb.Exec("INSERT INTO demo(`name`, `age`, `status`, `sex`, `email`, `mobile`, `qq`) "+
		"VALUES(?, ?, ?, ?, ?, ?, ?)", "孙策", "20", "2", "1", "jiangdong@126.com", "13366666666", "555333999")
	if Err == nil {
		// 获取最后插入ID
		num, Err := res.LastInsertId()
		if Err == nil {
			fmt.Println(num)
		} else {
			fmt.Println(Err.Error())
		}
	} else {
		fmt.Println(Err.Error())
	}
}

func MysqlUpdate() {
	res, Err := MysqlDb.Exec("UPDATE demo SET `name` = ? WHERE `id` = ?", "马超", 40)
	if Err == nil {
		// 获取影响行
		num, Err := res.RowsAffected()
		if Err == nil {
			fmt.Println(num)
		} else {
			fmt.Println(Err.Error())
		}
	} else {
		fmt.Println(Err.Error())
	}
}

func MysqlDelete() {
	res, Err := MysqlDb.Exec("DELETE FROM demo WHERE `id` = ?", 39)
	if Err == nil {
		// 获取影响行
		num, Err := res.RowsAffected()
		if Err == nil {
			fmt.Println(num)
		} else {
			fmt.Println(Err.Error())
		}
	} else {
		fmt.Println(Err.Error())
	}
}

func MysqlTx() {
	// 开启事务
	tx, Err := MysqlDb.Begin()
	if Err == nil {
		res1, _ := tx.Exec("UPDATE demo SET `name` = ? WHERE `id` = ?", "赵云", 40)
		res2, _ := tx.Exec("UPDATE demo SET `name` = ? WHERE `id` = ?", "赵云", 38)
		num1, _ := res1.RowsAffected()
		num2, _ := res2.RowsAffected()
		if num1 > 0 && num2 > 0 {
			// 提交
			Err = tx.Commit()
			if Err == nil {
				fmt.Println("Success")
			} else {
				fmt.Println(Err.Error())
			}
		} else {
			// 回滚
			Err = tx.Rollback()
			if Err == nil {
				fmt.Println("Fail")
			} else {
				fmt.Println(Err.Error())
			}
		}
	} else {
		fmt.Println(Err.Error())
	}
}

func main() {
	MysqlTx()
}
