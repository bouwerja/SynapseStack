package scrapper

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
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

	fName := "scrapper/JSONdata/business.json"

	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
	}
	defer file.Close()

	var businesses []Business

	c := colly.NewCollector(
		colly.AllowedDomains("brabys.com", "www.brabys.com"),
		colly.Async(true),
	)

	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2: true,
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
		r.Headers.Set("Accept-Language", "en-US,en;q=0.9")
		r.Headers.Set("Referer", "https://www.google.com/")
		r.Headers.Set("Sec-Ch-Ua", `"Chromium";v="122", "Not(A:Brand";v="24", "Google Chrome";v="122"`)
		r.Headers.Set("Sec-Fetch-Mode", "navigate")
	})

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*brabys.com",
		Parallelism: 5,
		RandomDelay: 1 * time.Second,
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("ERROR on %s: Status %d | %v\n", r.Request.URL, r.StatusCode, err)
	})

	c.OnHTML("#category-verified-business a.sub-category", func(e *colly.HTMLElement) {
		categoryName := strings.TrimSpace(e.Text)
		link := e.Request.AbsoluteURL(e.Attr("href"))
		// link := e.Attr("href")
		// fmt.Printf("Found Category: %s\n", e.Text)

		ctx := colly.NewContext()
		ctx.Put("category", categoryName)

		e.Request.Visit(link)
	})

	c.OnHTML("div.grid_element", func(e *colly.HTMLElement) {
		bizName := e.ChildAttr("a", "title")
		bizName = strings.Split(bizName, " - Business")[0]

		rawLocation := e.ChildText("span.member-search-location small")
		cleanLocation := strings.Join(strings.Fields(rawLocation), " ")

		category := e.Response.Ctx.Get("category")

		item := Business{
			Category:    category,
			CompanyName: bizName,
			Location:    cleanLocation,
		}

		mu.Lock()
		businesses = append(businesses, item)
		mu.Unlock()
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visting: ", r.URL.String())
	})

	c.Visit("https://www.brabys.com/categories/")
	c.Wait()

	writeJSON(fName, businesses)
}

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
