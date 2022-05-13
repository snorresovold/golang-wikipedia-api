package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func main() {
	router := gin.Default()

	router.GET("lol", func(c *gin.Context) {
		fmt.Println(c.FullPath())
	})
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
