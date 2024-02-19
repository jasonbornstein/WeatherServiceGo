package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jasonbornstein/WeatherServiceGo/config"
	"github.com/jasonbornstein/WeatherServiceGo/display"
)

// SetupRouter loads all the routes supported by the application and generates the appropriate responses.
func SetupRouter(conf config.Config) *gin.Engine {
	// Serve APIs using Gin API framework
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Load the HTML display template
	router.LoadHTMLFiles("index.html")

	// Add routes and their associated functions
	router.GET("/", func(c *gin.Context) {
		// Bind the query parameters into a struct
		urlReq := &urlRequest{}
		if err := c.Bind(urlReq); err != nil {
			log.Fatalf("error binding URL %v", err)
		}

		// Validation and building of the request and the response display
		urlReq.sanitize()

		// Get the APIKey to use: from the config or from the URL query parameters
		APIKey := conf.APIKey
		if conf.APIKey == nil {
			APIKey = urlReq.APIKey
		}

		data := gin.H{
			"APIKey": APIKey,
		}

		// Only proceed if there's an APIKey. Otherwise, an error will display on the HTML page.
		if APIKey == nil {
			c.HTML(http.StatusOK, "index.html", data)
			return
		}

		// Build the request
		weatherReq := &weatherRequest{
			APIKey: *APIKey,
			Units:  "imperial", // hardcode for now
		}

		data["locations"] = display.Locations // pass a pre-determined array of Locations to choose

		// If a photo was defined in the request, find its Latitude/Longitude. Otherwise, take the Latitude/Longitude that was given, but validate it first.

		if urlReq.hasPhoto() {
			// Valid
			data["photo"] = urlReq.Photo
			photoLocation := &display.Locations[urlReq.Photo]
			weatherReq.Latitude = photoLocation.Latitude
			weatherReq.Longitude = photoLocation.Longitude
		} else if urlReq.isNilCoordinates() {
			// Invalid
			data["error"] = errors.New("Location coordinates are missing")
			c.HTML(http.StatusOK, "index.html", data)
			return
		} else if !urlReq.isValidCoordinates() {
			// Invalid
			data["error"] = errors.New("Location coordinates are invalid")
			c.HTML(http.StatusOK, "index.html", data)
			return
		} else {
			// Valid
			weatherReq.Latitude = urlReq.Latitude
			weatherReq.Longitude = urlReq.Longitude
		}

		data["coordinates"] = display.Coordinates{
			Latitude:  weatherReq.Latitude,
			Longitude: weatherReq.Longitude,
		}

		// Make the request to openweathermap.org
		weatherResp := requestWeather(weatherReq)

		// If a response was returned, create the HTML display
		if weatherResp != nil {
			weatherDisplay := &display.WeatherDisplay{}
			weatherDisplay.Description = weatherResp.Weather[0].Description
			weatherDisplay.Temperature = display.CreateDisplayTemperature(weatherResp.Main.Temperature, weatherResp.Main.FeelsLike, weatherReq.Units)
			weatherDisplay.Humidity = display.CreateDisplayHumidity(weatherResp.Main.Humidity)
			weatherDisplay.Wind = display.CreateDisplayWind(weatherResp.Wind.Speed, weatherResp.Wind.Degrees, weatherReq.Units)
			weatherDisplay.Location = display.CreateDisplayLocation(weatherReq.Latitude, weatherReq.Longitude)
			weatherDisplay.Clouds = weatherResp.Clouds.All

			data["display"] = weatherDisplay
		}

		// Pass the parameters to the HTML template for final display
		c.HTML(http.StatusOK, "index.html", data)
	})

	return router
}

// requestWeather builds and invokes the request to openweathermap.org and parses the response into a struct
func requestWeather(weatherReq *weatherRequest) *weatherResponse {
	// Build the URL with the query parameters we want
	URL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&units=%v&appid=%v", weatherReq.Latitude, weatherReq.Longitude, weatherReq.Units, weatherReq.APIKey)

	// Invoke the URL, read the response, and unmarshal it into our weatherResponse struct
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatalf("error invoking openweathermap.org, %s\n", err)
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("client: could not read response body: %s\n", err)
	}
	weatherResp := weatherResponse{}
	err = json.Unmarshal(resBody, &weatherResp)
	if err != nil {
		log.Fatalf("error unmarshaling response body, %s\n", err)
	}

	return &weatherResp
}
