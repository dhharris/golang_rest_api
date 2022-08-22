package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    "sybo/web-service/handler"
    "sybo/web-service/data"
)

func main() {
    // TODO: remove
    fmt.Println("Hello world!"); 

    options := handler.Options{
        Storage: data.NewMockStorageHandler(),
    }
    handler := handler.NewWebServiceHandler(options)

    router := gin.Default()

    router.POST("/user", func(c *gin.Context) {
        var request data.NewUserRequest
        if err := c.BindJSON(&request); err != nil {
            return
        }
        c.IndentedJSON(http.StatusOK, handler.CreateUser(request))
    })

    router.Run(":8080")
}
