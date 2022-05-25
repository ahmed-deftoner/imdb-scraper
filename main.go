package main

import (
	"fmt"
	"strings"

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

	c.OnHTML("a.lister-page-next", func(e *colly.HTMLElement) {
		nextPage := e.Request.AbsoluteURL(e.Attr("href"))
		c.Visit(nextPage)
	})

	infoCollector.OnHTML("#content-2-wide", func(h *colly.HTMLElement) {
		topProfile := actor{}
		topProfile.name = h.ChildText("h1.header > span.itemprop")
		topProfile.photo = h.ChildAttr("#name-poster", "src")
		topProfile.jobTitle = h.ChildText("#name-job-categories > a > span.itemprop")
		topProfile.birthDate = h.ChildAttr("#name-born-info time", "datetime")
		topProfile.bio = strings.TrimSpace(h.ChildText("#name-bio-text > div.name-trivia-bio-text > div.inline"))

		h.ForEach("div.knownfor-title", func(_ int, ef *colly.HTMLElement) {
			tmpMovies := movies{}
			tmpMovies.Title = ef.ChildText("div.knownfor-title-role > a.knownfor-ellipsis")
			tmpMovies.Year = ef.ChildText("div.knownfor-year > span.knownfor-ellipsis")
			topProfile.TopMovies = append(topProfile.TopMovies, tmpMovies)
		})
	})

	startUrl := fmt.Sprintf("https://www.imdb.com/search/name/?birth_monthday=%d-%d", month, day)
	c.Visit(startUrl)
}

func main() {
	crawl(1, 1)
}
