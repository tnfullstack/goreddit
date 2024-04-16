package tempconv

// CToF converts a Celsius temperature to Farenheit.
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts a Farenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// KToC converts a Kevin temperature to Celsius.
func KToC(k Kevin) Celsius {
	return Celsius(k - 273.15)
}
