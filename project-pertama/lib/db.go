package lib

import (
	_"database/sql"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error){
	connectionString := "host=localhost port=5432 user=postgres password=hasibuan dbname=hacktiv sslmode=disable"
	// db, err := sql.Open("postgres", connectionString)
	// if err != nil {
	// 	return nil, err
	// }

	// err = db.Ping()
	// if err != nil {
	// 	return nil, err
	// }

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	return db, err
}