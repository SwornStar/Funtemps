package conv


func CelsiusToFahr(celsius float64) float64 {
    return ((celsius * (9/5)) + 32)
}

func FahrToKelvin(fahr float64) float64 {
    return ((fahr - 32) * 5/9 + 273.15)
}

func KelvinToCelsius(kelvin float64) float64 {
    return (kelvin - 273.15)
}

func FahrToCelsius(fahr float64) float64 {
    return ((fahr - 32) * 5/9)
}

func KelvinToFahr(kelvin float64) float64 {
    return ((kelvin - 273.15) * 9/5 + 32) 
}

func CelsiusToKelvin(celsius float64) float64 {
    return (celsius + 273.15)
}