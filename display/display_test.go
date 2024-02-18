package display

import (
	"testing"
)

func Test_getTemperatureDescription(t *testing.T) {
	type args struct {
		tempDegrees float32
		reqUnits    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Hot",
			args: args{
				tempDegrees: 0,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 1,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 2,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 3,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 4,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 5,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 6,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 7,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 8,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 9,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 10,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 11,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 12,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 13,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 14,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 15,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 16,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 17,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 18,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 19,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 20,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 21,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 22,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 23,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 24,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 25,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 26,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 27,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 28,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 29,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 30,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 31,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 32,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 33,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 34,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 35,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 36,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 37,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 38,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 39,
				reqUnits:    "imperial",
			},
			want: "Oh baby, it's cold outside!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 40,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 41,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 42,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 43,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 44,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 45,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 46,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 47,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 48,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 49,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 50,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 51,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 52,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 53,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 54,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 55,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 56,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 57,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 58,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 59,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 60,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 61,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 62,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 63,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 64,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 65,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 66,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 67,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 68,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 69,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 70,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 71,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 72,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 73,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 74,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 75,
				reqUnits:    "imperial",
			},
			want: "It's okay outside.",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 76,
				reqUnits:    "imperial",
			},
			want: "It's Hot, Hot Hot!!!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 77,
				reqUnits:    "imperial",
			},
			want: "It's Hot, Hot Hot!!!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 78,
				reqUnits:    "imperial",
			},
			want: "It's Hot, Hot Hot!!!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 79,
				reqUnits:    "imperial",
			},
			want: "It's Hot, Hot Hot!!!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 80,
				reqUnits:    "imperial",
			},
			want: "It's Hot, Hot Hot!!!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 81,
				reqUnits:    "imperial",
			},
			want: "It's Hot, Hot Hot!!!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 82,
				reqUnits:    "imperial",
			},
			want: "It's Hot, Hot Hot!!!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 83,
				reqUnits:    "imperial",
			},
			want: "It's Hot, Hot Hot!!!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 84,
				reqUnits:    "imperial",
			},
			want: "It's Hot, Hot Hot!!!",
		},
		{
			name: "Hot",
			args: args{
				tempDegrees: 85,
				reqUnits:    "imperial",
			},
			want: "It's Hot, Hot Hot!!!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTemperatureDescription(tt.args.tempDegrees, tt.args.reqUnits); got != tt.want {
				t.Errorf("getTemperatureDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHumidityDescription(t *testing.T) {
	type args struct {
		percentage float32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Dry",
			args: args{
				percentage: 0,
			},
			want: "The air is very dry.",
		},
		{
			name: "Dry",
			args: args{
				percentage: 1,
			},
			want: "The air is very dry.",
		},
		{
			name: "Dry",
			args: args{
				percentage: 2,
			},
			want: "The air is very dry.",
		},
		{
			name: "Dry",
			args: args{
				percentage: 3,
			},
			want: "The air is very dry.",
		},
		{
			name: "Dry",
			args: args{
				percentage: 4,
			},
			want: "The air is very dry.",
		},
		{
			name: "Dry",
			args: args{
				percentage: 5,
			},
			want: "The air is very dry.",
		},
		{
			name: "Dry",
			args: args{
				percentage: 6,
			},
			want: "The air is very dry.",
		},
		{
			name: "Dry",
			args: args{
				percentage: 7,
			},
			want: "The air is very dry.",
		},
		{
			name: "Dry",
			args: args{
				percentage: 8,
			},
			want: "The air is very dry.",
		},
		{
			name: "Dry",
			args: args{
				percentage: 9,
			},
			want: "The air is very dry.",
		},
		{
			name: "Dry",
			args: args{
				percentage: 10,
			},
			want: "The air is very dry.",
		},
		{
			name: "Dry",
			args: args{
				percentage: 11,
			},
			want: "The air is very dry.",
		},
		{
			name: "Dry",
			args: args{
				percentage: 12,
			},
			want: "The air is very dry.",
		},
		{
			name: "Dry",
			args: args{
				percentage: 13,
			},
			want: "The air is very dry.",
		},
		{
			name: "Dry",
			args: args{
				percentage: 14,
			},
			want: "The air is very dry.",
		},
		{
			name: "Dry",
			args: args{
				percentage: 15,
			},
			want: "The air is very dry.",
		},
		{
			name: "Dry",
			args: args{
				percentage: 16,
			},
			want: "The air is very dry.",
		},
		{
			name: "Dry",
			args: args{
				percentage: 17,
			},
			want: "The air is very dry.",
		},
		{
			name: "Dry",
			args: args{
				percentage: 18,
			},
			want: "The air is very dry.",
		},
		{
			name: "Dry",
			args: args{
				percentage: 19,
			},
			want: "The air is very dry.",
		},
		{
			name: "Dry",
			args: args{
				percentage: 20,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 21,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 22,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 23,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 24,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 25,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 26,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 27,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 28,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 29,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 30,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 31,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 32,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 33,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 34,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 35,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 36,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 37,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 38,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 39,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 40,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 41,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 42,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 43,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 44,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 45,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 46,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 47,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 48,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 49,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 50,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 51,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 52,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 53,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 54,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 55,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 56,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 57,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 58,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 59,
			},
			want: "The air is comfortable",
		},
		{
			name: "Comfortable",
			args: args{
				percentage: 60,
			},
			want: "The air is comfortable",
		},
		{
			name: "Humid",
			args: args{
				percentage: 61,
			},
			want: "The air is humid.",
		},
		{
			name: "Humid",
			args: args{
				percentage: 62,
			},
			want: "The air is humid.",
		},
		{
			name: "Humid",
			args: args{
				percentage: 63,
			},
			want: "The air is humid.",
		},
		{
			name: "Humid",
			args: args{
				percentage: 64,
			},
			want: "The air is humid.",
		},
		{
			name: "Humid",
			args: args{
				percentage: 65,
			},
			want: "The air is humid.",
		},
		{
			name: "Humid",
			args: args{
				percentage: 66,
			},
			want: "The air is humid.",
		},
		{
			name: "Humid",
			args: args{
				percentage: 67,
			},
			want: "The air is humid.",
		},
		{
			name: "Humid",
			args: args{
				percentage: 68,
			},
			want: "The air is humid.",
		},
		{
			name: "Humid",
			args: args{
				percentage: 69,
			},
			want: "The air is humid.",
		},
		{
			name: "Humid",
			args: args{
				percentage: 70,
			},
			want: "The air is humid.",
		},
		{
			name: "Humid",
			args: args{
				percentage: 71,
			},
			want: "The air is humid.",
		},
		{
			name: "Humid",
			args: args{
				percentage: 72,
			},
			want: "The air is humid.",
		},
		{
			name: "Humid",
			args: args{
				percentage: 73,
			},
			want: "The air is humid.",
		},
		{
			name: "Humid",
			args: args{
				percentage: 74,
			},
			want: "The air is humid.",
		},
		{
			name: "Humid",
			args: args{
				percentage: 75,
			},
			want: "The air is humid.",
		},
		{
			name: "Humid",
			args: args{
				percentage: 76,
			},
			want: "The air is humid.",
		},
		{
			name: "Humid",
			args: args{
				percentage: 77,
			},
			want: "The air is humid.",
		},
		{
			name: "Humid",
			args: args{
				percentage: 78,
			},
			want: "The air is humid.",
		},
		{
			name: "Humid",
			args: args{
				percentage: 79,
			},
			want: "The air is humid.",
		},
		{
			name: "Humid",
			args: args{
				percentage: 80,
			},
			want: "The air is humid.",
		},
		{
			name: "Swampy",
			args: args{
				percentage: 81,
			},
			want: "The air is like a swamp.",
		},
		{
			name: "Swampy",
			args: args{
				percentage: 82,
			},
			want: "The air is like a swamp.",
		},
		{
			name: "Swampy",
			args: args{
				percentage: 83,
			},
			want: "The air is like a swamp.",
		},
		{
			name: "Swampy",
			args: args{
				percentage: 84,
			},
			want: "The air is like a swamp.",
		},
		{
			name: "Swampy",
			args: args{
				percentage: 85,
			},
			want: "The air is like a swamp.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHumidityDescription(tt.args.percentage); got != tt.want {
				t.Errorf("getHumidityDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getWindDirection(t *testing.T) {
	type args struct {
		degrees float32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0",
			args: args{degrees: 0},
			want: "N",
		},
		{
			name: "1",
			args: args{degrees: 1},
			want: "N",
		},
		{
			name: "2",
			args: args{degrees: 2},
			want: "N",
		},
		{
			name: "3",
			args: args{degrees: 3},
			want: "N",
		},
		{
			name: "4",
			args: args{degrees: 4},
			want: "N",
		},
		{
			name: "5",
			args: args{degrees: 5},
			want: "N",
		},
		{
			name: "6",
			args: args{degrees: 6},
			want: "N",
		},
		{
			name: "7",
			args: args{degrees: 7},
			want: "N",
		},
		{
			name: "8",
			args: args{degrees: 8},
			want: "N",
		},
		{
			name: "9",
			args: args{degrees: 9},
			want: "N",
		},
		{
			name: "10",
			args: args{degrees: 10},
			want: "N",
		},
		{
			name: "11",
			args: args{degrees: 11},
			want: "N",
		},
		{
			name: "12",
			args: args{degrees: 12},
			want: "NNE",
		},
		{
			name: "13",
			args: args{degrees: 13},
			want: "NNE",
		},
		{
			name: "14",
			args: args{degrees: 14},
			want: "NNE",
		},
		{
			name: "15",
			args: args{degrees: 15},
			want: "NNE",
		},
		{
			name: "16",
			args: args{degrees: 16},
			want: "NNE",
		},
		{
			name: "17",
			args: args{degrees: 17},
			want: "NNE",
		},
		{
			name: "18",
			args: args{degrees: 18},
			want: "NNE",
		},
		{
			name: "19",
			args: args{degrees: 19},
			want: "NNE",
		},
		{
			name: "20",
			args: args{degrees: 20},
			want: "NNE",
		},
		{
			name: "21",
			args: args{degrees: 21},
			want: "NNE",
		},
		{
			name: "22",
			args: args{degrees: 22},
			want: "NNE",
		},
		{
			name: "23",
			args: args{degrees: 23},
			want: "NNE",
		},
		{
			name: "24",
			args: args{degrees: 24},
			want: "NNE",
		},
		{
			name: "25",
			args: args{degrees: 25},
			want: "NNE",
		},
		{
			name: "26",
			args: args{degrees: 26},
			want: "NNE",
		},
		{
			name: "27",
			args: args{degrees: 27},
			want: "NNE",
		},
		{
			name: "28",
			args: args{degrees: 28},
			want: "NNE",
		},
		{
			name: "29",
			args: args{degrees: 29},
			want: "NNE",
		},
		{
			name: "30",
			args: args{degrees: 30},
			want: "NNE",
		},
		{
			name: "31",
			args: args{degrees: 31},
			want: "NNE",
		},
		{
			name: "32",
			args: args{degrees: 32},
			want: "NNE",
		},
		{
			name: "33",
			args: args{degrees: 33},
			want: "NNE",
		},
		{
			name: "34",
			args: args{degrees: 34},
			want: "NE",
		},
		{
			name: "35",
			args: args{degrees: 35},
			want: "NE",
		},
		{
			name: "36",
			args: args{degrees: 36},
			want: "NE",
		},
		{
			name: "37",
			args: args{degrees: 37},
			want: "NE",
		},
		{
			name: "38",
			args: args{degrees: 38},
			want: "NE",
		},
		{
			name: "39",
			args: args{degrees: 39},
			want: "NE",
		},
		{
			name: "40",
			args: args{degrees: 40},
			want: "NE",
		},
		{
			name: "41",
			args: args{degrees: 41},
			want: "NE",
		},
		{
			name: "42",
			args: args{degrees: 42},
			want: "NE",
		},
		{
			name: "43",
			args: args{degrees: 43},
			want: "NE",
		},
		{
			name: "44",
			args: args{degrees: 44},
			want: "NE",
		},
		{
			name: "45",
			args: args{degrees: 45},
			want: "NE",
		},
		{
			name: "46",
			args: args{degrees: 46},
			want: "NE",
		},
		{
			name: "47",
			args: args{degrees: 47},
			want: "NE",
		},
		{
			name: "48",
			args: args{degrees: 48},
			want: "NE",
		},
		{
			name: "49",
			args: args{degrees: 49},
			want: "NE",
		},
		{
			name: "50",
			args: args{degrees: 50},
			want: "NE",
		},
		{
			name: "51",
			args: args{degrees: 51},
			want: "NE",
		},
		{
			name: "52",
			args: args{degrees: 52},
			want: "NE",
		},
		{
			name: "53",
			args: args{degrees: 53},
			want: "NE",
		},
		{
			name: "54",
			args: args{degrees: 54},
			want: "NE",
		},
		{
			name: "55",
			args: args{degrees: 55},
			want: "NE",
		},
		{
			name: "56",
			args: args{degrees: 56},
			want: "NE",
		},
		{
			name: "57",
			args: args{degrees: 57},
			want: "ENE",
		},
		{
			name: "58",
			args: args{degrees: 58},
			want: "ENE",
		},
		{
			name: "59",
			args: args{degrees: 59},
			want: "ENE",
		},
		{
			name: "60",
			args: args{degrees: 60},
			want: "ENE",
		},
		{
			name: "61",
			args: args{degrees: 61},
			want: "ENE",
		},
		{
			name: "62",
			args: args{degrees: 62},
			want: "ENE",
		},
		{
			name: "63",
			args: args{degrees: 63},
			want: "ENE",
		},
		{
			name: "64",
			args: args{degrees: 64},
			want: "ENE",
		},
		{
			name: "65",
			args: args{degrees: 65},
			want: "ENE",
		},
		{
			name: "66",
			args: args{degrees: 66},
			want: "ENE",
		},
		{
			name: "67",
			args: args{degrees: 67},
			want: "ENE",
		},
		{
			name: "68",
			args: args{degrees: 68},
			want: "ENE",
		},
		{
			name: "69",
			args: args{degrees: 69},
			want: "ENE",
		},
		{
			name: "70",
			args: args{degrees: 70},
			want: "ENE",
		},
		{
			name: "71",
			args: args{degrees: 71},
			want: "ENE",
		},
		{
			name: "72",
			args: args{degrees: 72},
			want: "ENE",
		},
		{
			name: "73",
			args: args{degrees: 73},
			want: "ENE",
		},
		{
			name: "74",
			args: args{degrees: 74},
			want: "ENE",
		},
		{
			name: "75",
			args: args{degrees: 75},
			want: "ENE",
		},
		{
			name: "76",
			args: args{degrees: 76},
			want: "ENE",
		},
		{
			name: "77",
			args: args{degrees: 77},
			want: "ENE",
		},
		{
			name: "78",
			args: args{degrees: 78},
			want: "ENE",
		},
		{
			name: "79",
			args: args{degrees: 79},
			want: "E",
		},
		{
			name: "80",
			args: args{degrees: 80},
			want: "E",
		},
		{
			name: "81",
			args: args{degrees: 81},
			want: "E",
		},
		{
			name: "82",
			args: args{degrees: 82},
			want: "E",
		},
		{
			name: "83",
			args: args{degrees: 83},
			want: "E",
		},
		{
			name: "84",
			args: args{degrees: 84},
			want: "E",
		},
		{
			name: "85",
			args: args{degrees: 85},
			want: "E",
		},
		{
			name: "86",
			args: args{degrees: 86},
			want: "E",
		},
		{
			name: "87",
			args: args{degrees: 87},
			want: "E",
		},
		{
			name: "88",
			args: args{degrees: 88},
			want: "E",
		},
		{
			name: "89",
			args: args{degrees: 89},
			want: "E",
		},
		{
			name: "90",
			args: args{degrees: 90},
			want: "E",
		},
		{
			name: "91",
			args: args{degrees: 91},
			want: "E",
		},
		{
			name: "92",
			args: args{degrees: 92},
			want: "E",
		},
		{
			name: "93",
			args: args{degrees: 93},
			want: "E",
		},
		{
			name: "94",
			args: args{degrees: 94},
			want: "E",
		},
		{
			name: "95",
			args: args{degrees: 95},
			want: "E",
		},
		{
			name: "96",
			args: args{degrees: 96},
			want: "E",
		},
		{
			name: "97",
			args: args{degrees: 97},
			want: "E",
		},
		{
			name: "98",
			args: args{degrees: 98},
			want: "E",
		},
		{
			name: "99",
			args: args{degrees: 99},
			want: "E",
		},
		{
			name: "100",
			args: args{degrees: 100},
			want: "E",
		},
		{
			name: "101",
			args: args{degrees: 101},
			want: "E",
		},
		{
			name: "102",
			args: args{degrees: 102},
			want: "ESE",
		},
		{
			name: "103",
			args: args{degrees: 103},
			want: "ESE",
		},
		{
			name: "104",
			args: args{degrees: 104},
			want: "ESE",
		},
		{
			name: "105",
			args: args{degrees: 105},
			want: "ESE",
		},
		{
			name: "106",
			args: args{degrees: 106},
			want: "ESE",
		},
		{
			name: "107",
			args: args{degrees: 107},
			want: "ESE",
		},
		{
			name: "108",
			args: args{degrees: 108},
			want: "ESE",
		},
		{
			name: "109",
			args: args{degrees: 109},
			want: "ESE",
		},
		{
			name: "110",
			args: args{degrees: 110},
			want: "ESE",
		},
		{
			name: "111",
			args: args{degrees: 111},
			want: "ESE",
		},
		{
			name: "112",
			args: args{degrees: 112},
			want: "ESE",
		},
		{
			name: "113",
			args: args{degrees: 113},
			want: "ESE",
		},
		{
			name: "114",
			args: args{degrees: 114},
			want: "ESE",
		},
		{
			name: "115",
			args: args{degrees: 115},
			want: "ESE",
		},
		{
			name: "116",
			args: args{degrees: 116},
			want: "ESE",
		},
		{
			name: "117",
			args: args{degrees: 117},
			want: "ESE",
		},
		{
			name: "118",
			args: args{degrees: 118},
			want: "ESE",
		},
		{
			name: "119",
			args: args{degrees: 119},
			want: "ESE",
		},
		{
			name: "120",
			args: args{degrees: 120},
			want: "ESE",
		},
		{
			name: "121",
			args: args{degrees: 121},
			want: "ESE",
		},
		{
			name: "122",
			args: args{degrees: 122},
			want: "ESE",
		},
		{
			name: "123",
			args: args{degrees: 123},
			want: "ESE",
		},
		{
			name: "124",
			args: args{degrees: 124},
			want: "SE",
		},
		{
			name: "125",
			args: args{degrees: 125},
			want: "SE",
		},
		{
			name: "126",
			args: args{degrees: 126},
			want: "SE",
		},
		{
			name: "127",
			args: args{degrees: 127},
			want: "SE",
		},
		{
			name: "128",
			args: args{degrees: 128},
			want: "SE",
		},
		{
			name: "129",
			args: args{degrees: 129},
			want: "SE",
		},
		{
			name: "130",
			args: args{degrees: 130},
			want: "SE",
		},
		{
			name: "131",
			args: args{degrees: 131},
			want: "SE",
		},
		{
			name: "132",
			args: args{degrees: 132},
			want: "SE",
		},
		{
			name: "133",
			args: args{degrees: 133},
			want: "SE",
		},
		{
			name: "134",
			args: args{degrees: 134},
			want: "SE",
		},
		{
			name: "135",
			args: args{degrees: 135},
			want: "SE",
		},
		{
			name: "136",
			args: args{degrees: 136},
			want: "SE",
		},
		{
			name: "137",
			args: args{degrees: 137},
			want: "SE",
		},
		{
			name: "138",
			args: args{degrees: 138},
			want: "SE",
		},
		{
			name: "139",
			args: args{degrees: 139},
			want: "SE",
		},
		{
			name: "140",
			args: args{degrees: 140},
			want: "SE",
		},
		{
			name: "141",
			args: args{degrees: 141},
			want: "SE",
		},
		{
			name: "142",
			args: args{degrees: 142},
			want: "SE",
		},
		{
			name: "143",
			args: args{degrees: 143},
			want: "SE",
		},
		{
			name: "144",
			args: args{degrees: 144},
			want: "SE",
		},
		{
			name: "145",
			args: args{degrees: 145},
			want: "SE",
		},
		{
			name: "146",
			args: args{degrees: 146},
			want: "SE",
		},
		{
			name: "147",
			args: args{degrees: 147},
			want: "SSE",
		},
		{
			name: "148",
			args: args{degrees: 148},
			want: "SSE",
		},
		{
			name: "149",
			args: args{degrees: 149},
			want: "SSE",
		},
		{
			name: "150",
			args: args{degrees: 150},
			want: "SSE",
		},
		{
			name: "151",
			args: args{degrees: 151},
			want: "SSE",
		},
		{
			name: "152",
			args: args{degrees: 152},
			want: "SSE",
		},
		{
			name: "153",
			args: args{degrees: 153},
			want: "SSE",
		},
		{
			name: "154",
			args: args{degrees: 154},
			want: "SSE",
		},
		{
			name: "155",
			args: args{degrees: 155},
			want: "SSE",
		},
		{
			name: "156",
			args: args{degrees: 156},
			want: "SSE",
		},
		{
			name: "157",
			args: args{degrees: 157},
			want: "SSE",
		},
		{
			name: "158",
			args: args{degrees: 158},
			want: "SSE",
		},
		{
			name: "159",
			args: args{degrees: 159},
			want: "SSE",
		},
		{
			name: "160",
			args: args{degrees: 160},
			want: "SSE",
		},
		{
			name: "161",
			args: args{degrees: 161},
			want: "SSE",
		},
		{
			name: "162",
			args: args{degrees: 162},
			want: "SSE",
		},
		{
			name: "163",
			args: args{degrees: 163},
			want: "SSE",
		},
		{
			name: "164",
			args: args{degrees: 164},
			want: "SSE",
		},
		{
			name: "165",
			args: args{degrees: 165},
			want: "SSE",
		},
		{
			name: "166",
			args: args{degrees: 166},
			want: "SSE",
		},
		{
			name: "167",
			args: args{degrees: 167},
			want: "SSE",
		},
		{
			name: "168",
			args: args{degrees: 168},
			want: "SSE",
		},
		{
			name: "169",
			args: args{degrees: 169},
			want: "S",
		},
		{
			name: "170",
			args: args{degrees: 170},
			want: "S",
		},
		{
			name: "171",
			args: args{degrees: 171},
			want: "S",
		},
		{
			name: "172",
			args: args{degrees: 172},
			want: "S",
		},
		{
			name: "173",
			args: args{degrees: 173},
			want: "S",
		},
		{
			name: "174",
			args: args{degrees: 174},
			want: "S",
		},
		{
			name: "175",
			args: args{degrees: 175},
			want: "S",
		},
		{
			name: "176",
			args: args{degrees: 176},
			want: "S",
		},
		{
			name: "177",
			args: args{degrees: 177},
			want: "S",
		},
		{
			name: "178",
			args: args{degrees: 178},
			want: "S",
		},
		{
			name: "179",
			args: args{degrees: 179},
			want: "S",
		},
		{
			name: "180",
			args: args{degrees: 180},
			want: "S",
		},
		{
			name: "181",
			args: args{degrees: 181},
			want: "S",
		},
		{
			name: "182",
			args: args{degrees: 182},
			want: "S",
		},
		{
			name: "183",
			args: args{degrees: 183},
			want: "S",
		},
		{
			name: "184",
			args: args{degrees: 184},
			want: "S",
		},
		{
			name: "185",
			args: args{degrees: 185},
			want: "S",
		},
		{
			name: "186",
			args: args{degrees: 186},
			want: "S",
		},
		{
			name: "187",
			args: args{degrees: 187},
			want: "S",
		},
		{
			name: "188",
			args: args{degrees: 188},
			want: "S",
		},
		{
			name: "189",
			args: args{degrees: 189},
			want: "S",
		},
		{
			name: "190",
			args: args{degrees: 190},
			want: "S",
		},
		{
			name: "191",
			args: args{degrees: 191},
			want: "S",
		},
		{
			name: "192",
			args: args{degrees: 192},
			want: "SSW",
		},
		{
			name: "193",
			args: args{degrees: 193},
			want: "SSW",
		},
		{
			name: "194",
			args: args{degrees: 194},
			want: "SSW",
		},
		{
			name: "195",
			args: args{degrees: 195},
			want: "SSW",
		},
		{
			name: "196",
			args: args{degrees: 196},
			want: "SSW",
		},
		{
			name: "197",
			args: args{degrees: 197},
			want: "SSW",
		},
		{
			name: "198",
			args: args{degrees: 198},
			want: "SSW",
		},
		{
			name: "199",
			args: args{degrees: 199},
			want: "SSW",
		},
		{
			name: "200",
			args: args{degrees: 200},
			want: "SSW",
		},
		{
			name: "201",
			args: args{degrees: 201},
			want: "SSW",
		},
		{
			name: "202",
			args: args{degrees: 202},
			want: "SSW",
		},
		{
			name: "203",
			args: args{degrees: 203},
			want: "SSW",
		},
		{
			name: "204",
			args: args{degrees: 204},
			want: "SSW",
		},
		{
			name: "205",
			args: args{degrees: 205},
			want: "SSW",
		},
		{
			name: "206",
			args: args{degrees: 206},
			want: "SSW",
		},
		{
			name: "207",
			args: args{degrees: 207},
			want: "SSW",
		},
		{
			name: "208",
			args: args{degrees: 208},
			want: "SSW",
		},
		{
			name: "209",
			args: args{degrees: 209},
			want: "SSW",
		},
		{
			name: "210",
			args: args{degrees: 210},
			want: "SSW",
		},
		{
			name: "211",
			args: args{degrees: 211},
			want: "SSW",
		},
		{
			name: "212",
			args: args{degrees: 212},
			want: "SSW",
		},
		{
			name: "213",
			args: args{degrees: 213},
			want: "SSW",
		},
		{
			name: "214",
			args: args{degrees: 214},
			want: "SW",
		},
		{
			name: "215",
			args: args{degrees: 215},
			want: "SW",
		},
		{
			name: "216",
			args: args{degrees: 216},
			want: "SW",
		},
		{
			name: "217",
			args: args{degrees: 217},
			want: "SW",
		},
		{
			name: "218",
			args: args{degrees: 218},
			want: "SW",
		},
		{
			name: "219",
			args: args{degrees: 219},
			want: "SW",
		},
		{
			name: "220",
			args: args{degrees: 220},
			want: "SW",
		},
		{
			name: "221",
			args: args{degrees: 221},
			want: "SW",
		},
		{
			name: "222",
			args: args{degrees: 222},
			want: "SW",
		},
		{
			name: "223",
			args: args{degrees: 223},
			want: "SW",
		},
		{
			name: "224",
			args: args{degrees: 224},
			want: "SW",
		},
		{
			name: "225",
			args: args{degrees: 225},
			want: "SW",
		},
		{
			name: "226",
			args: args{degrees: 226},
			want: "SW",
		},
		{
			name: "227",
			args: args{degrees: 227},
			want: "SW",
		},
		{
			name: "228",
			args: args{degrees: 228},
			want: "SW",
		},
		{
			name: "229",
			args: args{degrees: 229},
			want: "SW",
		},
		{
			name: "230",
			args: args{degrees: 230},
			want: "SW",
		},
		{
			name: "231",
			args: args{degrees: 231},
			want: "SW",
		},
		{
			name: "232",
			args: args{degrees: 232},
			want: "SW",
		},
		{
			name: "233",
			args: args{degrees: 233},
			want: "SW",
		},
		{
			name: "234",
			args: args{degrees: 234},
			want: "SW",
		},
		{
			name: "235",
			args: args{degrees: 235},
			want: "SW",
		},
		{
			name: "236",
			args: args{degrees: 236},
			want: "SW",
		},
		{
			name: "237",
			args: args{degrees: 237},
			want: "WSW",
		},
		{
			name: "238",
			args: args{degrees: 238},
			want: "WSW",
		},
		{
			name: "239",
			args: args{degrees: 239},
			want: "WSW",
		},
		{
			name: "240",
			args: args{degrees: 240},
			want: "WSW",
		},
		{
			name: "241",
			args: args{degrees: 241},
			want: "WSW",
		},
		{
			name: "242",
			args: args{degrees: 242},
			want: "WSW",
		},
		{
			name: "243",
			args: args{degrees: 243},
			want: "WSW",
		},
		{
			name: "244",
			args: args{degrees: 244},
			want: "WSW",
		},
		{
			name: "245",
			args: args{degrees: 245},
			want: "WSW",
		},
		{
			name: "246",
			args: args{degrees: 246},
			want: "WSW",
		},
		{
			name: "247",
			args: args{degrees: 247},
			want: "WSW",
		},
		{
			name: "248",
			args: args{degrees: 248},
			want: "WSW",
		},
		{
			name: "249",
			args: args{degrees: 249},
			want: "WSW",
		},
		{
			name: "250",
			args: args{degrees: 250},
			want: "WSW",
		},
		{
			name: "251",
			args: args{degrees: 251},
			want: "WSW",
		},
		{
			name: "252",
			args: args{degrees: 252},
			want: "WSW",
		},
		{
			name: "253",
			args: args{degrees: 253},
			want: "WSW",
		},
		{
			name: "254",
			args: args{degrees: 254},
			want: "WSW",
		},
		{
			name: "255",
			args: args{degrees: 255},
			want: "WSW",
		},
		{
			name: "256",
			args: args{degrees: 256},
			want: "WSW",
		},
		{
			name: "257",
			args: args{degrees: 257},
			want: "WSW",
		},
		{
			name: "258",
			args: args{degrees: 258},
			want: "WSW",
		},
		{
			name: "259",
			args: args{degrees: 259},
			want: "W",
		},
		{
			name: "260",
			args: args{degrees: 260},
			want: "W",
		},
		{
			name: "261",
			args: args{degrees: 261},
			want: "W",
		},
		{
			name: "262",
			args: args{degrees: 262},
			want: "W",
		},
		{
			name: "263",
			args: args{degrees: 263},
			want: "W",
		},
		{
			name: "264",
			args: args{degrees: 264},
			want: "W",
		},
		{
			name: "265",
			args: args{degrees: 265},
			want: "W",
		},
		{
			name: "266",
			args: args{degrees: 266},
			want: "W",
		},
		{
			name: "267",
			args: args{degrees: 267},
			want: "W",
		},
		{
			name: "268",
			args: args{degrees: 268},
			want: "W",
		},
		{
			name: "269",
			args: args{degrees: 269},
			want: "W",
		},
		{
			name: "270",
			args: args{degrees: 270},
			want: "W",
		},
		{
			name: "271",
			args: args{degrees: 271},
			want: "W",
		},
		{
			name: "272",
			args: args{degrees: 272},
			want: "W",
		},
		{
			name: "273",
			args: args{degrees: 273},
			want: "W",
		},
		{
			name: "274",
			args: args{degrees: 274},
			want: "W",
		},
		{
			name: "275",
			args: args{degrees: 275},
			want: "W",
		},
		{
			name: "276",
			args: args{degrees: 276},
			want: "W",
		},
		{
			name: "277",
			args: args{degrees: 277},
			want: "W",
		},
		{
			name: "278",
			args: args{degrees: 278},
			want: "W",
		},
		{
			name: "279",
			args: args{degrees: 279},
			want: "W",
		},
		{
			name: "280",
			args: args{degrees: 280},
			want: "W",
		},
		{
			name: "281",
			args: args{degrees: 281},
			want: "W",
		},
		{
			name: "282",
			args: args{degrees: 282},
			want: "WNW",
		},
		{
			name: "283",
			args: args{degrees: 283},
			want: "WNW",
		},
		{
			name: "284",
			args: args{degrees: 284},
			want: "WNW",
		},
		{
			name: "285",
			args: args{degrees: 285},
			want: "WNW",
		},
		{
			name: "286",
			args: args{degrees: 286},
			want: "WNW",
		},
		{
			name: "287",
			args: args{degrees: 287},
			want: "WNW",
		},
		{
			name: "288",
			args: args{degrees: 288},
			want: "WNW",
		},
		{
			name: "289",
			args: args{degrees: 289},
			want: "WNW",
		},
		{
			name: "290",
			args: args{degrees: 290},
			want: "WNW",
		},
		{
			name: "291",
			args: args{degrees: 291},
			want: "WNW",
		},
		{
			name: "292",
			args: args{degrees: 292},
			want: "WNW",
		},
		{
			name: "293",
			args: args{degrees: 293},
			want: "WNW",
		},
		{
			name: "294",
			args: args{degrees: 294},
			want: "WNW",
		},
		{
			name: "295",
			args: args{degrees: 295},
			want: "WNW",
		},
		{
			name: "296",
			args: args{degrees: 296},
			want: "WNW",
		},
		{
			name: "297",
			args: args{degrees: 297},
			want: "WNW",
		},
		{
			name: "298",
			args: args{degrees: 298},
			want: "WNW",
		},
		{
			name: "299",
			args: args{degrees: 299},
			want: "WNW",
		},
		{
			name: "300",
			args: args{degrees: 300},
			want: "WNW",
		},
		{
			name: "301",
			args: args{degrees: 301},
			want: "WNW",
		},
		{
			name: "302",
			args: args{degrees: 302},
			want: "WNW",
		},
		{
			name: "303",
			args: args{degrees: 303},
			want: "WNW",
		},
		{
			name: "304",
			args: args{degrees: 304},
			want: "NW",
		},
		{
			name: "305",
			args: args{degrees: 305},
			want: "NW",
		},
		{
			name: "306",
			args: args{degrees: 306},
			want: "NW",
		},
		{
			name: "307",
			args: args{degrees: 307},
			want: "NW",
		},
		{
			name: "308",
			args: args{degrees: 308},
			want: "NW",
		},
		{
			name: "309",
			args: args{degrees: 309},
			want: "NW",
		},
		{
			name: "310",
			args: args{degrees: 310},
			want: "NW",
		},
		{
			name: "311",
			args: args{degrees: 311},
			want: "NW",
		},
		{
			name: "312",
			args: args{degrees: 312},
			want: "NW",
		},
		{
			name: "313",
			args: args{degrees: 313},
			want: "NW",
		},
		{
			name: "314",
			args: args{degrees: 314},
			want: "NW",
		},
		{
			name: "315",
			args: args{degrees: 315},
			want: "NW",
		},
		{
			name: "316",
			args: args{degrees: 316},
			want: "NW",
		},
		{
			name: "317",
			args: args{degrees: 317},
			want: "NW",
		},
		{
			name: "318",
			args: args{degrees: 318},
			want: "NW",
		},
		{
			name: "319",
			args: args{degrees: 319},
			want: "NW",
		},
		{
			name: "320",
			args: args{degrees: 320},
			want: "NW",
		},
		{
			name: "321",
			args: args{degrees: 321},
			want: "NW",
		},
		{
			name: "322",
			args: args{degrees: 322},
			want: "NW",
		},
		{
			name: "323",
			args: args{degrees: 323},
			want: "NW",
		},
		{
			name: "324",
			args: args{degrees: 324},
			want: "NW",
		},
		{
			name: "325",
			args: args{degrees: 325},
			want: "NW",
		},
		{
			name: "326",
			args: args{degrees: 326},
			want: "NW",
		},
		{
			name: "327",
			args: args{degrees: 327},
			want: "NNW",
		},
		{
			name: "328",
			args: args{degrees: 328},
			want: "NNW",
		},
		{
			name: "329",
			args: args{degrees: 329},
			want: "NNW",
		},
		{
			name: "330",
			args: args{degrees: 330},
			want: "NNW",
		},
		{
			name: "331",
			args: args{degrees: 331},
			want: "NNW",
		},
		{
			name: "332",
			args: args{degrees: 332},
			want: "NNW",
		},
		{
			name: "333",
			args: args{degrees: 333},
			want: "NNW",
		},
		{
			name: "334",
			args: args{degrees: 334},
			want: "NNW",
		},
		{
			name: "335",
			args: args{degrees: 335},
			want: "NNW",
		},
		{
			name: "336",
			args: args{degrees: 336},
			want: "NNW",
		},
		{
			name: "337",
			args: args{degrees: 337},
			want: "NNW",
		},
		{
			name: "338",
			args: args{degrees: 338},
			want: "NNW",
		},
		{
			name: "339",
			args: args{degrees: 339},
			want: "NNW",
		},
		{
			name: "340",
			args: args{degrees: 340},
			want: "NNW",
		},
		{
			name: "341",
			args: args{degrees: 341},
			want: "NNW",
		},
		{
			name: "342",
			args: args{degrees: 342},
			want: "NNW",
		},
		{
			name: "343",
			args: args{degrees: 343},
			want: "NNW",
		},
		{
			name: "344",
			args: args{degrees: 344},
			want: "NNW",
		},
		{
			name: "345",
			args: args{degrees: 345},
			want: "NNW",
		},
		{
			name: "346",
			args: args{degrees: 346},
			want: "NNW",
		},
		{
			name: "347",
			args: args{degrees: 347},
			want: "NNW",
		},
		{
			name: "348",
			args: args{degrees: 348},
			want: "NNW",
		},
		{
			name: "349",
			args: args{degrees: 349},
			want: "N",
		},
		{
			name: "350",
			args: args{degrees: 350},
			want: "N",
		},
		{
			name: "351",
			args: args{degrees: 351},
			want: "N",
		},
		{
			name: "352",
			args: args{degrees: 352},
			want: "N",
		},
		{
			name: "353",
			args: args{degrees: 353},
			want: "N",
		},
		{
			name: "354",
			args: args{degrees: 354},
			want: "N",
		},
		{
			name: "355",
			args: args{degrees: 355},
			want: "N",
		},
		{
			name: "356",
			args: args{degrees: 356},
			want: "N",
		},
		{
			name: "357",
			args: args{degrees: 357},
			want: "N",
		},
		{
			name: "358",
			args: args{degrees: 358},
			want: "N",
		},
		{
			name: "359",
			args: args{degrees: 359},
			want: "N",
		},
		{
			name: "360",
			args: args{degrees: 360},
			want: "N",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWindDirection(tt.args.degrees); got != tt.want {
				t.Errorf("getWindDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getWindDescription(t *testing.T) {
	type args struct {
		speed    float32
		reqUnits string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Calm",
			args: args{
				speed:    0,
				reqUnits: "imperial",
			},
			want: "The wind is calm.",
		},
		{
			name: "Calm",
			args: args{
				speed:    1,
				reqUnits: "imperial",
			},
			want: "The wind is calm.",
		},
		{
			name: "Breezy",
			args: args{
				speed:    2,
				reqUnits: "imperial",
			},
			want: "There's a light breeze.",
		},
		{
			name: "Breezy",
			args: args{
				speed:    3,
				reqUnits: "imperial",
			},
			want: "There's a light breeze.",
		},
		{
			name: "Breezy",
			args: args{
				speed:    4,
				reqUnits: "imperial",
			},
			want: "There's a light breeze.",
		},
		{
			name: "Breezy",
			args: args{
				speed:    5,
				reqUnits: "imperial",
			},
			want: "There's a light breeze.",
		},
		{
			name: "Breezy",
			args: args{
				speed:    6,
				reqUnits: "imperial",
			},
			want: "There's a light breeze.",
		},
		{
			name: "Breezy",
			args: args{
				speed:    7,
				reqUnits: "imperial",
			},
			want: "There's a light breeze.",
		},
		{
			name: "Breezy",
			args: args{
				speed:    8,
				reqUnits: "imperial",
			},
			want: "There's a light breeze.",
		},
		{
			name: "Breezy",
			args: args{
				speed:    9,
				reqUnits: "imperial",
			},
			want: "There's a light breeze.",
		},
		{
			name: "Breezy",
			args: args{
				speed:    10,
				reqUnits: "imperial",
			},
			want: "There's a light breeze.",
		},
		{
			name: "Breezy",
			args: args{
				speed:    11,
				reqUnits: "imperial",
			},
			want: "There's a light breeze.",
		},
		{
			name: "Breezy",
			args: args{
				speed:    12,
				reqUnits: "imperial",
			},
			want: "There's a light breeze.",
		},
		{
			name: "Breezy",
			args: args{
				speed:    13,
				reqUnits: "imperial",
			},
			want: "There's a light breeze.",
		},
		{
			name: "Breezy",
			args: args{
				speed:    14,
				reqUnits: "imperial",
			},
			want: "There's a light breeze.",
		},
		{
			name: "Breezy",
			args: args{
				speed:    15,
				reqUnits: "imperial",
			},
			want: "There's a light breeze.",
		},
		{
			name: "Breezy",
			args: args{
				speed:    16,
				reqUnits: "imperial",
			},
			want: "There's a light breeze.",
		},
		{
			name: "Breezy",
			args: args{
				speed:    17,
				reqUnits: "imperial",
			},
			want: "There's a light breeze.",
		},
		{
			name: "Breezy",
			args: args{
				speed:    18,
				reqUnits: "imperial",
			},
			want: "There's a light breeze.",
		},
		{
			name: "Breezy",
			args: args{
				speed:    19,
				reqUnits: "imperial",
			},
			want: "There's a light breeze.",
		},
		{
			name: "Breezy",
			args: args{
				speed:    20,
				reqUnits: "imperial",
			},
			want: "There's a light breeze.",
		},
		{
			name: "Windy",
			args: args{
				speed:    21,
				reqUnits: "imperial",
			},
			want: "It's windy!",
		},
		{
			name: "Windy",
			args: args{
				speed:    22,
				reqUnits: "imperial",
			},
			want: "It's windy!",
		},
		{
			name: "Windy",
			args: args{
				speed:    23,
				reqUnits: "imperial",
			},
			want: "It's windy!",
		},
		{
			name: "Windy",
			args: args{
				speed:    24,
				reqUnits: "imperial",
			},
			want: "It's windy!",
		},
		{
			name: "Windy",
			args: args{
				speed:    25,
				reqUnits: "imperial",
			},
			want: "It's windy!",
		},
		{
			name: "Windy",
			args: args{
				speed:    26,
				reqUnits: "imperial",
			},
			want: "It's windy!",
		},
		{
			name: "Windy",
			args: args{
				speed:    27,
				reqUnits: "imperial",
			},
			want: "It's windy!",
		},
		{
			name: "Windy",
			args: args{
				speed:    28,
				reqUnits: "imperial",
			},
			want: "It's windy!",
		},
		{
			name: "Windy",
			args: args{
				speed:    29,
				reqUnits: "imperial",
			},
			want: "It's windy!",
		},
		{
			name: "Windy",
			args: args{
				speed:    30,
				reqUnits: "imperial",
			},
			want: "It's windy!",
		},
		{
			name: "Windy",
			args: args{
				speed:    31,
				reqUnits: "imperial",
			},
			want: "It's windy!",
		},
		{
			name: "Windy",
			args: args{
				speed:    32,
				reqUnits: "imperial",
			},
			want: "It's windy!",
		},
		{
			name: "Windy",
			args: args{
				speed:    33,
				reqUnits: "imperial",
			},
			want: "It's windy!",
		},
		{
			name: "Windy",
			args: args{
				speed:    34,
				reqUnits: "imperial",
			},
			want: "It's windy!",
		},
		{
			name: "Windy",
			args: args{
				speed:    35,
				reqUnits: "imperial",
			},
			want: "It's windy!",
		},
		{
			name: "Windy",
			args: args{
				speed:    36,
				reqUnits: "imperial",
			},
			want: "It's windy!",
		},
		{
			name: "Windy",
			args: args{
				speed:    37,
				reqUnits: "imperial",
			},
			want: "It's windy!",
		},
		{
			name: "Windy",
			args: args{
				speed:    38,
				reqUnits: "imperial",
			},
			want: "It's windy!",
		},
		{
			name: "Windy",
			args: args{
				speed:    39,
				reqUnits: "imperial",
			},
			want: "It's windy!",
		},
		{
			name: "Windy",
			args: args{
				speed:    40,
				reqUnits: "imperial",
			},
			want: "It's windy!",
		},
		{
			name: "Very Windy",
			args: args{
				speed:    41,
				reqUnits: "imperial",
			},
			want: "It's very windy!",
		},
		{
			name: "Very Windy",
			args: args{
				speed:    42,
				reqUnits: "imperial",
			},
			want: "It's very windy!",
		},
		{
			name: "Very Windy",
			args: args{
				speed:    43,
				reqUnits: "imperial",
			},
			want: "It's very windy!",
		},
		{
			name: "Very Windy",
			args: args{
				speed:    44,
				reqUnits: "imperial",
			},
			want: "It's very windy!",
		},
		{
			name: "Very Windy",
			args: args{
				speed:    45,
				reqUnits: "imperial",
			},
			want: "It's very windy!",
		},
		{
			name: "Very Windy",
			args: args{
				speed:    46,
				reqUnits: "imperial",
			},
			want: "It's very windy!",
		},
		{
			name: "Very Windy",
			args: args{
				speed:    47,
				reqUnits: "imperial",
			},
			want: "It's very windy!",
		},
		{
			name: "Very Windy",
			args: args{
				speed:    48,
				reqUnits: "imperial",
			},
			want: "It's very windy!",
		},
		{
			name: "Very Windy",
			args: args{
				speed:    49,
				reqUnits: "imperial",
			},
			want: "It's very windy!",
		},
		{
			name: "Very Windy",
			args: args{
				speed:    50,
				reqUnits: "imperial",
			},
			want: "It's very windy!",
		},
		{
			name: "Very Windy",
			args: args{
				speed:    51,
				reqUnits: "imperial",
			},
			want: "It's very windy!",
		},
		{
			name: "Very Windy",
			args: args{
				speed:    52,
				reqUnits: "imperial",
			},
			want: "It's very windy!",
		},
		{
			name: "Very Windy",
			args: args{
				speed:    53,
				reqUnits: "imperial",
			},
			want: "It's very windy!",
		},
		{
			name: "Very Windy",
			args: args{
				speed:    54,
				reqUnits: "imperial",
			},
			want: "It's very windy!",
		},
		{
			name: "Very Windy",
			args: args{
				speed:    55,
				reqUnits: "imperial",
			},
			want: "It's very windy!",
		},
		{
			name: "Very Windy",
			args: args{
				speed:    56,
				reqUnits: "imperial",
			},
			want: "It's very windy!",
		},
		{
			name: "Very Windy",
			args: args{
				speed:    57,
				reqUnits: "imperial",
			},
			want: "It's very windy!",
		},
		{
			name: "Very Windy",
			args: args{
				speed:    58,
				reqUnits: "imperial",
			},
			want: "It's very windy!",
		},
		{
			name: "Very Windy",
			args: args{
				speed:    59,
				reqUnits: "imperial",
			},
			want: "It's very windy!",
		},
		{
			name: "Very Windy",
			args: args{
				speed:    60,
				reqUnits: "imperial",
			},
			want: "It's very windy!",
		},
		{
			name: "Too Windy",
			args: args{
				speed:    61,
				reqUnits: "imperial",
			},
			want: "You shouldn't be outside!",
		},
		{
			name: "Too Windy",
			args: args{
				speed:    62,
				reqUnits: "imperial",
			},
			want: "You shouldn't be outside!",
		},
		{
			name: "Too Windy",
			args: args{
				speed:    63,
				reqUnits: "imperial",
			},
			want: "You shouldn't be outside!",
		},
		{
			name: "Too Windy",
			args: args{
				speed:    64,
				reqUnits: "imperial",
			},
			want: "You shouldn't be outside!",
		},
		{
			name: "Too Windy",
			args: args{
				speed:    65,
				reqUnits: "imperial",
			},
			want: "You shouldn't be outside!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWindDescription(tt.args.speed, tt.args.reqUnits); got != tt.want {
				t.Errorf("getWindDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}
