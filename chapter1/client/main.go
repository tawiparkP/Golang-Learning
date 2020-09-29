package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main(){
    resp, err := http.Get("http://localhost:8585/fastest-mirror")
    if err != nil{
        fmt.Println("Err $q",err)
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
}
