package db

import "database/sql"

func Init() *sql.DB {
	dbURL := "user=din password=743 dbname=bank-di sslmode=disable"
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil
	}
	return db
}
