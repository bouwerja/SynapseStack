package forums

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type RedditResp struct {
	Kind string `json:"kind"`
	Data struct {
		After     string `json:"after"`
		Dist      int    `json:"dist"`
		ModHash   string `json:"modhash"`
		GeoFilter string `json:"geo_filter"`
		Children  []struct {
			Kind      string `json:"kind"`
			ChildData struct {
				Title     string `json:"title"`
				PermaLink string `json:"permalink"`
			} `json:"data"`
		} `json:"children"`
		Before string `json:"before"`
	} `json:"data"`
}

func RedditResponse() {
	transport := &http.Transport{
		ForceAttemptHTTP2: true,
	}
	client := &http.Client{Transport: transport}

	req, err := http.NewRequest("GET", "https://www.reddit.com/r/startups/rising.json", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:148.0) Gecko/20100101 Firefox/148.0")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Sec-GPC", "1")

	response, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			panic(err)
		}
	}()

	if response.StatusCode != http.StatusOK {
		log.Fatal(response.StatusCode)
	}

	respData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var respObject RedditResp
	if err := json.Unmarshal(respData, &respObject); err != nil {
		fmt.Println("Error is here")
	}

	fmt.Printf("Found %d posts:\n", len(respObject.Data.Children))
	for i, child := range respObject.Data.Children {
		fmt.Printf("%d, %s\n", i+1, child.ChildData.Title)
		fmt.Printf(" Link: https://reddit.com%s\n", child.ChildData.PermaLink)
	}
}
