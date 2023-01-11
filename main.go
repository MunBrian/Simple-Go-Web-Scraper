package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// create a countyData struct
type countryData struct {
	CountryName string
	CapitalCity string
	Population  string
}

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

	c.OnHTML("div.col-md-4.country", func(h *colly.HTMLElement) {

		countryData := countryData{
			CountryName: h.ChildText("h3.country-name"),
			CapitalCity: h.ChildText("span.country-capital"),
			Population:  h.ChildText("span.country-population"),
		}

		fmt.Println(countryData)
	})

	c.Visit("https://www.scrapethissite.com/pages/simple/")

}
