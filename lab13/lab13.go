// Lab 13 -- maps
//
// Programmer name: Nigel S Adams
// Date completed:  11/19/19

package main

import (
	f "fmt"
	s "sort"
)

// StringIntMap explores usage of map with string keys
// and integer values.
func StringIntMap() {
	// NOTE: "pretty print" as much as possible
	items := [5]string{"apples", "sandwiches", "chips", "beverages", "cookies"}
	// 1. Use the make function to create a map with string keys and integer values
	SIM := make(map[string]int)
	//    NOTE: the map represents a shopping list of items (string) and quantities
	//          to purchase (integer)
	// 2. Use %v to print the map in its default format
	f.Printf("%v\n", SIM)

	// 3a. Use a "range" loop over items to populate the map with user input
	//     NOTE: each string from items will be used as a key for the map
	// 3b. Declare an integer variable
	// 3c. Prompt for and read in a user value (e.g., "How many apples? ")
	// 3d. Assign the user input to the map
	for i:=range items{
		amt:=0
		f.Printf("How Many %v:",items[i])
		f.Scanln(&amt)
		SIM[items[i]]=amt
	}
	// 4. Use %v to print the slice in its default format
	f.Printf("%v",SIM)
}

// StringFloatMap explores usage of map with string keys
// and float values.
func StringFloatMap() {
	// NOTE: "pretty print" as much as possible

	// 1. Use the make function to create a map with string keys and float values
	//    NOTE: the map represents a list of months (string) and their average
	//          precipitation (float) for the city of Chico
	SFM := make(map[string]float64)
	// 2. Use %v to print the map in its default format
	f.Printf("%v",SFM)
	// https://www.usclimatedata.com/climate/chico/california/united-states/usca0211
	//
	// 3. Using the data provided at the link above, enter the correct value for each
	//    month into the map (the month will be the key, the average precipitation will
	//    be the value)
SFM["January"]=4.84
SFM["February"]=4.41
SFM["March"]=4.29
SFM["April"]=1.73
SFM["May"]=1.02
SFM["June"]=0.47
SFM["July"]=0.04
SFM["August"]=0.08
SFM["September"]=0.43
SFM["October"]=1.42
SFM["November"]=3.27
SFM["December"]=4.61
	// 4. Use a "range" loop to display all of the months with their average precipitation
	//    values, in a formatted way (e.g., "Average January rainfall: 4.84 inches")

for i,obj:= range SFM{
	f.Printf("Average precipitation in %v:%v\n",SFM[i],obj)
}
}
// IntFloatMap explores usage of map with integer keys
// and float values. CHALLENGING.
func IntFloatMap() {
	// NOTE: "pretty print" as much as possible

	// Annual CPI averages from 2000-2018
	// https://inflationdata.com/Inflation/Consumer_Price_Index/HistoricalCPI.aspx
	cpiAverages := [19]float64{
		172.2, 177.1, 179.88, 183.96, 188.9, 195.3, 201.6, 207.342,
		215.303, 214.537, 218.056, 224.939, 229.594, 232.957,
		236.736, 237.017, 240.008, 245.12, 251.107,
	}

	// 1. Use the make function to create a map with integer keys and float values
	//    NOTE: the map represents a list of years (integer) and their Consumer Price
	//          Indices (float) for the USA for the years 2000-2018
	cpi:=make(map[int]float64)
	// 2. Use a "range" loop over cpiAverages to populate the map with year (key) and
	//    CPI (value); cpiAverages has CPI values in order from 2000-2018
	for i,obj :=range cpiAverages{
		cpi[i+2000]=obj
	}
	// 3. Use %v to print the map in its default format
	f.Printf("%v",cpi)
	// NOTE: these steps are needed to sort the values by year (maps are not by default
	//       sorted by their keys)
	// 4. Create an empty slice of integers to store years
		cpiYears:=make([]int,len(cpiAverages))
		// 5. Use a "range" loop to get the years from the map and append them to the newly
	//     created years slice
	for i:=range cpiAverages{
		cpiYears=append(cpiYears,int(cpi[i]))
	}
	// 6. Use sorts.Ints to sort the slice containing the years taken from the map
	//    (will require an import of "sort")
	s.Ints(cpiYears)
	// 7. Use a "range" loop over the years slice created in the previous step to
	//    key the map and display the CPIs for years when a CPI decreased compared
	//    to the prior year
	prioryear :=1999
	for i,year := range cpiYears{
		if cpiYears[i]<prioryear{
			f.Printf("%d -->%.2f\n",year,cpi[i])
		}
	}
}

// IntBoolMap explores the usage of map with integer keys
// and boolean values. CHALLENGING.
func IntBoolMap() {
	// NOTE: "pretty print" as much as possible
	numbers := [25]int{
		6992, 7224, 2572, 8290, 5021, 6631, 2690, 5462, 9970, 9933,
		8321, 7748, 2153, 5866, 1729, 6743, 9861, 5703, 2668, 9978,
		8148, 1821, 1963, 7999, 9440,
	}
	// 1. Use the make function to create a map with integer keys and boolean values
	//    NOTE: the map represents a list of numbers (integer) and a boolean value
	//          that flags the key as prime (true) or not prime (false)
	IBM:=make(map[int]bool)
	// 2a. Use a "range" loop over numbers to populate the map (the numbers will be the
	//     keys)
	// 2b. Use a nested "counting" loop to test the number to see if it is prime -- if
	//     prime, set the value for that number/key to true, otherwise set the value
	//     for that number/key to false

	for _,num:= range numbers{
	var factors int = 0
					for i:=1;i<=num;i++{
					if num%i==0{
						factors++
					}
				}

		   if factors>=2{
				IBM[num]=false
			 }else{
				IBM[num]=true
			 }
	}
	// https://www.random.org/integers/


	// 3. Use %v to print the map in its default format
f.Printf("%v",IBM)
	// 4. Use a "range" loop over the map to print only the prime numbers
	//    (the numbers/keys whose values are true)
	for num,prime := range IBM{
		if prime==true{
			f.Printf("%v:is prime: %v\n",num,prime)
			}else{
				f.Printf("%v:not prime:%v\n",num,prime)
			}
	}
}

func main() {

	StringIntMap()
	f.Printf("\n")

	StringFloatMap()
	f.Printf("\n")

	IntFloatMap()
	f.Printf("\n")

	IntBoolMap()
	f.Printf("\n")
}
