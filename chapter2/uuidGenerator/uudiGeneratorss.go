package main

import (
	"fmt"
    "crypto/rand"
    "net/http"
)

type Uuid struct {}

func (p *Uuid) ServeHTTP(w http.ResponseWriter, r *http.Request){
    if r.URL.Path == "/"{
        giverandomUUID(w,r)
        return
    }
    http.NotFound(w,r)
    return
}

func giverandomUUID(w http.ResponseWriter, r *http.Request){
    c := 10
    b := make([]byte, c)
    _, err := rand.Read(b)
    if err != nil {
        panic(err)
    }
    fmt.Fprintf(w, fmt.Sprintf("%x", b))
}