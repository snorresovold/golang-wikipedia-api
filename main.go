package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func getArticle(c *gin.Context) {
	csvFile, err := os.Open("data.csv")
	if err != nil {
		fmt.Println(err)
	}
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	c.IndentedJSON(http.StatusOK, csvLines)
	defer csvFile.Close()
}

func main() {
	url := "The_Batman_(film)"
	check_site(url)
	router := gin.Default()
	router.GET("/"+url, getArticle)
	router.Run("localhost:8080")
}
