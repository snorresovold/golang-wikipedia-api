package scraping

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	check_site("The_Batman_(film)")
}

func check_site(x string) { // x is a url to the site you want to scrape
	c := colly.NewCollector()
	// On every a element which has href attribute call callback
	c.OnHTML("p", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	// before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.Visit("https://en.wikipedia.org/wiki/" + x) // checks x aka the URL
}
