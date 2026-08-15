package main

type WeatherCondition int

const (
    Clear WeatherCondition = iota // Ключевое слово iota присваивает каждой константе числовое значение по порядку (0, 1, 2, 3 и т.д.)
    Rain
    HeavyRain
    Snow
)

type WeatherData struct {
    Condition WeatherCondition
    WindSpeed int
}

func GetWeatherMultiplier(weather WeatherData) float64 {
	var coefficient float64 = 1

	switch weather.Condition {
	case 1:
		coefficient += 0.125
	case 2:
		coefficient += 0.2
	case 3:
		coefficient += 0.15 
	}

	if weather.WindSpeed > 15 {
		coefficient += 0.1
	}
	return coefficient
}
