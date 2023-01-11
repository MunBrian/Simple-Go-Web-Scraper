package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	//colly setup
	c := colly.NewCollector(colly.AllowedDomains("www.scrapethissite.com"))

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Scraping:", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "nError:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Status:", r.StatusCode)
	})

	c.Visit("https://www.scrapethissite.com/pages/simple/")

}
