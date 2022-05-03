package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getArticle(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "hello")
}

func main() {

	check_site("The_Batman_(film)")
	router := gin.Default()
	router.GET("/articles", getArticle)
	router.Run("localhost:8080")
}
