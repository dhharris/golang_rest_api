package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    "sybo/web-service/handler"
    "sybo/web-service/data"
)

func main() {
    fmt.Println("Hello world!"); 
    handler := new(handler.WebServiceHandler)
}
