package main

import (
	"askUs/v1/package/advice"
	"askUs/v1/package/auth"
	"askUs/v1/package/user"
	"askUs/v1/util/cache"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	gpsql "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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
	tx = gdb.Exec("DROP SCHEMA IF EXISTS usr CASCADE")
	checkErr(tx.Error)
	tx = gdb.Exec("CREATE SCHEMA IF NOT EXISTS usr")
	checkErr(tx.Error)
	ok = gdb.AutoMigrate(
		&user.Doctor{},
		&user.Patient{},
		&user.Experience{},
		&user.Institution{},
		&user.Address{}, &user.Request{}, &user.FollowedByDoctor{},
		&user.FollowingDoctor{},
		&user.FollowedByPatient{},
		&user.FollowedDoctorsByPatient{},
	)
	checkErr(ok)
	tx = gdb.Exec("DROP SCHEMA IF EXISTS advice CASCADE")
	checkErr(tx.Error)
	tx = gdb.Exec("CREATE SCHEMA IF NOT EXISTS advice")
	checkErr(tx.Error)
	ok = gdb.AutoMigrate(&advice.Advice{}, &advice.Like{})
	checkErr(ok)
	clearCache()
}
func clearCache() {
	rdb, err := cache.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	rdb.DB.FlushAll(context.TODO())
}
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
