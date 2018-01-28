package space

type Planet string

const (
	secondsInDay     float64 = 24 * 60 * 60
	earthOrbitDays   float64 = 365.25
	earthSecondsYear float64 = 31557600
	mercuryOrbitDays float64 = 0.2408467
	venusOrbitDays   float64 = 0.61519726
	marsOrbitDays    float64 = 1.8808158
	jupiterOrbitDays float64 = 11.862615
	saturnOrbitDays  float64 = 29.447498
	uranusOrbitDays  float64 = 84.016846
	neptuneOrbitDays float64 = 164.79132
)

func Age(seconds float64, planet Planet) float64 {
	switch planet {
	case "Earth":
		return seconds / earthSecondsYear
	case "Mercury":
		return seconds / calculate(mercuryOrbitDays)
	case "Venus":
		return seconds / calculate(venusOrbitDays)
	case "Mars":
		return seconds / calculate(marsOrbitDays)
	case "Jupiter":
		return seconds / calculate(jupiterOrbitDays)
	case "Saturn":
		return seconds / calculate(saturnOrbitDays)
	case "Uranus":
		return seconds / calculate(uranusOrbitDays)
	case "Neptune":
		return seconds / calculate(neptuneOrbitDays)
	}
	return 0.0

}

func calculate(planetOrbitDays float64) float64 {
	planetOrbitPeriod := earthOrbitDays * planetOrbitDays
	orbitSeconds := planetOrbitPeriod * secondsInDay
	return orbitSeconds
}
