package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	check_site("The_Batman_(film)")
}

func check_site(x string) {
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

	c.OnHTML("p", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.Text,
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://en.wikipedia.org/wiki/" + x)

}
