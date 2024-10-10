package db

import "database/sql"

func Connect() *sql.DB {
	dsn := ""
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return db
}
