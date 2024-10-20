package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

type WeatherResponse struct {
	Main struct {
		Temp       float64 `json:"temp"`
		Pressure   int     `json:"pressure"`
		Humidity   int     `json:"humidity"`
	} `json:"main"`
	Visibility int     `json:"visibility"`
	Timezone   int     `json:"timezone"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

type VisitorInfo struct {
	IP      string  `json:"ip"`
	City    string  `json:"city"`
	Country string  `json:"country"`
	ASN     string  `json:"asn"`
	ISP     string  `json:"org"`
	Lat     float64 `json:"latitude"`
	Lon     float64 `json:"longitude"`
}

const (
	boxWidth = 60
)

func fetchWeather(lat, lon float64) (WeatherResponse, error) {
	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%.6f&lon=%.6f&units=metric&appid=%s", lat, lon, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return WeatherResponse{}, err
	}
	defer resp.Body.Close()

	var weather WeatherResponse
	if resp.StatusCode != http.StatusOK {
		return weather, fmt.Errorf("error fetching weather: %s", resp.Status)
	}
	json.NewDecoder(resp.Body).Decode(&weather)
	return weather, nil
}

func getVisitorInfo(r *http.Request) (*VisitorInfo, error) {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	url := fmt.Sprintf("https://ipapi.co/%s/json/", ip)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var visitorInfo VisitorInfo
	if err := json.NewDecoder(resp.Body).Decode(&visitorInfo); err != nil {
		return nil, err
	}
	return &visitorInfo, nil
}

func drawBox(content []string) string {
	header := fmt.Sprintf("╭%s╮\n", strings.Repeat("─", boxWidth-2))
	boxContent := header

	for _, lineContent := range content {
		contentLine := lineContent
		if utf8.RuneCountInString(contentLine) > boxWidth-2 {
			contentLine = contentLine[:boxWidth-5] + "..."
		}
		padding := strings.Repeat(" ", boxWidth-utf8.RuneCountInString(contentLine)-3)
		boxContent += fmt.Sprintf("│ %s%s│\n", contentLine, padding)
	}
	bottomLine := fmt.Sprintf("╰%s╯\n", strings.Repeat("─", boxWidth-2))

	return boxContent + bottomLine
}

func handler(w http.ResponseWriter, r *http.Request) {
	visitorInfo, err := getVisitorInfo(r)
	if err != nil {
		http.Error(w, "Unable to fetch visitor information", http.StatusInternalServerError)
		return
	}

	weather, err := fetchWeather(visitorInfo.Lat, visitorInfo.Lon)
	if err != nil {
		http.Error(w, "Unable to fetch weather data", http.StatusInternalServerError)
		return
	}

	//timezone is a number of seconds after UTC i.e. Perth is 28800 (GMT+8)
	localTime := time.Now().UTC().Add(time.Duration(weather.Timezone) * time.Second).Format("3:04 PM MST")

	content := []string{
		"IP Address: " + visitorInfo.IP,
		"ASN: " + strings.TrimSpace(visitorInfo.ASN),
		"ISP: " + strings.TrimSpace(visitorInfo.ISP),
		"Location: " + strings.TrimSpace(visitorInfo.City + ", " + visitorInfo.Country),
		"",
		fmt.Sprintf("Temperature: %.2f °C", weather.Main.Temp),
		fmt.Sprintf("Pressure: %d hPa", weather.Main.Pressure),
		fmt.Sprintf("Humidity: %d%%", weather.Main.Humidity),
		fmt.Sprintf("Visibility: %d m", weather.Visibility),
		fmt.Sprintf("Time: %s", localTime),
		fmt.Sprintf("Wind Speed: %.2f m/s", weather.Wind.Speed),
		"Description: " + weather.Weather[0].Description,
	}

	response := drawBox(content)
	fmt.Fprint(w, response)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}