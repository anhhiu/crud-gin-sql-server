package database

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB

// Kết nối cơ sở dữ liệu
func ConnectDB() {
	var err error
	connString := "server=LAPTOP-7CAHEI3Q\\HATHANHHAO;user id=sa;password=hao123;database=Student;"
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err.Error())
	}
}

// Trả về đối tượng db để sử dụng trong các hàm CRUD
func GetDB() *sql.DB {
	return db
}
