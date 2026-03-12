package scrapper

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	// "github.com/gocolly/colly/v2"
)

type Business struct {
	Cateogry    string
	CompanyName string
	Location    string
}

func BusinessRegisterScrapper() {
	/*
		fName := "/JSONdata/business.json"

		file, err := os.Create(fName)
		if err != nil {
			log.Fatalf("Cannot create file %q: %s\n", fName, err)
		}
		defer file.Close()
	*/

	c := colly.NewCollector(
		colly.AllowedDomains("brabys.com", "www.brabys.com"),
	)

	c.OnHTML("#category-verified-business a.sub-category", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Found Category: %s\n", e.Text)

		e.Request.Visit(link)
	})

	c.OnHTML("div.grid_element", func(e *colly.HTMLElement) {
		bizName := e.ChildAttr("a", "title")
		bizName = strings.Split(bizName, " - Business")[0]
		fmt.Printf("---- Scraping Business: %s ----\n", bizName)

		/*
			locationContainer := e.DOM.Find("span.member-search-location small")
			bizArea := locationContainer.Contents().First().Text()
			bizArea = strings.Split(strings.TrimSpace(bizArea), ",")[0]
			bizRegion := e.ChildText("span.inline-block")
			fmt.Println("Business Location: ")
			fmt.Printf("%s %s\n", bizArea, bizRegion)
		*/

		rawLocation := e.ChildText("span.member-search-location small")
		cleanLocation := strings.Join(strings.Fields(rawLocation), " ")
		fmt.Println("Business Locations: ")
		fmt.Println(cleanLocation)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visting: ", r.URL.String())
	})

	c.Visit("https://www.brabys.com/categories/")
}
