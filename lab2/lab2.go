/*
Author: Nigel Adams
email: Adams357@gmail.com
Description: This is a program for telling the time in golang using the TIME
function built into golang
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	//	p := fmt.Printf("") //calls built-in PrintLn() function from fmt package ...when it works!
	//  t := time.Now()     //calls current time from TIME built-in function...when it works!

	//global variable declarations and such go here.
	T1 := time.Now() //declares my time function as T1.

	var inp int //this is input for the if statement controlling the for loop
	var city string = "Chico"
	var temp float64 = 85.0

	for inp != 2 {
		fmt.Printf("\nWelcome to the Timenator.\n")
		fmt.Printf("Would you like to know what time it is now?\n")
		fmt.Printf("Type in your selection \n 1) Yes \n 2) No \n")
		fmt.Scanln(&inp)
		if inp != 1 && inp != 2 {
			fmt.Print("nice try I want a 1 or 2 guy")
		}
		if inp == 1 {
			fmt.Printf("the time is %v", T1)
			fmt.Printf("\nTomorrow's Temperature will be %v", temp)
			fmt.Printf(" in: %s", city)

		} else {
			fmt.Printf("K see ya next time!")
		}

	}

}
