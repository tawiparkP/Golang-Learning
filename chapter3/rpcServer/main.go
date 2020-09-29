package main

import (
	"net/rpc"
	"net/http"
	"net"
	"log"
	"time"
)


type Args struct{}

type TimeServer int64

func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error{
    *reply = time.Now().Unix()
    return nil
}

func main(){
    timeServer := new(TimeServer)
    rpc.Register(timeServer)
    rpc.HandleHTTP()
    l, e := net.Listen("tcp", ":12345")
    if e != nil {
        log.Fatal("Listen error:", e)
    }
    log.Fatal(http.Serve(l,nil))
}