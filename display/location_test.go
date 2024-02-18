package display

import (
	"testing"
)

func TestCreateDisplayLocation_photo(t *testing.T) {
	type args struct {
		latitude  float32
		longitude float32
	}
	type want struct {
		photo bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Stored Location not found",
			args: args{
				latitude:  90,
				longitude: 90,
			},
			want: want{
				photo: false,
			},
		},
		{
			name: "Stored Location found",
			args: args{
				latitude:  37.832225,
				longitude: -122.477817,
			},
			want: want{
				photo: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateDisplayLocation(tt.args.latitude, tt.args.longitude)
			if tt.want.photo == true && got.Photo == nil {
				t.Errorf("CreateDisplayLocation() = %v, want %v", got, tt.want)
			}
			if tt.want.photo == false && got.Photo != nil {
				t.Errorf("CreateDisplayLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLatitudeDescription(t *testing.T) {
	type args struct {
		lat float32
	}
	tests := []struct {
		name  string
		args  args
		want  float32
		want1 string
	}{
		{
			name: "Negative latitude",
			args: args{
				lat: -1,
			},
			want:  1,
			want1: "S",
		},
		{
			name: "Zero latitude",
			args: args{
				lat: 0,
			},
			want:  0,
			want1: "The Equator",
		},
		{
			name: "Positive latitude",
			args: args{
				lat: 1,
			},
			want:  1,
			want1: "N",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getLatitudeDescription(tt.args.lat)
			if got != tt.want {
				t.Errorf("getLatitudeDescription() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getLatitudeDescription() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getLongitudeDescription(t *testing.T) {
	type args struct {
		lon float32
	}
	tests := []struct {
		name  string
		args  args
		want  float32
		want1 string
	}{
		{
			name: "Negative longitude",
			args: args{
				lon: -1,
			},
			want:  1,
			want1: "W",
		},
		{
			name: "Zero longitude",
			args: args{
				lon: 0,
			},
			want:  0,
			want1: "The Prime Meridian",
		},
		{
			name: "Positive longitude",
			args: args{
				lon: 1,
			},
			want:  1,
			want1: "E",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getLongitudeDescription(tt.args.lon)
			if got != tt.want {
				t.Errorf("getLongitudeDescription() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getLongitudeDescription() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
