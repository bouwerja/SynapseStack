package main

import (
	"fmt"
	"time"

	f "backend/scrapper/Forums"
)

func main() {
	start := time.Now()

	fmt.Println("Initializing Scraper...")

	f.ScrapBW()

	duration := time.Since(start)
	fmt.Printf("\nScraping completed in: %v\n", duration)
}
