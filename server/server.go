package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld() {
	fmt.Println("Hello World!")
}

func Ping() {
	router := gin.Default()
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

}
