package main

import (
	"net/http"
	"math/rand"
	"fmt"
)

func main(){
    newMux := http.NewServeMux()
    newMux.HandleFunc("/randomFloat",func(w http.ResponseWriter,r *http.Request){
        fmt.Fprintln(w, rand.Float64())
    })
    newMux.HandleFunc("/randomInt",func(w http.ResponseWriter,r *http.Request){
        fmt.Fprintln(w, rand.Intn(100))
    })
    http.ListenAndServe(":8000", newMux)
}