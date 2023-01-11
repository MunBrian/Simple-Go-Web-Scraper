package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

// create a countyData struct
type countryData struct {
	CountryName string
	CapitalCity string
	Population  string
}

// define a slice to countryData struct type
var countryDataSlice = make([]countryData, 0)

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

		countryDataSlice = append(countryDataSlice, countryData)
	})

	c.Visit("https://www.scrapethissite.com/pages/simple/")

	//pass slice to json Marshall to change data to json
	content, err := json.Marshal(countryDataSlice)

	//handle error
	if err != nil {
		fmt.Println(err.Error())
	}

	//write a json file pass name, data and permission
	os.WriteFile("country-data.json", content, 0644)

}
