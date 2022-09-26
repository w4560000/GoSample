package main

import (
	"GOSAMPLE/Syntax/bxtest"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": string("123"), // cast it to string before showing
		})
	})

	router.Run(":7000")
}

func a() int {
	return bxtest.Add(1, 2)
}
