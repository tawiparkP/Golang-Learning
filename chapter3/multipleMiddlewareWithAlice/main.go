package main

import (
	"fmt"
	"time"
	"strconv"
	"net/http"
	"log"
    "encoding/json"
    "github.com/justinas/alice"
)

type city struct{
    Name string
    Area uint64 
}

func filterContentType(handler http.Handler) http.Handler{
    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
        log.Println("Currently in the check content type middleware")
        if r.Header.Get("Content-type") != "application/json"{
            w.WriteHeader(http.StatusUnsupportedMediaType)
            w.Write([]byte("415 - Unsupported Media Type. Please send JSON"))
            return
        }
        handler.ServeHTTP(w, r)
    })
}

func setServerTimeCookie(handler http.Handler) http.Handler{
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        cookie := http.Cookie{Name: "Server-Time(UTC)", Value: strconv.FormatInt(time.Now().Unix(), 10)}
        http.SetCookie(w, &cookie)
        log.Println("Currently in the set server time middleware")
        handler.ServeHTTP(w, r)
    })
}

func handle(w http.ResponseWriter, r *http.Request){
    if r.Method == "POST"{
        var tempCity city
        decoder := json.NewDecoder(r.Body)
        err := decoder.Decode(&tempCity)
        if err != nil{
            fmt.Println(err)
            panic(err)
        }
        defer r.Body.Close()
        fmt.Printf("Got %s city with area of %d sq miles!\n", tempCity.Name, tempCity.Area)
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("201 - Created"))
    }else{
        w.WriteHeader(http.StatusMethodNotAllowed)
        w.Write([]byte("405 - Method not allowed"))
    }
}

func main(){
    originalhandler := http.HandlerFunc(handle)
    chain := alice.New(filterContentType, setServerTimeCookie).Then(originalhandler)

    http.Handle("/city", chain)
    http.ListenAndServe(":8000",nil)
}
