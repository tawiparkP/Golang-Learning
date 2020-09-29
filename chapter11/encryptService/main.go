package main

import (
    "fmt"
    proto "github.com/tawiparkP/encrypt/encryption"
    micro "github.com/micro/go-micro"
)

func main(){
    service := micro.NewService(micro.Name("encrypter"))
    service.Init()
    proto.RegisterEncrypterHandler(service.Server(), new(Encrypter))

    if err := service.Run(); err != nil {
        fmt.Println(err)
    }
}