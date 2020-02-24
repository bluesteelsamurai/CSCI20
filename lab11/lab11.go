// Lab 11 -- slices
//
// Programmer name:Nigel S Adams
// Date completed:11/5/19

package main

import (
	"fmt"
	"strings"
)

// IntSlice explores usage of an integer slice.
func IntSlice() {
	/* NOTE: "pretty print" as much as possible
	   1. Use the make function to create a slice of 0 integers named "numbers"
	   2. Use %v to print the slice in its default format
	   3. Print the length of the slice
	   4a. Use a "counting" loop to prompt for and read in 5 values for "numbers"
	   4b. Use the append function to add the input values into "numbers"
	   5. Use %v to print the slice in its default format
	   6. Print the length of the slice*/
	numbers := make([]int, 0)
	fmt.Printf("Slice in default: %v\n", numbers)
	fmt.Printf("Length of slice: %v\n", len(numbers))
	for i := 0; i < 5; i++ {
		pInput := 0
		fmt.Printf("Please enter an integer value.\n")
		fmt.Scanln(&pInput)
		numbers = append(numbers, pInput)
	}
	fmt.Printf("slice default: %v\n", numbers)
	fmt.Printf("the length: %v\n", len(numbers))
	fmt.Scanln()
}
// FloatSlice explores usage of a float slice.
func FloatSlice() {
	/* NOTE: "pretty print" as much as possible
	   1. Use the make function to create a slice of 5 float64s named "temperatures"
	   2. Use %v to print the slice in its default format
	   3. Use a "range" loop to prompt for and read in 5 values for "temperatures"
	   4. Use a "range" loop (ignoring index) to compute the sum of the temperatures
	   5. Print the sum of temperatures
	   6. Compute the average of temperatures
	   7. Print the average of temperatures*/
	var sum float64 = 0.0
	temperatures := make([]float64, 5)
	fmt.Printf("Default slice: %v", temperatures)
	for i,_ := range temperatures {
		fmt.Printf("Please enter %v temperature down to 1/10th degree accuracy.\n", i+1)
		fmt.Scanln(&temperatures[i])
	}
	for _,temp := range temperatures {
		sum += temp
		}
	fmt.Printf("The Sum: %v\n", sum)
	average := sum / float64(len(temperatures))
	fmt.Printf("the Average: %v", average)
	fmt.Scanln()
}
// StringSlice explores usage of a string slice.
func StringSlice() {
	/* NOTE: "pretty print" as much as possible
	   1. Declare an array of 3 string named "originalGreetings" -- pre-populate
	      the array with any 3 greetings strings of your choice (EX: "Howdy")
	   2. Use the make function to create a slice of strings, of equal length to
	      "originalGreetings", name the slice "updatedGreetings"
	   3. Use %v to print the slice "updatedGreetings" in its default format
	   4. Use a "range" loop to overwrite each string in "updatedGreetings"
	      with its upper-case form (HINT: use strings.ToUpper)
	   5. Use %v to print the slice "updatedGreetings" in its default format*/
	 originalGreetings := [3]string{"guten tag.","buenos tardes.","'sup"}
	 updatedGreetings := make([]string,len(originalGreetings))
		fmt.Printf("%v",updatedGreetings)
		for i,_ := range updatedGreetings{
			updatedGreetings[i]=strings.ToUpper(originalGreetings[i])
				}
			fmt.Printf("%v",updatedGreetings)

	}
// SliceExpressions explores the usage of slice expressions.
func SliceExpressions() {

/*1. Declare an array of 5 string named "products" -- pre-populate
	   the array with "bread", "milk", "eggs", "butter", "sugar"
	2. Use %v and a slice expression to print the first two strings in "products"
	   (should be "bread" and "milk")
	3. Use %v and a slice expression to print the last two strings in "products"
	   (should be "butter" and "sugar")
	4. Use %v and a slice expression to print middle string in "products"
	   (should be "eggs")*/
		 products:=[5]string{"bread","milk","eggs","butter","sugar"}
		 fmt.Printf("%v\n",products[:2])
		 fmt.Printf("%v\n",products[3:])
		 fmt.Printf("[%v]\n",products[2])
		 fmt.Scanln()
}

func main() {

	IntSlice()
	fmt.Printf("\n")

	FloatSlice()
	fmt.Printf("\n")

	StringSlice()
	fmt.Printf("\n")

	SliceExpressions()
	fmt.Printf("\n")
}
