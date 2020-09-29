package main

import (
	"net/rpc"
	"log"
)

type Args struct{}

func main(){
    var reply int64
    args := Args{}
    client, err := rpc.DialHTTP("tcp", "localhost"+":12345")
    if err != nil {
        log.Fatal("dialing:", err)
    }

    err = client.Call("TimeServer.GiveServerTime", args, &reply)
    if err != nil {
        log.Fatal("arith error:", err)
    }

    log.Printf("%d", reply)
}