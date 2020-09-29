package main

import (
	"path/filepath"
	"os"
	"net/http"
	"log"
	"io/ioutil"
    jsonparse "encoding/json"

    "github.com/gorilla/mux"
    "github.com/gorilla/rpc"
    "github.com/gorilla/rpc/json"
)

type Args struct{
    ID string
}

type Book struct{
    ID string
    Name string
    Author string
}

type JSONServer struct{}

func (t *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error{
    var books []Book
    absPath, err := filepath.Abs("books.json")
    if err != nil {
        log.Println("file path error:", err)
        os.Exit(1)
    }

    raw, readerr := ioutil.ReadFile(absPath)
    if readerr != nil {
        log.Println("file read error:", err)
        os.Exit(1)
    }

    marshalerr := jsonparse.Unmarshal(raw, &books)
    if marshalerr != nil {
        log.Println("json error:", marshalerr)
        os.Exit(1)
    }

    for _, book := range books{
        if book.ID == args.ID{
            *reply = book
            break
        }
    }
    return nil
}

func main(){
    s := rpc.NewServer()
    s.RegisterCodec(json.NewCodec(), "application/json")
    s.RegisterService(new(JSONServer), "")
    r := mux.NewRouter()
    r.Handle("/rpc",s)
    http.ListenAndServe(":1234",r)
}