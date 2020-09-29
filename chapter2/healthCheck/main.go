package main

import (
	"time"
	"net/http"
	"log"
	"io"
)

func HealthCheck(w http.ResponseWriter,req *http.Request){
    currentTime := time.Now()
    io.WriteString(w, currentTime.String())
}

func main(){
    http.HandleFunc("/health", HealthCheck)
    log.Fatal(http.ListenAndServe(":8000",nil))
}