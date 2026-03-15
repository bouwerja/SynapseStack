package main

import (
	"fmt"
	"time"

	reg "backend/scrapper/Registers"
)

func main() {
	start := time.Now()

	fmt.Println("Initializing High-Speed Scraper...")

	reg.Overpass()

	duration := time.Since(start)
	fmt.Printf("\nScraping completed in: %v\n", duration)
}
