package space

type Planet string

const (
	Earth   Planet = "Earth"
	Mercury        = "Mercury"
	Venus          = "Venus"
	Mars           = "Mars"
	Jupiter        = "Jupiter"
	Saturn         = "Saturn"
	Uranus         = "Uranus"
	Neptune        = "Neptune"
)

const earthAgeSecond = 31557600

var planetPeriods = map[Planet]float64{
	Earth:   1,
	Mercury: 0.2408467,
	Venus:   0.61519726,
	Mars:    1.8808158,
	Jupiter: 11.862615,
	Saturn:  29.447498,
	Uranus:  84.016846,
	Neptune: 164.79132,
}

func Age(second float64, planet Planet) float64 {
	period := planetPeriods[planet] * earthAgeSecond

	return second / period
}
