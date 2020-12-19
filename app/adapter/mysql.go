package adapter

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbsql *gorm.DB
)

// LoadMySQL is load connection to mysql server
func LoadMySQL(url string) {
	dbsql = MySQL(url)
}

// MySQL is open connection to mysql server
func MySQL(url string) *gorm.DB {
	conn, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return conn
}

// DBSQL is open connection into database
func DBSQL() *gorm.DB {
	return dbsql
}
