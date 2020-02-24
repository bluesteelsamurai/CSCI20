// Lab 9 -- control flow (for)
//
// Programmer name: Nigel Adams
// Date completed:  10/17/19

package main

import (
	"fmt"
)

var max int = 0
var total, evens, odds, fizz, buzz, fizzBuzz int64 = 0, 2, 1, 0, 0, 0

// PrintEvens prints the even numbers from 1 up to max.
func PrintEvens(max int) {

	fmt.Printf("This is the evens counting up from 1...\n")
	for i := 2; i <= max; i = i + 2 { //counts by twos and prints value of i in the for loop as the even value 8-)
		if i%7 == 1 {
			fmt.Printf("\n")
		}
		fmt.Printf(" %v ", i)
		total++
	}
	fmt.Printf("\nThere are %v even numbers ranging from 2 to %v\n", total, max)
}

// PrintOdds prints the odd numbers from 1 up to max.
func PrintOdds(max int) {
	total = 0

	fmt.Printf("\nThis is the odds:\n")
	for i := 1; i <= max; i = i + 2 { //similar to PrintEvens, counts by twos from an odd starting point "1" counting each repetition for the total # of odd #s
		if i%8 == 1 {
			fmt.Printf("\n")
		}
		fmt.Printf(" %v ", i)
		total++
	}
	fmt.Printf("\nThere are %v Odd numbers ranging from 1 to %v\n", total, max)
}

// FizzBuzz solves the Fizz Buzz challenge from 1 up to max.
// https://leetcode.com/problems/fizz-buzz/description/
func FizzBuzz(max int) {

	fmt.Printf("Now for the FizzBuzz Challenge\n ")
	for i := 1; i < max+1; i++ {
		if i%8 == 1 {
			fmt.Printf("\n")
		}
		if i%15 == 0 {
			fmt.Printf(" FizzBuzz ")
			fizzBuzz++
		} else if i%5 == 0 {
			fmt.Printf(" Buzz ")
			buzz++
		} else if i%3 == 0 {
			fmt.Printf(" Fizz ")
			fizz++
		} else {
			fmt.Printf(" %v ", i)
		}

	}

	fmt.Printf("\n\n There are : %v, Fizzers;\n There are %v, Buzzers;\n And : %v, FizzBuzzers", fizz, buzz, fizzBuzz)
}

func main() {
	fmt.Printf("Hello please pick a max # and you will get a list of evens and odds\n WARNING!!! Large Values will take a while!\n")
	fmt.Scanln(&max)
	// call PrintEvens with any value for max
	PrintEvens(max)
	// call PrintOdds with any value for max
	PrintOdds(max)
	// call FizzBuzz with any value for max
	FizzBuzz(max)
	fmt.Scanln()
}
