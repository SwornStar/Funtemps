package main

import (
	"flag"
	"fmt"

	"github.com/SwornStar/funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var cels float64
var kelv float64
var out string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&cels, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelv, "K", 0.0, "temperatur i grader kelvin")

	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K - Kelvin")

}

func main() {

	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Println("Bruk følgene oppsett for å bruke koden:")
		fmt.Println("-*enhetskode* *temperatur* -out *enhetskode*")
		fmt.Println("eksempel: -F 32 -out C")
		fmt.Println("du vil da få 32 farenheit ut i celcius")
		fmt.Println("NB! MÅ BRUKE STORE BOKSTAVER (C, F og K)")
		return
	}

	// input output koden som kjøres til conv
	if out == "C" && isFlagPassed("F") {
		fmt.Println(conv.FahrToCelsius(fahr))
	}

	if out == "F" && isFlagPassed("C") {
		fmt.Println(conv.CelsiusToFahr(cels))
	}

	if out == "K" && isFlagPassed("C") {
		fmt.Println(conv.CelsiusToKelvin(cels))
	}

	if out == "C" && isFlagPassed("K") {
		fmt.Println(conv.KelvinToCelsius(kelv))
	}

	if out == "F" && isFlagPassed("K") {
		fmt.Println(conv.KelvinToFahr(kelv))
	}

	if out == "K" && isFlagPassed("F") {
		fmt.Println(conv.FahrToKelvin(fahr))
	}

}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// * Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
