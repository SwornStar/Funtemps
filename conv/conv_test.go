package conv

import (
	"testing"

	"github.com/SwornStar/funtemps/conv"
)

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{0, 32},
		{100, 212},
		{-40, -40},
	}

	for _, test := range tests {
		got := conv.CelsiusToFahrenheit(test.input)
		if got != test.want {
			t.Errorf("CelsiusToFahrenheit(%v) = %v; want %v", test.input, got, test.want)
		}
	}
}

func TestFahrenheitToKelvin(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{32, 273.15},
		{212, 373.15},
		{-40, 233.15},
	}

	for _, test := range tests {
		got := conv.FahrenheitToKelvin(test.input)
		if got != test.want {
			t.Errorf("FahrenheitToKelvin(%v) = %v; want %v", test.input, got, test.want)
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

func TestFahrenheitToCelsius(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{32, 0},
		{212, 100},
		{-40, -40},
	}

	for _, test := range tests {
		got := conv.FahrenheitToCelsius(test.input)
		if got != test.want {
			t.Errorf("FahrenheitToCelsius(%v) = %v; want %v", test.input, got, test.want)
		}
	}
}

func TestKelvinToFahrenheit(t *testing.T) {
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
			t.Errorf("KelvinToFahrenheit(%v) = %v; want %v", test.input, got, test.want)
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
