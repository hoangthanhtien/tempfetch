package utils

import (
	"encoding/json"
	"fmt"
	"github.com/guptarohit/asciigraph"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Location struct {
	Long     float64
	Lat      float64
	Timezone string
}

type IPRepsonse struct {
	Loc      string `json:"loc"`
	Timezone string `json:"timezone"`
}

type ForecastParam struct {
	Long     float64
	Lat      float64
	Hourly   string
	Timezone string `json:"timezone"`
}

type ForecastReponse struct {
	Lat                  float64          `json:"latitude"`
	Long                 float64          `json:"longitude"`
	GenerationTimeMsf    float64          `json:"generationtime_ms"`
	UTCOffsetSeconds     int              `json:"utc_offset_seconds"`
	Timezone             string           `json:"timezone"`
	TimezoneAbbreviation string           `json:"timezone_abbreviation"`
	Elevation            float64          `json:"elevation"`
	Hourly               HourlyTimeSeries `json:"hourly"`
}
type HourlyTimeSeries struct {
	Time          []string
	Temperature2M []float64 `json:"temperature_2m"`
}

func GetCurrentUserInfo() (Location, error) {
	resp, err := http.Get("https://ipinfo.io/json")
	var ipResponse IPRepsonse
	var loc Location
	if err != nil {
		fmt.Println("Error:", err)
		return loc, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&ipResponse); err != nil {
		fmt.Println("Error decoding response:", err)
		return loc, err
	}
	parts := strings.Split(ipResponse.Loc, ",")

	longStr, latStr := parts[1], parts[0]

	long, longErr := strconv.ParseFloat(longStr, 64)
	lat, latErr := strconv.ParseFloat(latStr, 64)

	if longErr != nil || latErr != nil {
		log.Fatal("Can't get long, lat of current user")
	}

	loc = Location{
		Long:     long,
		Lat:      lat,
		Timezone: ipResponse.Timezone,
	}

	return loc, err

}

func GetForecast(params ForecastParam) (ForecastReponse, error) {
	url := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%v&longitude=%v&hourly=%v&timezone=%v",
		params.Lat,
		params.Long,
		params.Hourly,
		params.Timezone)
	// fmt.Printf("url: %v", url)

	var result ForecastReponse
	resp, err := http.Get(url)
	if err != nil {
		return result, err
	}
	// fmt.Printf("Data from open-meteo: %+v\n", resp)
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("Error getting forecast from open-meteo api")
		return result, err
	}

	return result, err

}

func RenderASCIIChart(data ForecastReponse) {
	graph := asciigraph.Plot(data.Hourly.Temperature2M)
	fmt.Println(graph)
	fmt.Printf("%-18s", " ") // Padding
	for _, label := range data.Hourly.Time {
		if label[11:16] == "12:00" {
			fmt.Printf("%-24s", label[:10])
		}
	}
	fmt.Println()
}
