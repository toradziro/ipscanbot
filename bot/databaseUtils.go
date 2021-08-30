package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type userRequest struct {
	userName string
	Request string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "bot"
	password = "root"
	dbname   = "ip_scan_bot"
)

func collectDbName() string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " + "password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	return psqlInfo
}

func openDb() *sql.DB {
	db, err := sql.Open("postgres", collectDbName())
	if err != nil {
		log.Printf("\n\n%v\n\n", err)
	}
	return db
}

func isAdmin(db *sql.DB, username string) bool {
	var isExist bool
	err := db.QueryRow("SELECT * FROM admins WHERE admin_username = ($1)", username).Scan(&isExist)
	if err != nil {
		log.Printf("%v\n", err)
	}
	return isExist
}

func addInDatabase(username string, userId int, request string, db *sql.DB) {
	_, err := db.Exec("INSERT INTO user_requests(username, user_id, request) VALUES ($1, $2, $3)", username, userId, request)
	if err != nil {
		log.Printf("\n\n%v\n\n", err)
	}
}

func getLastReq(username string, db *sql.DB) string {
	var res, tmp string
	rows, err := db.Query("SELECT DISTINCT request FROM user_requests WHERE username = ($1)", username)
	if err != nil {
		log.Printf("\n\n%v\n\n", err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&tmp)
		res += tmp + "\n\n"
		if err != nil {
			log.Printf("\n\n%v\n\n", err)
		}
	}
	return res
}