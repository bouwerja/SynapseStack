package registers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	// "github.com/gocolly/colly"
	"github.com/gocolly/colly/v2"
)

type BrabysRepsonse struct {
	Data []Business
}

type Business struct {
	Category    string `json:"operating_sector"`
	CompanyName string `json:"company_name"`
	Location    string `json:"company_location"`
}

func BusinessRegisterScrapper() {
	var mu sync.Mutex
	fName := "scrapper/JSONdata/business.json"
	var businesses []Business

	c := colly.NewCollector(
		colly.AllowedDomains("brabys.com", "www.brabys.com"),
		colly.Async(false),
	)

	limitRules := &colly.LimitRule{
		DomainGlob:  "*brabys.com",
		Parallelism: 1,
		RandomDelay: 2 * time.Second,
		Delay:       2 * time.Second,
	}
	if err := c.Limit(limitRules); err != nil {
		log.Fatalf("Error in setting limit rules: %v\n", err)
	}

	c.OnHTML("#category-verified-business a.sub-category", func(e *colly.HTMLElement) {
		// categoryName := strings.TrimSpace(e.Text)
		link := e.Request.AbsoluteURL(e.Attr("href"))

		err := e.Request.Visit(link)
		if err != nil {
			log.Fatalf("Error in category-verified-business: %v\n", err)
		}
	})

	c.OnHTML("div.grid_element", func(e *colly.HTMLElement) {
		bizName := e.ChildAttr("a", "title")
		bizName = strings.Split(bizName, " - Business")[0]

		rawLocation := e.ChildText("span.member-search-location small")

		item := Business{
			CompanyName: strings.TrimSpace(bizName),
			Location:    strings.Join(strings.Fields(rawLocation), " "),
			Category:    "Trade/Service",
		}

		mu.Lock()
		businesses = append(businesses, item)
		mu.Unlock()
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL.String())
	})

	if err := c.Visit("https://www.brabys.com/categories/"); err != nil {
		log.Fatalf("Error visiting location: %v\n", err)
	}
	c.Wait()

	writeBrabysJSON(fName, businesses)
}

func writeBrabysJSON(fileName string, data []Business) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("error marshalling: %s", err)
	}

	err = os.WriteFile(fileName, jsonData, 0o644)
	if err != nil {
		log.Fatalf("error writing file: %s", err)
	}

	fmt.Printf("\nSuccessfully saved %d businesses to %s\n", len(data), fileName)
}
