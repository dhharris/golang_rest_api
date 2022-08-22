package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sybo/web-service/data"
	"sybo/web-service/handler"
)

func main() {
	// TODO: remove
	fmt.Println("Hello world!")

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
