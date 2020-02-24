/*
title: Project 1 - Change Maker
creator: Nigel S Adams
Date created: 9/19/2019
Description:This is a program using arithmetic operators and logic gates to prompt the user
	    for data involving US currency, logic will operate under the assumption that the
	    currency is base 10 and change made from dollars will break down by quarters
	    first then dimes then nickels etc.
*/

package main

import (
	"fmt"
)

func Handler() {

	var Menusel int = 0
	fmt.Printf("Welcome to the Changinator\n")
	fmt.Printf("Please make a selection\n")
	fmt.Printf("1-Change to Dollars-\n2-Dollars to Change-\n")
	fmt.Scanln(&Menusel)
	switch Menusel {
	case 1:
		ChangeToDollars()
	case 2:
		CentsToChange()
	default:
		fmt.Printf("OOOO make a proper selection ya rule breaker bye now!")
	}

}

/*this makes the quarter and dimes into whole dollar amounts*/
func ChangeToDollars() {
	Quarts, Dims, Nicks, Pens := 0.0, 0.0, 0.0, 0.0

	fmt.Printf("How many Quarters do you have?\n")
	fmt.Scanln(&Quarts)
	fmt.Printf("How many Dimes?\n")
	fmt.Scanln(&Dims)
	fmt.Printf("How many Nickels?\n")
	fmt.Scanln(&Nicks)
	fmt.Printf("How many Pennies?\n")
	fmt.Scanln(&Pens)
	var Dols float64 = (Quarts / 4) + (Dims / 10) + (Nicks / 20) + (Pens / 100)
	fmt.Printf("You have $%.2f ", Dols)

}

/*this makes the dollars into change */
func CentsToChange() {
	Cents := 0
	Nickels := 0
	Dimes := 0
	Pennies := 0

	fmt.Printf("How many cents do you have?\n")
	fmt.Scanln(&Cents)
	fmt.Printf("you have %d quarters and ", Cents/25)
	Dimes = Cents % 25
	fmt.Printf("%d dimes and ", Dimes/10)
	Nickels = Dimes % 5
	fmt.Printf("%d nickels and ", Nickels/5)
	Pennies = Nickels % 5
	fmt.Printf("%d pennies.", Pennies)
	fmt.Scanln()
}
func main() {

	Handler()

}
