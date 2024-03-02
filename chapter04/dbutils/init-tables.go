package dbutils

import (
	"database/sql"
	"log"
)

func Initialize(dbDriver *sql.DB) {
	statement, driverError := dbDriver.Prepare(train)
	if driverError != nil {
		log.Println(driverError)
	}
	// Create train table
	_, err := statement.Exec()
	if err != nil {
		log.Println("Table already exists!")
	}
	statement, _ = dbDriver.Prepare(station)
	_, err = statement.Exec()
	if err != nil {
		panic(err)
	}
	statement, _ = dbDriver.Prepare(schedule)
	_, err = statement.Exec()
	if err != nil {
		panic(err)
	}
	log.Println("All tables created/initialized successfully!")
}
