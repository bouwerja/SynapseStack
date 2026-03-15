package forums

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

type Post struct {
	Author  string
	Message string
	Date    string
}

type Thread struct {
	Title string
	URL   string
	Posts []Post
}

func ScrapBW() {
	c := colly.NewCollector(
		colly.AllowedDomains("bizwarriors.com", "www.bizwarriors.com"),
		colly.Async(false),
	)

	c.SetRequestTimeout(60 * time.Second)

	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		DisableKeepAlives:     true,
	})

	var mu sync.Mutex
	results := make(map[string]*Thread)

	tc := setupThreadCollector(c, &mu, results)
	sc := catCollector(c, tc)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
		r.Headers.Set("Accept-Language", "en-US,en;q=0.9")
		r.Headers.Set("Sec-Ch-Ua", `"Chromium";v="122", "Not(A:Brand";v="24", "Google Chrome";v="122"`)
		r.Headers.Set("Sec-Ch-Ua-Mobile", "?0")
		r.Headers.Set("Sec-Ch-Ua-Platform", `"Windows"`)
		r.Headers.Set("Upgrade-Insecure-Requests", "1")

		fmt.Println("Visiting", r.URL.String())
	})

	rule := &colly.LimitRule{
		Parallelism: 1,
		RandomDelay: 2 * time.Second,
		Delay:       2 * time.Second,
	}

	c.Limit(rule)
	sc.Limit(rule)
	tc.Limit(rule)

	c.OnHTML("h3.node-title", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		absLink := e.Request.AbsoluteURL(link)
		fmt.Printf("Visiting %s\n", absLink)
		if err := sc.Visit(absLink); err != nil {
			fmt.Printf("Skipping Category (likely already visited): %v\n", err)
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error on %s: %v (Status: %d)\n", r.Request.URL, err, r.StatusCode)
	})

	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode != 200 {
			fmt.Printf("[Main] Received Status: %d for URL: %s\n", r.StatusCode, r.Request.URL)
		}
	})

	sc.OnResponse(func(r *colly.Response) {
		if r.StatusCode != 200 {
			fmt.Printf("[Sub] Received Status: %d for URL: %s\n", r.StatusCode, r.Request.URL)
		}
	})

	tc.OnResponse(func(r *colly.Response) {
		if r.StatusCode != 200 {
			fmt.Printf("[Thread] Received Status: %d for URL: %s\n", r.StatusCode, r.Request.URL)
		}
	})

	if err := c.Visit("https://bizwarriors.com/forum/"); err != nil {
		log.Fatalf("Error visiting link: %v\n", err)
	}
	c.Wait()
	sc.Wait()
	tc.Wait()

	file, _ := json.MarshalIndent(results, "", "  ")
	_ = os.WriteFile("scrapper/JSONdata/bizwarriors_data.json", file, 0o644)

	fmt.Printf("Scraped %d threads total\n", len(results))
}

func catCollector(mainC *colly.Collector, threadC *colly.Collector) *colly.Collector {
	sc := mainC.Clone()

	sc.OnHTML("div.structItem-title", func(e *colly.HTMLElement) {
		threadLink := e.Request.AbsoluteURL(e.ChildAttr("a", "href"))
		fmt.Printf("    -> Visiting %s\n", threadLink)
		if err := threadC.Visit(threadLink); err != nil {
			fmt.Printf("Skipping Thread: %v\n", err)
		}
	})

	sc.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Sub Collector Error %s: %v (Status: %d)\n", r.Request.URL, err, r.StatusCode)
	})
	return sc
}

func setupThreadCollector(mainC *colly.Collector, mu *sync.Mutex, results map[string]*Thread) *colly.Collector {
	tc := mainC.Clone()

	tc.OnHTML("article.message", func(e *colly.HTMLElement) {
		threadURL := e.Request.URL.String()
		fmt.Printf("    -> Requesting from %s\n", threadURL)

		post := Post{
			Author:  e.ChildAttr("article.message", "data-author"),
			Message: e.ChildText("div.bbWrapper"),
			Date:    e.ChildText("time.u-dt"),
		}

		mu.Lock()
		if _, exists := results[threadURL]; !exists {
			results[threadURL] = &Thread{
				Title: e.ChildText("h1.p-title-value"),
				URL:   threadURL,
			}
		}
		results[threadURL].Posts = append(results[threadURL].Posts, post)
		mu.Unlock()
	})

	tc.OnError(func(r *colly.Response, err error) {
		if r.StatusCode == 502 || r.StatusCode == 0 {
			fmt.Printf("Ghosted or 502 on %s. Backing off 20s and retrying...\n", r.Request.URL)
			time.Sleep(20 * time.Second)
			err := r.Request.Retry()
			if err != nil {
				fmt.Printf("Error happend on Retry of Thread Collector: %v\n", err)
			}
		} else {
			fmt.Printf("Thread Collector Error %s: %v (Status: %d)\n", r.Request.URL, err, r.StatusCode)
		}
	})

	return tc
}
