package api

import "strconv"

// urlRequest is our page's query parameters as a struct
type urlRequest struct {
	LatitudeStr  *string `form:"Latitude"` // Latitude
	Latitude     float32
	LongitudeStr *string `form:"Longitude"` // Longitude
	Longitude    float32
	APIKey       *string `form:"APIKey"` // Your unique API key
	PhotoStr     *string `form:"photo"`  // Array index of the Photo selected
	Photo        int
}

func (req *urlRequest) sanitize() {
	if req.LatitudeStr != nil {
		lat, err := strconv.ParseFloat(*req.LatitudeStr, 32)
		if err != nil {
			req.LatitudeStr = nil
		} else {
			req.Latitude = float32(lat)
		}
	}
	if req.LongitudeStr != nil {
		lon, err := strconv.ParseFloat(*req.LongitudeStr, 32)
		if err != nil {
			req.LongitudeStr = nil
		} else {
			req.Longitude = float32(lon)
		}
	}
	if req.PhotoStr != nil {
		photo, err := strconv.ParseInt(*req.PhotoStr, 10, 4)
		if err != nil || photo == -1 {
			req.PhotoStr = nil
			req.Photo = -1
		} else {
			req.Photo = int(photo)
		}
	}
}
func (req *urlRequest) hasPhoto() bool {
	return req.PhotoStr != nil && req.Photo > -1
}
func (req *urlRequest) isNilCoordinates() bool {
	return req.LatitudeStr == nil || req.LongitudeStr == nil
}
func (req *urlRequest) isValidCoordinates() bool {
	if req.isNilCoordinates() {
		return false
	}

	return req.Latitude >= -90 && req.Latitude <= 90 && req.Longitude >= -180 && req.Longitude <= 180
}
