package main

import (
    "context"
    proto "github.com/tawiparkP/encrypt/encryption"
    "github.com/tawiparkP/encrypt/utils"
)

type Encrypter struct{}

func (g *Encrypter) Encrypt(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
    rsp.Result = utils.EncryptString(req.Key, req.Message)    
    return nil
}

func (g *Encrypter) Decrypt(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
    rsp.Result = utils.DecryptString(req.Key, req.Message)    
    return nil
}