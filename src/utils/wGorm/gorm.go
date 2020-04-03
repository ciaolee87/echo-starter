package wGorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func Connect() {

	// mysql, pstgreSQL, Sqlite3 ... 등등등
	dialect := os.Getenv("DB_DIALECT")

	//user:password@(localhost)/dbname?charset=utf8mb4&parseTime=True
	id := os.Getenv("DB_ID")
	passwd := os.Getenv("DB_PASSWD")
	host := os.Getenv("DB_HOST")
	schema := os.Getenv("DB_SCHEMA")
	args := os.Getenv("DB_ARGS")
	query := fmt.Sprintf("%s:%s@%s/%s?%s", id, passwd, host, schema, args)

	// DB 연결
	db, err := gorm.Open(dialect, query)

	defer db.Close()
	if err != nil {
		log.Fatal("Fail DB Connection")
	}

	DB = db
}
