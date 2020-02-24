package main
/*
import (
	"math"
	)
*/

// MilesPerGallon computes and returns gas mileage.
func MilesPerGallon(milesDriven, gallonsUsed float64) float64 {
	var Mpg float64 = 0//private scope MilesPerGallon float64 variable to return with
	Mpg = milesDriven/gallonsUsed
	return Mpg
}

// TotalWithTip computes and returns a check total including a tip.
// Gratuity is expected to be a value between 0 and 100%.
func TotalWithTip(total, gratuity float64) float64 {
	var TwT float64 = 0//private scope variable short for total with tip
	TwT = (total *gratuity/100)+total//this makes calls be ideally whole #'s and not use low #'s like 3 which become practically unuseable
	return TwT
}

// MaxCanPurchase computes and returns the maximum (whole) number of
// an item that can be purchased with a given amount of money at a
// given price per item.
func MaxCanPurchase(amountOfMoney, pricePerItem float64) int {
  return int(amountOfMoney/pricePerItem)
}

// ConvertToCelsius converts a Fahrenheit temperature to Celsius
// and returns the Celsius value.
func ConvertToCelsius(fahrenheit float64) float64 {
  var CtC float64 = 0.0
	CtC = ((fahrenheit-32.0)*5.0/9.0)
	return CtC
}

// ConvertToFahrenheit converts a Celsius temperature to Fahrenheit
// and returns the Fahrenheit value.
func ConvertToFahrenheit(celsius float64) float64 {
	var CtF float64 = 0.0
	CtF = ((celsius*9.0/5.0)+32.0)
	return CtF
}

/*
func TestMilesPerGallon(t *testing.T) {
	mileage := MilesPerGallon(350.0, 15.0)
	roundedMileage := math.Floor(mileage*100) / 100
	if roundedMileage != 23.33 {
		t.Error("Expected 23.33, got ", roundedMileage)
	}
}

func TestTotalWithTip(t *testing.T) {
	total := TotalWithTip(25.00, 15.00)
	if total != 28.75 {
		t.Error("Expected 28.75, got ", total)
	}
}

func TestConvertToFahrenheit(t *testing.T) {
	var freezing, boiling float64

	freezing = ConvertToFahrenheit(0.0)
	if freezing != 32.0 {
		t.Error("[C2F Freezing] expected 32.0, got ", freezing)
	}

	boiling = ConvertToFahrenheit(100.0)
	if boiling != 212.0 {
		t.Error("[C2F Boiling] expected 212.0, got ", freezing)
	}
}

func TestConvertToCelsius(t *testing.T) {
	var freezing, boiling float64

	freezing = ConvertToCelsius(32.0)
	if freezing != 0.0 {
		t.Error("[F2C Freezing] expected 0.0, got ", freezing)
	}

	boiling = ConvertToCelsius(212.0)
	if boiling != 100.0 {
		t.Error("[F2C Boiling] expected 100.0, got ", freezing)
	}
}

func TestMaxCanPurchase(t *testing.T) {
	var max int

	max = MaxCanPurchase(20, 1.99)

	if max != 10 {
		t.Error("Expected 10, got ", max)
	}
}
*/


func main() {


}
