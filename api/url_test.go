package api

import (
	"testing"

	"github.com/icza/gog"
)

func TestURLRequest_sanitize(t *testing.T) {
	type fields struct {
		LatitudeStr  *string
		LongitudeStr *string
		PhotoStr     *string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &urlRequest{
				LatitudeStr:  tt.fields.LatitudeStr,
				LongitudeStr: tt.fields.LongitudeStr,
				PhotoStr:     tt.fields.PhotoStr,
			}
			req.sanitize()
		})
	}
}

func TestURLRequest_hasPhoto(t *testing.T) {
	type fields struct {
		PhotoStr *string
		Photo    int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "No photo",
			fields: fields{},
			want:   false,
		},
		{
			name: "Invalid photo",
			fields: fields{
				Photo: 1,
			},
			want: false,
		},
		{
			name: "Has photo",
			fields: fields{
				PhotoStr: gog.Ptr("1"),
				Photo:    1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &urlRequest{
				PhotoStr: tt.fields.PhotoStr,
				Photo:    tt.fields.Photo,
			}
			if got := req.hasPhoto(); got != tt.want {
				t.Errorf("hasPhoto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLRequest_isNilCoordinates(t *testing.T) {
	type fields struct {
		LatitudeStr  *string
		LongitudeStr *string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Latitude nil",
			fields: fields{
				LongitudeStr: gog.Ptr("90"),
			},
			want: true,
		},
		{
			name: "Longitude nil",
			fields: fields{
				LatitudeStr: gog.Ptr("90"),
			},
			want: true,
		},
		{
			name: "Valid",
			fields: fields{
				LatitudeStr:  gog.Ptr("90"),
				LongitudeStr: gog.Ptr("90"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &urlRequest{
				LatitudeStr:  tt.fields.LatitudeStr,
				LongitudeStr: tt.fields.LongitudeStr,
			}
			if got := req.isNilCoordinates(); got != tt.want {
				t.Errorf("isNilCoordinates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLRequest_isValidCoordinates(t *testing.T) {
	type fields struct {
		LatitudeStr  *string
		Latitude     float32
		LongitudeStr *string
		Longitude    float32
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Latitude too small",
			fields: fields{
				LatitudeStr:  gog.Ptr("-90.1"),
				Latitude:     -90.1,
				LongitudeStr: gog.Ptr("0"),
				Longitude:    0,
			},
			want: false,
		},
		{
			name: "Latitude too large",
			fields: fields{
				LatitudeStr:  gog.Ptr("90.1"),
				Latitude:     90.1,
				LongitudeStr: gog.Ptr("0"),
				Longitude:    0,
			},
			want: false,
		},
		{
			name: "Longitude too small",
			fields: fields{
				LatitudeStr:  gog.Ptr("0"),
				Latitude:     0,
				LongitudeStr: gog.Ptr("-180.1"),
				Longitude:    -180.1,
			},
			want: false,
		},
		{
			name: "Longitude too large",
			fields: fields{
				LatitudeStr:  gog.Ptr("0"),
				Latitude:     0,
				LongitudeStr: gog.Ptr("180.1"),
				Longitude:    180.1,
			},
			want: false,
		},
		{
			name:   "Nil coordinates",
			fields: fields{},
			want:   false,
		},
		{
			name: "Valid",
			fields: fields{
				LatitudeStr:  gog.Ptr("90"),
				Latitude:     90,
				LongitudeStr: gog.Ptr("180"),
				Longitude:    180,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &urlRequest{
				LatitudeStr:  tt.fields.LatitudeStr,
				Latitude:     tt.fields.Latitude,
				LongitudeStr: tt.fields.LongitudeStr,
				Longitude:    tt.fields.Longitude,
			}
			if got := req.isValidCoordinates(); got != tt.want {
				t.Errorf("isValidCoordinates() = %v, want %v", got, tt.want)
			}
		})
	}
}
