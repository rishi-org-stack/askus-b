package main

import (
	"askUs/v1/package/auth"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	gpsql "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	psn := "postgresql://pgadmin:password@localhost/askus?sslmode=disable"

	sqlDB, err := sql.Open("postgres", psn)
	checkErr(err)
	gdb, err := gorm.Open(gpsql.New(gpsql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	tx := gdb.Exec("DROP SCHEMA IF EXISTS auth CASCADE")
	checkErr(tx.Error)
	tx = gdb.Exec("CREATE SCHEMA IF NOT EXISTS auth")
	checkErr(tx.Error)
	ok := gdb.AutoMigrate(&auth.AuthRequest{})
	checkErr(ok)
}
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
