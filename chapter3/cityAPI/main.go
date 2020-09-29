package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type city struct{
    Name string
    Area uint64 
}

func postHandler(w http.ResponseWriter, r *http.Request){
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
    http.HandleFunc("/city", postHandler)
    http.ListenAndServe(":8000", nil)
}