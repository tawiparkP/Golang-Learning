package main

import (
    "log"
    "github.com/tawiparkP/encrypt/utils"
    "encoding/hex"
)

func main(){
    key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
    message := "I am A Message"
    log.Println("Original message: ", message)
    encryptedString := utils.EncryptString(key, message)
    log.Println("Encrypted message: ",encryptedString)
    decryptedString := utils.DecryptString(key, encryptedString)
    log.Println("Decrypted message: ",decryptedString)   
}

