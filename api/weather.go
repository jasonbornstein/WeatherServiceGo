package api

// weatherRequest corresponds to the query parameters of the openweathermap.org "Current Weather data" request
type weatherRequest struct {
	Latitude  float32 `json:"lat"`             // Latitude
	Longitude float32 `json:"lon"`             // Longitude
	APIKey    string  `json:"appid"`           // Your unique API key
	Mode      *string `json:"mode,omitempty"`  // Response format (xml,html); default is json
	Units     string  `json:"units,omitempty"` // Units of measurement (standard, metric, imperial) default is standard
	Language  *string `json:"lang,omitempty"`  // You can use this parameter to get the output in your language
}

// weatherResponse corresponds to the JSON format API response fields of the openweathermap.org "Current Weather data" request
type weatherResponse struct {
	Coordinates coordinates   `json:"coord"`      //
	Weather     []weather     `json:"weather"`    //
	Base        string        `json:"base"`       // Internal parameter
	Main        main          `json:"main"`       //
	Visibility  int           `json:"visibility"` // Visibility (in meters); max value is 10000
	Wind        wind          `json:"wind"`       //
	Clouds      clouds        `json:"clouds"`     //
	Rain        precipitation `json:"rain"`       //
	Snow        precipitation `json:"snow"`       //
	Date        int64         `json:"dt"`         // Time of data calculation, unix, UTC
	Sys         system        `json:"sys"`        //
	Timezone    int64         `json:"timezone"`   // Shift in seconds from UTC
	ID          *int64        `json:"id"`         // City ID (deprecated)
	Name        *string       `json:"name"`       // City name (deprecated)
	COD         int64         `json:"cod"`        // Internal parameter
}

type coordinates struct {
	Latitude  float32 `json:"lat"` // Latitude of the location
	Longitude float32 `json:"lon"` // Longitude of the location
}
type weather struct {
	ID          int32  `json:"id"`          // Weather condition id
	Main        string `json:"main"`        // Group of Weather parameters (Precipitation, Snow, Clouds etc.)
	Description string `json:"description"` // Weather condition within the group
	Icon        string `json:"icon"`        // Weather icon id
}
type main struct {
	Temperature float32  `json:"temp"`               // Temperature (Kelvin for Units=standard, Celsius for Units=metric, Fahrenheit for units=imperial)
	FeelsLike   float32  `json:"feels_like"`         // This temperature parameter accounts for the human perception of Weather
	Pressure    float32  `json:"pressure"`           // Atmospheric pressure on the sea level (in hPa)
	Humidity    float32  `json:"humidity"`           // Percent Humidity
	TempMinimum *float32 `json:"temp_min,omitempty"` // Minimum temperature at the moment
	TempMaximum *float32 `json:"temp_max,omitempty"` // Maximum temperature at the moment
	SeaLevel    float32  `json:"sea_level"`          // Atmospheric pressure on the sea level (in hPa)
	GroundLevel float32  `json:"grnd_level"`         // Atmospheric pressure on the ground level (in hPa)
}
type wind struct {
	Speed   float32 `json:"speed"` // Wind speed (in meters/second or miles/hour)
	Degrees float32 `json:"deg"`   // Wind direction (in degrees (meteorological))
	Gust    float32 `json:"gust"`  // Wind gust (in meters/second or miles/hour)
}
type clouds struct {
	All int `json:"all"` // Percent Cloudiness
}
type precipitation struct {
	OneHour   *float32 `json:"1h"` // Precipitation volume for the last 1 hour (in mm)
	ThreeHour *float32 `json:"3h"` // Precipitation volume for the last 3 hours (in mm)
}
type system struct {
	Type    int32  `json:"type"`    // Internal parameter
	ID      int32  `json:"id"`      // Internal parameter
	Message string `json:"message"` // Internal parameter
	Country string `json:"country"` // Country code (GB, JP etc.)
	Sunrise int64  `json:"sunrise"` // Sunrise time, unix, UTC
	Sunset  int64  `json:"sunset"`  // Sunset time, unix, UTC
}
