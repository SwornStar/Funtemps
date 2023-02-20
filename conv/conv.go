package conv


func CelsiusToFahrenheit(celsius float64) float64 {
    return ((celsius * (9/5)) + 32)
}

func FahrenheitToKelvin(fahr float64) float64 {
    return ((fahr - 32) * 5/9 + 273.15)
}

func KelvinToCelsius(kelvin float64) float64 {
    return (kelvin - 273.15)
}

func FahrenheitToCelsius(fahr float64) float64 {
    return ((fahr - 32) * 5/9)
}

func KelvinToFahrenheit(kelvin float64) float64 {
    return ((kelvin - 273.15) * 9/5 + 32) 
}

func CelsiusToKelvin(celsius float64) float64 {
    return (celsius + 273.15)
}