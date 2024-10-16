package mydb

import (
	"database/sql"
	"log"
)

func DBInit(db *sql.DB) {
	_, errD := db.Exec(`CREATE DATABASE IF NOT EXISTS goo`)

	if errD != nil {
		log.Fatal(errD)
		panic(errD)
	}

	db.Exec(`USE goo`)

	_, errT := db.Exec(`CREATE TABLE IF NOT EXISTS user (
	id_user INT auto_increment PRIMARY KEY,
	name VARCHAR(50)
	)`)

	if errT != nil {
		panic(errT)
	}
}
