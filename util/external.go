package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func GetPlaceDetails(placeId string) ([]string, error) {
	mapApiKey := os.Getenv("MAP_API_KEY")
	placeDetailsApiUrl := os.Getenv("PLACE_DETAILS_API_URL")

	url := fmt.Sprintf("%s?placeid=%s&key=%s", placeDetailsApiUrl, placeId, mapApiKey)
	res, _ := http.Get(url)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)

	var result map[string]interface{}
	err := json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	photos := result["result"].(map[string]interface{})["photos"].([]interface{})

	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	var urls []string
	for _, photo := range photos {
		ref := photo.(map[string]interface{})["photo_reference"].(string)
		photoURL := fmt.Sprintf("https://maps.googleapis.com/maps/api/place/photo?maxwidth=800&photo_reference=%s&key=%s", ref, mapApiKey)
		resp, err := client.Head(photoURL)
		if err != nil {
			return nil, err
		}
		finalURL := resp.Request.URL.String()
		urls = append(urls, finalURL)
	}

	return urls, nil
}
