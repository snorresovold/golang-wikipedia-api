package main

/*
func main() {
	url := "The_Batman_(film)"
	check_site(url)
	router := gin.Default()
	router.GET("/"+url, getArticle)
	router.Run("localhost:8080")
}

func check_site(x string) { // x is a url to the site you want to scrape
	c := colly.NewCollector()

	fName := "data.csv"
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
	url := c.Request.URL.Query()
	fmt.Println("sus", url)
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
*/
