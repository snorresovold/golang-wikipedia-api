package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func main() {
	url := "The_Batman_(film)"
	check_site(url)
	router := gin.Default()
	router.GET("/"+url, getArticle)
	router.GET("/Foo", func(c *gin.Context) {
		fmt.Println("the url is:", c.Request.Host+c.Request.URL.Path)
	})
	fmt.Println(url)
	router.Run("localhost:8080")
}

func check_site(x string) { // x is a url to the site you want to scrape
	c := colly.NewCollector()

	fName := x + ".csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file, err :%q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	// On every a element which has href attribute call callback
	c.OnHTML("p", func(e *colly.HTMLElement) {
		//fmt.Println(e.Text)
		writer.Write([]string{
			e.Text,
		})
	})
	// before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.Visit("https://en.wikipedia.org/wiki/" + x) // checks x aka the URL
}
func getArticle(c *gin.Context) {
	csvFile, err := os.Open("data" + "csv")
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
