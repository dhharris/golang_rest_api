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

	router.GET("/user", func(c *gin.Context) {
		resp := make(map[string][]data.User)
		resp["users"] = handler.GetAllUsers()
		c.IndentedJSON(http.StatusOK, resp)
	})

	router.PUT("/user/:id/state", func(c *gin.Context) {
		id := c.Param("id")
		var state data.State
		if err := c.BindJSON(&state); err != nil {
			return
		}

		handler.SaveState(id, state)
		c.IndentedJSON(http.StatusOK, "Game saved")
	})

	router.GET("/user/:id/state", func(c *gin.Context) {
		id := c.Param("id")
		state, err := handler.LoadState(id)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, err)
		} else {
			c.IndentedJSON(http.StatusOK, state)
		}
	})

	router.PUT("/user/:id/friends", func(c *gin.Context) {
		id := c.Param("id")
		var friends data.Friends
		if err := c.BindJSON(&friends); err != nil {
			return
		}

		handler.UpdateFriends(id, friends)
		c.IndentedJSON(http.StatusOK, "Friends updated")
	})

	router.GET("/user/:id/friends", func(c *gin.Context) {
		id := c.Param("id")
		friends, err := handler.GetFriends(id)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, err)
		} else {
			c.IndentedJSON(http.StatusOK, friends)
		}
	})

	router.Run(":8080")
}
