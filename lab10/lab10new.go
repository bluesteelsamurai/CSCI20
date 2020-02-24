// Lab 10 -- arrays and range
//
// Programmer name: __________
// Date completed:  __________

package main

import (
	"fmt"
	"time"
)

// IntArray explores usage of an integer array.
func IntArray() {
	// NOTE: "pretty print" as much as possible

	/* 1. Declare an array of 10 integers named "numbers" in its zero state
	2. Use %v to print the array in its default format
	3. Print the length of the array
	4. Use array indexing to individually populate
	   the array with the values 10 (index 0), 20, 30, 40
	   50, 60, 70, 80, 90, 100
	5. Use %v to print the updated array in its default format */
	var numbers [10]int
	fmt.Printf("Empty array of 10:%v\n", numbers)
	time.Sleep(2 * time.Second)
	fmt.Printf("Length of array:%v\n", len(numbers))
	time.Sleep(2 * time.Second)

	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50
	numbers[5] = 60
	numbers[6] = 70
	numbers[7] = 80
	numbers[8] = 90
	numbers[9] = 100

	fmt.Printf("Values in array: %v\n", numbers)

}

// FloatArray explores usage of a float array.
func FloatArray() {
	/* NOTE: "pretty print" as much as possible
	   1. Declare an array of 5 float64 named "scores" in its zero state
	   2. Use a "counting" loop to prompt for and read in values for the array
	   3. Use %v to print the array in its default format
	   4. Use a "range" loop (ignoring index) to compute the sum of the scores
	   5. Print the sum of scores
	   6. Compute the average of scores
	   7. Print the average of scores*/
	var scores [5]float64
	var sum float64 = 0.0

	for i := 0; i < len(scores); i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("enter score for applicant %v:\n", i+1)
		fmt.Scanln(&scores[i])
		fmt.Printf("\n")
	}
	for score := range scores{
		sum +=scores[int(score)]
	}

	fmt.Printf("sum: %.1f\n", sum)
	average := sum / float64(len(scores))
	fmt.Printf("Average: %.1f\n", average)


}

// StringArray explores usage of a string array.
func StringArray() {
	/* NOTE: "pretty print" as much as possible
	 1. Declare an array of 3 string named "greetings" -- pre-populate
	    the array with any 3 greetings strings of your choice (EX: "Howdy")
	 2. Use a "range" loop (including index) to print each greeting, preceded
	    by its position/index in the array
	    EXAMPLE:     [0] Howdy
	                 [1] Hello
	                 [2] Hi
	     BETTER:     [1] Howdy
	                 [2] Hello
					 [3] Hi */
	var greetings [3]string
	greetings[0] = "Guten Morgen!"
	greetings[1] = "Salutations!"
	greetings[2] = "Maj Po Ghargh"

	for i, greeting := range greetings {
		fmt.Printf("[%v]:%v\n",i, greeting)
	}
	fmt.Scanln()
}

	/* ParallelArrays explores usage of arrays "in parallel"
	 (i.e., the arrays have related data).*/

func ParallelArrays() {
	/* NOTE: "pretty print" as much as possible
	 1. Declare an array of 5 string named "products" -- pre-populate
	    the array with "bread", "milk", "eggs", "butter", "sugar"
	 2. Declare an array of 5 float64 named "prices" -- pre-populate
	    the array with reasonable prices for each item; prices should
	    "parallel" the products (i.e., prices[0] is the price for products[0])
	    (Sample prices: http://www.thepeoplehistory.com/pricebasket.html)
	 3. Use a "counting loop" to print the products with their prices
	    EXAMPLE:      bread: $1.98
	                   milk: $3.22
	                   eggs: $0.68
	                 butter: $2.68
	                  sugar: $1.43*/
	products := []string {"loaf of bread", "1 gallon milk", "1 dozen eggs", "butter, 1lb ", "sugar, 4lbs"}
	prices := []float64 {1.99,3.49,1.99,2.49,3.99}
	fmt.Printf("The Store:\n")
	for i,_ := range products{
		fmt.Printf("%v: $%.2f\n",products[i],prices[i])
	}
	fmt.Printf("ERROR! USER NOT RECOGNIZED PRESS [ENTER] TO QUIT.\nJK!!!!!!!!!!!hahahahahaha!\n")
	fmt.Scanln()
}

func main() {

	// You do not need to change anything in main.

	IntArray()
	fmt.Printf("\n")

	FloatArray()
	fmt.Printf("\n")

	StringArray()
	fmt.Printf("\n")

	ParallelArrays()
	fmt.Printf("\n")
}
