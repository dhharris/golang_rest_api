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

	mySqlOptions := data.MysqlOptions{
		DbName:     "sybo",
		DbUser:     "root",
		DbPassword: "password",
	}

	storage := data.MysqlStorageHandler{
		Driver: mySqlOptions.GetDriver(),
	}

	handler := handler.NewWebServiceHandler(
		handler.Options{
			Storage: storage,
		})

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
