package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

var MysqlDb *gorm.DB

func ConnectMysql() {
	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	database := os.Getenv("MYSQL_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	MysqlDb, _ = gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if MysqlDb.Error != nil {
		panic(MysqlDb.Error)
	}
}
