package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getArticle(c *gin.Context) {

	c.IndentedJSON(http.StatusOK)
}

func main() {
	router := gin.Default()
	scraper.check_site()

}
