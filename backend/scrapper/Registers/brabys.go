package registers

import (
	"fmt"
	"strings"
	"sync"
	"time"

	// "github.com/gocolly/colly"
	"github.com/gocolly/colly/v2"
)

type Business struct {
	Category    string `json:"operating_sector"`
	CompanyName string `json:"company_name"`
	Location    string `json:"company_location"`
}

func BusinessRegisterScrapper() {
	var mu sync.Mutex
	// fName := "scrapper/JSONdata/business.json"
	var businesses []Business

	c := colly.NewCollector(
		colly.AllowedDomains("brabys.com", "www.brabys.com"),
		colly.Async(true),
	)

	// Set up limits to avoid getting banned
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*brabys.com",
		Parallelism: 2, // Start slow
		RandomDelay: 2 * time.Second,
	})

	// 1. SCRAPE CATEGORIES
	c.OnHTML("#category-verified-business a.sub-category", func(e *colly.HTMLElement) {
		// categoryName := strings.TrimSpace(e.Text)
		// link := e.Request.AbsoluteURL(e.Attr("href"))

		// Create a new request and ATTACH the context to it
		// err := e.Request.Visit(link)
		// Note: Colly context is actually shared across requests in the same "branch"
		// Better: Put it in the request before visiting
	})

	// 2. SCRAPE BUSINESSES
	c.OnHTML("div.grid_element", func(e *colly.HTMLElement) {
		bizName := e.ChildAttr("a", "title")
		bizName = strings.Split(bizName, " - Business")[0]

		// If the category was on the previous page, we'd use e.Response.Ctx
		// But Brabys usually shows the category on the result page too.
		rawLocation := e.ChildText("span.member-search-location small")

		item := Business{
			CompanyName: strings.TrimSpace(bizName),
			Location:    strings.Join(strings.Fields(rawLocation), " "),
			Category:    "Trade/Service", // You can refine this by scraping the H1 of the page
		}

		mu.Lock()
		businesses = append(businesses, item)
		mu.Unlock()
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL.String())
	})

	c.Visit("https://www.brabys.com/categories/")
	c.Wait()

	// writeJSON(fName, businesses)
}

/*
func writeJSON(fileName string, data []Business) {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("error marshalling: %s", err)
	}

	err = os.WriteFile(fileName, file, 0o644)
	if err != nil {
		log.Fatalf("error writing file: ", err)
	}

	fmt.Printf("\nSuccessfully saved %d buinsesses to %s\n", len(data), fileName)
}
*/
