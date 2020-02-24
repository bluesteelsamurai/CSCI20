/*
Lab 4 Arithmetic
Programmer Name: Nigel S Adams
Date Completed: 9/10/19
*/

package main

import "fmt"


func main(){


var Name string =""
var xvar int =0
var yvar int =0
var flovar1 float32 =0.0
var flovar2 float32 =0.0
var uir string =""
var flotorint bool = false
var inp int =0

fmt.Printf("...Hello.\n")
fmt.Printf("What is your name?\n")
fmt.Scanln(&Name)
fmt.Printf("%s, lets do some math,\n", Name)
fmt.Printf("select your numbers to work with.\n")
fmt.Printf("1)Whole Numbers.\n 2)Decimal Numbers\n")
fmt.Scanln(&inp)
if inp == 1{ flotorint=false}
if inp == 2{flotorint=true}
if flotorint==false{
fmt.Printf("please offer a whole #.\n")
fmt.Scanln(&xvar)
fmt.Printf("please offer another whole #\n")
fmt.Scanln(&yvar)
fmt.Printf("you entered %d, and %d , is this correct?\n", xvar,yvar)
fmt.Scanln(&uir)
 if uir =="N"||uir=="n"{
   fmt.Printf("...Too bad.\n")
   fmt.Scanln()
 }
fmt.Printf("your supplied whole values as their arithmetic results are as such:\n")
fmt.Printf("Addition: %d\n",xvar+yvar)
fmt.Printf("Subtraction, both ways: %d, %d\n",xvar-yvar, yvar-xvar)
fmt.Printf("Multiplication: %d\n",xvar*yvar)
fmt.Printf("Division both ways: %d, %d\n",xvar/yvar, yvar/xvar)
fmt.Printf("Modulus remainders both ways: %d, %d",xvar%yvar, yvar%xvar)
}
if flotorint==true{
  fmt.Printf("please offer a decimal #\n")
  fmt.Scanln(&flovar1)
  fmt.Printf("please offer another decimal#\n")
  fmt.Scanln(&flovar2)
  fmt.Printf("you entered %.2f, and %.2f , is this correct?\n", flovar1,flovar2)
  fmt.Scanln(&uir)
   if uir =="N"||uir=="n"{
     fmt.Printf("...Too bad.\n")
     fmt.Scanln()
   }
  fmt.Printf("your supplied Decimal values as their arithmetic results are as such:\n")
  fmt.Printf("Addition: %.2f\n",flovar1+flovar2)
  fmt.Printf("Subtraction, both ways: %.2f, %.2f\n",flovar1-flovar2, flovar1-flovar2)
  fmt.Printf("Multiplication: %.2f\n",flovar1*flovar2)
  fmt.Printf("Division both ways: %.2f, %.2f\n",flovar1/flovar2, flovar2/flovar1)



}

// declare a string variable to represent a name##
// prompt for and read in a value for the name##

// print a message: "Hello, NAME!"##
// (fill in NAME with your variable)##

// declare two integer variables##
// prompt for and read in values for both integer variables##

// print a message: "You entered X and Y."##
// (fill in X/Y with your variables)##

// compute and display the results of applying all arithmetic operators
// to the two integer variables (+ - * / %)
// print messages showing each result (e.g., "X + Y = Z")
// (fill in X/Y/Z with your variables and the computed result)

// declare two float32 variables
// prompt for and read in values for both float variables

// print a message: "You entered A and B."
// (fill in A/B with your variables)

// compute and display the results of applying all arithmetic operators
// to the two float variables (+ - * /)
// print messages showing each result (e.g., "A + B = C")
// (fill in A/B/C with your variables and the computed result)



}
