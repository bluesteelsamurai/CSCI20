// Lab 14 -- structs
//
// Programmer name: Nigel S Adams
// Date completed:  11/25/19

package main

import f "fmt"

// For use in the CalculateWeight function.
const earthGravity = 32.1

// Planet represents a planet in our solar system.
type Planet struct {
	Name            string
	Gravity         float64
	MeanTemperature float64
	NumberOfMoons   int
	// Add fields Name (string), Gravity (float64), MeanTemperature (float64),
	// NumberOfMoons (int)
}

// Describe accepts a planet and "pretty prints" the name, mean temperature,
// and number of moons for the planet.
// (EXAMPLE: "Mars: mean temperate -85.0F, 2 moons")
func (p Planet) Describe() {
	f.Printf("Name:%s\nGravitic-Val:%.2f\nAvgTemp:%.2f\nNumOfMoons:%v\n", p.Name, p.Gravity, p.MeanTemperature, p.NumberOfMoons)
}

// MostMoons accepts a slice of planets. Computes and displays the planet
// with the largest number of moons.
func MostMoons(planets []Planet) {
	mostMoon := planets[0]
	for _, planet := range planets {
		if planet.NumberOfMoons > mostMoon.NumberOfMoons {
			mostMoon = planet
		}
	}
	f.Print("This planet has the most moons:\n")
	mostMoon.Describe()

}

// ColdestPlanet accepts a slice of planets. Computes and displays the planet
// that is the coldest.
func ColdestPlanet(planets []Planet) {
	CoPlanet := planets[0]
	for _, planet := range planets {
		if planet.MeanTemperature < CoPlanet.MeanTemperature {
			CoPlanet = planet
		}
	}
	f.Print("This planet is the coldest:\n")
	CoPlanet.Describe()
}

// HottestPlanet accepts a slice of planets. Computes and displays the planet
// that is the hottest.
func HottestPlanet(planets []Planet) {
	HoPlanet := planets[0]
	for _, planet := range planets {
		if planet.MeanTemperature > HoPlanet.MeanTemperature {
			HoPlanet = planet
		}
	}
	f.Print("This planet is the Hottest:\n")
	HoPlanet.Describe()
}

// CalculateWeight accepts a planet and a person's weight on Earth.
// Computes and returns the person's weight on the given planet.
// weightOnOtherPlanet = weightOnEarth * (planetGravity / earthGravity)
func CalculateWeight(p Planet, earthWeight float64) float64 {
	return earthWeight * (p.Gravity / earthGravity)
}

// LeastWeight accepts a slice of planets and a person's weight on Earth.
// Computes and displays the planet on which the person would weigh the least.
// (HINT: use CalculateWeight as a helper)
func LeastWeight(planets []Planet, earthWeight float64) {
	lwPlanet := planets[0]
	for _, planet := range planets {
		if CalculateWeight(planet, earthWeight) < CalculateWeight(lwPlanet, earthWeight) {
			lwPlanet = planet
		}
	}
	f.Printf("you weigh:%.2f\non planet:%s\n", CalculateWeight(lwPlanet, earthWeight), lwPlanet.Name)
}

// MostWeight accepts a slice of planets and a person's weight on Earth.
// Computes and displays the planet on which the person would weigh the most.
// (HINT: use CalculateWeight as a helper)
func MostWeight(planets []Planet, earthWeight float64) {
	MwPlanet := planets[0]
	for _, planet := range planets {
		if CalculateWeight(planet, earthWeight) > CalculateWeight(MwPlanet, earthWeight) {
			MwPlanet = planet
		}
	}
	f.Printf("u weigh:%.2f\nOn planet:%s\n\n", CalculateWeight(MwPlanet, earthWeight), MwPlanet.Name)
}

func main() {
	// Create a slice of Planets containing data for Mercury, Venus, Mars, Jupiter,
	// Saturn, Uranus, Neptune, Pluto
	// Data source: https://nssdc.gsfc.nasa.gov/planetary/factsheet/planet_table_british.html
	planets := []Planet{
		Planet{Name: "Mercury", Gravity: 29.1, MeanTemperature: 333, NumberOfMoons: 0},
		Planet{Name: "Venus", Gravity: 29.1, MeanTemperature: 867, NumberOfMoons: 0},
		Planet{Name: "Earth", Gravity: earthGravity, MeanTemperature: 59, NumberOfMoons: 1},
		Planet{Name: "Mars", Gravity: 12.1, MeanTemperature: -85, NumberOfMoons: 2},
		Planet{Name: "Jupiter", Gravity: 75.9, MeanTemperature: -166, NumberOfMoons: 79},
		Planet{Name: "Saturn", Gravity: 29.4, MeanTemperature: -220, NumberOfMoons: 82},
		Planet{Name: "Uranus", Gravity: 28.5, MeanTemperature: -320, NumberOfMoons: 27},
		Planet{Name: "Neptune", Gravity: 36.0, MeanTemperature: -330, NumberOfMoons: 14},
		Planet{Name: "Pluto", Gravity: 2.3, MeanTemperature: -375, NumberOfMoons: 5},
	}
	// Use a "range" loop to Describe each of the planets
	for _, p := range planets {
		p.Describe()
	}
	f.Printf("\n")
	// call HottestPlanet -- pass planets
	HottestPlanet(planets)
	f.Printf("\n")
	// call ColdestPlanet -- pass planets
	ColdestPlanet(planets)
	f.Printf("\n")
	// call MostMoons -- pass planets
	MostMoons(planets)
	f.Printf("\n")
	// prompt for and read in a person's weight in lbs
	f.Print("Please enter a weight.\n")
	input := 0.0
	f.Scanln(&input)
	for i, planet := range planets {
		f.Printf("you weigh this much on %v:%.2f lbs\n ", planets[i].Name, CalculateWeight(planet, input))
	}
	f.Printf("\n")
	// call LeastWeight -- pass planets and person's weight
	LeastWeight(planets, input)
	f.Printf("\n")
	// call MostWeight -- pass planets and person's weight
	MostWeight(planets, input)
}
