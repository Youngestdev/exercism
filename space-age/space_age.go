package space

type Planet string

const earthRevolution = 31557600

func Age(seconds float64, planet Planet) float64 {
	Planets := map[Planet]float64{
		"Earth":   earthRevolution,
		"Mercury": earthRevolution * 0.2408467,
		"Venus":   earthRevolution * 0.61519726,
		"Mars":    earthRevolution * 1.8808158,
		"Jupiter": earthRevolution * 11.862615,
		"Saturn":  earthRevolution * 29.447498,
		"Uranus":  earthRevolution * 84.016846,
		"Neptune": earthRevolution * 164.79132,
	}
	return getAge(Planets[planet], seconds)
}

func getAge(rev float64, secs float64) float64 {
	return secs / rev
}
