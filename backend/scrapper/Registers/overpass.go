/* Package registers package used to get business list and location data
 */
package registers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

type OverpassResponse struct {
	Elements []struct {
		ID     int64   `json:"id"`
		Lat    float64 `json:"lat"`
		Lon    float64 `json:"lon"`
		Center struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"center"`
		Tags map[string]string `json:"tags"`
	} `json:"elements"`
}

func Overpass() {
	query := `
[out:json][timeout:180];

// 1. Get South Africa
area["name"="South Africa"]->.country;

// 2. Get Gauteng within South Africa
area["name"="Gauteng"](area.country)->.province;

// 3. Get Johannesburg within Gauteng 
// Note: We add admin_level=6 for the Metropolitan Municipality (the whole city)
area["name"="City of Johannesburg Metropolitan Municipality"]["admin_level"="6"](area.province)->.searchArea;

(
  nwr["craft"~"."](area.searchArea);
  nwr["shop"~"."](area.searchArea);
  nwr["office"~"."](area.searchArea);
  nwr["amenity"~"."](area.searchArea);
);
out center;
	`

	apiEndpoint := "https://overpass-api.de/api/interpreter"
	data := url.Values{}
	data.Set("data", query)

	resp, err := http.PostForm(apiEndpoint, data)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatalf("Error closing response body: %s\n", err)
		}
	}()

	var result OverpassResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		panic(err)
	}

	filteredElements := result.Elements[:0]
	for _, el := range result.Elements {
		if name, ok := el.Tags["name"]; ok && name != "" {
			filteredElements = append(filteredElements, el)
		}
	}
	result.Elements = filteredElements
	/*
		for _, el := range result.Elements {
			category := "Business"
			keys := []string{"craft", "office", "shop", "amenity"}
			for _, k := range keys {
				if val, ok := el.Tags[k]; ok {
					category = val
					break
				}
			}

			lat, lon := el.Lat, el.Lon
			if lat == 0 {
				lat = el.Center.Lat
				lon = el.Center.Lon
			}

			fmt.Printf("▶ %s [%s]\n", el.Tags["name"], category)
			fmt.Printf("  Location: %f, %f\n", lat, lon)
			fmt.Println("-------------------------------------------")
		}
	*/

	fmt.Printf("Found %d businesses\n", len(result.Elements))

	fName := "scrapper/JSONdata/overpass.json"
	writeJSON(fName, result)
}

func writeJSON(fileName string, data OverpassResponse) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("error marshalling: %s", err)
	}

	err = os.WriteFile(fileName, jsonData, 0o644)
	if err != nil {
		log.Fatalf("error writing file: %s", err)
	}

	fmt.Printf("\nSuccessfully saved %d businesses to %s\n", len(data.Elements), fileName)
}
