package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func main() {
	collyObj := colly.NewCollector(
		colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	)

	collyObj.OnHTML("a[href", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println("Link found: %q -> %s\n", e.Text, link)

		collyObj.Visit(e.Request.AbsoluteURL(link))
	})

	collyObj.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	collyObj.Visit("https://hackerspaces.org/")
}
