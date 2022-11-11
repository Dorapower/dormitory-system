package database

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var MysqlDb *gorm.DB

func ConnectMysql() {
	mysqlDsn := "root:abc123321A@tcp(47.92.123.159:8080)/login_demo?charset=utf8mb4&parseTime=True&loc=Local"
	sqlDB, err := sql.Open("mysql", mysqlDsn)
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
