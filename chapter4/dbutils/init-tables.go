package dbutils

import (
	"database/sql"
	"log"
)

func Initialize(dbDriver *sql.DB){
    statement, driverError := dbDriver.Prepare(trains)

    if driverError != nil {
        log.Println(driverError)
    }

    _,statementError := statement.Exec()

    if statementError != nil {
        log.Println(statementError)
    }

    statement, driverError = dbDriver.Prepare(station)

    if driverError != nil {
        log.Println(driverError)
    }

    _,statementError = statement.Exec()

    if statementError != nil {
        log.Println(statementError)
    }

    statement, driverError = dbDriver.Prepare(schedule)

    if driverError != nil {
        log.Println(driverError)
    }

    _,statementError = statement.Exec()

    if statementError != nil {
        log.Println(statementError)
    }

    log.Println("All tables created/initialized successfully")
}