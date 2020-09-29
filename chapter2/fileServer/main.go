package main

import (
	"net/http"
    "log"
    
    "github.com/julienschmidt/httprouter"
)

func main(){
    r := httprouter.New()
    r.ServeFiles("/static/*filepath", http.Dir("static"))
    log.Fatal(http.ListenAndServe(":8000",r))
}

