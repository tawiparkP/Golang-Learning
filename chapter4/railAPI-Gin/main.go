package main

import (
	"time"
	"net/http"
	"log"
    "database/sql"
    "github.com/gin-gonic/gin"
    _ "github.com/mattn/go-sqlite3"
    "github.com/tawiparkP/dbutils"
)

var DB *sql.DB

type StationResource struct {
    ID int `json:id`
    Name string `json:name`
    OpeningTine time.Time `json:opening_time`
    ClosingTime time.Time `json:closing_time`
}

func GetStation(c *gin.Context){
    var station StationResource
    id := c.Param("station_id")
    err := DB.QueryRow("SELECT ID, NAME, OPENING_TIME, CLOSING_TIME FROM station WHERE id=?",id).Scan(&station.ID, &station.Name, &station.OpeningTine, &station.ClosingTime)
    if err != nil {
        log.Println(err)
        c.JSON(500, gin.H{"error": err.Error()})
    } else {
        c.JSON(200, gin.H{"result": station})
    }
}

func CreateStation(c *gin.Context){
    var station StationResource
    if err := c.BindJSON(&station); err == nil {
        statement, _ := DB.Prepare("INSERT INTO station (NAME, OPENING_TIME, CLOSING_TIME) VALUES (?,?,?)")
        result, err := statement.Exec(station.Name,station.OpeningTine,station.ClosingTime)
        if err == nil {
            newID, _ := result.LastInsertId()
            station.ID = int(newID)
            c.JSON(http.StatusOK, gin.H{"result": station})
        } else {
            c.String(http.StatusInternalServerError, err.Error())
        }
    }
}

func main(){
    var err error
    DB, err = sql.Open("sqlite3", "./railapi.db?parseTime=true")
    if err != nil {
        log.Println("Database Driver creation failed!")
    }

    dbutils.Initialize(DB)

    r := gin.Default()
    r.GET("/v1/stations/:station_id", GetStation)
    r.POST("/v1/stations", CreateStation)

    r.Run(":8000")
}


