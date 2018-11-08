package space

type Planet string

const earth_years_seconds = 31557600.0

// orbital period in earth years
var planetsMap = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// return age according planet
func Age(seconds float64, planet Planet) float64 {
	planet_years := seconds / (earth_years_seconds * planetsMap[planet])
	return planet_years
}
