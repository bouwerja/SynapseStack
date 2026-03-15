// Package forums
// Scraps forums like reddit for business startegy ideas
package forums

import (
	"fmt"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

type Item struct {
	Title     string
	StoryURL  string
	Source    string
	Comments  string
	CrawledAt time.Time
}

func scrapReddit(targetURL string) {
	stories := []Item{}
	var mu sync.Mutex

	c := colly.NewCollector(
		colly.AllowedDomains("old.reddit.com", "www.old.reddit.com", "reddit.com"),
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
		r.Headers.Set("Accept-Language", "en-US,en;q=0.9")
		r.Headers.Set("Sec-Ch-Ua", `"Chromium";v="122", "Not(A:Brand";v="24", "Google Chrome";v="122"`)
		r.Headers.Set("Sec-Ch-Ua-Mobile", "?0")
		r.Headers.Set("Sec-Ch-Ua-Platform", `"Windows"`)
		r.Headers.Set("Upgrade-Insecure-Requests", "1")

		fmt.Println("Visiting", r.URL.String())
	})

	c.OnHTML("div.top-matter", func(e *colly.HTMLElement) {
		temp := Item{
			// Target the <a> tag with the title attribute specifically
			Title:     e.ChildText("a[data-event-action=title]"),
			StoryURL:  e.Request.AbsoluteURL(e.ChildAttr("a[data-event-action=title]", "href")),
			Comments:  e.ChildAttr("a[data-event-action=comments]", "href"),
			Source:    e.Request.URL.String(),
			CrawledAt: time.Now(),
		}

		// Validation: Don't append empty results
		if temp.Title != "" {
			mu.Lock()
			stories = append(stories, temp)
			mu.Unlock()
		}
	})

	c.OnHTML("span.next-button a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		e.Request.Visit(link)
	})

	// Set max Parallelism and introduce a Random Delay
	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error on %s: %v (Status: %d)\n", r.Request.URL, err, r.StatusCode)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Crawl all reddits the user passes in
	c.Visit(targetURL)

	c.Wait()

	fmt.Printf("Successfully scraped %d stories\n", len(stories))
}
