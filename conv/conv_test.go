package conv

import (
	"testing"

	"github.com/SwornStar/funtemps/conv"
)

func TestCelsiusToFahr(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{0, 32},
		{100, 212},
		{-40, -40},
	}

	for _, test := range tests {
		got := conv.CelsiusToFahrtest.input)
		if got != test.want {
			t.Errorf("CelsiusToFahr(%v) = %v; want %v", test.input, got, test.want)
		}
	}
}

func TestFahrToKelvin(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{32, 273.15},
		{212, 373.15},
		{-40, 233.15},
	}

	for _, test := range tests {
		got := conv.FahrToKelvin(test.input)
		if got != test.want {
			t.Errorf("FahrToKelvin(%v) = %v; want %v", test.input, got, test.want)
		}
	}
}

func TestKelvinToCelsius(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{273.15, 0},
		{373.15, 100},
		{233.15, -40},
	}

	for _, test := range tests {
		got := conv.KelvinToCelsius(test.input)
		if got != test.want {
			t.Errorf("KelvinToCelsius(%v) = %v; want %v", test.input, got, test.want)
		}
	}
}

func TestFahrToCelsius(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{32, 0},
		{212, 100},
		{-40, -40},
	}

	for _, test := range tests {
		got := conv.FahrToCelsius(test.input)
		if got != test.want {
			t.Errorf("FahrToCelsius(%v) = %v; want %v", test.input, got, test.want)
		}
	}
}




func TestKelvinToFahr(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{273.15, 32},
		{373.15, 212},
		{233.15, -40},
	}

	for _, test := range tests {
		got := conv.KelvinToFahrenheit(test.input)
		if got != test.want {
			t.Errorf("KelvinToFahr(%v) = %v; want %v", test.input, got, test.want)
		}
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{0, 273.15},
		{100, 373.15},
		{-40, 233.15},
	}

	for _, test := range tests {
		got := conv.CelsiusToKelvin(test.input)
		if got != test.want {
			t.Errorf("CelsiusToKelvin(%v) = %v; want %v", test.input, got, test.want)
		}
	}
}
