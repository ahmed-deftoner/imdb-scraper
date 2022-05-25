package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type actor struct {
	name      string
	photo     string
	jobTitle  string
	bio       string
	birthDate string
	TopMovies []movies
}

type movies struct {
	Title string
	Year  string
}

func crawl(month int, day int) {
	c := colly.NewCollector(
		colly.AllowedDomains("imdb.com", "www.imdb.com"),
	)
	infoCollector := c.Clone()

	c.OnHTML(".mode-detail", func(e *colly.HTMLElement) {
		profileUrl := e.ChildAttr("div.lister-item-image > a", "href")
		profileUrl = e.Request.AbsoluteURL(profileUrl)
		infoCollector.Visit(profileUrl)
	})

	startUrl := fmt.Sprintf("https://www.imdb.com/search/name/?birth_monthday=%d-%d", month, day)
	c.Visit(startUrl)
}

func main() {
	crawl(1, 1)
}
