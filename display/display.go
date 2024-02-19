package display

// tempUnits is a conversion between the `units` field in the request and a display of temperature units
var tempUnits = map[string]string{
	"standard": "K",
	"metric":   "C",
	"imperial": "F",
}

// speedUnits is a conversion between the `units` field in the request and a display of speed units
var speedUnits = map[string]string{
	"standard": "meters/second",
	"metric":   "meters/second",
	"imperial": "miles/hour",
}

// WeatherDisplay holds the parameters that make up the formatted display of weather at a location
type WeatherDisplay struct {
	Description string
	Temperature Temperature
	Humidity    Humidity
	Wind        Wind
	Location    Location
	Clouds      int
}

type Temperature struct {
	Degrees     float32
	FeelsLike   float32
	Units       string
	Description string
}
type Humidity struct {
	Percentage  float32
	Description string
}
type Wind struct {
	Speed       float32
	Degrees     float32
	Direction   string
	Units       string
	Description string
}

func CreateDisplayTemperature(tempDegrees float32, feelsLike float32, reqUnits string) Temperature {
	return Temperature{
		Degrees:     tempDegrees,
		FeelsLike:   feelsLike,
		Units:       tempUnits[reqUnits],
		Description: getTemperatureDescription(tempDegrees, reqUnits),
	}
}
func getTemperatureDescription(tempDegrees float32, reqUnits string) string {
	// TODO JB: this is based on imperial units
	if tempDegrees > 75 {
		return "It's Hot, Hot Hot!!!"
	}
	if tempDegrees < 40 {
		return "Oh baby, it's cold outside!"
	}
	return "It's okay outside."
}

func CreateDisplayHumidity(percentage float32) Humidity {
	return Humidity{
		Percentage:  percentage,
		Description: getHumidityDescription(percentage),
	}
}
func getHumidityDescription(percentage float32) string {
	if percentage < 20 {
		return "The air is very dry."
	}
	if percentage > 80 {
		return "The air is like a swamp."
	}
	if percentage > 60 {
		return "The air is humid."
	}
	return "The air is comfortable"
}

func CreateDisplayWind(windSpeed float32, windDegrees float32, reqUnits string) Wind {
	return Wind{
		Speed:       windSpeed,
		Degrees:     windDegrees,
		Units:       speedUnits[reqUnits],
		Direction:   getWindDirection(windDegrees),
		Description: getWindDescription(windSpeed, reqUnits),
	}
}
func getWindDirection(degrees float32) string {
	if degrees < 11.25 {
		return "N"
	}
	if degrees < 33.75 {
		return "NNE"
	}
	if degrees < 56.25 {
		return "NE"
	}
	if degrees < 78.75 {
		return "ENE"
	}
	if degrees < 101.25 {
		return "E"
	}
	if degrees < 123.75 {
		return "ESE"
	}
	if degrees < 146.25 {
		return "SE"
	}
	if degrees < 168.75 {
		return "SSE"
	}
	if degrees < 191.25 {
		return "S"
	}
	if degrees < 213.75 {
		return "SSW"
	}
	if degrees < 236.25 {
		return "SW"
	}
	if degrees < 258.75 {
		return "WSW"
	}
	if degrees < 281.25 {
		return "W"
	}
	if degrees < 303.75 {
		return "WNW"
	}
	if degrees < 326.25 {
		return "NW"
	}
	if degrees < 348.75 {
		return "NNW"
	}
	return "N"
}
func getWindDescription(speed float32, reqUnits string) string {
	// TODO JB: this is based on imperial units
	if speed < 2 {
		return "The wind is calm."
	}
	if speed > 60 {
		return "You shouldn't be outside!"
	}
	if speed > 40 {
		return "It's very windy!"
	}
	if speed > 20 {
		return "It's windy!"
	}
	return "There's a light breeze."
}
