package main

import (
	"database/sql"
	"ecom/cmd/api"
	"ecom/config"
	db2 "ecom/db"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	server := api.NewAPIServer(":8080", nil)

	db, err := db2.NewMySQQStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	_ = err
	_ = db

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfull connected! ")
}
