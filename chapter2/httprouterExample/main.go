package main

import (
	"os/exec"
	"net/http"
	"log"
	"io"
    "fmt"
    "github.com/julienschmidt/httprouter"
)

func goVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params){
    resp := getCommandOutput("go","version")
    io.WriteString(w, resp)
    return
}

func getFileContent(w http.ResponseWriter, r *http.Request, params httprouter.Params){
    fmt.Println(params.ByName("name"))
    data := getCommandOutput("cat",params.ByName("name"))
    fmt.Fprintf(w, data)
}

func getCommandOutput(command string, argu ...string) string{
    out,err := exec.Command(command, argu...).Output()
    if err != nil{
        fmt.Println(err)
    }
    return string(out)
}

func main(){
    router := httprouter.New()
    router.GET("/api/v1/go-version",goVersion)
    router.GET("/api/v1/show-file/:name",getFileContent)
    log.Fatal(http.ListenAndServe(":8000",router))
}

