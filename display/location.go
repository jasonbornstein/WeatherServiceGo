package display

import (
	"math"
)

type Coordinates struct {
	Latitude             float32
	LatitudeDescription  float32
	LatitudeDirection    string
	Longitude            float32
	LongitudeDescription float32
	LongitudeDirection   string
}
type Photo struct {
	PhotoURL string
	ImgURL   string
	ImgTitle string
}

// Location holds the parameters for a geographic location which is associated to "local" weather
type Location struct {
	Coordinates
	*Photo
}

var Locations = []Location{
	{
		Coordinates: Coordinates{
			Latitude:  37.832225,
			Longitude: -122.477817,
		},
		Photo: &Photo{
			PhotoURL: "https://www.flickr.com/photos/jasonbornsteinphoto/53308431182/",
			ImgURL:   "https://live.staticflickr.com/65535/53308431182_6429db0308_m.jpg",
			ImgTitle: "ILCE-6400-0011956",
		},
	},
	{
		Coordinates: Coordinates{
			Latitude:  41.077219,
			Longitude: -73.869206,
		},
		Photo: &Photo{
			PhotoURL: "https://www.flickr.com/photos/jasonbornsteinphoto/53409351484/",
			ImgURL:   "https://live.staticflickr.com/65535/53409351484_3f082fb613_m.jpg",
			ImgTitle: "ILCE-6400-0019688",
		},
	},
	{
		Coordinates: Coordinates{
			Latitude:  40.688655,
			Longitude: -74.04392,
		},
		Photo: &Photo{
			PhotoURL: "https://www.flickr.com/photos/jasonbornsteinphoto/53412081012/",
			ImgURL:   "https://live.staticflickr.com/65535/53412081012_01075bc19b_m.jpg",
			ImgTitle: "ILCE-6400-0018323",
		},
	},
	{
		Coordinates: Coordinates{
			Latitude:  40.82641,
			Longitude: -74.246764,
		},
		Photo: &Photo{
			PhotoURL: "https://www.flickr.com/photos/jasonbornsteinphoto/53337991763/",
			ImgURL:   "https://live.staticflickr.com/65535/53337991763_c6b85cf073_m.jpg",
			ImgTitle: "ILCE-6400-0013012",
		},
	},
}

func CreateDisplayLocation(latitude float32, longitude float32) Location {
	loc := Location{
		Coordinates: Coordinates{
			Latitude:  latitude,
			Longitude: longitude,
		},
	}
	for indx := range Locations {
		if Locations[indx].Coordinates == loc.Coordinates {
			loc = Locations[indx]
			break
		}
	}
	loc.LatitudeDescription, loc.LatitudeDirection = getLatitudeDescription(loc.Latitude)
	loc.LongitudeDescription, loc.LongitudeDirection = getLongitudeDescription(loc.Longitude)
	return loc
}
func getLatitudeDescription(lat float32) (float32, string) {
	if lat < 0 {
		return float32(math.Abs(float64(lat))), "S"
	}
	if lat > 0 {
		return lat, "N"
	}
	return lat, "The Equator"
}
func getLongitudeDescription(lon float32) (float32, string) {
	if lon < 0 {
		return float32(math.Abs(float64(lon))), "W"
	}
	if lon > 0 {
		return lon, "E"
	}
	return lon, "The Prime Meridian"
}
