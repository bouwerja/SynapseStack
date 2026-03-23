package main

import (
	"fmt"
	"time"

	// f "backend/scrapper/Forums"
	bra "backend/scrapper/Registers"
)

func main() {
	start := time.Now()

	fmt.Println("Initializing Scraper...")

	bra.BrabysScrapper()

	duration := time.Since(start)
	fmt.Printf("\nScraping completed in: %v\n", duration)
}
