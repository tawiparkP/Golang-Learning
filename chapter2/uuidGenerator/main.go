package main

import (
    "net/http"
)

func main(){
    mux := &Uuid{}
    http.ListenAndServe(":8000",mux)
}