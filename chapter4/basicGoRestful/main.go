package main

import (
	"time"
	"net/http"
	"io"
    "fmt"
    "github.com/emicklei/go-restful"
)

func pingTime(req *restful.Request, resp *restful.Response){
    io.WriteString(resp, fmt.Sprintf("%s",time.Now()))
}

func main(){
    ws := new(restful.WebService)
    ws.Route(ws.GET("/ping").To(pingTime))
    restful.Add(ws)
    http.ListenAndServe(":8000",nil)
}

