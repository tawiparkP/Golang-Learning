package main

import (
	"time"
	"net/http"
	"encoding/json"
	"log"
    "database/sql"
    "github.com/emicklei/go-restful"
    _ "github.com/mattn/go-sqlite3"
    "github.com/tawiparkP/dbutils"
)

var DB *sql.DB

type TrainResource struct{
    ID int
    DriverName string
    OperatingStatus bool
}

type StationResource struct {
    ID int
    Name string
    OpeningTine time.Time
    ClosingTime time.Time
}

type ScheduleResource struct {
    ID int
    TrainID int
    StationID int
    Arrivalttime time.Time
}

func (t *TrainResource) Register(container *restful.Container){
    ws := new(restful.WebService)
    ws.Path("/v1/trains").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
    ws.Route(ws.GET("/{train_id}").To(t.getTrain))
    ws.Route(ws.POST("").To(t.createTrain))
    ws.Route(ws.DELETE("/{train_id}").To(t.removeTrain))
    container.Add(ws)
}

func (t *TrainResource) getTrain(request *restful.Request, response *restful.Response){
    id := request.PathParameter("train_id")
    err := DB.QueryRow("SELECT ID, DRIVER_NAME, OPERATING_STATUS FROM train WHERE id=?", id).Scan(&t.ID, &t.DriverName, &t.OperatingStatus)
    if err != nil {
        log.Println(err)
        response.AddHeader("Content-Type","text/plain")
        response.WriteErrorString(http.StatusNotFound,"Train could not be found")
    }else{
        response.WriteEntity(t)
    }
}

func (t *TrainResource) createTrain(request *restful.Request, response *restful.Response){
    log.Println(request.Request.Body)
    decoder := json.NewDecoder(request.Request.Body)
    var b TrainResource
    err := decoder.Decode(&b)
    statement, _ := DB.Prepare("INSERT INTO train (DRIVER_NAME, OPERATING_STATUS) VALUES (?,?)")
    result, err := statement.Exec(b.DriverName,b.OperatingStatus)
    
    if err == nil {
        newID, _ := result.LastInsertId()
        b.ID = int(newID)
        response.WriteHeaderAndEntity(http.StatusCreated, b)
    }else{
        response.AddHeader("Content-Type", "text-plain")
        response.WriteErrorString(http.StatusInternalServerError, err.Error())
    }
}

func (t *TrainResource) removeTrain(request *restful.Request, response *restful.Response){
    id := request.PathParameter("train_id")
    statement, _ := DB.Prepare("DELETE FROM train WHERE id=?")
    _, err := statement.Exec(id)

    if err == nil {
        response.WriteHeader(http.StatusOK)
    }else{
        response.WriteErrorString(http.StatusInternalServerError, err.Error())
    }

}

func main(){
    var err error
    DB, err = sql.Open("sqlite3", "./railapi.db")
    if err != nil {
        log.Println("Driver creation failed")
    }
    dbutils.Initialize(DB)

    ws := restful.NewContainer()
    ws.Router(restful.CurlyRouter{})
    t := TrainResource{}
    t.Register(ws)
    log.Printf("Start listening on lcalhost:8000")
    server := &http.Server{Addr: ":8000", Handler: ws}
    log.Fatal(server.ListenAndServe())
}
