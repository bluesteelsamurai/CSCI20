/*
Programmer Name: Nigel S Adams
Date:9/10/19
Project Name: Lab 3
Description:
*/
package main

import (
	"fmt"
)

func main() {

	dow := ""
	mo := ""
	day := 0
	year := 0000

	// declare a string variable to represent the day of the week##
	// prompt for and read in a value for the day of the week##
	fmt.Printf("What day is it?\n")
	fmt.Scanln(&dow)
	// declare a string variable to represent the month##
	// prompt for and read in a value for the month##
	fmt.Printf("What is the month?\n")
	fmt.Scanln(&mo)
	// declare an integer variable to represent the date##
	// prompt for and read in a value for the date##
	fmt.Printf("what is the day of the month is this?\n")
	fmt.Scanln(&day)
	// declare an integer variable to represent the year##
	// prompt for and read in a value for the year##
	fmt.Printf("what is the year?\n")
	fmt.Scanln(&year)
	// print a message: "Today is DAY, MONTH DATE, YEAR."##
	// (fill in DAY/MONTH/DATE/YEAR with your variables)##
	fmt.Printf("Today is %s, ", dow)
	fmt.Printf(", %s", mo)
	fmt.Printf(", %d", day)
	fmt.Printf(", %d", year)
	// declare a string variable to represent a city##
	// prompt for and read in a value for the city##
	var where string = ""
	fmt.Printf("\n where you based?\n")
	fmt.Scanln(&where)
	// declare a float variable to represent a temperature
	// prompt for and read in a value for the temperature
	var temp float64 = 0.0
	var laytemp string = ""
	fmt.Printf("how hot/cold is it today?(decimals accepted ie 98.5)\n")
	fmt.Scanln(&temp)
	if temp <= 65.0 {
		laytemp = "its pretty dang cold.\n"
	}
	if temp >= 80.0 {
		laytemp = "too damn hot out stay inside!\n"
	} else {
		laytemp = "must be nice out why are you inside bro?!\n"
	}
	// print a message: "The temperature in CITY is currently TEMPERATURE degrees."
	// (fill in CITY/TEMPERATURE with your variables)
	fmt.Printf("Based on your report it is:\n")
	fmt.Printf("location:%s ", where)
	fmt.Printf("\ntemp: %f", temp)
	fmt.Printf("\nadvice: %s", laytemp)

}
